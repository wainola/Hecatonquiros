// Code generated by protoc-gen-go. DO NOT EDIT.
// source: list.proto

package list

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

type List struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Content              string   `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *List) Reset()         { *m = List{} }
func (m *List) String() string { return proto.CompactTextString(m) }
func (*List) ProtoMessage()    {}
func (*List) Descriptor() ([]byte, []int) {
	return fileDescriptor_af793ce248ee1bf0, []int{0}
}

func (m *List) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_List.Unmarshal(m, b)
}
func (m *List) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_List.Marshal(b, m, deterministic)
}
func (m *List) XXX_Merge(src proto.Message) {
	xxx_messageInfo_List.Merge(m, src)
}
func (m *List) XXX_Size() int {
	return xxx_messageInfo_List.Size(m)
}
func (m *List) XXX_DiscardUnknown() {
	xxx_messageInfo_List.DiscardUnknown(m)
}

var xxx_messageInfo_List proto.InternalMessageInfo

func (m *List) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *List) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type GetListReq struct {
	ListId               string   `protobuf:"bytes,1,opt,name=list_id,json=listId,proto3" json:"list_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetListReq) Reset()         { *m = GetListReq{} }
func (m *GetListReq) String() string { return proto.CompactTextString(m) }
func (*GetListReq) ProtoMessage()    {}
func (*GetListReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_af793ce248ee1bf0, []int{1}
}

func (m *GetListReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetListReq.Unmarshal(m, b)
}
func (m *GetListReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetListReq.Marshal(b, m, deterministic)
}
func (m *GetListReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetListReq.Merge(m, src)
}
func (m *GetListReq) XXX_Size() int {
	return xxx_messageInfo_GetListReq.Size(m)
}
func (m *GetListReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetListReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetListReq proto.InternalMessageInfo

func (m *GetListReq) GetListId() string {
	if m != nil {
		return m.ListId
	}
	return ""
}

type ListResp struct {
	Items                []*List  `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListResp) Reset()         { *m = ListResp{} }
func (m *ListResp) String() string { return proto.CompactTextString(m) }
func (*ListResp) ProtoMessage()    {}
func (*ListResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_af793ce248ee1bf0, []int{2}
}

func (m *ListResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResp.Unmarshal(m, b)
}
func (m *ListResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResp.Marshal(b, m, deterministic)
}
func (m *ListResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResp.Merge(m, src)
}
func (m *ListResp) XXX_Size() int {
	return xxx_messageInfo_ListResp.Size(m)
}
func (m *ListResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResp.DiscardUnknown(m)
}

var xxx_messageInfo_ListResp proto.InternalMessageInfo

func (m *ListResp) GetItems() []*List {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterType((*List)(nil), "list.List")
	proto.RegisterType((*GetListReq)(nil), "list.GetListReq")
	proto.RegisterType((*ListResp)(nil), "list.ListResp")
}

func init() {
	proto.RegisterFile("list.proto", fileDescriptor_af793ce248ee1bf0)
}

var fileDescriptor_af793ce248ee1bf0 = []byte{
	// 173 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xca, 0xc9, 0x2c, 0x2e,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0x0c, 0xb8, 0x58, 0x7c, 0x80,
	0xb4, 0x10, 0x1f, 0x17, 0x53, 0x66, 0x8a, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x90, 0x25,
	0x24, 0xc1, 0xc5, 0x9e, 0x9c, 0x9f, 0x57, 0x92, 0x9a, 0x57, 0x22, 0xc1, 0x04, 0x16, 0x84, 0x71,
	0x95, 0x54, 0xb9, 0xb8, 0xdc, 0x53, 0x4b, 0x40, 0x9a, 0x82, 0x52, 0x0b, 0x85, 0xc4, 0xb9, 0xd8,
	0x41, 0xe6, 0xc4, 0xc3, 0x35, 0xb3, 0x81, 0xb8, 0x9e, 0x29, 0x4a, 0x3a, 0x5c, 0x1c, 0x10, 0x35,
	0xc5, 0x05, 0x42, 0x0a, 0x5c, 0xac, 0x99, 0x25, 0xa9, 0xb9, 0xc5, 0x40, 0x25, 0xcc, 0x1a, 0xdc,
	0x46, 0x5c, 0x7a, 0x60, 0x67, 0x80, 0xa5, 0x21, 0x12, 0x46, 0x36, 0x5c, 0xdc, 0x20, 0x6e, 0x70,
	0x6a, 0x51, 0x59, 0x66, 0x72, 0xaa, 0x90, 0x2e, 0x17, 0x3b, 0xd4, 0x0e, 0x21, 0x01, 0x88, 0x62,
	0x84, 0x95, 0x52, 0x7c, 0x48, 0xda, 0x81, 0xa6, 0x2b, 0x31, 0x24, 0xb1, 0x81, 0x7d, 0x64, 0x0c,
	0x08, 0x00, 0x00, 0xff, 0xff, 0x34, 0x7f, 0x5e, 0x64, 0xdf, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ListServiceClient is the client API for ListService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ListServiceClient interface {
	GetList(ctx context.Context, in *GetListReq, opts ...grpc.CallOption) (*ListResp, error)
}

type listServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewListServiceClient(cc grpc.ClientConnInterface) ListServiceClient {
	return &listServiceClient{cc}
}

func (c *listServiceClient) GetList(ctx context.Context, in *GetListReq, opts ...grpc.CallOption) (*ListResp, error) {
	out := new(ListResp)
	err := c.cc.Invoke(ctx, "/list.ListService/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ListServiceServer is the server API for ListService service.
type ListServiceServer interface {
	GetList(context.Context, *GetListReq) (*ListResp, error)
}

// UnimplementedListServiceServer can be embedded to have forward compatible implementations.
type UnimplementedListServiceServer struct {
}

func (*UnimplementedListServiceServer) GetList(ctx context.Context, req *GetListReq) (*ListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}

func RegisterListServiceServer(s *grpc.Server, srv ListServiceServer) {
	s.RegisterService(&_ListService_serviceDesc, srv)
}

func _ListService_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListServiceServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/list.ListService/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListServiceServer).GetList(ctx, req.(*GetListReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _ListService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "list.ListService",
	HandlerType: (*ListServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetList",
			Handler:    _ListService_GetList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "list.proto",
}