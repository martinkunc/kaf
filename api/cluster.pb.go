// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/cluster.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type CreateClusterRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateClusterRequest) Reset()         { *m = CreateClusterRequest{} }
func (m *CreateClusterRequest) String() string { return proto.CompactTextString(m) }
func (*CreateClusterRequest) ProtoMessage()    {}
func (*CreateClusterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_963a9d505aa9dcea, []int{0}
}

func (m *CreateClusterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateClusterRequest.Unmarshal(m, b)
}
func (m *CreateClusterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateClusterRequest.Marshal(b, m, deterministic)
}
func (m *CreateClusterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateClusterRequest.Merge(m, src)
}
func (m *CreateClusterRequest) XXX_Size() int {
	return xxx_messageInfo_CreateClusterRequest.Size(m)
}
func (m *CreateClusterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateClusterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateClusterRequest proto.InternalMessageInfo

type CreateClusterResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateClusterResponse) Reset()         { *m = CreateClusterResponse{} }
func (m *CreateClusterResponse) String() string { return proto.CompactTextString(m) }
func (*CreateClusterResponse) ProtoMessage()    {}
func (*CreateClusterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_963a9d505aa9dcea, []int{1}
}

func (m *CreateClusterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateClusterResponse.Unmarshal(m, b)
}
func (m *CreateClusterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateClusterResponse.Marshal(b, m, deterministic)
}
func (m *CreateClusterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateClusterResponse.Merge(m, src)
}
func (m *CreateClusterResponse) XXX_Size() int {
	return xxx_messageInfo_CreateClusterResponse.Size(m)
}
func (m *CreateClusterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateClusterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateClusterResponse proto.InternalMessageInfo

type ListClustersRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListClustersRequest) Reset()         { *m = ListClustersRequest{} }
func (m *ListClustersRequest) String() string { return proto.CompactTextString(m) }
func (*ListClustersRequest) ProtoMessage()    {}
func (*ListClustersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_963a9d505aa9dcea, []int{2}
}

