//go:build cgo && go1.24

// Copyright 2019, 2025 The Godror Authors
//
//
// SPDX-License-Identifier: UPL-1.0 OR Apache-2.0

package godror

// See https://github.com/golang/go/issues/56378
// For pain, see https://github.com/zhanbolat18/godror/issues/365 .

/*
#cgo nocallback dpiConn_breakExecution
#cgo nocallback dpiConn_commit
#cgo nocallback dpiConn_create
#cgo nocallback dpiConn_getCurrentSchema
#cgo nocallback dpiConn_getDbDomain
#cgo nocallback dpiConn_getDbName
#cgo nocallback dpiConn_getEdition
#cgo nocallback dpiConn_getIsHealthy
#cgo nocallback dpiConn_getObjectType
#cgo nocallback dpiConn_getServerVersion
#cgo nocallback dpiConn_getServiceName
#cgo nocallback dpiConn_newMsgProps
#cgo nocallback dpiConn_newQueue
#cgo nocallback dpiConn_newTempLob
#cgo nocallback dpiConn_newVar
#cgo nocallback dpiConn_ping
#cgo nocallback dpiConn_prepareStmt
#cgo nocallback dpiConn_release
#cgo nocallback dpiConn_rollback
#cgo nocallback dpiConn_setAction
#cgo nocallback dpiConn_setCallTimeout
// #cgo nocallback dpiConn_setClientIdentifier
#cgo nocallback dpiConn_setClientInfo
#cgo nocallback dpiConn_setCurrentSchema
#cgo nocallback dpiConn_setDbOp
#cgo nocallback dpiConn_setModule
#cgo nocallback dpiConn_shutdownDatabase
#cgo nocallback dpiConn_startupDatabase
#cgo nocallback dpiContext_createWithParams
#cgo nocallback dpiContext_destroy
#cgo nocallback dpiContext_getClientVersion
#cgo nocallback dpiContext_getError
#cgo nocallback dpiContext_initCommonCreateParams
#cgo nocallback dpiContext_initConnCreateParams
#cgo nocallback dpiContext_initPoolCreateParams
#cgo nocallback dpiContext_initSubscrCreateParams
// #cgo nocallback dpiData_getBool
// #cgo nocallback dpiData_getBytes
// #cgo nocallback dpiData_getDouble
// #cgo nocallback dpiData_getFloat
// #cgo nocallback dpiData_getInt64
// #cgo nocallback dpiData_getIntervalDS
// #cgo nocallback dpiData_getIntervalYM
// #cgo nocallback dpiData_getIsNull
#cgo nocallback dpiData_getJson
#cgo nocallback dpiData_getJsonArray
#cgo nocallback dpiData_getJsonObject
#cgo nocallback dpiData_getLOB
#cgo nocallback dpiData_getObject
// #cgo nocallback dpiData_getRowid
#cgo nocallback dpiData_getStmt
// #cgo nocallback dpiData_getTimestamp
// #cgo nocallback dpiData_getUint64
#cgo nocallback dpiData_setBool
#cgo nocallback dpiData_setBytes
#cgo nocallback dpiData_setDouble
#cgo nocallback dpiData_setFloat
#cgo nocallback dpiData_setInt64
#cgo nocallback dpiData_setIntervalDS
#cgo nocallback dpiData_setIntervalYM
#cgo nocallback dpiData_setLOB
// #cgo nocallback dpiData_setNull
#cgo nocallback dpiData_setObject
#cgo nocallback dpiData_setStmt
#cgo nocallback dpiData_setTimestamp
#cgo nocallback dpiData_setUint64
#cgo nocallback dpiDeqOptions_getCondition
#cgo nocallback dpiDeqOptions_getConsumerName
#cgo nocallback dpiDeqOptions_getCorrelation
#cgo nocallback dpiDeqOptions_getMode
#cgo nocallback dpiDeqOptions_getMsgId
#cgo nocallback dpiDeqOptions_getNavigation
#cgo nocallback dpiDeqOptions_getTransformation
#cgo nocallback dpiDeqOptions_getVisibility
#cgo nocallback dpiDeqOptions_getWait
#cgo nocallback dpiDeqOptions_setCondition
#cgo nocallback dpiDeqOptions_setConsumerName
#cgo nocallback dpiDeqOptions_setCorrelation
#cgo nocallback dpiDeqOptions_setDeliveryMode
#cgo nocallback dpiDeqOptions_setMode
#cgo nocallback dpiDeqOptions_setMsgId
#cgo nocallback dpiDeqOptions_setNavigation
#cgo nocallback dpiDeqOptions_setTransformation
#cgo nocallback dpiDeqOptions_setVisibility
#cgo nocallback dpiDeqOptions_setWait
#cgo nocallback dpiEnqOptions_getTransformation
#cgo nocallback dpiEnqOptions_getVisibility
#cgo nocallback dpiEnqOptions_setDeliveryMode
#cgo nocallback dpiEnqOptions_setTransformation
#cgo nocallback dpiEnqOptions_setVisibility
#cgo nocallback dpiJson_setFromText
#cgo nocallback dpiJson_setValue
#cgo nocallback dpiLob_close
#cgo nocallback dpiLob_closeResource
#cgo nocallback dpiLob_getChunkSize
#cgo nocallback dpiLob_getDirectoryAndFileName
#cgo nocallback dpiLob_getIsResourceOpen
#cgo nocallback dpiLob_getSize
#cgo nocallback dpiLob_getType
#cgo nocallback dpiLob_openResource
#cgo nocallback dpiLob_readBytes
#cgo nocallback dpiLob_release
#cgo nocallback dpiLob_setFromBytes
#cgo nocallback dpiLob_trim
#cgo nocallback dpiLob_writeBytes
#cgo nocallback dpiMsgProps_getCorrelation
#cgo nocallback dpiMsgProps_getDelay
#cgo nocallback dpiMsgProps_getDeliveryMode
#cgo nocallback dpiMsgProps_getEnqTime
#cgo nocallback dpiMsgProps_getExceptionQ
#cgo nocallback dpiMsgProps_getExpiration
#cgo nocallback dpiMsgProps_getMsgId
#cgo nocallback dpiMsgProps_getNumAttempts
#cgo nocallback dpiMsgProps_getOriginalMsgId
#cgo nocallback dpiMsgProps_getPayload
#cgo nocallback dpiMsgProps_getPriority
#cgo nocallback dpiMsgProps_getState
#cgo nocallback dpiMsgProps_release
#cgo nocallback dpiMsgProps_setCorrelation
#cgo nocallback dpiMsgProps_setDelay
#cgo nocallback dpiMsgProps_setExceptionQ
#cgo nocallback dpiMsgProps_setExpiration
#cgo nocallback dpiMsgProps_setOriginalMsgId
#cgo nocallback dpiMsgProps_setPayloadBytes
#cgo nocallback dpiMsgProps_setPayloadObject
#cgo nocallback dpiMsgProps_setPriority
#cgo nocallback dpiObject_addRef
#cgo nocallback dpiObject_appendElement
#cgo nocallback dpiObjectAttr_getInfo
#cgo nocallback dpiObjectAttr_release
#cgo nocallback dpiObject_deleteElementByIndex
#cgo nocallback dpiObject_getAttributeValue
#cgo nocallback dpiObject_getElementExistsByIndex
#cgo nocallback dpiObject_getElementValueByIndex
#cgo nocallback dpiObject_getFirstIndex
#cgo nocallback dpiObject_getLastIndex
#cgo nocallback dpiObject_getNextIndex
#cgo nocallback dpiObject_getSize
#cgo nocallback dpiObject_release
#cgo nocallback dpiObject_setAttributeValue
#cgo nocallback dpiObject_setElementValueByIndex
#cgo nocallback dpiObject_trim
#cgo nocallback dpiObjectType_addRef
#cgo nocallback dpiObjectType_createObject
#cgo nocallback dpiObjectType_getAttributes
#cgo nocallback dpiObjectType_getInfo
#cgo nocallback dpiObjectType_release
#cgo nocallback dpiPool_close
#cgo nocallback dpiPool_create
#cgo nocallback dpiPool_getBusyCount
#cgo nocallback dpiPool_getMaxLifetimeSession
#cgo nocallback dpiPool_getOpenCount
#cgo nocallback dpiPool_getTimeout
#cgo nocallback dpiPool_getWaitTimeout
#cgo nocallback dpiPool_release
#cgo nocallback dpiPool_setStmtCacheSize
#cgo nocallback dpiQueue_deqMany
#cgo nocallback dpiQueue_deqOne
#cgo nocallback dpiQueue_enqMany
#cgo nocallback dpiQueue_enqOne
#cgo nocallback dpiQueue_getDeqOptions
#cgo nocallback dpiQueue_getEnqOptions
#cgo nocallback dpiQueue_release
#cgo nocallback dpiRowid_getStringValue
#cgo nocallback dpiStmt_addRef
#cgo nocallback dpiStmt_bindByName
#cgo nocallback dpiStmt_bindByPos
#cgo nocallback dpiStmt_define
#cgo nocallback dpiStmt_deleteFromCache
#cgo nocallback dpiStmt_execute
#cgo nocallback dpiStmt_executeMany
#cgo nocallback dpiStmt_fetchRows
#cgo nocallback dpiStmt_getBatchErrorCount
#cgo nocallback dpiStmt_getBatchErrors
#cgo nocallback dpiStmt_getBindCount
#cgo nocallback dpiStmt_getBindNames
#cgo nocallback dpiStmt_getImplicitResult
#cgo nocallback dpiStmt_getInfo
#cgo nocallback dpiStmt_getNumQueryColumns
#cgo nocallback dpiStmt_getQueryInfo
#cgo nocallback dpiStmt_getRowCount
#cgo nocallback dpiStmt_getSubscrQueryId
#cgo nocallback dpiStmt_release
#cgo nocallback dpiStmt_setFetchArraySize
#cgo nocallback dpiStmt_setPrefetchRows
#cgo nocallback dpiVar_getNumElementsInArray
#cgo nocallback dpiVar_getReturnedData
#cgo nocallback dpiVar_release
#cgo nocallback dpiVar_setFromBytes
#cgo nocallback dpiVar_setFromJson
#cgo nocallback dpiVar_setFromLob
#cgo nocallback dpiVar_setFromObject
#cgo nocallback dpiVar_setNumElementsInArray
#cgo nocallback godror_allocate_dpiNode
#cgo nocallback godror_dpiasJsonArray
#cgo nocallback godror_dpiasJsonObject
#cgo nocallback godror_dpiJsonArray_initialize
#cgo nocallback godror_dpiJsonfreeMem
#cgo nocallback godror_dpiJsonObject_initialize
#cgo nocallback godror_dpiJsonObject_setKey
#cgo nocallback godror_dpiJson_setBool
#cgo nocallback godror_dpiJson_setBytes
#cgo nocallback godror_dpiJson_setDouble
#cgo nocallback godror_dpiJson_setInt64
#cgo nocallback godror_dpiJson_setIntervalDS
#cgo nocallback godror_dpiJson_setNumber
#cgo nocallback godror_dpiJson_setString
#cgo nocallback godror_dpiJson_setTime
#cgo nocallback godror_dpiJson_setUint64
#cgo nocallback godror_getAnnotation
#cgo nocallback godror_setArrayElements
#cgo nocallback godror_setFromString
#cgo nocallback godror_setObjectFields
*/
import "C"
