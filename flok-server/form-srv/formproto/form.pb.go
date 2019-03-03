// Code generated by protoc-gen-go. DO NOT EDIT.
// source: form.proto

package formproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ProtocolForm struct {
	ConflictInterestTa       string   `protobuf:"bytes,1,opt,name=conflictInterestTa,proto3" json:"conflictInterestTa,omitempty"`
	CbPublications           bool     `protobuf:"varint,2,opt,name=cbPublications,proto3" json:"cbPublications,omitempty"`
	TaPublications           string   `protobuf:"bytes,3,opt,name=taPublications,proto3" json:"taPublications,omitempty"`
	ProjectInvestigator      string   `protobuf:"bytes,4,opt,name=projectInvestigator,proto3" json:"projectInvestigator,omitempty"`
	ProjectCoInvestigator    []string `protobuf:"bytes,5,rep,name=projectCoInvestigator,proto3" json:"projectCoInvestigator,omitempty"`
	ProjectPeriod            string   `protobuf:"bytes,6,opt,name=projectPeriod,proto3" json:"projectPeriod,omitempty"`
	DnbName                  string   `protobuf:"bytes,7,opt,name=dnbName,proto3" json:"dnbName,omitempty"`
	DnbContact               string   `protobuf:"bytes,8,opt,name=dnbContact,proto3" json:"dnbContact,omitempty"`
	DnbDesignation           string   `protobuf:"bytes,9,opt,name=dnbDesignation,proto3" json:"dnbDesignation,omitempty"`
	DnbEmail                 string   `protobuf:"bytes,10,opt,name=dnbEmail,proto3" json:"dnbEmail,omitempty"`
	TimeDataAnalysis         string   `protobuf:"bytes,11,opt,name=timeDataAnalysis,proto3" json:"timeDataAnalysis,omitempty"`
	TimeIndividualPatient    string   `protobuf:"bytes,12,opt,name=timeIndividualPatient,proto3" json:"timeIndividualPatient,omitempty"`
	TimeProspectiveStudies   string   `protobuf:"bytes,13,opt,name=timeProspectiveStudies,proto3" json:"timeProspectiveStudies,omitempty"`
	TimeRetrospectiveStudies string   `protobuf:"bytes,14,opt,name=timeRetrospectiveStudies,proto3" json:"timeRetrospectiveStudies,omitempty"`
	TimeTotalDuration        string   `protobuf:"bytes,15,opt,name=timeTotalDuration,proto3" json:"timeTotalDuration,omitempty"`
	TimeWriteUps             string   `protobuf:"bytes,16,opt,name=timeWriteUps,proto3" json:"timeWriteUps,omitempty"`
	ProjectCode              string   `protobuf:"bytes,17,opt,name=projectCode,proto3" json:"projectCode,omitempty"`
	ProjectName              string   `protobuf:"bytes,18,opt,name=projectName,proto3" json:"projectName,omitempty"`
	BudgetEstimate           string   `protobuf:"bytes,19,opt,name=budgetEstimate,proto3" json:"budgetEstimate,omitempty"`
	UploadFile               string   `protobuf:"bytes,20,opt,name=uploadFile,proto3" json:"uploadFile,omitempty"`
	ConflictInterest         bool     `protobuf:"varint,21,opt,name=conflictInterest,proto3" json:"conflictInterest,omitempty"`
	XXX_NoUnkeyedLiteral     struct{} `json:"-"`
	XXX_unrecognized         []byte   `json:"-"`
	XXX_sizecache            int32    `json:"-"`
}

func (m *ProtocolForm) Reset()         { *m = ProtocolForm{} }
func (m *ProtocolForm) String() string { return proto.CompactTextString(m) }
func (*ProtocolForm) ProtoMessage()    {}
func (*ProtocolForm) Descriptor() ([]byte, []int) {
	return fileDescriptor_form_210950787b632e91, []int{0}
}
func (m *ProtocolForm) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProtocolForm.Unmarshal(m, b)
}
func (m *ProtocolForm) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProtocolForm.Marshal(b, m, deterministic)
}
func (dst *ProtocolForm) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtocolForm.Merge(dst, src)
}
func (m *ProtocolForm) XXX_Size() int {
	return xxx_messageInfo_ProtocolForm.Size(m)
}
func (m *ProtocolForm) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtocolForm.DiscardUnknown(m)
}

