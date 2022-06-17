// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: paloma/evm/turnstone/turnstone.proto

package turnstone

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Valset struct {
	HexAddress []string `protobuf:"bytes,1,rep,name=hexAddress,proto3" json:"hexAddress,omitempty"`
	Powers     []uint32 `protobuf:"varint,2,rep,packed,name=powers,proto3" json:"powers,omitempty"`
	ValsetID   uint64   `protobuf:"varint,3,opt,name=valsetID,proto3" json:"valsetID,omitempty"`
}

func (m *Valset) Reset()         { *m = Valset{} }
func (m *Valset) String() string { return proto.CompactTextString(m) }
func (*Valset) ProtoMessage()    {}
func (*Valset) Descriptor() ([]byte, []int) {
	return fileDescriptor_0af0c7ee2b25d4ee, []int{0}
}
func (m *Valset) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Valset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Valset.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Valset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Valset.Merge(m, src)
}
func (m *Valset) XXX_Size() int {
	return m.Size()
}
func (m *Valset) XXX_DiscardUnknown() {
	xxx_messageInfo_Valset.DiscardUnknown(m)
}

var xxx_messageInfo_Valset proto.InternalMessageInfo

func (m *Valset) GetHexAddress() []string {
	if m != nil {
		return m.HexAddress
	}
	return nil
}

func (m *Valset) GetPowers() []uint32 {
	if m != nil {
		return m.Powers
	}
	return nil
}

func (m *Valset) GetValsetID() uint64 {
	if m != nil {
		return m.ValsetID
	}
	return 0
}

type SubmitLogicCall struct {
	HexContractAddress string `protobuf:"bytes,1,opt,name=hexContractAddress,proto3" json:"hexContractAddress,omitempty"`
	Abi                []byte `protobuf:"bytes,2,opt,name=abi,proto3" json:"abi,omitempty"`
	Payload            []byte `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	Deadline           int64  `protobuf:"varint,4,opt,name=deadline,proto3" json:"deadline,omitempty"`
}

func (m *SubmitLogicCall) Reset()         { *m = SubmitLogicCall{} }
func (m *SubmitLogicCall) String() string { return proto.CompactTextString(m) }
func (*SubmitLogicCall) ProtoMessage()    {}
func (*SubmitLogicCall) Descriptor() ([]byte, []int) {
	return fileDescriptor_0af0c7ee2b25d4ee, []int{1}
}
func (m *SubmitLogicCall) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SubmitLogicCall) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SubmitLogicCall.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SubmitLogicCall) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubmitLogicCall.Merge(m, src)
}
func (m *SubmitLogicCall) XXX_Size() int {
	return m.Size()
}
func (m *SubmitLogicCall) XXX_DiscardUnknown() {
	xxx_messageInfo_SubmitLogicCall.DiscardUnknown(m)
}

var xxx_messageInfo_SubmitLogicCall proto.InternalMessageInfo

func (m *SubmitLogicCall) GetHexContractAddress() string {
	if m != nil {
		return m.HexContractAddress
	}
	return ""
}

func (m *SubmitLogicCall) GetAbi() []byte {
	if m != nil {
		return m.Abi
	}
	return nil
}

func (m *SubmitLogicCall) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *SubmitLogicCall) GetDeadline() int64 {
	if m != nil {
		return m.Deadline
	}
	return 0
}

type UpdateValset struct {
	Valset *Valset `protobuf:"bytes,1,opt,name=valset,proto3" json:"valset,omitempty"`
}

func (m *UpdateValset) Reset()         { *m = UpdateValset{} }
func (m *UpdateValset) String() string { return proto.CompactTextString(m) }
func (*UpdateValset) ProtoMessage()    {}
func (*UpdateValset) Descriptor() ([]byte, []int) {
	return fileDescriptor_0af0c7ee2b25d4ee, []int{2}
}
func (m *UpdateValset) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UpdateValset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UpdateValset.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UpdateValset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateValset.Merge(m, src)
}
func (m *UpdateValset) XXX_Size() int {
	return m.Size()
}
func (m *UpdateValset) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateValset.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateValset proto.InternalMessageInfo

func (m *UpdateValset) GetValset() *Valset {
	if m != nil {
		return m.Valset
	}
	return nil
}

type UploadSmartContract struct {
	Bytecode []byte `protobuf:"bytes,1,opt,name=bytecode,proto3" json:"bytecode,omitempty"`
}

func (m *UploadSmartContract) Reset()         { *m = UploadSmartContract{} }
func (m *UploadSmartContract) String() string { return proto.CompactTextString(m) }
func (*UploadSmartContract) ProtoMessage()    {}
func (*UploadSmartContract) Descriptor() ([]byte, []int) {
	return fileDescriptor_0af0c7ee2b25d4ee, []int{3}
}
func (m *UploadSmartContract) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UploadSmartContract) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UploadSmartContract.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UploadSmartContract) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadSmartContract.Merge(m, src)
}
func (m *UploadSmartContract) XXX_Size() int {
	return m.Size()
}
func (m *UploadSmartContract) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadSmartContract.DiscardUnknown(m)
}

var xxx_messageInfo_UploadSmartContract proto.InternalMessageInfo

func (m *UploadSmartContract) GetBytecode() []byte {
	if m != nil {
		return m.Bytecode
	}
	return nil
}

type Message struct {
	TurnstoneID uint64 `protobuf:"varint,1,opt,name=turnstoneID,proto3" json:"turnstoneID,omitempty"`
	ChainID     string `protobuf:"bytes,2,opt,name=chainID,proto3" json:"chainID,omitempty"`
	// Types that are valid to be assigned to Action:
	//	*Message_SubmitLogicCall
	//	*Message_UpdateValset
	//	*Message_UploadSmartContract
	Action isMessage_Action `protobuf_oneof:"action"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_0af0c7ee2b25d4ee, []int{4}
}
func (m *Message) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Message.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return m.Size()
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

