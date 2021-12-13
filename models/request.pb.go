// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1-devel
// 	protoc        v3.19.1
// source: request.proto

package models

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

type RequestParameter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApiKey string `protobuf:"bytes,1,opt,name=ApiKey,proto3" json:"ApiKey,omitempty"`
	ImdbID string `protobuf:"bytes,2,opt,name=ImdbID,proto3" json:"ImdbID,omitempty"`
	Title  string `protobuf:"bytes,3,opt,name=Title,proto3" json:"Title,omitempty"`
	Page   string `protobuf:"bytes,4,opt,name=Page,proto3" json:"Page,omitempty"`
}

func (x *RequestParameter) Reset() {
	*x = RequestParameter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestParameter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestParameter) ProtoMessage() {}

func (x *RequestParameter) ProtoReflect() protoreflect.Message {
	mi := &file_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestParameter.ProtoReflect.Descriptor instead.
func (*RequestParameter) Descriptor() ([]byte, []int) {
	return file_request_proto_rawDescGZIP(), []int{0}
}

func (x *RequestParameter) GetApiKey() string {
	if x != nil {
		return x.ApiKey
	}
	return ""
}

func (x *RequestParameter) GetImdbID() string {
	if x != nil {
		return x.ImdbID
	}
	return ""
}

func (x *RequestParameter) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *RequestParameter) GetPage() string {
	if x != nil {
		return x.Page
	}
	return ""
}

type ListFilm struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Search      []*DetailFilm `protobuf:"bytes,1,rep,name=Search,proto3" json:"Search,omitempty"`
	TotalResult int32         `protobuf:"varint,2,opt,name=TotalResult,proto3" json:"TotalResult,omitempty"`
	Response    string        `protobuf:"bytes,3,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *ListFilm) Reset() {
	*x = ListFilm{}
	if protoimpl.UnsafeEnabled {
		mi := &file_request_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFilm) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFilm) ProtoMessage() {}

func (x *ListFilm) ProtoReflect() protoreflect.Message {
	mi := &file_request_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFilm.ProtoReflect.Descriptor instead.
func (*ListFilm) Descriptor() ([]byte, []int) {
	return file_request_proto_rawDescGZIP(), []int{1}
}

func (x *ListFilm) GetSearch() []*DetailFilm {
	if x != nil {
		return x.Search
	}
	return nil
}

func (x *ListFilm) GetTotalResult() int32 {
	if x != nil {
		return x.TotalResult
	}
	return 0
}

func (x *ListFilm) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

type Film struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title  string `protobuf:"bytes,1,opt,name=Title,proto3" json:"Title,omitempty"`
	Year   string `protobuf:"bytes,2,opt,name=Year,proto3" json:"Year,omitempty"`
	ImdbID string `protobuf:"bytes,3,opt,name=imdbID,proto3" json:"imdbID,omitempty"`
	Poster string `protobuf:"bytes,4,opt,name=Poster,proto3" json:"Poster,omitempty"`
}

func (x *Film) Reset() {
	*x = Film{}
	if protoimpl.UnsafeEnabled {
		mi := &file_request_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Film) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Film) ProtoMessage() {}

func (x *Film) ProtoReflect() protoreflect.Message {
	mi := &file_request_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Film.ProtoReflect.Descriptor instead.
func (*Film) Descriptor() ([]byte, []int) {
	return file_request_proto_rawDescGZIP(), []int{2}
}

func (x *Film) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Film) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

func (x *Film) GetImdbID() string {
	if x != nil {
		return x.ImdbID
	}
	return ""
}

func (x *Film) GetPoster() string {
	if x != nil {
		return x.Poster
	}
	return ""
}

type DetailFilm struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title      string `protobuf:"bytes,1,opt,name=Title,proto3" json:"Title,omitempty"`
	Year       string `protobuf:"bytes,2,opt,name=Year,proto3" json:"Year,omitempty"`
	Rated      string `protobuf:"bytes,3,opt,name=Rated,proto3" json:"Rated,omitempty"`
	Release    string `protobuf:"bytes,4,opt,name=Release,proto3" json:"Release,omitempty"`
	Runtime    string `protobuf:"bytes,5,opt,name=Runtime,proto3" json:"Runtime,omitempty"`
	Genre      string `protobuf:"bytes,6,opt,name=Genre,proto3" json:"Genre,omitempty"`
	Director   string `protobuf:"bytes,7,opt,name=Director,proto3" json:"Director,omitempty"`
	Wrter      string `protobuf:"bytes,8,opt,name=Wrter,proto3" json:"Wrter,omitempty"`
	Actory     string `protobuf:"bytes,9,opt,name=Actory,proto3" json:"Actory,omitempty"`
	Language   string `protobuf:"bytes,10,opt,name=Language,proto3" json:"Language,omitempty"`
	Country    string `protobuf:"bytes,11,opt,name=Country,proto3" json:"Country,omitempty"`
	Awards     string `protobuf:"bytes,12,opt,name=Awards,proto3" json:"Awards,omitempty"`
	Poster     string `protobuf:"bytes,13,opt,name=Poster,proto3" json:"Poster,omitempty"`
	ImdbRating string `protobuf:"bytes,14,opt,name=ImdbRating,proto3" json:"ImdbRating,omitempty"`
	ImdbID     string `protobuf:"bytes,15,opt,name=ImdbID,proto3" json:"ImdbID,omitempty"`
	Response   string `protobuf:"bytes,16,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *DetailFilm) Reset() {
	*x = DetailFilm{}
	if protoimpl.UnsafeEnabled {
		mi := &file_request_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetailFilm) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetailFilm) ProtoMessage() {}

