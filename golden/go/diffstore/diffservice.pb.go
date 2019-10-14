// Code generated by protoc-gen-go. DO NOT EDIT.
// source: diffservice.proto

package diffstore

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_7be44b2c6ca656f4, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type GetDiffsRequest struct {
	MainDigest           string   `protobuf:"bytes,2,opt,name=mainDigest,proto3" json:"mainDigest,omitempty"`
	RightDigests         []string `protobuf:"bytes,3,rep,name=rightDigests,proto3" json:"rightDigests,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDiffsRequest) Reset()         { *m = GetDiffsRequest{} }
func (m *GetDiffsRequest) String() string { return proto.CompactTextString(m) }
func (*GetDiffsRequest) ProtoMessage()    {}
func (*GetDiffsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7be44b2c6ca656f4, []int{1}
}

func (m *GetDiffsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDiffsRequest.Unmarshal(m, b)
}
func (m *GetDiffsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDiffsRequest.Marshal(b, m, deterministic)
}
func (m *GetDiffsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDiffsRequest.Merge(m, src)
}
func (m *GetDiffsRequest) XXX_Size() int {
	return xxx_messageInfo_GetDiffsRequest.Size(m)
}
func (m *GetDiffsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDiffsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDiffsRequest proto.InternalMessageInfo

func (m *GetDiffsRequest) GetMainDigest() string {
	if m != nil {
		return m.MainDigest
	}
	return ""
}

func (m *GetDiffsRequest) GetRightDigests() []string {
	if m != nil {
		return m.RightDigests
	}
	return nil
}

type GetDiffsResponse struct {
	Diffs                []byte   `protobuf:"bytes,1,opt,name=diffs,proto3" json:"diffs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDiffsResponse) Reset()         { *m = GetDiffsResponse{} }
func (m *GetDiffsResponse) String() string { return proto.CompactTextString(m) }
func (*GetDiffsResponse) ProtoMessage()    {}
func (*GetDiffsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7be44b2c6ca656f4, []int{2}
}

func (m *GetDiffsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDiffsResponse.Unmarshal(m, b)
}
func (m *GetDiffsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDiffsResponse.Marshal(b, m, deterministic)
}
func (m *GetDiffsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDiffsResponse.Merge(m, src)
}
func (m *GetDiffsResponse) XXX_Size() int {
	return xxx_messageInfo_GetDiffsResponse.Size(m)
}
func (m *GetDiffsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDiffsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetDiffsResponse proto.InternalMessageInfo

func (m *GetDiffsResponse) GetDiffs() []byte {
	if m != nil {
		return m.Diffs
	}
	return nil
}

type PurgeDigestsRequest struct {
	Digests              []string `protobuf:"bytes,1,rep,name=digests,proto3" json:"digests,omitempty"`
	PurgeGCS             bool     `protobuf:"varint,2,opt,name=purgeGCS,proto3" json:"purgeGCS,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PurgeDigestsRequest) Reset()         { *m = PurgeDigestsRequest{} }
func (m *PurgeDigestsRequest) String() string { return proto.CompactTextString(m) }
func (*PurgeDigestsRequest) ProtoMessage()    {}
func (*PurgeDigestsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7be44b2c6ca656f4, []int{3}
}

func (m *PurgeDigestsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PurgeDigestsRequest.Unmarshal(m, b)
}
func (m *PurgeDigestsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PurgeDigestsRequest.Marshal(b, m, deterministic)
}
func (m *PurgeDigestsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PurgeDigestsRequest.Merge(m, src)
}
func (m *PurgeDigestsRequest) XXX_Size() int {
	return xxx_messageInfo_PurgeDigestsRequest.Size(m)
}
func (m *PurgeDigestsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PurgeDigestsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PurgeDigestsRequest proto.InternalMessageInfo

func (m *PurgeDigestsRequest) GetDigests() []string {
	if m != nil {
		return m.Digests
	}
	return nil
}

func (m *PurgeDigestsRequest) GetPurgeGCS() bool {
	if m != nil {
		return m.PurgeGCS
	}
	return false
}

type UnavailableDigestsResponse struct {
	DigestFailures       map[string]*DigestFailureResponse `protobuf:"bytes,1,rep,name=digestFailures,proto3" json:"digestFailures,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *UnavailableDigestsResponse) Reset()         { *m = UnavailableDigestsResponse{} }
func (m *UnavailableDigestsResponse) String() string { return proto.CompactTextString(m) }
func (*UnavailableDigestsResponse) ProtoMessage()    {}
func (*UnavailableDigestsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7be44b2c6ca656f4, []int{4}
}

func (m *UnavailableDigestsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnavailableDigestsResponse.Unmarshal(m, b)
}
func (m *UnavailableDigestsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnavailableDigestsResponse.Marshal(b, m, deterministic)
}
func (m *UnavailableDigestsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnavailableDigestsResponse.Merge(m, src)
}
func (m *UnavailableDigestsResponse) XXX_Size() int {
	return xxx_messageInfo_UnavailableDigestsResponse.Size(m)
}
func (m *UnavailableDigestsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UnavailableDigestsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UnavailableDigestsResponse proto.InternalMessageInfo

func (m *UnavailableDigestsResponse) GetDigestFailures() map[string]*DigestFailureResponse {
	if m != nil {
		return m.DigestFailures
	}
	return nil
}

type DigestFailureResponse struct {
	Digest               string   `protobuf:"bytes,1,opt,name=Digest,proto3" json:"Digest,omitempty"`
	Reason               string   `protobuf:"bytes,2,opt,name=Reason,proto3" json:"Reason,omitempty"`
	TS                   int64    `protobuf:"varint,3,opt,name=TS,proto3" json:"TS,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DigestFailureResponse) Reset()         { *m = DigestFailureResponse{} }
func (m *DigestFailureResponse) String() string { return proto.CompactTextString(m) }
func (*DigestFailureResponse) ProtoMessage()    {}
func (*DigestFailureResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7be44b2c6ca656f4, []int{5}
}

func (m *DigestFailureResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DigestFailureResponse.Unmarshal(m, b)
}
func (m *DigestFailureResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DigestFailureResponse.Marshal(b, m, deterministic)
}
func (m *DigestFailureResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DigestFailureResponse.Merge(m, src)
}
func (m *DigestFailureResponse) XXX_Size() int {
	return xxx_messageInfo_DigestFailureResponse.Size(m)
}
func (m *DigestFailureResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DigestFailureResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DigestFailureResponse proto.InternalMessageInfo

func (m *DigestFailureResponse) GetDigest() string {
	if m != nil {
		return m.Digest
	}
	return ""
}

func (m *DigestFailureResponse) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func (m *DigestFailureResponse) GetTS() int64 {
	if m != nil {
		return m.TS
	}
	return 0
}

func init() {
	proto.RegisterType((*Empty)(nil), "diffstore.Empty")
	proto.RegisterType((*GetDiffsRequest)(nil), "diffstore.GetDiffsRequest")
	proto.RegisterType((*GetDiffsResponse)(nil), "diffstore.GetDiffsResponse")
	proto.RegisterType((*PurgeDigestsRequest)(nil), "diffstore.PurgeDigestsRequest")
	proto.RegisterType((*UnavailableDigestsResponse)(nil), "diffstore.UnavailableDigestsResponse")
	proto.RegisterMapType((map[string]*DigestFailureResponse)(nil), "diffstore.UnavailableDigestsResponse.DigestFailuresEntry")
	proto.RegisterType((*DigestFailureResponse)(nil), "diffstore.DigestFailureResponse")
}

func init() { proto.RegisterFile("diffservice.proto", fileDescriptor_7be44b2c6ca656f4) }

var fileDescriptor_7be44b2c6ca656f4 = []byte{
	// 405 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xc1, 0xae, 0x93, 0x40,
	0x14, 0x2d, 0x60, 0xdf, 0x2b, 0xb7, 0xcd, 0xb3, 0xde, 0xa7, 0x86, 0x60, 0xf2, 0x42, 0x26, 0x31,
	0x61, 0x61, 0x58, 0xd4, 0xc4, 0xa8, 0x4b, 0x6d, 0x7d, 0x0b, 0x17, 0x36, 0x43, 0x1b, 0xd7, 0xd3,
	0x76, 0x8a, 0x13, 0x29, 0x20, 0x0c, 0x4d, 0xfa, 0x25, 0xfe, 0xa0, 0x1f, 0x62, 0x98, 0x81, 0x96,
	0xb6, 0x98, 0xb8, 0xe3, 0xdc, 0xb9, 0xe7, 0xdc, 0x33, 0x73, 0x0f, 0xf0, 0x6c, 0x23, 0xb6, 0xdb,
	0x82, 0xe7, 0x7b, 0xb1, 0xe6, 0x41, 0x96, 0xa7, 0x32, 0x45, 0x5b, 0x95, 0x64, 0x9a, 0x73, 0x72,
	0x0b, 0xfd, 0xd9, 0x2e, 0x93, 0x07, 0xb2, 0x84, 0xa7, 0x8f, 0x5c, 0x4e, 0xab, 0x03, 0xca, 0x7f,
	0x95, 0xbc, 0x90, 0xf8, 0x00, 0xb0, 0x63, 0x22, 0x99, 0x8a, 0x88, 0x17, 0xd2, 0x31, 0x3d, 0xc3,
	0xb7, 0x69, 0xab, 0x82, 0x04, 0x46, 0xb9, 0x88, 0x7e, 0x48, 0x0d, 0x0b, 0xc7, 0xf2, 0x2c, 0xdf,
	0xa6, 0x67, 0x35, 0xe2, 0xc3, 0xf8, 0x24, 0x5b, 0x64, 0x69, 0x52, 0x70, 0x7c, 0x0e, 0x7d, 0x65,
	0xc0, 0x31, 0x3c, 0xc3, 0x1f, 0x51, 0x0d, 0xc8, 0x57, 0xb8, 0x9f, 0x97, 0x79, 0xc4, 0x6b, 0x66,
	0x63, 0xc2, 0x81, 0xdb, 0x4d, 0xad, 0x6f, 0x28, 0xfd, 0x06, 0xa2, 0x0b, 0x83, 0xac, 0x22, 0x3c,
	0x7e, 0x0e, 0x95, 0xb9, 0x01, 0x3d, 0x62, 0xf2, 0xc7, 0x00, 0x77, 0x99, 0xb0, 0x3d, 0x13, 0x31,
	0x5b, 0xc5, 0x27, 0xcd, 0xda, 0x01, 0x83, 0x3b, 0xad, 0xf2, 0x85, 0x89, 0xb8, 0xcc, 0xb9, 0xd6,
	0x1e, 0x4e, 0x3e, 0x04, 0xc7, 0x97, 0x09, 0xfe, 0x4d, 0x0f, 0xa6, 0x67, 0xdc, 0x59, 0x22, 0xf3,
	0x03, 0xbd, 0x10, 0x74, 0xd7, 0x70, 0xdf, 0xd1, 0x86, 0x63, 0xb0, 0x7e, 0xf2, 0x83, 0xba, 0xb9,
	0x4d, 0xab, 0x4f, 0x7c, 0x07, 0xfd, 0x3d, 0x8b, 0x4b, 0xae, 0xee, 0x30, 0x9c, 0x78, 0x2d, 0x0b,
	0x67, 0x02, 0xcd, 0x74, 0xaa, 0xdb, 0x3f, 0x9a, 0xef, 0x0d, 0xf2, 0x1d, 0x5e, 0x74, 0xf6, 0xe0,
	0x4b, 0xb8, 0xa9, 0xd7, 0xa6, 0x27, 0xd5, 0xa8, 0xaa, 0x53, 0xce, 0x8a, 0x34, 0xa9, 0xd7, 0x59,
	0x23, 0xbc, 0x03, 0x73, 0x11, 0x3a, 0x96, 0x67, 0xf8, 0x16, 0x35, 0x17, 0xe1, 0xe4, 0xb7, 0x09,
	0xc3, 0x6a, 0x69, 0xa1, 0xce, 0x0d, 0xce, 0x60, 0xd0, 0xac, 0x11, 0xdd, 0x96, 0xc3, 0x8b, 0xc8,
	0xb8, 0xaf, 0x3a, 0xcf, 0xb4, 0x29, 0xd2, 0xc3, 0x6f, 0x80, 0xd7, 0xcf, 0x8a, 0xe3, 0x16, 0x49,
	0x85, 0xd1, 0x7d, 0xfd, 0x5f, 0x7b, 0x20, 0x3d, 0xfc, 0x04, 0xa3, 0x76, 0x68, 0xf0, 0xa1, 0x45,
	0xec, 0x48, 0x93, 0x7b, 0x35, 0x8a, 0xf4, 0xf0, 0x0d, 0x3c, 0x99, 0x8b, 0x24, 0xea, 0xb0, 0xd1,
	0xd1, 0xbd, 0xba, 0x51, 0xbf, 0xd0, 0xdb, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x61, 0x1f, 0xa3,
	0x61, 0x57, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DiffServiceClient is the client API for DiffService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DiffServiceClient interface {
	// Same functionality as Get in the diff.DiffStore interface.
	GetDiffs(ctx context.Context, in *GetDiffsRequest, opts ...grpc.CallOption) (*GetDiffsResponse, error)
	// Same functionality asSee UnavailableDigests in the diff.DiffStore interface.
	UnavailableDigests(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*UnavailableDigestsResponse, error)
	//Same functionality asSee PurgeDigestset in the diff.DiffStore interface.
	PurgeDigests(ctx context.Context, in *PurgeDigestsRequest, opts ...grpc.CallOption) (*Empty, error)
	// Ping is used to test connection.
	Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
}

type diffServiceClient struct {
	cc *grpc.ClientConn
}

func NewDiffServiceClient(cc *grpc.ClientConn) DiffServiceClient {
	return &diffServiceClient{cc}
}

func (c *diffServiceClient) GetDiffs(ctx context.Context, in *GetDiffsRequest, opts ...grpc.CallOption) (*GetDiffsResponse, error) {
	out := new(GetDiffsResponse)
	err := c.cc.Invoke(ctx, "/diffstore.DiffService/GetDiffs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *diffServiceClient) UnavailableDigests(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*UnavailableDigestsResponse, error) {
	out := new(UnavailableDigestsResponse)
	err := c.cc.Invoke(ctx, "/diffstore.DiffService/UnavailableDigests", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *diffServiceClient) PurgeDigests(ctx context.Context, in *PurgeDigestsRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/diffstore.DiffService/PurgeDigests", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *diffServiceClient) Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/diffstore.DiffService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DiffServiceServer is the server API for DiffService service.
type DiffServiceServer interface {
	// Same functionality as Get in the diff.DiffStore interface.
	GetDiffs(context.Context, *GetDiffsRequest) (*GetDiffsResponse, error)
	// Same functionality asSee UnavailableDigests in the diff.DiffStore interface.
	UnavailableDigests(context.Context, *Empty) (*UnavailableDigestsResponse, error)
	//Same functionality asSee PurgeDigestset in the diff.DiffStore interface.
	PurgeDigests(context.Context, *PurgeDigestsRequest) (*Empty, error)
	// Ping is used to test connection.
	Ping(context.Context, *Empty) (*Empty, error)
}

// UnimplementedDiffServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDiffServiceServer struct {
}

func (*UnimplementedDiffServiceServer) GetDiffs(ctx context.Context, req *GetDiffsRequest) (*GetDiffsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDiffs not implemented")
}
func (*UnimplementedDiffServiceServer) UnavailableDigests(ctx context.Context, req *Empty) (*UnavailableDigestsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnavailableDigests not implemented")
}
func (*UnimplementedDiffServiceServer) PurgeDigests(ctx context.Context, req *PurgeDigestsRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PurgeDigests not implemented")
}
func (*UnimplementedDiffServiceServer) Ping(ctx context.Context, req *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}

func RegisterDiffServiceServer(s *grpc.Server, srv DiffServiceServer) {
	s.RegisterService(&_DiffService_serviceDesc, srv)
}

func _DiffService_GetDiffs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDiffsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiffServiceServer).GetDiffs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/diffstore.DiffService/GetDiffs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiffServiceServer).GetDiffs(ctx, req.(*GetDiffsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiffService_UnavailableDigests_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiffServiceServer).UnavailableDigests(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/diffstore.DiffService/UnavailableDigests",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiffServiceServer).UnavailableDigests(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiffService_PurgeDigests_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PurgeDigestsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiffServiceServer).PurgeDigests(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/diffstore.DiffService/PurgeDigests",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiffServiceServer).PurgeDigests(ctx, req.(*PurgeDigestsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiffService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiffServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/diffstore.DiffService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiffServiceServer).Ping(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _DiffService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "diffstore.DiffService",
	HandlerType: (*DiffServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDiffs",
			Handler:    _DiffService_GetDiffs_Handler,
		},
		{
			MethodName: "UnavailableDigests",
			Handler:    _DiffService_UnavailableDigests_Handler,
		},
		{
			MethodName: "PurgeDigests",
			Handler:    _DiffService_PurgeDigests_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _DiffService_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "diffservice.proto",
}