type isMessage_Action interface {
	isMessage_Action()
	MarshalTo([]byte) (int, error)
	Size() int
}

type Message_SubmitLogicCall struct {
	SubmitLogicCall *SubmitLogicCall `protobuf:"bytes,3,opt,name=submitLogicCall,proto3,oneof" json:"submitLogicCall,omitempty"`
}
type Message_UpdateValset struct {
	UpdateValset *UpdateValset `protobuf:"bytes,4,opt,name=updateValset,proto3,oneof" json:"updateValset,omitempty"`
}
type Message_UploadSmartContract struct {
	UploadSmartContract *UploadSmartContract `protobuf:"bytes,5,opt,name=uploadSmartContract,proto3,oneof" json:"uploadSmartContract,omitempty"`
}

func (*Message_SubmitLogicCall) isMessage_Action()     {}
func (*Message_UpdateValset) isMessage_Action()        {}
func (*Message_UploadSmartContract) isMessage_Action() {}

func (m *Message) GetAction() isMessage_Action {
	if m != nil {
		return m.Action
	}
	return nil
}

func (m *Message) GetTurnstoneID() uint64 {
	if m != nil {
		return m.TurnstoneID
	}
	return 0
}

func (m *Message) GetChainID() string {
	if m != nil {
		return m.ChainID
	}
	return ""
}

func (m *Message) GetSubmitLogicCall() *SubmitLogicCall {
	if x, ok := m.GetAction().(*Message_SubmitLogicCall); ok {
		return x.SubmitLogicCall
	}
	return nil
}

func (m *Message) GetUpdateValset() *UpdateValset {
	if x, ok := m.GetAction().(*Message_UpdateValset); ok {
		return x.UpdateValset
	}
	return nil
}

func (m *Message) GetUploadSmartContract() *UploadSmartContract {
	if x, ok := m.GetAction().(*Message_UploadSmartContract); ok {
		return x.UploadSmartContract
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Message) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Message_SubmitLogicCall)(nil),
		(*Message_UpdateValset)(nil),
		(*Message_UploadSmartContract)(nil),
	}
}

