// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.12.4
// source: protobuf/v1/auth.proto

package v1

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

type AuthoriseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JwtToken string `protobuf:"bytes,1,opt,name=JwtToken,proto3" json:"JwtToken,omitempty"`
	Email    string `protobuf:"bytes,2,opt,name=Email,proto3" json:"Email,omitempty"`
}

func (x *AuthoriseRequest) Reset() {
	*x = AuthoriseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_v1_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthoriseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthoriseRequest) ProtoMessage() {}

func (x *AuthoriseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_v1_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthoriseRequest.ProtoReflect.Descriptor instead.
func (*AuthoriseRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_v1_auth_proto_rawDescGZIP(), []int{0}
}

func (x *AuthoriseRequest) GetJwtToken() string {
	if x != nil {
		return x.JwtToken
	}
	return ""
}

func (x *AuthoriseRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type AuthoriseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObjectId string `protobuf:"bytes,1,opt,name=ObjectId,proto3" json:"ObjectId,omitempty"`
	Email    string `protobuf:"bytes,2,opt,name=Email,proto3" json:"Email,omitempty"`
	Name     string `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	Role     string `protobuf:"bytes,4,opt,name=Role,proto3" json:"Role,omitempty"`
}

func (x *AuthoriseResponse) Reset() {
	*x = AuthoriseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_v1_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthoriseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthoriseResponse) ProtoMessage() {}

func (x *AuthoriseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_v1_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthoriseResponse.ProtoReflect.Descriptor instead.
func (*AuthoriseResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_v1_auth_proto_rawDescGZIP(), []int{1}
}

func (x *AuthoriseResponse) GetObjectId() string {
	if x != nil {
		return x.ObjectId
	}
	return ""
}

func (x *AuthoriseResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *AuthoriseResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AuthoriseResponse) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

// get user details based on email Id or object id
type UserDetailsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObjectId string `protobuf:"bytes,1,opt,name=ObjectId,proto3" json:"ObjectId,omitempty"`
}

func (x *UserDetailsRequest) Reset() {
	*x = UserDetailsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_v1_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserDetailsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDetailsRequest) ProtoMessage() {}

func (x *UserDetailsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_v1_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDetailsRequest.ProtoReflect.Descriptor instead.
func (*UserDetailsRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_v1_auth_proto_rawDescGZIP(), []int{2}
}

func (x *UserDetailsRequest) GetObjectId() string {
	if x != nil {
		return x.ObjectId
	}
	return ""
}

// it should be no response or a complete user detail
type UserDetailsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error string `protobuf:"bytes,1,opt,name=Error,proto3" json:"Error,omitempty"`
	User  *User  `protobuf:"bytes,2,opt,name=User,proto3" json:"User,omitempty"`
}

func (x *UserDetailsResponse) Reset() {
	*x = UserDetailsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_v1_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserDetailsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDetailsResponse) ProtoMessage() {}

func (x *UserDetailsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_v1_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDetailsResponse.ProtoReflect.Descriptor instead.
func (*UserDetailsResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_v1_auth_proto_rawDescGZIP(), []int{3}
}

func (x *UserDetailsResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *UserDetailsResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObjectId string `protobuf:"bytes,1,opt,name=ObjectId,proto3" json:"ObjectId,omitempty"`
	Email    string `protobuf:"bytes,2,opt,name=Email,proto3" json:"Email,omitempty"`
	Phone    string `protobuf:"bytes,3,opt,name=Phone,proto3" json:"Phone,omitempty"`
	Name     string `protobuf:"bytes,4,opt,name=Name,proto3" json:"Name,omitempty"`
	Role     string `protobuf:"bytes,5,opt,name=Role,proto3" json:"Role,omitempty"`
	Standard string `protobuf:"bytes,6,opt,name=Standard,proto3" json:"Standard,omitempty"`
	City     string `protobuf:"bytes,7,opt,name=City,proto3" json:"City,omitempty"`
	State    string `protobuf:"bytes,8,opt,name=State,proto3" json:"State,omitempty"`
	Country  string `protobuf:"bytes,9,opt,name=Country,proto3" json:"Country,omitempty"`
	Address  string `protobuf:"bytes,10,opt,name=Address,proto3" json:"Address,omitempty"`
	Zip      string `protobuf:"bytes,11,opt,name=Zip,proto3" json:"Zip,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_v1_auth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_v1_auth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_protobuf_v1_auth_proto_rawDescGZIP(), []int{4}
}

func (x *User) GetObjectId() string {
	if x != nil {
		return x.ObjectId
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *User) GetStandard() string {
	if x != nil {
		return x.Standard
	}
	return ""
}

func (x *User) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *User) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *User) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *User) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *User) GetZip() string {
	if x != nil {
		return x.Zip
	}
	return ""
}

// issue book to user
type BookIssueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserJwtToken string `protobuf:"bytes,1,opt,name=UserJwtToken,proto3" json:"UserJwtToken,omitempty"` // i will get user details from this token only
	BookId       string `protobuf:"bytes,2,opt,name=BookId,proto3" json:"BookId,omitempty"`
}

func (x *BookIssueRequest) Reset() {
	*x = BookIssueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_v1_auth_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookIssueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookIssueRequest) ProtoMessage() {}

