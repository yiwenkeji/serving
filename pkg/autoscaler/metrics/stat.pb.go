/*
Copyright 2020 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pkg/autoscaler/metrics/stat.proto

package metrics

import (
	encoding_binary "encoding/binary"
	fmt "fmt"
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

// Stat defines a single measurement at a point in time
type Stat struct {
	// The unique identity of this pod.  Used to count how many pods
	// are contributing to the metrics.
	PodName string `protobuf:"bytes,1,opt,name=pod_name,json=podName,proto3" json:"pod_name,omitempty"`
	// Average number of requests currently being handled by this pod.
	AverageConcurrentRequests float64 `protobuf:"fixed64,2,opt,name=average_concurrent_requests,json=averageConcurrentRequests,proto3" json:"average_concurrent_requests,omitempty"`
	// Part of AverageConcurrentRequests, for requests going through a proxy.
	AverageProxiedConcurrentRequests float64 `protobuf:"fixed64,3,opt,name=average_proxied_concurrent_requests,json=averageProxiedConcurrentRequests,proto3" json:"average_proxied_concurrent_requests,omitempty"`
	// Number of requests received since last Stat (approximately requests per second).
	RequestCount float64 `protobuf:"fixed64,4,opt,name=request_count,json=requestCount,proto3" json:"request_count,omitempty"`
	// Part of RequestCount, for requests going through a proxy.
	ProxiedRequestCount float64 `protobuf:"fixed64,5,opt,name=proxied_request_count,json=proxiedRequestCount,proto3" json:"proxied_request_count,omitempty"`
	// Process uptime in seconds.
	ProcessUptime float64 `protobuf:"fixed64,6,opt,name=process_uptime,json=processUptime,proto3" json:"process_uptime,omitempty"`
}

func (m *Stat) Reset()         { *m = Stat{} }
func (m *Stat) String() string { return proto.CompactTextString(m) }
func (*Stat) ProtoMessage()    {}
func (*Stat) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf216df9f6fff44c, []int{0}
}
func (m *Stat) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Stat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Stat.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Stat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Stat.Merge(m, src)
}
func (m *Stat) XXX_Size() int {
	return m.Size()
}
func (m *Stat) XXX_DiscardUnknown() {
	xxx_messageInfo_Stat.DiscardUnknown(m)
}

var xxx_messageInfo_Stat proto.InternalMessageInfo

func (m *Stat) GetPodName() string {
	if m != nil {
		return m.PodName
	}
	return ""
}

func (m *Stat) GetAverageConcurrentRequests() float64 {
	if m != nil {
		return m.AverageConcurrentRequests
	}
	return 0
}

func (m *Stat) GetAverageProxiedConcurrentRequests() float64 {
	if m != nil {
		return m.AverageProxiedConcurrentRequests
	}
	return 0
}

func (m *Stat) GetRequestCount() float64 {
	if m != nil {
		return m.RequestCount
	}
	return 0
}

func (m *Stat) GetProxiedRequestCount() float64 {
	if m != nil {
		return m.ProxiedRequestCount
	}
	return 0
}

func (m *Stat) GetProcessUptime() float64 {
	if m != nil {
		return m.ProcessUptime
	}
	return 0
}

// WireStatMessage is a copy of the StatMessage Golang type, exploding the fields of
// `types.NamespacedName` to make it compatible with protobufs.
type WireStatMessage struct {
	// Namespace is the namespace of the entity the stat belongs to.
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// Name is the name of the entity the stat belongs to.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Stat are the actual metrics being sent.
	Stat *Stat `protobuf:"bytes,3,opt,name=stat,proto3" json:"stat,omitempty"`
}

func (m *WireStatMessage) Reset()         { *m = WireStatMessage{} }
func (m *WireStatMessage) String() string { return proto.CompactTextString(m) }
func (*WireStatMessage) ProtoMessage()    {}
func (*WireStatMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf216df9f6fff44c, []int{1}
}
func (m *WireStatMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WireStatMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WireStatMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WireStatMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WireStatMessage.Merge(m, src)
}
func (m *WireStatMessage) XXX_Size() int {
	return m.Size()
}
func (m *WireStatMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_WireStatMessage.DiscardUnknown(m)
}

var xxx_messageInfo_WireStatMessage proto.InternalMessageInfo

func (m *WireStatMessage) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *WireStatMessage) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *WireStatMessage) GetStat() *Stat {
	if m != nil {
		return m.Stat
	}
	return nil
}

// WireStatMessages is a wrapper to send multiple WireStatMessages at once.
type WireStatMessages struct {
	// Messages is a list of WireStatMessages.
	Messages []*WireStatMessage `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (m *WireStatMessages) Reset()         { *m = WireStatMessages{} }
func (m *WireStatMessages) String() string { return proto.CompactTextString(m) }
func (*WireStatMessages) ProtoMessage()    {}
func (*WireStatMessages) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf216df9f6fff44c, []int{2}
}
func (m *WireStatMessages) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WireStatMessages) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WireStatMessages.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WireStatMessages) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WireStatMessages.Merge(m, src)
}
func (m *WireStatMessages) XXX_Size() int {
	return m.Size()
}
func (m *WireStatMessages) XXX_DiscardUnknown() {
	xxx_messageInfo_WireStatMessages.DiscardUnknown(m)
}

var xxx_messageInfo_WireStatMessages proto.InternalMessageInfo

func (m *WireStatMessages) GetMessages() []*WireStatMessage {
	if m != nil {
		return m.Messages
	}
	return nil
}

func init() {
	proto.RegisterType((*Stat)(nil), "metrics.Stat")
	proto.RegisterType((*WireStatMessage)(nil), "metrics.WireStatMessage")
	proto.RegisterType((*WireStatMessages)(nil), "metrics.WireStatMessages")
}

func init() { proto.RegisterFile("pkg/autoscaler/metrics/stat.proto", fileDescriptor_cf216df9f6fff44c) }

var fileDescriptor_cf216df9f6fff44c = []byte{
	// 349 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xcf, 0x4e, 0xf2, 0x40,
	0x14, 0xc5, 0x19, 0xe8, 0xc7, 0x9f, 0xcb, 0x87, 0x9a, 0x31, 0x26, 0x25, 0x9a, 0xa6, 0x40, 0x4c,
	0x58, 0x41, 0x82, 0xae, 0x5d, 0xc8, 0xc6, 0x0d, 0xc6, 0xd4, 0x18, 0x97, 0xcd, 0x38, 0x5c, 0x49,
	0xa3, 0xed, 0x8c, 0x33, 0x53, 0xe3, 0x63, 0xf8, 0x28, 0x3e, 0x86, 0x4b, 0x96, 0x2e, 0x0d, 0xbc,
	0x88, 0xe9, 0x38, 0xa0, 0x12, 0x76, 0x93, 0x73, 0x7f, 0xe7, 0xdc, 0xf4, 0xdc, 0x42, 0x47, 0x3e,
	0xcc, 0x86, 0x2c, 0x37, 0x42, 0x73, 0xf6, 0x88, 0x6a, 0x98, 0xa2, 0x51, 0x09, 0xd7, 0x43, 0x6d,
	0x98, 0x19, 0x48, 0x25, 0x8c, 0xa0, 0x35, 0xa7, 0x75, 0xdf, 0xca, 0xe0, 0x5d, 0x1b, 0x66, 0x68,
	0x1b, 0xea, 0x52, 0x4c, 0xe3, 0x8c, 0xa5, 0xe8, 0x93, 0x90, 0xf4, 0x1b, 0x51, 0x4d, 0x8a, 0xe9,
	0x25, 0x4b, 0x91, 0x9e, 0xc1, 0x21, 0x7b, 0x46, 0xc5, 0x66, 0x18, 0x73, 0x91, 0xf1, 0x5c, 0x29,
	0xcc, 0x4c, 0xac, 0xf0, 0x29, 0x47, 0x6d, 0xb4, 0x5f, 0x0e, 0x49, 0x9f, 0x44, 0x6d, 0x87, 0x8c,
	0xd7, 0x44, 0xe4, 0x00, 0x3a, 0x81, 0xde, 0xca, 0x2f, 0x95, 0x78, 0x49, 0x70, 0xba, 0x35, 0xa7,
	0x62, 0x73, 0x42, 0x87, 0x5e, 0x7d, 0x93, 0x5b, 0xe2, 0x7a, 0xd0, 0x72, 0x9e, 0x98, 0x8b, 0x3c,
	0x33, 0xbe, 0x67, 0x8d, 0xff, 0x9d, 0x38, 0x2e, 0x34, 0x3a, 0x82, 0x83, 0xd5, 0xae, 0xbf, 0xf0,
	0x3f, 0x0b, 0xef, 0xbb, 0x61, 0xf4, 0xdb, 0x73, 0x0c, 0x3b, 0x52, 0x09, 0x8e, 0x5a, 0xc7, 0xb9,
	0x34, 0x49, 0x8a, 0x7e, 0xd5, 0xc2, 0x2d, 0xa7, 0xde, 0x58, 0xb1, 0x7b, 0x0f, 0xbb, 0xb7, 0x89,
	0xc2, 0xa2, 0xb5, 0x09, 0x6a, 0xcd, 0x66, 0x48, 0x8f, 0xa0, 0x51, 0x14, 0xa7, 0x25, 0xe3, 0xab,
	0xf6, 0x7e, 0x04, 0x4a, 0xc1, 0xb3, 0xb5, 0x96, 0xed, 0xc0, 0xbe, 0x69, 0x07, 0xbc, 0xe2, 0x1c,
	0xf6, 0xa3, 0x9b, 0xa3, 0xd6, 0xc0, 0xdd, 0x63, 0x50, 0xa4, 0x46, 0x76, 0xd4, 0xbd, 0x80, 0xbd,
	0x8d, 0x3d, 0x9a, 0x9e, 0x42, 0x3d, 0x75, 0x6f, 0x9f, 0x84, 0x95, 0x7e, 0x73, 0xe4, 0xaf, 0xad,
	0x1b, 0x70, 0xb4, 0x26, 0xcf, 0xfd, 0xf7, 0x45, 0x40, 0xe6, 0x8b, 0x80, 0x7c, 0x2e, 0x02, 0xf2,
	0xba, 0x0c, 0x4a, 0xf3, 0x65, 0x50, 0xfa, 0x58, 0x06, 0xa5, 0xbb, 0xaa, 0xfd, 0x1d, 0x4e, 0xbe,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x26, 0xd8, 0x3d, 0xa8, 0x33, 0x02, 0x00, 0x00,
}

func (m *Stat) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Stat) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Stat) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ProcessUptime != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.ProcessUptime))))
		i--
		dAtA[i] = 0x31
	}
	if m.ProxiedRequestCount != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.ProxiedRequestCount))))
		i--
		dAtA[i] = 0x29
	}
	if m.RequestCount != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.RequestCount))))
		i--
		dAtA[i] = 0x21
	}
	if m.AverageProxiedConcurrentRequests != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.AverageProxiedConcurrentRequests))))
		i--
		dAtA[i] = 0x19
	}
	if m.AverageConcurrentRequests != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.AverageConcurrentRequests))))
		i--
		dAtA[i] = 0x11
	}
	if len(m.PodName) > 0 {
		i -= len(m.PodName)
		copy(dAtA[i:], m.PodName)
		i = encodeVarintStat(dAtA, i, uint64(len(m.PodName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *WireStatMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WireStatMessage) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WireStatMessage) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Stat != nil {
		{
			size, err := m.Stat.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintStat(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintStat(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Namespace) > 0 {
		i -= len(m.Namespace)
		copy(dAtA[i:], m.Namespace)
		i = encodeVarintStat(dAtA, i, uint64(len(m.Namespace)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *WireStatMessages) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WireStatMessages) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WireStatMessages) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Messages) > 0 {
		for iNdEx := len(m.Messages) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Messages[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintStat(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintStat(dAtA []byte, offset int, v uint64) int {
	offset -= sovStat(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Stat) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PodName)
	if l > 0 {
		n += 1 + l + sovStat(uint64(l))
	}
	if m.AverageConcurrentRequests != 0 {
		n += 9
	}
	if m.AverageProxiedConcurrentRequests != 0 {
		n += 9
	}
	if m.RequestCount != 0 {
		n += 9
	}
	if m.ProxiedRequestCount != 0 {
		n += 9
	}
	if m.ProcessUptime != 0 {
		n += 9
	}
	return n
}

func (m *WireStatMessage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Namespace)
	if l > 0 {
		n += 1 + l + sovStat(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovStat(uint64(l))
	}
	if m.Stat != nil {
		l = m.Stat.Size()
		n += 1 + l + sovStat(uint64(l))
	}
	return n
}

func (m *WireStatMessages) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Messages) > 0 {
		for _, e := range m.Messages {
			l = e.Size()
			n += 1 + l + sovStat(uint64(l))
		}
	}
	return n
}

func sovStat(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStat(x uint64) (n int) {
	return sovStat(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Stat) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStat
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
			return fmt.Errorf("proto: Stat: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Stat: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PodName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStat
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
				return ErrInvalidLengthStat
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStat
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PodName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field AverageConcurrentRequests", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.AverageConcurrentRequests = float64(math.Float64frombits(v))
		case 3:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field AverageProxiedConcurrentRequests", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.AverageProxiedConcurrentRequests = float64(math.Float64frombits(v))
		case 4:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestCount", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.RequestCount = float64(math.Float64frombits(v))
		case 5:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProxiedRequestCount", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.ProxiedRequestCount = float64(math.Float64frombits(v))
		case 6:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProcessUptime", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.ProcessUptime = float64(math.Float64frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skipStat(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStat
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthStat
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
func (m *WireStatMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStat
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
			return fmt.Errorf("proto: WireStatMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WireStatMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStat
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
				return ErrInvalidLengthStat
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStat
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Namespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStat
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
				return ErrInvalidLengthStat
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStat
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stat", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStat
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
				return ErrInvalidLengthStat
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStat
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Stat == nil {
				m.Stat = &Stat{}
			}
			if err := m.Stat.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStat(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStat
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthStat
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
func (m *WireStatMessages) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStat
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
			return fmt.Errorf("proto: WireStatMessages: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WireStatMessages: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Messages", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStat
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
				return ErrInvalidLengthStat
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStat
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Messages = append(m.Messages, &WireStatMessage{})
			if err := m.Messages[len(m.Messages)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStat(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStat
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthStat
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
func skipStat(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStat
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
					return 0, ErrIntOverflowStat
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
					return 0, ErrIntOverflowStat
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
				return 0, ErrInvalidLengthStat
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStat
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStat
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStat        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStat          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStat = fmt.Errorf("proto: unexpected end of group")
)