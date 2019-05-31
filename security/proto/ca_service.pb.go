// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: security/proto/ca_service.proto

/*
	Package istio_v1_auth is a generated protocol buffer package.

	It is generated from these files:
		security/proto/ca_service.proto

	It has these top-level messages:
		CsrRequest
		CsrResponse
*/
package istio_v1_auth

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import google_rpc "github.com/gogo/googleapis/google/rpc"
import _ "github.com/gogo/protobuf/gogoproto"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type CsrRequest struct {
	// PEM-encoded certificate signing request
	CsrPem []byte `protobuf:"bytes,1,opt,name=csr_pem,json=csrPem,proto3" json:"csr_pem,omitempty"`
	// opaque credential for node agent
	NodeAgentCredential []byte `protobuf:"bytes,2,opt,name=node_agent_credential,json=nodeAgentCredential,proto3" json:"node_agent_credential,omitempty"`
	// type of the node_agent_credential (aws/gcp/onprem/custom...)
	CredentialType string `protobuf:"bytes,3,opt,name=credential_type,json=credentialType,proto3" json:"credential_type,omitempty"`
	// the requested ttl of the certificate in minutes
	RequestedTtlMinutes int32 `protobuf:"varint,4,opt,name=requested_ttl_minutes,json=requestedTtlMinutes,proto3" json:"requested_ttl_minutes,omitempty"`
}

func (m *CsrRequest) Reset()                    { *m = CsrRequest{} }
func (*CsrRequest) ProtoMessage()               {}
func (*CsrRequest) Descriptor() ([]byte, []int) { return fileDescriptorCaService, []int{0} }

type CsrResponse struct {
	// Whether the CSR is approved.
	IsApproved bool               `protobuf:"varint,1,opt,name=is_approved,json=isApproved,proto3" json:"is_approved,omitempty"`
	Status     *google_rpc.Status `protobuf:"bytes,2,opt,name=status" json:"status,omitempty"`
	// The signed target cert.
	SignedCert []byte `protobuf:"bytes,3,opt,name=signed_cert,json=signedCert,proto3" json:"signed_cert,omitempty"`
	// The cert chain up to the trusted root cert. It includes all the certs between the
	// newly signed cert and the root cert.
	CertChain []byte `protobuf:"bytes,4,opt,name=cert_chain,json=certChain,proto3" json:"cert_chain,omitempty"`
}

func (m *CsrResponse) Reset()                    { *m = CsrResponse{} }
func (*CsrResponse) ProtoMessage()               {}
func (*CsrResponse) Descriptor() ([]byte, []int) { return fileDescriptorCaService, []int{1} }

func init() {
	proto.RegisterType((*CsrRequest)(nil), "istio.v1.auth.CsrRequest")
	proto.RegisterType((*CsrResponse)(nil), "istio.v1.auth.CsrResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for IstioCAService service

type IstioCAServiceClient interface {
	// A request object includes a PEM-encoded certificate signing request that
	// is generated on the Node Agent. Additionally credential can be attached
	// within the request object for a server to authenticate the originating
	// node agent.
	HandleCSR(ctx context.Context, in *CsrRequest, opts ...grpc.CallOption) (*CsrResponse, error)
}

type istioCAServiceClient struct {
	cc *grpc.ClientConn
}

func NewIstioCAServiceClient(cc *grpc.ClientConn) IstioCAServiceClient {
	return &istioCAServiceClient{cc}
}

func (c *istioCAServiceClient) HandleCSR(ctx context.Context, in *CsrRequest, opts ...grpc.CallOption) (*CsrResponse, error) {
	out := new(CsrResponse)
	err := grpc.Invoke(ctx, "/istio.v1.auth.IstioCAService/HandleCSR", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for IstioCAService service

type IstioCAServiceServer interface {
	// A request object includes a PEM-encoded certificate signing request that
	// is generated on the Node Agent. Additionally credential can be attached
	// within the request object for a server to authenticate the originating
	// node agent.
	HandleCSR(context.Context, *CsrRequest) (*CsrResponse, error)
}

func RegisterIstioCAServiceServer(s *grpc.Server, srv IstioCAServiceServer) {
	s.RegisterService(&_IstioCAService_serviceDesc, srv)
}

func _IstioCAService_HandleCSR_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CsrRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IstioCAServiceServer).HandleCSR(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/istio.v1.auth.IstioCAService/HandleCSR",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IstioCAServiceServer).HandleCSR(ctx, req.(*CsrRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _IstioCAService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "istio.v1.auth.IstioCAService",
	HandlerType: (*IstioCAServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HandleCSR",
			Handler:    _IstioCAService_HandleCSR_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "security/proto/ca_service.proto",
}

func (m *CsrRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CsrRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.CsrPem) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCaService(dAtA, i, uint64(len(m.CsrPem)))
		i += copy(dAtA[i:], m.CsrPem)
	}
	if len(m.NodeAgentCredential) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCaService(dAtA, i, uint64(len(m.NodeAgentCredential)))
		i += copy(dAtA[i:], m.NodeAgentCredential)
	}
	if len(m.CredentialType) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintCaService(dAtA, i, uint64(len(m.CredentialType)))
		i += copy(dAtA[i:], m.CredentialType)
	}
	if m.RequestedTtlMinutes != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintCaService(dAtA, i, uint64(m.RequestedTtlMinutes))
	}
	return i, nil
}

