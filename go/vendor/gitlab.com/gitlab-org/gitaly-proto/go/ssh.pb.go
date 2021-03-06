// Code generated by protoc-gen-go.
// source: ssh.proto
// DO NOT EDIT!

package gitaly

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

type SSHUploadPackRequest struct {
	// 'repository' must be present in the first message.
	Repository *Repository `protobuf:"bytes,1,opt,name=repository" json:"repository,omitempty"`
	// A chunk of raw data to be copied to 'git upload-pack' standard input
	Stdin []byte `protobuf:"bytes,2,opt,name=stdin,proto3" json:"stdin,omitempty"`
}

func (m *SSHUploadPackRequest) Reset()                    { *m = SSHUploadPackRequest{} }
func (m *SSHUploadPackRequest) String() string            { return proto.CompactTextString(m) }
func (*SSHUploadPackRequest) ProtoMessage()               {}
func (*SSHUploadPackRequest) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

func (m *SSHUploadPackRequest) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *SSHUploadPackRequest) GetStdin() []byte {
	if m != nil {
		return m.Stdin
	}
	return nil
}

type SSHUploadPackResponse struct {
	// A chunk of raw data from 'git upload-pack' standard output
	Stdout []byte `protobuf:"bytes,1,opt,name=stdout,proto3" json:"stdout,omitempty"`
	// A chunk of raw data from 'git upload-pack' standard error
	Stderr []byte `protobuf:"bytes,2,opt,name=stderr,proto3" json:"stderr,omitempty"`
	// This field may be nil. This is intentional: only when the remote
	// command has finished can we return its exit status.
	ExitStatus *ExitStatus `protobuf:"bytes,3,opt,name=exit_status,json=exitStatus" json:"exit_status,omitempty"`
}

func (m *SSHUploadPackResponse) Reset()                    { *m = SSHUploadPackResponse{} }
func (m *SSHUploadPackResponse) String() string            { return proto.CompactTextString(m) }
func (*SSHUploadPackResponse) ProtoMessage()               {}
func (*SSHUploadPackResponse) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{1} }

func (m *SSHUploadPackResponse) GetStdout() []byte {
	if m != nil {
		return m.Stdout
	}
	return nil
}

func (m *SSHUploadPackResponse) GetStderr() []byte {
	if m != nil {
		return m.Stderr
	}
	return nil
}

func (m *SSHUploadPackResponse) GetExitStatus() *ExitStatus {
	if m != nil {
		return m.ExitStatus
	}
	return nil
}

type SSHReceivePackRequest struct {
	// 'repository' must be present in the first message.
	Repository *Repository `protobuf:"bytes,1,opt,name=repository" json:"repository,omitempty"`
	// A chunk of raw data to be copied to 'git upload-pack' standard input
	Stdin []byte `protobuf:"bytes,2,opt,name=stdin,proto3" json:"stdin,omitempty"`
	// Contents of GL_ID environment variable for 'git receive-pack'
	GlId string `protobuf:"bytes,3,opt,name=gl_id,json=glId" json:"gl_id,omitempty"`
}

func (m *SSHReceivePackRequest) Reset()                    { *m = SSHReceivePackRequest{} }
func (m *SSHReceivePackRequest) String() string            { return proto.CompactTextString(m) }
func (*SSHReceivePackRequest) ProtoMessage()               {}
func (*SSHReceivePackRequest) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{2} }

func (m *SSHReceivePackRequest) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *SSHReceivePackRequest) GetStdin() []byte {
	if m != nil {
		return m.Stdin
	}
	return nil
}

func (m *SSHReceivePackRequest) GetGlId() string {
	if m != nil {
		return m.GlId
	}
	return ""
}

type SSHReceivePackResponse struct {
	// A chunk of raw data from 'git receive-pack' standard output
	Stdout []byte `protobuf:"bytes,1,opt,name=stdout,proto3" json:"stdout,omitempty"`
	// A chunk of raw data from 'git receive-pack' standard error
	Stderr []byte `protobuf:"bytes,2,opt,name=stderr,proto3" json:"stderr,omitempty"`
	// This field may be nil. This is intentional: only when the remote
	// command has finished can we return its exit status.
	ExitStatus *ExitStatus `protobuf:"bytes,3,opt,name=exit_status,json=exitStatus" json:"exit_status,omitempty"`
}

func (m *SSHReceivePackResponse) Reset()                    { *m = SSHReceivePackResponse{} }
func (m *SSHReceivePackResponse) String() string            { return proto.CompactTextString(m) }
func (*SSHReceivePackResponse) ProtoMessage()               {}
func (*SSHReceivePackResponse) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{3} }

