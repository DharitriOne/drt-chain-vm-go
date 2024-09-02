package wasmer2

// Code generated by vmhooks generator. DO NOT EDIT.

// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
// !!!!!!!!!!!!!!!!!!!!!! AUTO-GENERATED FILE !!!!!!!!!!!!!!!!!!!!!!
// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

var empty struct{}

var functionNames = map[string]struct{}{
	"getGasLeft": empty,
	"getSCAddress": empty,
	"getOwnerAddress": empty,
	"getShardOfAddress": empty,
	"isSmartContract": empty,
	"signalError": empty,
	"getExternalBalance": empty,
	"getBlockHash": empty,
	"getDCDTBalance": empty,
	"getDCDTNFTNameLength": empty,
	"getDCDTNFTAttributeLength": empty,
	"getDCDTNFTURILength": empty,
	"getDCDTTokenData": empty,
	"getDCDTLocalRoles": empty,
	"validateTokenIdentifier": empty,
	"transferValue": empty,
	"transferValueExecute": empty,
	"transferDCDTExecute": empty,
	"transferDCDTNFTExecute": empty,
	"multiTransferDCDTNFTExecute": empty,
	"createAsyncCall": empty,
	"setAsyncContextCallback": empty,
	"upgradeContract": empty,
	"upgradeFromSourceContract": empty,
	"deleteContract": empty,
	"asyncCall": empty,
	"getArgumentLength": empty,
	"getArgument": empty,
	"getFunction": empty,
	"getNumArguments": empty,
	"storageStore": empty,
	"storageLoadLength": empty,
	"storageLoadFromAddress": empty,
	"storageLoad": empty,
	"setStorageLock": empty,
	"getStorageLock": empty,
	"isStorageLocked": empty,
	"clearStorageLock": empty,
	"getCaller": empty,
	"checkNoPayment": empty,
	"getCallValue": empty,
	"getDCDTValue": empty,
	"getDCDTValueByIndex": empty,
	"getDCDTTokenName": empty,
	"getDCDTTokenNameByIndex": empty,
	"getDCDTTokenNonce": empty,
	"getDCDTTokenNonceByIndex": empty,
	"getCurrentDCDTNFTNonce": empty,
	"getDCDTTokenType": empty,
	"getDCDTTokenTypeByIndex": empty,
	"getNumDCDTTransfers": empty,
	"getCallValueTokenName": empty,
	"getCallValueTokenNameByIndex": empty,
	"writeLog": empty,
	"writeEventLog": empty,
	"getBlockTimestamp": empty,
	"getBlockNonce": empty,
	"getBlockRound": empty,
	"getBlockEpoch": empty,
	"getBlockRandomSeed": empty,
	"getStateRootHash": empty,
	"getPrevBlockTimestamp": empty,
	"getPrevBlockNonce": empty,
	"getPrevBlockRound": empty,
	"getPrevBlockEpoch": empty,
	"getPrevBlockRandomSeed": empty,
	"finish": empty,
	"executeOnSameContext": empty,
	"executeOnDestContext": empty,
	"executeReadOnly": empty,
	"createContract": empty,
	"deployFromSourceContract": empty,
	"getNumReturnData": empty,
	"getReturnDataSize": empty,
	"getReturnData": empty,
	"cleanReturnData": empty,
	"deleteFromReturnData": empty,
	"getOriginalTxHash": empty,
	"getCurrentTxHash": empty,
	"getPrevTxHash": empty,
	"managedSCAddress": empty,
	"managedOwnerAddress": empty,
	"managedCaller": empty,
	"managedSignalError": empty,
	"managedWriteLog": empty,
	"managedGetOriginalTxHash": empty,
	"managedGetStateRootHash": empty,
	"managedGetBlockRandomSeed": empty,
	"managedGetPrevBlockRandomSeed": empty,
	"managedGetReturnData": empty,
	"managedGetMultiDCDTCallValue": empty,
	"managedGetBackTransfers": empty,
	"managedGetDCDTBalance": empty,
	"managedGetDCDTTokenData": empty,
	"managedAsyncCall": empty,
	"managedCreateAsyncCall": empty,
	"managedGetCallbackClosure": empty,
	"managedUpgradeFromSourceContract": empty,
	"managedUpgradeContract": empty,
	"managedDeleteContract": empty,
	"managedDeployFromSourceContract": empty,
	"managedCreateContract": empty,
	"managedExecuteReadOnly": empty,
	"managedExecuteOnSameContext": empty,
	"managedExecuteOnDestContext": empty,
	"managedMultiTransferDCDTNFTExecute": empty,
	"managedTransferValueExecute": empty,
	"managedIsDCDTFrozen": empty,
	"managedIsDCDTLimitedTransfer": empty,
	"managedIsDCDTPaused": empty,
	"managedBufferToHex": empty,
	"managedGetCodeMetadata": empty,
	"managedIsBuiltinFunction": empty,
	"bigFloatNewFromParts": empty,
	"bigFloatNewFromFrac": empty,
	"bigFloatNewFromSci": empty,
	"bigFloatAdd": empty,
	"bigFloatSub": empty,
	"bigFloatMul": empty,
	"bigFloatDiv": empty,
	"bigFloatNeg": empty,
	"bigFloatClone": empty,
	"bigFloatCmp": empty,
	"bigFloatAbs": empty,
	"bigFloatSign": empty,
	"bigFloatSqrt": empty,
	"bigFloatPow": empty,
	"bigFloatFloor": empty,
	"bigFloatCeil": empty,
	"bigFloatTruncate": empty,
	"bigFloatSetInt64": empty,
	"bigFloatIsInt": empty,
	"bigFloatSetBigInt": empty,
	"bigFloatGetConstPi": empty,
	"bigFloatGetConstE": empty,
	"bigIntGetUnsignedArgument": empty,
	"bigIntGetSignedArgument": empty,
	"bigIntStorageStoreUnsigned": empty,
	"bigIntStorageLoadUnsigned": empty,
	"bigIntGetCallValue": empty,
	"bigIntGetDCDTCallValue": empty,
	"bigIntGetDCDTCallValueByIndex": empty,
	"bigIntGetExternalBalance": empty,
	"bigIntGetDCDTExternalBalance": empty,
	"bigIntNew": empty,
	"bigIntUnsignedByteLength": empty,
	"bigIntSignedByteLength": empty,
	"bigIntGetUnsignedBytes": empty,
	"bigIntGetSignedBytes": empty,
	"bigIntSetUnsignedBytes": empty,
	"bigIntSetSignedBytes": empty,
	"bigIntIsInt64": empty,
	"bigIntGetInt64": empty,
	"bigIntSetInt64": empty,
	"bigIntAdd": empty,
	"bigIntSub": empty,
	"bigIntMul": empty,
	"bigIntTDiv": empty,
	"bigIntTMod": empty,
	"bigIntEDiv": empty,
	"bigIntEMod": empty,
	"bigIntSqrt": empty,
	"bigIntPow": empty,
	"bigIntLog2": empty,
	"bigIntAbs": empty,
	"bigIntNeg": empty,
	"bigIntSign": empty,
	"bigIntCmp": empty,
	"bigIntNot": empty,
	"bigIntAnd": empty,
	"bigIntOr": empty,
	"bigIntXor": empty,
	"bigIntShr": empty,
	"bigIntShl": empty,
	"bigIntFinishUnsigned": empty,
	"bigIntFinishSigned": empty,
	"bigIntToString": empty,
	"mBufferNew": empty,
	"mBufferNewFromBytes": empty,
	"mBufferGetLength": empty,
	"mBufferGetBytes": empty,
	"mBufferGetByteSlice": empty,
	"mBufferCopyByteSlice": empty,
	"mBufferEq": empty,
	"mBufferSetBytes": empty,
	"mBufferSetByteSlice": empty,
	"mBufferAppend": empty,
	"mBufferAppendBytes": empty,
	"mBufferToBigIntUnsigned": empty,
	"mBufferToBigIntSigned": empty,
	"mBufferFromBigIntUnsigned": empty,
	"mBufferFromBigIntSigned": empty,
	"mBufferToBigFloat": empty,
	"mBufferFromBigFloat": empty,
	"mBufferStorageStore": empty,
	"mBufferStorageLoad": empty,
	"mBufferStorageLoadFromAddress": empty,
	"mBufferGetArgument": empty,
	"mBufferFinish": empty,
	"mBufferSetRandom": empty,
	"managedMapNew": empty,
	"managedMapPut": empty,
	"managedMapGet": empty,
	"managedMapRemove": empty,
	"managedMapContains": empty,
	"smallIntGetUnsignedArgument": empty,
	"smallIntGetSignedArgument": empty,
	"smallIntFinishUnsigned": empty,
	"smallIntFinishSigned": empty,
	"smallIntStorageStoreUnsigned": empty,
	"smallIntStorageStoreSigned": empty,
	"smallIntStorageLoadUnsigned": empty,
	"smallIntStorageLoadSigned": empty,
	"int64getArgument": empty,
	"int64finish": empty,
	"int64storageStore": empty,
	"int64storageLoad": empty,
	"sha256": empty,
	"managedSha256": empty,
	"keccak256": empty,
	"managedKeccak256": empty,
	"ripemd160": empty,
	"managedRipemd160": empty,
	"verifyBLS": empty,
	"managedVerifyBLS": empty,
	"verifyEd25519": empty,
	"managedVerifyEd25519": empty,
	"verifyCustomSecp256k1": empty,
	"managedVerifyCustomSecp256k1": empty,
	"verifySecp256k1": empty,
	"managedVerifySecp256k1": empty,
	"encodeSecp256k1DerSignature": empty,
	"managedEncodeSecp256k1DerSignature": empty,
	"addEC": empty,
	"doubleEC": empty,
	"isOnCurveEC": empty,
	"scalarBaseMultEC": empty,
	"managedScalarBaseMultEC": empty,
	"scalarMultEC": empty,
	"managedScalarMultEC": empty,
	"marshalEC": empty,
	"managedMarshalEC": empty,
	"marshalCompressedEC": empty,
	"managedMarshalCompressedEC": empty,
	"unmarshalEC": empty,
	"managedUnmarshalEC": empty,
	"unmarshalCompressedEC": empty,
	"managedUnmarshalCompressedEC": empty,
	"generateKeyEC": empty,
	"managedGenerateKeyEC": empty,
	"createEC": empty,
	"managedCreateEC": empty,
	"getCurveLengthEC": empty,
	"getPrivKeyByteLengthEC": empty,
	"ellipticCurveGetValues": empty,
}
