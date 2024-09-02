package vmhooks

import (
	"encoding/hex"
	"errors"

	"github.com/DharitriOne/drt-chain-core-go/core/check"
	"github.com/DharitriOne/drt-chain-vm-common-go/builtInFunctions"

	"github.com/DharitriOne/drt-chain-vm-go/executor"
	"github.com/DharitriOne/drt-chain-vm-go/math"
	"github.com/DharitriOne/drt-chain-vm-go/vmhost"
)

const (
	managedSCAddressName                    = "managedSCAddress"
	managedOwnerAddressName                 = "managedOwnerAddress"
	managedCallerName                       = "managedCaller"
	managedSignalErrorName                  = "managedSignalError"
	managedWriteLogName                     = "managedWriteLog"
	managedMultiTransferDCDTNFTExecuteName  = "managedMultiTransferDCDTNFTExecute"
	managedTransferValueExecuteName         = "managedTransferValueExecute"
	managedExecuteOnDestContextName         = "managedExecuteOnDestContext"
	managedExecuteOnDestContextByCallerName = "managedExecuteOnDestContextByCaller"
	managedExecuteOnSameContextName         = "managedExecuteOnSameContext"
	managedExecuteReadOnlyName              = "managedExecuteReadOnly"
	managedCreateContractName               = "managedCreateContract"
	managedDeployFromSourceContractName     = "managedDeployFromSourceContract"
	managedUpgradeContractName              = "managedUpgradeContract"
	managedUpgradeFromSourceContractName    = "managedUpgradeFromSourceContract"
	managedAsyncCallName                    = "managedAsyncCall"
	managedCreateAsyncCallName              = "managedCreateAsyncCall"
	managedGetCallbackClosure               = "managedGetCallbackClosure"
	managedGetMultiDCDTCallValueName        = "managedGetMultiDCDTCallValue"
	managedGetDCDTBalanceName               = "managedGetDCDTBalance"
	managedGetDCDTTokenDataName             = "managedGetDCDTTokenData"
	managedGetReturnDataName                = "managedGetReturnData"
	managedGetPrevBlockRandomSeedName       = "managedGetPrevBlockRandomSeed"
	managedGetBlockRandomSeedName           = "managedGetBlockRandomSeed"
	managedGetStateRootHashName             = "managedGetStateRootHash"
	managedGetOriginalTxHashName            = "managedGetOriginalTxHash"
	managedIsDCDTFrozenName                 = "managedIsDCDTFrozen"
	managedIsDCDTLimitedTransferName        = "managedIsDCDTLimitedTransfer"
	managedIsDCDTPausedName                 = "managedIsDCDTPaused"
	managedBufferToHexName                  = "managedBufferToHex"
	managedGetCodeMetadataName              = "managedGetCodeMetadata"
	managedIsBuiltinFunction                = "managedIsBuiltinFunction"
)

// ManagedSCAddress VMHooks implementation.
// @autogenerate(VMHooks)
func (context *VMHooksImpl) ManagedSCAddress(destinationHandle int32) {
	managedType := context.GetManagedTypesContext()
	runtime := context.GetRuntimeContext()
	metering := context.GetMeteringContext()

	gasToUse := metering.GasSchedule().BaseOpsAPICost.GetSCAddress
	err := metering.UseGasBoundedAndAddTracedGas(managedSCAddressName, gasToUse)
	if context.WithFault(err, runtime.BaseOpsErrorShouldFailExecution()) {
		return
	}

	scAddress := runtime.GetContextAddress()

	managedType.SetBytes(destinationHandle, scAddress)
}

// ManagedOwnerAddress VMHooks implementation.
// @autogenerate(VMHooks)
func (context *VMHooksImpl) ManagedOwnerAddress(destinationHandle int32) {
	managedType := context.GetManagedTypesContext()
	blockchain := context.GetBlockchainContext()
	runtime := context.GetRuntimeContext()
	metering := context.GetMeteringContext()

	gasToUse := metering.GasSchedule().BaseOpsAPICost.GetOwnerAddress
	err := metering.UseGasBoundedAndAddTracedGas(managedOwnerAddressName, gasToUse)
	if context.WithFault(err, runtime.BaseOpsErrorShouldFailExecution()) {
		return
	}

	owner, err := blockchain.GetOwnerAddress()
	if context.WithFault(err, runtime.BaseOpsErrorShouldFailExecution()) {
		return
	}

	managedType.SetBytes(destinationHandle, owner)
}

// ManagedCaller VMHooks implementation.
// @autogenerate(VMHooks)
func (context *VMHooksImpl) ManagedCaller(destinationHandle int32) {
	managedType := context.GetManagedTypesContext()
	runtime := context.GetRuntimeContext()
	metering := context.GetMeteringContext()

	gasToUse := metering.GasSchedule().BaseOpsAPICost.GetCaller
	err := metering.UseGasBoundedAndAddTracedGas(managedCallerName, gasToUse)
	if context.WithFault(err, runtime.BaseOpsErrorShouldFailExecution()) {
		return
	}

	caller := runtime.GetVMInput().CallerAddr
	managedType.SetBytes(destinationHandle, caller)
}

// ManagedSignalError VMHooks implementation.
// @autogenerate(VMHooks)
func (context *VMHooksImpl) ManagedSignalError(errHandle int32) {
	managedType := context.GetManagedTypesContext()
	runtime := context.GetRuntimeContext()
	metering := context.GetMeteringContext()
	metering.StartGasTracing(managedSignalErrorName)

	gasToUse := metering.GasSchedule().BaseOpsAPICost.SignalError
	err := metering.UseGasBounded(gasToUse)
	if context.WithFault(err, runtime.UseGasBoundedShouldFailExecution()) {
		return
	}

	errBytes, err := managedType.GetBytes(errHandle)
	if context.WithFault(err, runtime.ManagedBufferAPIErrorShouldFailExecution()) {
		return
	}

	err = managedType.ConsumeGasForBytes(errBytes)
	if context.WithFault(err, runtime.UseGasBoundedShouldFailExecution()) {
		return
	}

	gasToUse = metering.GasSchedule().BaseOperationCost.PersistPerByte * uint64(len(errBytes))
	err = metering.UseGasBounded(gasToUse)
	if err != nil && context.WithFault(err, runtime.UseGasBoundedShouldFailExecution()) {
		return
	}

	runtime.SignalUserError(string(errBytes))
}

