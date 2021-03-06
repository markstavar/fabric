// Code generated by protoc-gen-go.
// source: peer/chaincode.proto
// DO NOT EDIT!

package peer

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf1 "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Confidentiality Levels
type ConfidentialityLevel int32

const (
	ConfidentialityLevel_PUBLIC       ConfidentialityLevel = 0
	ConfidentialityLevel_CONFIDENTIAL ConfidentialityLevel = 1
)

var ConfidentialityLevel_name = map[int32]string{
	0: "PUBLIC",
	1: "CONFIDENTIAL",
}
var ConfidentialityLevel_value = map[string]int32{
	"PUBLIC":       0,
	"CONFIDENTIAL": 1,
}

func (x ConfidentialityLevel) String() string {
	return proto.EnumName(ConfidentialityLevel_name, int32(x))
}
func (ConfidentialityLevel) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type ChaincodeSpec_Type int32

const (
	ChaincodeSpec_UNDEFINED ChaincodeSpec_Type = 0
	ChaincodeSpec_GOLANG    ChaincodeSpec_Type = 1
	ChaincodeSpec_NODE      ChaincodeSpec_Type = 2
	ChaincodeSpec_CAR       ChaincodeSpec_Type = 3
	ChaincodeSpec_JAVA      ChaincodeSpec_Type = 4
)

var ChaincodeSpec_Type_name = map[int32]string{
	0: "UNDEFINED",
	1: "GOLANG",
	2: "NODE",
	3: "CAR",
	4: "JAVA",
}
var ChaincodeSpec_Type_value = map[string]int32{
	"UNDEFINED": 0,
	"GOLANG":    1,
	"NODE":      2,
	"CAR":       3,
	"JAVA":      4,
}

func (x ChaincodeSpec_Type) String() string {
	return proto.EnumName(ChaincodeSpec_Type_name, int32(x))
}
func (ChaincodeSpec_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{2, 0} }

type ChaincodeDeploymentSpec_ExecutionEnvironment int32

const (
	ChaincodeDeploymentSpec_DOCKER ChaincodeDeploymentSpec_ExecutionEnvironment = 0
	ChaincodeDeploymentSpec_SYSTEM ChaincodeDeploymentSpec_ExecutionEnvironment = 1
)

var ChaincodeDeploymentSpec_ExecutionEnvironment_name = map[int32]string{
	0: "DOCKER",
	1: "SYSTEM",
}
var ChaincodeDeploymentSpec_ExecutionEnvironment_value = map[string]int32{
	"DOCKER": 0,
	"SYSTEM": 1,
}

func (x ChaincodeDeploymentSpec_ExecutionEnvironment) String() string {
	return proto.EnumName(ChaincodeDeploymentSpec_ExecutionEnvironment_name, int32(x))
}
func (ChaincodeDeploymentSpec_ExecutionEnvironment) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor1, []int{3, 0}
}

