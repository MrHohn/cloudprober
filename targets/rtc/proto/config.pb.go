// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/google/cloudprober/targets/rtc/proto/config.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// TargetsConf represents listing targets from a key/val set stored in the RTC
// API under a certain provided configuration. It always only lists the keys
// (hostnames) and provides a resolver which will resolve to a single tag. This
// can be used as a kind of drop-in replacement for GCE Instances when that
// would not be possible. For example, one could set "resolve_tag" to "ip", and
// this lister will now list all VM names, with a resolver that resolves those
// names to their ip addresses.
//
// For more information on RTC see
// https://cloud.google.com/deployment-manager/runtime-configurator/reference/rest/
//
// It is assumed that this is being curated, and indeed can be controlled with
// settings in a cloudprober configuration.  See
// cloudprober/config/config.proto.
type TargetsConf struct {
	// Config-name to list hosts from.
	Cfg *string `protobuf:"bytes,1,opt,name=cfg" json:"cfg,omitempty"`
	// Expiration time in ms for RTC variables. Variables older will be ignored.
	ExpireMsec *int32 `protobuf:"varint,2,opt,name=expire_msec,json=expireMsec,def=30000" json:"expire_msec,omitempty"`
	// Which address tag to resolve a hostname to. Hosts report multiple addresses
	// for themselves, each with a "tag" (like "PUBLIC_IP" or "PRIVATE_IP"). For
	// more information, see the "tag" field of the RtcTargetInfo.Address message.
	ResolveTag *string `protobuf:"bytes,3,opt,name=resolve_tag,json=resolveTag" json:"resolve_tag,omitempty"`
	// groups is a list of groups that should be included by the lister. Hosts can
	// assign themselves a set "groups" such as "Zone1" or "HasMultiNIC". If this
	// field is set, then we will filter out hosts that don't have a group that
	// matches any groups listed here. If groups is not set, no filtering will
	// occure.
	//
	// For more information, see the "groups" field of the RtcTargetInfo message.
	//
	// Example :
	//   Host1 has group "G1" and "G2"
	//   Host2 has groups "G2" and "G3"
	//   Host3 has no groups.
	//   The following table shows values for groups and which hosts will be
	//   listed.
	//   groups      | Listed Hosts
	//   ------------|--------------------
	//   ["G1"]      | Host1
	//   ["G2"]      | Host1, Host2
	//   ["G1","G3"] | Host1, Host2
	//   ["G4"]      | None
	//   []          | Host1, Host2, Host3
	Groups []string `protobuf:"bytes,4,rep,name=groups" json:"groups,omitempty"`
	// How often targets should be evaluated. Any number less than or equal to 0
	// will result in no target caching (targets will be reevaluated on demand).
	ReEvalSec *int32 `protobuf:"varint,6,opt,name=re_eval_sec,json=reEvalSec,def=0" json:"re_eval_sec,omitempty"`
	// The alphanumeric ID of the project which this instance is member of
	// (not project number)
	// e.g. bbmc-cloudprober-instance-cxkdsa123
	ProjectId            []string `protobuf:"bytes,7,rep,name=project_id,json=projectId" json:"project_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TargetsConf) Reset()         { *m = TargetsConf{} }
func (m *TargetsConf) String() string { return proto.CompactTextString(m) }
func (*TargetsConf) ProtoMessage()    {}
func (*TargetsConf) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_740388aebedb8f5c, []int{0}
}
func (m *TargetsConf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TargetsConf.Unmarshal(m, b)
}
func (m *TargetsConf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TargetsConf.Marshal(b, m, deterministic)
}
func (dst *TargetsConf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TargetsConf.Merge(dst, src)
}
func (m *TargetsConf) XXX_Size() int {
	return xxx_messageInfo_TargetsConf.Size(m)
}
func (m *TargetsConf) XXX_DiscardUnknown() {
	xxx_messageInfo_TargetsConf.DiscardUnknown(m)
}

var xxx_messageInfo_TargetsConf proto.InternalMessageInfo

const Default_TargetsConf_ExpireMsec int32 = 30000
const Default_TargetsConf_ReEvalSec int32 = 0

func (m *TargetsConf) GetCfg() string {
	if m != nil && m.Cfg != nil {
		return *m.Cfg
	}
	return ""
}

func (m *TargetsConf) GetExpireMsec() int32 {
	if m != nil && m.ExpireMsec != nil {
		return *m.ExpireMsec
	}
	return Default_TargetsConf_ExpireMsec
}

func (m *TargetsConf) GetResolveTag() string {
	if m != nil && m.ResolveTag != nil {
		return *m.ResolveTag
	}
	return ""
}

func (m *TargetsConf) GetGroups() []string {
	if m != nil {
		return m.Groups
	}
	return nil
}

func (m *TargetsConf) GetReEvalSec() int32 {
	if m != nil && m.ReEvalSec != nil {
		return *m.ReEvalSec
	}
	return Default_TargetsConf_ReEvalSec
}

func (m *TargetsConf) GetProjectId() []string {
	if m != nil {
		return m.ProjectId
	}
	return nil
}

func init() {
	proto.RegisterType((*TargetsConf)(nil), "cloudprober.targets.rtc.TargetsConf")
}

func init() {
	proto.RegisterFile("github.com/google/cloudprober/targets/rtc/proto/config.proto", fileDescriptor_config_740388aebedb8f5c)
}

var fileDescriptor_config_740388aebedb8f5c = []byte{
	// 236 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8d, 0xbf, 0x4e, 0x84, 0x40,
	0x10, 0xc6, 0x83, 0x78, 0x67, 0x18, 0x1a, 0xb3, 0x85, 0xd2, 0x18, 0xd1, 0xc2, 0x50, 0xb1, 0x24,
	0x76, 0xc6, 0xce, 0x58, 0x58, 0xd8, 0xe0, 0xf5, 0x84, 0x1b, 0x86, 0x15, 0xb3, 0x77, 0xb3, 0x99,
	0x5d, 0x88, 0xcf, 0xe7, 0x93, 0x19, 0x38, 0x8a, 0xeb, 0xe6, 0xf7, 0xcd, 0xf7, 0x07, 0x5e, 0xcd,
	0x10, 0xbe, 0xc7, 0x7d, 0x89, 0x7c, 0xd0, 0x86, 0xd9, 0x58, 0xd2, 0x68, 0x79, 0xec, 0x9c, 0xf0,
	0x9e, 0x44, 0x87, 0x56, 0x0c, 0x05, 0xaf, 0x25, 0xa0, 0x76, 0xc2, 0x81, 0x35, 0xf2, 0xb1, 0x1f,
	0x4c, 0xb9, 0x80, 0xba, 0x3d, 0xf3, 0x96, 0xab, 0xb7, 0x94, 0x80, 0x8f, 0x7f, 0x11, 0xa4, 0xbb,
	0x13, 0xbf, 0xf1, 0xb1, 0x57, 0xd7, 0x10, 0x63, 0x6f, 0xb2, 0x28, 0x8f, 0x8a, 0xa4, 0x9e, 0x4f,
	0xf5, 0x04, 0x29, 0xfd, 0xba, 0x41, 0xa8, 0x39, 0x78, 0xc2, 0xec, 0x22, 0x8f, 0x8a, 0xcd, 0xcb,
	0xe6, 0xb9, 0xaa, 0xaa, 0xaa, 0x86, 0xd3, 0xe7, 0xd3, 0x13, 0xaa, 0x7b, 0x48, 0x85, 0x3c, 0xdb,
	0x89, 0x9a, 0xd0, 0x9a, 0x2c, 0x5e, 0x1a, 0x60, 0x95, 0x76, 0xad, 0x51, 0x37, 0xb0, 0x35, 0xc2,
	0xa3, 0xf3, 0xd9, 0x65, 0x1e, 0x17, 0x49, 0xbd, 0x92, 0x7a, 0x98, 0x83, 0x0d, 0x4d, 0xad, 0x6d,
	0xe6, 0x81, 0xed, 0x32, 0x10, 0x55, 0x75, 0x22, 0xf4, 0x3e, 0xb5, 0xf6, 0x8b, 0x50, 0xdd, 0x01,
	0x38, 0xe1, 0x1f, 0xc2, 0xd0, 0x0c, 0x5d, 0x76, 0xb5, 0xc4, 0x93, 0x55, 0xf9, 0xe8, 0xfe, 0x03,
	0x00, 0x00, 0xff, 0xff, 0x1c, 0x0a, 0xc3, 0x47, 0x1c, 0x01, 0x00, 0x00,
}
