// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/common/v1/account.proto

package v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Account struct {
	Revision             int64               `protobuf:"varint,1,opt,name=revision,proto3" json:"revision,omitempty"`
	Profile              *AccountProfile     `protobuf:"bytes,2,opt,name=profile,proto3" json:"profile,omitempty"`
	Properties           *AccountProperties  `protobuf:"bytes,3,opt,name=properties,proto3" json:"properties,omitempty"`
	ConnectedAccounts    []*ConnectedAccount `protobuf:"bytes,4,rep,name=connected_accounts,json=connectedAccounts,proto3" json:"connected_accounts,omitempty"`
	State                string              `protobuf:"bytes,5,opt,name=state,proto3" json:"state,omitempty"`
	Owner                string              `protobuf:"bytes,6,opt,name=owner,proto3" json:"owner,omitempty"`
	Ui                   map[string]string   `protobuf:"bytes,7,rep,name=ui,proto3" json:"ui,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b29259e5b96683f, []int{0}
}

func (m *Account) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Account.Unmarshal(m, b)
}
func (m *Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Account.Marshal(b, m, deterministic)
}
func (m *Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account.Merge(m, src)
}
func (m *Account) XXX_Size() int {
	return xxx_messageInfo_Account.Size(m)
}
func (m *Account) XXX_DiscardUnknown() {
	xxx_messageInfo_Account.DiscardUnknown(m)
}

var xxx_messageInfo_Account proto.InternalMessageInfo

func (m *Account) GetRevision() int64 {
	if m != nil {
		return m.Revision
	}
	return 0
}

func (m *Account) GetProfile() *AccountProfile {
	if m != nil {
		return m.Profile
	}
	return nil
}

func (m *Account) GetProperties() *AccountProperties {
	if m != nil {
		return m.Properties
	}
	return nil
}

func (m *Account) GetConnectedAccounts() []*ConnectedAccount {
	if m != nil {
		return m.ConnectedAccounts
	}
	return nil
}

func (m *Account) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *Account) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Account) GetUi() map[string]string {
	if m != nil {
		return m.Ui
	}
	return nil
}

type AccountProperties struct {
	Subject              string   `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	EmailVerified        bool     `protobuf:"varint,3,opt,name=email_verified,json=emailVerified,proto3" json:"email_verified,omitempty"`
	Created              float64  `protobuf:"fixed64,4,opt,name=created,proto3" json:"created,omitempty"`
	Modified             float64  `protobuf:"fixed64,5,opt,name=modified,proto3" json:"modified,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountProperties) Reset()         { *m = AccountProperties{} }
func (m *AccountProperties) String() string { return proto.CompactTextString(m) }
func (*AccountProperties) ProtoMessage()    {}
func (*AccountProperties) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b29259e5b96683f, []int{1}
}

func (m *AccountProperties) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountProperties.Unmarshal(m, b)
}
func (m *AccountProperties) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountProperties.Marshal(b, m, deterministic)
}
func (m *AccountProperties) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountProperties.Merge(m, src)
}
func (m *AccountProperties) XXX_Size() int {
	return xxx_messageInfo_AccountProperties.Size(m)
}
func (m *AccountProperties) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountProperties.DiscardUnknown(m)
}

var xxx_messageInfo_AccountProperties proto.InternalMessageInfo

func (m *AccountProperties) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *AccountProperties) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *AccountProperties) GetEmailVerified() bool {
	if m != nil {
		return m.EmailVerified
	}
	return false
}

func (m *AccountProperties) GetCreated() float64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *AccountProperties) GetModified() float64 {
	if m != nil {
		return m.Modified
	}
	return 0
}

type AccountProfile struct {
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	GivenName            string   `protobuf:"bytes,4,opt,name=given_name,json=givenName,proto3" json:"given_name,omitempty"`
	FamilyName           string   `protobuf:"bytes,5,opt,name=family_name,json=familyName,proto3" json:"family_name,omitempty"`
	MiddleName           string   `protobuf:"bytes,6,opt,name=middle_name,json=middleName,proto3" json:"middle_name,omitempty"`
	Profile              string   `protobuf:"bytes,7,opt,name=profile,proto3" json:"profile,omitempty"`
	Picture              string   `protobuf:"bytes,8,opt,name=picture,proto3" json:"picture,omitempty"`
	ZoneInfo             string   `protobuf:"bytes,9,opt,name=zone_info,json=zoneInfo,proto3" json:"zone_info,omitempty"`
	Locale               string   `protobuf:"bytes,10,opt,name=locale,proto3" json:"locale,omitempty"`
	FormattedName        string   `protobuf:"bytes,11,opt,name=formatted_name,json=formattedName,proto3" json:"formatted_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountProfile) Reset()         { *m = AccountProfile{} }
func (m *AccountProfile) String() string { return proto.CompactTextString(m) }
func (*AccountProfile) ProtoMessage()    {}
func (*AccountProfile) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b29259e5b96683f, []int{2}
}