// ChaincodeID contains the path as specified by the deploy transaction
// that created it as well as the hashCode that is generated by the
// system for the path. From the user level (ie, CLI, REST API and so on)
// deploy transaction is expected to provide the path and other requests
// are expected to provide the hashCode. The other value will be ignored.
// Internally, the structure could contain both values. For instance, the
// hashCode will be set when first generated using the path
type ChaincodeID struct {
	// deploy transaction will use the path
	Path string `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	// all other requests will use the name (really a hashcode) generated by
	// the deploy transaction
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	// user friendly version name for the chaincode
	Version string `protobuf:"bytes,3,opt,name=version" json:"version,omitempty"`
}

func (m *ChaincodeID) Reset()                    { *m = ChaincodeID{} }
func (m *ChaincodeID) String() string            { return proto.CompactTextString(m) }
func (*ChaincodeID) ProtoMessage()               {}
func (*ChaincodeID) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

// Carries the chaincode function and its arguments.
// UnmarshalJSON in transaction.go converts the string-based REST/JSON input to
// the []byte-based current ChaincodeInput structure.
type ChaincodeInput struct {
	Args [][]byte `protobuf:"bytes,1,rep,name=args,proto3" json:"args,omitempty"`
}

func (m *ChaincodeInput) Reset()                    { *m = ChaincodeInput{} }
func (m *ChaincodeInput) String() string            { return proto.CompactTextString(m) }
func (*ChaincodeInput) ProtoMessage()               {}
func (*ChaincodeInput) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

// Carries the chaincode specification. This is the actual metadata required for
// defining a chaincode.
type ChaincodeSpec struct {
	Type        ChaincodeSpec_Type `protobuf:"varint,1,opt,name=type,enum=protos.ChaincodeSpec_Type" json:"type,omitempty"`
	ChaincodeId *ChaincodeID       `protobuf:"bytes,2,opt,name=chaincode_id,json=chaincodeId" json:"chaincode_id,omitempty"`
	Input       *ChaincodeInput    `protobuf:"bytes,3,opt,name=input" json:"input,omitempty"`
	Timeout     int32              `protobuf:"varint,4,opt,name=timeout" json:"timeout,omitempty"`
}

func (m *ChaincodeSpec) Reset()                    { *m = ChaincodeSpec{} }
func (m *ChaincodeSpec) String() string            { return proto.CompactTextString(m) }
func (*ChaincodeSpec) ProtoMessage()               {}
func (*ChaincodeSpec) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *ChaincodeSpec) GetChaincodeId() *ChaincodeID {
	if m != nil {
		return m.ChaincodeId
	}
	return nil
}

func (m *ChaincodeSpec) GetInput() *ChaincodeInput {
	if m != nil {
		return m.Input
	}
	return nil
}

// Specify the deployment of a chaincode.
// TODO: Define `codePackage`.
type ChaincodeDeploymentSpec struct {
	ChaincodeSpec *ChaincodeSpec `protobuf:"bytes,1,opt,name=chaincode_spec,json=chaincodeSpec" json:"chaincode_spec,omitempty"`
	// Controls when the chaincode becomes executable.
	EffectiveDate *google_protobuf1.Timestamp                  `protobuf:"bytes,2,opt,name=effective_date,json=effectiveDate" json:"effective_date,omitempty"`
	CodePackage   []byte                                       `protobuf:"bytes,3,opt,name=code_package,json=codePackage,proto3" json:"code_package,omitempty"`
	ExecEnv       ChaincodeDeploymentSpec_ExecutionEnvironment `protobuf:"varint,4,opt,name=exec_env,json=execEnv,enum=protos.ChaincodeDeploymentSpec_ExecutionEnvironment" json:"exec_env,omitempty"`
}

func (m *ChaincodeDeploymentSpec) Reset()                    { *m = ChaincodeDeploymentSpec{} }
func (m *ChaincodeDeploymentSpec) String() string            { return proto.CompactTextString(m) }
func (*ChaincodeDeploymentSpec) ProtoMessage()               {}
func (*ChaincodeDeploymentSpec) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *ChaincodeDeploymentSpec) GetChaincodeSpec() *ChaincodeSpec {
	if m != nil {
		return m.ChaincodeSpec
	}
	return nil
}

func (m *ChaincodeDeploymentSpec) GetEffectiveDate() *google_protobuf1.Timestamp {
	if m != nil {
		return m.EffectiveDate
	}
	return nil
}

// Carries the chaincode function and its arguments.
type ChaincodeInvocationSpec struct {
	ChaincodeSpec *ChaincodeSpec `protobuf:"bytes,1,opt,name=chaincode_spec,json=chaincodeSpec" json:"chaincode_spec,omitempty"`
	// This field can contain a user-specified ID generation algorithm
	// If supplied, this will be used to generate a ID
	// If not supplied (left empty), sha256base64 will be used
	// The algorithm consists of two parts:
	//  1, a hash function
	//  2, a decoding used to decode user (string) input to bytes
	// Currently, SHA256 with BASE64 is supported (e.g. idGenerationAlg='sha256base64')
	IdGenerationAlg string `protobuf:"bytes,2,opt,name=id_generation_alg,json=idGenerationAlg" json:"id_generation_alg,omitempty"`
}

func (m *ChaincodeInvocationSpec) Reset()                    { *m = ChaincodeInvocationSpec{} }
func (m *ChaincodeInvocationSpec) String() string            { return proto.CompactTextString(m) }
func (*ChaincodeInvocationSpec) ProtoMessage()               {}
func (*ChaincodeInvocationSpec) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *ChaincodeInvocationSpec) GetChaincodeSpec() *ChaincodeSpec {
	if m != nil {
		return m.ChaincodeSpec
	}
	return nil
}

func init() {
	proto.RegisterType((*ChaincodeID)(nil), "protos.ChaincodeID")
	proto.RegisterType((*ChaincodeInput)(nil), "protos.ChaincodeInput")
	proto.RegisterType((*ChaincodeSpec)(nil), "protos.ChaincodeSpec")
	proto.RegisterType((*ChaincodeDeploymentSpec)(nil), "protos.ChaincodeDeploymentSpec")
	proto.RegisterType((*ChaincodeInvocationSpec)(nil), "protos.ChaincodeInvocationSpec")
	proto.RegisterEnum("protos.ConfidentialityLevel", ConfidentialityLevel_name, ConfidentialityLevel_value)
	proto.RegisterEnum("protos.ChaincodeSpec_Type", ChaincodeSpec_Type_name, ChaincodeSpec_Type_value)
	proto.RegisterEnum("protos.ChaincodeDeploymentSpec_ExecutionEnvironment", ChaincodeDeploymentSpec_ExecutionEnvironment_name, ChaincodeDeploymentSpec_ExecutionEnvironment_value)
}

func init() { proto.RegisterFile("peer/chaincode.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 591 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x93, 0x5d, 0x6f, 0xd3, 0x3e,
	0x14, 0xc6, 0x97, 0xb5, 0x7b, 0x3b, 0x7d, 0xf9, 0xe7, 0x6f, 0x06, 0x54, 0xbb, 0x61, 0x44, 0x5c,
	0x8c, 0x09, 0xa5, 0x52, 0x99, 0xb8, 0x42, 0x48, 0x59, 0x93, 0x4d, 0x81, 0xd2, 0x4c, 0x59, 0x87,
	0x04, 0x37, 0x95, 0x9b, 0x9c, 0xa6, 0x16, 0xa9, 0x1d, 0x25, 0x6e, 0xb4, 0x5e, 0xf3, 0xbd, 0xf8,
	0x6a, 0x20, 0x3b, 0x6b, 0xb7, 0xa9, 0xbb, 0xe4, 0x2a, 0xf6, 0x93, 0xe7, 0xd8, 0xcf, 0xf9, 0xe9,
	0x18, 0x0e, 0x33, 0xc4, 0xbc, 0x1b, 0xcd, 0x28, 0xe3, 0x91, 0x88, 0xd1, 0xce, 0x72, 0x21, 0x05,
	0xd9, 0xd5, 0x9f, 0xe2, 0xe8, 0x55, 0x22, 0x44, 0x92, 0x62, 0x57, 0x6f, 0x27, 0x8b, 0x69, 0x57,
	0xb2, 0x39, 0x16, 0x92, 0xce, 0xb3, 0xca, 0x68, 0x05, 0xd0, 0xe8, 0xaf, 0x6a, 0x7d, 0x97, 0x10,
	0xa8, 0x67, 0x54, 0xce, 0x3a, 0xc6, 0xb1, 0x71, 0x72, 0x10, 0xea, 0xb5, 0xd2, 0x38, 0x9d, 0x63,
	0x67, 0xbb, 0xd2, 0xd4, 0x9a, 0x74, 0x60, 0xaf, 0xc4, 0xbc, 0x60, 0x82, 0x77, 0x6a, 0x5a, 0x5e,
	0x6d, 0xad, 0x37, 0xd0, 0xbe, 0x3f, 0x90, 0x67, 0x0b, 0xa9, 0xea, 0x69, 0x9e, 0x14, 0x1d, 0xe3,
	0xb8, 0x76, 0xd2, 0x0c, 0xf5, 0xda, 0xfa, 0x63, 0x40, 0x6b, 0x6d, 0xbb, 0xce, 0x30, 0x22, 0x36,
	0xd4, 0xe5, 0x32, 0x43, 0x7d, 0x73, 0xbb, 0x77, 0x54, 0xc5, 0x2b, 0xec, 0x47, 0x26, 0x7b, 0xb4,
	0xcc, 0x30, 0xd4, 0x3e, 0xf2, 0x01, 0x9a, 0xeb, 0xa6, 0xc7, 0x2c, 0xd6, 0xe9, 0x1a, 0xbd, 0x67,
	0x1b, 0x75, 0xbe, 0x1b, 0x36, 0xd6, 0x46, 0x3f, 0x26, 0xef, 0x60, 0x87, 0xa9, 0x58, 0x3a, 0x77,
	0xa3, 0xf7, 0x62, 0xb3, 0x40, 0xfd, 0x0d, 0x2b, 0x93, 0xea, 0x53, 0x11, 0x13, 0x0b, 0xd9, 0xa9,
	0x1f, 0x1b, 0x27, 0x3b, 0xe1, 0x6a, 0x6b, 0x7d, 0x82, 0xba, 0x4a, 0x43, 0x5a, 0x70, 0x70, 0x33,
	0x74, 0xbd, 0x0b, 0x7f, 0xe8, 0xb9, 0xe6, 0x16, 0x01, 0xd8, 0xbd, 0x0c, 0x06, 0xce, 0xf0, 0xd2,
	0x34, 0xc8, 0x3e, 0xd4, 0x87, 0x81, 0xeb, 0x99, 0xdb, 0x64, 0x0f, 0x6a, 0x7d, 0x27, 0x34, 0x6b,
	0x4a, 0xfa, 0xec, 0x7c, 0x73, 0xcc, 0xba, 0xf5, 0x7b, 0x1b, 0x5e, 0xae, 0xef, 0x74, 0x31, 0x4b,
	0xc5, 0x72, 0x8e, 0x5c, 0x6a, 0x16, 0x1f, 0xa1, 0x7d, 0xdf, 0x5b, 0x91, 0x61, 0xa4, 0xa9, 0x34,
	0x7a, 0xcf, 0x9f, 0xa4, 0x12, 0xb6, 0xa2, 0x47, 0x24, 0x1d, 0x68, 0xe3, 0x74, 0x8a, 0x91, 0x64,
	0x25, 0x8e, 0x63, 0x2a, 0xf1, 0x8e, 0xcd, 0x91, 0x5d, 0x0d, 0x83, 0xbd, 0x1a, 0x06, 0x7b, 0xb4,
	0x1a, 0x86, 0xb0, 0xb5, 0xae, 0x70, 0xa9, 0x44, 0xf2, 0x1a, 0x9a, 0xfa, 0xee, 0x8c, 0x46, 0x3f,
	0x69, 0x82, 0x9a, 0x55, 0x33, 0x6c, 0x28, 0xed, 0xaa, 0x92, 0x48, 0x00, 0xfb, 0x78, 0x8b, 0xd1,
	0x18, 0x79, 0xa9, 0xd1, 0xb4, 0x7b, 0x67, 0x1b, 0xe9, 0x1e, 0xb7, 0x65, 0x7b, 0xb7, 0x18, 0x2d,
	0x24, 0x13, 0xdc, 0xe3, 0x25, 0xcb, 0x05, 0x57, 0x3f, 0xc2, 0x3d, 0x75, 0x8a, 0xc7, 0x4b, 0xcb,
	0x86, 0xc3, 0xa7, 0x0c, 0x8a, 0xa8, 0x1b, 0xf4, 0xbf, 0x78, 0x61, 0x45, 0xf7, 0xfa, 0xfb, 0xf5,
	0xc8, 0xfb, 0x6a, 0x1a, 0xd6, 0x2f, 0xe3, 0x01, 0x40, 0x9f, 0x97, 0x22, 0xa2, 0xaa, 0xf4, 0x1f,
	0x00, 0x3c, 0x85, 0xff, 0x59, 0x3c, 0x4e, 0x90, 0x63, 0xae, 0x8f, 0x1c, 0xd3, 0x34, 0xb9, 0x9b,
	0xfe, 0xff, 0x58, 0x7c, 0xb9, 0xd6, 0x9d, 0x34, 0x39, 0x3d, 0x83, 0xc3, 0xbe, 0xe0, 0x53, 0x16,
	0x23, 0x97, 0x8c, 0xa6, 0x4c, 0x2e, 0x07, 0x58, 0x62, 0xaa, 0x92, 0x5e, 0xdd, 0x9c, 0x0f, 0xfc,
	0xbe, 0xb9, 0x45, 0x4c, 0x68, 0xf6, 0x83, 0xe1, 0x85, 0xef, 0x7a, 0xc3, 0x91, 0xef, 0x0c, 0x4c,
	0xe3, 0x3c, 0x00, 0x4b, 0xe4, 0x89, 0x3d, 0x5b, 0x66, 0x98, 0xa7, 0x18, 0x27, 0x98, 0xdb, 0x53,
	0x3a, 0xc9, 0x59, 0xb4, 0xca, 0xa7, 0x1e, 0xf5, 0x8f, 0xb7, 0x09, 0x93, 0xb3, 0xc5, 0xc4, 0x8e,
	0xc4, 0xbc, 0xfb, 0xc0, 0xda, 0xad, 0xac, 0xd5, 0x9b, 0x2e, 0xba, 0xca, 0x3a, 0xa9, 0xde, 0xfb,
	0xfb, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x05, 0x18, 0x06, 0x44, 0x0e, 0x04, 0x00, 0x00,
}
