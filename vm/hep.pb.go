// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: vm/hep.proto

/*
	Package vm is a generated protocol buffer package.

	It is generated from these files:
		vm/hep.proto

	It has these top-level messages:
		HEP
*/
package vm

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

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

// HEP represents HEP packet
type HEP struct {
	Version   uint32 `protobuf:"varint,1,req,name=Version" json:"Version"`
	Protocol  uint32 `protobuf:"varint,2,req,name=Protocol" json:"Protocol"`
	SrcIP     string `protobuf:"bytes,3,req,name=SrcIP" json:"SrcIP"`
	DstIP     string `protobuf:"bytes,4,req,name=DstIP" json:"DstIP"`
	SrcPort   uint32 `protobuf:"varint,5,req,name=SrcPort" json:"SrcPort"`
	DstPort   uint32 `protobuf:"varint,6,req,name=DstPort" json:"DstPort"`
	Tsec      uint32 `protobuf:"varint,7,req,name=Tsec" json:"Tsec"`
	Tmsec     uint32 `protobuf:"varint,8,req,name=Tmsec" json:"Tmsec"`
	ProtoType uint32 `protobuf:"varint,9,req,name=ProtoType" json:"ProtoType"`
	NodeID    uint32 `protobuf:"varint,10,req,name=NodeID" json:"NodeID"`
	NodePW    string `protobuf:"bytes,11,req,name=NodePW" json:"NodePW"`
	Payload   string `protobuf:"bytes,12,req,name=Payload" json:"Payload"`
	CID       string `protobuf:"bytes,13,req,name=CID" json:"CID"`
	Vlan      uint32 `protobuf:"varint,14,req,name=Vlan" json:"Vlan"`
}

func (m *HEP) Reset()                    { *m = HEP{} }
func (*HEP) ProtoMessage()               {}
func (*HEP) Descriptor() ([]byte, []int) { return fileDescriptorHep, []int{0} }

func (m *HEP) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *HEP) GetProtocol() uint32 {
	if m != nil {
		return m.Protocol
	}
	return 0
}

func (m *HEP) GetSrcIP() string {
	if m != nil {
		return m.SrcIP
	}
	return ""
}

func (m *HEP) GetDstIP() string {
	if m != nil {
		return m.DstIP
	}
	return ""
}

func (m *HEP) GetSrcPort() uint32 {
	if m != nil {
		return m.SrcPort
	}
	return 0
}

func (m *HEP) GetDstPort() uint32 {
	if m != nil {
		return m.DstPort
	}
	return 0
}

func (m *HEP) GetTsec() uint32 {
	if m != nil {
		return m.Tsec
	}
	return 0
}

func (m *HEP) GetTmsec() uint32 {
	if m != nil {
		return m.Tmsec
	}
	return 0
}

func (m *HEP) GetProtoType() uint32 {
	if m != nil {
		return m.ProtoType
	}
	return 0
}

func (m *HEP) GetNodeID() uint32 {
	if m != nil {
		return m.NodeID
	}
	return 0
}

func (m *HEP) GetNodePW() string {
	if m != nil {
		return m.NodePW
	}
	return ""
}

func (m *HEP) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

func (m *HEP) GetCID() string {
	if m != nil {
		return m.CID
	}
	return ""
}

func (m *HEP) GetVlan() uint32 {
	if m != nil {
		return m.Vlan
	}
	return 0
}

