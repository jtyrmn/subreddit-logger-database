//
//this file is the API contract that all services using
//subreddit-logger-database will share (they will all use
//this same file for code gen)

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.4
// source: pb/proto/ListingsDatabase.proto

package pb

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

// A listing object that's stored in + returned from the database.
type RedditContent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string                        `protobuf:"bytes,1,opt,name=id,json=_id,proto3" json:"id,omitempty"`
	MetaData *RedditContent_MetaData       `protobuf:"bytes,2,opt,name=meta_data,json=listing,proto3" json:"meta_data,omitempty"`
	Entries  []*RedditContent_ListingEntry `protobuf:"bytes,3,rep,name=entries,proto3" json:"entries,omitempty"`
}

func (x *RedditContent) Reset() {
	*x = RedditContent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_proto_ListingsDatabase_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RedditContent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RedditContent) ProtoMessage() {}

func (x *RedditContent) ProtoReflect() protoreflect.Message {
	mi := &file_pb_proto_ListingsDatabase_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RedditContent.ProtoReflect.Descriptor instead.
func (*RedditContent) Descriptor() ([]byte, []int) {
	return file_pb_proto_ListingsDatabase_proto_rawDescGZIP(), []int{0}
}

func (x *RedditContent) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RedditContent) GetMetaData() *RedditContent_MetaData {
	if x != nil {
		return x.MetaData
	}
	return nil
}

func (x *RedditContent) GetEntries() []*RedditContent_ListingEntry {
	if x != nil {
		return x.Entries
	}
	return nil
}

type SaveListingsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SaveListingsResponse) Reset() {
	*x = SaveListingsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_proto_ListingsDatabase_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveListingsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveListingsResponse) ProtoMessage() {}

func (x *SaveListingsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_proto_ListingsDatabase_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveListingsResponse.ProtoReflect.Descriptor instead.
func (*SaveListingsResponse) Descriptor() ([]byte, []int) {
	return file_pb_proto_ListingsDatabase_proto_rawDescGZIP(), []int{1}
}

type UpdateListingsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateListingsResponse) Reset() {
	*x = UpdateListingsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_proto_ListingsDatabase_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateListingsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateListingsResponse) ProtoMessage() {}

func (x *UpdateListingsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_proto_ListingsDatabase_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateListingsResponse.ProtoReflect.Descriptor instead.
func (*UpdateListingsResponse) Descriptor() ([]byte, []int) {
	return file_pb_proto_ListingsDatabase_proto_rawDescGZIP(), []int{2}
}

type CullListingsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MaxAge uint64 `protobuf:"varint,1,opt,name=max_age,json=maxAge,proto3" json:"max_age,omitempty"` // max age is in seconds
}

func (x *CullListingsRequest) Reset() {
	*x = CullListingsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_proto_ListingsDatabase_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CullListingsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CullListingsRequest) ProtoMessage() {}

func (x *CullListingsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_proto_ListingsDatabase_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CullListingsRequest.ProtoReflect.Descriptor instead.
func (*CullListingsRequest) Descriptor() ([]byte, []int) {
	return file_pb_proto_ListingsDatabase_proto_rawDescGZIP(), []int{3}
}

func (x *CullListingsRequest) GetMaxAge() uint64 {
	if x != nil {
		return x.MaxAge
	}
	return 0
}

type CullListingsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NumDeleted uint32 `protobuf:"varint,1,opt,name=num_deleted,json=numDeleted,proto3" json:"num_deleted,omitempty"`
}

func (x *CullListingsResponse) Reset() {
	*x = CullListingsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_proto_ListingsDatabase_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CullListingsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CullListingsResponse) ProtoMessage() {}

