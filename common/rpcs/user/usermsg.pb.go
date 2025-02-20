// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: source/rpc/asset-user/usermsg.proto

package user

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

type PageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit       int64  `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset      int64  `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	SearchCount bool   `protobuf:"varint,3,opt,name=searchCount,proto3" json:"searchCount,omitempty"`
	Filter      string `protobuf:"bytes,4,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *PageRequest) Reset() {
	*x = PageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_source_rpc_asset_user_usermsg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageRequest) ProtoMessage() {}

func (x *PageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_source_rpc_asset_user_usermsg_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageRequest.ProtoReflect.Descriptor instead.
func (*PageRequest) Descriptor() ([]byte, []int) {
	return file_source_rpc_asset_user_usermsg_proto_rawDescGZIP(), []int{0}
}

func (x *PageRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *PageRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *PageRequest) GetSearchCount() bool {
	if x != nil {
		return x.SearchCount
	}
	return false
}

func (x *PageRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

type UserByAccountReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//标识
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	//账户
	Account string `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
	//密码
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *UserByAccountReq) Reset() {
	*x = UserByAccountReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_source_rpc_asset_user_usermsg_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserByAccountReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserByAccountReq) ProtoMessage() {}

func (x *UserByAccountReq) ProtoReflect() protoreflect.Message {
	mi := &file_source_rpc_asset_user_usermsg_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserByAccountReq.ProtoReflect.Descriptor instead.
func (*UserByAccountReq) Descriptor() ([]byte, []int) {
	return file_source_rpc_asset_user_usermsg_proto_rawDescGZIP(), []int{1}
}

func (x *UserByAccountReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserByAccountReq) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *UserByAccountReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type CommonRpcRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//返回Json
	Json string `protobuf:"bytes,1,opt,name=json,proto3" json:"json,omitempty"`
}

func (x *CommonRpcRes) Reset() {
	*x = CommonRpcRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_source_rpc_asset_user_usermsg_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonRpcRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonRpcRes) ProtoMessage() {}

func (x *CommonRpcRes) ProtoReflect() protoreflect.Message {
	mi := &file_source_rpc_asset_user_usermsg_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonRpcRes.ProtoReflect.Descriptor instead.
func (*CommonRpcRes) Descriptor() ([]byte, []int) {
	return file_source_rpc_asset_user_usermsg_proto_rawDescGZIP(), []int{2}
}

func (x *CommonRpcRes) GetJson() string {
	if x != nil {
		return x.Json
	}
	return ""
}

type UserReRegReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//电话
	PhoneNumber string `protobuf:"bytes,1,opt,name=phoneNumber,proto3" json:"phoneNumber,omitempty"`
	//姓名
	RealName string `protobuf:"bytes,2,opt,name=realName,proto3" json:"realName,omitempty"`
	//密码
	Pwd string `protobuf:"bytes,3,opt,name=pwd,proto3" json:"pwd,omitempty"`
}

func (x *UserReRegReq) Reset() {
	*x = UserReRegReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_source_rpc_asset_user_usermsg_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserReRegReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserReRegReq) ProtoMessage() {}

func (x *UserReRegReq) ProtoReflect() protoreflect.Message {
	mi := &file_source_rpc_asset_user_usermsg_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserReRegReq.ProtoReflect.Descriptor instead.
func (*UserReRegReq) Descriptor() ([]byte, []int) {
	return file_source_rpc_asset_user_usermsg_proto_rawDescGZIP(), []int{3}
}

func (x *UserReRegReq) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *UserReRegReq) GetRealName() string {
	if x != nil {
		return x.RealName
	}
	return ""
}

func (x *UserReRegReq) GetPwd() string {
	if x != nil {
		return x.Pwd
	}
	return ""
}

type UpdatePasswdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//用户ID
	UserId int64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	//旧密码
	OldPwd string `protobuf:"bytes,2,opt,name=oldPwd,proto3" json:"oldPwd,omitempty"`
	//新密码
	NewPwd string `protobuf:"bytes,3,opt,name=newPwd,proto3" json:"newPwd,omitempty"`
}

func (x *UpdatePasswdReq) Reset() {
	*x = UpdatePasswdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_source_rpc_asset_user_usermsg_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePasswdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePasswdReq) ProtoMessage() {}