func init() {
	proto.RegisterType((*HEP)(nil), "vm.HEP")
}
func (this *HEP) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*HEP)
	if !ok {
		that2, ok := that.(HEP)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Version != that1.Version {
		return false
	}
	if this.Protocol != that1.Protocol {
		return false
	}
	if this.SrcIP != that1.SrcIP {
		return false
	}
	if this.DstIP != that1.DstIP {
		return false
	}
	if this.SrcPort != that1.SrcPort {
		return false
	}
	if this.DstPort != that1.DstPort {
		return false
	}
	if this.Tsec != that1.Tsec {
		return false
	}
	if this.Tmsec != that1.Tmsec {
		return false
	}
	if this.ProtoType != that1.ProtoType {
		return false
	}
	if this.NodeID != that1.NodeID {
		return false
	}
	if this.NodePW != that1.NodePW {
		return false
	}
	if this.Payload != that1.Payload {
		return false
	}
	if this.CID != that1.CID {
		return false
	}
	if this.Vlan != that1.Vlan {
		return false
	}
	return true
}
func (this *HEP) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 18)
	s = append(s, "&vm.HEP{")
	s = append(s, "Version: "+fmt.Sprintf("%#v", this.Version)+",\n")
	s = append(s, "Protocol: "+fmt.Sprintf("%#v", this.Protocol)+",\n")
	s = append(s, "SrcIP: "+fmt.Sprintf("%#v", this.SrcIP)+",\n")
	s = append(s, "DstIP: "+fmt.Sprintf("%#v", this.DstIP)+",\n")
	s = append(s, "SrcPort: "+fmt.Sprintf("%#v", this.SrcPort)+",\n")
	s = append(s, "DstPort: "+fmt.Sprintf("%#v", this.DstPort)+",\n")
	s = append(s, "Tsec: "+fmt.Sprintf("%#v", this.Tsec)+",\n")
	s = append(s, "Tmsec: "+fmt.Sprintf("%#v", this.Tmsec)+",\n")
	s = append(s, "ProtoType: "+fmt.Sprintf("%#v", this.ProtoType)+",\n")
	s = append(s, "NodeID: "+fmt.Sprintf("%#v", this.NodeID)+",\n")
	s = append(s, "NodePW: "+fmt.Sprintf("%#v", this.NodePW)+",\n")
	s = append(s, "Payload: "+fmt.Sprintf("%#v", this.Payload)+",\n")
	s = append(s, "CID: "+fmt.Sprintf("%#v", this.CID)+",\n")
	s = append(s, "Vlan: "+fmt.Sprintf("%#v", this.Vlan)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringHep(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *HEP) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HEP) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintHep(dAtA, i, uint64(m.Version))
	dAtA[i] = 0x10
	i++
	i = encodeVarintHep(dAtA, i, uint64(m.Protocol))
	dAtA[i] = 0x1a
	i++
	i = encodeVarintHep(dAtA, i, uint64(len(m.SrcIP)))
	i += copy(dAtA[i:], m.SrcIP)
	dAtA[i] = 0x22
	i++
	i = encodeVarintHep(dAtA, i, uint64(len(m.DstIP)))
	i += copy(dAtA[i:], m.DstIP)
	dAtA[i] = 0x28
	i++
	i = encodeVarintHep(dAtA, i, uint64(m.SrcPort))
	dAtA[i] = 0x30
	i++
	i = encodeVarintHep(dAtA, i, uint64(m.DstPort))
	dAtA[i] = 0x38
	i++
	i = encodeVarintHep(dAtA, i, uint64(m.Tsec))
	dAtA[i] = 0x40
	i++
	i = encodeVarintHep(dAtA, i, uint64(m.Tmsec))
	dAtA[i] = 0x48
	i++
	i = encodeVarintHep(dAtA, i, uint64(m.ProtoType))
	dAtA[i] = 0x50
	i++
	i = encodeVarintHep(dAtA, i, uint64(m.NodeID))
	dAtA[i] = 0x5a
	i++
	i = encodeVarintHep(dAtA, i, uint64(len(m.NodePW)))
	i += copy(dAtA[i:], m.NodePW)
	dAtA[i] = 0x62
	i++
	i = encodeVarintHep(dAtA, i, uint64(len(m.Payload)))
	i += copy(dAtA[i:], m.Payload)
	dAtA[i] = 0x6a
	i++
	i = encodeVarintHep(dAtA, i, uint64(len(m.CID)))
	i += copy(dAtA[i:], m.CID)
	dAtA[i] = 0x70
	i++
	i = encodeVarintHep(dAtA, i, uint64(m.Vlan))
	return i, nil
}