// ManagedWriteLog VMHooks implementation.
// @autogenerate(VMHooks)
func (context *VMHooksImpl) ManagedWriteLog(
	topicsHandle int32,
	dataHandle int32,
) {
	runtime := context.GetRuntimeContext()
	output := context.GetOutputContext()
	metering := context.GetMeteringContext()
	managedType := context.GetManagedTypesContext()
	metering.StartGasTracing(managedWriteLogName)

	topics, sumOfTopicByteLengths, err := managedType.ReadManagedVecOfManagedBuffers(topicsHandle)
	if context.WithFault(err, runtime.BaseOpsErrorShouldFailExecution()) {
		return
	}

	dataBytes, err := managedType.GetBytes(dataHandle)
	if context.WithFault(err, runtime.ManagedBufferAPIErrorShouldFailExecution()) {
		return
	}

	err = managedType.ConsumeGasForBytes(dataBytes)
	if context.WithFault(err, runtime.ManagedBufferAPIErrorShouldFailExecution()) {
		return
	}

	dataByteLen := uint64(len(dataBytes))

	gasToUse := metering.GasSchedule().BaseOpsAPICost.Log
	gasForData := math.MulUint64(
		metering.GasSchedule().BaseOperationCost.DataCopyPerByte,
		sumOfTopicByteLengths+dataByteLen)
	gasToUse = math.AddUint64(gasToUse, gasForData)
	err = metering.UseGasBounded(gasToUse)
	if context.WithFault(err, runtime.ManagedBufferAPIErrorShouldFailExecution()) {
		return
	}

	output.WriteLog(runtime.GetContextAddress(), topics, [][]byte{dataBytes})
}

// ManagedGetOriginalTxHash VMHooks implementation.
// @autogenerate(VMHooks)
func (context *VMHooksImpl) ManagedGetOriginalTxHash(resultHandle int32) {
	runtime := context.GetRuntimeContext()
	metering := context.GetMeteringContext()
	managedType := context.GetManagedTypesContext()

	gasToUse := metering.GasSchedule().BaseOpsAPICost.GetOriginalTxHash
	err := metering.UseGasBounded(gasToUse)
	if context.WithFault(err, runtime.ManagedBufferAPIErrorShouldFailExecution()) {
		return
	}

	managedType.SetBytes(resultHandle, runtime.GetOriginalTxHash())
}

// ManagedGetStateRootHash VMHooks implementation.
// @autogenerate(VMHooks)
func (context *VMHooksImpl) ManagedGetStateRootHash(resultHandle int32) {
	blockchain := context.GetBlockchainContext()
	metering := context.GetMeteringContext()
	managedType := context.GetManagedTypesContext()

	gasToUse := metering.GasSchedule().BaseOpsAPICost.GetStateRootHash
	err := metering.UseGasBoundedAndAddTracedGas(managedGetStateRootHashName, gasToUse)
	if context.WithFault(err, context.GetRuntimeContext().BaseOpsErrorShouldFailExecution()) {
		return
	}

	managedType.SetBytes(resultHandle, blockchain.GetStateRootHash())
}

// ManagedGetBlockRandomSeed VMHooks implementation.
// @autogenerate(VMHooks)
func (context *VMHooksImpl) ManagedGetBlockRandomSeed(resultHandle int32) {
	blockchain := context.GetBlockchainContext()
	metering := context.GetMeteringContext()
	managedType := context.GetManagedTypesContext()

	gasToUse := metering.GasSchedule().BaseOpsAPICost.GetBlockRandomSeed
	err := metering.UseGasBoundedAndAddTracedGas(managedGetBlockRandomSeedName, gasToUse)
	if context.WithFault(err, context.GetRuntimeContext().BaseOpsErrorShouldFailExecution()) {
		return
	}

	managedType.SetBytes(resultHandle, blockchain.CurrentRandomSeed())
}

// ManagedGetPrevBlockRandomSeed VMHooks implementation.
// @autogenerate(VMHooks)
func (context *VMHooksImpl) ManagedGetPrevBlockRandomSeed(resultHandle int32) {
	blockchain := context.GetBlockchainContext()
	metering := context.GetMeteringContext()
	managedType := context.GetManagedTypesContext()

	gasToUse := metering.GasSchedule().BaseOpsAPICost.GetBlockRandomSeed
	err := metering.UseGasBoundedAndAddTracedGas(managedGetPrevBlockRandomSeedName, gasToUse)
	if context.WithFault(err, context.GetRuntimeContext().BaseOpsErrorShouldFailExecution()) {
		return
	}

	managedType.SetBytes(resultHandle, blockchain.LastRandomSeed())
}

// ManagedGetReturnData VMHooks implementation.
// @autogenerate(VMHooks)
func (context *VMHooksImpl) ManagedGetReturnData(resultID int32, resultHandle int32) {
	runtime := context.GetRuntimeContext()
	output := context.GetOutputContext()
	metering := context.GetMeteringContext()
	managedType := context.GetManagedTypesContext()

	gasToUse := metering.GasSchedule().BaseOpsAPICost.GetReturnData
	err := metering.UseGasBoundedAndAddTracedGas(managedGetReturnDataName, gasToUse)
	if context.WithFault(err, runtime.BaseOpsErrorShouldFailExecution()) {
		return
	}

	returnData := output.ReturnData()
	if resultID >= int32(len(returnData)) || resultID < 0 {
		_ = context.WithFault(vmhost.ErrArgOutOfRange, runtime.BaseOpsErrorShouldFailExecution())
		return
	}

	managedType.SetBytes(resultHandle, returnData[resultID])
}

