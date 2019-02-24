// Code generated by protoc-gen-go. DO NOT EDIT.
// source: talks.proto

package talks

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type Talk struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Repository           string   `protobuf:"bytes,4,opt,name=repository,proto3" json:"repository,omitempty"`
	Date                 int64    `protobuf:"varint,5,opt,name=date,proto3" json:"date,omitempty"`
	Tags                 string   `protobuf:"bytes,6,opt,name=tags,proto3" json:"tags,omitempty"`
	UserId               string   `protobuf:"bytes,7,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	CreatedAt            int64    `protobuf:"varint,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            int64    `protobuf:"varint,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Talk) Reset()         { *m = Talk{} }
func (m *Talk) String() string { return proto.CompactTextString(m) }
func (*Talk) ProtoMessage()    {}
func (*Talk) Descriptor() ([]byte, []int) {
	return fileDescriptor_6541a5a4ec1fb622, []int{0}
}

func (m *Talk) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Talk.Unmarshal(m, b)
}
func (m *Talk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Talk.Marshal(b, m, deterministic)
}
func (m *Talk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Talk.Merge(m, src)
}
func (m *Talk) XXX_Size() int {
	return xxx_messageInfo_Talk.Size(m)
}
func (m *Talk) XXX_DiscardUnknown() {
	xxx_messageInfo_Talk.DiscardUnknown(m)
}

var xxx_messageInfo_Talk proto.InternalMessageInfo

func (m *Talk) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Talk) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Talk) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Talk) GetRepository() string {
	if m != nil {
		return m.Repository
	}
	return ""
}

func (m *Talk) GetDate() int64 {
	if m != nil {
		return m.Date
	}
	return 0
}

func (m *Talk) GetTags() string {
	if m != nil {
		return m.Tags
	}
	return ""
}

func (m *Talk) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Talk) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *Talk) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