func (m *CsrResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CsrResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.IsApproved {
		dAtA[i] = 0x8
		i++
		if m.IsApproved {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.Status != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCaService(dAtA, i, uint64(m.Status.Size()))
		n1, err := m.Status.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.SignedCert) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintCaService(dAtA, i, uint64(len(m.SignedCert)))
		i += copy(dAtA[i:], m.SignedCert)
	}
	if len(m.CertChain) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintCaService(dAtA, i, uint64(len(m.CertChain)))
		i += copy(dAtA[i:], m.CertChain)
	}
	return i, nil
}

func encodeVarintCaService(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *CsrRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.CsrPem)
	if l > 0 {
		n += 1 + l + sovCaService(uint64(l))
	}
	l = len(m.NodeAgentCredential)
	if l > 0 {
		n += 1 + l + sovCaService(uint64(l))
	}
	l = len(m.CredentialType)
	if l > 0 {
		n += 1 + l + sovCaService(uint64(l))
	}
	if m.RequestedTtlMinutes != 0 {
		n += 1 + sovCaService(uint64(m.RequestedTtlMinutes))
	}
	return n
}

func (m *CsrResponse) Size() (n int) {
	var l int
	_ = l
	if m.IsApproved {
		n += 2
	}
	if m.Status != nil {
		l = m.Status.Size()
		n += 1 + l + sovCaService(uint64(l))
	}
	l = len(m.SignedCert)
	if l > 0 {
		n += 1 + l + sovCaService(uint64(l))
	}
	l = len(m.CertChain)
	if l > 0 {
		n += 1 + l + sovCaService(uint64(l))
	}
	return n
}