func (m *AccountProfile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountProfile.Unmarshal(m, b)
}
func (m *AccountProfile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountProfile.Marshal(b, m, deterministic)
}
func (m *AccountProfile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountProfile.Merge(m, src)
}
func (m *AccountProfile) XXX_Size() int {
	return xxx_messageInfo_AccountProfile.Size(m)
}
func (m *AccountProfile) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountProfile.DiscardUnknown(m)
}

var xxx_messageInfo_AccountProfile proto.InternalMessageInfo

func (m *AccountProfile) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *AccountProfile) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AccountProfile) GetGivenName() string {
	if m != nil {
		return m.GivenName
	}
	return ""
}

func (m *AccountProfile) GetFamilyName() string {
	if m != nil {
		return m.FamilyName
	}
	return ""
}

func (m *AccountProfile) GetMiddleName() string {
	if m != nil {
		return m.MiddleName
	}
	return ""
}

func (m *AccountProfile) GetProfile() string {
	if m != nil {
		return m.Profile
	}
	return ""
}

func (m *AccountProfile) GetPicture() string {
	if m != nil {
		return m.Picture
	}
	return ""
}

func (m *AccountProfile) GetZoneInfo() string {
	if m != nil {
		return m.ZoneInfo
	}
	return ""
}

func (m *AccountProfile) GetLocale() string {
	if m != nil {
		return m.Locale
	}
	return ""
}

func (m *AccountProfile) GetFormattedName() string {
	if m != nil {
		return m.FormattedName
	}
	return ""
}

type ConnectedAccount struct {
	Profile                  *AccountProfile    `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
	Properties               *AccountProperties `protobuf:"bytes,2,opt,name=properties,proto3" json:"properties,omitempty"`
	Provider                 string             `protobuf:"bytes,3,opt,name=provider,proto3" json:"provider,omitempty"`
	Refreshed                float64            `protobuf:"fixed64,4,opt,name=refreshed,proto3" json:"refreshed,omitempty"`
	Revision                 int64              `protobuf:"varint,5,opt,name=revision,proto3" json:"revision,omitempty"`
	LinkRevision             int64              `protobuf:"varint,6,opt,name=link_revision,json=linkRevision,proto3" json:"link_revision,omitempty"`
	Passport                 *Passport          `protobuf:"bytes,7,opt,name=passport,proto3" json:"passport,omitempty"`
	ComputedIdentityProvider *IdentityProvider  `protobuf:"bytes,9,opt,name=computed_identity_provider,json=identityProvider,proto3" json:"computed_identity_provider,omitempty"`
	ComputedLoginHint        string             `protobuf:"bytes,10,opt,name=computed_login_hint,json=loginHint,proto3" json:"computed_login_hint,omitempty"`
	Primary                  bool               `protobuf:"varint,11,opt,name=primary,proto3" json:"primary,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}           `json:"-"`
	XXX_unrecognized         []byte             `json:"-"`
	XXX_sizecache            int32              `json:"-"`
}