func init() {
	proto.RegisterType((*Valset)(nil), "palomachain.paloma.evm.turnstone.Valset")
	proto.RegisterType((*SubmitLogicCall)(nil), "palomachain.paloma.evm.turnstone.SubmitLogicCall")
	proto.RegisterType((*UpdateValset)(nil), "palomachain.paloma.evm.turnstone.UpdateValset")
	proto.RegisterType((*UploadSmartContract)(nil), "palomachain.paloma.evm.turnstone.UploadSmartContract")
	proto.RegisterType((*Message)(nil), "palomachain.paloma.evm.turnstone.Message")
}

func init() {
	proto.RegisterFile("paloma/evm/turnstone/turnstone.proto", fileDescriptor_0af0c7ee2b25d4ee)
}

var fileDescriptor_0af0c7ee2b25d4ee = []byte{
	// 467 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x4d, 0x6f, 0x9b, 0x40,
	0x10, 0x05, 0xdb, 0x25, 0xf1, 0x98, 0x2a, 0xd5, 0xa6, 0xaa, 0x50, 0x0e, 0x08, 0xa1, 0x1e, 0x38,
	0x2d, 0x8a, 0xa3, 0xde, 0xdb, 0xc4, 0x07, 0x5b, 0x6a, 0xa4, 0x6a, 0xd3, 0xf4, 0x50, 0xb5, 0x87,
	0x05, 0x56, 0x18, 0x09, 0x58, 0xc4, 0x2e, 0xae, 0xfd, 0x13, 0x7a, 0xeb, 0xcf, 0xea, 0x31, 0xc7,
	0x1e, 0x2b, 0xfb, 0x07, 0xf4, 0x2f, 0x54, 0x2c, 0x36, 0x26, 0x1f, 0x92, 0x73, 0x9b, 0xb7, 0xcc,
	0xbc, 0xf7, 0xe6, 0xa1, 0x81, 0xb7, 0x05, 0x4d, 0x79, 0x46, 0x7d, 0xb6, 0xc8, 0x7c, 0x59, 0x95,
	0xb9, 0x90, 0x3c, 0x67, 0xfb, 0x0a, 0x17, 0x25, 0x97, 0x1c, 0x39, 0x4d, 0x57, 0x38, 0xa7, 0x49,
	0x8e, 0x9b, 0x1a, 0xb3, 0x45, 0x86, 0xdb, 0xbe, 0xb3, 0xd7, 0x31, 0x8f, 0xb9, 0x6a, 0xf6, 0xeb,
	0xaa, 0x99, 0x73, 0xbf, 0x81, 0xf1, 0x85, 0xa6, 0x82, 0x49, 0x64, 0x03, 0xcc, 0xd9, 0xf2, 0x43,
	0x14, 0x95, 0x4c, 0x08, 0x4b, 0x77, 0xfa, 0xde, 0x90, 0x74, 0x5e, 0xd0, 0x1b, 0x30, 0x0a, 0xfe,
	0x83, 0x95, 0xc2, 0xea, 0x39, 0x7d, 0xef, 0x25, 0xd9, 0x22, 0x74, 0x06, 0xc7, 0x0b, 0xc5, 0x30,
	0x9b, 0x58, 0x7d, 0x47, 0xf7, 0x06, 0xa4, 0xc5, 0xee, 0x4f, 0x1d, 0x4e, 0x6e, 0xaa, 0x20, 0x4b,
	0xe4, 0x47, 0x1e, 0x27, 0xe1, 0x15, 0x4d, 0x53, 0x84, 0x01, 0xcd, 0xd9, 0xf2, 0x8a, 0xe7, 0xb2,
	0xa4, 0xa1, 0xdc, 0xeb, 0xe9, 0xde, 0x90, 0x3c, 0xf1, 0x05, 0xbd, 0x82, 0x3e, 0x0d, 0x12, 0xab,
	0xe7, 0xe8, 0x9e, 0x49, 0xea, 0x12, 0x59, 0x70, 0x54, 0xd0, 0x55, 0xca, 0x69, 0xa4, 0x04, 0x4d,
	0xb2, 0x83, 0xb5, 0x97, 0x88, 0xd1, 0x28, 0x4d, 0x72, 0x66, 0x0d, 0x1c, 0xdd, 0xeb, 0x93, 0x16,
	0xbb, 0x9f, 0xc0, 0xbc, 0x2d, 0x22, 0x2a, 0xd9, 0x76, 0xdf, 0xf7, 0x60, 0x34, 0x3e, 0x95, 0xf6,
	0x68, 0xec, 0xe1, 0x43, 0x11, 0xe2, 0x66, 0x92, 0x6c, 0xe7, 0xdc, 0x73, 0x38, 0xbd, 0x2d, 0x6a,
	0xdd, 0x9b, 0x8c, 0x96, 0x72, 0xe7, 0xbb, 0x36, 0x11, 0xac, 0x24, 0x0b, 0x79, 0xc4, 0x14, 0xb5,
	0x49, 0x5a, 0xec, 0xfe, 0xeb, 0xc1, 0xd1, 0x35, 0x13, 0x82, 0xc6, 0x0c, 0x39, 0x30, 0x6a, 0xa9,
	0x67, 0x13, 0xd5, 0x3a, 0x20, 0xdd, 0xa7, 0x7a, 0x51, 0xe5, 0x66, 0x36, 0x51, 0xeb, 0x0f, 0xc9,
	0x0e, 0xa2, 0xef, 0x70, 0x22, 0xee, 0xe7, 0xaa, 0xa2, 0x18, 0x8d, 0xcf, 0x0f, 0x6f, 0xf1, 0xe0,
	0x87, 0x4c, 0x35, 0xf2, 0x90, 0x0b, 0x7d, 0x06, 0xb3, 0xea, 0x64, 0xa5, 0xb2, 0x1c, 0x8d, 0xf1,
	0x61, 0xee, 0x6e, 0xc2, 0x53, 0x8d, 0xdc, 0x63, 0x41, 0x09, 0x9c, 0x56, 0x8f, 0xf3, 0xb2, 0x5e,
	0x28, 0xf2, 0x77, 0xcf, 0x21, 0x7f, 0x34, 0x3c, 0xd5, 0xc8, 0x53, 0x9c, 0x97, 0xc7, 0x60, 0xd0,
	0x50, 0x26, 0x3c, 0xbf, 0xbc, 0xfe, 0xbd, 0xb6, 0xf5, 0xbb, 0xb5, 0xad, 0xff, 0x5d, 0xdb, 0xfa,
	0xaf, 0x8d, 0xad, 0xdd, 0x6d, 0x6c, 0xed, 0xcf, 0xc6, 0xd6, 0xbe, 0x5e, 0xc4, 0x89, 0x9c, 0x57,
	0x01, 0x0e, 0x79, 0xe6, 0x77, 0xb4, 0xb7, 0xb5, 0xbf, 0x6c, 0x2e, 0x6e, 0x55, 0x30, 0xb1, 0xbf,
	0xb6, 0xc0, 0x50, 0x67, 0x73, 0xf1, 0x3f, 0x00, 0x00, 0xff, 0xff, 0xed, 0xa9, 0x8e, 0xc5, 0x96,
	0x03, 0x00, 0x00,
}