func (x *UpdatePasswdReq) ProtoReflect() protoreflect.Message {
	mi := &file_source_rpc_asset_user_usermsg_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePasswdReq.ProtoReflect.Descriptor instead.
func (*UpdatePasswdReq) Descriptor() ([]byte, []int) {
	return file_source_rpc_asset_user_usermsg_proto_rawDescGZIP(), []int{4}
}

func (x *UpdatePasswdReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UpdatePasswdReq) GetOldPwd() string {
	if x != nil {
		return x.OldPwd
	}
	return ""
}

func (x *UpdatePasswdReq) GetNewPwd() string {
	if x != nil {
		return x.NewPwd
	}
	return ""
}

type TenantCodeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//用户ID
	UserId int64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	//账户
	Account string `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
	//租户编号
	TenantCode string `protobuf:"bytes,3,opt,name=tenantCode,proto3" json:"tenantCode,omitempty"`
}

func (x *TenantCodeReq) Reset() {
	*x = TenantCodeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_source_rpc_asset_user_usermsg_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TenantCodeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TenantCodeReq) ProtoMessage() {}

func (x *TenantCodeReq) ProtoReflect() protoreflect.Message {
	mi := &file_source_rpc_asset_user_usermsg_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TenantCodeReq.ProtoReflect.Descriptor instead.
func (*TenantCodeReq) Descriptor() ([]byte, []int) {
	return file_source_rpc_asset_user_usermsg_proto_rawDescGZIP(), []int{5}
}

func (x *TenantCodeReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *TenantCodeReq) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *TenantCodeReq) GetTenantCode() string {
	if x != nil {
		return x.TenantCode
	}
	return ""
}

type FindUserListPageReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//人员ID
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	//人员名称
	UserName string `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty"`
	//是否删除
	IsDeleted int64 `protobuf:"varint,3,opt,name=isDeleted,proto3" json:"isDeleted,omitempty"`
	//电话号码
	PhoneNumber string `protobuf:"bytes,4,opt,name=phoneNumber,proto3" json:"phoneNumber,omitempty"`
	//租户Code
	TenantCode string `protobuf:"bytes,5,opt,name=tenantCode,proto3" json:"tenantCode,omitempty"`
	//分页参数
	Page *PageRequest `protobuf:"bytes,6,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *FindUserListPageReq) Reset() {
	*x = FindUserListPageReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_source_rpc_asset_user_usermsg_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindUserListPageReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindUserListPageReq) ProtoMessage() {}

func (x *FindUserListPageReq) ProtoReflect() protoreflect.Message {
	mi := &file_source_rpc_asset_user_usermsg_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindUserListPageReq.ProtoReflect.Descriptor instead.
func (*FindUserListPageReq) Descriptor() ([]byte, []int) {
	return file_source_rpc_asset_user_usermsg_proto_rawDescGZIP(), []int{6}
}

func (x *FindUserListPageReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FindUserListPageReq) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *FindUserListPageReq) GetIsDeleted() int64 {
	if x != nil {
		return x.IsDeleted
	}
	return 0
}

func (x *FindUserListPageReq) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *FindUserListPageReq) GetTenantCode() string {
	if x != nil {
		return x.TenantCode
	}
	return ""
}

func (x *FindUserListPageReq) GetPage() *PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

type SwitchTenantReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//账户
	Account string `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
	//单位编号
	TenantCode string `protobuf:"bytes,3,opt,name=tenantCode,proto3" json:"tenantCode,omitempty"`
}

func (x *SwitchTenantReq) Reset() {
	*x = SwitchTenantReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_source_rpc_asset_user_usermsg_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SwitchTenantReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SwitchTenantReq) ProtoMessage() {}

func (x *SwitchTenantReq) ProtoReflect() protoreflect.Message {
	mi := &file_source_rpc_asset_user_usermsg_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SwitchTenantReq.ProtoReflect.Descriptor instead.
func (*SwitchTenantReq) Descriptor() ([]byte, []int) {
	return file_source_rpc_asset_user_usermsg_proto_rawDescGZIP(), []int{7}
}

func (x *SwitchTenantReq) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *SwitchTenantReq) GetTenantCode() string {
	if x != nil {
		return x.TenantCode
	}
	return ""
}

type FindserTenantUserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//人员ID
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	//人员名称
	UserName string `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty"`
	//是否删除
	TenantApplyinStateRange []int64 `protobuf:"varint,3,rep,packed,name=tenantApplyinStateRange,proto3" json:"tenantApplyinStateRange,omitempty"`
	//租户Code
	TenantCode string `protobuf:"bytes,4,opt,name=tenantCode,proto3" json:"tenantCode,omitempty"`
	//分页参数
	Page *PageRequest `protobuf:"bytes,5,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *FindserTenantUserReq) Reset() {
	*x = FindserTenantUserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_source_rpc_asset_user_usermsg_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindserTenantUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindserTenantUserReq) ProtoMessage() {}

func (x *FindserTenantUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_source_rpc_asset_user_usermsg_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindserTenantUserReq.ProtoReflect.Descriptor instead.
func (*FindserTenantUserReq) Descriptor() ([]byte, []int) {
	return file_source_rpc_asset_user_usermsg_proto_rawDescGZIP(), []int{8}
}

func (x *FindserTenantUserReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FindserTenantUserReq) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *FindserTenantUserReq) GetTenantApplyinStateRange() []int64 {
	if x != nil {
		return x.TenantApplyinStateRange
	}
	return nil
}

func (x *FindserTenantUserReq) GetTenantCode() string {
	if x != nil {
		return x.TenantCode
	}
	return ""
}

func (x *FindserTenantUserReq) GetPage() *PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

type DeleteUserInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//人员ID列表
	UserIds []int64 `protobuf:"varint,1,rep,packed,name=userIds,proto3" json:"userIds,omitempty"`
	//租户编号
	TenantCode string `protobuf:"bytes,2,opt,name=tenantCode,proto3" json:"tenantCode,omitempty"`
}

func (x *DeleteUserInfoReq) Reset() {
	*x = DeleteUserInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_source_rpc_asset_user_usermsg_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserInfoReq) ProtoMessage() {}

func (x *DeleteUserInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_source_rpc_asset_user_usermsg_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserInfoReq.ProtoReflect.Descriptor instead.
func (*DeleteUserInfoReq) Descriptor() ([]byte, []int) {
	return file_source_rpc_asset_user_usermsg_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteUserInfoReq) GetUserIds() []int64 {
	if x != nil {
		return x.UserIds
	}
	return nil
}

func (x *DeleteUserInfoReq) GetTenantCode() string {
	if x != nil {
		return x.TenantCode
	}
	return ""
}

type AuditUserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//人员ID列表
	UserIds []int64 `protobuf:"varint,1,rep,packed,name=userIds,proto3" json:"userIds,omitempty"`
	//是否通过
	IsPass bool `protobuf:"varint,2,opt,name=isPass,proto3" json:"isPass,omitempty"`
	//租户编号
	TenantCode string `protobuf:"bytes,3,opt,name=tenantCode,proto3" json:"tenantCode,omitempty"`
}

func (x *AuditUserReq) Reset() {
	*x = AuditUserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_source_rpc_asset_user_usermsg_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuditUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuditUserReq) ProtoMessage() {}

func (x *AuditUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_source_rpc_asset_user_usermsg_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuditUserReq.ProtoReflect.Descriptor instead.
func (*AuditUserReq) Descriptor() ([]byte, []int) {
	return file_source_rpc_asset_user_usermsg_proto_rawDescGZIP(), []int{10}
}

func (x *AuditUserReq) GetUserIds() []int64 {
	if x != nil {
		return x.UserIds
	}
	return nil
}

func (x *AuditUserReq) GetIsPass() bool {
	if x != nil {
		return x.IsPass
	}
	return false
}

func (x *AuditUserReq) GetTenantCode() string {
	if x != nil {
		return x.TenantCode
	}
	return ""
}

type CancelJoinTenantReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//账户
	Account string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	//单位编号
	TenantCodes []string `protobuf:"bytes,2,rep,name=tenantCodes,proto3" json:"tenantCodes,omitempty"`
}

func (x *CancelJoinTenantReq) Reset() {
	*x = CancelJoinTenantReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_source_rpc_asset_user_usermsg_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelJoinTenantReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelJoinTenantReq) ProtoMessage() {}

func (x *CancelJoinTenantReq) ProtoReflect() protoreflect.Message {
	mi := &file_source_rpc_asset_user_usermsg_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelJoinTenantReq.ProtoReflect.Descriptor instead.
func (*CancelJoinTenantReq) Descriptor() ([]byte, []int) {
	return file_source_rpc_asset_user_usermsg_proto_rawDescGZIP(), []int{11}
}

func (x *CancelJoinTenantReq) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *CancelJoinTenantReq) GetTenantCodes() []string {
	if x != nil {
		return x.TenantCodes
	}
	return nil
}

type RemoveUserByIdsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//用户ID
	Ids []int64 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (x *RemoveUserByIdsReq) Reset() {
	*x = RemoveUserByIdsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_source_rpc_asset_user_usermsg_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveUserByIdsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveUserByIdsReq) ProtoMessage() {}

func (x *RemoveUserByIdsReq) ProtoReflect() protoreflect.Message {
	mi := &file_source_rpc_asset_user_usermsg_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveUserByIdsReq.ProtoReflect.Descriptor instead.
func (*RemoveUserByIdsReq) Descriptor() ([]byte, []int) {
	return file_source_rpc_asset_user_usermsg_proto_rawDescGZIP(), []int{12}
}

func (x *RemoveUserByIdsReq) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type GetMarketAppInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//应用key
	Appkey string `protobuf:"bytes,1,opt,name=appkey,proto3" json:"appkey,omitempty"`
	//应用secret
	Appsecret string `protobuf:"bytes,2,opt,name=appsecret,proto3" json:"appsecret,omitempty"`
}

func (x *GetMarketAppInfoReq) Reset() {
	*x = GetMarketAppInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_source_rpc_asset_user_usermsg_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMarketAppInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMarketAppInfoReq) ProtoMessage() {}

func (x *GetMarketAppInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_source_rpc_asset_user_usermsg_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMarketAppInfoReq.ProtoReflect.Descriptor instead.
func (*GetMarketAppInfoReq) Descriptor() ([]byte, []int) {
	return file_source_rpc_asset_user_usermsg_proto_rawDescGZIP(), []int{13}
}

func (x *GetMarketAppInfoReq) GetAppkey() string {
	if x != nil {
		return x.Appkey
	}
	return ""
}

func (x *GetMarketAppInfoReq) GetAppsecret() string {
	if x != nil {
		return x.Appsecret
	}
	return ""
}

var File_source_rpc_asset_user_usermsg_proto protoreflect.FileDescriptor

var file_source_rpc_asset_user_usermsg_proto_rawDesc = []byte{
	0x0a, 0x23, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x2d, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x6d, 0x73, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x75, 0x0a, 0x0b, 0x50,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x73,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x22, 0x58, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x22, 0x0a, 0x0c,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x70, 0x63, 0x52, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x6a, 0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6a, 0x73, 0x6f, 0x6e,
	0x22, 0x5e, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x52, 0x65, 0x67, 0x52, 0x65, 0x71,
	0x12, 0x20, 0x0a, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x70, 0x77, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x77, 0x64,
	0x22, 0x59, 0x0a, 0x0f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x64,
	0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6f,
	0x6c, 0x64, 0x50, 0x77, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x6c, 0x64,
	0x50, 0x77, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x65, 0x77, 0x50, 0x77, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x65, 0x77, 0x50, 0x77, 0x64, 0x22, 0x61, 0x0a, 0x0d, 0x54,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e,
	0x0a, 0x0a, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x22, 0xc8,
	0x01, 0x0a, 0x13, 0x46, 0x69, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x69, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x12, 0x20, 0x0a, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x25, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0x4b, 0x0a, 0x0f, 0x53, 0x77, 0x69,
	0x74, 0x63, 0x68, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x43, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x22, 0xc3, 0x01, 0x0a, 0x14, 0x46, 0x69, 0x6e, 0x64, 0x73,
	0x65, 0x72, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x38, 0x0a, 0x17, 0x74,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x69, 0x6e, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x03, 0x52, 0x17, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x69, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x43,
	0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x25, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0x4d, 0x0a, 0x11,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x71, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x03, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x74,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x60, 0x0a, 0x0c, 0x41,
	0x75, 0x64, 0x69, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x50, 0x61, 0x73, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x50, 0x61, 0x73, 0x73, 0x12, 0x1e, 0x0a,
	0x0a, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x51, 0x0a,
	0x13, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4a, 0x6f, 0x69, 0x6e, 0x54, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x20,
	0x0a, 0x0b, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x73,
	0x22, 0x26, 0x0a, 0x12, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79,
	0x49, 0x64, 0x73, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x4b, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4d,
	0x61, 0x72, 0x6b, 0x65, 0x74, 0x41, 0x70, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x61, 0x70, 0x70, 0x6b, 0x65, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x70, 0x70, 0x73, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x70, 0x70, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_source_rpc_asset_user_usermsg_proto_rawDescOnce sync.Once
	file_source_rpc_asset_user_usermsg_proto_rawDescData = file_source_rpc_asset_user_usermsg_proto_rawDesc
)

func file_source_rpc_asset_user_usermsg_proto_rawDescGZIP() []byte {
	file_source_rpc_asset_user_usermsg_proto_rawDescOnce.Do(func() {
		file_source_rpc_asset_user_usermsg_proto_rawDescData = protoimpl.X.CompressGZIP(file_source_rpc_asset_user_usermsg_proto_rawDescData)
	})
	return file_source_rpc_asset_user_usermsg_proto_rawDescData
}

var file_source_rpc_asset_user_usermsg_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_source_rpc_asset_user_usermsg_proto_goTypes = []interface{}{
	(*PageRequest)(nil),          // 0: user.PageRequest
	(*UserByAccountReq)(nil),     // 1: user.UserByAccountReq
	(*CommonRpcRes)(nil),         // 2: user.CommonRpcRes
	(*UserReRegReq)(nil),         // 3: user.UserReRegReq
	(*UpdatePasswdReq)(nil),      // 4: user.updatePasswdReq
	(*TenantCodeReq)(nil),        // 5: user.TenantCodeReq
	(*FindUserListPageReq)(nil),  // 6: user.FindUserListPageReq
	(*SwitchTenantReq)(nil),      // 7: user.SwitchTenantReq
	(*FindserTenantUserReq)(nil), // 8: user.FindserTenantUserReq
	(*DeleteUserInfoReq)(nil),    // 9: user.DeleteUserInfoReq
	(*AuditUserReq)(nil),         // 10: user.AuditUserReq
	(*CancelJoinTenantReq)(nil),  // 11: user.CancelJoinTenantReq
	(*RemoveUserByIdsReq)(nil),   // 12: user.RemoveUserByIdsReq
	(*GetMarketAppInfoReq)(nil),  // 13: user.GetMarketAppInfoReq
}
var file_source_rpc_asset_user_usermsg_proto_depIdxs = []int32{
	0, // 0: user.FindUserListPageReq.page:type_name -> user.PageRequest
	0, // 1: user.FindserTenantUserReq.page:type_name -> user.PageRequest
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_source_rpc_asset_user_usermsg_proto_init() }
func file_source_rpc_asset_user_usermsg_proto_init() {
	if File_source_rpc_asset_user_usermsg_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_source_rpc_asset_user_usermsg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageRequest); i {
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
		file_source_rpc_asset_user_usermsg_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserByAccountReq); i {
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
		file_source_rpc_asset_user_usermsg_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonRpcRes); i {
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
		file_source_rpc_asset_user_usermsg_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserReRegReq); i {
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
		file_source_rpc_asset_user_usermsg_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePasswdReq); i {
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
		file_source_rpc_asset_user_usermsg_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TenantCodeReq); i {
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
		file_source_rpc_asset_user_usermsg_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindUserListPageReq); i {
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
		file_source_rpc_asset_user_usermsg_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SwitchTenantReq); i {
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
		file_source_rpc_asset_user_usermsg_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindserTenantUserReq); i {
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
		file_source_rpc_asset_user_usermsg_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserInfoReq); i {
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
		file_source_rpc_asset_user_usermsg_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuditUserReq); i {
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
		file_source_rpc_asset_user_usermsg_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelJoinTenantReq); i {
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
		file_source_rpc_asset_user_usermsg_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveUserByIdsReq); i {
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
		file_source_rpc_asset_user_usermsg_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMarketAppInfoReq); i {
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
			RawDescriptor: file_source_rpc_asset_user_usermsg_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_source_rpc_asset_user_usermsg_proto_goTypes,
		DependencyIndexes: file_source_rpc_asset_user_usermsg_proto_depIdxs,
		MessageInfos:      file_source_rpc_asset_user_usermsg_proto_msgTypes,
	}.Build()
	File_source_rpc_asset_user_usermsg_proto = out.File
	file_source_rpc_asset_user_usermsg_proto_rawDesc = nil
	file_source_rpc_asset_user_usermsg_proto_goTypes = nil
	file_source_rpc_asset_user_usermsg_proto_depIdxs = nil
}
