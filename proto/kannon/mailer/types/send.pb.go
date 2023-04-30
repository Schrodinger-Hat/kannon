// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: kannon/mailer/types/send.proto

package types

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Sender struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Alias string `protobuf:"bytes,2,opt,name=alias,proto3" json:"alias,omitempty"`
}

func (x *Sender) Reset() {
	*x = Sender{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kannon_mailer_types_send_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sender) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sender) ProtoMessage() {}

func (x *Sender) ProtoReflect() protoreflect.Message {
	mi := &file_kannon_mailer_types_send_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sender.ProtoReflect.Descriptor instead.
func (*Sender) Descriptor() ([]byte, []int) {
	return file_kannon_mailer_types_send_proto_rawDescGZIP(), []int{0}
}

func (x *Sender) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Sender) GetAlias() string {
	if x != nil {
		return x.Alias
	}
	return ""
}

type Recipient struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email  string            `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Fields map[string]string `protobuf:"bytes,2,rep,name=fields,proto3" json:"fields,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Recipient) Reset() {
	*x = Recipient{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kannon_mailer_types_send_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Recipient) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Recipient) ProtoMessage() {}

func (x *Recipient) ProtoReflect() protoreflect.Message {
	mi := &file_kannon_mailer_types_send_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Recipient.ProtoReflect.Descriptor instead.
func (*Recipient) Descriptor() ([]byte, []int) {
	return file_kannon_mailer_types_send_proto_rawDescGZIP(), []int{1}
}

func (x *Recipient) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Recipient) GetFields() map[string]string {
	if x != nil {
		return x.Fields
	}
	return nil
}

var File_kannon_mailer_types_send_proto protoreflect.FileDescriptor

var file_kannon_mailer_types_send_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x6b, 0x61, 0x6e, 0x6e, 0x6f, 0x6e, 0x2f, 0x6d, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x73, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x17, 0x70, 0x6b, 0x67, 0x2e, 0x6b, 0x61, 0x6e, 0x6e, 0x6f, 0x6e, 0x2e, 0x6d, 0x61, 0x69,
	0x6c, 0x65, 0x72, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x22, 0x34, 0x0a, 0x06, 0x53, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c, 0x69,
	0x61, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x22,
	0xa4, 0x01, 0x0a, 0x09, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x46, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x6b, 0x61, 0x6e, 0x6e, 0x6f, 0x6e,
	0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x52, 0x65,
	0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x75, 0x64, 0x75, 0x73, 0x72, 0x75, 0x73, 0x73, 0x6f, 0x2f,
	0x6b, 0x61, 0x6e, 0x6e, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6b, 0x61, 0x6e,
	0x6e, 0x6f, 0x6e, 0x2f, 0x6d, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kannon_mailer_types_send_proto_rawDescOnce sync.Once
	file_kannon_mailer_types_send_proto_rawDescData = file_kannon_mailer_types_send_proto_rawDesc
)

func file_kannon_mailer_types_send_proto_rawDescGZIP() []byte {
	file_kannon_mailer_types_send_proto_rawDescOnce.Do(func() {
		file_kannon_mailer_types_send_proto_rawDescData = protoimpl.X.CompressGZIP(file_kannon_mailer_types_send_proto_rawDescData)
	})
	return file_kannon_mailer_types_send_proto_rawDescData
}

var file_kannon_mailer_types_send_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_kannon_mailer_types_send_proto_goTypes = []interface{}{
	(*Sender)(nil),    // 0: pkg.kannon.mailer.types.Sender
	(*Recipient)(nil), // 1: pkg.kannon.mailer.types.Recipient
	nil,               // 2: pkg.kannon.mailer.types.Recipient.FieldsEntry
}
var file_kannon_mailer_types_send_proto_depIdxs = []int32{
	2, // 0: pkg.kannon.mailer.types.Recipient.fields:type_name -> pkg.kannon.mailer.types.Recipient.FieldsEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_kannon_mailer_types_send_proto_init() }
func file_kannon_mailer_types_send_proto_init() {
	if File_kannon_mailer_types_send_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kannon_mailer_types_send_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sender); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_kannon_mailer_types_send_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Recipient); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_kannon_mailer_types_send_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kannon_mailer_types_send_proto_goTypes,
		DependencyIndexes: file_kannon_mailer_types_send_proto_depIdxs,
		MessageInfos:      file_kannon_mailer_types_send_proto_msgTypes,
	}.Build()
	File_kannon_mailer_types_send_proto = out.File
	file_kannon_mailer_types_send_proto_rawDesc = nil
	file_kannon_mailer_types_send_proto_goTypes = nil
	file_kannon_mailer_types_send_proto_depIdxs = nil
}