var xxx_messageInfo_ProtocolForm proto.InternalMessageInfo

func (m *ProtocolForm) GetConflictInterestTa() string {
	if m != nil {
		return m.ConflictInterestTa
	}
	return ""
}

func (m *ProtocolForm) GetCbPublications() bool {
	if m != nil {
		return m.CbPublications
	}
	return false
}

func (m *ProtocolForm) GetTaPublications() string {
	if m != nil {
		return m.TaPublications
	}
	return ""
}

func (m *ProtocolForm) GetProjectInvestigator() string {
	if m != nil {
		return m.ProjectInvestigator
	}
	return ""
}

func (m *ProtocolForm) GetProjectCoInvestigator() []string {
	if m != nil {
		return m.ProjectCoInvestigator
	}
	return nil
}

func (m *ProtocolForm) GetProjectPeriod() string {
	if m != nil {
		return m.ProjectPeriod
	}
	return ""
}

func (m *ProtocolForm) GetDnbName() string {
	if m != nil {
		return m.DnbName
	}
	return ""
}

func (m *ProtocolForm) GetDnbContact() string {
	if m != nil {
		return m.DnbContact
	}
	return ""
}

func (m *ProtocolForm) GetDnbDesignation() string {
	if m != nil {
		return m.DnbDesignation
	}
	return ""
}

func (m *ProtocolForm) GetDnbEmail() string {
	if m != nil {
		return m.DnbEmail
	}
	return ""
}

func (m *ProtocolForm) GetTimeDataAnalysis() string {
	if m != nil {
		return m.TimeDataAnalysis
	}
	return ""
}

func (m *ProtocolForm) GetTimeIndividualPatient() string {
	if m != nil {
		return m.TimeIndividualPatient
	}
	return ""
}

func (m *ProtocolForm) GetTimeProspectiveStudies() string {
	if m != nil {
		return m.TimeProspectiveStudies
	}
	return ""
}

func (m *ProtocolForm) GetTimeRetrospectiveStudies() string {
	if m != nil {
		return m.TimeRetrospectiveStudies
	}
	return ""
}

func (m *ProtocolForm) GetTimeTotalDuration() string {
	if m != nil {
		return m.TimeTotalDuration
	}
	return ""
}

func (m *ProtocolForm) GetTimeWriteUps() string {
	if m != nil {
		return m.TimeWriteUps
	}
	return ""
}

func (m *ProtocolForm) GetProjectCode() string {
	if m != nil {
		return m.ProjectCode
	}
	return ""
}

func (m *ProtocolForm) GetProjectName() string {
	if m != nil {
		return m.ProjectName
	}
	return ""
}

func (m *ProtocolForm) GetBudgetEstimate() string {
	if m != nil {
		return m.BudgetEstimate
	}
	return ""
}

func (m *ProtocolForm) GetUploadFile() string {
	if m != nil {
		return m.UploadFile
	}
	return ""
}

func (m *ProtocolForm) GetConflictInterest() bool {
	if m != nil {
		return m.ConflictInterest
	}
	return false
}