// ManagedGetMultiDCDTCallValue VMHooks implementation.
// @autogenerate(VMHooks)
func (context *VMHooksImpl) ManagedGetMultiDCDTCallValue(multiCallValueHandle int32) {
	runtime := context.GetRuntimeContext()
	metering := context.GetMeteringContext()
	managedType := context.GetManagedTypesContext()

	gasToUse := metering.GasSchedule().BaseOpsAPICost.GetCallValue
	err := metering.UseGasBoundedAndAddTracedGas(managedGetMultiDCDTCallValueName, gasToUse)
	if context.WithFault(err, runtime.BaseOpsErrorShouldFailExecution()) {
		return
	}

	dcdtTransfers := runtime.GetVMInput().DCDTTransfers
	multiCallBytes := writeDCDTTransfersToBytes(managedType, dcdtTransfers)
	err = managedType.ConsumeGasForBytes(multiCallBytes)
	if context.WithFault(err, runtime.BaseOpsErrorShouldFailExecution()) {
		return
	}

	managedType.SetBytes(multiCallValueHandle, multiCallBytes)
}

// ManagedGetBackTransfers VMHooks implementation.
// @autogenerate(VMHooks)
func (context *VMHooksImpl) ManagedGetBackTransfers(dcdtTransfersValueHandle int32, rewaValueHandle int32) {
	metering := context.GetMeteringContext()
	managedType := context.GetManagedTypesContext()

	gasToUse := metering.GasSchedule().BaseOpsAPICost.GetCallValue
	err := metering.UseGasBoundedAndAddTracedGas(managedGetMultiDCDTCallValueName, gasToUse)
	if context.WithFault(err, context.GetRuntimeContext().BaseOpsErrorShouldFailExecution()) {
		return
	}

	dcdtTransfers, transferValue := managedType.GetBackTransfers()
	multiCallBytes := writeDCDTTransfersToBytes(managedType, dcdtTransfers)
	err = managedType.ConsumeGasForBytes(multiCallBytes)
	if context.WithFault(err, context.GetRuntimeContext().BaseOpsErrorShouldFailExecution()) {
		return
	}

	managedType.SetBytes(dcdtTransfersValueHandle, multiCallBytes)
	rewaValue := managedType.GetBigIntOrCreate(rewaValueHandle)
	rewaValue.SetBytes(transferValue.Bytes())
}

// ManagedGetDCDTBalance VMHooks implementation.
// @autogenerate(VMHooks)
func (context *VMHooksImpl) ManagedGetDCDTBalance(addressHandle int32, tokenIDHandle int32, nonce int64, valueHandle int32) {
	runtime := context.GetRuntimeContext()
	metering := context.GetMeteringContext()
	blockchain := context.GetBlockchainContext()
	managedType := context.GetManagedTypesContext()

	gasToUse := metering.GasSchedule().BaseOpsAPICost.GetExternalBalance
	err := metering.UseGasBoundedAndAddTracedGas(managedGetDCDTBalanceName, gasToUse)
	if context.WithFault(err, runtime.BaseOpsErrorShouldFailExecution()) {
		return
	}

	address, err := managedType.GetBytes(addressHandle)
	if err != nil {
		_ = context.WithFault(vmhost.ErrArgOutOfRange, runtime.BaseOpsErrorShouldFailExecution())
		return
	}
	tokenID, err := managedType.GetBytes(tokenIDHandle)
	if err != nil {
		_ = context.WithFault(vmhost.ErrArgOutOfRange, runtime.BaseOpsErrorShouldFailExecution())
		return
	}

	dcdtToken, err := blockchain.GetDCDTToken(address, tokenID, uint64(nonce))
	if err != nil {
		_ = context.WithFault(vmhost.ErrArgOutOfRange, runtime.BaseOpsErrorShouldFailExecution())
		return
	}

	value := managedType.GetBigIntOrCreate(valueHandle)
	value.Set(dcdtToken.Value)
}

// ManagedGetDCDTTokenData VMHooks implementation.
// @autogenerate(VMHooks)
func (context *VMHooksImpl) ManagedGetDCDTTokenData(
	addressHandle int32,
	tokenIDHandle int32,
	nonce int64,
	valueHandle, propertiesHandle, hashHandle, nameHandle, attributesHandle, creatorHandle, royaltiesHandle, urisHandle int32) {
	host := context.GetVMHost()
	ManagedGetDCDTTokenDataWithHost(
		host,
		addressHandle,
		tokenIDHandle,
		nonce,
		valueHandle, propertiesHandle, hashHandle, nameHandle, attributesHandle, creatorHandle, royaltiesHandle, urisHandle)

}

func ManagedGetDCDTTokenDataWithHost(
	host vmhost.VMHost,
	addressHandle int32,
	tokenIDHandle int32,
	nonce int64,
	valueHandle, propertiesHandle, hashHandle, nameHandle, attributesHandle, creatorHandle, royaltiesHandle, urisHandle int32) {
	runtime := host.Runtime()
	metering := host.Metering()
	blockchain := host.Blockchain()
	managedType := host.ManagedTypes()
	metering.StartGasTracing(managedGetDCDTTokenDataName)

	gasToUse := metering.GasSchedule().BaseOpsAPICost.GetExternalBalance
	err := metering.UseGasBounded(gasToUse)
	if WithFaultAndHost(host, err, runtime.ManagedBufferAPIErrorShouldFailExecution()) {
		return
	}

	address, err := managedType.GetBytes(addressHandle)
	if err != nil {
		_ = WithFaultAndHost(host, vmhost.ErrArgOutOfRange, runtime.BaseOpsErrorShouldFailExecution())
		return
	}
	tokenID, err := managedType.GetBytes(tokenIDHandle)
	if err != nil {
		_ = WithFaultAndHost(host, vmhost.ErrArgOutOfRange, runtime.BaseOpsErrorShouldFailExecution())
		return
	}

	dcdtToken, err := blockchain.GetDCDTToken(address, tokenID, uint64(nonce))
	if err != nil {
		_ = WithFaultAndHost(host, vmhost.ErrArgOutOfRange, runtime.BaseOpsErrorShouldFailExecution())
		return
	}

	value := managedType.GetBigIntOrCreate(valueHandle)
	value.Set(dcdtToken.Value)

	managedType.SetBytes(propertiesHandle, dcdtToken.Properties)
	if dcdtToken.TokenMetaData != nil {
		managedType.SetBytes(hashHandle, dcdtToken.TokenMetaData.Hash)
		err = managedType.ConsumeGasForBytes(dcdtToken.TokenMetaData.Hash)
		if WithFaultAndHost(host, err, runtime.BaseOpsErrorShouldFailExecution()) {
			return
		}
		managedType.SetBytes(nameHandle, dcdtToken.TokenMetaData.Name)
		err = managedType.ConsumeGasForBytes(dcdtToken.TokenMetaData.Name)
		if WithFaultAndHost(host, err, runtime.BaseOpsErrorShouldFailExecution()) {
			return
		}
		managedType.SetBytes(attributesHandle, dcdtToken.TokenMetaData.Attributes)
		err = managedType.ConsumeGasForBytes(dcdtToken.TokenMetaData.Attributes)
		if WithFaultAndHost(host, err, runtime.BaseOpsErrorShouldFailExecution()) {
			return
		}
		managedType.SetBytes(creatorHandle, dcdtToken.TokenMetaData.Creator)
		err = managedType.ConsumeGasForBytes(dcdtToken.TokenMetaData.Creator)
		if WithFaultAndHost(host, err, runtime.BaseOpsErrorShouldFailExecution()) {
			return
		}
		royalties := managedType.GetBigIntOrCreate(royaltiesHandle)
		royalties.SetUint64(uint64(dcdtToken.TokenMetaData.Royalties))

		err = managedType.WriteManagedVecOfManagedBuffers(dcdtToken.TokenMetaData.URIs, urisHandle)
		if WithFaultAndHost(host, err, runtime.BaseOpsErrorShouldFailExecution()) {
			return
		}
	}

}

