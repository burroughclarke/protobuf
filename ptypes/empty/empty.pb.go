// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/golang/protobuf/ptypes/empty/empty.proto

package empty

import (
	proto "github.com/golang/protobuf/proto"
	protoapi "github.com/golang/protobuf/protoapi"
	protoreflect "github.com/golang/protobuf/v2/reflect/protoreflect"
	protoimpl "github.com/golang/protobuf/v2/runtime/protoimpl"
	known "github.com/golang/protobuf/v2/types/known"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Symbols defined in public import of google/protobuf/empty.proto

type Empty = known.Empty

func init() {
	proto.RegisterFile("github.com/golang/protobuf/ptypes/empty/empty.proto", xxx_File_github_com_golang_protobuf_ptypes_empty_empty_proto_rawdesc_gzipped)
}

var xxx_File_github_com_golang_protobuf_ptypes_empty_empty_proto_rawdesc = []byte{
	// 141 bytes of the wire-encoded FileDescriptorProto
	0x0a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x70, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x70, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x3b, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x50, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var xxx_File_github_com_golang_protobuf_ptypes_empty_empty_proto_rawdesc_gzipped = protoapi.CompressGZIP(xxx_File_github_com_golang_protobuf_ptypes_empty_empty_proto_rawdesc)

const _ = protoimpl.EnforceVersion(protoimpl.Version - 0)

var File_github_com_golang_protobuf_ptypes_empty_empty_proto protoreflect.FileDescriptor

var xxx_File_github_com_golang_protobuf_ptypes_empty_empty_proto_goTypes = []interface{}{}
var xxx_File_github_com_golang_protobuf_ptypes_empty_empty_proto_depIdxs = []int32{}

func init() { xxx_File_github_com_golang_protobuf_ptypes_empty_empty_proto_init() }
func xxx_File_github_com_golang_protobuf_ptypes_empty_empty_proto_init() {
	if File_github_com_golang_protobuf_ptypes_empty_empty_proto != nil {
		return
	}
	File_github_com_golang_protobuf_ptypes_empty_empty_proto = protoimpl.FileBuilder{
		RawDescriptor:     xxx_File_github_com_golang_protobuf_ptypes_empty_empty_proto_rawdesc,
		GoTypes:           xxx_File_github_com_golang_protobuf_ptypes_empty_empty_proto_goTypes,
		DependencyIndexes: xxx_File_github_com_golang_protobuf_ptypes_empty_empty_proto_depIdxs,
	}.Init()
	xxx_File_github_com_golang_protobuf_ptypes_empty_empty_proto_goTypes = nil
	xxx_File_github_com_golang_protobuf_ptypes_empty_empty_proto_depIdxs = nil
}
