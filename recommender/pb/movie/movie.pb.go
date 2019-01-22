// Code generated by protoc-gen-go. DO NOT EDIT.
// source: movie/movie.proto

package movie

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
	return fileDescriptor_cc162e4bb218d9c6, []int{0}
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
	return fileDescriptor_cc162e4bb218d9c6, []int{1}
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
	return fileDescriptor_cc162e4bb218d9c6, []int{2}
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
	return fileDescriptor_cc162e4bb218d9c6, []int{3}
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
	return fileDescriptor_cc162e4bb218d9c6, []int{4}
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
	proto.RegisterType((*Movie)(nil), "movie.Movie")
	proto.RegisterType((*RecommendRequest)(nil), "movie.RecommendRequest")
	proto.RegisterType((*RecommendResponse)(nil), "movie.RecommendResponse")
	proto.RegisterType((*UpdateRequest)(nil), "movie.UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "movie.UpdateResponse")
}

func init() { proto.RegisterFile("movie/movie.proto", fileDescriptor_cc162e4bb218d9c6) }

var fileDescriptor_cc162e4bb218d9c6 = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x51, 0x4b, 0xfb, 0x30,
	0x14, 0xc5, 0xc9, 0xf6, 0xef, 0xfe, 0xdb, 0x55, 0xe7, 0x16, 0x2a, 0x0b, 0x05, 0xa1, 0x14, 0x1f,
	0x2a, 0x48, 0xc5, 0xf9, 0xa4, 0xaf, 0xa2, 0xe0, 0x83, 0x20, 0x81, 0x3d, 0x97, 0x2e, 0xb9, 0x6a,
	0x60, 0x6d, 0xba, 0x24, 0xf3, 0x7b, 0xf8, 0x8d, 0xa5, 0x69, 0x2d, 0x3a, 0x5f, 0x7c, 0x09, 0x39,
	0xbf, 0x9b, 0x9b, 0x73, 0x4f, 0x02, 0xf3, 0x52, 0xbf, 0x2b, 0xbc, 0xf4, 0x6b, 0x56, 0x1b, 0xed,
	0x34, 0x0d, 0xbc, 0x48, 0xb6, 0x10, 0x3c, 0x35, 0x1b, 0x3a, 0x85, 0x81, 0x92, 0x8c, 0xc4, 0x24,
	0x1d, 0xf2, 0x81, 0x92, 0x34, 0x84, 0xc0, 0x29, 0xb7, 0x41, 0x36, 0x88, 0x49, 0x3a, 0xe1, 0xad,
	0xa0, 0x0b, 0xf8, 0xaf, 0x4a, 0xb9, 0xce, 0x95, 0x64, 0x43, 0xcf, 0x47, 0x8d, 0x7c, 0x94, 0xf4,
	0x1c, 0x66, 0xb5, 0x41, 0xa9, 0x84, 0x43, 0x99, 0x9b, 0xc2, 0xa9, 0xea, 0x95, 0xfd, 0x8b, 0x49,
	0x4a, 0xf8, 0x71, 0xcf, 0xb9, 0xc7, 0xc9, 0x15, 0xcc, 0x38, 0x0a, 0x5d, 0x96, 0x58, 0x49, 0x8e,
	0xdb, 0x1d, 0x5a, 0x47, 0x4f, 0x01, 0x76, 0x16, 0x4d, 0x8e, 0x65, 0xa1, 0x36, 0x7e, 0x8a, 0x09,
	0x9f, 0x34, 0xe4, 0xbe, 0x01, 0xc9, 0x0d, 0xcc, 0xbf, 0xb5, 0xd8, 0x5a, 0x57, 0x16, 0xe9, 0x19,
	0x8c, 0x7c, 0x06, 0xcb, 0x48, 0x3c, 0x4c, 0x0f, 0x96, 0x87, 0x59, 0x9b, 0xcf, 0xe7, 0xe1, 0x5d,
	0x2d, 0xc9, 0xe0, 0x68, 0x55, 0xcb, 0xc2, 0xe1, 0x1f, 0xad, 0x2e, 0x60, 0xfa, 0x75, 0xbe, 0xf3,
	0x89, 0x60, 0x5c, 0x08, 0x81, 0xb5, 0xc3, 0xf6, 0x7d, 0xc6, 0xbc, 0xd7, 0xcb, 0x0f, 0x02, 0xd3,
	0x7e, 0xb2, 0xc2, 0x29, 0x5d, 0xd1, 0x5b, 0x08, 0x1e, 0xd0, 0x89, 0x37, 0xba, 0xe8, 0xe6, 0xd9,
	0x0f, 0x1b, 0xb1, 0xdf, 0x85, 0xce, 0xea, 0x0e, 0xc2, 0xd6, 0x7c, 0x65, 0xd1, 0x3c, 0x1b, 0x7c,
	0x41, 0x83, 0x95, 0x40, 0x1a, 0x76, 0x1d, 0x3f, 0x92, 0x44, 0x27, 0x7b, 0xb4, 0xbd, 0x64, 0x3d,
	0xf2, 0x1f, 0x7c, 0xfd, 0x19, 0x00, 0x00, 0xff, 0xff, 0x2f, 0xdc, 0x56, 0xeb, 0xf5, 0x01, 0x00,
	0x00,
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
	err := c.cc.Invoke(ctx, "/movie.Recommendation/Fetch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recommendationClient) UpdateUserPreference(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/movie.Recommendation/UpdateUserPreference", in, out, opts...)
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
		FullMethod: "/movie.Recommendation/Fetch",
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
		FullMethod: "/movie.Recommendation/UpdateUserPreference",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecommendationServer).UpdateUserPreference(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Recommendation_serviceDesc = grpc.ServiceDesc{
	ServiceName: "movie.Recommendation",
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
	Metadata: "movie/movie.proto",
}