// ManagedAsyncCall VMHooks implementation.
// @autogenerate(VMHooks)
func (context *VMHooksImpl) ManagedAsyncCall(
	destHandle int32,
	valueHandle int32,
	functionHandle int32,
	argumentsHandle int32) {
	host := context.GetVMHost()
	ManagedAsyncCallWithHost(
		host,
		destHandle,
		valueHandle,
		functionHandle,
		argumentsHandle)
}

func ManagedAsyncCallWithHost(
	host vmhost.VMHost,
	destHandle int32,
	valueHandle int32,
	functionHandle int32,
	argumentsHandle int32) {
	runtime := host.Runtime()
	async := host.Async()
	metering := host.Metering()
	managedType := host.ManagedTypes()
	metering.StartGasTracing(managedAsyncCallName)

	gasSchedule := metering.GasSchedule()
	gasToUse := gasSchedule.BaseOpsAPICost.AsyncCallStep
	err := metering.UseGasBounded(gasToUse)
	if WithFaultAndHost(host, err, runtime.UseGasBoundedShouldFailExecution()) {
		return
	}

	vmInput, err := readDestinationFunctionArguments(host, destHandle, functionHandle, argumentsHandle)
	if WithFaultAndHost(host, err, host.Runtime().BaseOpsErrorShouldFailExecution()) {
		return
	}

	data := makeCrossShardCallFromInput(vmInput.function, vmInput.arguments)

	value, err := managedType.GetBigInt(valueHandle)
	if err != nil {
		_ = WithFaultAndHost(host, vmhost.ErrArgOutOfRange, host.Runtime().BaseOpsErrorShouldFailExecution())
		return
	}

	gasToUse = math.MulUint64(gasSchedule.BaseOperationCost.DataCopyPerByte, uint64(len(data)))
	err = metering.UseGasBounded(gasToUse)
	if WithFaultAndHost(host, err, runtime.UseGasBoundedShouldFailExecution()) {
		return
	}

	err = async.RegisterLegacyAsyncCall(vmInput.destination, []byte(data), value.Bytes())
	if errors.Is(err, vmhost.ErrNotEnoughGas) {
		runtime.SetRuntimeBreakpointValue(vmhost.BreakpointOutOfGas)
		return
	}
	if WithFaultAndHost(host, err, host.Runtime().BaseOpsErrorShouldFailExecution()) {
		return
	}
}

// ManagedCreateAsyncCall VMHooks implementation.
// @autogenerate(VMHooks)
func (context *VMHooksImpl) ManagedCreateAsyncCall(
	destHandle int32,
	valueHandle int32,
	functionHandle int32,
	argumentsHandle int32,
	successOffset executor.MemPtr,
	successLength executor.MemLength,
	errorOffset executor.MemPtr,
	errorLength executor.MemLength,
	gas int64,
	extraGasForCallback int64,
	callbackClosureHandle int32,
) int32 {

	host := context.GetVMHost()
	runtime := host.Runtime()
	managedType := host.ManagedTypes()

	vmInput, err := readDestinationFunctionArguments(host, destHandle, functionHandle, argumentsHandle)
	if WithFaultAndHost(host, err, host.Runtime().BaseOpsErrorShouldFailExecution()) {
		return 1
	}

	data := makeCrossShardCallFromInput(vmInput.function, vmInput.arguments)

	value, err := managedType.GetBigInt(valueHandle)
	if err != nil {
		_ = context.WithFault(vmhost.ErrArgOutOfRange, runtime.BaseOpsErrorShouldFailExecution())
		return 1
	}

	successFunc, err := context.MemLoad(successOffset, successLength)
	if WithFaultAndHost(host, err, runtime.BaseOpsErrorShouldFailExecution()) {
		return 1
	}

	errorFunc, err := context.MemLoad(errorOffset, errorLength)
	if WithFaultAndHost(host, err, runtime.BaseOpsErrorShouldFailExecution()) {
		return 1
	}

	callbackClosure, err := managedType.GetBytes(callbackClosureHandle)
	if WithFaultAndHost(host, err, runtime.BaseOpsErrorShouldFailExecution()) {
		return 1
	}

	return CreateAsyncCallWithTypedArgs(host,
		vmInput.destination,
		value.Bytes(),
		[]byte(data),
		successFunc,
		errorFunc,
		gas,
		extraGasForCallback,
		callbackClosure)
}

// ManagedGetCallbackClosure VMHooks implementation.
// @autogenerate(VMHooks)
func (context *VMHooksImpl) ManagedGetCallbackClosure(
	callbackClosureHandle int32,
) {
	host := context.GetVMHost()
	GetCallbackClosureWithHost(host, callbackClosureHandle)
}