func (m *ConnectedAccount) Reset()         { *m = ConnectedAccount{} }
func (m *ConnectedAccount) String() string { return proto.CompactTextString(m) }
func (*ConnectedAccount) ProtoMessage()    {}
func (*ConnectedAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b29259e5b96683f, []int{3}
}

func (m *ConnectedAccount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectedAccount.Unmarshal(m, b)
}
func (m *ConnectedAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectedAccount.Marshal(b, m, deterministic)
}
func (m *ConnectedAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectedAccount.Merge(m, src)
}
func (m *ConnectedAccount) XXX_Size() int {
	return xxx_messageInfo_ConnectedAccount.Size(m)
}
func (m *ConnectedAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectedAccount.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectedAccount proto.InternalMessageInfo

func (m *ConnectedAccount) GetProfile() *AccountProfile {
	if m != nil {
		return m.Profile
	}
	return nil
}

func (m *ConnectedAccount) GetProperties() *AccountProperties {
	if m != nil {
		return m.Properties
	}
	return nil
}

func (m *ConnectedAccount) GetProvider() string {
	if m != nil {
		return m.Provider
	}
	return ""
}

func (m *ConnectedAccount) GetRefreshed() float64 {
	if m != nil {
		return m.Refreshed
	}
	return 0
}

func (m *ConnectedAccount) GetRevision() int64 {
	if m != nil {
		return m.Revision
	}
	return 0
}

func (m *ConnectedAccount) GetLinkRevision() int64 {
	if m != nil {
		return m.LinkRevision
	}
	return 0
}

func (m *ConnectedAccount) GetPassport() *Passport {
	if m != nil {
		return m.Passport
	}
	return nil
}

func (m *ConnectedAccount) GetComputedIdentityProvider() *IdentityProvider {
	if m != nil {
		return m.ComputedIdentityProvider
	}
	return nil
}

func (m *ConnectedAccount) GetComputedLoginHint() string {
	if m != nil {
		return m.ComputedLoginHint
	}
	return ""
}

func (m *ConnectedAccount) GetPrimary() bool {
	if m != nil {
		return m.Primary
	}
	return false
}

type AccountLookup struct {
	Subject              string   `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty"`
	Revision             int64    `protobuf:"varint,2,opt,name=revision,proto3" json:"revision,omitempty"`
	CommitTime           float64  `protobuf:"fixed64,3,opt,name=commit_time,json=commitTime,proto3" json:"commit_time,omitempty"`
	State                string   `protobuf:"bytes,4,opt,name=state,proto3" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountLookup) Reset()         { *m = AccountLookup{} }
func (m *AccountLookup) String() string { return proto.CompactTextString(m) }
func (*AccountLookup) ProtoMessage()    {}
func (*AccountLookup) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b29259e5b96683f, []int{4}
}

func (m *AccountLookup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountLookup.Unmarshal(m, b)
}
func (m *AccountLookup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountLookup.Marshal(b, m, deterministic)
}
func (m *AccountLookup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountLookup.Merge(m, src)
}
func (m *AccountLookup) XXX_Size() int {
	return xxx_messageInfo_AccountLookup.Size(m)
}
func (m *AccountLookup) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountLookup.DiscardUnknown(m)
}

var xxx_messageInfo_AccountLookup proto.InternalMessageInfo

func (m *AccountLookup) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *AccountLookup) GetRevision() int64 {
	if m != nil {
		return m.Revision
	}
	return 0
}

func (m *AccountLookup) GetCommitTime() float64 {
	if m != nil {
		return m.CommitTime
	}
	return 0
}

func (m *AccountLookup) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

type IdentityProvider struct {
	Issuer               string            `protobuf:"bytes,1,opt,name=issuer,proto3" json:"issuer,omitempty"`
	AuthorizeUrl         string            `protobuf:"bytes,2,opt,name=authorize_url,json=authorizeUrl,proto3" json:"authorize_url,omitempty"`
	ResponseType         string            `protobuf:"bytes,3,opt,name=response_type,json=responseType,proto3" json:"response_type,omitempty"`
	TokenUrl             string            `protobuf:"bytes,4,opt,name=token_url,json=tokenUrl,proto3" json:"token_url,omitempty"`
	Scopes               []string          `protobuf:"bytes,5,rep,name=scopes,proto3" json:"scopes,omitempty"`
	TranslateUsing       string            `protobuf:"bytes,6,opt,name=translate_using,json=translateUsing,proto3" json:"translate_using,omitempty"`
	ClientId             string            `protobuf:"bytes,7,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Ui                   map[string]string `protobuf:"bytes,8,rep,name=ui,proto3" json:"ui,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *IdentityProvider) Reset()         { *m = IdentityProvider{} }
func (m *IdentityProvider) String() string { return proto.CompactTextString(m) }
func (*IdentityProvider) ProtoMessage()    {}
func (*IdentityProvider) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b29259e5b96683f, []int{5}
}

