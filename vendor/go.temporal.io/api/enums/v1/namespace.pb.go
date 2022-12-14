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
// source: temporal/api/enums/v1/namespace.proto

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

type NamespaceState int32

const (
	NAMESPACE_STATE_UNSPECIFIED NamespaceState = 0
	NAMESPACE_STATE_REGISTERED  NamespaceState = 1
	NAMESPACE_STATE_DEPRECATED  NamespaceState = 2
	NAMESPACE_STATE_DELETED     NamespaceState = 3
)

var NamespaceState_name = map[int32]string{
	0: "Unspecified",
	1: "Registered",
	2: "Deprecated",
	3: "Deleted",
}

var NamespaceState_value = map[string]int32{
	"Unspecified": 0,
	"Registered":  1,
	"Deprecated":  2,
	"Deleted":     3,
}

func (NamespaceState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d8c23ee067d676e7, []int{0}
}

type ArchivalState int32

const (
	ARCHIVAL_STATE_UNSPECIFIED ArchivalState = 0
	ARCHIVAL_STATE_DISABLED    ArchivalState = 1
	ARCHIVAL_STATE_ENABLED     ArchivalState = 2
)

var ArchivalState_name = map[int32]string{
	0: "Unspecified",
	1: "Disabled",
	2: "Enabled",
}

var ArchivalState_value = map[string]int32{
	"Unspecified": 0,
	"Disabled":    1,
	"Enabled":     2,
}

func (ArchivalState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d8c23ee067d676e7, []int{1}
}

type ReplicationState int32

const (
	REPLICATION_STATE_UNSPECIFIED ReplicationState = 0
	REPLICATION_STATE_NORMAL      ReplicationState = 1
	REPLICATION_STATE_HANDOVER    ReplicationState = 2
)

var ReplicationState_name = map[int32]string{
	0: "Unspecified",
	1: "Normal",
	2: "Handover",
}

var ReplicationState_value = map[string]int32{
	"Unspecified": 0,
	"Normal":      1,
	"Handover":    2,
}

func (ReplicationState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d8c23ee067d676e7, []int{2}
}

func init() {
	proto.RegisterEnum("temporal.api.enums.v1.NamespaceState", NamespaceState_name, NamespaceState_value)
	proto.RegisterEnum("temporal.api.enums.v1.ArchivalState", ArchivalState_name, ArchivalState_value)
	proto.RegisterEnum("temporal.api.enums.v1.ReplicationState", ReplicationState_name, ReplicationState_value)
}

func init() {
	proto.RegisterFile("temporal/api/enums/v1/namespace.proto", fileDescriptor_d8c23ee067d676e7)
}