func GetCallbackClosureWithHost(
	host vmhost.VMHost,
	callbackClosureHandle int32,
) {
	runtime := host.Runtime()
	async := host.Async()
	metering := host.Metering()
	managedTypes := host.ManagedTypes()

	metering.StartGasTracing(managedGetCallbackClosure)

	gasToUse := metering.GasSchedule().BaseOpsAPICost.GetCallbackClosure
	err := metering.UseGasBounded(gasToUse)
	if WithFaultAndHost(host, err, runtime.ManagedBufferAPIErrorShouldFailExecution()) {
		return
	}

	callbackClosure, err := async.GetCallbackClosure()
	if WithFaultAndHost(host, err, runtime.BaseOpsErrorShouldFailExecution()) {
		return
	}

	managedTypes.SetBytes(callbackClosureHandle, callbackClosure)
}

// ManagedUpgradeFromSourceContract VMHooks implementation.
// @autogenerate(VMHooks)
func (context *VMHooksImpl) ManagedUpgradeFromSourceContract(
	destHandle int32,
	gas int64,
	valueHandle int32,
	addressHandle int32,
	codeMetadataHandle int32,
	argumentsHandle int32,
	resultHandle int32,
) {
	host := context.GetVMHost()
	runtime := host.Runtime()
	metering := host.Metering()
	managedType := host.ManagedTypes()
	metering.StartGasTracing(managedUpgradeFromSourceContractName)

	gasToUse := metering.GasSchedule().BaseOpsAPICost.CreateContract
	err := metering.UseGasBounded(gasToUse)
	if context.WithFault(err, runtime.UseGasBoundedShouldFailExecution()) {
		return
	}

	vmInput, err := readDestinationValueArguments(host, destHandle, valueHandle, argumentsHandle)
	if WithFaultAndHost(host, err, runtime.BaseOpsErrorShouldFailExecution()) {
		return
	}

	sourceContractAddress, err := managedType.GetBytes(addressHandle)
	if WithFaultAndHost(host, err, runtime.BaseOpsErrorShouldFailExecution()) {
		return
	}

	codeMetadata, err := managedType.GetBytes(codeMetadataHandle)
	if WithFaultAndHost(host, err, runtime.BaseOpsErrorShouldFailExecution()) {
		return
	}

	lenReturnData := len(host.Output().ReturnData())

	UpgradeFromSourceContractWithTypedArgs(
		host,
		sourceContractAddress,
		vmInput.destination,
		vmInput.value.Bytes(),
		vmInput.arguments,
		gas,
		codeMetadata,
	)
	err = setReturnDataIfExists(host, lenReturnData, resultHandle)
	if WithFaultAndHost(host, err, runtime.UseGasBoundedShouldFailExecution()) {
		return
	}
}

// ManagedUpgradeContract VMHooks implementation.
// @autogenerate(VMHooks)
func (context *VMHooksImpl) ManagedUpgradeContract(
	destHandle int32,
	gas int64,
	valueHandle int32,
	codeHandle int32,
	codeMetadataHandle int32,
	argumentsHandle int32,
	resultHandle int32,
) {
	host := context.GetVMHost()
	runtime := host.Runtime()
	metering := host.Metering()
	managedType := host.ManagedTypes()
	metering.StartGasTracing(managedUpgradeContractName)

	gasToUse := metering.GasSchedule().BaseOpsAPICost.CreateContract
	err := metering.UseGasBounded(gasToUse)
	if context.WithFault(err, runtime.UseGasBoundedShouldFailExecution()) {
		return
	}

	vmInput, err := readDestinationValueArguments(host, destHandle, valueHandle, argumentsHandle)
	if WithFaultAndHost(host, err, runtime.BaseOpsErrorShouldFailExecution()) {
		return
	}

	codeMetadata, err := managedType.GetBytes(codeMetadataHandle)
	if WithFaultAndHost(host, err, runtime.BaseOpsErrorShouldFailExecution()) {
		return
	}

	code, err := managedType.GetBytes(codeHandle)
	if WithFaultAndHost(host, err, runtime.BaseOpsErrorShouldFailExecution()) {
		return
	}

	if WithFaultAndHost(host, err, runtime.BaseOpsErrorShouldFailExecution()) {
		return
	}

	lenReturnData := len(host.Output().ReturnData())

	upgradeContract(host, vmInput.destination, code, codeMetadata, vmInput.value.Bytes(), vmInput.arguments, gas)
	err = setReturnDataIfExists(host, lenReturnData, resultHandle)
	if WithFaultAndHost(host, err, runtime.UseGasBoundedShouldFailExecution()) {
		return
	}
}

// ManagedDeleteContract VMHooks implementation.
// @autogenerate(VMHooks)
func (context *VMHooksImpl) ManagedDeleteContract(
	destHandle int32,
	gasLimit int64,
	argumentsHandle int32,
) {
	host := context.GetVMHost()
	ManagedDeleteContractWithHost(
		host,
		destHandle,
		gasLimit,
		argumentsHandle,
	)
}

func ManagedDeleteContractWithHost(
	host vmhost.VMHost,
	destHandle int32,
	gasLimit int64,
	argumentsHandle int32,
) {
	runtime := host.Runtime()
	metering := host.Metering()
	managedType := host.ManagedTypes()
	metering.StartGasTracing(deleteContractName)

	gasToUse := metering.GasSchedule().BaseOpsAPICost.CreateContract
	err := metering.UseGasBounded(gasToUse)
	if WithFaultAndHost(host, err, runtime.UseGasBoundedShouldFailExecution()) {
		return
	}

	calledSCAddress, err := managedType.GetBytes(destHandle)
	if WithFaultAndHost(host, err, runtime.BaseOpsErrorShouldFailExecution()) {
		return
	}

	data, _, err := managedType.ReadManagedVecOfManagedBuffers(argumentsHandle)
	if WithFaultAndHost(host, err, runtime.BaseOpsErrorShouldFailExecution()) {
		return
	}

	deleteContract(
		host,
		calledSCAddress,
		data,
		gasLimit,
	)
}