func sovCaService(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozCaService(x uint64) (n int) {
	return sovCaService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *CsrRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CsrRequest{`,
		`CsrPem:` + fmt.Sprintf("%v", this.CsrPem) + `,`,
		`NodeAgentCredential:` + fmt.Sprintf("%v", this.NodeAgentCredential) + `,`,
		`CredentialType:` + fmt.Sprintf("%v", this.CredentialType) + `,`,
		`RequestedTtlMinutes:` + fmt.Sprintf("%v", this.RequestedTtlMinutes) + `,`,
		`}`,
	}, "")
	return s
}
func (this *CsrResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CsrResponse{`,
		`IsApproved:` + fmt.Sprintf("%v", this.IsApproved) + `,`,
		`Status:` + strings.Replace(fmt.Sprintf("%v", this.Status), "Status", "google_rpc.Status", 1) + `,`,
		`SignedCert:` + fmt.Sprintf("%v", this.SignedCert) + `,`,
		`CertChain:` + fmt.Sprintf("%v", this.CertChain) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringCaService(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *CsrRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCaService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CsrRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CsrRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CsrPem", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCaService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCaService
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CsrPem = append(m.CsrPem[:0], dAtA[iNdEx:postIndex]...)
			if m.CsrPem == nil {
				m.CsrPem = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeAgentCredential", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCaService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCaService
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NodeAgentCredential = append(m.NodeAgentCredential[:0], dAtA[iNdEx:postIndex]...)
			if m.NodeAgentCredential == nil {
				m.NodeAgentCredential = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CredentialType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCaService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCaService
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CredentialType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestedTtlMinutes", wireType)
			}
			m.RequestedTtlMinutes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCaService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RequestedTtlMinutes |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCaService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCaService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CsrResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCaService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CsrResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CsrResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsApproved", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCaService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsApproved = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCaService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCaService
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Status == nil {
				m.Status = &google_rpc.Status{}
			}
			if err := m.Status.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignedCert", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCaService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCaService
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SignedCert = append(m.SignedCert[:0], dAtA[iNdEx:postIndex]...)
			if m.SignedCert == nil {
				m.SignedCert = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CertChain", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCaService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCaService
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CertChain = append(m.CertChain[:0], dAtA[iNdEx:postIndex]...)
			if m.CertChain == nil {
				m.CertChain = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCaService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCaService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipCaService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCaService
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCaService
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCaService
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthCaService
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCaService
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipCaService(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthCaService = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCaService   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("security/proto/ca_service.proto", fileDescriptorCaService) }

var fileDescriptorCaService = []byte{
	// 428 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x3f, 0x8e, 0x13, 0x31,
	0x18, 0xc5, 0xc7, 0xfc, 0x09, 0xc4, 0x09, 0x8b, 0xe4, 0x05, 0x6d, 0x88, 0x84, 0x37, 0xda, 0x86,
	0x88, 0x62, 0xa2, 0x0d, 0x07, 0x40, 0xd9, 0x69, 0xa0, 0x40, 0x42, 0x4e, 0x2a, 0x1a, 0x6b, 0xf0,
	0x7c, 0x9a, 0xb5, 0x34, 0xb1, 0x8d, 0xed, 0x89, 0x94, 0x8e, 0x23, 0x50, 0x73, 0x02, 0x4e, 0xc0,
	0x19, 0xb6, 0xdc, 0x92, 0x92, 0x0c, 0x0d, 0xe5, 0x1e, 0x01, 0xd9, 0x5e, 0x25, 0x42, 0xda, 0x6e,
	0xf4, 0x7e, 0x4f, 0x4f, 0x6f, 0xbe, 0x67, 0x7c, 0xea, 0x40, 0xb4, 0x56, 0xfa, 0xed, 0xcc, 0x58,
	0xed, 0xf5, 0x4c, 0x94, 0xdc, 0x81, 0xdd, 0x48, 0x01, 0x79, 0x14, 0xc8, 0x13, 0xe9, 0xbc, 0xd4,
	0xf9, 0xe6, 0x3c, 0x2f, 0x5b, 0x7f, 0x39, 0x3e, 0xa9, 0xb5, 0xae, 0x1b, 0x98, 0x59, 0x23, 0x66,
	0xce, 0x97, 0xbe, 0x75, 0xc9, 0x37, 0x7e, 0x56, 0xeb, 0x5a, 0xa7, 0x8c, 0xf0, 0x95, 0xd4, 0xb3,
	0x9f, 0x08, 0xe3, 0xc2, 0x59, 0x06, 0x5f, 0x5a, 0x70, 0x9e, 0x9c, 0xe0, 0x47, 0xc2, 0x59, 0x6e,
	0x60, 0x3d, 0x42, 0x13, 0x34, 0x1d, 0xb2, 0x9e, 0x70, 0xf6, 0x23, 0xac, 0xc9, 0x1c, 0x3f, 0x57,
	0xba, 0x02, 0x5e, 0xd6, 0xa0, 0x3c, 0x17, 0x16, 0x2a, 0x50, 0x5e, 0x96, 0xcd, 0xe8, 0x5e, 0xb4,
	0x1d, 0x07, 0xb8, 0x08, 0xac, 0xd8, 0x23, 0xf2, 0x0a, 0x3f, 0x3d, 0x18, 0xb9, 0xdf, 0x1a, 0x18,
	0xdd, 0x9f, 0xa0, 0x69, 0x9f, 0x1d, 0x1d, 0xe4, 0xd5, 0xd6, 0x40, 0x08, 0xb7, 0xa9, 0x00, 0x54,
	0xdc, 0xfb, 0x86, 0xaf, 0xa5, 0x6a, 0x3d, 0xb8, 0xd1, 0x83, 0x09, 0x9a, 0x3e, 0x64, 0xc7, 0x7b,
	0xb8, 0xf2, 0xcd, 0x87, 0x84, 0xce, 0xbe, 0x23, 0x3c, 0x88, 0xc5, 0x9d, 0xd1, 0xca, 0x01, 0x39,
	0xc5, 0x03, 0xe9, 0x78, 0x69, 0x8c, 0xd5, 0x1b, 0xa8, 0x62, 0xfb, 0xc7, 0x0c, 0x4b, 0xb7, 0xb8,
	0x55, 0xc8, 0x6b, 0xdc, 0x4b, 0xf7, 0x88, 0x95, 0x07, 0x73, 0x92, 0xa7, 0x4b, 0xe5, 0xd6, 0x88,
	0x7c, 0x19, 0x09, 0xbb, 0x75, 0x84, 0x30, 0x27, 0x6b, 0x05, 0x15, 0x17, 0x60, 0x7d, 0x6c, 0x3d,
	0x64, 0x38, 0x49, 0x05, 0x58, 0x4f, 0x5e, 0x62, 0x1c, 0x08, 0x17, 0x97, 0xa5, 0x54, 0xb1, 0xe6,
	0x90, 0xf5, 0x83, 0x52, 0x04, 0x61, 0xbe, 0xc2, 0x47, 0xef, 0xc3, 0x2a, 0xc5, 0x62, 0x99, 0xb6,
	0x22, 0x17, 0xb8, 0xff, 0xae, 0x54, 0x55, 0x03, 0xc5, 0x92, 0x91, 0x17, 0xf9, 0x7f, 0x9b, 0xe5,
	0x87, 0x01, 0xc6, 0xe3, 0xbb, 0x50, 0xfa, 0xc5, 0x8b, 0xb7, 0x57, 0x3b, 0x9a, 0x5d, 0xef, 0x68,
	0xf6, 0x6b, 0x47, 0xb3, 0x9b, 0x1d, 0xcd, 0xbe, 0x76, 0x14, 0xfd, 0xe8, 0x68, 0x76, 0xd5, 0x51,
	0x74, 0xdd, 0x51, 0xf4, 0xbb, 0xa3, 0xe8, 0x6f, 0x47, 0xb3, 0x9b, 0x8e, 0xa2, 0x6f, 0x7f, 0x68,
	0xf6, 0x29, 0xbd, 0x0d, 0xbe, 0x39, 0xe7, 0x21, 0xec, 0x73, 0x2f, 0x6e, 0xfe, 0xe6, 0x5f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x7c, 0xd3, 0x3f, 0x22, 0x54, 0x02, 0x00, 0x00,
}
