// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: template.proto

package pbtemplate

import (
	base "github.com/TencentBlueKing/bk-bscp/pkg/protocol/core/base"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
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

// Template source resource reference: pkg/dal/table/template.go
type Template struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint32              `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Spec       *TemplateSpec       `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	Attachment *TemplateAttachment `protobuf:"bytes,3,opt,name=attachment,proto3" json:"attachment,omitempty"`
	Revision   *base.Revision      `protobuf:"bytes,4,opt,name=revision,proto3" json:"revision,omitempty"`
}

func (x *Template) Reset() {
	*x = Template{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Template) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Template) ProtoMessage() {}

func (x *Template) ProtoReflect() protoreflect.Message {
	mi := &file_template_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Template.ProtoReflect.Descriptor instead.
func (*Template) Descriptor() ([]byte, []int) {
	return file_template_proto_rawDescGZIP(), []int{0}
}

func (x *Template) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Template) GetSpec() *TemplateSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *Template) GetAttachment() *TemplateAttachment {
	if x != nil {
		return x.Attachment
	}
	return nil
}

func (x *Template) GetRevision() *base.Revision {
	if x != nil {
		return x.Revision
	}
	return nil
}

// TemplateSpec source resource reference: pkg/dal/table/template.go
type TemplateSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Path string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Memo string `protobuf:"bytes,3,opt,name=memo,proto3" json:"memo,omitempty"`
}

func (x *TemplateSpec) Reset() {
	*x = TemplateSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemplateSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemplateSpec) ProtoMessage() {}

func (x *TemplateSpec) ProtoReflect() protoreflect.Message {
	mi := &file_template_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemplateSpec.ProtoReflect.Descriptor instead.
func (*TemplateSpec) Descriptor() ([]byte, []int) {
	return file_template_proto_rawDescGZIP(), []int{1}
}

func (x *TemplateSpec) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TemplateSpec) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *TemplateSpec) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

// TemplateAttachment source resource reference: pkg/dal/table/template.go
type TemplateAttachment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BizId           uint32 `protobuf:"varint,1,opt,name=biz_id,json=bizId,proto3" json:"biz_id,omitempty"`
	TemplateSpaceId uint32 `protobuf:"varint,2,opt,name=template_space_id,json=templateSpaceId,proto3" json:"template_space_id,omitempty"`
}

func (x *TemplateAttachment) Reset() {
	*x = TemplateAttachment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemplateAttachment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemplateAttachment) ProtoMessage() {}

func (x *TemplateAttachment) ProtoReflect() protoreflect.Message {
	mi := &file_template_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemplateAttachment.ProtoReflect.Descriptor instead.
func (*TemplateAttachment) Descriptor() ([]byte, []int) {
	return file_template_proto_rawDescGZIP(), []int{2}
}

func (x *TemplateAttachment) GetBizId() uint32 {
	if x != nil {
		return x.BizId
	}
	return 0
}

func (x *TemplateAttachment) GetTemplateSpaceId() uint32 {
	if x != nil {
		return x.TemplateSpaceId
	}
	return 0
}

var File_template_proto protoreflect.FileDescriptor

var file_template_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0a, 0x70, 0x62, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x1a, 0x21, 0x70, 0x6b,
	0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x62, 0x61, 0x73, 0x65, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e,
	0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xcb, 0x01, 0x0a, 0x08, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x23, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x13, 0x92, 0x41, 0x10, 0x32, 0x0e, 0xe6,
	0xa8, 0xa1, 0xe6, 0x9d, 0xbf, 0xe6, 0x96, 0x87, 0xe4, 0xbb, 0xb6, 0x49, 0x44, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x2c, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x70, 0x62, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x54, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12,
	0x3e, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x62, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x2c, 0x0a, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x70, 0x62, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x80, 0x01,
	0x0a, 0x0c, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x53, 0x70, 0x65, 0x63, 0x12, 0x22,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0e, 0x92, 0x41,
	0x0b, 0x32, 0x09, 0xe6, 0x96, 0x87, 0xe4, 0xbb, 0xb6, 0xe5, 0x90, 0x8d, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x25, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x11, 0x92, 0x41, 0x0e, 0x32, 0x0c, 0xe6, 0x96, 0x87, 0xe4, 0xbb, 0xb6, 0xe8, 0xb7, 0xaf,
	0xe5, 0xbe, 0x84, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x25, 0x0a, 0x04, 0x6d, 0x65, 0x6d,
	0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x11, 0x92, 0x41, 0x0e, 0x32, 0x0c, 0xe6, 0x96,
	0x87, 0xe4, 0xbb, 0xb6, 0xe6, 0x8f, 0x8f, 0xe8, 0xbf, 0xb0, 0x52, 0x04, 0x6d, 0x65, 0x6d, 0x6f,
	0x22, 0x7b, 0x0a, 0x12, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x41, 0x74, 0x74, 0x61,
	0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x06, 0x62, 0x69, 0x7a, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x0d, 0x92, 0x41, 0x0a, 0x32, 0x08, 0xe4, 0xb8, 0x9a,
	0xe5, 0x8a, 0xa1, 0x49, 0x44, 0x52, 0x05, 0x62, 0x69, 0x7a, 0x49, 0x64, 0x12, 0x3f, 0x0a, 0x11,
	0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x13, 0x92, 0x41, 0x10, 0x32, 0x0e, 0xe6, 0xa8,
	0xa1, 0xe6, 0x9d, 0xbf, 0xe7, 0xa9, 0xba, 0xe9, 0x97, 0xb4, 0x49, 0x44, 0x52, 0x0f, 0x74, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x53, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x42, 0x4a, 0x5a,
	0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x54, 0x65, 0x6e, 0x63,
	0x65, 0x6e, 0x74, 0x42, 0x6c, 0x75, 0x65, 0x4b, 0x69, 0x6e, 0x67, 0x2f, 0x62, 0x6b, 0x2d, 0x62,
	0x73, 0x63, 0x70, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x3b, 0x70,
	0x62, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_template_proto_rawDescOnce sync.Once
	file_template_proto_rawDescData = file_template_proto_rawDesc
)

func file_template_proto_rawDescGZIP() []byte {
	file_template_proto_rawDescOnce.Do(func() {
		file_template_proto_rawDescData = protoimpl.X.CompressGZIP(file_template_proto_rawDescData)
	})
	return file_template_proto_rawDescData
}

var file_template_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_template_proto_goTypes = []interface{}{
	(*Template)(nil),           // 0: pbtemplate.Template
	(*TemplateSpec)(nil),       // 1: pbtemplate.TemplateSpec
	(*TemplateAttachment)(nil), // 2: pbtemplate.TemplateAttachment
	(*base.Revision)(nil),      // 3: pbbase.Revision
}
var file_template_proto_depIdxs = []int32{
	1, // 0: pbtemplate.Template.spec:type_name -> pbtemplate.TemplateSpec
	2, // 1: pbtemplate.Template.attachment:type_name -> pbtemplate.TemplateAttachment
	3, // 2: pbtemplate.Template.revision:type_name -> pbbase.Revision
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_template_proto_init() }
func file_template_proto_init() {
	if File_template_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_template_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Template); i {
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
		file_template_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemplateSpec); i {
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
		file_template_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemplateAttachment); i {
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
			RawDescriptor: file_template_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_template_proto_goTypes,
		DependencyIndexes: file_template_proto_depIdxs,
		MessageInfos:      file_template_proto_msgTypes,
	}.Build()
	File_template_proto = out.File
	file_template_proto_rawDesc = nil
	file_template_proto_goTypes = nil
	file_template_proto_depIdxs = nil
}
