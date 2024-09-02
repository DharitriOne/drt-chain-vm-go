package vmhooksgenerate

import (
	"fmt"
)

func WriteVMHooksWrapper(out *eiGenWriter, eiMetadata *EIMetadata) {
	out.WriteString(`package executorwrapper

// Code generated by vmhooks generator. DO NOT EDIT.

// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
// !!!!!!!!!!!!!!!!!!!!!! AUTO-GENERATED FILE !!!!!!!!!!!!!!!!!!!!!!
// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

import (
	"fmt"

	"github.com/DharitriOne/drt-chain-vm-go/executor"
)

// WrapperVMHooks wraps a VMHooks instance and optionally performs some logging.
type WrapperVMHooks struct {
	logger         ExecutorLogger
	wrappedVMHooks executor.VMHooks
}
`)

	for _, funcMetadata := range eiMetadata.AllFunctions {
		out.WriteString(fmt.Sprintf("\n// %s VM hook wrapper", upperInitial(funcMetadata.Name)))
		out.WriteString(fmt.Sprintf("\nfunc (w *WrapperVMHooks) %s(", upperInitial(funcMetadata.Name)))
		for argIndex, arg := range funcMetadata.Arguments {
			if argIndex > 0 {
				out.WriteString(", ")
			}
			out.WriteString(fmt.Sprintf("%s %s", arg.Name, vmHooksWrapperType(arg.Type)))
		}
		out.WriteString(")")
		if funcMetadata.Result != nil {
			out.WriteString(fmt.Sprintf(" %s", vmHooksWrapperType(funcMetadata.Result.Type)))
		}
		out.WriteString(" {")
		writeCallInfo(out, funcMetadata)
		out.WriteString("\n\tw.logger.LogVMHookCallBefore(callInfo)")
		out.WriteString("\n\t")
		if funcMetadata.Result != nil {
			out.WriteString("result := ")
		}
		out.WriteString(fmt.Sprintf("w.wrappedVMHooks.%s(", upperInitial(funcMetadata.Name)))
		writeCommaSeparatedArgumentNames(out, funcMetadata.Arguments)
		out.WriteString(")")
		out.WriteString("\n\tw.logger.LogVMHookCallAfter(callInfo)")
		if funcMetadata.Result != nil {
			out.WriteString("\n\treturn result")
		}
		out.WriteString("\n}\n")
	}
}

func writeCallInfo(out *eiGenWriter, funcMetadata *EIFunction) {
	if len(funcMetadata.Arguments) == 0 {
		out.WriteString(fmt.Sprintf("\n\tcallInfo := \"%s()\"", upperInitial(funcMetadata.Name)))
	} else {
		out.WriteString(fmt.Sprintf("\n\tcallInfo := fmt.Sprintf(\"%s(", upperInitial(funcMetadata.Name)))
		for argIndex := range funcMetadata.Arguments {
			if argIndex > 0 {
				out.WriteString(", ")
			}
			out.WriteString("%d")
		}
		out.WriteString(")\", ")
		writeCommaSeparatedArgumentNames(out, funcMetadata.Arguments)
		out.WriteString(")")
	}
}

func writeCommaSeparatedArgumentNames(out *eiGenWriter, arguments []*EIFunctionArg) {
	for argIndex, arg := range arguments {
		if argIndex > 0 {
			out.WriteString(", ")
		}
		out.WriteString(fmt.Sprintf("%s", arg.Name))
	}
}