func (m *ListClustersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListClustersRequest.Unmarshal(m, b)
}
func (m *ListClustersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListClustersRequest.Marshal(b, m, deterministic)
}
func (m *ListClustersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListClustersRequest.Merge(m, src)
}
func (m *ListClustersRequest) XXX_Size() int {
	return xxx_messageInfo_ListClustersRequest.Size(m)
}
func (m *ListClustersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListClustersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListClustersRequest proto.InternalMessageInfo

type ListClustersResponse struct {
	Clusters             []*Cluster `protobuf:"bytes,1,rep,name=clusters,proto3" json:"clusters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ListClustersResponse) Reset()         { *m = ListClustersResponse{} }
func (m *ListClustersResponse) String() string { return proto.CompactTextString(m) }
func (*ListClustersResponse) ProtoMessage()    {}
func (*ListClustersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_963a9d505aa9dcea, []int{3}
}

func (m *ListClustersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListClustersResponse.Unmarshal(m, b)
}
func (m *ListClustersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListClustersResponse.Marshal(b, m, deterministic)
}
func (m *ListClustersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListClustersResponse.Merge(m, src)
}
func (m *ListClustersResponse) XXX_Size() int {
	return xxx_messageInfo_ListClustersResponse.Size(m)
}
func (m *ListClustersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListClustersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListClustersResponse proto.InternalMessageInfo

func (m *ListClustersResponse) GetClusters() []*Cluster {
	if m != nil {
		return m.Clusters
	}
	return nil
}

type GetClusterRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetClusterRequest) Reset()         { *m = GetClusterRequest{} }
func (m *GetClusterRequest) String() string { return proto.CompactTextString(m) }
func (*GetClusterRequest) ProtoMessage()    {}
func (*GetClusterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_963a9d505aa9dcea, []int{4}
}

func (m *GetClusterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetClusterRequest.Unmarshal(m, b)
}
func (m *GetClusterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetClusterRequest.Marshal(b, m, deterministic)
}
func (m *GetClusterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetClusterRequest.Merge(m, src)
}
func (m *GetClusterRequest) XXX_Size() int {
	return xxx_messageInfo_GetClusterRequest.Size(m)
}
func (m *GetClusterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetClusterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetClusterRequest proto.InternalMessageInfo

type DeleteClusterRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteClusterRequest) Reset()         { *m = DeleteClusterRequest{} }
func (m *DeleteClusterRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteClusterRequest) ProtoMessage()    {}
func (*DeleteClusterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_963a9d505aa9dcea, []int{5}
}

func (m *DeleteClusterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteClusterRequest.Unmarshal(m, b)
}
func (m *DeleteClusterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteClusterRequest.Marshal(b, m, deterministic)
}
func (m *DeleteClusterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteClusterRequest.Merge(m, src)
}
func (m *DeleteClusterRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteClusterRequest.Size(m)
}
func (m *DeleteClusterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteClusterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteClusterRequest proto.InternalMessageInfo

type UpdateClusterRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateClusterRequest) Reset()         { *m = UpdateClusterRequest{} }
func (m *UpdateClusterRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateClusterRequest) ProtoMessage()    {}
func (*UpdateClusterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_963a9d505aa9dcea, []int{6}
}

func (m *UpdateClusterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateClusterRequest.Unmarshal(m, b)
}
func (m *UpdateClusterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateClusterRequest.Marshal(b, m, deterministic)
}
func (m *UpdateClusterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateClusterRequest.Merge(m, src)
}
func (m *UpdateClusterRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateClusterRequest.Size(m)
}
func (m *UpdateClusterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateClusterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateClusterRequest proto.InternalMessageInfo

type Cluster struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Cluster) Reset()         { *m = Cluster{} }
func (m *Cluster) String() string { return proto.CompactTextString(m) }
func (*Cluster) ProtoMessage()    {}
func (*Cluster) Descriptor() ([]byte, []int) {
	return fileDescriptor_963a9d505aa9dcea, []int{7}
}

func (m *Cluster) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cluster.Unmarshal(m, b)
}
func (m *Cluster) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cluster.Marshal(b, m, deterministic)
}
func (m *Cluster) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cluster.Merge(m, src)
}
func (m *Cluster) XXX_Size() int {
	return xxx_messageInfo_Cluster.Size(m)
}
func (m *Cluster) XXX_DiscardUnknown() {
	xxx_messageInfo_Cluster.DiscardUnknown(m)
}

var xxx_messageInfo_Cluster proto.InternalMessageInfo

func (m *Cluster) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateClusterRequest)(nil), "kaf.api.CreateClusterRequest")
	proto.RegisterType((*CreateClusterResponse)(nil), "kaf.api.CreateClusterResponse")
	proto.RegisterType((*ListClustersRequest)(nil), "kaf.api.ListClustersRequest")
	proto.RegisterType((*ListClustersResponse)(nil), "kaf.api.ListClustersResponse")
	proto.RegisterType((*GetClusterRequest)(nil), "kaf.api.GetClusterRequest")
	proto.RegisterType((*DeleteClusterRequest)(nil), "kaf.api.DeleteClusterRequest")
	proto.RegisterType((*UpdateClusterRequest)(nil), "kaf.api.UpdateClusterRequest")
	proto.RegisterType((*Cluster)(nil), "kaf.api.Cluster")
}

func init() { proto.RegisterFile("api/cluster.proto", fileDescriptor_963a9d505aa9dcea) }

var fileDescriptor_963a9d505aa9dcea = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x41, 0x4f, 0xc2, 0x40,
	0x10, 0x85, 0x53, 0x51, 0xd1, 0xd1, 0x1a, 0x59, 0x4a, 0x25, 0xd5, 0x26, 0xa4, 0x27, 0x0e, 0x66,
	0x9b, 0xe0, 0xcd, 0x13, 0x11, 0xd4, 0x83, 0x9e, 0x30, 0x5e, 0xbc, 0x2d, 0x38, 0x90, 0xc6, 0xc2,
	0xae, 0xdd, 0xad, 0x89, 0x3f, 0xd8, 0xff, 0x61, 0x6c, 0xb7, 0xa5, 0xa5, 0xeb, 0xad, 0x79, 0xf3,
	0x76, 0xfa, 0xbd, 0x99, 0x81, 0x0e, 0x13, 0x51, 0xb8, 0x88, 0x53, 0xa9, 0x30, 0xa1, 0x22, 0xe1,
	0x8a, 0x93, 0xf6, 0x07, 0x5b, 0x52, 0x26, 0x22, 0xef, 0x72, 0xc5, 0xf9, 0x2a, 0xc6, 0x30, 0x93,
	0xe7, 0xe9, 0x32, 0xc4, 0xb5, 0x50, 0xdf, 0xb9, 0x2b, 0x70, 0xc1, 0x99, 0x24, 0xc8, 0x14, 0x4e,
	0xf2, 0xc7, 0x33, 0xfc, 0x4c, 0x51, 0xaa, 0xe0, 0x02, 0x7a, 0x3b, 0xba, 0x14, 0x7c, 0x23, 0x31,
	0xe8, 0x41, 0xf7, 0x39, 0x92, 0x4a, 0xcb, 0xb2, 0xf0, 0x4f, 0xc1, 0xa9, 0xcb, 0xb9, 0x9d, 0x5c,
	0xc3, 0x91, 0xc6, 0x92, 0x7d, 0x6b, 0xd0, 0x1a, 0x9e, 0x8c, 0xce, 0xa9, 0x06, 0xa3, 0x45, 0xeb,
	0xd2, 0x11, 0x74, 0xa1, 0xf3, 0x88, 0x6a, 0x07, 0xc5, 0x05, 0x67, 0x8a, 0x31, 0x36, 0x10, 0x5d,
	0x70, 0x5e, 0xc5, 0x7b, 0x13, 0xdd, 0x87, 0xb6, 0x56, 0x08, 0x81, 0xfd, 0x0d, 0x5b, 0x63, 0xdf,
	0x1a, 0x58, 0xc3, 0xe3, 0x59, 0xf6, 0x3d, 0xfa, 0xd9, 0x83, 0x33, 0x5d, 0x7f, 0xc1, 0xe4, 0x2b,
	0x5a, 0x20, 0x19, 0x83, 0x5d, 0x0b, 0x4b, 0xfc, 0x2d, 0xa3, 0x61, 0x38, 0x5e, 0x23, 0x02, 0xb9,
	0x05, 0xd8, 0x82, 0x13, 0xaf, 0xac, 0x37, 0xd2, 0x18, 0xde, 0x8e, 0xc1, 0xae, 0xe5, 0xa8, 0xfc,
	0xdd, 0x94, 0xcf, 0xd0, 0xe1, 0x09, 0x4e, 0xab, 0xc3, 0x27, 0x57, 0xa5, 0xc3, 0xb0, 0x2a, 0xcf,
	0xff, 0xa7, 0xaa, 0x37, 0xf6, 0x00, 0x76, 0x6d, 0xdc, 0x15, 0x1c, 0xd3, 0x1a, 0x3c, 0x97, 0xe6,
	0xf7, 0x45, 0x8b, 0xfb, 0xa2, 0xf7, 0x7f, 0xf7, 0x75, 0x77, 0xf0, 0xd6, 0x62, 0x22, 0x9a, 0x1f,
	0x66, 0xf2, 0xcd, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xcd, 0xa1, 0x92, 0xb4, 0xa2, 0x02, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ClusterServiceClient is the client API for ClusterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ClusterServiceClient interface {
	CreateCluster(ctx context.Context, in *CreateClusterRequest, opts ...grpc.CallOption) (*Cluster, error)
	GetCluster(ctx context.Context, in *GetClusterRequest, opts ...grpc.CallOption) (*Cluster, error)
	UpdateCluster(ctx context.Context, in *UpdateClusterRequest, opts ...grpc.CallOption) (*Cluster, error)
	ListClusters(ctx context.Context, in *ListClustersRequest, opts ...grpc.CallOption) (*ListClustersResponse, error)
	DeleteCluster(ctx context.Context, in *DeleteClusterRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type clusterServiceClient struct {
	cc *grpc.ClientConn
}

func NewClusterServiceClient(cc *grpc.ClientConn) ClusterServiceClient {
	return &clusterServiceClient{cc}
}

func (c *clusterServiceClient) CreateCluster(ctx context.Context, in *CreateClusterRequest, opts ...grpc.CallOption) (*Cluster, error) {
	out := new(Cluster)
	err := c.cc.Invoke(ctx, "/kaf.api.ClusterService/CreateCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) GetCluster(ctx context.Context, in *GetClusterRequest, opts ...grpc.CallOption) (*Cluster, error) {
	out := new(Cluster)
	err := c.cc.Invoke(ctx, "/kaf.api.ClusterService/GetCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) UpdateCluster(ctx context.Context, in *UpdateClusterRequest, opts ...grpc.CallOption) (*Cluster, error) {
	out := new(Cluster)
	err := c.cc.Invoke(ctx, "/kaf.api.ClusterService/UpdateCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) ListClusters(ctx context.Context, in *ListClustersRequest, opts ...grpc.CallOption) (*ListClustersResponse, error) {
	out := new(ListClustersResponse)
	err := c.cc.Invoke(ctx, "/kaf.api.ClusterService/ListClusters", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) DeleteCluster(ctx context.Context, in *DeleteClusterRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/kaf.api.ClusterService/DeleteCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClusterServiceServer is the server API for ClusterService service.
type ClusterServiceServer interface {
	CreateCluster(context.Context, *CreateClusterRequest) (*Cluster, error)
	GetCluster(context.Context, *GetClusterRequest) (*Cluster, error)
	UpdateCluster(context.Context, *UpdateClusterRequest) (*Cluster, error)
	ListClusters(context.Context, *ListClustersRequest) (*ListClustersResponse, error)
	DeleteCluster(context.Context, *DeleteClusterRequest) (*empty.Empty, error)
}

// UnimplementedClusterServiceServer can be embedded to have forward compatible implementations.
type UnimplementedClusterServiceServer struct {
}

func (*UnimplementedClusterServiceServer) CreateCluster(ctx context.Context, req *CreateClusterRequest) (*Cluster, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCluster not implemented")
}
func (*UnimplementedClusterServiceServer) GetCluster(ctx context.Context, req *GetClusterRequest) (*Cluster, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCluster not implemented")
}
func (*UnimplementedClusterServiceServer) UpdateCluster(ctx context.Context, req *UpdateClusterRequest) (*Cluster, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCluster not implemented")
}
func (*UnimplementedClusterServiceServer) ListClusters(ctx context.Context, req *ListClustersRequest) (*ListClustersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListClusters not implemented")
}
func (*UnimplementedClusterServiceServer) DeleteCluster(ctx context.Context, req *DeleteClusterRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCluster not implemented")
}

func RegisterClusterServiceServer(s *grpc.Server, srv ClusterServiceServer) {
	s.RegisterService(&_ClusterService_serviceDesc, srv)
}

func _ClusterService_CreateCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).CreateCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kaf.api.ClusterService/CreateCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).CreateCluster(ctx, req.(*CreateClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_GetCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).GetCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kaf.api.ClusterService/GetCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).GetCluster(ctx, req.(*GetClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_UpdateCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).UpdateCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kaf.api.ClusterService/UpdateCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).UpdateCluster(ctx, req.(*UpdateClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_ListClusters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListClustersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).ListClusters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kaf.api.ClusterService/ListClusters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).ListClusters(ctx, req.(*ListClustersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_DeleteCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).DeleteCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kaf.api.ClusterService/DeleteCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).DeleteCluster(ctx, req.(*DeleteClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ClusterService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kaf.api.ClusterService",
	HandlerType: (*ClusterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCluster",
			Handler:    _ClusterService_CreateCluster_Handler,
		},
		{
			MethodName: "GetCluster",
			Handler:    _ClusterService_GetCluster_Handler,
		},
		{
			MethodName: "UpdateCluster",
			Handler:    _ClusterService_UpdateCluster_Handler,
		},
		{
			MethodName: "ListClusters",
			Handler:    _ClusterService_ListClusters_Handler,
		},
		{
			MethodName: "DeleteCluster",
			Handler:    _ClusterService_DeleteCluster_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/cluster.proto",
}
