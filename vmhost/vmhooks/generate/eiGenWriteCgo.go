package vmhooksgenerate

import (
	"fmt"
)

type cgoWriter struct {
	goPackage string
	cgoPrefix string
}

// C types
func cgoType(eiType EIType) string {
	switch eiType {
	case EITypeMemPtr:
		fallthrough
	case EITypeMemLength:
		fallthrough
	case EITypeInt32:
		return "int32_t"
	case EITypeInt64:
		return "long long"
	default:
		panic("invalid type")
	}
}

// Go types equvalent to the cgo C types
func cgoExportType(eiType EIType) string {
	switch eiType {
	case EITypeMemPtr:
		fallthrough
	case EITypeMemLength:
		fallthrough
	case EITypeInt32:
		return "int32"
	case EITypeInt64:
		return "int64"
	default:
		panic("invalid type")
	}
}

// Go types equvalent to the cgo C types
func cgoExportConversion(arg *EIFunctionArg) string {
	switch arg.Type {
	case EITypeMemPtr:
		return fmt.Sprintf("executor.MemPtr(%s)", arg.Name)
	default:
		return arg.Name
	}
}

// Go types present in the VMHooks interface
func vmHooksType(eiType EIType) string {
	switch eiType {
	case EITypeMemPtr:
		return "MemPtr"
	case EITypeMemLength:
		return "MemLength"
	case EITypeInt32:
		return "int32"
	case EITypeInt64:
		return "int64"
	default:
		panic("invalid type")
	}
}

// Go types present in the VMHooksWrapper
func vmHooksWrapperType(eiType EIType) string {
	switch eiType {
	case EITypeMemPtr:
		return "executor.MemPtr"
	case EITypeMemLength:
		return "executor.MemLength"
	case EITypeInt32:
		return "int32"
	case EITypeInt64:
		return "int64"
	default:
		panic("invalid type")
	}
}

func (writer *cgoWriter) cgoFuncName(funcMetadata *EIFunction) string {
	return writer.cgoPrefix + lowerInitial(funcMetadata.Name)
}

func (writer *cgoWriter) cgoImportName(funcMetadata *EIFunction) string {
	return fmt.Sprintf("C.%s", writer.cgoFuncName(funcMetadata))
}

// WriteWasmer1Cgo writes the metadata in the provided file
func WriteWasmer1Cgo(out *eiGenWriter, eiMetadata *EIMetadata) {
	writer := &cgoWriter{
		goPackage: "wasmer",
		cgoPrefix: "v1_5_",
	}
	writer.writeHeader(out, eiMetadata)
	writer.writeCgoFunctions(out, eiMetadata)
	writer.writePopulateImports(out, eiMetadata)
	writer.writeGoExports(out, eiMetadata)
}

// WriteWasmer2Cgo writes the metadata in the provided file
func WriteWasmer2Cgo(out *eiGenWriter, eiMetadata *EIMetadata) {
	writer := &cgoWriter{
		goPackage: "wasmer2",
		cgoPrefix: "w2_",
	}
	writer.writeHeader(out, eiMetadata)
	writer.writeCgoFunctions(out, eiMetadata)
	writer.writePopulateFuncPointers(out, eiMetadata)
	writer.writeGoExports(out, eiMetadata)
}

func (writer *cgoWriter) writeHeader(out *eiGenWriter, eiMetadata *EIMetadata) {
	autoGeneratedGoHeader(out, writer.goPackage)
	out.WriteString(`
// // Declare the function signatures (see [cgo](https://golang.org/cmd/cgo/)).
//
// #include <stdlib.h>
// typedef int int32_t;
//
`)
}

func (writer *cgoWriter) writeCgoFunctions(out *eiGenWriter, eiMetadata *EIMetadata) {
	for _, funcMetadata := range eiMetadata.AllFunctions {
		out.WriteString(fmt.Sprintf("// extern %-9s %s(void* context",
			externResult(funcMetadata.Result),
			writer.cgoFuncName(funcMetadata),
		))
		for _, arg := range funcMetadata.Arguments {
			out.WriteString(fmt.Sprintf(", %s %s", cgoType(arg.Type), arg.Name))
		}
		out.WriteString(");\n")
	}

	out.WriteString(`import "C"

import (
	"unsafe"

	"github.com/DharitriOne/drt-chain-vm-go/executor"
)

`)
}

func (writer *cgoWriter) writePopulateImports(out *eiGenWriter, eiMetadata *EIMetadata) {
	out.WriteString(`// populateWasmerImports populates imports with the BaseOpsAPI API methods
func populateWasmerImports(imports *wasmerImports) error {
	var err error
`)

	for _, funcMetadata := range eiMetadata.AllFunctions {
		out.WriteString(fmt.Sprintf("\terr = imports.append(\"%s\", %s, %s)\n",
			lowerInitial(funcMetadata.Name),
			writer.cgoFuncName(funcMetadata),
			writer.cgoImportName(funcMetadata),
		))
		out.WriteString("\tif err != nil {\n")
		out.WriteString("\t\treturn err\n")
		out.WriteString("\t}\n\n")
	}
	out.WriteString("\treturn nil\n")
	out.WriteString("}\n")
}

func (writer *cgoWriter) writePopulateFuncPointers(out *eiGenWriter, eiMetadata *EIMetadata) {
	out.WriteString(`// populateCgoFunctionPointers populates imports with the BaseOpsAPI API methods
func populateCgoFunctionPointers() *cWasmerVmHookPointers {
	return &cWasmerVmHookPointers{`)

	for _, funcMetadata := range eiMetadata.AllFunctions {
		out.WriteString(fmt.Sprintf("\n\t\t%s: funcPointer(C.%s),",
			cgoFuncPointerFieldName(funcMetadata),
			writer.cgoFuncName(funcMetadata),
		))
	}
	out.WriteString(`
	}
}
`)
}

func (writer *cgoWriter) writeGoExports(out *eiGenWriter, eiMetadata *EIMetadata) {
	for _, funcMetadata := range eiMetadata.AllFunctions {
		out.WriteString(fmt.Sprintf("\n//export %s\n",
			writer.cgoFuncName(funcMetadata),
		))
		out.WriteString(fmt.Sprintf("func %s(context unsafe.Pointer",
			writer.cgoFuncName(funcMetadata),
		))
		for _, arg := range funcMetadata.Arguments {
			out.WriteString(fmt.Sprintf(", %s %s", arg.Name, cgoExportType(arg.Type)))
		}
		out.WriteString(")")
		if funcMetadata.Result != nil {
			out.WriteString(fmt.Sprintf(" %s", cgoExportType(funcMetadata.Result.Type)))
		}
		out.WriteString(" {\n")
		out.WriteString("\tvmHooks := getVMHooksFromContextRawPtr(context)\n")
		out.WriteString("\t")
		if funcMetadata.Result != nil {
			out.WriteString("return ")
		}
		out.WriteString(fmt.Sprintf("vmHooks.%s(",
			upperInitial(funcMetadata.Name),
		))
		for argIndex, arg := range funcMetadata.Arguments {
			if argIndex > 0 {
				out.WriteString(", ")
			}
			out.WriteString(cgoExportConversion(arg))
		}
		out.WriteString(")\n")

		out.WriteString("}\n")
	}
}