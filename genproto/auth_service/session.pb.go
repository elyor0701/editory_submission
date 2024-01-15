// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.17.3
// source: session.proto

package auth_service

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

type LoginReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XRole    string `protobuf:"bytes,3,opt,name=x_role,json=xRole,proto3" json:"x_role,omitempty"`
}

func (x *LoginReq) Reset() {
	*x = LoginReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_session_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReq) ProtoMessage() {}

func (x *LoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_session_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReq.ProtoReflect.Descriptor instead.
func (*LoginReq) Descriptor() ([]byte, []int) {
	return file_session_proto_rawDescGZIP(), []int{0}
}

func (x *LoginReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *LoginReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *LoginReq) GetXRole() string {
	if x != nil {
		return x.XRole
	}
	return ""
}

type LoginRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserFound bool       `protobuf:"varint,1,opt,name=user_found,json=userFound,proto3" json:"user_found,omitempty"`
	User      *User      `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	RoleId    string     `protobuf:"bytes,3,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	Token     *Token     `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
	Sessions  []*Session `protobuf:"bytes,5,rep,name=sessions,proto3" json:"sessions,omitempty"`
	Roles     []*Role    `protobuf:"bytes,6,rep,name=roles,proto3" json:"roles,omitempty"`
}

func (x *LoginRes) Reset() {
	*x = LoginRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_session_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRes) ProtoMessage() {}

func (x *LoginRes) ProtoReflect() protoreflect.Message {
	mi := &file_session_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRes.ProtoReflect.Descriptor instead.
func (*LoginRes) Descriptor() ([]byte, []int) {
	return file_session_proto_rawDescGZIP(), []int{1}
}

func (x *LoginRes) GetUserFound() bool {
	if x != nil {
		return x.UserFound
	}
	return false
}

func (x *LoginRes) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *LoginRes) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

func (x *LoginRes) GetToken() *Token {
	if x != nil {
		return x.Token
	}
	return nil
}

func (x *LoginRes) GetSessions() []*Session {
	if x != nil {
		return x.Sessions
	}
	return nil
}

func (x *LoginRes) GetRoles() []*Role {
	if x != nil {
		return x.Roles
	}
	return nil
}

type LogoutReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
}

func (x *LogoutReq) Reset() {
	*x = LogoutReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_session_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogoutReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutReq) ProtoMessage() {}

func (x *LogoutReq) ProtoReflect() protoreflect.Message {
	mi := &file_session_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutReq.ProtoReflect.Descriptor instead.
func (*LogoutReq) Descriptor() ([]byte, []int) {
	return file_session_proto_rawDescGZIP(), []int{2}
}

func (x *LogoutReq) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

type RefreshTokenReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RefreshToken string `protobuf:"bytes,1,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	RoleId       string `protobuf:"bytes,2,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
}

func (x *RefreshTokenReq) Reset() {
	*x = RefreshTokenReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_session_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshTokenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshTokenReq) ProtoMessage() {}

func (x *RefreshTokenReq) ProtoReflect() protoreflect.Message {
	mi := &file_session_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshTokenReq.ProtoReflect.Descriptor instead.
func (*RefreshTokenReq) Descriptor() ([]byte, []int) {
	return file_session_proto_rawDescGZIP(), []int{3}
}

func (x *RefreshTokenReq) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *RefreshTokenReq) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

type RefreshTokenRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token *Token `protobuf:"bytes,6,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *RefreshTokenRes) Reset() {
	*x = RefreshTokenRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_session_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshTokenRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshTokenRes) ProtoMessage() {}

func (x *RefreshTokenRes) ProtoReflect() protoreflect.Message {
	mi := &file_session_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshTokenRes.ProtoReflect.Descriptor instead.
func (*RefreshTokenRes) Descriptor() ([]byte, []int) {
	return file_session_proto_rawDescGZIP(), []int{4}
}

func (x *RefreshTokenRes) GetToken() *Token {
	if x != nil {
		return x.Token
	}
	return nil
}

type HasAccessReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	Path        string `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`
	Method      string `protobuf:"bytes,5,opt,name=method,proto3" json:"method,omitempty"`
}

func (x *HasAccessReq) Reset() {
	*x = HasAccessReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_session_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HasAccessReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HasAccessReq) ProtoMessage() {}

func (x *HasAccessReq) ProtoReflect() protoreflect.Message {
	mi := &file_session_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HasAccessReq.ProtoReflect.Descriptor instead.
func (*HasAccessReq) Descriptor() ([]byte, []int) {
	return file_session_proto_rawDescGZIP(), []int{5}
}

func (x *HasAccessReq) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *HasAccessReq) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *HasAccessReq) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

type HasAccessRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId    string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	RoleId    string `protobuf:"bytes,3,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	Ip        string `protobuf:"bytes,4,opt,name=ip,proto3" json:"ip,omitempty"`
	Data      string `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
	ExpiresAt string `protobuf:"bytes,6,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	CreatedAt string `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *HasAccessRes) Reset() {
	*x = HasAccessRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_session_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HasAccessRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HasAccessRes) ProtoMessage() {}