func (x *DetailFilm) ProtoReflect() protoreflect.Message {
	mi := &file_request_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetailFilm.ProtoReflect.Descriptor instead.
func (*DetailFilm) Descriptor() ([]byte, []int) {
	return file_request_proto_rawDescGZIP(), []int{3}
}

func (x *DetailFilm) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *DetailFilm) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

func (x *DetailFilm) GetRated() string {
	if x != nil {
		return x.Rated
	}
	return ""
}

func (x *DetailFilm) GetRelease() string {
	if x != nil {
		return x.Release
	}
	return ""
}

func (x *DetailFilm) GetRuntime() string {
	if x != nil {
		return x.Runtime
	}
	return ""
}

func (x *DetailFilm) GetGenre() string {
	if x != nil {
		return x.Genre
	}
	return ""
}

func (x *DetailFilm) GetDirector() string {
	if x != nil {
		return x.Director
	}
	return ""
}

func (x *DetailFilm) GetWrter() string {
	if x != nil {
		return x.Wrter
	}
	return ""
}

func (x *DetailFilm) GetActory() string {
	if x != nil {
		return x.Actory
	}
	return ""
}

func (x *DetailFilm) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *DetailFilm) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *DetailFilm) GetAwards() string {
	if x != nil {
		return x.Awards
	}
	return ""
}

func (x *DetailFilm) GetPoster() string {
	if x != nil {
		return x.Poster
	}
	return ""
}

func (x *DetailFilm) GetImdbRating() string {
	if x != nil {
		return x.ImdbRating
	}
	return ""
}

func (x *DetailFilm) GetImdbID() string {
	if x != nil {
		return x.ImdbID
	}
	return ""
}

func (x *DetailFilm) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

var File_request_proto protoreflect.FileDescriptor

var file_request_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x06, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x22, 0x6c, 0x0a, 0x10, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x41,
	0x70, 0x69, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x70, 0x69,
	0x4b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x49, 0x6d, 0x64, 0x62, 0x49, 0x44, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x49, 0x6d, 0x64, 0x62, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x54,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x50, 0x61, 0x67, 0x65, 0x22, 0x74, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x69, 0x6c,
	0x6d, 0x12, 0x2a, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x46, 0x69, 0x6c, 0x6d, 0x52, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x20, 0x0a,
	0x0b, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0b, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x60, 0x0a, 0x04, 0x46,
	0x69, 0x6c, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x59, 0x65, 0x61,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x59, 0x65, 0x61, 0x72, 0x12, 0x16, 0x0a,
	0x06, 0x69, 0x6d, 0x64, 0x62, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69,
	0x6d, 0x64, 0x62, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x22, 0x9a, 0x03,
	0x0a, 0x0a, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x46, 0x69, 0x6c, 0x6d, 0x12, 0x14, 0x0a, 0x05,
	0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x59, 0x65, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x59, 0x65, 0x61, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x52, 0x61, 0x74, 0x65, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x52, 0x61, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x52,
	0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x57, 0x72, 0x74, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x57, 0x72, 0x74, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x6f,
	0x72, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x79,
	0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x77, 0x61, 0x72, 0x64, 0x73,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x77, 0x61, 0x72, 0x64, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x50, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x49, 0x6d, 0x64, 0x62, 0x52, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x49, 0x6d, 0x64, 0x62,
	0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x49, 0x6d, 0x64, 0x62, 0x49, 0x44,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x49, 0x6d, 0x64, 0x62, 0x49, 0x44, 0x12, 0x1a,
	0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x98, 0x01, 0x0a, 0x0e, 0x46,
	0x69, 0x6c, 0x6d, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x41, 0x0a,
	0x11, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x42, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x47, 0x72,
	0x70, 0x63, 0x12, 0x18, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x1a, 0x10, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x6d, 0x22, 0x00,
	0x12, 0x43, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x46, 0x69, 0x6c,
	0x6d, 0x47, 0x72, 0x70, 0x63, 0x12, 0x18, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x1a,
	0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x46,
	0x69, 0x6c, 0x6d, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_request_proto_rawDescOnce sync.Once
	file_request_proto_rawDescData = file_request_proto_rawDesc
)

func file_request_proto_rawDescGZIP() []byte {
	file_request_proto_rawDescOnce.Do(func() {
		file_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_request_proto_rawDescData)
	})
	return file_request_proto_rawDescData
}

var file_request_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_request_proto_goTypes = []interface{}{
	(*RequestParameter)(nil), // 0: models.RequestParameter
	(*ListFilm)(nil),         // 1: models.ListFilm
	(*Film)(nil),             // 2: models.Film
	(*DetailFilm)(nil),       // 3: models.DetailFilm
}
var file_request_proto_depIdxs = []int32{
	3, // 0: models.ListFilm.Search:type_name -> models.DetailFilm
	0, // 1: models.FilmManagement.SearchByParamGrpc:input_type -> models.RequestParameter
	0, // 2: models.FilmManagement.GetDetailFilmGrpc:input_type -> models.RequestParameter
	1, // 3: models.FilmManagement.SearchByParamGrpc:output_type -> models.ListFilm
	3, // 4: models.FilmManagement.GetDetailFilmGrpc:output_type -> models.DetailFilm
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_request_proto_init() }
func file_request_proto_init() {
	if File_request_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestParameter); i {
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
		file_request_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFilm); i {
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
		file_request_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Film); i {
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
		file_request_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetailFilm); i {
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
			RawDescriptor: file_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_request_proto_goTypes,
		DependencyIndexes: file_request_proto_depIdxs,
		MessageInfos:      file_request_proto_msgTypes,
	}.Build()
	File_request_proto = out.File
	file_request_proto_rawDesc = nil
	file_request_proto_goTypes = nil
	file_request_proto_depIdxs = nil
}