func (x *CullListingsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_proto_ListingsDatabase_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CullListingsResponse.ProtoReflect.Descriptor instead.
func (*CullListingsResponse) Descriptor() ([]byte, []int) {
	return file_pb_proto_ListingsDatabase_proto_rawDescGZIP(), []int{4}
}

func (x *CullListingsResponse) GetNumDeleted() uint32 {
	if x != nil {
		return x.NumDeleted
	}
	return 0
}

type ManyListingsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit uint32 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"` // this value shouldn't be too high
	Skip  uint32 `protobuf:"varint,2,opt,name=skip,proto3" json:"skip,omitempty"`   // for pagination. # of listings to skip over
}

func (x *ManyListingsRequest) Reset() {
	*x = ManyListingsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_proto_ListingsDatabase_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManyListingsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManyListingsRequest) ProtoMessage() {}

func (x *ManyListingsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_proto_ListingsDatabase_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManyListingsRequest.ProtoReflect.Descriptor instead.
func (*ManyListingsRequest) Descriptor() ([]byte, []int) {
	return file_pb_proto_ListingsDatabase_proto_rawDescGZIP(), []int{5}
}

func (x *ManyListingsRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ManyListingsRequest) GetSkip() uint32 {
	if x != nil {
		return x.Skip
	}
	return 0
}

type ManyListingsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Listings []*RedditContent `protobuf:"bytes,1,rep,name=listings,proto3" json:"listings,omitempty"`
}

func (x *ManyListingsResponse) Reset() {
	*x = ManyListingsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_proto_ListingsDatabase_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManyListingsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManyListingsResponse) ProtoMessage() {}

func (x *ManyListingsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_proto_ListingsDatabase_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManyListingsResponse.ProtoReflect.Descriptor instead.
func (*ManyListingsResponse) Descriptor() ([]byte, []int) {
	return file_pb_proto_ListingsDatabase_proto_rawDescGZIP(), []int{6}
}

func (x *ManyListingsResponse) GetListings() []*RedditContent {
	if x != nil {
		return x.Listings
	}
	return nil
}

type FetchListingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *FetchListingRequest) Reset() {
	*x = FetchListingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_proto_ListingsDatabase_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchListingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchListingRequest) ProtoMessage() {}

func (x *FetchListingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_proto_ListingsDatabase_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchListingRequest.ProtoReflect.Descriptor instead.
func (*FetchListingRequest) Descriptor() ([]byte, []int) {
	return file_pb_proto_ListingsDatabase_proto_rawDescGZIP(), []int{7}
}

func (x *FetchListingRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type RetrieveListingsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MaxAge uint64 `protobuf:"varint,1,opt,name=max_age,json=maxAge,proto3" json:"max_age,omitempty"`
}

func (x *RetrieveListingsRequest) Reset() {
	*x = RetrieveListingsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_proto_ListingsDatabase_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetrieveListingsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetrieveListingsRequest) ProtoMessage() {}

func (x *RetrieveListingsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_proto_ListingsDatabase_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetrieveListingsRequest.ProtoReflect.Descriptor instead.
func (*RetrieveListingsRequest) Descriptor() ([]byte, []int) {
	return file_pb_proto_ListingsDatabase_proto_rawDescGZIP(), []int{8}
}

func (x *RetrieveListingsRequest) GetMaxAge() uint64 {
	if x != nil {
		return x.MaxAge
	}
	return 0
}

type RedditContent_MetaData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContentType string `protobuf:"bytes,1,opt,name=content_type,json=contenttype,proto3" json:"content_type,omitempty"`
	Id          string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Title       string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Upvotes     uint32 `protobuf:"varint,4,opt,name=upvotes,proto3" json:"upvotes,omitempty"`
	Comments    uint32 `protobuf:"varint,5,opt,name=comments,proto3" json:"comments,omitempty"`
	DateCreated uint64 `protobuf:"varint,6,opt,name=date_created,json=date,proto3" json:"date_created,omitempty"`
	DateQueried uint64 `protobuf:"varint,7,opt,name=date_queried,json=querydate,proto3" json:"date_queried,omitempty"`
}

func (x *RedditContent_MetaData) Reset() {
	*x = RedditContent_MetaData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_proto_ListingsDatabase_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RedditContent_MetaData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RedditContent_MetaData) ProtoMessage() {}

func (x *RedditContent_MetaData) ProtoReflect() protoreflect.Message {
	mi := &file_pb_proto_ListingsDatabase_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RedditContent_MetaData.ProtoReflect.Descriptor instead.
func (*RedditContent_MetaData) Descriptor() ([]byte, []int) {
	return file_pb_proto_ListingsDatabase_proto_rawDescGZIP(), []int{0, 0}
}

func (x *RedditContent_MetaData) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

func (x *RedditContent_MetaData) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RedditContent_MetaData) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *RedditContent_MetaData) GetUpvotes() uint32 {
	if x != nil {
		return x.Upvotes
	}
	return 0
}

