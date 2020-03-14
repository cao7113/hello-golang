// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user/user.proto

package pbuser

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type User struct {
	Name                 string                `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email                string                `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	IsBoy                bool                  `protobuf:"varint,3,opt,name=is_boy,json=isBoy,proto3" json:"is_boy,omitempty"`
	IsGirl               *wrappers.BoolValue   `protobuf:"bytes,4,opt,name=is_girl,json=isGirl,proto3" json:"is_girl,omitempty"`
	Job                  *wrappers.StringValue `protobuf:"bytes,5,opt,name=job,proto3" json:"job,omitempty"`
	UpdatedAt            *timestamp.Timestamp  `protobuf:"bytes,11,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed89022014131a74, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetIsBoy() bool {
	if m != nil {
		return m.IsBoy
	}
	return false
}

func (m *User) GetIsGirl() *wrappers.BoolValue {
	if m != nil {
		return m.IsGirl
	}
	return nil
}

func (m *User) GetJob() *wrappers.StringValue {
	if m != nil {
		return m.Job
	}
	return nil
}

func (m *User) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "pb.user.User")
}

func init() { proto.RegisterFile("user/user.proto", fileDescriptor_ed89022014131a74) }

var fileDescriptor_ed89022014131a74 = []byte{
	// 248 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x4f, 0x4b, 0xc4, 0x30,
	0x14, 0xc4, 0x89, 0xdb, 0x76, 0x77, 0xb3, 0x07, 0x21, 0x28, 0x84, 0x22, 0x5a, 0x3c, 0xf5, 0x94,
	0x82, 0x7b, 0xf2, 0x68, 0x2f, 0xde, 0xeb, 0x9f, 0x83, 0x97, 0x92, 0xb0, 0xb1, 0x3c, 0x49, 0xf7,
	0x85, 0x24, 0x45, 0xf6, 0x13, 0xfb, 0x35, 0xa4, 0x89, 0xbd, 0x28, 0x5e, 0x42, 0x66, 0xde, 0x6f,
	0x18, 0x18, 0x7a, 0x3e, 0x79, 0xed, 0x9a, 0xf9, 0x11, 0xd6, 0x61, 0x40, 0xb6, 0xb6, 0x4a, 0xcc,
	0xb2, 0xbc, 0x19, 0x10, 0x07, 0xa3, 0x9b, 0x68, 0xab, 0xe9, 0xbd, 0x09, 0x30, 0x6a, 0x1f, 0xe4,
	0x68, 0x13, 0x59, 0x5e, 0xff, 0x06, 0x3e, 0x9d, 0xb4, 0x56, 0x3b, 0x9f, 0xee, 0xb7, 0x5f, 0x84,
	0x66, 0x2f, 0x5e, 0x3b, 0xc6, 0x68, 0x76, 0x94, 0xa3, 0xe6, 0xa4, 0x22, 0xf5, 0xb6, 0x8b, 0x7f,
	0x76, 0x41, 0x73, 0x3d, 0x4a, 0x30, 0xfc, 0x2c, 0x9a, 0x49, 0xb0, 0x4b, 0x5a, 0x80, 0xef, 0x15,
	0x9e, 0xf8, 0xaa, 0x22, 0xf5, 0xa6, 0xcb, 0xc1, 0xb7, 0x78, 0x62, 0x7b, 0xba, 0x06, 0xdf, 0x0f,
	0xe0, 0x0c, 0xcf, 0x2a, 0x52, 0xef, 0xee, 0x4a, 0x91, 0xba, 0xc5, 0xd2, 0x2d, 0x5a, 0x44, 0xf3,
	0x2a, 0xcd, 0xa4, 0xbb, 0x02, 0xfc, 0x23, 0x38, 0xc3, 0x04, 0x5d, 0x7d, 0xa0, 0xe2, 0x79, 0x0c,
	0x5c, 0xfd, 0x09, 0x3c, 0x05, 0x07, 0xc7, 0x21, 0x45, 0x66, 0x90, 0xdd, 0x53, 0x3a, 0xd9, 0x83,
	0x0c, 0xfa, 0xd0, 0xcb, 0xc0, 0x77, 0xff, 0xf4, 0x3c, 0x2f, 0x23, 0x74, 0xdb, 0x1f, 0xfa, 0x21,
	0xb4, 0x9b, 0xb7, 0xc2, 0xaa, 0x79, 0x34, 0x55, 0x44, 0x70, 0xff, 0x1d, 0x00, 0x00, 0xff, 0xff,
	0xe8, 0x7d, 0x31, 0x37, 0x57, 0x01, 0x00, 0x00,
}