// ManagedDeployFromSourceContract VMHooks implementation.
// @autogenerate(VMHooks)
func (context *VMHooksImpl) ManagedDeployFromSourceContract(
	gas int64,
	valueHandle int32,
	addressHandle int32,
	codeMetadataHandle int32,
	argumentsHandle int32,
	resultAddressHandle int32,
	resultHandle int32,
) int32 {
	host := context.GetVMHost()
	runtime := host.Runtime()
	metering := host.Metering()
	managedType := host.ManagedTypes()
	metering.StartGasTracing(managedDeployFromSourceContractName)

	gasToUse := metering.GasSchedule().BaseOpsAPICost.CreateContract
	err := metering.UseGasBounded(gasToUse)
	if context.WithFault(err, runtime.UseGasBoundedShouldFailExecution()) {
		return -1
	}

	vmInput, err := readDestinationValueArguments(host, addressHandle, valueHandle, argumentsHandle)
	if WithFaultAndHost(host, err, runtime.BaseOpsErrorShouldFailExecution()) {
		return 1
	}

	codeMetadata, err := managedType.GetBytes(codeMetadataHandle)
	if WithFaultAndHost(host, err, runtime.BaseOpsErrorShouldFailExecution()) {
		return 1
	}

	lenReturnData := len(host.Output().ReturnData())

	newAddress, err := DeployFromSourceContractWithTypedArgs(
		host,
		vmInput.destination,
		codeMetadata,
		vmInput.value,
		vmInput.arguments,
		gas,
	)
	if WithFaultAndHost(host, err, runtime.BaseOpsErrorShouldFailExecution()) {
		return 1
	}

	managedType.SetBytes(resultAddressHandle, newAddress)
	err = setReturnDataIfExists(host, lenReturnData, resultHandle)
	if WithFaultAndHost(host, err, runtime.UseGasBoundedShouldFailExecution()) {
		return 1
	}

	return 0
}

// ManagedCreateContract VMHooks implementation.
// @autogenerate(VMHooks)
func (context *VMHooksImpl) ManagedCreateContract(
	gas int64,
	valueHandle int32,
	codeHandle int32,
	codeMetadataHandle int32,
	argumentsHandle int32,
	resultAddressHandle int32,
	resultHandle int32,
) int32 {
	host := context.GetVMHost()
	runtime := host.Runtime()
	metering := host.Metering()
	managedType := host.ManagedTypes()
	metering.StartGasTracing(managedCreateContractName)

	gasToUse := metering.GasSchedule().BaseOpsAPICost.CreateContract
	err := metering.UseGasBounded(gasToUse)
	if context.WithFault(err, runtime.UseGasBoundedShouldFailExecution()) {
		return -1
	}

	sender := runtime.GetContextAddress()
	value, err := managedType.GetBigInt(valueHandle)
	if WithFaultAndHost(host, err, runtime.BaseOpsErrorShouldFailExecution()) {
		return 1
	}

	data, actualLen, err := managedType.ReadManagedVecOfManagedBuffers(argumentsHandle)
	if WithFaultAndHost(host, err, runtime.BaseOpsErrorShouldFailExecution()) {
		return 1
	}

	gasToUse = math.MulUint64(metering.GasSchedule().BaseOperationCost.DataCopyPerByte, actualLen)
	err = metering.UseGasBounded(gasToUse)
	if context.WithFault(err, runtime.UseGasBoundedShouldFailExecution()) {
		return -1
	}

	codeMetadata, err := managedType.GetBytes(codeMetadataHandle)
	if WithFaultAndHost(host, err, runtime.BaseOpsErrorShouldFailExecution()) {
		return 1
	}

	code, err := managedType.GetBytes(codeHandle)
	if WithFaultAndHost(host, err, runtime.BaseOpsErrorShouldFailExecution()) {
		return 1
	}

	lenReturnData := len(host.Output().ReturnData())
	newAddress, err := createContract(sender, data, value, gas, code, codeMetadata, host, CreateContract)
	if WithFaultAndHost(host, err, runtime.BaseOpsErrorShouldFailExecution()) {
		return 1
	}

	managedType.SetBytes(resultAddressHandle, newAddress)
	err = setReturnDataIfExists(host, lenReturnData, resultHandle)
	if WithFaultAndHost(host, err, runtime.UseGasBoundedShouldFailExecution()) {
		return 1
	}

	return 0
}

func setReturnDataIfExists(
	host vmhost.VMHost,
	oldLen int,
	resultHandle int32,
) error {
	returnData := host.Output().ReturnData()
	if len(returnData) > oldLen {
		return host.ManagedTypes().WriteManagedVecOfManagedBuffers(returnData[oldLen:], resultHandle)
	}

	host.ManagedTypes().SetBytes(resultHandle, make([]byte, 0))
	return nil
}

// ManagedExecuteReadOnly VMHooks implementation.
// @autogenerate(VMHooks)
func (context *VMHooksImpl) ManagedExecuteReadOnly(
	gas int64,
	addressHandle int32,
	functionHandle int32,
	argumentsHandle int32,
	resultHandle int32,
) int32 {
	host := context.GetVMHost()
	metering := host.Metering()
	metering.StartGasTracing(managedExecuteReadOnlyName)

	vmInput, err := readDestinationFunctionArguments(host, addressHandle, functionHandle, argumentsHandle)
	if WithFaultAndHost(host, err, host.Runtime().BaseOpsErrorShouldFailExecution()) {
		return -1
	}

	lenReturnData := len(host.Output().ReturnData())
	returnVal := ExecuteReadOnlyWithTypedArguments(
		host,
		gas,
		[]byte(vmInput.function),
		vmInput.destination,
		vmInput.arguments,
	)
	err = setReturnDataIfExists(host, lenReturnData, resultHandle)
	if WithFaultAndHost(host, err, host.Runtime().UseGasBoundedShouldFailExecution()) {
		return -1
	}

	return returnVal
}

// ManagedExecuteOnSameContext VMHooks implementation.
// @autogenerate(VMHooks)
func (context *VMHooksImpl) ManagedExecuteOnSameContext(
	gas int64,
	addressHandle int32,
	valueHandle int32,
	functionHandle int32,
	argumentsHandle int32,
	resultHandle int32,
) int32 {
	host := context.GetVMHost()
	metering := host.Metering()
	metering.StartGasTracing(managedExecuteOnSameContextName)

	vmInput, err := readDestinationValueFunctionArguments(host, addressHandle, valueHandle, functionHandle, argumentsHandle)
	if WithFaultAndHost(host, err, host.Runtime().BaseOpsErrorShouldFailExecution()) {
		return -1
	}

	lenReturnData := len(host.Output().ReturnData())
	returnVal := ExecuteOnSameContextWithTypedArgs(
		host,
		gas,
		vmInput.value,
		[]byte(vmInput.function),
		vmInput.destination,
		vmInput.arguments,
	)
	err = setReturnDataIfExists(host, lenReturnData, resultHandle)
	if WithFaultAndHost(host, err, host.Runtime().UseGasBoundedShouldFailExecution()) {
		return -1
	}

	return returnVal
}

