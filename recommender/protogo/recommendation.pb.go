// Code generated by protoc-gen-go. DO NOT EDIT.
// source: recommendation.proto

package protogo

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

type RecommendRequest struct {
	UserEmail            string   `protobuf:"bytes,1,opt,name=user_email,json=userEmail,proto3" json:"user_email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RecommendRequest) Reset()         { *m = RecommendRequest{} }
func (m *RecommendRequest) String() string { return proto.CompactTextString(m) }
func (*RecommendRequest) ProtoMessage()    {}
func (*RecommendRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_44af9f958e489253, []int{0}
}

func (m *RecommendRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecommendRequest.Unmarshal(m, b)
}
func (m *RecommendRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecommendRequest.Marshal(b, m, deterministic)
}
func (m *RecommendRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecommendRequest.Merge(m, src)
}
func (m *RecommendRequest) XXX_Size() int {
	return xxx_messageInfo_RecommendRequest.Size(m)
}
func (m *RecommendRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RecommendRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RecommendRequest proto.InternalMessageInfo

func (m *RecommendRequest) GetUserEmail() string {
	if m != nil {
		return m.UserEmail
	}
	return ""
}

type Movie struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	ImdbId               string   `protobuf:"bytes,3,opt,name=imdb_id,json=imdbId,proto3" json:"imdb_id,omitempty"`
	PredictedRating      float64  `protobuf:"fixed64,4,opt,name=predicted_rating,json=predictedRating,proto3" json:"predicted_rating,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Movie) Reset()         { *m = Movie{} }
func (m *Movie) String() string { return proto.CompactTextString(m) }
func (*Movie) ProtoMessage()    {}
func (*Movie) Descriptor() ([]byte, []int) {
	return fileDescriptor_44af9f958e489253, []int{1}
}

func (m *Movie) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Movie.Unmarshal(m, b)
}
func (m *Movie) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Movie.Marshal(b, m, deterministic)
}
func (m *Movie) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Movie.Merge(m, src)
}
func (m *Movie) XXX_Size() int {
	return xxx_messageInfo_Movie.Size(m)
}
func (m *Movie) XXX_DiscardUnknown() {
	xxx_messageInfo_Movie.DiscardUnknown(m)
}

var xxx_messageInfo_Movie proto.InternalMessageInfo

func (m *Movie) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Movie) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Movie) GetImdbId() string {
	if m != nil {
		return m.ImdbId
	}
	return ""
}

func (m *Movie) GetPredictedRating() float64 {
	if m != nil {
		return m.PredictedRating
	}
	return 0
}

