// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: protos/v1/post/post.proto

package post

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

type PostMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostId string   `protobuf:"bytes,1,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	Author string   `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`
	Title  string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Body   string   `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	Tags   []string `protobuf:"bytes,5,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *PostMessage) Reset() {
	*x = PostMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_v1_post_post_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostMessage) ProtoMessage() {}

func (x *PostMessage) ProtoReflect() protoreflect.Message {
	mi := &file_protos_v1_post_post_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostMessage.ProtoReflect.Descriptor instead.
func (*PostMessage) Descriptor() ([]byte, []int) {
	return file_protos_v1_post_post_proto_rawDescGZIP(), []int{0}
}

func (x *PostMessage) GetPostId() string {
	if x != nil {
		return x.PostId
	}
	return ""
}

func (x *PostMessage) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *PostMessage) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *PostMessage) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *PostMessage) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

type ListPostsByUserIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *ListPostsByUserIdRequest) Reset() {
	*x = ListPostsByUserIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_v1_post_post_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPostsByUserIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPostsByUserIdRequest) ProtoMessage() {}

func (x *ListPostsByUserIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_v1_post_post_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPostsByUserIdRequest.ProtoReflect.Descriptor instead.
func (*ListPostsByUserIdRequest) Descriptor() ([]byte, []int) {
	return file_protos_v1_post_post_proto_rawDescGZIP(), []int{1}
}

func (x *ListPostsByUserIdRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type ListPostsByUserIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostMessages []*PostMessage `protobuf:"bytes,1,rep,name=post_messages,json=postMessages,proto3" json:"post_messages,omitempty"`
}

func (x *ListPostsByUserIdResponse) Reset() {
	*x = ListPostsByUserIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_v1_post_post_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPostsByUserIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPostsByUserIdResponse) ProtoMessage() {}

func (x *ListPostsByUserIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_v1_post_post_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPostsByUserIdResponse.ProtoReflect.Descriptor instead.
func (*ListPostsByUserIdResponse) Descriptor() ([]byte, []int) {
	return file_protos_v1_post_post_proto_rawDescGZIP(), []int{2}
}

func (x *ListPostsByUserIdResponse) GetPostMessages() []*PostMessage {
	if x != nil {
		return x.PostMessages
	}
	return nil
}

type ListPostsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListPostsRequest) Reset() {
	*x = ListPostsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_v1_post_post_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPostsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPostsRequest) ProtoMessage() {}

func (x *ListPostsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_v1_post_post_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPostsRequest.ProtoReflect.Descriptor instead.
func (*ListPostsRequest) Descriptor() ([]byte, []int) {
	return file_protos_v1_post_post_proto_rawDescGZIP(), []int{3}
}

type ListPostsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostMessages []*PostMessage `protobuf:"bytes,1,rep,name=post_messages,json=postMessages,proto3" json:"post_messages,omitempty"`
}

func (x *ListPostsResponse) Reset() {
	*x = ListPostsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_v1_post_post_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPostsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPostsResponse) ProtoMessage() {}

func (x *ListPostsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_v1_post_post_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_protos_v1_post_post_proto_rawDescGZIP(), []int{4}
}

func (x *ListPostsResponse) GetPostMessages() []*PostMessage {
	if x != nil {
		return x.PostMessages
	}
	return nil
}

var File_protos_v1_post_post_proto protoreflect.FileDescriptor

var file_protos_v1_post_post_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x73, 0x74,
	0x2f, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x76, 0x31, 0x2e,
	0x70, 0x6f, 0x73, 0x74, 0x22, 0x7c, 0x0a, 0x0b, 0x50, 0x6f, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f,
	0x64, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61,
	0x67, 0x73, 0x22, 0x33, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x42,
	0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x56, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x6f, 0x73, 0x74, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x0d, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x76, 0x31,
	0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x0c, 0x70, 0x6f, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x22,
	0x12, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x4e, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x0d, 0x70, 0x6f, 0x73, 0x74,
	0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0c, 0x70, 0x6f, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x32, 0xa6, 0x01, 0x0a, 0x04, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x5a, 0x0a, 0x11,
	0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x21, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x50, 0x6f, 0x73, 0x74, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74,
	0x50, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x19, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1a, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2f, 0x5a, 0x2d,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x42, 0x61, 0x63, 0x6b, 0x65,
	0x6e, 0x64, 0x2d, 0x47, 0x72, 0x70, 0x63, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_v1_post_post_proto_rawDescOnce sync.Once
	file_protos_v1_post_post_proto_rawDescData = file_protos_v1_post_post_proto_rawDesc
)

func file_protos_v1_post_post_proto_rawDescGZIP() []byte {
	file_protos_v1_post_post_proto_rawDescOnce.Do(func() {
		file_protos_v1_post_post_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_v1_post_post_proto_rawDescData)
	})
	return file_protos_v1_post_post_proto_rawDescData
}

var file_protos_v1_post_post_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_protos_v1_post_post_proto_goTypes = []interface{}{
	(*PostMessage)(nil),               // 0: v1.post.PostMessage
	(*ListPostsByUserIdRequest)(nil),  // 1: v1.post.ListPostsByUserIdRequest
	(*ListPostsByUserIdResponse)(nil), // 2: v1.post.ListPostsByUserIdResponse
	(*ListPostsRequest)(nil),          // 3: v1.post.ListPostsRequest
	(*ListPostsResponse)(nil),         // 4: v1.post.ListPostsResponse
}
var file_protos_v1_post_post_proto_depIdxs = []int32{
	0, // 0: v1.post.ListPostsByUserIdResponse.post_messages:type_name -> v1.post.PostMessage
	0, // 1: v1.post.ListPostsResponse.post_messages:type_name -> v1.post.PostMessage
	1, // 2: v1.post.Post.ListPostsByUserId:input_type -> v1.post.ListPostsByUserIdRequest
	3, // 3: v1.post.Post.ListPosts:input_type -> v1.post.ListPostsRequest
	2, // 4: v1.post.Post.ListPostsByUserId:output_type -> v1.post.ListPostsByUserIdResponse
	4, // 5: v1.post.Post.ListPosts:output_type -> v1.post.ListPostsResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protos_v1_post_post_proto_init() }
func file_protos_v1_post_post_proto_init() {
	if File_protos_v1_post_post_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_v1_post_post_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostMessage); i {
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
		file_protos_v1_post_post_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPostsByUserIdRequest); i {
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
		file_protos_v1_post_post_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPostsByUserIdResponse); i {
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
		file_protos_v1_post_post_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPostsRequest); i {
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
		file_protos_v1_post_post_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPostsResponse); i {
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
			RawDescriptor: file_protos_v1_post_post_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_v1_post_post_proto_goTypes,
		DependencyIndexes: file_protos_v1_post_post_proto_depIdxs,
		MessageInfos:      file_protos_v1_post_post_proto_msgTypes,
	}.Build()
	File_protos_v1_post_post_proto = out.File
	file_protos_v1_post_post_proto_rawDesc = nil
	file_protos_v1_post_post_proto_goTypes = nil
	file_protos_v1_post_post_proto_depIdxs = nil
}