func (x *RedditContent_MetaData) GetComments() uint32 {
	if x != nil {
		return x.Comments
	}
	return 0
}

func (x *RedditContent_MetaData) GetDateCreated() uint64 {
	if x != nil {
		return x.DateCreated
	}
	return 0
}

func (x *RedditContent_MetaData) GetDateQueried() uint64 {
	if x != nil {
		return x.DateQueried
	}
	return 0
}

type RedditContent_ListingEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Upvotes     uint32 `protobuf:"varint,1,opt,name=upvotes,proto3" json:"upvotes,omitempty"`
	Comments    uint32 `protobuf:"varint,2,opt,name=comments,proto3" json:"comments,omitempty"`
	DateQueried uint64 `protobuf:"varint,3,opt,name=date_queried,json=date,proto3" json:"date_queried,omitempty"`
}

func (x *RedditContent_ListingEntry) Reset() {
	*x = RedditContent_ListingEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_proto_ListingsDatabase_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RedditContent_ListingEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RedditContent_ListingEntry) ProtoMessage() {}

func (x *RedditContent_ListingEntry) ProtoReflect() protoreflect.Message {
	mi := &file_pb_proto_ListingsDatabase_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RedditContent_ListingEntry.ProtoReflect.Descriptor instead.
func (*RedditContent_ListingEntry) Descriptor() ([]byte, []int) {
	return file_pb_proto_ListingsDatabase_proto_rawDescGZIP(), []int{0, 1}
}

func (x *RedditContent_ListingEntry) GetUpvotes() uint32 {
	if x != nil {
		return x.Upvotes
	}
	return 0
}

func (x *RedditContent_ListingEntry) GetComments() uint32 {
	if x != nil {
		return x.Comments
	}
	return 0
}

func (x *RedditContent_ListingEntry) GetDateQueried() uint64 {
	if x != nil {
		return x.DateQueried
	}
	return 0
}

var File_pb_proto_ListingsDatabase_proto protoreflect.FileDescriptor

var file_pb_proto_ListingsDatabase_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x4c, 0x69, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xb7, 0x03, 0x0a, 0x0d, 0x52, 0x65, 0x64, 0x64, 0x69, 0x74, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x12, 0x0f, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x5f, 0x69, 0x64, 0x12, 0x33, 0x0a, 0x09, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x52, 0x65, 0x64, 0x64, 0x69, 0x74,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x07, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x35, 0x0a, 0x07, 0x65, 0x6e, 0x74,
	0x72, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x52, 0x65, 0x64,
	0x64, 0x69, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73,
	0x1a, 0xc6, 0x01, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x12, 0x21, 0x0a,
	0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x70, 0x76, 0x6f, 0x74, 0x65,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x75, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x73,
	0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x1a, 0x0a, 0x0c,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0c, 0x64, 0x61, 0x74, 0x65,
	0x5f, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x64, 0x61, 0x74, 0x65, 0x1a, 0x60, 0x0a, 0x0c, 0x4c, 0x69, 0x73,
	0x74, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x70, 0x76,
	0x6f, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x75, 0x70, 0x76, 0x6f,
	0x74, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12,
	0x1a, 0x0a, 0x0c, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x22, 0x16, 0x0a, 0x14, 0x53,
	0x61, 0x76, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x18, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2e, 0x0a,
	0x13, 0x43, 0x75, 0x6c, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6d, 0x61, 0x78, 0x5f, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6d, 0x61, 0x78, 0x41, 0x67, 0x65, 0x22, 0x37, 0x0a,
	0x14, 0x43, 0x75, 0x6c, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x75, 0x6d, 0x5f, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x6e, 0x75, 0x6d, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x22, 0x3f, 0x0a, 0x13, 0x4d, 0x61, 0x6e, 0x79, 0x4c, 0x69,
	0x73, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6b, 0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x04, 0x73, 0x6b, 0x69, 0x70, 0x22, 0x42, 0x0a, 0x14, 0x4d, 0x61, 0x6e, 0x79, 0x4c,
	0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2a, 0x0a, 0x08, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x52, 0x65, 0x64, 0x64, 0x69, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x52, 0x08, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x22, 0x25, 0x0a, 0x13, 0x46,
	0x65, 0x74, 0x63, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x32, 0x0a, 0x17, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x6d, 0x61, 0x78, 0x5f, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06,
	0x6d, 0x61, 0x78, 0x41, 0x67, 0x65, 0x32, 0x84, 0x03, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x0c, 0x53,
	0x61, 0x76, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x0e, 0x2e, 0x52, 0x65,
	0x64, 0x64, 0x69, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x1a, 0x15, 0x2e, 0x53, 0x61,
	0x76, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x3d, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x0e, 0x2e, 0x52, 0x65, 0x64, 0x64, 0x69,
	0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x1a, 0x17, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x3d, 0x0a, 0x0c, 0x43, 0x75, 0x6c, 0x6c, 0x4c, 0x69, 0x73,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x14, 0x2e, 0x43, 0x75, 0x6c, 0x6c, 0x4c, 0x69, 0x73, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x43, 0x75,
	0x6c, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0c, 0x4d, 0x61, 0x6e, 0x79, 0x4c, 0x69, 0x73, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x12, 0x14, 0x2e, 0x4d, 0x61, 0x6e, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x4d, 0x61, 0x6e,
	0x79, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x10, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x18, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65,
	0x76, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0e, 0x2e, 0x52, 0x65, 0x64, 0x64, 0x69, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x22, 0x00, 0x30, 0x01, 0x12, 0x36, 0x0a, 0x0c, 0x46, 0x65, 0x74, 0x63, 0x68, 0x4c, 0x69,
	0x73, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x14, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x4c, 0x69, 0x73,
	0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x52, 0x65,
	0x64, 0x64, 0x69, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x42, 0x07, 0x5a,
	0x05, 0x2e, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_proto_ListingsDatabase_proto_rawDescOnce sync.Once
	file_pb_proto_ListingsDatabase_proto_rawDescData = file_pb_proto_ListingsDatabase_proto_rawDesc
)

