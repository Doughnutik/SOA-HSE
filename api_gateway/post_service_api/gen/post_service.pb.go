// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: post_service.proto

package gen

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Post struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	CreatorId     string                 `protobuf:"bytes,4,opt,name=creator_id,json=creatorId,proto3" json:"creator_id,omitempty"`
	CreatedAt     string                 `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     string                 `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	IsPrivate     bool                   `protobuf:"varint,7,opt,name=is_private,json=isPrivate,proto3" json:"is_private,omitempty"`
	Tags          []string               `protobuf:"bytes,8,rep,name=tags,proto3" json:"tags,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Post) Reset() {
	*x = Post{}
	mi := &file_post_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Post) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Post) ProtoMessage() {}

func (x *Post) ProtoReflect() protoreflect.Message {
	mi := &file_post_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Post.ProtoReflect.Descriptor instead.
func (*Post) Descriptor() ([]byte, []int) {
	return file_post_service_proto_rawDescGZIP(), []int{0}
}

func (x *Post) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Post) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Post) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Post) GetCreatorId() string {
	if x != nil {
		return x.CreatorId
	}
	return ""
}

func (x *Post) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Post) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Post) GetIsPrivate() bool {
	if x != nil {
		return x.IsPrivate
	}
	return false
}

func (x *Post) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

type PostCreateData struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Title         string                 `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description   string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	CreatorId     string                 `protobuf:"bytes,3,opt,name=creator_id,json=creatorId,proto3" json:"creator_id,omitempty"`
	IsPrivate     bool                   `protobuf:"varint,4,opt,name=is_private,json=isPrivate,proto3" json:"is_private,omitempty"`
	Tags          []string               `protobuf:"bytes,5,rep,name=tags,proto3" json:"tags,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PostCreateData) Reset() {
	*x = PostCreateData{}
	mi := &file_post_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PostCreateData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostCreateData) ProtoMessage() {}

func (x *PostCreateData) ProtoReflect() protoreflect.Message {
	mi := &file_post_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostCreateData.ProtoReflect.Descriptor instead.
func (*PostCreateData) Descriptor() ([]byte, []int) {
	return file_post_service_proto_rawDescGZIP(), []int{1}
}

func (x *PostCreateData) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *PostCreateData) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *PostCreateData) GetCreatorId() string {
	if x != nil {
		return x.CreatorId
	}
	return ""
}

func (x *PostCreateData) GetIsPrivate() bool {
	if x != nil {
		return x.IsPrivate
	}
	return false
}

func (x *PostCreateData) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

type PostCreateResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PostCreateResponse) Reset() {
	*x = PostCreateResponse{}
	mi := &file_post_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PostCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostCreateResponse) ProtoMessage() {}

func (x *PostCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_post_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostCreateResponse.ProtoReflect.Descriptor instead.
func (*PostCreateResponse) Descriptor() ([]byte, []int) {
	return file_post_service_proto_rawDescGZIP(), []int{2}
}

func (x *PostCreateResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type PostGetData struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatorId     string                 `protobuf:"bytes,2,opt,name=creator_id,json=creatorId,proto3" json:"creator_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PostGetData) Reset() {
	*x = PostGetData{}
	mi := &file_post_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PostGetData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostGetData) ProtoMessage() {}

func (x *PostGetData) ProtoReflect() protoreflect.Message {
	mi := &file_post_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostGetData.ProtoReflect.Descriptor instead.
func (*PostGetData) Descriptor() ([]byte, []int) {
	return file_post_service_proto_rawDescGZIP(), []int{3}
}

func (x *PostGetData) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PostGetData) GetCreatorId() string {
	if x != nil {
		return x.CreatorId
	}
	return ""
}

type PostGetResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Post          *Post                  `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PostGetResponse) Reset() {
	*x = PostGetResponse{}
	mi := &file_post_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PostGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostGetResponse) ProtoMessage() {}

func (x *PostGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_post_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostGetResponse.ProtoReflect.Descriptor instead.
func (*PostGetResponse) Descriptor() ([]byte, []int) {
	return file_post_service_proto_rawDescGZIP(), []int{4}
}

func (x *PostGetResponse) GetPost() *Post {
	if x != nil {
		return x.Post
	}
	return nil
}

type PostUpdateData struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	CreatorId     string                 `protobuf:"bytes,4,opt,name=creator_id,json=creatorId,proto3" json:"creator_id,omitempty"`
	IsPrivate     bool                   `protobuf:"varint,5,opt,name=is_private,json=isPrivate,proto3" json:"is_private,omitempty"`
	Tags          []string               `protobuf:"bytes,6,rep,name=tags,proto3" json:"tags,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PostUpdateData) Reset() {
	*x = PostUpdateData{}
	mi := &file_post_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PostUpdateData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostUpdateData) ProtoMessage() {}

func (x *PostUpdateData) ProtoReflect() protoreflect.Message {
	mi := &file_post_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostUpdateData.ProtoReflect.Descriptor instead.
func (*PostUpdateData) Descriptor() ([]byte, []int) {
	return file_post_service_proto_rawDescGZIP(), []int{5}
}

func (x *PostUpdateData) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PostUpdateData) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *PostUpdateData) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *PostUpdateData) GetCreatorId() string {
	if x != nil {
		return x.CreatorId
	}
	return ""
}

func (x *PostUpdateData) GetIsPrivate() bool {
	if x != nil {
		return x.IsPrivate
	}
	return false
}

func (x *PostUpdateData) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

type PostUpdateResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PostUpdateResponse) Reset() {
	*x = PostUpdateResponse{}
	mi := &file_post_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PostUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostUpdateResponse) ProtoMessage() {}

func (x *PostUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_post_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostUpdateResponse.ProtoReflect.Descriptor instead.
func (*PostUpdateResponse) Descriptor() ([]byte, []int) {
	return file_post_service_proto_rawDescGZIP(), []int{6}
}

func (x *PostUpdateResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type PostDeleteData struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatorId     string                 `protobuf:"bytes,2,opt,name=creator_id,json=creatorId,proto3" json:"creator_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PostDeleteData) Reset() {
	*x = PostDeleteData{}
	mi := &file_post_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PostDeleteData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostDeleteData) ProtoMessage() {}

func (x *PostDeleteData) ProtoReflect() protoreflect.Message {
	mi := &file_post_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostDeleteData.ProtoReflect.Descriptor instead.
func (*PostDeleteData) Descriptor() ([]byte, []int) {
	return file_post_service_proto_rawDescGZIP(), []int{7}
}

func (x *PostDeleteData) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PostDeleteData) GetCreatorId() string {
	if x != nil {
		return x.CreatorId
	}
	return ""
}

type PostDeleteResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PostDeleteResponse) Reset() {
	*x = PostDeleteResponse{}
	mi := &file_post_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PostDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostDeleteResponse) ProtoMessage() {}

func (x *PostDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_post_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostDeleteResponse.ProtoReflect.Descriptor instead.
func (*PostDeleteResponse) Descriptor() ([]byte, []int) {
	return file_post_service_proto_rawDescGZIP(), []int{8}
}

func (x *PostDeleteResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type ListPostsData struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Page          int32                  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Limit         int32                  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	CreatorId     string                 `protobuf:"bytes,3,opt,name=creator_id,json=creatorId,proto3" json:"creator_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListPostsData) Reset() {
	*x = ListPostsData{}
	mi := &file_post_service_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListPostsData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPostsData) ProtoMessage() {}

func (x *ListPostsData) ProtoReflect() protoreflect.Message {
	mi := &file_post_service_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPostsData.ProtoReflect.Descriptor instead.
func (*ListPostsData) Descriptor() ([]byte, []int) {
	return file_post_service_proto_rawDescGZIP(), []int{9}
}

func (x *ListPostsData) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListPostsData) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListPostsData) GetCreatorId() string {
	if x != nil {
		return x.CreatorId
	}
	return ""
}

type ListPostsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Posts         []*Post                `protobuf:"bytes,1,rep,name=posts,proto3" json:"posts,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListPostsResponse) Reset() {
	*x = ListPostsResponse{}
	mi := &file_post_service_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListPostsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPostsResponse) ProtoMessage() {}

func (x *ListPostsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_post_service_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPostsResponse.ProtoReflect.Descriptor instead.
func (*ListPostsResponse) Descriptor() ([]byte, []int) {
	return file_post_service_proto_rawDescGZIP(), []int{10}
}

func (x *ListPostsResponse) GetPosts() []*Post {
	if x != nil {
		return x.Posts
	}
	return nil
}

var File_post_service_proto protoreflect.FileDescriptor

const file_post_service_proto_rawDesc = "" +
	"\n" +
	"\x12post_service.proto\x12\vpostservice\"\xde\x01\n" +
	"\x04Post\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x14\n" +
	"\x05title\x18\x02 \x01(\tR\x05title\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\x12\x1d\n" +
	"\n" +
	"creator_id\x18\x04 \x01(\tR\tcreatorId\x12\x1d\n" +
	"\n" +
	"created_at\x18\x05 \x01(\tR\tcreatedAt\x12\x1d\n" +
	"\n" +
	"updated_at\x18\x06 \x01(\tR\tupdatedAt\x12\x1d\n" +
	"\n" +
	"is_private\x18\a \x01(\bR\tisPrivate\x12\x12\n" +
	"\x04tags\x18\b \x03(\tR\x04tags\"\x9a\x01\n" +
	"\x0ePostCreateData\x12\x14\n" +
	"\x05title\x18\x01 \x01(\tR\x05title\x12 \n" +
	"\vdescription\x18\x02 \x01(\tR\vdescription\x12\x1d\n" +
	"\n" +
	"creator_id\x18\x03 \x01(\tR\tcreatorId\x12\x1d\n" +
	"\n" +
	"is_private\x18\x04 \x01(\bR\tisPrivate\x12\x12\n" +
	"\x04tags\x18\x05 \x03(\tR\x04tags\"$\n" +
	"\x12PostCreateResponse\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"<\n" +
	"\vPostGetData\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x1d\n" +
	"\n" +
	"creator_id\x18\x02 \x01(\tR\tcreatorId\"8\n" +
	"\x0fPostGetResponse\x12%\n" +
	"\x04post\x18\x01 \x01(\v2\x11.postservice.PostR\x04post\"\xaa\x01\n" +
	"\x0ePostUpdateData\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x14\n" +
	"\x05title\x18\x02 \x01(\tR\x05title\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\x12\x1d\n" +
	"\n" +
	"creator_id\x18\x04 \x01(\tR\tcreatorId\x12\x1d\n" +
	"\n" +
	"is_private\x18\x05 \x01(\bR\tisPrivate\x12\x12\n" +
	"\x04tags\x18\x06 \x03(\tR\x04tags\".\n" +
	"\x12PostUpdateResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\"?\n" +
	"\x0ePostDeleteData\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x1d\n" +
	"\n" +
	"creator_id\x18\x02 \x01(\tR\tcreatorId\".\n" +
	"\x12PostDeleteResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\"X\n" +
	"\rListPostsData\x12\x12\n" +
	"\x04page\x18\x01 \x01(\x05R\x04page\x12\x14\n" +
	"\x05limit\x18\x02 \x01(\x05R\x05limit\x12\x1d\n" +
	"\n" +
	"creator_id\x18\x03 \x01(\tR\tcreatorId\"<\n" +
	"\x11ListPostsResponse\x12'\n" +
	"\x05posts\x18\x01 \x03(\v2\x11.postservice.PostR\x05posts2\xfd\x02\n" +
	"\vPostService\x12J\n" +
	"\n" +
	"CreatePost\x12\x1b.postservice.PostCreateData\x1a\x1f.postservice.PostCreateResponse\x12A\n" +
	"\aGetPost\x12\x18.postservice.PostGetData\x1a\x1c.postservice.PostGetResponse\x12J\n" +
	"\n" +
	"UpdatePost\x12\x1b.postservice.PostUpdateData\x1a\x1f.postservice.PostUpdateResponse\x12J\n" +
	"\n" +
	"DeletePost\x12\x1b.postservice.PostDeleteData\x1a\x1f.postservice.PostDeleteResponse\x12G\n" +
	"\tListPosts\x12\x1a.postservice.ListPostsData\x1a\x1e.postservice.ListPostsResponseB\aZ\x05./genb\x06proto3"

var (
	file_post_service_proto_rawDescOnce sync.Once
	file_post_service_proto_rawDescData []byte
)

func file_post_service_proto_rawDescGZIP() []byte {
	file_post_service_proto_rawDescOnce.Do(func() {
		file_post_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_post_service_proto_rawDesc), len(file_post_service_proto_rawDesc)))
	})
	return file_post_service_proto_rawDescData
}

var file_post_service_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_post_service_proto_goTypes = []any{
	(*Post)(nil),               // 0: postservice.Post
	(*PostCreateData)(nil),     // 1: postservice.PostCreateData
	(*PostCreateResponse)(nil), // 2: postservice.PostCreateResponse
	(*PostGetData)(nil),        // 3: postservice.PostGetData
	(*PostGetResponse)(nil),    // 4: postservice.PostGetResponse
	(*PostUpdateData)(nil),     // 5: postservice.PostUpdateData
	(*PostUpdateResponse)(nil), // 6: postservice.PostUpdateResponse
	(*PostDeleteData)(nil),     // 7: postservice.PostDeleteData
	(*PostDeleteResponse)(nil), // 8: postservice.PostDeleteResponse
	(*ListPostsData)(nil),      // 9: postservice.ListPostsData
	(*ListPostsResponse)(nil),  // 10: postservice.ListPostsResponse
}
var file_post_service_proto_depIdxs = []int32{
	0,  // 0: postservice.PostGetResponse.post:type_name -> postservice.Post
	0,  // 1: postservice.ListPostsResponse.posts:type_name -> postservice.Post
	1,  // 2: postservice.PostService.CreatePost:input_type -> postservice.PostCreateData
	3,  // 3: postservice.PostService.GetPost:input_type -> postservice.PostGetData
	5,  // 4: postservice.PostService.UpdatePost:input_type -> postservice.PostUpdateData
	7,  // 5: postservice.PostService.DeletePost:input_type -> postservice.PostDeleteData
	9,  // 6: postservice.PostService.ListPosts:input_type -> postservice.ListPostsData
	2,  // 7: postservice.PostService.CreatePost:output_type -> postservice.PostCreateResponse
	4,  // 8: postservice.PostService.GetPost:output_type -> postservice.PostGetResponse
	6,  // 9: postservice.PostService.UpdatePost:output_type -> postservice.PostUpdateResponse
	8,  // 10: postservice.PostService.DeletePost:output_type -> postservice.PostDeleteResponse
	10, // 11: postservice.PostService.ListPosts:output_type -> postservice.ListPostsResponse
	7,  // [7:12] is the sub-list for method output_type
	2,  // [2:7] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_post_service_proto_init() }
func file_post_service_proto_init() {
	if File_post_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_post_service_proto_rawDesc), len(file_post_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_post_service_proto_goTypes,
		DependencyIndexes: file_post_service_proto_depIdxs,
		MessageInfos:      file_post_service_proto_msgTypes,
	}.Build()
	File_post_service_proto = out.File
	file_post_service_proto_goTypes = nil
	file_post_service_proto_depIdxs = nil
}