// ManagedExecuteOnDestContext VMHooks implementation.
// @autogenerate(VMHooks)
func (context *VMHooksImpl) ManagedExecuteOnDestContext(
	gas int64,
	addressHandle int32,
	valueHandle int32,
	functionHandle int32,
	argumentsHandle int32,
	resultHandle int32,
) int32 {
	host := context.GetVMHost()
	metering := host.Metering()
	metering.StartGasTracing(managedExecuteOnDestContextName)

	vmInput, err := readDestinationValueFunctionArguments(host, addressHandle, valueHandle, functionHandle, argumentsHandle)
	if WithFaultAndHost(host, err, host.Runtime().BaseOpsErrorShouldFailExecution()) {
		return -1
	}

	lenReturnData := len(host.Output().ReturnData())
	returnVal := ExecuteOnDestContextWithTypedArgs(
		host,
		gas,
		vmInput.value,
		[]byte(vmInput.function),
		vmInput.destination,
		vmInput.arguments,
	)
	err = setReturnDataIfExists(host, lenReturnData, resultHandle)
	if WithFaultAndHost(host, err, host.Runtime().UseGasBoundedShouldFailExecution()) {
		return -1
	}

	return returnVal
}

// ManagedMultiTransferDCDTNFTExecute VMHooks implementation.
// @autogenerate(VMHooks)
func (context *VMHooksImpl) ManagedMultiTransferDCDTNFTExecute(
	dstHandle int32,
	tokenTransfersHandle int32,
	gasLimit int64,
	functionHandle int32,
	argumentsHandle int32,
) int32 {
	host := context.GetVMHost()
	managedType := host.ManagedTypes()
	runtime := host.Runtime()
	metering := host.Metering()
	metering.StartGasTracing(managedMultiTransferDCDTNFTExecuteName)

	vmInput, err := readDestinationFunctionArguments(host, dstHandle, functionHandle, argumentsHandle)
	if WithFaultAndHost(host, err, runtime.BaseOpsErrorShouldFailExecution()) {
		return -1
	}

	transfers, err := readDCDTTransfers(managedType, runtime, tokenTransfersHandle)
	if WithFaultAndHost(host, err, runtime.BaseOpsErrorShouldFailExecution()) {
		return -1
	}

	return TransferDCDTNFTExecuteWithTypedArgs(
		host,
		vmInput.destination,
		transfers,
		gasLimit,
		[]byte(vmInput.function),
		vmInput.arguments,
	)
}

// ManagedTransferValueExecute VMHooks implementation.
// @autogenerate(VMHooks)
func (context *VMHooksImpl) ManagedTransferValueExecute(
	dstHandle int32,
	valueHandle int32,
	gasLimit int64,
	functionHandle int32,
	argumentsHandle int32,
) int32 {
	host := context.GetVMHost()
	metering := host.Metering()
	metering.StartGasTracing(managedTransferValueExecuteName)

	vmInput, err := readDestinationValueFunctionArguments(host, dstHandle, valueHandle, functionHandle, argumentsHandle)
	if WithFaultAndHost(host, err, host.Runtime().BaseOpsErrorShouldFailExecution()) {
		return -1
	}

	return TransferValueExecuteWithTypedArgs(
		host,
		vmInput.destination,
		vmInput.value,
		gasLimit,
		[]byte(vmInput.function),
		vmInput.arguments,
	)
}

// ManagedIsDCDTFrozen VMHooks implementation.
// @autogenerate(VMHooks)
func (context *VMHooksImpl) ManagedIsDCDTFrozen(
	addressHandle int32,
	tokenIDHandle int32,
	nonce int64) int32 {
	host := context.GetVMHost()
	return ManagedIsDCDTFrozenWithHost(host, addressHandle, tokenIDHandle, nonce)
}

func ManagedIsDCDTFrozenWithHost(
	host vmhost.VMHost,
	addressHandle int32,
	tokenIDHandle int32,
	nonce int64) int32 {
	runtime := host.Runtime()
	metering := host.Metering()
	blockchain := host.Blockchain()
	managedType := host.ManagedTypes()

	gasToUse := metering.GasSchedule().BaseOpsAPICost.GetExternalBalance
	err := metering.UseGasBoundedAndAddTracedGas(managedIsDCDTFrozenName, gasToUse)
	if WithFaultAndHost(host, err, runtime.BaseOpsErrorShouldFailExecution()) {
		return -1
	}

	address, err := managedType.GetBytes(addressHandle)
	if err != nil {
		_ = WithFaultAndHost(host, err, runtime.BaseOpsErrorShouldFailExecution())
		return -1
	}
	tokenID, err := managedType.GetBytes(tokenIDHandle)
	if err != nil {
		_ = WithFaultAndHost(host, err, runtime.BaseOpsErrorShouldFailExecution())
		return -1
	}

	dcdtToken, err := blockchain.GetDCDTToken(address, tokenID, uint64(nonce))
	if err != nil {
		_ = WithFaultAndHost(host, err, runtime.BaseOpsErrorShouldFailExecution())
		return -1
	}

	dcdtUserData := builtInFunctions.DCDTUserMetadataFromBytes(dcdtToken.Properties)
	if dcdtUserData.Frozen {
		return 1
	}
	return 0
}

// ManagedIsDCDTLimitedTransfer VMHooks implementation.
// @autogenerate(VMHooks)
func (context *VMHooksImpl) ManagedIsDCDTLimitedTransfer(tokenIDHandle int32) int32 {
	host := context.GetVMHost()
	return ManagedIsDCDTLimitedTransferWithHost(host, tokenIDHandle)
}