func (m *IdentityProvider) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IdentityProvider.Unmarshal(m, b)
}
func (m *IdentityProvider) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IdentityProvider.Marshal(b, m, deterministic)
}
func (m *IdentityProvider) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdentityProvider.Merge(m, src)
}
func (m *IdentityProvider) XXX_Size() int {
	return xxx_messageInfo_IdentityProvider.Size(m)
}
func (m *IdentityProvider) XXX_DiscardUnknown() {
	xxx_messageInfo_IdentityProvider.DiscardUnknown(m)
}

var xxx_messageInfo_IdentityProvider proto.InternalMessageInfo

func (m *IdentityProvider) GetIssuer() string {
	if m != nil {
		return m.Issuer
	}
	return ""
}

func (m *IdentityProvider) GetAuthorizeUrl() string {
	if m != nil {
		return m.AuthorizeUrl
	}
	return ""
}

func (m *IdentityProvider) GetResponseType() string {
	if m != nil {
		return m.ResponseType
	}
	return ""
}

func (m *IdentityProvider) GetTokenUrl() string {
	if m != nil {
		return m.TokenUrl
	}
	return ""
}

func (m *IdentityProvider) GetScopes() []string {
	if m != nil {
		return m.Scopes
	}
	return nil
}

func (m *IdentityProvider) GetTranslateUsing() string {
	if m != nil {
		return m.TranslateUsing
	}
	return ""
}

func (m *IdentityProvider) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *IdentityProvider) GetUi() map[string]string {
	if m != nil {
		return m.Ui
	}
	return nil
}

func init() {
	proto.RegisterType((*Account)(nil), "common.Account")
	proto.RegisterMapType((map[string]string)(nil), "common.Account.UiEntry")
	proto.RegisterType((*AccountProperties)(nil), "common.AccountProperties")
	proto.RegisterType((*AccountProfile)(nil), "common.AccountProfile")
	proto.RegisterType((*ConnectedAccount)(nil), "common.ConnectedAccount")
	proto.RegisterType((*AccountLookup)(nil), "common.AccountLookup")
	proto.RegisterType((*IdentityProvider)(nil), "common.IdentityProvider")
	proto.RegisterMapType((map[string]string)(nil), "common.IdentityProvider.UiEntry")
}

func init() {
	proto.RegisterFile("proto/common/v1/account.proto", fileDescriptor_9b29259e5b96683f)
}