func (m *SSHReceivePackResponse) GetStdout() []byte {
	if m != nil {
		return m.Stdout
	}
	return nil
}

func (m *SSHReceivePackResponse) GetStderr() []byte {
	if m != nil {
		return m.Stderr
	}
	return nil
}

func (m *SSHReceivePackResponse) GetExitStatus() *ExitStatus {
	if m != nil {
		return m.ExitStatus
	}
	return nil
}

func init() {
	proto.RegisterType((*SSHUploadPackRequest)(nil), "gitaly.SSHUploadPackRequest")
	proto.RegisterType((*SSHUploadPackResponse)(nil), "gitaly.SSHUploadPackResponse")
	proto.RegisterType((*SSHReceivePackRequest)(nil), "gitaly.SSHReceivePackRequest")
	proto.RegisterType((*SSHReceivePackResponse)(nil), "gitaly.SSHReceivePackResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for SSH service

type SSHClient interface {
	// To forward 'git upload-pack' to Gitaly for SSH sessions
	SSHUploadPack(ctx context.Context, opts ...grpc.CallOption) (SSH_SSHUploadPackClient, error)
	// To forward 'git receive-pack' to Gitaly for SSH sessions
	SSHReceivePack(ctx context.Context, opts ...grpc.CallOption) (SSH_SSHReceivePackClient, error)
}

type sSHClient struct {
	cc *grpc.ClientConn
}

func NewSSHClient(cc *grpc.ClientConn) SSHClient {
	return &sSHClient{cc}
}

func (c *sSHClient) SSHUploadPack(ctx context.Context, opts ...grpc.CallOption) (SSH_SSHUploadPackClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_SSH_serviceDesc.Streams[0], c.cc, "/gitaly.SSH/SSHUploadPack", opts...)
	if err != nil {
		return nil, err
	}
	x := &sSHSSHUploadPackClient{stream}
	return x, nil
}

type SSH_SSHUploadPackClient interface {
	Send(*SSHUploadPackRequest) error
	Recv() (*SSHUploadPackResponse, error)
	grpc.ClientStream
}

type sSHSSHUploadPackClient struct {
	grpc.ClientStream
}

func (x *sSHSSHUploadPackClient) Send(m *SSHUploadPackRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sSHSSHUploadPackClient) Recv() (*SSHUploadPackResponse, error) {
	m := new(SSHUploadPackResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sSHClient) SSHReceivePack(ctx context.Context, opts ...grpc.CallOption) (SSH_SSHReceivePackClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_SSH_serviceDesc.Streams[1], c.cc, "/gitaly.SSH/SSHReceivePack", opts...)
	if err != nil {
		return nil, err
	}
	x := &sSHSSHReceivePackClient{stream}
	return x, nil
}

type SSH_SSHReceivePackClient interface {
	Send(*SSHReceivePackRequest) error
	Recv() (*SSHReceivePackResponse, error)
	grpc.ClientStream
}

type sSHSSHReceivePackClient struct {
	grpc.ClientStream
}

func (x *sSHSSHReceivePackClient) Send(m *SSHReceivePackRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sSHSSHReceivePackClient) Recv() (*SSHReceivePackResponse, error) {
	m := new(SSHReceivePackResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for SSH service

type SSHServer interface {
	// To forward 'git upload-pack' to Gitaly for SSH sessions
	SSHUploadPack(SSH_SSHUploadPackServer) error
	// To forward 'git receive-pack' to Gitaly for SSH sessions
	SSHReceivePack(SSH_SSHReceivePackServer) error
}

func RegisterSSHServer(s *grpc.Server, srv SSHServer) {
	s.RegisterService(&_SSH_serviceDesc, srv)
}

func _SSH_SSHUploadPack_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SSHServer).SSHUploadPack(&sSHSSHUploadPackServer{stream})
}

type SSH_SSHUploadPackServer interface {
	Send(*SSHUploadPackResponse) error
	Recv() (*SSHUploadPackRequest, error)
	grpc.ServerStream
}

type sSHSSHUploadPackServer struct {
	grpc.ServerStream
}

func (x *sSHSSHUploadPackServer) Send(m *SSHUploadPackResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sSHSSHUploadPackServer) Recv() (*SSHUploadPackRequest, error) {
	m := new(SSHUploadPackRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SSH_SSHReceivePack_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SSHServer).SSHReceivePack(&sSHSSHReceivePackServer{stream})
}

type SSH_SSHReceivePackServer interface {
	Send(*SSHReceivePackResponse) error
	Recv() (*SSHReceivePackRequest, error)
	grpc.ServerStream
}

type sSHSSHReceivePackServer struct {
	grpc.ServerStream
}

func (x *sSHSSHReceivePackServer) Send(m *SSHReceivePackResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sSHSSHReceivePackServer) Recv() (*SSHReceivePackRequest, error) {
	m := new(SSHReceivePackRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _SSH_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gitaly.SSH",
	HandlerType: (*SSHServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SSHUploadPack",
			Handler:       _SSH_SSHUploadPack_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "SSHReceivePack",
			Handler:       _SSH_SSHReceivePack_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "ssh.proto",
}

func init() { proto.RegisterFile("ssh.proto", fileDescriptor6) }

var fileDescriptor6 = []byte{
	// 284 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x52, 0xcd, 0x4a, 0xc3, 0x40,
	0x10, 0x76, 0xad, 0x0d, 0x74, 0x1a, 0x3d, 0xac, 0xb5, 0x84, 0xa0, 0x52, 0x72, 0xca, 0x29, 0x48,
	0xfa, 0x0c, 0x42, 0xbd, 0xc9, 0x86, 0x9e, 0x6b, 0xec, 0x2e, 0xe9, 0x62, 0xe8, 0xc6, 0x9d, 0x49,
	0x69, 0x41, 0xdf, 0xc9, 0x47, 0x14, 0x92, 0x58, 0x92, 0x6a, 0x8f, 0xf6, 0xb6, 0x33, 0xdf, 0xce,
	0xf7, 0x33, 0xbb, 0x30, 0x40, 0x5c, 0x45, 0x85, 0x35, 0x64, 0xb8, 0x93, 0x69, 0x4a, 0xf3, 0x9d,
	0xef, 0xe2, 0x2a, 0xb5, 0x4a, 0xd6, 0xdd, 0xe0, 0x05, 0x46, 0x49, 0x32, 0x9b, 0x17, 0xb9, 0x49,
	0xe5, 0x73, 0xba, 0x7c, 0x13, 0xea, 0xbd, 0x54, 0x48, 0x3c, 0x06, 0xb0, 0xaa, 0x30, 0xa8, 0xc9,
	0xd8, 0x9d, 0xc7, 0x26, 0x2c, 0x1c, 0xc6, 0x3c, 0xaa, 0x29, 0x22, 0xb1, 0x47, 0x44, 0xeb, 0x16,
	0x1f, 0x41, 0x1f, 0x49, 0xea, 0xb5, 0x77, 0x3e, 0x61, 0xa1, 0x2b, 0xea, 0x22, 0xf8, 0x80, 0x9b,
	0x03, 0x05, 0x2c, 0xcc, 0x1a, 0x15, 0x1f, 0x83, 0x83, 0x24, 0x4d, 0x49, 0x15, 0xbd, 0x2b, 0x9a,
	0xaa, 0xe9, 0x2b, 0x6b, 0x1b, 0x9e, 0xa6, 0xe2, 0x53, 0x18, 0xaa, 0xad, 0xa6, 0x05, 0x52, 0x4a,
	0x25, 0x7a, 0xbd, 0xae, 0xa7, 0xc7, 0xad, 0xa6, 0xa4, 0x42, 0x04, 0xa8, 0xfd, 0x39, 0xd8, 0x54,
	0xea, 0x42, 0x2d, 0x95, 0xde, 0xa8, 0x7f, 0x09, 0xc8, 0xaf, 0xa1, 0x9f, 0xe5, 0x0b, 0x2d, 0x2b,
	0x47, 0x03, 0x71, 0x91, 0xe5, 0x4f, 0x32, 0xf8, 0x84, 0xf1, 0xa1, 0xee, 0x09, 0x63, 0xc7, 0x5f,
	0x0c, 0x7a, 0x49, 0x32, 0xe3, 0x02, 0x2e, 0x3b, 0xcb, 0xe7, 0xb7, 0x3f, 0x83, 0x7f, 0xbd, 0xba,
	0x7f, 0x77, 0x04, 0xad, 0xad, 0x07, 0x67, 0x21, 0x7b, 0x60, 0x7c, 0x0e, 0x57, 0xdd, 0x68, 0xbc,
	0x3d, 0xf6, 0x7b, 0xd5, 0xfe, 0xfd, 0x31, 0xb8, 0x4d, 0xfb, 0xea, 0x54, 0x1f, 0x72, 0xfa, 0x1d,
	0x00, 0x00, 0xff, 0xff, 0x71, 0xde, 0xae, 0x4f, 0xb3, 0x02, 0x00, 0x00,
}