func (x *HasAccessRes) ProtoReflect() protoreflect.Message {
	mi := &file_session_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HasAccessRes.ProtoReflect.Descriptor instead.
func (*HasAccessRes) Descriptor() ([]byte, []int) {
	return file_session_proto_rawDescGZIP(), []int{6}
}

func (x *HasAccessRes) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *HasAccessRes) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *HasAccessRes) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

func (x *HasAccessRes) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *HasAccessRes) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *HasAccessRes) GetExpiresAt() string {
	if x != nil {
		return x.ExpiresAt
	}
	return ""
}

func (x *HasAccessRes) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *HasAccessRes) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type SessionAndTokenReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LoginData *LoginRes `protobuf:"bytes,1,opt,name=login_data,json=loginData,proto3" json:"login_data,omitempty"`
}

func (x *SessionAndTokenReq) Reset() {
	*x = SessionAndTokenReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_session_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionAndTokenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionAndTokenReq) ProtoMessage() {}

func (x *SessionAndTokenReq) ProtoReflect() protoreflect.Message {
	mi := &file_session_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionAndTokenReq.ProtoReflect.Descriptor instead.
func (*SessionAndTokenReq) Descriptor() ([]byte, []int) {
	return file_session_proto_rawDescGZIP(), []int{7}
}

func (x *SessionAndTokenReq) GetLoginData() *LoginRes {
	if x != nil {
		return x.LoginData
	}
	return nil
}

type SessionPrimaryKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SessionPrimaryKey) Reset() {
	*x = SessionPrimaryKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_session_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionPrimaryKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionPrimaryKey) ProtoMessage() {}

func (x *SessionPrimaryKey) ProtoReflect() protoreflect.Message {
	mi := &file_session_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionPrimaryKey.ProtoReflect.Descriptor instead.
func (*SessionPrimaryKey) Descriptor() ([]byte, []int) {
	return file_session_proto_rawDescGZIP(), []int{8}
}

func (x *SessionPrimaryKey) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type SessionGetList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  int32  `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int32  `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Search string `protobuf:"bytes,3,opt,name=search,proto3" json:"search,omitempty"`
}

func (x *SessionGetList) Reset() {
	*x = SessionGetList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_session_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionGetList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionGetList) ProtoMessage() {}

func (x *SessionGetList) ProtoReflect() protoreflect.Message {
	mi := &file_session_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionGetList.ProtoReflect.Descriptor instead.
func (*SessionGetList) Descriptor() ([]byte, []int) {
	return file_session_proto_rawDescGZIP(), []int{9}
}

func (x *SessionGetList) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *SessionGetList) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *SessionGetList) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

type GetSessionListRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sessions []*Session `protobuf:"bytes,1,rep,name=sessions,proto3" json:"sessions,omitempty"`
	Count    int32      `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *GetSessionListRes) Reset() {
	*x = GetSessionListRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_session_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSessionListRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSessionListRes) ProtoMessage() {}

func (x *GetSessionListRes) ProtoReflect() protoreflect.Message {
	mi := &file_session_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSessionListRes.ProtoReflect.Descriptor instead.
func (*GetSessionListRes) Descriptor() ([]byte, []int) {
	return file_session_proto_rawDescGZIP(), []int{10}
}

func (x *GetSessionListRes) GetSessions() []*Session {
	if x != nil {
		return x.Sessions
	}
	return nil
}

func (x *GetSessionListRes) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type CreateSessionReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	RoleId    string `protobuf:"bytes,2,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	Ip        string `protobuf:"bytes,3,opt,name=ip,proto3" json:"ip,omitempty"`
	Data      string `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	ExpiresAt string `protobuf:"bytes,5,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
}

