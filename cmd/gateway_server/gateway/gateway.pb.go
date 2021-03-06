// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gateway.proto

package gateway

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

type SXGRequest struct {
	// URL to fetch from publisher's internal server. ie. PublisherServer.
	FetchUrl string `protobuf:"bytes,1,opt,name=fetch_url,json=fetchUrl,proto3" json:"fetch_url,omitempty"`
	// URL to sign in sxg package.
	SignUrl string `protobuf:"bytes,2,opt,name=sign_url,json=signUrl,proto3" json:"sign_url,omitempty"`
	// ECC private key.
	PrivateKey []byte `protobuf:"bytes,3,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	// X509 certificate which may or may not contain CanSignHTTPExchange
	// extension.
	PublicCert []byte `protobuf:"bytes,4,opt,name=public_cert,json=publicCert,proto3" json:"public_cert,omitempty"`
	// Extra request headers to pass to Packager. singer.go->ServeHTTP.
	PackagerServerRequestHeaders map[string]string `protobuf:"bytes,5,rep,name=packager_server_request_headers,json=packagerServerRequestHeaders,proto3" json:"packager_server_request_headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral         struct{}          `json:"-"`
	XXX_unrecognized             []byte            `json:"-"`
	XXX_sizecache                int32             `json:"-"`
}

func (m *SXGRequest) Reset()         { *m = SXGRequest{} }
func (m *SXGRequest) String() string { return proto.CompactTextString(m) }
func (*SXGRequest) ProtoMessage()    {}
func (*SXGRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1a937782ebbded5, []int{0}
}