func (m *Valset) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Valset) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Valset) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ValsetID != 0 {
		i = encodeVarintTurnstone(dAtA, i, uint64(m.ValsetID))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Powers) > 0 {
		dAtA2 := make([]byte, len(m.Powers)*10)
		var j1 int
		for _, num := range m.Powers {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintTurnstone(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x12
	}
	if len(m.HexAddress) > 0 {
		for iNdEx := len(m.HexAddress) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.HexAddress[iNdEx])
			copy(dAtA[i:], m.HexAddress[iNdEx])
			i = encodeVarintTurnstone(dAtA, i, uint64(len(m.HexAddress[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *SubmitLogicCall) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SubmitLogicCall) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SubmitLogicCall) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Deadline != 0 {
		i = encodeVarintTurnstone(dAtA, i, uint64(m.Deadline))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Payload) > 0 {
		i -= len(m.Payload)
		copy(dAtA[i:], m.Payload)
		i = encodeVarintTurnstone(dAtA, i, uint64(len(m.Payload)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Abi) > 0 {
		i -= len(m.Abi)
		copy(dAtA[i:], m.Abi)
		i = encodeVarintTurnstone(dAtA, i, uint64(len(m.Abi)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.HexContractAddress) > 0 {
		i -= len(m.HexContractAddress)
		copy(dAtA[i:], m.HexContractAddress)
		i = encodeVarintTurnstone(dAtA, i, uint64(len(m.HexContractAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *UpdateValset) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UpdateValset) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UpdateValset) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Valset != nil {
		{
			size, err := m.Valset.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTurnstone(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *UploadSmartContract) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UploadSmartContract) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UploadSmartContract) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Bytecode) > 0 {
		i -= len(m.Bytecode)
		copy(dAtA[i:], m.Bytecode)
		i = encodeVarintTurnstone(dAtA, i, uint64(len(m.Bytecode)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Message) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Message) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Message) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Action != nil {
		{
			size := m.Action.Size()
			i -= size
			if _, err := m.Action.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	if len(m.ChainID) > 0 {
		i -= len(m.ChainID)
		copy(dAtA[i:], m.ChainID)
		i = encodeVarintTurnstone(dAtA, i, uint64(len(m.ChainID)))
		i--
		dAtA[i] = 0x12
	}
	if m.TurnstoneID != 0 {
		i = encodeVarintTurnstone(dAtA, i, uint64(m.TurnstoneID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Message_SubmitLogicCall) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Message_SubmitLogicCall) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.SubmitLogicCall != nil {
		{
			size, err := m.SubmitLogicCall.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTurnstone(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}
func (m *Message_UpdateValset) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Message_UpdateValset) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.UpdateValset != nil {
		{
			size, err := m.UpdateValset.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTurnstone(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	return len(dAtA) - i, nil
}
func (m *Message_UploadSmartContract) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Message_UploadSmartContract) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.UploadSmartContract != nil {
		{
			size, err := m.UploadSmartContract.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTurnstone(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	return len(dAtA) - i, nil
}
func encodeVarintTurnstone(dAtA []byte, offset int, v uint64) int {
	offset -= sovTurnstone(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Valset) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.HexAddress) > 0 {
		for _, s := range m.HexAddress {
			l = len(s)
			n += 1 + l + sovTurnstone(uint64(l))
		}
	}
	if len(m.Powers) > 0 {
		l = 0
		for _, e := range m.Powers {
			l += sovTurnstone(uint64(e))
		}
		n += 1 + sovTurnstone(uint64(l)) + l
	}
	if m.ValsetID != 0 {
		n += 1 + sovTurnstone(uint64(m.ValsetID))
	}
	return n
}

func (m *SubmitLogicCall) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.HexContractAddress)
	if l > 0 {
		n += 1 + l + sovTurnstone(uint64(l))
	}
	l = len(m.Abi)
	if l > 0 {
		n += 1 + l + sovTurnstone(uint64(l))
	}
	l = len(m.Payload)
	if l > 0 {
		n += 1 + l + sovTurnstone(uint64(l))
	}
	if m.Deadline != 0 {
		n += 1 + sovTurnstone(uint64(m.Deadline))
	}
	return n
}

func (m *UpdateValset) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Valset != nil {
		l = m.Valset.Size()
		n += 1 + l + sovTurnstone(uint64(l))
	}
	return n
}

func (m *UploadSmartContract) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Bytecode)
	if l > 0 {
		n += 1 + l + sovTurnstone(uint64(l))
	}
	return n
}

func (m *Message) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TurnstoneID != 0 {
		n += 1 + sovTurnstone(uint64(m.TurnstoneID))
	}
	l = len(m.ChainID)
	if l > 0 {
		n += 1 + l + sovTurnstone(uint64(l))
	}
	if m.Action != nil {
		n += m.Action.Size()
	}
	return n
}

func (m *Message_SubmitLogicCall) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SubmitLogicCall != nil {
		l = m.SubmitLogicCall.Size()
		n += 1 + l + sovTurnstone(uint64(l))
	}
	return n
}
func (m *Message_UpdateValset) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.UpdateValset != nil {
		l = m.UpdateValset.Size()
		n += 1 + l + sovTurnstone(uint64(l))
	}
	return n
}
func (m *Message_UploadSmartContract) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.UploadSmartContract != nil {
		l = m.UploadSmartContract.Size()
		n += 1 + l + sovTurnstone(uint64(l))
	}
	return n
}

func sovTurnstone(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTurnstone(x uint64) (n int) {
	return sovTurnstone(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Valset) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTurnstone
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Valset: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Valset: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HexAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTurnstone
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTurnstone
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTurnstone
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HexAddress = append(m.HexAddress, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType == 0 {
				var v uint32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTurnstone
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Powers = append(m.Powers, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTurnstone
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthTurnstone
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthTurnstone
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.Powers) == 0 {
					m.Powers = make([]uint32, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTurnstone
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint32(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Powers = append(m.Powers, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Powers", wireType)
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValsetID", wireType)
			}
			m.ValsetID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTurnstone
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ValsetID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTurnstone(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTurnstone
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
func (m *SubmitLogicCall) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTurnstone
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SubmitLogicCall: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SubmitLogicCall: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HexContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTurnstone
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTurnstone
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTurnstone
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HexContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Abi", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTurnstone
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTurnstone
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTurnstone
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Abi = append(m.Abi[:0], dAtA[iNdEx:postIndex]...)
			if m.Abi == nil {
				m.Abi = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTurnstone
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTurnstone
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTurnstone
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payload = append(m.Payload[:0], dAtA[iNdEx:postIndex]...)
			if m.Payload == nil {
				m.Payload = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deadline", wireType)
			}
			m.Deadline = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTurnstone
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Deadline |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTurnstone(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTurnstone
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
func (m *UpdateValset) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTurnstone
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: UpdateValset: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UpdateValset: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Valset", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTurnstone
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTurnstone
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTurnstone
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Valset == nil {
				m.Valset = &Valset{}
			}
			if err := m.Valset.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTurnstone(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTurnstone
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
func (m *UploadSmartContract) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTurnstone
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: UploadSmartContract: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UploadSmartContract: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bytecode", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTurnstone
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTurnstone
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTurnstone
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bytecode = append(m.Bytecode[:0], dAtA[iNdEx:postIndex]...)
			if m.Bytecode == nil {
				m.Bytecode = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTurnstone(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTurnstone
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
func (m *Message) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTurnstone
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Message: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Message: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TurnstoneID", wireType)
			}
			m.TurnstoneID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTurnstone
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TurnstoneID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTurnstone
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTurnstone
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTurnstone
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubmitLogicCall", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTurnstone
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTurnstone
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTurnstone
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &SubmitLogicCall{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Action = &Message_SubmitLogicCall{v}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdateValset", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTurnstone
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTurnstone
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTurnstone
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &UpdateValset{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Action = &Message_UpdateValset{v}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UploadSmartContract", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTurnstone
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTurnstone
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTurnstone
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &UploadSmartContract{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Action = &Message_UploadSmartContract{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTurnstone(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTurnstone
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
func skipTurnstone(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTurnstone
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
					return 0, ErrIntOverflowTurnstone
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTurnstone
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
			if length < 0 {
				return 0, ErrInvalidLengthTurnstone
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTurnstone
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTurnstone
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTurnstone        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTurnstone          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTurnstone = fmt.Errorf("proto: unexpected end of group")
)