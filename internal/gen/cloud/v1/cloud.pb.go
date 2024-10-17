// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: cloud/v1/cloud.proto

package cloudv1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

// UserProfile represents a GitHub user's profile
type Profile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login     string `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	Id        int32  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	AvatarUrl string `protobuf:"bytes,3,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`
	HtmlUrl   string `protobuf:"bytes,4,opt,name=html_url,json=htmlUrl,proto3" json:"html_url,omitempty"`
	Type      string `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *Profile) Reset() {
	*x = Profile{}
	mi := &file_cloud_v1_cloud_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Profile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Profile) ProtoMessage() {}

func (x *Profile) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_v1_cloud_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Profile.ProtoReflect.Descriptor instead.
func (*Profile) Descriptor() ([]byte, []int) {
	return file_cloud_v1_cloud_proto_rawDescGZIP(), []int{0}
}

func (x *Profile) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *Profile) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Profile) GetAvatarUrl() string {
	if x != nil {
		return x.AvatarUrl
	}
	return ""
}

func (x *Profile) GetHtmlUrl() string {
	if x != nil {
		return x.HtmlUrl
	}
	return ""
}

func (x *Profile) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

// Repository represents a GitHub repository
type Repository struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	FullName string `protobuf:"bytes,3,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Private  bool   `protobuf:"varint,4,opt,name=private,proto3" json:"private,omitempty"`
	Owner    int32  `protobuf:"varint,5,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (x *Repository) Reset() {
	*x = Repository{}
	mi := &file_cloud_v1_cloud_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Repository) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Repository) ProtoMessage() {}

func (x *Repository) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_v1_cloud_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Repository.ProtoReflect.Descriptor instead.
func (*Repository) Descriptor() ([]byte, []int) {
	return file_cloud_v1_cloud_proto_rawDescGZIP(), []int{1}
}

func (x *Repository) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Repository) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Repository) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *Repository) GetPrivate() bool {
	if x != nil {
		return x.Private
	}
	return false
}

func (x *Repository) GetOwner() int32 {
	if x != nil {
		return x.Owner
	}
	return 0
}

// Commit represents a GitHub commit
type Commit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sha    string                `protobuf:"bytes,1,opt,name=sha,proto3" json:"sha,omitempty"`
	Commit *Commit_CommitDetails `protobuf:"bytes,2,opt,name=commit,proto3" json:"commit,omitempty"`
}

func (x *Commit) Reset() {
	*x = Commit{}
	mi := &file_cloud_v1_cloud_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Commit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Commit) ProtoMessage() {}

func (x *Commit) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_v1_cloud_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Commit.ProtoReflect.Descriptor instead.
func (*Commit) Descriptor() ([]byte, []int) {
	return file_cloud_v1_cloud_proto_rawDescGZIP(), []int{2}
}

func (x *Commit) GetSha() string {
	if x != nil {
		return x.Sha
	}
	return ""
}

func (x *Commit) GetCommit() *Commit_CommitDetails {
	if x != nil {
		return x.Commit
	}
	return nil
}

type Commit_CommitDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Author  *Commit_CommitDetails_Author `protobuf:"bytes,1,opt,name=author,proto3" json:"author,omitempty"`
	Message string                       `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Commit_CommitDetails) Reset() {
	*x = Commit_CommitDetails{}
	mi := &file_cloud_v1_cloud_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Commit_CommitDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Commit_CommitDetails) ProtoMessage() {}

func (x *Commit_CommitDetails) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_v1_cloud_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Commit_CommitDetails.ProtoReflect.Descriptor instead.
func (*Commit_CommitDetails) Descriptor() ([]byte, []int) {
	return file_cloud_v1_cloud_proto_rawDescGZIP(), []int{2, 0}
}

func (x *Commit_CommitDetails) GetAuthor() *Commit_CommitDetails_Author {
	if x != nil {
		return x.Author
	}
	return nil
}

