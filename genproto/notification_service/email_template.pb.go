// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.17.3
// source: email_template.proto

package notification_service

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EmailTmp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Type        string `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Text        string `protobuf:"bytes,5,opt,name=text,proto3" json:"text,omitempty"`
	CreatedAt   string `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *EmailTmp) Reset() {
	*x = EmailTmp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_email_template_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmailTmp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailTmp) ProtoMessage() {}

func (x *EmailTmp) ProtoReflect() protoreflect.Message {
	mi := &file_email_template_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailTmp.ProtoReflect.Descriptor instead.
func (*EmailTmp) Descriptor() ([]byte, []int) {
	return file_email_template_proto_rawDescGZIP(), []int{0}
}

func (x *EmailTmp) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *EmailTmp) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *EmailTmp) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *EmailTmp) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *EmailTmp) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *EmailTmp) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type CreateEmailTmpReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Type        string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Text        string `protobuf:"bytes,4,opt,name=text,proto3" json:"text,omitempty"`
	CreatedAt   string `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *CreateEmailTmpReq) Reset() {
	*x = CreateEmailTmpReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_email_template_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEmailTmpReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEmailTmpReq) ProtoMessage() {}

func (x *CreateEmailTmpReq) ProtoReflect() protoreflect.Message {
	mi := &file_email_template_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEmailTmpReq.ProtoReflect.Descriptor instead.
func (*CreateEmailTmpReq) Descriptor() ([]byte, []int) {
	return file_email_template_proto_rawDescGZIP(), []int{1}
}

func (x *CreateEmailTmpReq) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateEmailTmpReq) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateEmailTmpReq) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *CreateEmailTmpReq) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *CreateEmailTmpReq) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type CreateEmailTmpRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Type        string `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Text        string `protobuf:"bytes,5,opt,name=text,proto3" json:"text,omitempty"`
	CreatedAt   string `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *CreateEmailTmpRes) Reset() {
	*x = CreateEmailTmpRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_email_template_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEmailTmpRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEmailTmpRes) ProtoMessage() {}

func (x *CreateEmailTmpRes) ProtoReflect() protoreflect.Message {
	mi := &file_email_template_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEmailTmpRes.ProtoReflect.Descriptor instead.
func (*CreateEmailTmpRes) Descriptor() ([]byte, []int) {
	return file_email_template_proto_rawDescGZIP(), []int{2}
}

func (x *CreateEmailTmpRes) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateEmailTmpRes) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateEmailTmpRes) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateEmailTmpRes) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *CreateEmailTmpRes) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *CreateEmailTmpRes) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type GetEmailTmpReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetEmailTmpReq) Reset() {
	*x = GetEmailTmpReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_email_template_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEmailTmpReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEmailTmpReq) ProtoMessage() {}

func (x *GetEmailTmpReq) ProtoReflect() protoreflect.Message {
	mi := &file_email_template_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEmailTmpReq.ProtoReflect.Descriptor instead.
func (*GetEmailTmpReq) Descriptor() ([]byte, []int) {
	return file_email_template_proto_rawDescGZIP(), []int{3}
}

func (x *GetEmailTmpReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetEmailTmpRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Type        string `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Text        string `protobuf:"bytes,5,opt,name=text,proto3" json:"text,omitempty"`
	CreatedAt   string `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *GetEmailTmpRes) Reset() {
	*x = GetEmailTmpRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_email_template_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEmailTmpRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEmailTmpRes) ProtoMessage() {}

func (x *GetEmailTmpRes) ProtoReflect() protoreflect.Message {
	mi := &file_email_template_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEmailTmpRes.ProtoReflect.Descriptor instead.
func (*GetEmailTmpRes) Descriptor() ([]byte, []int) {
	return file_email_template_proto_rawDescGZIP(), []int{4}
}

func (x *GetEmailTmpRes) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetEmailTmpRes) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *GetEmailTmpRes) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *GetEmailTmpRes) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *GetEmailTmpRes) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *GetEmailTmpRes) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type GetEmailTmpListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  int32  `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int32  `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Search string `protobuf:"bytes,3,opt,name=search,proto3" json:"search,omitempty"`
	Type   string `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *GetEmailTmpListReq) Reset() {
	*x = GetEmailTmpListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_email_template_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEmailTmpListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEmailTmpListReq) ProtoMessage() {}

func (x *GetEmailTmpListReq) ProtoReflect() protoreflect.Message {
	mi := &file_email_template_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEmailTmpListReq.ProtoReflect.Descriptor instead.
func (*GetEmailTmpListReq) Descriptor() ([]byte, []int) {
	return file_email_template_proto_rawDescGZIP(), []int{5}
}

func (x *GetEmailTmpListReq) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetEmailTmpListReq) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetEmailTmpListReq) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