type GetRequest struct {
	TalkId               string   `protobuf:"bytes,1,opt,name=talk_id,json=talkId,proto3" json:"talk_id,omitempty"`
	UserId               string   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6541a5a4ec1fb622, []int{1}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetTalkId() string {
	if m != nil {
		return m.TalkId
	}
	return ""
}

func (m *GetRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type GetResponse struct {
	Talk                 *Talk    `protobuf:"bytes,1,opt,name=talk,proto3" json:"talk,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6541a5a4ec1fb622, []int{2}
}

func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (m *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(m, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetTalk() *Talk {
	if m != nil {
		return m.Talk
	}
	return nil
}

type SelectRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SelectRequest) Reset()         { *m = SelectRequest{} }
func (m *SelectRequest) String() string { return proto.CompactTextString(m) }
func (*SelectRequest) ProtoMessage()    {}
func (*SelectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6541a5a4ec1fb622, []int{3}
}

func (m *SelectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SelectRequest.Unmarshal(m, b)
}
func (m *SelectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SelectRequest.Marshal(b, m, deterministic)
}
func (m *SelectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SelectRequest.Merge(m, src)
}
func (m *SelectRequest) XXX_Size() int {
	return xxx_messageInfo_SelectRequest.Size(m)
}
func (m *SelectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SelectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SelectRequest proto.InternalMessageInfo

type SelectResponse struct {
	Talk                 []*Talk  `protobuf:"bytes,1,rep,name=talk,proto3" json:"talk,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SelectResponse) Reset()         { *m = SelectResponse{} }
func (m *SelectResponse) String() string { return proto.CompactTextString(m) }
func (*SelectResponse) ProtoMessage()    {}
func (*SelectResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6541a5a4ec1fb622, []int{4}
}

func (m *SelectResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SelectResponse.Unmarshal(m, b)
}
func (m *SelectResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SelectResponse.Marshal(b, m, deterministic)
}
func (m *SelectResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SelectResponse.Merge(m, src)
}
func (m *SelectResponse) XXX_Size() int {
	return xxx_messageInfo_SelectResponse.Size(m)
}
func (m *SelectResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SelectResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SelectResponse proto.InternalMessageInfo

func (m *SelectResponse) GetTalk() []*Talk {
	if m != nil {
		return m.Talk
	}
	return nil
}

type CreateRequest struct {
	Talk                 *Talk    `protobuf:"bytes,1,opt,name=talk,proto3" json:"talk,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6541a5a4ec1fb622, []int{5}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetTalk() *Talk {
	if m != nil {
		return m.Talk
	}
	return nil
}

type CreateResponse struct {
	Talk                 *Talk    `protobuf:"bytes,1,opt,name=talk,proto3" json:"talk,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6541a5a4ec1fb622, []int{6}
}

func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetTalk() *Talk {
	if m != nil {
		return m.Talk
	}
	return nil
}

type UpdateRequest struct {
	Data                 *Talk    `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6541a5a4ec1fb622, []int{7}
}

func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (m *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(m, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetData() *Talk {
	if m != nil {
		return m.Data
	}
	return nil
}

type UpdateResponse struct {
	Data                 *Talk    `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateResponse) Reset()         { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()    {}
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6541a5a4ec1fb622, []int{8}
}

func (m *UpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateResponse.Unmarshal(m, b)
}
func (m *UpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateResponse.Marshal(b, m, deterministic)
}
func (m *UpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateResponse.Merge(m, src)
}
func (m *UpdateResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateResponse.Size(m)
}
func (m *UpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateResponse proto.InternalMessageInfo

func (m *UpdateResponse) GetData() *Talk {
	if m != nil {
		return m.Data
	}
	return nil
}

type DeleteRequest struct {
	Data                 *Talk    `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6541a5a4ec1fb622, []int{9}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetData() *Talk {
	if m != nil {
		return m.Data
	}
	return nil
}

type DeleteResponse struct {
	Data                 *Talk    `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6541a5a4ec1fb622, []int{10}
}

func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (m *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(m, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

func (m *DeleteResponse) GetData() *Talk {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*Talk)(nil), "talks.Talk")
	proto.RegisterType((*GetRequest)(nil), "talks.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "talks.GetResponse")
	proto.RegisterType((*SelectRequest)(nil), "talks.SelectRequest")
	proto.RegisterType((*SelectResponse)(nil), "talks.SelectResponse")
	proto.RegisterType((*CreateRequest)(nil), "talks.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "talks.CreateResponse")
	proto.RegisterType((*UpdateRequest)(nil), "talks.UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "talks.UpdateResponse")
	proto.RegisterType((*DeleteRequest)(nil), "talks.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "talks.DeleteResponse")
}

func init() { proto.RegisterFile("talks.proto", fileDescriptor_6541a5a4ec1fb622) }

var fileDescriptor_6541a5a4ec1fb622 = []byte{
	// 411 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xc1, 0xee, 0xd2, 0x40,
	0x10, 0xc6, 0x69, 0x0b, 0x45, 0xa6, 0x29, 0xc6, 0x0d, 0xc6, 0x0d, 0x89, 0x4a, 0x7a, 0xe2, 0xd4,
	0x00, 0x1e, 0xbc, 0x99, 0x10, 0x4d, 0x08, 0xd7, 0xaa, 0x67, 0xb2, 0xb2, 0x13, 0xb2, 0xa1, 0xa1,
	0xb5, 0x3b, 0x1c, 0x7c, 0x18, 0x1f, 0xd1, 0x77, 0x30, 0xbb, 0xdb, 0x96, 0x56, 0x89, 0xf2, 0xbf,
	0x75, 0xbe, 0xd9, 0xdf, 0x7c, 0xbb, 0xdf, 0x24, 0x85, 0x88, 0x44, 0x7e, 0xd6, 0x69, 0x59, 0x15,
	0x54, 0xb0, 0x91, 0x2d, 0x92, 0x5f, 0x1e, 0x0c, 0xbf, 0x88, 0xfc, 0xcc, 0xa6, 0xe0, 0x2b, 0xc9,
	0xbd, 0x85, 0xb7, 0x9c, 0x64, 0xbe, 0x92, 0x6c, 0x06, 0x23, 0x52, 0x94, 0x23, 0xf7, 0xad, 0xe4,
	0x0a, 0xb6, 0x80, 0x48, 0xa2, 0x3e, 0x56, 0xaa, 0x24, 0x55, 0x5c, 0x78, 0x60, 0x7b, 0x5d, 0x89,
	0xbd, 0x01, 0xa8, 0xb0, 0x2c, 0xb4, 0xa2, 0xa2, 0xfa, 0xc1, 0x87, 0xf6, 0x40, 0x47, 0x61, 0x0c,
	0x86, 0x52, 0x10, 0xf2, 0xd1, 0xc2, 0x5b, 0x06, 0x99, 0xfd, 0x36, 0x1a, 0x89, 0x93, 0xe6, 0xa1,
	0x3d, 0x6d, 0xbf, 0xd9, 0x2b, 0x18, 0x5f, 0x35, 0x56, 0x07, 0x25, 0xf9, 0xd8, 0xca, 0xa1, 0x29,
	0xf7, 0x92, 0xbd, 0x06, 0x38, 0x56, 0x28, 0x08, 0xe5, 0x41, 0x10, 0x7f, 0x66, 0xc7, 0x4c, 0x6a,
	0x65, 0x4b, 0xa6, 0x7d, 0x2d, 0x65, 0xd3, 0x9e, 0xb8, 0x76, 0xad, 0x6c, 0x29, 0xf9, 0x00, 0xb0,
	0x43, 0xca, 0xf0, 0xfb, 0x15, 0x35, 0x19, 0x13, 0x13, 0xc3, 0xa1, 0x7d, 0x79, 0x68, 0xca, 0xbd,
	0xec, 0xba, 0xfb, 0x5d, 0xf7, 0x24, 0x85, 0xc8, 0xf2, 0xba, 0x2c, 0x2e, 0x1a, 0xd9, 0x5b, 0x73,
	0xf3, 0xfc, 0x6c, 0xe9, 0x68, 0x13, 0xa5, 0x2e, 0x61, 0x13, 0x68, 0x66, 0x1b, 0xc9, 0x73, 0x88,
	0x3f, 0x63, 0x8e, 0xc7, 0xc6, 0x32, 0x59, 0xc3, 0xb4, 0x11, 0xfe, 0x9a, 0x11, 0xdc, 0x9f, 0xb1,
	0x82, 0xf8, 0xa3, 0x7d, 0x5f, 0x73, 0xed, 0xff, 0xba, 0xae, 0x61, 0xda, 0x10, 0x8f, 0x5e, 0x74,
	0x05, 0xf1, 0x57, 0x9b, 0x52, 0xc7, 0x44, 0x0a, 0x12, 0x77, 0x09, 0xd3, 0x30, 0x26, 0x0d, 0x71,
	0x33, 0xf9, 0x37, 0xb2, 0x82, 0xf8, 0x13, 0xe6, 0xf8, 0x34, 0x93, 0x86, 0x78, 0xd0, 0x64, 0xf3,
	0xd3, 0x87, 0xb1, 0x29, 0xd5, 0xe5, 0xc4, 0x52, 0x08, 0x76, 0x48, 0xec, 0x45, 0x7d, 0xea, 0xb6,
	0xfa, 0x39, 0xeb, 0x4a, 0x6e, 0x74, 0x32, 0x60, 0xef, 0x21, 0x74, 0xdb, 0x61, 0xb3, 0xba, 0xdf,
	0xdb, 0xde, 0xfc, 0xe5, 0x1f, 0x6a, 0x17, 0x74, 0x89, 0xb7, 0x60, 0x6f, 0x65, 0x2d, 0xd8, 0x5f,
	0x8b, 0x03, 0x5d, 0x8a, 0x2d, 0xd8, 0x5b, 0x43, 0x0b, 0xf6, 0xa3, 0x76, 0xa0, 0x4b, 0xa6, 0x05,
	0x7b, 0xd1, 0xb6, 0x60, 0x3f, 0xbe, 0x64, 0xf0, 0x2d, 0xb4, 0x3f, 0x80, 0x77, 0xbf, 0x03, 0x00,
	0x00, 0xff, 0xff, 0x30, 0x19, 0x5f, 0x21, 0x0f, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TalkingClient is the client API for Talking service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TalkingClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	Select(ctx context.Context, in *SelectRequest, opts ...grpc.CallOption) (*SelectResponse, error)
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type talkingClient struct {
	cc *grpc.ClientConn
}

func NewTalkingClient(cc *grpc.ClientConn) TalkingClient {
	return &talkingClient{cc}
}

func (c *talkingClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/talks.Talking/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *talkingClient) Select(ctx context.Context, in *SelectRequest, opts ...grpc.CallOption) (*SelectResponse, error) {
	out := new(SelectResponse)
	err := c.cc.Invoke(ctx, "/talks.Talking/Select", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *talkingClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/talks.Talking/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *talkingClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/talks.Talking/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *talkingClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/talks.Talking/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TalkingServer is the server API for Talking service.
type TalkingServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
	Select(context.Context, *SelectRequest) (*SelectResponse, error)
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
}

func RegisterTalkingServer(s *grpc.Server, srv TalkingServer) {
	s.RegisterService(&_Talking_serviceDesc, srv)
}

func _Talking_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TalkingServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/talks.Talking/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TalkingServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Talking_Select_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SelectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TalkingServer).Select(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/talks.Talking/Select",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TalkingServer).Select(ctx, req.(*SelectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Talking_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TalkingServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/talks.Talking/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TalkingServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Talking_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TalkingServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/talks.Talking/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TalkingServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Talking_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TalkingServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/talks.Talking/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TalkingServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Talking_serviceDesc = grpc.ServiceDesc{
	ServiceName: "talks.Talking",
	HandlerType: (*TalkingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Talking_Get_Handler,
		},
		{
			MethodName: "Select",
			Handler:    _Talking_Select_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Talking_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Talking_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Talking_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "talks.proto",
}
