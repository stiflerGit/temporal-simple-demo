// The MIT License
//
// Copyright (c) 2022 Temporal Technologies Inc.  All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: temporal/api/enums/v1/query.proto

package enums

import (
	fmt "fmt"
	math "math"
	strconv "strconv"

	proto "github.com/gogo/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type QueryResultType int32

const (
	QUERY_RESULT_TYPE_UNSPECIFIED QueryResultType = 0
	QUERY_RESULT_TYPE_ANSWERED    QueryResultType = 1
	QUERY_RESULT_TYPE_FAILED      QueryResultType = 2
)

var QueryResultType_name = map[int32]string{
	0: "Unspecified",
	1: "Answered",
	2: "Failed",
}

var QueryResultType_value = map[string]int32{
	"Unspecified": 0,
	"Answered":    1,
	"Failed":      2,
}

func (QueryResultType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b9a616a97224ce1d, []int{0}
}

type QueryRejectCondition int32

const (
	QUERY_REJECT_CONDITION_UNSPECIFIED QueryRejectCondition = 0
	// None indicates that query should not be rejected.
	QUERY_REJECT_CONDITION_NONE QueryRejectCondition = 1
	// NotOpen indicates that query should be rejected if workflow is not open.
	QUERY_REJECT_CONDITION_NOT_OPEN QueryRejectCondition = 2
	// NotCompletedCleanly indicates that query should be rejected if workflow did not complete cleanly.
	QUERY_REJECT_CONDITION_NOT_COMPLETED_CLEANLY QueryRejectCondition = 3
)

var QueryRejectCondition_name = map[int32]string{
	0: "Unspecified",
	1: "None",
	2: "NotOpen",
	3: "NotCompletedCleanly",
}

var QueryRejectCondition_value = map[string]int32{
	"Unspecified":         0,
	"None":                1,
	"NotOpen":             2,
	"NotCompletedCleanly": 3,
}

func (QueryRejectCondition) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b9a616a97224ce1d, []int{1}
}

func init() {
	proto.RegisterEnum("temporal.api.enums.v1.QueryResultType", QueryResultType_name, QueryResultType_value)
	proto.RegisterEnum("temporal.api.enums.v1.QueryRejectCondition", QueryRejectCondition_name, QueryRejectCondition_value)
}

func init() { proto.RegisterFile("temporal/api/enums/v1/query.proto", fileDescriptor_b9a616a97224ce1d) }

var fileDescriptor_b9a616a97224ce1d = []byte{
	// 375 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0xd1, 0x4f, 0x4b, 0xe3, 0x40,
	0x18, 0x06, 0xf0, 0x4c, 0x16, 0xf6, 0x30, 0x97, 0x0d, 0x61, 0x0b, 0x65, 0xff, 0x4c, 0xe9, 0x2e,
	0xec, 0xa1, 0x2c, 0x89, 0xc1, 0x5b, 0x3c, 0xa5, 0xc9, 0x5b, 0x88, 0xc4, 0x49, 0x9a, 0x4e, 0x2b,
	0xf5, 0x12, 0xaa, 0x06, 0x19, 0x69, 0x9b, 0x98, 0xa6, 0x85, 0xde, 0xfc, 0x08, 0x7e, 0x0c, 0xf1,
	0xe8, 0xc1, 0xcf, 0xe0, 0xb1, 0xc7, 0x1e, 0x6d, 0x7a, 0x11, 0x4f, 0xfd, 0x08, 0xd2, 0x60, 0x40,
	0xad, 0xf5, 0x36, 0xcc, 0xfb, 0x9b, 0x79, 0xe1, 0x79, 0x70, 0x35, 0x0d, 0x07, 0x71, 0x94, 0xf4,
	0xfa, 0x6a, 0x2f, 0xe6, 0x6a, 0x38, 0x1c, 0x0f, 0x46, 0xea, 0x44, 0x53, 0x2f, 0xc6, 0x61, 0x32,
	0x55, 0xe2, 0x24, 0x4a, 0x23, 0xb9, 0x54, 0x10, 0xa5, 0x17, 0x73, 0x25, 0x27, 0xca, 0x44, 0xab,
	0x25, 0xf8, 0x5b, 0x73, 0xad, 0xfc, 0x70, 0x34, 0xee, 0xa7, 0x6c, 0x1a, 0x87, 0x72, 0x15, 0xff,
	0x6e, 0xb6, 0xc1, 0xef, 0x06, 0x3e, 0xb4, 0xda, 0x0e, 0x0b, 0x58, 0xd7, 0x83, 0xa0, 0x4d, 0x5b,
	0x1e, 0x98, 0x76, 0xc3, 0x06, 0x4b, 0x12, 0x64, 0x82, 0x7f, 0x6c, 0x12, 0x83, 0xb6, 0x0e, 0xc1,
	0x07, 0x4b, 0x42, 0xf2, 0x2f, 0x5c, 0xde, 0x9c, 0x37, 0x0c, 0xdb, 0x01, 0x4b, 0x12, 0x6b, 0x77,
	0x08, 0x7f, 0x7f, 0x59, 0x7a, 0x1e, 0x9e, 0xa4, 0x66, 0x34, 0x3c, 0xe5, 0x29, 0x8f, 0x86, 0xf2,
	0x3f, 0xfc, 0xa7, 0x78, 0xb6, 0x0f, 0x26, 0x0b, 0x4c, 0x97, 0x5a, 0x36, 0xb3, 0x5d, 0xfa, 0x6e,
	0x7d, 0x05, 0xff, 0xdc, 0xe2, 0xa8, 0x4b, 0x41, 0x42, 0xf2, 0x5f, 0x5c, 0xd9, 0x0a, 0x58, 0xe0,
	0x7a, 0x40, 0x25, 0x51, 0xde, 0xc1, 0xff, 0x3f, 0x41, 0xa6, 0x7b, 0xe0, 0x39, 0xc0, 0xc0, 0x0a,
	0x4c, 0x07, 0x0c, 0xea, 0x74, 0xa5, 0x2f, 0xf5, 0x5b, 0x34, 0x5b, 0x10, 0x61, 0xbe, 0x20, 0xc2,
	0x6a, 0x41, 0xd0, 0x65, 0x46, 0xd0, 0x75, 0x46, 0xd0, 0x7d, 0x46, 0xd0, 0x2c, 0x23, 0xe8, 0x21,
	0x23, 0xe8, 0x31, 0x23, 0xc2, 0x2a, 0x23, 0xe8, 0x6a, 0x49, 0x84, 0xd9, 0x92, 0x08, 0xf3, 0x25,
	0x11, 0x70, 0x99, 0x47, 0xca, 0x87, 0xe9, 0xd7, 0x71, 0x1e, 0x83, 0xb7, 0x2e, 0xc8, 0x43, 0x47,
	0xd5, 0xb3, 0x57, 0x8e, 0x47, 0x6f, 0xba, 0xdc, 0xcb, 0x0f, 0x37, 0x62, 0x89, 0x15, 0xc0, 0x88,
	0xb9, 0x02, 0xf9, 0x47, 0x1d, 0xed, 0x49, 0x2c, 0x17, 0xf7, 0xba, 0x6e, 0xc4, 0x5c, 0xd7, 0xf3,
	0x89, 0xae, 0x77, 0xb4, 0xe3, 0xaf, 0x79, 0xff, 0xbb, 0xcf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x78,
	0x9d, 0x45, 0x60, 0x24, 0x02, 0x00, 0x00,
}

func (x QueryResultType) String() string {
	s, ok := QueryResultType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x QueryRejectCondition) String() string {
	s, ok := QueryRejectCondition_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}