func (x *CreateSessionReq) Reset() {
	*x = CreateSessionReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_session_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSessionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSessionReq) ProtoMessage() {}

func (x *CreateSessionReq) ProtoReflect() protoreflect.Message {
	mi := &file_session_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSessionReq.ProtoReflect.Descriptor instead.
func (*CreateSessionReq) Descriptor() ([]byte, []int) {
	return file_session_proto_rawDescGZIP(), []int{11}
}

func (x *CreateSessionReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateSessionReq) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

func (x *CreateSessionReq) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *CreateSessionReq) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *CreateSessionReq) GetExpiresAt() string {
	if x != nil {
		return x.ExpiresAt
	}
	return ""
}

type UpdateSessionReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId    string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	RoleId    string `protobuf:"bytes,3,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	Ip        string `protobuf:"bytes,4,opt,name=ip,proto3" json:"ip,omitempty"`
	Data      string `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
	ExpiresAt string `protobuf:"bytes,6,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
}

func (x *UpdateSessionReq) Reset() {
	*x = UpdateSessionReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_session_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSessionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSessionReq) ProtoMessage() {}

func (x *UpdateSessionReq) ProtoReflect() protoreflect.Message {
	mi := &file_session_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSessionReq.ProtoReflect.Descriptor instead.
func (*UpdateSessionReq) Descriptor() ([]byte, []int) {
	return file_session_proto_rawDescGZIP(), []int{12}
}

func (x *UpdateSessionReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateSessionReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UpdateSessionReq) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

func (x *UpdateSessionReq) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *UpdateSessionReq) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *UpdateSessionReq) GetExpiresAt() string {
	if x != nil {
		return x.ExpiresAt
	}
	return ""
}

var File_session_proto protoreflect.FileDescriptor

var file_session_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0c, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x0a, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x53, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x78, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x78, 0x52, 0x6f, 0x6c, 0x65, 0x22, 0xf2, 0x01, 0x0a, 0x08,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x75, 0x73,
	0x65, 0x72, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x26, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12,
	0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x31, 0x0a, 0x08, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x28, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73,
	0x22, 0x2e, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x12, 0x21, 0x0a,
	0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x4f, 0x0a, 0x0f, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x65, 0x71, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72,
	0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49,
	0x64, 0x22, 0x3c, 0x0a, 0x0f, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x65, 0x73, 0x12, 0x29, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22,
	0x5d, 0x0a, 0x0c, 0x48, 0x61, 0x73, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x12,
	0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x22, 0xd1,
	0x01, 0x0a, 0x0c, 0x48, 0x61, 0x73, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49,
	0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x70, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73,
	0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72,
	0x65, 0x73, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x22, 0x4b, 0x0a, 0x12, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x41, 0x6e, 0x64,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x35, 0x0a, 0x0a, 0x6c, 0x6f, 0x67, 0x69,
	0x6e, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x52, 0x65, 0x73, 0x52, 0x09, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x22,
	0x23, 0x0a, 0x11, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72,
	0x79, 0x4b, 0x65, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x56, 0x0a, 0x0e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x47,
	0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x22, 0x5c, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x12, 0x31, 0x0a, 0x08, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x73, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x87, 0x01, 0x0a, 0x10, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49,
	0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x70, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73,
	0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72,
	0x65, 0x73, 0x41, 0x74, 0x22, 0x97, 0x01, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x32, 0x9f,
	0x02, 0x0a, 0x0e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x39, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x16, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x1a, 0x16, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x06,
	0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x12, 0x17, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x0c, 0x52, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x09, 0x48, 0x61, 0x73,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x48, 0x61, 0x73, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52,
	0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x48, 0x61, 0x73, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x22, 0x00,
	0x42, 0x17, 0x5a, 0x15, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_session_proto_rawDescOnce sync.Once
	file_session_proto_rawDescData = file_session_proto_rawDesc
)

func file_session_proto_rawDescGZIP() []byte {
	file_session_proto_rawDescOnce.Do(func() {
		file_session_proto_rawDescData = protoimpl.X.CompressGZIP(file_session_proto_rawDescData)
	})
	return file_session_proto_rawDescData
}

var file_session_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_session_proto_goTypes = []interface{}{
	(*LoginReq)(nil),           // 0: auth_service.LoginReq
	(*LoginRes)(nil),           // 1: auth_service.LoginRes
	(*LogoutReq)(nil),          // 2: auth_service.LogoutReq
	(*RefreshTokenReq)(nil),    // 3: auth_service.RefreshTokenReq
	(*RefreshTokenRes)(nil),    // 4: auth_service.RefreshTokenRes
	(*HasAccessReq)(nil),       // 5: auth_service.HasAccessReq
	(*HasAccessRes)(nil),       // 6: auth_service.HasAccessRes
	(*SessionAndTokenReq)(nil), // 7: auth_service.SessionAndTokenReq
	(*SessionPrimaryKey)(nil),  // 8: auth_service.SessionPrimaryKey
	(*SessionGetList)(nil),     // 9: auth_service.SessionGetList
	(*GetSessionListRes)(nil),  // 10: auth_service.GetSessionListRes
	(*CreateSessionReq)(nil),   // 11: auth_service.CreateSessionReq
	(*UpdateSessionReq)(nil),   // 12: auth_service.UpdateSessionReq
	(*User)(nil),               // 13: auth_service.User
	(*Token)(nil),              // 14: auth_service.Token
	(*Session)(nil),            // 15: auth_service.Session
	(*Role)(nil),               // 16: auth_service.Role
	(*emptypb.Empty)(nil),      // 17: google.protobuf.Empty
}
var file_session_proto_depIdxs = []int32{
	13, // 0: auth_service.LoginRes.user:type_name -> auth_service.User
	14, // 1: auth_service.LoginRes.token:type_name -> auth_service.Token
	15, // 2: auth_service.LoginRes.sessions:type_name -> auth_service.Session
	16, // 3: auth_service.LoginRes.roles:type_name -> auth_service.Role
	14, // 4: auth_service.RefreshTokenRes.token:type_name -> auth_service.Token
	1,  // 5: auth_service.SessionAndTokenReq.login_data:type_name -> auth_service.LoginRes
	15, // 6: auth_service.GetSessionListRes.sessions:type_name -> auth_service.Session
	0,  // 7: auth_service.SessionService.Login:input_type -> auth_service.LoginReq
	2,  // 8: auth_service.SessionService.Logout:input_type -> auth_service.LogoutReq
	3,  // 9: auth_service.SessionService.RefreshToken:input_type -> auth_service.RefreshTokenReq
	5,  // 10: auth_service.SessionService.HasAccess:input_type -> auth_service.HasAccessReq
	1,  // 11: auth_service.SessionService.Login:output_type -> auth_service.LoginRes
	17, // 12: auth_service.SessionService.Logout:output_type -> google.protobuf.Empty
	4,  // 13: auth_service.SessionService.RefreshToken:output_type -> auth_service.RefreshTokenRes
	6,  // 14: auth_service.SessionService.HasAccess:output_type -> auth_service.HasAccessRes
	11, // [11:15] is the sub-list for method output_type
	7,  // [7:11] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_session_proto_init() }
func file_session_proto_init() {
	if File_session_proto != nil {
		return
	}
	file_auth_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_session_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginReq); i {
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
		file_session_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRes); i {
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
		file_session_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogoutReq); i {
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
		file_session_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefreshTokenReq); i {
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
		file_session_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefreshTokenRes); i {
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
		file_session_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HasAccessReq); i {
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
		file_session_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HasAccessRes); i {
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
		file_session_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionAndTokenReq); i {
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
		file_session_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionPrimaryKey); i {
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
		file_session_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionGetList); i {
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
		file_session_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSessionListRes); i {
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
		file_session_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSessionReq); i {
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
		file_session_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSessionReq); i {
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
			RawDescriptor: file_session_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_session_proto_goTypes,
		DependencyIndexes: file_session_proto_depIdxs,
		MessageInfos:      file_session_proto_msgTypes,
	}.Build()
	File_session_proto = out.File
	file_session_proto_rawDesc = nil
	file_session_proto_goTypes = nil
	file_session_proto_depIdxs = nil
}