var fileDescriptor_9b29259e5b96683f = []byte{
	// 864 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xdd, 0x6e, 0xdc, 0x44,
	0x14, 0x96, 0xbd, 0x7f, 0xf6, 0xd9, 0x24, 0x6c, 0x07, 0x14, 0xcc, 0xd2, 0xaa, 0xab, 0x45, 0xd0,
	0xbd, 0x20, 0xd9, 0x12, 0x84, 0x04, 0xdc, 0x41, 0x85, 0x4a, 0x24, 0x54, 0x45, 0xa6, 0xe1, 0x82,
	0x1b, 0x6b, 0x62, 0x9f, 0xdd, 0x1d, 0x62, 0xcf, 0x58, 0x33, 0xe3, 0x45, 0x5b, 0xf1, 0x30, 0xbc,
	0x02, 0xb7, 0x3c, 0x09, 0x6f, 0xc0, 0x6b, 0xa0, 0xf9, 0xb1, 0x93, 0xba, 0x02, 0x09, 0xc4, 0xdd,
	0x7c, 0xdf, 0xf9, 0xe6, 0xf8, 0xcc, 0xf9, 0x33, 0x3c, 0xaa, 0xa5, 0xd0, 0x62, 0x9d, 0x8b, 0xaa,
	0x12, 0x7c, 0xbd, 0xff, 0x64, 0x4d, 0xf3, 0x5c, 0x34, 0x5c, 0x9f, 0x5b, 0x9e, 0x8c, 0x9d, 0x61,
	0xfe, 0xb0, 0x2f, 0x73, 0x27, 0xa7, 0x5a, 0xfe, 0x19, 0xc2, 0xe4, 0x2b, 0x77, 0x8f, 0xcc, 0x21,
	0x92, 0xb8, 0x67, 0x8a, 0x09, 0x9e, 0x04, 0x8b, 0x60, 0x35, 0x48, 0x3b, 0x4c, 0x9e, 0xc2, 0xa4,
	0x96, 0x62, 0xc3, 0x4a, 0x4c, 0xc2, 0x45, 0xb0, 0x9a, 0x5e, 0x9c, 0x9e, 0x7b, 0x3f, 0xfe, 0xf6,
	0x95, 0xb3, 0xa6, 0xad, 0x8c, 0x7c, 0x01, 0x50, 0x4b, 0x51, 0xa3, 0xd4, 0x0c, 0x55, 0x32, 0xb0,
	0x97, 0xde, 0x7b, 0xf3, 0x92, 0x17, 0xa4, 0xf7, 0xc4, 0xe4, 0x39, 0x90, 0x5c, 0x70, 0x8e, 0xb9,
	0xc6, 0x22, 0xf3, 0xaf, 0x52, 0xc9, 0x70, 0x31, 0x58, 0x4d, 0x2f, 0x92, 0xd6, 0xc5, 0xb3, 0x56,
	0xe1, 0x7d, 0xa5, 0x0f, 0xf2, 0x1e, 0xa3, 0xc8, 0x3b, 0x30, 0x52, 0x9a, 0x6a, 0x4c, 0x46, 0x8b,
	0x60, 0x15, 0xa7, 0x0e, 0x18, 0x56, 0xfc, 0xcc, 0x51, 0x26, 0x63, 0xc7, 0x5a, 0x40, 0x9e, 0x40,
	0xd8, 0xb0, 0x64, 0x62, 0x3f, 0xf2, 0x6e, 0x2f, 0xce, 0xf3, 0x6b, 0xf6, 0x0d, 0xd7, 0xf2, 0x90,
	0x86, 0x0d, 0x9b, 0x7f, 0x06, 0x13, 0x0f, 0xc9, 0x0c, 0x06, 0xb7, 0x78, 0xb0, 0xc9, 0x8a, 0x53,
	0x73, 0x34, 0xbe, 0xf7, 0xb4, 0x6c, 0x5c, 0x96, 0xe2, 0xd4, 0x81, 0x2f, 0xc3, 0xcf, 0x83, 0xe5,
	0xaf, 0x01, 0x3c, 0x78, 0xe3, 0xd9, 0x24, 0x81, 0x89, 0x6a, 0x6e, 0x7e, 0xc2, 0x5c, 0x7b, 0x2f,
	0x2d, 0x34, 0x9e, 0xb0, 0xa2, 0xac, 0x6c, 0x3d, 0x59, 0x40, 0x3e, 0x84, 0x13, 0x7b, 0xc8, 0xf6,
	0x28, 0xd9, 0x86, 0x61, 0x61, 0x33, 0x1b, 0xa5, 0xc7, 0x96, 0xfd, 0xc1, 0x93, 0xc6, 0x6d, 0x2e,
	0x91, 0x6a, 0x2c, 0x92, 0xe1, 0x22, 0x58, 0x05, 0x69, 0x0b, 0x4d, 0x91, 0x2b, 0x51, 0xb8, 0xab,
	0x23, 0x6b, 0xea, 0xf0, 0xf2, 0xb7, 0x10, 0x4e, 0x5e, 0x2f, 0xa7, 0x91, 0x37, 0x0a, 0x25, 0xa7,
	0x55, 0xfb, 0xa4, 0x0e, 0x13, 0x02, 0x43, 0xcb, 0x0f, 0x2c, 0x6f, 0xcf, 0xe4, 0x11, 0xc0, 0x96,
	0xed, 0x91, 0x67, 0xd6, 0x32, 0xb4, 0x96, 0xd8, 0x32, 0x2f, 0x8c, 0xf9, 0x31, 0x4c, 0x37, 0xb4,
	0x62, 0xe5, 0xc1, 0xd9, 0x5d, 0x59, 0xc0, 0x51, 0xad, 0xa0, 0x62, 0x45, 0x51, 0xa2, 0x13, 0xb8,
	0x0a, 0x81, 0xa3, 0xac, 0x20, 0xb9, 0x6b, 0xc4, 0x89, 0x4b, 0x58, 0xdb, 0x70, 0xc6, 0xc2, 0x72,
	0xdd, 0x48, 0x4c, 0x22, 0x6f, 0x71, 0x90, 0xbc, 0x0f, 0xf1, 0x2b, 0xc1, 0x31, 0x63, 0x7c, 0x23,
	0x92, 0xd8, 0xbd, 0xc2, 0x10, 0x97, 0x7c, 0x23, 0xc8, 0x29, 0x8c, 0x4b, 0x91, 0xd3, 0x12, 0x13,
	0xb0, 0x16, 0x8f, 0x4c, 0xa6, 0x37, 0x42, 0x56, 0x54, 0x9b, 0x26, 0xb4, 0xc1, 0x4c, 0xad, 0xfd,
	0xb8, 0x63, 0x4d, 0x3c, 0xcb, 0xdf, 0x07, 0x30, 0xeb, 0xb7, 0xe2, 0xfd, 0x69, 0x09, 0xfe, 0xcb,
	0xb4, 0x84, 0xff, 0x66, 0x5a, 0xe6, 0x10, 0xd5, 0x52, 0xec, 0x59, 0x81, 0xd2, 0x97, 0xa2, 0xc3,
	0xe4, 0x21, 0xc4, 0x12, 0x37, 0x12, 0xd5, 0xae, 0xeb, 0x84, 0x3b, 0xe2, 0xb5, 0x81, 0x1f, 0xf5,
	0x06, 0xfe, 0x03, 0x38, 0x2e, 0x19, 0xbf, 0xcd, 0x3a, 0xc1, 0xd8, 0x0a, 0x8e, 0x0c, 0x99, 0xb6,
	0xa2, 0x8f, 0x21, 0xaa, 0xa9, 0x52, 0xb5, 0x90, 0xda, 0x56, 0x63, 0x7a, 0x31, 0x6b, 0x63, 0xbe,
	0xf2, 0x7c, 0xda, 0x29, 0xc8, 0x0b, 0x98, 0xe7, 0xa2, 0xaa, 0x1b, 0x93, 0x50, 0x56, 0x20, 0xd7,
	0x4c, 0x1f, 0xb2, 0x2e, 0xf4, 0xd8, 0xde, 0xef, 0xc6, 0xfb, 0xd2, 0x0b, 0xae, 0xbc, 0x3d, 0x9d,
	0xb1, 0x1e, 0x43, 0x3e, 0x82, 0xb7, 0x3b, 0x7f, 0xa5, 0xd8, 0x32, 0x9e, 0xed, 0x18, 0xd7, 0xbe,
	0x8c, 0xb1, 0x65, 0xbe, 0x65, 0x5c, 0xbb, 0x96, 0x61, 0x15, 0x95, 0x07, 0x5b, 0xc2, 0x28, 0x6d,
	0xe1, 0xf2, 0x17, 0x38, 0xf6, 0xb9, 0xfd, 0x4e, 0x88, 0xdb, 0xa6, 0xfe, 0x87, 0x71, 0xbc, 0x9f,
	0xab, 0xb0, 0x97, 0xab, 0xc7, 0x30, 0x35, 0x51, 0x33, 0x9d, 0x69, 0xe6, 0xe7, 0x21, 0x48, 0xc1,
	0x51, 0x2f, 0x59, 0x85, 0x77, 0x7b, 0x68, 0x78, 0x6f, 0x0f, 0x2d, 0xff, 0x08, 0x61, 0xd6, 0x7f,
	0xa6, 0x69, 0x47, 0xa6, 0x54, 0x83, 0xd2, 0x07, 0xe0, 0x91, 0xa9, 0x07, 0x6d, 0xf4, 0x4e, 0x48,
	0xf6, 0x0a, 0xb3, 0x46, 0xb6, 0x6b, 0xe1, 0xa8, 0x23, 0xaf, 0x65, 0x69, 0x44, 0x12, 0x55, 0x2d,
	0xb8, 0xc2, 0x4c, 0x1f, 0xea, 0x76, 0x34, 0x8f, 0x5a, 0xf2, 0xe5, 0xa1, 0xb6, 0xd3, 0xa0, 0xc5,
	0x2d, 0x72, 0xeb, 0xc5, 0x05, 0x14, 0x59, 0xc2, 0x78, 0x38, 0x85, 0xb1, 0xca, 0x45, 0x8d, 0x2a,
	0x19, 0x2d, 0x06, 0xe6, 0xf3, 0x0e, 0x91, 0x27, 0xf0, 0x96, 0x96, 0x94, 0xab, 0x92, 0x6a, 0xcc,
	0x1a, 0xc5, 0xf8, 0xd6, 0xcf, 0xe6, 0x49, 0x47, 0x5f, 0x1b, 0xd6, 0x78, 0xcf, 0x4b, 0x86, 0x5c,
	0x67, 0xac, 0xf0, 0x13, 0x1a, 0x39, 0xe2, 0xb2, 0x20, 0x4f, 0xed, 0x8e, 0x8d, 0xec, 0x8e, 0x5d,
	0xfc, 0x5d, 0xa5, 0xff, 0x87, 0x65, 0xfb, 0xf5, 0xf5, 0x8f, 0xdf, 0x6f, 0x99, 0xde, 0x35, 0x37,
	0xe6, 0x23, 0xeb, 0xe7, 0x42, 0x6c, 0x4b, 0x7c, 0x56, 0x8a, 0xa6, 0xb8, 0x2a, 0xa9, 0x36, 0x03,
	0xbc, 0xde, 0x21, 0x2d, 0xf5, 0x2e, 0xa7, 0x12, 0xcf, 0x36, 0x58, 0xa0, 0x34, 0xcb, 0xf1, 0x8c,
	0xe6, 0x39, 0x2a, 0x75, 0xa6, 0x50, 0xee, 0x59, 0x8e, 0x6a, 0xdd, 0xfb, 0x75, 0xde, 0x8c, 0x2d,
	0xf1, 0xe9, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x1d, 0x02, 0x08, 0x64, 0x7b, 0x07, 0x00, 0x00,
}