type Form struct {
	// @inject_tag: bson:"_id"
	Id                   string        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" bson:"_id"`
	ProtocolForm         *ProtocolForm `protobuf:"bytes,2,opt,name=protocolForm,proto3" json:"protocolForm,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Form) Reset()         { *m = Form{} }
func (m *Form) String() string { return proto.CompactTextString(m) }
func (*Form) ProtoMessage()    {}
func (*Form) Descriptor() ([]byte, []int) {
	return fileDescriptor_form_210950787b632e91, []int{1}
}
func (m *Form) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Form.Unmarshal(m, b)
}
func (m *Form) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Form.Marshal(b, m, deterministic)
}
func (dst *Form) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Form.Merge(dst, src)
}
func (m *Form) XXX_Size() int {
	return xxx_messageInfo_Form.Size(m)
}
func (m *Form) XXX_DiscardUnknown() {
	xxx_messageInfo_Form.DiscardUnknown(m)
}

var xxx_messageInfo_Form proto.InternalMessageInfo

func (m *Form) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Form) GetProtocolForm() *ProtocolForm {
	if m != nil {
		return m.ProtocolForm
	}
	return nil
}

type CreateRequest struct {
	Form                 *Form    `protobuf:"bytes,1,opt,name=form,proto3" json:"form,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_form_210950787b632e91, []int{2}
}
func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (dst *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(dst, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetForm() *Form {
	if m != nil {
		return m.Form
	}
	return nil
}

type CreateResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_form_210950787b632e91, []int{3}
}
func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (dst *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(dst, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ViewRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ViewRequest) Reset()         { *m = ViewRequest{} }
func (m *ViewRequest) String() string { return proto.CompactTextString(m) }
func (*ViewRequest) ProtoMessage()    {}
func (*ViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_form_210950787b632e91, []int{4}
}
func (m *ViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ViewRequest.Unmarshal(m, b)
}
func (m *ViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ViewRequest.Marshal(b, m, deterministic)
}
func (dst *ViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ViewRequest.Merge(dst, src)
}
func (m *ViewRequest) XXX_Size() int {
	return xxx_messageInfo_ViewRequest.Size(m)
}
func (m *ViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ViewRequest proto.InternalMessageInfo

func (m *ViewRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ViewResponse struct {
	Form                 *Form    `protobuf:"bytes,1,opt,name=form,proto3" json:"form,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ViewResponse) Reset()         { *m = ViewResponse{} }
func (m *ViewResponse) String() string { return proto.CompactTextString(m) }
func (*ViewResponse) ProtoMessage()    {}
func (*ViewResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_form_210950787b632e91, []int{5}
}
func (m *ViewResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ViewResponse.Unmarshal(m, b)
}
func (m *ViewResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ViewResponse.Marshal(b, m, deterministic)
}
func (dst *ViewResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ViewResponse.Merge(dst, src)
}
func (m *ViewResponse) XXX_Size() int {
	return xxx_messageInfo_ViewResponse.Size(m)
}
func (m *ViewResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ViewResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ViewResponse proto.InternalMessageInfo

func (m *ViewResponse) GetForm() *Form {
	if m != nil {
		return m.Form
	}
	return nil
}

type UpdateRequest struct {
	Form                 *Form    `protobuf:"bytes,1,opt,name=form,proto3" json:"form,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_form_210950787b632e91, []int{6}
}
func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(dst, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetForm() *Form {
	if m != nil {
		return m.Form
	}
	return nil
}

type UpdateResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateResponse) Reset()         { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()    {}
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_form_210950787b632e91, []int{7}
}
func (m *UpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateResponse.Unmarshal(m, b)
}
func (m *UpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateResponse.Marshal(b, m, deterministic)
}
func (dst *UpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateResponse.Merge(dst, src)
}
func (m *UpdateResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateResponse.Size(m)
}
func (m *UpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ProtocolForm)(nil), "formproto.ProtocolForm")
	proto.RegisterType((*Form)(nil), "formproto.Form")
	proto.RegisterType((*CreateRequest)(nil), "formproto.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "formproto.CreateResponse")
	proto.RegisterType((*ViewRequest)(nil), "formproto.ViewRequest")
	proto.RegisterType((*ViewResponse)(nil), "formproto.ViewResponse")
	proto.RegisterType((*UpdateRequest)(nil), "formproto.UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "formproto.UpdateResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FormServiceClient is the client API for FormService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FormServiceClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	View(ctx context.Context, in *ViewRequest, opts ...grpc.CallOption) (*ViewResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
}

type formServiceClient struct {
	cc *grpc.ClientConn
}

func NewFormServiceClient(cc *grpc.ClientConn) FormServiceClient {
	return &formServiceClient{cc}
}

func (c *formServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/formproto.FormService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *formServiceClient) View(ctx context.Context, in *ViewRequest, opts ...grpc.CallOption) (*ViewResponse, error) {
	out := new(ViewResponse)
	err := c.cc.Invoke(ctx, "/formproto.FormService/View", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *formServiceClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/formproto.FormService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FormServiceServer is the server API for FormService service.
type FormServiceServer interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	View(context.Context, *ViewRequest) (*ViewResponse, error)
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
}

func RegisterFormServiceServer(s *grpc.Server, srv FormServiceServer) {
	s.RegisterService(&_FormService_serviceDesc, srv)
}

func _FormService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FormServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/formproto.FormService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FormServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FormService_View_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FormServiceServer).View(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/formproto.FormService/View",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FormServiceServer).View(ctx, req.(*ViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FormService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FormServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/formproto.FormService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FormServiceServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FormService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "formproto.FormService",
	HandlerType: (*FormServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _FormService_Create_Handler,
		},
		{
			MethodName: "View",
			Handler:    _FormService_View_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _FormService_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "form.proto",
}

func init() { proto.RegisterFile("form.proto", fileDescriptor_form_210950787b632e91) }

var fileDescriptor_form_210950787b632e91 = []byte{
	// 612 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x5d, 0x6b, 0x13, 0x4d,
	0x14, 0x6e, 0xda, 0xbc, 0x69, 0x73, 0x36, 0x49, 0xdb, 0xe9, 0xdb, 0x76, 0x2c, 0x28, 0x61, 0x95,
	0x52, 0x44, 0x82, 0xb4, 0x45, 0x50, 0x2f, 0x8a, 0xf6, 0x03, 0x7a, 0x23, 0x21, 0x6d, 0x15, 0xbc,
	0x9b, 0xdd, 0x39, 0x0d, 0x23, 0x9b, 0x99, 0x75, 0x66, 0x36, 0xe2, 0x5f, 0xf4, 0x27, 0x79, 0x25,
	0x33, 0x9b, 0xa4, 0x3b, 0x49, 0x0a, 0x7a, 0x95, 0xd9, 0xe7, 0xe3, 0x9c, 0x39, 0x73, 0x78, 0x02,
	0x70, 0xaf, 0xf4, 0xa8, 0x97, 0x6b, 0x65, 0x15, 0x69, 0xba, 0xb3, 0x3f, 0xc6, 0xbf, 0x1b, 0xd0,
	0xea, 0xbb, 0x53, 0xaa, 0xb2, 0x2b, 0xa5, 0x47, 0xa4, 0x07, 0x24, 0x55, 0xf2, 0x3e, 0x13, 0xa9,
	0xbd, 0x96, 0x16, 0x35, 0x1a, 0x7b, 0xcb, 0x68, 0xad, 0x5b, 0x3b, 0x6a, 0x0e, 0x96, 0x30, 0xe4,
	0x10, 0x3a, 0x69, 0xd2, 0x2f, 0x92, 0x4c, 0xa4, 0xcc, 0x0a, 0x25, 0x0d, 0x5d, 0xed, 0xd6, 0x8e,
	0x36, 0x06, 0x73, 0xa8, 0xd3, 0x59, 0x16, 0xe8, 0xd6, 0x7c, 0xcd, 0x39, 0x94, 0xbc, 0x86, 0x9d,
	0x5c, 0xab, 0x6f, 0xe8, 0x9a, 0x8c, 0xd1, 0x58, 0x31, 0x64, 0x56, 0x69, 0x5a, 0xf7, 0xe2, 0x65,
	0x14, 0x39, 0x85, 0xdd, 0x09, 0x7c, 0xae, 0x02, 0xcf, 0x7f, 0xdd, 0xb5, 0xa3, 0xe6, 0x60, 0x39,
	0x49, 0x5e, 0x40, 0x7b, 0x42, 0xf4, 0x51, 0x0b, 0xc5, 0x69, 0xc3, 0x77, 0x08, 0x41, 0x42, 0x61,
	0x9d, 0xcb, 0xe4, 0x13, 0x1b, 0x21, 0x5d, 0xf7, 0xfc, 0xf4, 0x93, 0x3c, 0x03, 0xe0, 0x32, 0x39,
	0x57, 0xd2, 0xb2, 0xd4, 0xd2, 0x0d, 0x4f, 0x56, 0x10, 0x37, 0x2f, 0x97, 0xc9, 0x05, 0x1a, 0x31,
	0x94, 0x7e, 0x34, 0xda, 0x2c, 0xe7, 0x0d, 0x51, 0x72, 0x00, 0x1b, 0x5c, 0x26, 0x97, 0x23, 0x26,
	0x32, 0x0a, 0x5e, 0x31, 0xfb, 0x26, 0x2f, 0x61, 0xcb, 0x8a, 0x11, 0x5e, 0x30, 0xcb, 0x3e, 0x48,
	0x96, 0xfd, 0x34, 0xc2, 0xd0, 0xc8, 0x6b, 0x16, 0x70, 0xf7, 0x0a, 0x0e, 0xbb, 0x96, 0x5c, 0x8c,
	0x05, 0x2f, 0x58, 0xd6, 0x67, 0x56, 0xa0, 0xb4, 0xb4, 0xe5, 0x0d, 0xcb, 0x49, 0xf2, 0x06, 0xf6,
	0x1c, 0xd1, 0xd7, 0xca, 0xe4, 0x98, 0x5a, 0x31, 0xc6, 0x1b, 0x5b, 0x70, 0x81, 0x86, 0xb6, 0xbd,
	0xed, 0x11, 0x96, 0xbc, 0x03, 0xea, 0x98, 0x01, 0xda, 0x45, 0x67, 0xc7, 0x3b, 0x1f, 0xe5, 0xc9,
	0x2b, 0xd8, 0x76, 0xdc, 0xad, 0xb2, 0x2c, 0xbb, 0x28, 0x74, 0xf9, 0x38, 0x9b, 0xde, 0xb4, 0x48,
	0x90, 0x18, 0x5a, 0x0e, 0xfc, 0xa2, 0x85, 0xc5, 0xbb, 0xdc, 0xd0, 0x2d, 0x2f, 0x0c, 0x30, 0xd2,
	0x85, 0x68, 0xb6, 0x64, 0x8e, 0x74, 0xdb, 0x4b, 0xaa, 0x50, 0x45, 0xe1, 0x77, 0x49, 0x02, 0x85,
	0xdf, 0xe7, 0x21, 0x74, 0x92, 0x82, 0x0f, 0xd1, 0x5e, 0x1a, 0x2b, 0x46, 0xcc, 0x22, 0xdd, 0x29,
	0xf7, 0x15, 0xa2, 0x6e, 0xef, 0x45, 0x9e, 0x29, 0xc6, 0xaf, 0x44, 0x86, 0xf4, 0xff, 0x72, 0xef,
	0x0f, 0x88, 0xdb, 0xd9, 0x7c, 0x4a, 0xe8, 0xae, 0x4f, 0xc4, 0x02, 0x1e, 0xdf, 0x40, 0xdd, 0x67,
	0xae, 0x03, 0xab, 0x82, 0x4f, 0x32, 0xb6, 0x2a, 0x38, 0x79, 0x0f, 0xad, 0xbc, 0x92, 0x49, 0x9f,
	0xa8, 0xe8, 0x78, 0xbf, 0x37, 0x8b, 0x6d, 0xaf, 0x1a, 0xd9, 0x41, 0x20, 0x8e, 0x4f, 0xa1, 0x7d,
	0xae, 0x91, 0x59, 0x1c, 0xe0, 0xf7, 0x02, 0x8d, 0x25, 0xcf, 0xa1, 0xee, 0x8c, 0xbe, 0x7e, 0x74,
	0xbc, 0x59, 0xa9, 0xe2, 0xdd, 0x9e, 0x8c, 0xbb, 0xd0, 0x99, 0xba, 0x4c, 0xae, 0xa4, 0xc1, 0xf9,
	0x4b, 0xc5, 0x4f, 0x21, 0xfa, 0x2c, 0xf0, 0xc7, 0xb4, 0xea, 0x3c, 0x7d, 0x02, 0xad, 0x92, 0x9e,
	0xd8, 0xff, 0xaa, 0xeb, 0x29, 0xb4, 0xef, 0x72, 0xfe, 0xaf, 0x77, 0xdd, 0x82, 0xce, 0xd4, 0x55,
	0x36, 0x3b, 0xfe, 0x55, 0x83, 0xc8, 0x09, 0x6e, 0x50, 0x8f, 0x45, 0x8a, 0xe4, 0x0c, 0x1a, 0xe5,
	0x34, 0x84, 0x56, 0x4a, 0x04, 0xcf, 0x72, 0xf0, 0x64, 0x09, 0x53, 0x96, 0x8b, 0x57, 0xc8, 0x5b,
	0xa8, 0xbb, 0x69, 0xc8, 0x5e, 0x45, 0x54, 0x99, 0xfe, 0x60, 0x7f, 0x01, 0x9f, 0x59, 0xcf, 0xa0,
	0x51, 0xde, 0x2e, 0xe8, 0x1d, 0x8c, 0x19, 0xf4, 0x0e, 0x47, 0x89, 0x57, 0x3e, 0x46, 0x5f, 0x1f,
	0xfe, 0x9f, 0x93, 0x86, 0xff, 0x39, 0xf9, 0x13, 0x00, 0x00, 0xff, 0xff, 0x31, 0x9d, 0xe1, 0x96,
	0xbf, 0x05, 0x00, 0x00,
}