func (x *Commit_CommitDetails) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Commit_CommitDetails_Author struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Date  string `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *Commit_CommitDetails_Author) Reset() {
	*x = Commit_CommitDetails_Author{}
	mi := &file_cloud_v1_cloud_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Commit_CommitDetails_Author) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Commit_CommitDetails_Author) ProtoMessage() {}

func (x *Commit_CommitDetails_Author) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_v1_cloud_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Commit_CommitDetails_Author.ProtoReflect.Descriptor instead.
func (*Commit_CommitDetails_Author) Descriptor() ([]byte, []int) {
	return file_cloud_v1_cloud_proto_rawDescGZIP(), []int{2, 0, 0}
}

func (x *Commit_CommitDetails_Author) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Commit_CommitDetails_Author) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Commit_CommitDetails_Author) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

var File_cloud_v1_cloud_proto protoreflect.FileDescriptor

var file_cloud_v1_cloud_proto_rawDesc = []byte{
	0x0a, 0x14, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31,
	0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7d, 0x0a, 0x07, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x68, 0x74, 0x6d,
	0x6c, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x68, 0x74, 0x6d,
	0x6c, 0x55, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x7d, 0x0a, 0x0a, 0x52, 0x65, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75,
	0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66,
	0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x22, 0x85, 0x02, 0x0a, 0x06, 0x43, 0x6f, 0x6d, 0x6d,
	0x69, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x68, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x73, 0x68, 0x61, 0x12, 0x36, 0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x52, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x1a, 0xb0, 0x01, 0x0a,
	0x0d, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x3d,
	0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x46, 0x0a, 0x06, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x42,
	0x80, 0x01, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31,
	0x42, 0x0a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x23,
	0x6b, 0x61, 0x66, 0x6b, 0x61, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x58, 0x58, 0xaa, 0x02, 0x08, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x08, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x56, 0x31, 0xe2,
	0x02, 0x14, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x09, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_v1_cloud_proto_rawDescOnce sync.Once
	file_cloud_v1_cloud_proto_rawDescData = file_cloud_v1_cloud_proto_rawDesc
)

func file_cloud_v1_cloud_proto_rawDescGZIP() []byte {
	file_cloud_v1_cloud_proto_rawDescOnce.Do(func() {
		file_cloud_v1_cloud_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_v1_cloud_proto_rawDescData)
	})
	return file_cloud_v1_cloud_proto_rawDescData
}

var file_cloud_v1_cloud_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_cloud_v1_cloud_proto_goTypes = []any{
	(*Profile)(nil),                     // 0: cloud.v1.Profile
	(*Repository)(nil),                  // 1: cloud.v1.Repository
	(*Commit)(nil),                      // 2: cloud.v1.Commit
	(*Commit_CommitDetails)(nil),        // 3: cloud.v1.Commit.CommitDetails
	(*Commit_CommitDetails_Author)(nil), // 4: cloud.v1.Commit.CommitDetails.Author
}
var file_cloud_v1_cloud_proto_depIdxs = []int32{
	3, // 0: cloud.v1.Commit.commit:type_name -> cloud.v1.Commit.CommitDetails
	4, // 1: cloud.v1.Commit.CommitDetails.author:type_name -> cloud.v1.Commit.CommitDetails.Author
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_cloud_v1_cloud_proto_init() }
func file_cloud_v1_cloud_proto_init() {
	if File_cloud_v1_cloud_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cloud_v1_cloud_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_v1_cloud_proto_goTypes,
		DependencyIndexes: file_cloud_v1_cloud_proto_depIdxs,
		MessageInfos:      file_cloud_v1_cloud_proto_msgTypes,
	}.Build()
	File_cloud_v1_cloud_proto = out.File
	file_cloud_v1_cloud_proto_rawDesc = nil
	file_cloud_v1_cloud_proto_goTypes = nil
	file_cloud_v1_cloud_proto_depIdxs = nil
}