func ManagedIsDCDTLimitedTransferWithHost(host vmhost.VMHost, tokenIDHandle int32) int32 {
	runtime := host.Runtime()
	metering := host.Metering()
	blockchain := host.Blockchain()
	managedType := host.ManagedTypes()

	gasToUse := metering.GasSchedule().BaseOpsAPICost.GetExternalBalance
	err := metering.UseGasBoundedAndAddTracedGas(managedIsDCDTLimitedTransferName, gasToUse)
	if WithFaultAndHost(host, err, runtime.BaseOpsErrorShouldFailExecution()) {
		return -1
	}

	tokenID, err := managedType.GetBytes(tokenIDHandle)
	if err != nil {
		_ = WithFaultAndHost(host, err, runtime.BaseOpsErrorShouldFailExecution())
		return -1
	}

	if blockchain.IsLimitedTransfer(tokenID) {
		return 1
	}

	return 0
}

// ManagedIsDCDTPaused VMHooks implementation.
// @autogenerate(VMHooks)
func (context *VMHooksImpl) ManagedIsDCDTPaused(tokenIDHandle int32) int32 {
	host := context.GetVMHost()
	return ManagedIsDCDTPausedWithHost(host, tokenIDHandle)
}

func ManagedIsDCDTPausedWithHost(host vmhost.VMHost, tokenIDHandle int32) int32 {
	runtime := host.Runtime()
	metering := host.Metering()
	blockchain := host.Blockchain()
	managedType := host.ManagedTypes()

	gasToUse := metering.GasSchedule().BaseOpsAPICost.GetExternalBalance
	err := metering.UseGasBoundedAndAddTracedGas(managedIsDCDTPausedName, gasToUse)
	if WithFaultAndHost(host, err, runtime.BaseOpsErrorShouldFailExecution()) {
		return -1
	}

	tokenID, err := managedType.GetBytes(tokenIDHandle)
	if err != nil {
		_ = WithFaultAndHost(host, err, runtime.BaseOpsErrorShouldFailExecution())
		return -1
	}

	if blockchain.IsPaused(tokenID) {
		return 1
	}

	return 0
}

// ManagedBufferToHex VMHooks implementation.
// @autogenerate(VMHooks)
func (context *VMHooksImpl) ManagedBufferToHex(sourceHandle int32, destHandle int32) {
	host := context.GetVMHost()
	ManagedBufferToHexWithHost(host, sourceHandle, destHandle)
}

func ManagedBufferToHexWithHost(host vmhost.VMHost, sourceHandle int32, destHandle int32) {
	runtime := host.Runtime()
	metering := host.Metering()
	managedType := host.ManagedTypes()

	gasToUse := metering.GasSchedule().ManagedBufferAPICost.MBufferSetBytes
	err := metering.UseGasBoundedAndAddTracedGas(managedBufferToHexName, gasToUse)
	if WithFaultAndHost(host, err, runtime.BaseOpsErrorShouldFailExecution()) {
		return
	}

	mBuff, err := managedType.GetBytes(sourceHandle)
	if err != nil {
		WithFaultAndHost(host, err, runtime.BaseOpsErrorShouldFailExecution())
		return
	}

	encoded := hex.EncodeToString(mBuff)
	managedType.SetBytes(destHandle, []byte(encoded))
}

// ManagedGetCodeMetadata VMHooks implementation.
// @autogenerate(VMHooks)
func (context *VMHooksImpl) ManagedGetCodeMetadata(addressHandle int32, responseHandle int32) {
	host := context.GetVMHost()
	ManagedGetCodeMetadataWithHost(host, addressHandle, responseHandle)
}

func ManagedGetCodeMetadataWithHost(host vmhost.VMHost, addressHandle int32, responseHandle int32) {
	runtime := host.Runtime()
	metering := host.Metering()
	managedType := host.ManagedTypes()

	gasToUse := metering.GasSchedule().BaseOpsAPICost.GetCodeMetadata
	err := metering.UseGasBoundedAndAddTracedGas(managedGetCodeMetadataName, gasToUse)
	if WithFaultAndHost(host, err, runtime.BaseOpsErrorShouldFailExecution()) {
		return
	}

	gasToUse = metering.GasSchedule().ManagedBufferAPICost.MBufferSetBytes
	err = metering.UseGasBoundedAndAddTracedGas(managedGetCodeMetadataName, gasToUse)
	if WithFaultAndHost(host, err, runtime.BaseOpsErrorShouldFailExecution()) {
		return
	}

	mBuffAddress, err := managedType.GetBytes(addressHandle)
	if err != nil {
		WithFaultAndHost(host, err, runtime.BaseOpsErrorShouldFailExecution())
		return
	}

	contract, err := host.Blockchain().GetUserAccount(mBuffAddress)
	if err != nil || check.IfNil(contract) {
		WithFaultAndHost(host, err, runtime.BaseOpsErrorShouldFailExecution())
		return
	}

	codeMetadata := contract.GetCodeMetadata()

	managedType.SetBytes(responseHandle, codeMetadata)
}

// ManagedIsBuiltinFunction VMHooks implementation.
// @autogenerate(VMHooks)
func (context *VMHooksImpl) ManagedIsBuiltinFunction(functionNameHandle int32) int32 {
	host := context.GetVMHost()
	return ManagedIsBuiltinFunctionWithHost(host, functionNameHandle)
}

func ManagedIsBuiltinFunctionWithHost(host vmhost.VMHost, functionNameHandle int32) int32 {
	runtime := host.Runtime()
	metering := host.Metering()
	managedType := host.ManagedTypes()

	gasToUse := metering.GasSchedule().BaseOpsAPICost.IsBuiltinFunction
	err := metering.UseGasBoundedAndAddTracedGas(managedIsBuiltinFunction, gasToUse)
	if WithFaultAndHost(host, err, runtime.BaseOpsErrorShouldFailExecution()) {
		return -1
	}

	mBuffFunctionName, err := managedType.GetBytes(functionNameHandle)
	if err != nil {
		WithFaultAndHost(host, err, runtime.BaseOpsErrorShouldFailExecution())
		return -1
	}

	isBuiltinFunction := host.IsBuiltinFunctionName(string(mBuffFunctionName))
	if isBuiltinFunction {
		return 1
	}

	return 0
}