type RecommendResponse struct {
	Movies               []*Movie `protobuf:"bytes,1,rep,name=movies,proto3" json:"movies,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RecommendResponse) Reset()         { *m = RecommendResponse{} }
func (m *RecommendResponse) String() string { return proto.CompactTextString(m) }
func (*RecommendResponse) ProtoMessage()    {}
func (*RecommendResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_44af9f958e489253, []int{2}
}

func (m *RecommendResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecommendResponse.Unmarshal(m, b)
}
func (m *RecommendResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecommendResponse.Marshal(b, m, deterministic)
}
func (m *RecommendResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecommendResponse.Merge(m, src)
}
func (m *RecommendResponse) XXX_Size() int {
	return xxx_messageInfo_RecommendResponse.Size(m)
}
func (m *RecommendResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RecommendResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RecommendResponse proto.InternalMessageInfo

func (m *RecommendResponse) GetMovies() []*Movie {
	if m != nil {
		return m.Movies
	}
	return nil
}

type UpdateRequest struct {
	UserEmail            string   `protobuf:"bytes,1,opt,name=user_email,json=userEmail,proto3" json:"user_email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_44af9f958e489253, []int{3}
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

func (m *UpdateRequest) GetUserEmail() string {
	if m != nil {
		return m.UserEmail
	}
	return ""
}

type UpdateResponse struct {
	Accepted             bool     `protobuf:"varint,1,opt,name=accepted,proto3" json:"accepted,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateResponse) Reset()         { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()    {}
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_44af9f958e489253, []int{4}
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

func (m *UpdateResponse) GetAccepted() bool {
	if m != nil {
		return m.Accepted
	}
	return false
}

func init() {
	proto.RegisterType((*RecommendRequest)(nil), "RecommendRequest")
	proto.RegisterType((*Movie)(nil), "Movie")
	proto.RegisterType((*RecommendResponse)(nil), "RecommendResponse")
	proto.RegisterType((*UpdateRequest)(nil), "UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "UpdateResponse")
}

func init() { proto.RegisterFile("recommendation.proto", fileDescriptor_44af9f958e489253) }

var fileDescriptor_44af9f958e489253 = []byte{
	// 295 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xc1, 0x4b, 0xc3, 0x30,
	0x14, 0x87, 0xe9, 0xe6, 0xba, 0xf5, 0x89, 0xdd, 0x16, 0x0a, 0x96, 0x82, 0x52, 0x7a, 0xaa, 0x20,
	0x01, 0xb7, 0x83, 0x77, 0x41, 0xc1, 0x83, 0x20, 0x81, 0x5d, 0xbc, 0x94, 0x2e, 0x79, 0xce, 0xc0,
	0xda, 0x74, 0x49, 0x26, 0xf8, 0xdf, 0x4b, 0xd3, 0x6d, 0x58, 0xbd, 0x78, 0xcb, 0xef, 0x4b, 0xc2,
	0xfb, 0xde, 0x7b, 0x10, 0x69, 0xe4, 0xaa, 0xaa, 0xb0, 0x16, 0xa5, 0x95, 0xaa, 0xa6, 0x8d, 0x56,
	0x56, 0x65, 0x77, 0x30, 0x63, 0x47, 0xce, 0x70, 0xb7, 0x47, 0x63, 0xc9, 0x15, 0xc0, 0xde, 0xa0,
	0x2e, 0xb0, 0x2a, 0xe5, 0x36, 0xf6, 0x52, 0x2f, 0x0f, 0x58, 0xd0, 0x92, 0xc7, 0x16, 0x64, 0x3b,
	0x18, 0xbd, 0xa8, 0x4f, 0x89, 0x24, 0x84, 0x81, 0x14, 0xee, 0x7e, 0xc8, 0x06, 0x52, 0x90, 0x08,
	0x46, 0x56, 0xda, 0x2d, 0xc6, 0x03, 0xf7, 0xa5, 0x0b, 0xe4, 0x12, 0xc6, 0xb2, 0x12, 0xeb, 0x42,
	0x8a, 0x78, 0xe8, 0xb8, 0xdf, 0xc6, 0x67, 0x41, 0x6e, 0x60, 0xd6, 0x68, 0x14, 0x92, 0x5b, 0x14,
	0x85, 0x2e, 0xad, 0xac, 0x37, 0xf1, 0x59, 0xea, 0xe5, 0x1e, 0x9b, 0x9e, 0x38, 0x73, 0x38, 0x5b,
	0xc2, 0xfc, 0x87, 0xa5, 0x69, 0x54, 0x6d, 0x90, 0x5c, 0x83, 0x5f, 0xb5, 0x1e, 0x26, 0xf6, 0xd2,
	0x61, 0x7e, 0xbe, 0xf0, 0xa9, 0xd3, 0x62, 0x07, 0x9a, 0x51, 0xb8, 0x58, 0x35, 0xa2, 0xb4, 0xf8,
	0xcf, 0xbe, 0x6e, 0x21, 0x3c, 0xbe, 0x3f, 0x54, 0x48, 0x60, 0x52, 0x72, 0x8e, 0x8d, 0xc5, 0xae,
	0xcd, 0x09, 0x3b, 0xe5, 0xc5, 0x17, 0x84, 0xac, 0x37, 0x50, 0x42, 0x61, 0xf4, 0x84, 0x96, 0x7f,
	0x90, 0x39, 0xfd, 0x3d, 0xd2, 0x84, 0xd0, 0xbf, 0xfe, 0xf7, 0x10, 0x75, 0xf5, 0x56, 0x06, 0xf5,
	0xab, 0xc6, 0x77, 0xd4, 0x58, 0x73, 0x24, 0x21, 0xed, 0x69, 0x27, 0x53, 0xda, 0xd7, 0x7a, 0x08,
	0xde, 0xc6, 0x6e, 0x79, 0x1b, 0xb5, 0xf6, 0xdd, 0x61, 0xf9, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x5c,
	0x74, 0x52, 0x6c, 0xdd, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RecommendationClient is the client API for Recommendation service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RecommendationClient interface {
	Fetch(ctx context.Context, in *RecommendRequest, opts ...grpc.CallOption) (*RecommendResponse, error)
	UpdateUserPreference(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
}

type recommendationClient struct {
	cc *grpc.ClientConn
}

func NewRecommendationClient(cc *grpc.ClientConn) RecommendationClient {
	return &recommendationClient{cc}
}

func (c *recommendationClient) Fetch(ctx context.Context, in *RecommendRequest, opts ...grpc.CallOption) (*RecommendResponse, error) {
	out := new(RecommendResponse)
	err := c.cc.Invoke(ctx, "/Recommendation/Fetch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recommendationClient) UpdateUserPreference(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/Recommendation/UpdateUserPreference", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RecommendationServer is the server API for Recommendation service.
type RecommendationServer interface {
	Fetch(context.Context, *RecommendRequest) (*RecommendResponse, error)
	UpdateUserPreference(context.Context, *UpdateRequest) (*UpdateResponse, error)
}

func RegisterRecommendationServer(s *grpc.Server, srv RecommendationServer) {
	s.RegisterService(&_Recommendation_serviceDesc, srv)
}

func _Recommendation_Fetch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecommendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecommendationServer).Fetch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Recommendation/Fetch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecommendationServer).Fetch(ctx, req.(*RecommendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Recommendation_UpdateUserPreference_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecommendationServer).UpdateUserPreference(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Recommendation/UpdateUserPreference",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecommendationServer).UpdateUserPreference(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Recommendation_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Recommendation",
	HandlerType: (*RecommendationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Fetch",
			Handler:    _Recommendation_Fetch_Handler,
		},
		{
			MethodName: "UpdateUserPreference",
			Handler:    _Recommendation_UpdateUserPreference_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "recommendation.proto",
}
