// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: enums.proto

package types

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GitProviders int32

const (
	GitProviders_UNKNOWN   GitProviders = 0
	GitProviders_GITHUB    GitProviders = 1
	GitProviders_GITLAB    GitProviders = 2
	GitProviders_BITBUCKET GitProviders = 3
)

// Enum value maps for GitProviders.
var (
	GitProviders_name = map[int32]string{
		0: "UNKNOWN",
		1: "GITHUB",
		2: "GITLAB",
		3: "BITBUCKET",
	}
	GitProviders_value = map[string]int32{
		"UNKNOWN":   0,
		"GITHUB":    1,
		"GITLAB":    2,
		"BITBUCKET": 3,
	}
)

func (x GitProviders) Enum() *GitProviders {
	p := new(GitProviders)
	*p = x
	return p
}

func (x GitProviders) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GitProviders) Descriptor() protoreflect.EnumDescriptor {
	return file_enums_proto_enumTypes[0].Descriptor()
}

func (GitProviders) Type() protoreflect.EnumType {
	return &file_enums_proto_enumTypes[0]
}

func (x GitProviders) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GitProviders.Descriptor instead.
func (GitProviders) EnumDescriptor() ([]byte, []int) {
	return file_enums_proto_rawDescGZIP(), []int{0}
}

var File_enums_proto protoreflect.FileDescriptor

var file_enums_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61,
	0x6c, 0x66, 0x72, 0x65, 0x64, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2a, 0x42, 0x0a, 0x0c, 0x47,
	0x69, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x47, 0x49, 0x54, 0x48,
	0x55, 0x42, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x47, 0x49, 0x54, 0x4c, 0x41, 0x42, 0x10, 0x02,
	0x12, 0x0d, 0x0a, 0x09, 0x42, 0x49, 0x54, 0x42, 0x55, 0x43, 0x4b, 0x45, 0x54, 0x10, 0x03, 0x42,
	0x15, 0x5a, 0x13, 0x61, 0x6c, 0x66, 0x72, 0x65, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_enums_proto_rawDescOnce sync.Once
	file_enums_proto_rawDescData = file_enums_proto_rawDesc
)

func file_enums_proto_rawDescGZIP() []byte {
	file_enums_proto_rawDescOnce.Do(func() {
		file_enums_proto_rawDescData = protoimpl.X.CompressGZIP(file_enums_proto_rawDescData)
	})
	return file_enums_proto_rawDescData
}

var file_enums_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_enums_proto_goTypes = []interface{}{
	(GitProviders)(0), // 0: alfred.types.GitProviders
}
var file_enums_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_enums_proto_init() }
func file_enums_proto_init() {
	if File_enums_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_enums_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_enums_proto_goTypes,
		DependencyIndexes: file_enums_proto_depIdxs,
		EnumInfos:         file_enums_proto_enumTypes,
	}.Build()
	File_enums_proto = out.File
	file_enums_proto_rawDesc = nil
	file_enums_proto_goTypes = nil
	file_enums_proto_depIdxs = nil
}