func (m *SXGRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SXGRequest.Unmarshal(m, b)
}
func (m *SXGRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SXGRequest.Marshal(b, m, deterministic)
}
func (m *SXGRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SXGRequest.Merge(m, src)
}
func (m *SXGRequest) XXX_Size() int {
	return xxx_messageInfo_SXGRequest.Size(m)
}
func (m *SXGRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SXGRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SXGRequest proto.InternalMessageInfo

func (m *SXGRequest) GetFetchUrl() string {
	if m != nil {
		return m.FetchUrl
	}
	return ""
}

func (m *SXGRequest) GetSignUrl() string {
	if m != nil {
		return m.SignUrl
	}
	return ""
}

func (m *SXGRequest) GetPrivateKey() []byte {
	if m != nil {
		return m.PrivateKey
	}
	return nil
}

func (m *SXGRequest) GetPublicCert() []byte {
	if m != nil {
		return m.PublicCert
	}
	return nil
}

func (m *SXGRequest) GetPackagerServerRequestHeaders() map[string]string {
	if m != nil {
		return m.PackagerServerRequestHeaders
	}
	return nil
}

type SXGResponse struct {
	Sxg              []byte `protobuf:"bytes,1,opt,name=sxg,proto3" json:"sxg,omitempty"`
	Error            bool   `protobuf:"varint,2,opt,name=error,proto3" json:"error,omitempty"`
	ErrorDescription string `protobuf:"bytes,3,opt,name=error_description,json=errorDescription,proto3" json:"error_description,omitempty"`
	// application/cert-chain+cbor format certificate response.
	Cbor []byte `protobuf:"bytes,4,opt,name=cbor,proto3" json:"cbor,omitempty"`
	// HTTP headers returned by the packager.
	HttpHeaders          map[string]string `protobuf:"bytes,5,rep,name=http_headers,json=httpHeaders,proto3" json:"http_headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SXGResponse) Reset()         { *m = SXGResponse{} }
func (m *SXGResponse) String() string { return proto.CompactTextString(m) }
func (*SXGResponse) ProtoMessage()    {}
func (*SXGResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1a937782ebbded5, []int{1}
}

func (m *SXGResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SXGResponse.Unmarshal(m, b)
}
func (m *SXGResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SXGResponse.Marshal(b, m, deterministic)
}
func (m *SXGResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SXGResponse.Merge(m, src)
}
func (m *SXGResponse) XXX_Size() int {
	return xxx_messageInfo_SXGResponse.Size(m)
}
func (m *SXGResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SXGResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SXGResponse proto.InternalMessageInfo

func (m *SXGResponse) GetSxg() []byte {
	if m != nil {
		return m.Sxg
	}
	return nil
}

func (m *SXGResponse) GetError() bool {
	if m != nil {
		return m.Error
	}
	return false
}

func (m *SXGResponse) GetErrorDescription() string {
	if m != nil {
		return m.ErrorDescription
	}
	return ""
}

func (m *SXGResponse) GetCbor() []byte {
	if m != nil {
		return m.Cbor
	}
	return nil
}

func (m *SXGResponse) GetHttpHeaders() map[string]string {
	if m != nil {
		return m.HttpHeaders
	}
	return nil
}

func init() {
	proto.RegisterType((*SXGRequest)(nil), "gateway.SXGRequest")
	proto.RegisterMapType((map[string]string)(nil), "gateway.SXGRequest.PackagerServerRequestHeadersEntry")
	proto.RegisterType((*SXGResponse)(nil), "gateway.SXGResponse")
	proto.RegisterMapType((map[string]string)(nil), "gateway.SXGResponse.HttpHeadersEntry")
}

func init() { proto.RegisterFile("gateway.proto", fileDescriptor_f1a937782ebbded5) }

var fileDescriptor_f1a937782ebbded5 = []byte{
	// 391 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x4f, 0xcb, 0xd3, 0x40,
	0x10, 0xc6, 0x4d, 0xd2, 0xda, 0x76, 0x52, 0xa5, 0xae, 0x3d, 0xc4, 0x2a, 0xb4, 0x16, 0x84, 0x82,
	0x90, 0x43, 0x45, 0x90, 0x1e, 0x3c, 0xf8, 0x87, 0x14, 0x14, 0x94, 0x14, 0xa1, 0xb7, 0xb0, 0x4d,
	0xc7, 0x24, 0x34, 0x24, 0xeb, 0xec, 0xa6, 0x1a, 0xf0, 0xe0, 0x97, 0xf4, 0xbb, 0x78, 0x94, 0xdd,
	0x44, 0x4a, 0x4b, 0xf1, 0xe5, 0xbd, 0xcd, 0x3c, 0x33, 0x4f, 0xf2, 0xcc, 0x2f, 0x81, 0x7b, 0x09,
	0x57, 0xf8, 0x9d, 0xd7, 0xbe, 0xa0, 0x52, 0x95, 0xac, 0xd7, 0xb6, 0xf3, 0xdf, 0x36, 0xc0, 0x66,
	0x1b, 0x84, 0xf8, 0xad, 0x42, 0xa9, 0xd8, 0x63, 0x18, 0x7c, 0x45, 0x15, 0xa7, 0x51, 0x45, 0xb9,
	0x67, 0xcd, 0xac, 0xc5, 0x20, 0xec, 0x1b, 0xe1, 0x0b, 0xe5, 0xec, 0x11, 0xf4, 0x65, 0x96, 0x14,
	0x66, 0x66, 0x9b, 0x59, 0x4f, 0xf7, 0x7a, 0x34, 0x05, 0x57, 0x50, 0x76, 0xe4, 0x0a, 0xa3, 0x03,
	0xd6, 0x9e, 0x33, 0xb3, 0x16, 0xc3, 0x10, 0x5a, 0xe9, 0x03, 0xd6, 0x66, 0xa1, 0xda, 0xe5, 0x59,
	0x1c, 0xc5, 0x48, 0xca, 0xeb, 0xb4, 0x0b, 0x46, 0x7a, 0x8b, 0xa4, 0xd8, 0x4f, 0x98, 0x0a, 0x1e,
	0x1f, 0x78, 0x82, 0x14, 0x49, 0xa4, 0x23, 0x52, 0x44, 0x4d, 0xa8, 0x28, 0x45, 0xbe, 0x47, 0x92,
	0x5e, 0x77, 0xe6, 0x2c, 0xdc, 0xe5, 0x4b, 0xff, 0xdf, 0x29, 0xa7, 0xdc, 0xfe, 0xe7, 0xd6, 0xba,
	0x31, 0xce, 0x56, 0x5d, 0x37, 0xbe, 0xf7, 0x85, 0xa2, 0x3a, 0x7c, 0x22, 0xfe, 0xb3, 0x32, 0xf9,
	0x04, 0x4f, 0x6f, 0x7c, 0x04, 0x1b, 0x81, 0xa3, 0x8f, 0x6b, 0xb0, 0xe8, 0x92, 0x8d, 0xa1, 0x7b,
	0xe4, 0x79, 0x85, 0x2d, 0x8e, 0xa6, 0x59, 0xd9, 0xaf, 0xac, 0xf9, 0x2f, 0x1b, 0x5c, 0x93, 0x4f,
	0x8a, 0xb2, 0x90, 0xa8, 0xbd, 0xf2, 0x47, 0x62, 0xbc, 0xc3, 0x50, 0x97, 0xda, 0x8b, 0x44, 0x25,
	0x19, 0x6f, 0x3f, 0x6c, 0x1a, 0xf6, 0x1c, 0x1e, 0x98, 0x22, 0xda, 0xa3, 0x8c, 0x29, 0x13, 0x2a,
	0x2b, 0x0b, 0x83, 0x73, 0x10, 0x8e, 0xcc, 0xe0, 0xdd, 0x49, 0x67, 0x0c, 0x3a, 0xf1, 0xae, 0xa4,
	0x96, 0xa6, 0xa9, 0xd9, 0x1a, 0x86, 0xa9, 0x52, 0xe2, 0x02, 0xda, 0xb3, 0x73, 0x68, 0x4d, 0x28,
	0x7f, 0xad, 0x94, 0x38, 0x83, 0xe4, 0xa6, 0x27, 0x65, 0xf2, 0x1a, 0x46, 0x97, 0x0b, 0xb7, 0x41,
	0xb0, 0xfc, 0x08, 0xf7, 0x83, 0xe6, 0xa5, 0x1a, 0x69, 0x16, 0x23, 0x5b, 0x81, 0x1b, 0x60, 0x81,
	0xc4, 0x15, 0x6e, 0xb6, 0x01, 0x7b, 0x78, 0xe5, 0x4b, 0x4e, 0xc6, 0xd7, 0x92, 0xce, 0xef, 0xbc,
	0x71, 0xfe, 0x58, 0xd6, 0xee, 0xae, 0xf9, 0x7b, 0x5f, 0xfc, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x15,
	0x84, 0x4b, 0x32, 0xce, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GatewayServiceClient is the client API for GatewayService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GatewayServiceClient interface {
	// Calls Signer.go->FetchUrl method.
	GenerateSXG(ctx context.Context, in *SXGRequest, opts ...grpc.CallOption) (*SXGResponse, error)
}

type gatewayServiceClient struct {
	cc *grpc.ClientConn
}

func NewGatewayServiceClient(cc *grpc.ClientConn) GatewayServiceClient {
	return &gatewayServiceClient{cc}
}

func (c *gatewayServiceClient) GenerateSXG(ctx context.Context, in *SXGRequest, opts ...grpc.CallOption) (*SXGResponse, error) {
	out := new(SXGResponse)
	err := c.cc.Invoke(ctx, "/gateway.GatewayService/GenerateSXG", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServiceServer is the server API for GatewayService service.
type GatewayServiceServer interface {
	// Calls Signer.go->FetchUrl method.
	GenerateSXG(context.Context, *SXGRequest) (*SXGResponse, error)
}

func RegisterGatewayServiceServer(s *grpc.Server, srv GatewayServiceServer) {
	s.RegisterService(&_GatewayService_serviceDesc, srv)
}

func _GatewayService_GenerateSXG_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SXGRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).GenerateSXG(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gateway.GatewayService/GenerateSXG",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).GenerateSXG(ctx, req.(*SXGRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GatewayService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gateway.GatewayService",
	HandlerType: (*GatewayServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateSXG",
			Handler:    _GatewayService_GenerateSXG_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gateway.proto",
}
