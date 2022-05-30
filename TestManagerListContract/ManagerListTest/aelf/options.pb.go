// Code generated by protoc-gen-go. DO NOT EDIT.
// source: options.proto

package aelf

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

var E_Identity = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         500001,
	Name:          "aelf.identity",
	Tag:           "bytes,500001,opt,name=identity",
	Filename:      "options.proto",
}

var E_Base = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.ServiceOptions)(nil),
	ExtensionType: ([]string)(nil),
	Field:         505001,
	Name:          "aelf.base",
	Tag:           "bytes,505001,rep,name=base",
	Filename:      "options.proto",
}

var E_CsharpState = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.ServiceOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         505030,
	Name:          "aelf.csharp_state",
	Tag:           "bytes,505030,opt,name=csharp_state",
	Filename:      "options.proto",
}

var E_IsView = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         506001,
	Name:          "aelf.is_view",
	Tag:           "varint,506001,opt,name=is_view",
	Filename:      "options.proto",
}

var E_IsEvent = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         50100,
	Name:          "aelf.is_event",
	Tag:           "varint,50100,opt,name=is_event",
	Filename:      "options.proto",
}

var E_IsIndexed = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         502001,
	Name:          "aelf.is_indexed",
	Tag:           "varint,502001,opt,name=is_indexed",
	Filename:      "options.proto",
}

func init() {
	// proto.RegisterExtension(E_Identity)
	// proto.RegisterExtension(E_Base)
	// proto.RegisterExtension(E_CsharpState)
	// proto.RegisterExtension(E_IsView)
	// proto.RegisterExtension(E_IsEvent)
	// proto.RegisterExtension(E_IsIndexed)
}

func init() {  }

var fileDescriptor_110d40819f1994f9 = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0xd0, 0x3f, 0x4b, 0xc3, 0x40,
	0x18, 0xc7, 0x71, 0xb4, 0xa1, 0x49, 0x4e, 0x5d, 0x32, 0x89, 0x68, 0xcc, 0xd8, 0x29, 0x1d, 0xc4,
	0xc1, 0x43, 0x04, 0xc5, 0x0a, 0x0e, 0x45, 0x48, 0xc1, 0x35, 0xe4, 0xcf, 0x93, 0xf4, 0x81, 0x23,
	0x17, 0xee, 0x39, 0x53, 0x7d, 0x05, 0xce, 0xbe, 0x04, 0x47, 0x9d, 0x7d, 0x05, 0x0e, 0xee, 0xbe,
	0x8c, 0x8e, 0x7d, 0x07, 0x92, 0x1c, 0xcd, 0xd2, 0x0e, 0xae, 0xc7, 0xef, 0xf3, 0x7c, 0xe1, 0xd8,
	0x81, 0xac, 0x35, 0xca, 0x8a, 0xc2, 0x5a, 0x49, 0x2d, 0x3d, 0x2b, 0x01, 0x51, 0x1c, 0x05, 0xa5,
	0x94, 0xa5, 0x80, 0x71, 0xf7, 0x96, 0x3e, 0x15, 0xe3, 0x1c, 0x28, 0x53, 0x58, 0x6b, 0xa9, 0xcc,
	0x8e, 0x73, 0xe6, 0x60, 0x0e, 0x95, 0x46, 0xfd, 0xe2, 0x1d, 0x87, 0x66, 0x1e, 0xae, 0xe7, 0xe1,
	0x1d, 0x0a, 0x78, 0x30, 0x77, 0x0f, 0xdf, 0xbf, 0xfd, 0x60, 0x67, 0xe4, 0x46, 0xfd, 0x9e, 0x9f,
	0x33, 0x2b, 0x4d, 0x08, 0xbc, 0xd3, 0x0d, 0x37, 0x03, 0xd5, 0x60, 0xd6, 0xd3, 0x8f, 0xa5, 0x1f,
	0x0c, 0x46, 0x6e, 0xd4, 0xcd, 0xf9, 0x2d, 0xdb, 0xcf, 0x68, 0x9e, 0xa8, 0x3a, 0x26, 0x9d, 0xe8,
	0x7f, 0xf0, 0x9f, 0xa5, 0x29, 0xef, 0x19, 0x36, 0x6b, 0x15, 0xbf, 0x60, 0x36, 0x52, 0xdc, 0x20,
	0x2c, 0x3c, 0x7f, 0xe3, 0xc0, 0x14, 0xf4, 0x5c, 0xe6, 0x6b, 0xff, 0xb6, 0x6a, 0xbd, 0x13, 0x0d,
	0x91, 0x1e, 0x11, 0x16, 0xfc, 0x92, 0x39, 0x48, 0x31, 0x34, 0x50, 0xe9, 0x2d, 0xf1, 0x29, 0x10,
	0x25, 0x65, 0x1f, 0xff, 0x7a, 0x1d, 0x74, 0xd8, 0x46, 0x9a, 0xb4, 0x82, 0x5f, 0x31, 0x86, 0x14,
	0x63, 0x95, 0xc3, 0x33, 0xe4, 0xde, 0xc9, 0x96, 0x3f, 0x03, 0xd1, 0xa7, 0x57, 0xbf, 0x26, 0xed,
	0x22, 0xdd, 0x1b, 0x71, 0x63, 0x7f, 0xee, 0x5a, 0xd7, 0x13, 0x51, 0xa4, 0xc3, 0x8e, 0x9c, 0xfd,
	0x05, 0x00, 0x00, 0xff, 0xff, 0xd7, 0xf7, 0x4b, 0x7f, 0xba, 0x01, 0x00, 0x00,
}