func file_pb_proto_ListingsDatabase_proto_rawDescGZIP() []byte {
	file_pb_proto_ListingsDatabase_proto_rawDescOnce.Do(func() {
		file_pb_proto_ListingsDatabase_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_proto_ListingsDatabase_proto_rawDescData)
	})
	return file_pb_proto_ListingsDatabase_proto_rawDescData
}

var file_pb_proto_ListingsDatabase_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_pb_proto_ListingsDatabase_proto_goTypes = []interface{}{
	(*RedditContent)(nil),              // 0: RedditContent
	(*SaveListingsResponse)(nil),       // 1: SaveListingsResponse
	(*UpdateListingsResponse)(nil),     // 2: UpdateListingsResponse
	(*CullListingsRequest)(nil),        // 3: CullListingsRequest
	(*CullListingsResponse)(nil),       // 4: CullListingsResponse
	(*ManyListingsRequest)(nil),        // 5: ManyListingsRequest
	(*ManyListingsResponse)(nil),       // 6: ManyListingsResponse
	(*FetchListingRequest)(nil),        // 7: FetchListingRequest
	(*RetrieveListingsRequest)(nil),    // 8: RetrieveListingsRequest
	(*RedditContent_MetaData)(nil),     // 9: RedditContent.MetaData
	(*RedditContent_ListingEntry)(nil), // 10: RedditContent.ListingEntry
}
var file_pb_proto_ListingsDatabase_proto_depIdxs = []int32{
	9,  // 0: RedditContent.meta_data:type_name -> RedditContent.MetaData
	10, // 1: RedditContent.entries:type_name -> RedditContent.ListingEntry
	0,  // 2: ManyListingsResponse.listings:type_name -> RedditContent
	0,  // 3: ListingsDatabase.SaveListings:input_type -> RedditContent
	0,  // 4: ListingsDatabase.UpdateListings:input_type -> RedditContent
	3,  // 5: ListingsDatabase.CullListings:input_type -> CullListingsRequest
	5,  // 6: ListingsDatabase.ManyListings:input_type -> ManyListingsRequest
	8,  // 7: ListingsDatabase.RetrieveListings:input_type -> RetrieveListingsRequest
	7,  // 8: ListingsDatabase.FetchListing:input_type -> FetchListingRequest
	1,  // 9: ListingsDatabase.SaveListings:output_type -> SaveListingsResponse
	2,  // 10: ListingsDatabase.UpdateListings:output_type -> UpdateListingsResponse
	4,  // 11: ListingsDatabase.CullListings:output_type -> CullListingsResponse
	6,  // 12: ListingsDatabase.ManyListings:output_type -> ManyListingsResponse
	0,  // 13: ListingsDatabase.RetrieveListings:output_type -> RedditContent
	0,  // 14: ListingsDatabase.FetchListing:output_type -> RedditContent
	9,  // [9:15] is the sub-list for method output_type
	3,  // [3:9] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_pb_proto_ListingsDatabase_proto_init() }
func file_pb_proto_ListingsDatabase_proto_init() {
	if File_pb_proto_ListingsDatabase_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_proto_ListingsDatabase_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RedditContent); i {
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
		file_pb_proto_ListingsDatabase_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveListingsResponse); i {
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
		file_pb_proto_ListingsDatabase_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateListingsResponse); i {
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
		file_pb_proto_ListingsDatabase_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CullListingsRequest); i {
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
		file_pb_proto_ListingsDatabase_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CullListingsResponse); i {
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
		file_pb_proto_ListingsDatabase_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ManyListingsRequest); i {
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
		file_pb_proto_ListingsDatabase_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ManyListingsResponse); i {
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
		file_pb_proto_ListingsDatabase_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchListingRequest); i {
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
		file_pb_proto_ListingsDatabase_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetrieveListingsRequest); i {
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
		file_pb_proto_ListingsDatabase_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RedditContent_MetaData); i {
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
		file_pb_proto_ListingsDatabase_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RedditContent_ListingEntry); i {
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
			RawDescriptor: file_pb_proto_ListingsDatabase_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_proto_ListingsDatabase_proto_goTypes,
		DependencyIndexes: file_pb_proto_ListingsDatabase_proto_depIdxs,
		MessageInfos:      file_pb_proto_ListingsDatabase_proto_msgTypes,
	}.Build()
	File_pb_proto_ListingsDatabase_proto = out.File
	file_pb_proto_ListingsDatabase_proto_rawDesc = nil
	file_pb_proto_ListingsDatabase_proto_goTypes = nil
	file_pb_proto_ListingsDatabase_proto_depIdxs = nil
}