func encodeVarintHep(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedHEP(r randyHep, easy bool) *HEP {
	this := &HEP{}
	this.Version = uint32(r.Uint32())
	this.Protocol = uint32(r.Uint32())
	this.SrcIP = string(randStringHep(r))
	this.DstIP = string(randStringHep(r))
	this.SrcPort = uint32(r.Uint32())
	this.DstPort = uint32(r.Uint32())
	this.Tsec = uint32(r.Uint32())
	this.Tmsec = uint32(r.Uint32())
	this.ProtoType = uint32(r.Uint32())
	this.NodeID = uint32(r.Uint32())
	this.NodePW = string(randStringHep(r))
	this.Payload = string(randStringHep(r))
	this.CID = string(randStringHep(r))
	this.Vlan = uint32(r.Uint32())
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyHep interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneHep(r randyHep) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringHep(r randyHep) string {
	v1 := r.Intn(100)
	tmps := make([]rune, v1)
	for i := 0; i < v1; i++ {
		tmps[i] = randUTF8RuneHep(r)
	}
	return string(tmps)
}
func randUnrecognizedHep(r randyHep, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldHep(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldHep(dAtA []byte, r randyHep, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateHep(dAtA, uint64(key))
		v2 := r.Int63()
		if r.Intn(2) == 0 {
			v2 *= -1
		}
		dAtA = encodeVarintPopulateHep(dAtA, uint64(v2))
	case 1:
		dAtA = encodeVarintPopulateHep(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateHep(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateHep(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateHep(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateHep(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *HEP) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovHep(uint64(m.Version))
	n += 1 + sovHep(uint64(m.Protocol))
	l = len(m.SrcIP)
	n += 1 + l + sovHep(uint64(l))
	l = len(m.DstIP)
	n += 1 + l + sovHep(uint64(l))
	n += 1 + sovHep(uint64(m.SrcPort))
	n += 1 + sovHep(uint64(m.DstPort))
	n += 1 + sovHep(uint64(m.Tsec))
	n += 1 + sovHep(uint64(m.Tmsec))
	n += 1 + sovHep(uint64(m.ProtoType))
	n += 1 + sovHep(uint64(m.NodeID))
	l = len(m.NodePW)
	n += 1 + l + sovHep(uint64(l))
	l = len(m.Payload)
	n += 1 + l + sovHep(uint64(l))
	l = len(m.CID)
	n += 1 + l + sovHep(uint64(l))
	n += 1 + sovHep(uint64(m.Vlan))
	return n
}

func sovHep(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozHep(x uint64) (n int) {
	return sovHep(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *HEP) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&HEP{`,
		`Version:` + fmt.Sprintf("%v", this.Version) + `,`,
		`Protocol:` + fmt.Sprintf("%v", this.Protocol) + `,`,
		`SrcIP:` + fmt.Sprintf("%v", this.SrcIP) + `,`,
		`DstIP:` + fmt.Sprintf("%v", this.DstIP) + `,`,
		`SrcPort:` + fmt.Sprintf("%v", this.SrcPort) + `,`,
		`DstPort:` + fmt.Sprintf("%v", this.DstPort) + `,`,
		`Tsec:` + fmt.Sprintf("%v", this.Tsec) + `,`,
		`Tmsec:` + fmt.Sprintf("%v", this.Tmsec) + `,`,
		`ProtoType:` + fmt.Sprintf("%v", this.ProtoType) + `,`,
		`NodeID:` + fmt.Sprintf("%v", this.NodeID) + `,`,
		`NodePW:` + fmt.Sprintf("%v", this.NodePW) + `,`,
		`Payload:` + fmt.Sprintf("%v", this.Payload) + `,`,
		`CID:` + fmt.Sprintf("%v", this.CID) + `,`,
		`Vlan:` + fmt.Sprintf("%v", this.Vlan) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringHep(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *HEP) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHep
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
			return fmt.Errorf("proto: HEP: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HEP: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHep
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Protocol", wireType)
			}
			m.Protocol = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHep
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Protocol |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			hasFields[0] |= uint64(0x00000002)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SrcIP", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHep
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
				return ErrInvalidLengthHep
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SrcIP = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000004)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DstIP", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHep
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
				return ErrInvalidLengthHep
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DstIP = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000008)
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SrcPort", wireType)
			}
			m.SrcPort = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHep
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SrcPort |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			hasFields[0] |= uint64(0x00000010)
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DstPort", wireType)
			}
			m.DstPort = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHep
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DstPort |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			hasFields[0] |= uint64(0x00000020)
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tsec", wireType)
			}
			m.Tsec = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHep
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Tsec |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			hasFields[0] |= uint64(0x00000040)
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tmsec", wireType)
			}
			m.Tmsec = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHep
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Tmsec |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			hasFields[0] |= uint64(0x00000080)
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProtoType", wireType)
			}
			m.ProtoType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHep
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProtoType |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			hasFields[0] |= uint64(0x00000100)
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeID", wireType)
			}
			m.NodeID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHep
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NodeID |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			hasFields[0] |= uint64(0x00000200)
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodePW", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHep
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
				return ErrInvalidLengthHep
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NodePW = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000400)
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHep
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
				return ErrInvalidLengthHep
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payload = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000800)
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHep
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
				return ErrInvalidLengthHep
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00001000)
		case 14:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vlan", wireType)
			}
			m.Vlan = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHep
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Vlan |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			hasFields[0] |= uint64(0x00002000)
		default:
			iNdEx = preIndex
			skippy, err := skipHep(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHep
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return proto.NewRequiredNotSetError("Version")
	}
	if hasFields[0]&uint64(0x00000002) == 0 {
		return proto.NewRequiredNotSetError("Protocol")
	}
	if hasFields[0]&uint64(0x00000004) == 0 {
		return proto.NewRequiredNotSetError("SrcIP")
	}
	if hasFields[0]&uint64(0x00000008) == 0 {
		return proto.NewRequiredNotSetError("DstIP")
	}
	if hasFields[0]&uint64(0x00000010) == 0 {
		return proto.NewRequiredNotSetError("SrcPort")
	}
	if hasFields[0]&uint64(0x00000020) == 0 {
		return proto.NewRequiredNotSetError("DstPort")
	}
	if hasFields[0]&uint64(0x00000040) == 0 {
		return proto.NewRequiredNotSetError("Tsec")
	}
	if hasFields[0]&uint64(0x00000080) == 0 {
		return proto.NewRequiredNotSetError("Tmsec")
	}
	if hasFields[0]&uint64(0x00000100) == 0 {
		return proto.NewRequiredNotSetError("ProtoType")
	}
	if hasFields[0]&uint64(0x00000200) == 0 {
		return proto.NewRequiredNotSetError("NodeID")
	}
	if hasFields[0]&uint64(0x00000400) == 0 {
		return proto.NewRequiredNotSetError("NodePW")
	}
	if hasFields[0]&uint64(0x00000800) == 0 {
		return proto.NewRequiredNotSetError("Payload")
	}
	if hasFields[0]&uint64(0x00001000) == 0 {
		return proto.NewRequiredNotSetError("CID")
	}
	if hasFields[0]&uint64(0x00002000) == 0 {
		return proto.NewRequiredNotSetError("Vlan")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipHep(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHep
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
					return 0, ErrIntOverflowHep
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
					return 0, ErrIntOverflowHep
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
				return 0, ErrInvalidLengthHep
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowHep
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
				next, err := skipHep(dAtA[start:])
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
	ErrInvalidLengthHep = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHep   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("vm/hep.proto", fileDescriptorHep) }

var fileDescriptorHep = []byte{
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0xcf, 0x3d, 0x4e, 0x02, 0x41,
	0x14, 0xc0, 0x71, 0x86, 0xe5, 0x73, 0x14, 0x13, 0xa7, 0x30, 0x2f, 0xc4, 0x3c, 0x09, 0x15, 0x85,
	0xc2, 0x1d, 0x70, 0x4d, 0xdc, 0xc6, 0x4c, 0x84, 0x60, 0x0d, 0xcb, 0x0a, 0x24, 0x0b, 0xb3, 0xd9,
	0x5d, 0x0a, 0x3a, 0x8f, 0xe0, 0x2d, 0xf4, 0x08, 0x96, 0x96, 0x94, 0x96, 0x56, 0x46, 0xc6, 0x0b,
	0x58, 0x5a, 0x9a, 0x19, 0x76, 0xf5, 0x75, 0xbb, 0xbf, 0xff, 0x4c, 0xde, 0x3c, 0x7e, 0x1c, 0xad,
	0x27, 0xe1, 0x22, 0x99, 0xf7, 0xe6, 0x41, 0xd4, 0x8d, 0x62, 0x95, 0x2a, 0x51, 0xcd, 0xa8, 0x79,
	0x31, 0x5b, 0xa4, 0xf3, 0xf5, 0xa4, 0xeb, 0xab, 0x65, 0x6f, 0xa6, 0x66, 0xaa, 0x67, 0xfb, 0x64,
	0x7d, 0x6f, 0xff, 0xec, 0x8f, 0xfd, 0xda, 0xdf, 0x6b, 0x3f, 0x39, 0xdc, 0xb9, 0xbe, 0x92, 0x02,
	0x79, 0x75, 0x14, 0xc4, 0xc9, 0x42, 0xad, 0x80, 0xb5, 0x8a, 0x9d, 0x46, 0xbf, 0xb4, 0xfd, 0x38,
	0x2b, 0xdc, 0xe6, 0x28, 0x5a, 0xbc, 0x26, 0xcd, 0x05, 0x5f, 0x85, 0x50, 0x24, 0x07, 0xfe, 0x54,
	0x34, 0x79, 0x79, 0x10, 0xfb, 0x9e, 0x04, 0xa7, 0x55, 0xec, 0xd4, 0xb3, 0xbc, 0x27, 0xd3, 0xdc,
	0x24, 0xf5, 0x24, 0x94, 0x68, 0xb3, 0x64, 0x26, 0x0f, 0x62, 0x5f, 0xaa, 0x38, 0x85, 0x32, 0x9d,
	0x9c, 0xa1, 0xe9, 0x6e, 0x92, 0xda, 0x5e, 0xa1, 0x3d, 0x43, 0x01, 0xbc, 0x34, 0x4c, 0x02, 0x1f,
	0xaa, 0x24, 0x5a, 0x31, 0x53, 0x87, 0x4b, 0x93, 0x6a, 0x24, 0xed, 0x49, 0xb4, 0x79, 0xdd, 0xbe,
	0x7c, 0xb8, 0x89, 0x02, 0xa8, 0x93, 0xfe, 0xcf, 0xe2, 0x94, 0x57, 0x6e, 0xd4, 0x34, 0xf0, 0x5c,
	0xe0, 0xe4, 0x40, 0x66, 0x79, 0x95, 0x77, 0x70, 0x40, 0x96, 0xca, 0xcc, 0xbc, 0x5a, 0x8e, 0x37,
	0xa1, 0x1a, 0x4f, 0xe1, 0x90, 0xe4, 0x1c, 0xc5, 0x09, 0x77, 0x2e, 0x3d, 0x17, 0x1a, 0xa4, 0x19,
	0x30, 0xdb, 0x8c, 0xc2, 0xf1, 0x0a, 0x8e, 0xe8, 0x36, 0x46, 0xfa, 0xe7, 0xef, 0x3b, 0x2c, 0x7c,
	0xef, 0x90, 0xfd, 0xec, 0x90, 0x3d, 0x68, 0x64, 0xcf, 0x1a, 0xd9, 0x8b, 0x46, 0xf6, 0xaa, 0x91,
	0x6d, 0x35, 0xb2, 0x37, 0x8d, 0xec, 0x53, 0x23, 0x7b, 0xfc, 0xc2, 0xc2, 0x6f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xdd, 0x5d, 0x6b, 0xc3, 0x23, 0x02, 0x00, 0x00,
}
