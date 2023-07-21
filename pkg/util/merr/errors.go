// Licensed to the LF AI & Data foundation under one
// or more contributor license agreements. See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership. The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License. You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package merr

import (
	"context"

	"github.com/cockroachdb/errors"
)

const (
	rootCoordBits = (iota + 1) << 16
	dataCoordBits
	queryCoordBits
	dataNodeBits
	queryNodeBits
	indexNodeBits
	proxyBits
	standaloneBits
	embededBits

	retriableFlag      = 1 << 20
	RootReasonCodeMask = (1 << 16) - 1

	CanceledCode int32 = 10000
	TimeoutCode  int32 = 10001
)

// Define leaf errors here,
// WARN: take care to add new error,
// check whehter you can use the erorrs below before adding a new one.
// Name: Err + related prefix + error name
var (
	// Service related
	ErrServiceNotReady             = newMilvusError("service not ready", 1, true) // This indicates the service is still in init
	ErrServiceUnavailable          = newMilvusError("service unavailable", 2, true)
	ErrServiceMemoryLimitExceeded  = newMilvusError("memory limit exceeded", 3, false)
	ErrServiceRequestLimitExceeded = newMilvusError("request limit exceeded", 4, true)
	ErrServiceInternal             = newMilvusError("service internal error", 5, false) // Never return this error out of Milvus
	ErrCrossClusterRouting         = newMilvusError("cross cluster routing", 6, false)

	// Collection related
	ErrCollectionNotFound         = newMilvusError("collection not found", 100, false)
	ErrCollectionNotLoaded        = newMilvusError("collection not loaded", 101, false)
	ErrCollectionNumLimitExceeded = newMilvusError("exceeded the limit number of collections", 102, false)
	ErrCollectionNotFullyLoaded   = newMilvusError("collection not fully loaded", 103, true)

	// Partition related
	ErrPartitionNotFound       = newMilvusError("partition not found", 202, false)
	ErrPartitionNotLoaded      = newMilvusError("partition not loaded", 203, false)
	ErrPartitionNotFullyLoaded = newMilvusError("collection not fully loaded", 103, true)

	// ResourceGroup related
	ErrResourceGroupNotFound = newMilvusError("resource group not found", 300, false)

	// Replica related
	ErrReplicaNotFound     = newMilvusError("replica not found", 400, false)
	ErrReplicaNotAvailable = newMilvusError("replica not available", 401, false)

	// Channel related
	ErrChannelNotFound     = newMilvusError("channel not found", 500, false)
	ErrChannelLack         = newMilvusError("channel lacks", 501, false)
	ErrChannelReduplicate  = newMilvusError("channel reduplicates", 502, false)
	ErrChannelNotAvailable = newMilvusError("channel not available", 503, false)

	// Segment related
	ErrSegmentNotFound    = newMilvusError("segment not found", 600, false)
	ErrSegmentNotLoaded   = newMilvusError("segment not loaded", 601, false)
	ErrSegmentLack        = newMilvusError("segment lacks", 602, false)
	ErrSegmentReduplicate = newMilvusError("segment reduplicates", 603, false)

	// Index related
	ErrIndexNotFound = newMilvusError("index not found", 700, false)

	// Database related
	ErrDatabaseNotfound         = newMilvusError("database not found", 800, false)
	ErrDatabaseNumLimitExceeded = newMilvusError("exceeded the limit number of database", 801, false)
	ErrInvalidedDatabaseName    = newMilvusError("invalided database name", 802, false)

	// Node related
	ErrNodeNotFound     = newMilvusError("node not found", 901, false)
	ErrNodeOffline      = newMilvusError("node offline", 902, false)
	ErrNodeLack         = newMilvusError("node lacks", 903, false)
	ErrNodeNotMatch     = newMilvusError("node not match", 904, false)
	ErrNodeNotAvailable = newMilvusError("node not available", 905, false)

	// IO related
	ErrIoKeyNotFound = newMilvusError("key not found", 1000, false)
	ErrIoFailed      = newMilvusError("IO failed", 1001, false)

	// Parameter related
	ErrParameterInvalid = newMilvusError("invalid parameter", 1100, false)

	// Metrics related
	ErrMetricNotFound = newMilvusError("metric not found", 1200, false)

	// Topic related
	ErrTopicNotFound = newMilvusError("topic not found", 1300, false)
	ErrTopicNotEmpty = newMilvusError("topic not empty", 1301, false)

	// shard delegator related
	ErrShardDelegatorNotFound        = newMilvusError("shard delegator not found", 1500, false)
	ErrShardDelegatorAccessFailed    = newMilvusError("fail to access shard delegator", 1501, true)
	ErrShardDelegatorSearchFailed    = newMilvusError("fail to search on all shard leaders", 1502, true)
	ErrShardDelegatorQueryFailed     = newMilvusError("fail to query on all shard leaders", 1503, true)
	ErrShardDelegatorStatisticFailed = newMilvusError("get statistics on all shard leaders", 1504, true)

	// field related
	ErrFieldNotFound = newMilvusError("field not found", 1700, false)

	// high-level restful api related
	ErrNeedAuthenticate          = newMilvusError("user hasn't authenticate", 1800, false)
	ErrIncorrectParameterFormat  = newMilvusError("can only accept json format request", 1801, false)
	ErrMissingRequiredParameters = newMilvusError("missing required parameters", 1802, false)
	ErrMarshalCollectionSchema   = newMilvusError("fail to marshal collection schema", 1803, false)
	ErrInvalidInsertData         = newMilvusError("fail to deal the insert data", 1804, false)
	ErrInvalidSearchResult       = newMilvusError("fail to parse search result", 1805, false)
	ErrCheckPrimaryKey           = newMilvusError("please check the primary key and its' type can only in [int, string]", 1806, false)

	// Do NOT export this,
	// never allow programmer using this, keep only for converting unknown error to milvusError
	errUnexpected = newMilvusError("unexpected error", (1<<16)-1, false)
)

func maskComponentBits(code int32) int32 {
	return code | proxyBits
}

type milvusError struct {
	msg     string
	errCode int32
}

func newMilvusError(msg string, code int32, retriable bool) milvusError {
	if retriable {
		code |= retriableFlag
	}
	return milvusError{
		msg:     msg,
		errCode: code,
	}
}

func (e milvusError) code() int32 {
	return maskComponentBits(e.errCode)
}

func (e milvusError) Error() string {
	return e.msg
}

// Code returns the error code of the given error,
// WARN: DO NOT use this for now
func Code(err error) int32 {
	if err == nil {
		return 0
	}

	cause := errors.Cause(err)
	switch cause := cause.(type) {
	case milvusError:
		return cause.code()

	default:
		if errors.Is(cause, context.Canceled) {
			return CanceledCode
		} else if errors.Is(cause, context.DeadlineExceeded) {
			return TimeoutCode
		} else {
			return errUnexpected.code()
		}
	}
}