func (x *BookIssueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_v1_auth_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookIssueRequest.ProtoReflect.Descriptor instead.
func (*BookIssueRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_v1_auth_proto_rawDescGZIP(), []int{5}
}

func (x *BookIssueRequest) GetUserJwtToken() string {
	if x != nil {
		return x.UserJwtToken
	}
	return ""
}

func (x *BookIssueRequest) GetBookId() string {
	if x != nil {
		return x.BookId
	}
	return ""
}

// just response which tells if book is issued or not.
type BookIssueResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Issued bool   `protobuf:"varint,1,opt,name=Issued,proto3" json:"Issued,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=Error,proto3" json:"Error,omitempty"` // if not issued then what error occured
}

func (x *BookIssueResponse) Reset() {
	*x = BookIssueResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_v1_auth_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookIssueResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookIssueResponse) ProtoMessage() {}

func (x *BookIssueResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_v1_auth_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookIssueResponse.ProtoReflect.Descriptor instead.
func (*BookIssueResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_v1_auth_proto_rawDescGZIP(), []int{6}
}

func (x *BookIssueResponse) GetIssued() bool {
	if x != nil {
		return x.Issued
	}
	return false
}

func (x *BookIssueResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_protobuf_v1_auth_proto protoreflect.FileDescriptor

var file_protobuf_v1_auth_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x44, 0x0a, 0x10, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x4a, 0x77, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x4a, 0x77, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x6d,
	0x0a, 0x11, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x52, 0x6f, 0x6c,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x22, 0x30, 0x0a,
	0x12, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x22,
	0x46, 0x0a, 0x13, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x19, 0x0a, 0x04,
	0x55, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x04, 0x55, 0x73, 0x65, 0x72, 0x22, 0x82, 0x02, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x1a, 0x0a, 0x08, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x52, 0x6f, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x52, 0x6f, 0x6c, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x43, 0x69, 0x74, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x43, 0x69, 0x74, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x18, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x5a, 0x69,
	0x70, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x5a, 0x69, 0x70, 0x22, 0x4e, 0x0a, 0x10,
	0x42, 0x6f, 0x6f, 0x6b, 0x49, 0x73, 0x73, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x22, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x4a, 0x77, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x4a, 0x77, 0x74, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x42, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x42, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x22, 0x41, 0x0a, 0x11,
	0x42, 0x6f, 0x6f, 0x6b, 0x49, 0x73, 0x73, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x49, 0x73, 0x73, 0x75, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x49, 0x73, 0x73, 0x75, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x32,
	0xaf, 0x01, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x36, 0x0a, 0x0d, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x73, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x11, 0x2e, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3b, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x12, 0x13, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a,
	0x09, 0x42, 0x6f, 0x6f, 0x6b, 0x49, 0x73, 0x73, 0x75, 0x65, 0x12, 0x11, 0x2e, 0x42, 0x6f, 0x6f,
	0x6b, 0x49, 0x73, 0x73, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e,
	0x42, 0x6f, 0x6f, 0x6b, 0x49, 0x73, 0x73, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_protobuf_v1_auth_proto_rawDescOnce sync.Once
	file_protobuf_v1_auth_proto_rawDescData = file_protobuf_v1_auth_proto_rawDesc
)

func file_protobuf_v1_auth_proto_rawDescGZIP() []byte {
	file_protobuf_v1_auth_proto_rawDescOnce.Do(func() {
		file_protobuf_v1_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_v1_auth_proto_rawDescData)
	})
	return file_protobuf_v1_auth_proto_rawDescData
}

var file_protobuf_v1_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_protobuf_v1_auth_proto_goTypes = []interface{}{
	(*AuthoriseRequest)(nil),    // 0: AuthoriseRequest
	(*AuthoriseResponse)(nil),   // 1: AuthoriseResponse
	(*UserDetailsRequest)(nil),  // 2: UserDetailsRequest
	(*UserDetailsResponse)(nil), // 3: UserDetailsResponse
	(*User)(nil),                // 4: User
	(*BookIssueRequest)(nil),    // 5: BookIssueRequest
	(*BookIssueResponse)(nil),   // 6: BookIssueResponse
}
var file_protobuf_v1_auth_proto_depIdxs = []int32{
	4, // 0: UserDetailsResponse.User:type_name -> User
	0, // 1: Auth.AuthoriseUser:input_type -> AuthoriseRequest
	2, // 2: Auth.GetUserDetails:input_type -> UserDetailsRequest
	5, // 3: Auth.BookIssue:input_type -> BookIssueRequest
	1, // 4: Auth.AuthoriseUser:output_type -> AuthoriseResponse
	3, // 5: Auth.GetUserDetails:output_type -> UserDetailsResponse
	6, // 6: Auth.BookIssue:output_type -> BookIssueResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protobuf_v1_auth_proto_init() }
func file_protobuf_v1_auth_proto_init() {
	if File_protobuf_v1_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuf_v1_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthoriseRequest); i {
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
		file_protobuf_v1_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthoriseResponse); i {
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
		file_protobuf_v1_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserDetailsRequest); i {
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
		file_protobuf_v1_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserDetailsResponse); i {
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
		file_protobuf_v1_auth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_protobuf_v1_auth_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookIssueRequest); i {
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
		file_protobuf_v1_auth_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookIssueResponse); i {
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
			RawDescriptor: file_protobuf_v1_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protobuf_v1_auth_proto_goTypes,
		DependencyIndexes: file_protobuf_v1_auth_proto_depIdxs,
		MessageInfos:      file_protobuf_v1_auth_proto_msgTypes,
	}.Build()
	File_protobuf_v1_auth_proto = out.File
	file_protobuf_v1_auth_proto_rawDesc = nil
	file_protobuf_v1_auth_proto_goTypes = nil
	file_protobuf_v1_auth_proto_depIdxs = nil
}
