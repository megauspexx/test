// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.3
// source: company.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Req struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Req) Reset() {
	*x = Req{}
	if protoimpl.UnsafeEnabled {
		mi := &file_company_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Req) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Req) ProtoMessage() {}

func (x *Req) ProtoReflect() protoreflect.Message {
	mi := &file_company_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Req.ProtoReflect.Descriptor instead.
func (*Req) Descriptor() ([]byte, []int) {
	return file_company_proto_rawDescGZIP(), []int{0}
}

func (x *Req) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type Resp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response string `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *Resp) Reset() {
	*x = Resp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_company_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resp) ProtoMessage() {}

func (x *Resp) ProtoReflect() protoreflect.Message {
	mi := &file_company_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resp.ProtoReflect.Descriptor instead.
func (*Resp) Descriptor() ([]byte, []int) {
	return file_company_proto_rawDescGZIP(), []int{1}
}

func (x *Resp) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

var File_company_proto protoreflect.FileDescriptor

var file_company_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x19, 0x0a, 0x03, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x22, 0x0a, 0x04, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xe7, 0x01, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x12, 0x27, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12,
	0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x1a, 0x0b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x27, 0x0a, 0x0a, 0x47, 0x65,
	0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x52, 0x65, 0x71, 0x1a, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x57, 0x69, 0x74, 0x68, 0x46, 0x6f, 0x72, 0x6d, 0x12, 0x0a, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x1a, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x2a, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x52, 0x65, 0x71, 0x1a, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x22, 0x00, 0x12, 0x2a, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x12, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x71,
	0x1a, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_company_proto_rawDescOnce sync.Once
	file_company_proto_rawDescData = file_company_proto_rawDesc
)

func file_company_proto_rawDescGZIP() []byte {
	file_company_proto_rawDescOnce.Do(func() {
		file_company_proto_rawDescData = protoimpl.X.CompressGZIP(file_company_proto_rawDescData)
	})
	return file_company_proto_rawDescData
}

var file_company_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_company_proto_goTypes = []interface{}{
	(*Req)(nil),  // 0: proto.Req
	(*Resp)(nil), // 1: proto.Resp
}
var file_company_proto_depIdxs = []int32{
	0, // 0: proto.company.AddCompany:input_type -> proto.Req
	0, // 1: proto.company.GetCompany:input_type -> proto.Req
	0, // 2: proto.company.UpdateCompanyWithForm:input_type -> proto.Req
	0, // 3: proto.company.UpdateCompany:input_type -> proto.Req
	0, // 4: proto.company.DeleteCompany:input_type -> proto.Req
	1, // 5: proto.company.AddCompany:output_type -> proto.Resp
	1, // 6: proto.company.GetCompany:output_type -> proto.Resp
	1, // 7: proto.company.UpdateCompanyWithForm:output_type -> proto.Resp
	1, // 8: proto.company.UpdateCompany:output_type -> proto.Resp
	1, // 9: proto.company.DeleteCompany:output_type -> proto.Resp
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_company_proto_init() }
func file_company_proto_init() {
	if File_company_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_company_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Req); i {
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
		file_company_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resp); i {
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
			RawDescriptor: file_company_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_company_proto_goTypes,
		DependencyIndexes: file_company_proto_depIdxs,
		MessageInfos:      file_company_proto_msgTypes,
	}.Build()
	File_company_proto = out.File
	file_company_proto_rawDesc = nil
	file_company_proto_goTypes = nil
	file_company_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CompanyClient is the client API for Company service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CompanyClient interface {
	AddCompany(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Resp, error)
	GetCompany(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Resp, error)
	UpdateCompanyWithForm(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Resp, error)
	UpdateCompany(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Resp, error)
	DeleteCompany(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Resp, error)
}

type companyClient struct {
	cc grpc.ClientConnInterface
}

func NewCompanyClient(cc grpc.ClientConnInterface) CompanyClient {
	return &companyClient{cc}
}

func (c *companyClient) AddCompany(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Resp, error) {
	out := new(Resp)
	err := c.cc.Invoke(ctx, "/proto.company/AddCompany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyClient) GetCompany(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Resp, error) {
	out := new(Resp)
	err := c.cc.Invoke(ctx, "/proto.company/GetCompany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyClient) UpdateCompanyWithForm(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Resp, error) {
	out := new(Resp)
	err := c.cc.Invoke(ctx, "/proto.company/UpdateCompanyWithForm", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyClient) UpdateCompany(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Resp, error) {
	out := new(Resp)
	err := c.cc.Invoke(ctx, "/proto.company/UpdateCompany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyClient) DeleteCompany(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Resp, error) {
	out := new(Resp)
	err := c.cc.Invoke(ctx, "/proto.company/DeleteCompany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CompanyServer is the server API for Company service.
type CompanyServer interface {
	AddCompany(context.Context, *Req) (*Resp, error)
	GetCompany(context.Context, *Req) (*Resp, error)
	UpdateCompanyWithForm(context.Context, *Req) (*Resp, error)
	UpdateCompany(context.Context, *Req) (*Resp, error)
	DeleteCompany(context.Context, *Req) (*Resp, error)
}

// UnimplementedCompanyServer can be embedded to have forward compatible implementations.
type UnimplementedCompanyServer struct {
}

func (*UnimplementedCompanyServer) AddCompany(context.Context, *Req) (*Resp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCompany not implemented")
}
func (*UnimplementedCompanyServer) GetCompany(context.Context, *Req) (*Resp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCompany not implemented")
}
func (*UnimplementedCompanyServer) UpdateCompanyWithForm(context.Context, *Req) (*Resp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCompanyWithForm not implemented")
}
func (*UnimplementedCompanyServer) UpdateCompany(context.Context, *Req) (*Resp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCompany not implemented")
}
func (*UnimplementedCompanyServer) DeleteCompany(context.Context, *Req) (*Resp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCompany not implemented")
}

func RegisterCompanyServer(s *grpc.Server, srv CompanyServer) {
	s.RegisterService(&_Company_serviceDesc, srv)
}

func _Company_AddCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServer).AddCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.company/AddCompany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServer).AddCompany(ctx, req.(*Req))
	}
	return interceptor(ctx, in, info, handler)
}

func _Company_GetCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServer).GetCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.company/GetCompany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServer).GetCompany(ctx, req.(*Req))
	}
	return interceptor(ctx, in, info, handler)
}

func _Company_UpdateCompanyWithForm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServer).UpdateCompanyWithForm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.company/UpdateCompanyWithForm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServer).UpdateCompanyWithForm(ctx, req.(*Req))
	}
	return interceptor(ctx, in, info, handler)
}

func _Company_UpdateCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServer).UpdateCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.company/UpdateCompany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServer).UpdateCompany(ctx, req.(*Req))
	}
	return interceptor(ctx, in, info, handler)
}

func _Company_DeleteCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServer).DeleteCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.company/DeleteCompany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServer).DeleteCompany(ctx, req.(*Req))
	}
	return interceptor(ctx, in, info, handler)
}

var _Company_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.company",
	HandlerType: (*CompanyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddCompany",
			Handler:    _Company_AddCompany_Handler,
		},
		{
			MethodName: "GetCompany",
			Handler:    _Company_GetCompany_Handler,
		},
		{
			MethodName: "UpdateCompanyWithForm",
			Handler:    _Company_UpdateCompanyWithForm_Handler,
		},
		{
			MethodName: "UpdateCompany",
			Handler:    _Company_UpdateCompany_Handler,
		},
		{
			MethodName: "DeleteCompany",
			Handler:    _Company_DeleteCompany_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "company.proto",
}