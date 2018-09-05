package mongerr

import (
	"sync"

	"github.com/globalsign/mgo"
)

func FromMongoErr(queryError *mgo.QueryError) Error {
	mapLock.RLock()
	ferr, ok := errMap[queryError.Code]
	mapLock.RLock()
	if ok {
		return ferr()
	}

	var err = Error{
		Code:    queryError.Code,
		Message: queryError.Message,
	}
	mapLock.Lock()
	errMap[queryError.Code] = func() Error {
		return err
	}
	mapLock.Unlock()
	return err
}

var mapLock = sync.RWMutex{}

var errMap = map[int]func() Error{
	0:     ErrOK,
	1:     ErrInternalError,
	2:     ErrBadValue,
	3:     ErrOBSOLETE_DuplicateKey,
	4:     ErrNoSuchKey,
	5:     ErrGraphContainsCycle,
	6:     ErrHostUnreachable,
	7:     ErrHostNotFound,
	8:     ErrUnknownError,
	9:     ErrFailedToParse,
	10:    ErrCannotMutateObject,
	11:    ErrUserNotFound,
	12:    ErrUnsupportedFormat,
	13:    ErrUnauthorized,
	14:    ErrTypeMismatch,
	15:    ErrOverflow,
	16:    ErrInvalidLength,
	17:    ErrProtocolError,
	18:    ErrAuthenticationFailed,
	19:    ErrCannotReuseObject,
	20:    ErrIllegalOperation,
	21:    ErrEmptyArrayOperation,
	22:    ErrInvalidBSON,
	23:    ErrAlreadyInitialized,
	24:    ErrLockTimeout,
	25:    ErrRemoteValidationError,
	26:    ErrNamespaceNotFound,
	27:    ErrIndexNotFound,
	28:    ErrPathNotViable,
	29:    ErrNonExistentPath,
	30:    ErrInvalidPath,
	31:    ErrRoleNotFound,
	32:    ErrRolesNotRelated,
	33:    ErrPrivilegeNotFound,
	34:    ErrCannotBackfillArray,
	35:    ErrUserModificationFailed,
	36:    ErrRemoteChangeDetected,
	37:    ErrFileRenameFailed,
	38:    ErrFileNotOpen,
	39:    ErrFileStreamFailed,
	40:    ErrConflictingUpdateOperators,
	41:    ErrFileAlreadyOpen,
	42:    ErrLogWriteFailed,
	43:    ErrCursorNotFound,
	45:    ErrUserDataInconsistent,
	46:    ErrLockBusy,
	47:    ErrNoMatchingDocument,
	48:    ErrNamespaceExists,
	49:    ErrInvalidRoleModification,
	50:    ErrExceededTimeLimit,
	51:    ErrManualInterventionRequired,
	52:    ErrDollarPrefixedFieldName,
	53:    ErrInvalidIdField,
	54:    ErrNotSingleValueField,
	55:    ErrInvalidDBRef,
	56:    ErrEmptyFieldName,
	57:    ErrDottedFieldName,
	58:    ErrRoleModificationFailed,
	59:    ErrCommandNotFound,
	60:    ErrOBSOLETE_DatabaseNotFound,
	61:    ErrShardKeyNotFound,
	62:    ErrOplogOperationUnsupported,
	63:    ErrStaleShardVersion,
	64:    ErrWriteConcernFailed,
	65:    ErrMultipleErrorsOccurred,
	66:    ErrImmutableField,
	67:    ErrCannotCreateIndex,
	68:    ErrIndexAlreadyExists,
	69:    ErrAuthSchemaIncompatible,
	70:    ErrShardNotFound,
	71:    ErrReplicaSetNotFound,
	72:    ErrInvalidOptions,
	73:    ErrInvalidNamespace,
	74:    ErrNodeNotFound,
	75:    ErrWriteConcernLegacyOK,
	76:    ErrNoReplicationEnabled,
	77:    ErrOperationIncomplete,
	78:    ErrCommandResultSchemaViolation,
	79:    ErrUnknownReplWriteConcern,
	80:    ErrRoleDataInconsistent,
	81:    ErrNoMatchParseContext,
	82:    ErrNoProgressMade,
	83:    ErrRemoteResultsUnavailable,
	84:    ErrDuplicateKeyValue,
	85:    ErrIndexOptionsConflict,
	86:    ErrIndexKeySpecsConflict,
	87:    ErrCannotSplit,
	88:    ErrSplitFailed_OBSOLETE,
	89:    ErrNetworkTimeout,
	90:    ErrCallbackCanceled,
	91:    ErrShutdownInProgress,
	92:    ErrSecondaryAheadOfPrimary,
	93:    ErrInvalidReplicaSetConfig,
	94:    ErrNotYetInitialized,
	95:    ErrNotSecondary,
	96:    ErrOperationFailed,
	97:    ErrNoProjectionFound,
	98:    ErrDBPathInUse,
	100:   ErrCannotSatisfyWriteConcern,
	101:   ErrOutdatedClient,
	102:   ErrIncompatibleAuditMetadata,
	103:   ErrNewReplicaSetConfigurationIncompatible,
	104:   ErrNodeNotElectable,
	105:   ErrIncompatibleShardingMetadata,
	106:   ErrDistributedClockSkewed,
	107:   ErrLockFailed,
	108:   ErrInconsistentReplicaSetNames,
	109:   ErrConfigurationInProgress,
	110:   ErrCannotInitializeNodeWithData,
	111:   ErrNotExactValueField,
	112:   ErrWriteConflict,
	113:   ErrInitialSyncFailure,
	114:   ErrInitialSyncOplogSourceMissing,
	115:   ErrCommandNotSupported,
	116:   ErrDocTooLargeForCapped,
	117:   ErrConflictingOperationInProgress,
	118:   ErrNamespaceNotSharded,
	119:   ErrInvalidSyncSource,
	120:   ErrOplogStartMissing,
	121:   ErrDocumentValidationFailure,
	122:   ErrOBSOLETE_ReadAfterOptimeTimeout,
	123:   ErrNotAReplicaSet,
	124:   ErrIncompatibleElectionProtocol,
	125:   ErrCommandFailed,
	126:   ErrRPCProtocolNegotiationFailed,
	127:   ErrUnrecoverableRollbackError,
	128:   ErrLockNotFound,
	129:   ErrLockStateChangeFailed,
	130:   ErrSymbolNotFound,
	131:   ErrRLPInitializationFailed,
	132:   ErrOBSOLETE_ConfigServersInconsistent,
	133:   ErrFailedToSatisfyReadPreference,
	134:   ErrReadConcernMajorityNotAvailableYet,
	135:   ErrStaleTerm,
	136:   ErrCappedPositionLost,
	137:   ErrIncompatibleShardingConfigVersion,
	138:   ErrRemoteOplogStale,
	139:   ErrJSInterpreterFailure,
	140:   ErrInvalidSSLConfiguration,
	141:   ErrSSLHandshakeFailed,
	142:   ErrJSUncatchableError,
	143:   ErrCursorInUse,
	144:   ErrIncompatibleCatalogManager,
	145:   ErrPooledConnectionsDropped,
	146:   ErrExceededMemoryLimit,
	147:   ErrZLibError,
	148:   ErrReadConcernMajorityNotEnabled,
	149:   ErrNoConfigMaster,
	150:   ErrStaleEpoch,
	151:   ErrOperationCannotBeBatched,
	152:   ErrOplogOutOfOrder,
	153:   ErrChunkTooBig,
	154:   ErrInconsistentShardIdentity,
	155:   ErrCannotApplyOplogWhilePrimary,
	156:   ErrNeedsDocumentMove,
	157:   ErrCanRepairToDowngrade,
	158:   ErrMustUpgrade,
	159:   ErrDurationOverflow,
	160:   ErrMaxStalenessOutOfRange,
	161:   ErrIncompatibleCollationVersion,
	162:   ErrCollectionIsEmpty,
	163:   ErrZoneStillInUse,
	164:   ErrInitialSyncActive,
	165:   ErrViewDepthLimitExceeded,
	166:   ErrCommandNotSupportedOnView,
	167:   ErrOptionNotSupportedOnView,
	168:   ErrInvalidPipelineOperator,
	169:   ErrCommandOnShardedViewNotSupportedOnMongod,
	170:   ErrTooManyMatchingDocuments,
	171:   ErrCannotIndexParallelArrays,
	172:   ErrTransportSessionClosed,
	173:   ErrTransportSessionNotFound,
	174:   ErrTransportSessionUnknown,
	175:   ErrQueryPlanKilled,
	176:   ErrFileOpenFailed,
	177:   ErrZoneNotFound,
	178:   ErrRangeOverlapConflict,
	179:   ErrWindowsPdhError,
	180:   ErrBadPerfCounterPath,
	181:   ErrAmbiguousIndexKeyPattern,
	182:   ErrInvalidViewDefinition,
	183:   ErrClientMetadataMissingField,
	184:   ErrClientMetadataAppNameTooLarge,
	185:   ErrClientMetadataDocumentTooLarge,
	186:   ErrClientMetadataCannotBeMutated,
	187:   ErrLinearizableReadConcernError,
	188:   ErrIncompatibleServerVersion,
	189:   ErrPrimarySteppedDown,
	190:   ErrMasterSlaveConnectionFailure,
	191:   ErrOBSOLETE_BalancerLostDistributedLock,
	192:   ErrFailPointEnabled,
	193:   ErrNoShardingEnabled,
	194:   ErrBalancerInterrupted,
	195:   ErrViewPipelineMaxSizeExceeded,
	197:   ErrInvalidIndexSpecificationOption,
	198:   ErrOBSOLETE_ReceivedOpReplyMessage,
	199:   ErrReplicaSetMonitorRemoved,
	200:   ErrChunkRangeCleanupPending,
	201:   ErrCannotBuildIndexKeys,
	202:   ErrNetworkInterfaceExceededTimeLimit,
	203:   ErrShardingStateNotInitialized,
	204:   ErrTimeProofMismatch,
	205:   ErrClusterTimeFailsRateLimiter,
	206:   ErrNoSuchSession,
	207:   ErrInvalidUUID,
	208:   ErrTooManyLocks,
	209:   ErrStaleClusterTime,
	210:   ErrCannotVerifyAndSignLogicalTime,
	211:   ErrKeyNotFound,
	212:   ErrIncompatibleRollbackAlgorithm,
	213:   ErrDuplicateSession,
	214:   ErrAuthenticationRestrictionUnmet,
	215:   ErrDatabaseDropPending,
	216:   ErrElectionInProgress,
	217:   ErrIncompleteTransactionHistory,
	218:   ErrUpdateOperationFailed,
	219:   ErrFTDCPathNotSet,
	220:   ErrFTDCPathAlreadySet,
	221:   ErrIndexModified,
	222:   ErrCloseChangeStream,
	223:   ErrIllegalOpMsgFlag,
	224:   ErrQueryFeatureNotAllowed,
	225:   ErrTransactionTooOld,
	226:   ErrAtomicityFailure,
	227:   ErrCannotImplicitlyCreateCollection,
	228:   ErrSessionTransferIncomplete,
	229:   ErrMustDowngrade,
	230:   ErrDNSHostNotFound,
	231:   ErrDNSProtocolError,
	232:   ErrMaxSubPipelineDepthExceeded,
	233:   ErrTooManyDocumentSequences,
	234:   ErrRetryChangeStream,
	235:   ErrInternalErrorNotSupported,
	236:   ErrForTestingErrorExtraInfo,
	237:   ErrCursorKilled,
	238:   ErrNotImplemented,
	239:   ErrSnapshotTooOld,
	240:   ErrDNSRecordTypeMismatch,
	241:   ErrConversionFailure,
	242:   ErrCannotCreateCollection,
	243:   ErrIncompatibleWithUpgradedServer,
	244:   ErrNOT_YET_AVAILABLE_TransactionAborted,
	245:   ErrBrokenPromise,
	246:   ErrSnapshotUnavailable,
	247:   ErrProducerConsumerQueueBatchTooLarge,
	248:   ErrProducerConsumerQueueEndClosed,
	249:   ErrStaleDbVersion,
	250:   ErrStaleChunkHistory,
	251:   ErrNoSuchTransaction,
	252:   ErrReentrancyNotAllowed,
	253:   ErrFreeMonHttpInFlight,
	254:   ErrFreeMonHttpTemporaryFailure,
	255:   ErrFreeMonHttpPermanentFailure,
	256:   ErrTransactionCommitted,
	257:   ErrTransactionTooLarge,
	258:   ErrUnknownFeatureCompatibilityVersion,
	9001:  ErrSocketException,
	9996:  ErrOBSOLETE_RecvStaleConfig,
	10107: ErrNotMaster,
	10003: ErrCannotGrowDocumentInCappedNamespace,
	10334: ErrBSONObjectTooLarge,
	11000: ErrDuplicateKey,
	11600: ErrInterruptedAtShutdown,
	11601: ErrInterrupted,
	11602: ErrInterruptedDueToReplStateChange,
	14031: ErrOutOfDiskSpace,
	17280: ErrKeyTooLong,
	12586: ErrBackgroundOperationInProgressForDatabase,
	12587: ErrBackgroundOperationInProgressForNamespace,
	13436: ErrNotMasterOrSecondary,
	13435: ErrNotMasterNoSlaveOk,
	13334: ErrShardKeyTooBig,
	13388: ErrStaleConfig,
	13297: ErrDatabaseDifferCase,
	13104: ErrOBSOLETE_PrepareConfigsFailed,
}