func (x *GetEmailTmpListReq) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type GetEmailTmpListRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EmailTmps []*EmailTmp `protobuf:"bytes,1,rep,name=email_tmps,json=emailTmps,proto3" json:"email_tmps,omitempty"`
	Count     int32       `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *GetEmailTmpListRes) Reset() {
	*x = GetEmailTmpListRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_email_template_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEmailTmpListRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEmailTmpListRes) ProtoMessage() {}

func (x *GetEmailTmpListRes) ProtoReflect() protoreflect.Message {
	mi := &file_email_template_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEmailTmpListRes.ProtoReflect.Descriptor instead.
func (*GetEmailTmpListRes) Descriptor() ([]byte, []int) {
	return file_email_template_proto_rawDescGZIP(), []int{6}
}

func (x *GetEmailTmpListRes) GetEmailTmps() []*EmailTmp {
	if x != nil {
		return x.EmailTmps
	}
	return nil
}

func (x *GetEmailTmpListRes) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type UpdateEmailTmpReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Type        string `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Text        string `protobuf:"bytes,5,opt,name=text,proto3" json:"text,omitempty"`
	CreatedAt   string `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *UpdateEmailTmpReq) Reset() {
	*x = UpdateEmailTmpReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_email_template_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEmailTmpReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEmailTmpReq) ProtoMessage() {}

func (x *UpdateEmailTmpReq) ProtoReflect() protoreflect.Message {
	mi := &file_email_template_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEmailTmpReq.ProtoReflect.Descriptor instead.
func (*UpdateEmailTmpReq) Descriptor() ([]byte, []int) {
	return file_email_template_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateEmailTmpReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateEmailTmpReq) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateEmailTmpReq) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateEmailTmpReq) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *UpdateEmailTmpReq) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *UpdateEmailTmpReq) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type UpdateEmailTmpRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Type        string `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Text        string `protobuf:"bytes,5,opt,name=text,proto3" json:"text,omitempty"`
	CreatedAt   string `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *UpdateEmailTmpRes) Reset() {
	*x = UpdateEmailTmpRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_email_template_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEmailTmpRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEmailTmpRes) ProtoMessage() {}

func (x *UpdateEmailTmpRes) ProtoReflect() protoreflect.Message {
	mi := &file_email_template_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEmailTmpRes.ProtoReflect.Descriptor instead.
func (*UpdateEmailTmpRes) Descriptor() ([]byte, []int) {
	return file_email_template_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateEmailTmpRes) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateEmailTmpRes) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateEmailTmpRes) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateEmailTmpRes) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *UpdateEmailTmpRes) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *UpdateEmailTmpRes) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type DeleteEmailTmpReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteEmailTmpReq) Reset() {
	*x = DeleteEmailTmpReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_email_template_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteEmailTmpReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEmailTmpReq) ProtoMessage() {}

func (x *DeleteEmailTmpReq) ProtoReflect() protoreflect.Message {
	mi := &file_email_template_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEmailTmpReq.ProtoReflect.Descriptor instead.
func (*DeleteEmailTmpReq) Descriptor() ([]byte, []int) {
	return file_email_template_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteEmailTmpReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_email_template_proto protoreflect.FileDescriptor

var file_email_template_proto_rawDesc = []byte{
	0x0a, 0x14, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x99, 0x01, 0x0a, 0x08, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x54, 0x6d, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x92, 0x01, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x54, 0x6d, 0x70, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xa2, 0x01, 0x0a, 0x11, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x54, 0x6d, 0x70, 0x52, 0x65, 0x73,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22,
	0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x54, 0x6d, 0x70, 0x52, 0x65,
	0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x9f, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x54, 0x6d,
	0x70, 0x52, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x22, 0x6e, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x54,
	0x6d, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x22, 0x69, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x54,
	0x6d, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x12, 0x3d, 0x0a, 0x0a, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x5f, 0x74, 0x6d, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x54, 0x6d, 0x70, 0x52, 0x09, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x54, 0x6d, 0x70, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xa2,
	0x01, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x54, 0x6d,
	0x70, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x22, 0xa2, 0x01, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x54, 0x6d, 0x70, 0x52, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x23, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x54, 0x6d, 0x70, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0xf8, 0x03,
	0x0a, 0x0f, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x54, 0x6d, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x64, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x54, 0x6d, 0x70, 0x12, 0x27, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x54, 0x6d, 0x70, 0x52, 0x65, 0x71, 0x1a, 0x27, 0x2e, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x54,
	0x6d, 0x70, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x5b, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x54, 0x6d, 0x70, 0x12, 0x24, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65,
	0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x54, 0x6d, 0x70, 0x52, 0x65, 0x71, 0x1a, 0x24, 0x2e, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x54, 0x6d, 0x70, 0x52,
	0x65, 0x73, 0x22, 0x00, 0x12, 0x67, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x54, 0x6d, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x28, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47,
	0x65, 0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x54, 0x6d, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x1a, 0x28, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x54, 0x6d, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x64, 0x0a,
	0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x54, 0x6d, 0x70, 0x12,
	0x27, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x54, 0x6d, 0x70, 0x52, 0x65, 0x71, 0x1a, 0x27, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x54, 0x6d, 0x70, 0x52, 0x65,
	0x73, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x54, 0x6d, 0x70, 0x12, 0x27, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x54, 0x6d, 0x70, 0x52, 0x65, 0x71, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x1f, 0x5a, 0x1d, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_email_template_proto_rawDescOnce sync.Once
	file_email_template_proto_rawDescData = file_email_template_proto_rawDesc
)

func file_email_template_proto_rawDescGZIP() []byte {
	file_email_template_proto_rawDescOnce.Do(func() {
		file_email_template_proto_rawDescData = protoimpl.X.CompressGZIP(file_email_template_proto_rawDescData)
	})
	return file_email_template_proto_rawDescData
}

var file_email_template_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_email_template_proto_goTypes = []interface{}{
	(*EmailTmp)(nil),           // 0: notification_service.EmailTmp
	(*CreateEmailTmpReq)(nil),  // 1: notification_service.CreateEmailTmpReq
	(*CreateEmailTmpRes)(nil),  // 2: notification_service.CreateEmailTmpRes
	(*GetEmailTmpReq)(nil),     // 3: notification_service.GetEmailTmpReq
	(*GetEmailTmpRes)(nil),     // 4: notification_service.GetEmailTmpRes
	(*GetEmailTmpListReq)(nil), // 5: notification_service.GetEmailTmpListReq
	(*GetEmailTmpListRes)(nil), // 6: notification_service.GetEmailTmpListRes
	(*UpdateEmailTmpReq)(nil),  // 7: notification_service.UpdateEmailTmpReq
	(*UpdateEmailTmpRes)(nil),  // 8: notification_service.UpdateEmailTmpRes
	(*DeleteEmailTmpReq)(nil),  // 9: notification_service.DeleteEmailTmpReq
	(*emptypb.Empty)(nil),      // 10: google.protobuf.Empty
}
var file_email_template_proto_depIdxs = []int32{
	0,  // 0: notification_service.GetEmailTmpListRes.email_tmps:type_name -> notification_service.EmailTmp
	1,  // 1: notification_service.EmailTmpService.CreateEmailTmp:input_type -> notification_service.CreateEmailTmpReq
	3,  // 2: notification_service.EmailTmpService.GetEmailTmp:input_type -> notification_service.GetEmailTmpReq
	5,  // 3: notification_service.EmailTmpService.GetEmailTmpList:input_type -> notification_service.GetEmailTmpListReq
	7,  // 4: notification_service.EmailTmpService.UpdateEmailTmp:input_type -> notification_service.UpdateEmailTmpReq
	9,  // 5: notification_service.EmailTmpService.DeleteEmailTmp:input_type -> notification_service.DeleteEmailTmpReq
	2,  // 6: notification_service.EmailTmpService.CreateEmailTmp:output_type -> notification_service.CreateEmailTmpRes
	4,  // 7: notification_service.EmailTmpService.GetEmailTmp:output_type -> notification_service.GetEmailTmpRes
	6,  // 8: notification_service.EmailTmpService.GetEmailTmpList:output_type -> notification_service.GetEmailTmpListRes
	8,  // 9: notification_service.EmailTmpService.UpdateEmailTmp:output_type -> notification_service.UpdateEmailTmpRes
	10, // 10: notification_service.EmailTmpService.DeleteEmailTmp:output_type -> google.protobuf.Empty
	6,  // [6:11] is the sub-list for method output_type
	1,  // [1:6] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_email_template_proto_init() }
func file_email_template_proto_init() {
	if File_email_template_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_email_template_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmailTmp); i {
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
		file_email_template_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEmailTmpReq); i {
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
		file_email_template_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEmailTmpRes); i {
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
		file_email_template_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEmailTmpReq); i {
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
		file_email_template_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEmailTmpRes); i {
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
		file_email_template_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEmailTmpListReq); i {
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
		file_email_template_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEmailTmpListRes); i {
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
		file_email_template_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateEmailTmpReq); i {
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
		file_email_template_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateEmailTmpRes); i {
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
		file_email_template_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteEmailTmpReq); i {
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
			RawDescriptor: file_email_template_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_email_template_proto_goTypes,
		DependencyIndexes: file_email_template_proto_depIdxs,
		MessageInfos:      file_email_template_proto_msgTypes,
	}.Build()
	File_email_template_proto = out.File
	file_email_template_proto_rawDesc = nil
	file_email_template_proto_goTypes = nil
	file_email_template_proto_depIdxs = nil
}