var fileDescriptor_d8c23ee067d676e7 = []byte{
	// 391 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xc1, 0x8a, 0xda, 0x40,
	0x18, 0xc7, 0x33, 0x29, 0xf4, 0x30, 0xd0, 0x12, 0x02, 0x56, 0xd1, 0x76, 0x8a, 0x87, 0x5e, 0x3c,
	0x24, 0x84, 0xde, 0xd2, 0xd3, 0x98, 0x7c, 0xad, 0x81, 0x18, 0xc3, 0x24, 0xcd, 0xa1, 0x17, 0x99,
	0x4a, 0xa8, 0x03, 0x6a, 0x06, 0x4d, 0x3d, 0xf7, 0x09, 0x4a, 0x1f, 0xa3, 0xf4, 0x01, 0xfa, 0x0c,
	0x7b, 0xf4, 0xe8, 0x71, 0x8d, 0x97, 0x65, 0x4f, 0x3e, 0xc2, 0x62, 0x42, 0xd8, 0x5d, 0xcd, 0x6d,
	0x98, 0xdf, 0x8f, 0xf9, 0xfe, 0xc3, 0xff, 0xc3, 0x1f, 0xf2, 0x74, 0x29, 0xb3, 0x35, 0x5f, 0x98,
	0x5c, 0x0a, 0x33, 0x5d, 0xfd, 0x5c, 0x6e, 0xcc, 0xad, 0x65, 0xae, 0xf8, 0x32, 0xdd, 0x48, 0x3e,
	0x4b, 0x0d, 0xb9, 0xce, 0xf2, 0x4c, 0x6f, 0xd5, 0x9a, 0xc1, 0xa5, 0x30, 0x4a, 0xcd, 0xd8, 0x5a,
	0x83, 0xdf, 0x08, 0xbf, 0x0e, 0x6a, 0x35, 0xca, 0x79, 0x9e, 0xea, 0xef, 0x71, 0x2f, 0xa0, 0x63,
	0x88, 0x42, 0xea, 0xc0, 0x34, 0x8a, 0x69, 0x0c, 0xd3, 0xaf, 0x41, 0x14, 0x82, 0xe3, 0x7d, 0xf6,
	0xc0, 0xd5, 0x14, 0x9d, 0xe0, 0xee, 0xa5, 0xc0, 0xe0, 0x8b, 0x17, 0xc5, 0xc0, 0xc0, 0xd5, 0x50,
	0x13, 0x77, 0x21, 0x64, 0xe0, 0xd0, 0x18, 0x5c, 0x4d, 0xd5, 0x7b, 0xb8, 0x7d, 0xcd, 0x7d, 0x38,
	0xc3, 0x17, 0x83, 0x39, 0x7e, 0x45, 0xd7, 0xb3, 0xb9, 0xd8, 0xf2, 0x45, 0x15, 0x87, 0xe0, 0x2e,
	0x65, 0xce, 0xc8, 0x4b, 0xa8, 0xdf, 0x98, 0xa6, 0x87, 0xdb, 0x17, 0xdc, 0xf5, 0x22, 0x3a, 0xf4,
	0xcb, 0x28, 0x5d, 0xfc, 0xe6, 0x02, 0x42, 0x50, 0x31, 0x75, 0xb0, 0xc1, 0x1a, 0x4b, 0xe5, 0x42,
	0xcc, 0x78, 0x2e, 0xb2, 0x55, 0x35, 0xac, 0x8f, 0xdf, 0x31, 0x08, 0x7d, 0xcf, 0xa1, 0xb1, 0x37,
	0x09, 0x1a, 0xe7, 0xbd, 0xc5, 0x9d, 0x6b, 0x25, 0x98, 0xb0, 0x31, 0xf5, 0xab, 0xbf, 0x5f, 0xd3,
	0x11, 0x0d, 0xdc, 0x49, 0x02, 0x4c, 0x53, 0x87, 0xff, 0xd1, 0xee, 0x40, 0x94, 0xfd, 0x81, 0x28,
	0xa7, 0x03, 0x41, 0xbf, 0x0a, 0x82, 0xfe, 0x16, 0x04, 0xdd, 0x14, 0x04, 0xed, 0x0a, 0x82, 0x6e,
	0x0b, 0x82, 0xee, 0x0a, 0xa2, 0x9c, 0x0a, 0x82, 0xfe, 0x1c, 0x89, 0xb2, 0x3b, 0x12, 0x65, 0x7f,
	0x24, 0x0a, 0xee, 0x88, 0xcc, 0x68, 0x2c, 0x70, 0xf8, 0xd8, 0x5e, 0x78, 0xee, 0x39, 0x44, 0xdf,
	0xfa, 0x3f, 0x9e, 0xb8, 0x22, 0x7b, 0xb6, 0x16, 0x9f, 0xca, 0xc3, 0x3f, 0xb5, 0x15, 0xd7, 0x02,
	0x95, 0xc2, 0x80, 0xf2, 0xb1, 0xc4, 0xba, 0x57, 0x3b, 0xf5, 0xbd, 0x6d, 0x53, 0x29, 0x6c, 0xbb,
	0x24, 0xb6, 0x9d, 0x58, 0xdf, 0x5f, 0x96, 0x6b, 0xf4, 0xf1, 0x21, 0x00, 0x00, 0xff, 0xff, 0x62,
	0xbb, 0x0f, 0x1a, 0x6f, 0x02, 0x00, 0x00,
}

func (x NamespaceState) String() string {
	s, ok := NamespaceState_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x ArchivalState) String() string {
	s, ok := ArchivalState_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x ReplicationState) String() string {
	s, ok := ReplicationState_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
