// Copyright (C) 2020 Finogeeks Co., Ltd
//
// This program is free software: you can redistribute it and/or  modify
// it under the terms of the GNU Affero General Public License, version 3,
// as published by the Free Software Foundation.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// Code generated by capnpc-go. DO NOT EDIT.

package external

import (
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
)

type GetAvatarURLRequestCapn struct{ capnp.Struct }

// GetAvatarURLRequestCapn_TypeID is the unique identifier for the type GetAvatarURLRequestCapn.
const GetAvatarURLRequestCapn_TypeID = 0xbe5e69974aa5e7af

func NewGetAvatarURLRequestCapn(s *capnp.Segment) (GetAvatarURLRequestCapn, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return GetAvatarURLRequestCapn{st}, err
}

func NewRootGetAvatarURLRequestCapn(s *capnp.Segment) (GetAvatarURLRequestCapn, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return GetAvatarURLRequestCapn{st}, err
}

func ReadRootGetAvatarURLRequestCapn(msg *capnp.Message) (GetAvatarURLRequestCapn, error) {
	root, err := msg.RootPtr()
	return GetAvatarURLRequestCapn{root.Struct()}, err
}

func (s GetAvatarURLRequestCapn) String() string {
	str, _ := text.Marshal(0xbe5e69974aa5e7af, s.Struct)
	return str
}

func (s GetAvatarURLRequestCapn) UserID() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s GetAvatarURLRequestCapn) HasUserID() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s GetAvatarURLRequestCapn) UserIDBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s GetAvatarURLRequestCapn) SetUserID(v string) error {
	return s.Struct.SetText(0, v)
}

// GetAvatarURLRequestCapn_List is a list of GetAvatarURLRequestCapn.
type GetAvatarURLRequestCapn_List struct{ capnp.List }

// NewGetAvatarURLRequestCapn creates a new list of GetAvatarURLRequestCapn.
func NewGetAvatarURLRequestCapn_List(s *capnp.Segment, sz int32) (GetAvatarURLRequestCapn_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return GetAvatarURLRequestCapn_List{l}, err
}

func (s GetAvatarURLRequestCapn_List) At(i int) GetAvatarURLRequestCapn {
	return GetAvatarURLRequestCapn{s.List.Struct(i)}
}

func (s GetAvatarURLRequestCapn_List) Set(i int, v GetAvatarURLRequestCapn) error {
	return s.List.SetStruct(i, v.Struct)
}

func (s GetAvatarURLRequestCapn_List) String() string {
	str, _ := text.MarshalList(0xbe5e69974aa5e7af, s.List)
	return str
}

// GetAvatarURLRequestCapn_Promise is a wrapper for a GetAvatarURLRequestCapn promised by a client call.
type GetAvatarURLRequestCapn_Promise struct{ *capnp.Pipeline }

func (p GetAvatarURLRequestCapn_Promise) Struct() (GetAvatarURLRequestCapn, error) {
	s, err := p.Pipeline.Struct()
	return GetAvatarURLRequestCapn{s}, err
}

type GetAvatarURLResponseCapn struct{ capnp.Struct }

// GetAvatarURLResponseCapn_TypeID is the unique identifier for the type GetAvatarURLResponseCapn.
const GetAvatarURLResponseCapn_TypeID = 0x97c59b2360f2bd24

func NewGetAvatarURLResponseCapn(s *capnp.Segment) (GetAvatarURLResponseCapn, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return GetAvatarURLResponseCapn{st}, err
}

func NewRootGetAvatarURLResponseCapn(s *capnp.Segment) (GetAvatarURLResponseCapn, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return GetAvatarURLResponseCapn{st}, err
}

func ReadRootGetAvatarURLResponseCapn(msg *capnp.Message) (GetAvatarURLResponseCapn, error) {
	root, err := msg.RootPtr()
	return GetAvatarURLResponseCapn{root.Struct()}, err
}

func (s GetAvatarURLResponseCapn) String() string {
	str, _ := text.Marshal(0x97c59b2360f2bd24, s.Struct)
	return str
}

func (s GetAvatarURLResponseCapn) AvatarURL() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s GetAvatarURLResponseCapn) HasAvatarURL() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s GetAvatarURLResponseCapn) AvatarURLBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s GetAvatarURLResponseCapn) SetAvatarURL(v string) error {
	return s.Struct.SetText(0, v)
}

// GetAvatarURLResponseCapn_List is a list of GetAvatarURLResponseCapn.
type GetAvatarURLResponseCapn_List struct{ capnp.List }

// NewGetAvatarURLResponseCapn creates a new list of GetAvatarURLResponseCapn.
func NewGetAvatarURLResponseCapn_List(s *capnp.Segment, sz int32) (GetAvatarURLResponseCapn_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return GetAvatarURLResponseCapn_List{l}, err
}

func (s GetAvatarURLResponseCapn_List) At(i int) GetAvatarURLResponseCapn {
	return GetAvatarURLResponseCapn{s.List.Struct(i)}
}

func (s GetAvatarURLResponseCapn_List) Set(i int, v GetAvatarURLResponseCapn) error {
	return s.List.SetStruct(i, v.Struct)
}

func (s GetAvatarURLResponseCapn_List) String() string {
	str, _ := text.MarshalList(0x97c59b2360f2bd24, s.List)
	return str
}

// GetAvatarURLResponseCapn_Promise is a wrapper for a GetAvatarURLResponseCapn promised by a client call.
type GetAvatarURLResponseCapn_Promise struct{ *capnp.Pipeline }

func (p GetAvatarURLResponseCapn_Promise) Struct() (GetAvatarURLResponseCapn, error) {
	s, err := p.Pipeline.Struct()
	return GetAvatarURLResponseCapn{s}, err
}

type GetDisplayNameRequestCapn struct{ capnp.Struct }

// GetDisplayNameRequestCapn_TypeID is the unique identifier for the type GetDisplayNameRequestCapn.
const GetDisplayNameRequestCapn_TypeID = 0x87aeb8fbedaaf0a1

func NewGetDisplayNameRequestCapn(s *capnp.Segment) (GetDisplayNameRequestCapn, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return GetDisplayNameRequestCapn{st}, err
}

func NewRootGetDisplayNameRequestCapn(s *capnp.Segment) (GetDisplayNameRequestCapn, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return GetDisplayNameRequestCapn{st}, err
}

func ReadRootGetDisplayNameRequestCapn(msg *capnp.Message) (GetDisplayNameRequestCapn, error) {
	root, err := msg.RootPtr()
	return GetDisplayNameRequestCapn{root.Struct()}, err
}

func (s GetDisplayNameRequestCapn) String() string {
	str, _ := text.Marshal(0x87aeb8fbedaaf0a1, s.Struct)
	return str
}

func (s GetDisplayNameRequestCapn) UserID() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s GetDisplayNameRequestCapn) HasUserID() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s GetDisplayNameRequestCapn) UserIDBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s GetDisplayNameRequestCapn) SetUserID(v string) error {
	return s.Struct.SetText(0, v)
}

// GetDisplayNameRequestCapn_List is a list of GetDisplayNameRequestCapn.
type GetDisplayNameRequestCapn_List struct{ capnp.List }

// NewGetDisplayNameRequestCapn creates a new list of GetDisplayNameRequestCapn.
func NewGetDisplayNameRequestCapn_List(s *capnp.Segment, sz int32) (GetDisplayNameRequestCapn_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return GetDisplayNameRequestCapn_List{l}, err
}

func (s GetDisplayNameRequestCapn_List) At(i int) GetDisplayNameRequestCapn {
	return GetDisplayNameRequestCapn{s.List.Struct(i)}
}

func (s GetDisplayNameRequestCapn_List) Set(i int, v GetDisplayNameRequestCapn) error {
	return s.List.SetStruct(i, v.Struct)
}

func (s GetDisplayNameRequestCapn_List) String() string {
	str, _ := text.MarshalList(0x87aeb8fbedaaf0a1, s.List)
	return str
}

// GetDisplayNameRequestCapn_Promise is a wrapper for a GetDisplayNameRequestCapn promised by a client call.
type GetDisplayNameRequestCapn_Promise struct{ *capnp.Pipeline }

func (p GetDisplayNameRequestCapn_Promise) Struct() (GetDisplayNameRequestCapn, error) {
	s, err := p.Pipeline.Struct()
	return GetDisplayNameRequestCapn{s}, err
}

type GetDisplayNameResponseCapn struct{ capnp.Struct }

// GetDisplayNameResponseCapn_TypeID is the unique identifier for the type GetDisplayNameResponseCapn.
const GetDisplayNameResponseCapn_TypeID = 0xb46aae264582eaec

func NewGetDisplayNameResponseCapn(s *capnp.Segment) (GetDisplayNameResponseCapn, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return GetDisplayNameResponseCapn{st}, err
}

func NewRootGetDisplayNameResponseCapn(s *capnp.Segment) (GetDisplayNameResponseCapn, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return GetDisplayNameResponseCapn{st}, err
}

func ReadRootGetDisplayNameResponseCapn(msg *capnp.Message) (GetDisplayNameResponseCapn, error) {
	root, err := msg.RootPtr()
	return GetDisplayNameResponseCapn{root.Struct()}, err
}

func (s GetDisplayNameResponseCapn) String() string {
	str, _ := text.Marshal(0xb46aae264582eaec, s.Struct)
	return str
}

func (s GetDisplayNameResponseCapn) DisplayName() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s GetDisplayNameResponseCapn) HasDisplayName() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s GetDisplayNameResponseCapn) DisplayNameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s GetDisplayNameResponseCapn) SetDisplayName(v string) error {
	return s.Struct.SetText(0, v)
}

// GetDisplayNameResponseCapn_List is a list of GetDisplayNameResponseCapn.
type GetDisplayNameResponseCapn_List struct{ capnp.List }

// NewGetDisplayNameResponseCapn creates a new list of GetDisplayNameResponseCapn.
func NewGetDisplayNameResponseCapn_List(s *capnp.Segment, sz int32) (GetDisplayNameResponseCapn_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return GetDisplayNameResponseCapn_List{l}, err
}

func (s GetDisplayNameResponseCapn_List) At(i int) GetDisplayNameResponseCapn {
	return GetDisplayNameResponseCapn{s.List.Struct(i)}
}

func (s GetDisplayNameResponseCapn_List) Set(i int, v GetDisplayNameResponseCapn) error {
	return s.List.SetStruct(i, v.Struct)
}

func (s GetDisplayNameResponseCapn_List) String() string {
	str, _ := text.MarshalList(0xb46aae264582eaec, s.List)
	return str
}

// GetDisplayNameResponseCapn_Promise is a wrapper for a GetDisplayNameResponseCapn promised by a client call.
type GetDisplayNameResponseCapn_Promise struct{ *capnp.Pipeline }

func (p GetDisplayNameResponseCapn_Promise) Struct() (GetDisplayNameResponseCapn, error) {
	s, err := p.Pipeline.Struct()
	return GetDisplayNameResponseCapn{s}, err
}

type GetProfileRequestCapn struct{ capnp.Struct }

// GetProfileRequestCapn_TypeID is the unique identifier for the type GetProfileRequestCapn.
const GetProfileRequestCapn_TypeID = 0xade4a4f4ea93c7c2

func NewGetProfileRequestCapn(s *capnp.Segment) (GetProfileRequestCapn, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return GetProfileRequestCapn{st}, err
}

func NewRootGetProfileRequestCapn(s *capnp.Segment) (GetProfileRequestCapn, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return GetProfileRequestCapn{st}, err
}

func ReadRootGetProfileRequestCapn(msg *capnp.Message) (GetProfileRequestCapn, error) {
	root, err := msg.RootPtr()
	return GetProfileRequestCapn{root.Struct()}, err
}

func (s GetProfileRequestCapn) String() string {
	str, _ := text.Marshal(0xade4a4f4ea93c7c2, s.Struct)
	return str
}

func (s GetProfileRequestCapn) UserID() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s GetProfileRequestCapn) HasUserID() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s GetProfileRequestCapn) UserIDBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s GetProfileRequestCapn) SetUserID(v string) error {
	return s.Struct.SetText(0, v)
}

// GetProfileRequestCapn_List is a list of GetProfileRequestCapn.
type GetProfileRequestCapn_List struct{ capnp.List }

// NewGetProfileRequestCapn creates a new list of GetProfileRequestCapn.
func NewGetProfileRequestCapn_List(s *capnp.Segment, sz int32) (GetProfileRequestCapn_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return GetProfileRequestCapn_List{l}, err
}

func (s GetProfileRequestCapn_List) At(i int) GetProfileRequestCapn {
	return GetProfileRequestCapn{s.List.Struct(i)}
}

func (s GetProfileRequestCapn_List) Set(i int, v GetProfileRequestCapn) error {
	return s.List.SetStruct(i, v.Struct)
}

func (s GetProfileRequestCapn_List) String() string {
	str, _ := text.MarshalList(0xade4a4f4ea93c7c2, s.List)
	return str
}

// GetProfileRequestCapn_Promise is a wrapper for a GetProfileRequestCapn promised by a client call.
type GetProfileRequestCapn_Promise struct{ *capnp.Pipeline }

func (p GetProfileRequestCapn_Promise) Struct() (GetProfileRequestCapn, error) {
	s, err := p.Pipeline.Struct()
	return GetProfileRequestCapn{s}, err
}

type GetProfileResponseCapn struct{ capnp.Struct }

// GetProfileResponseCapn_TypeID is the unique identifier for the type GetProfileResponseCapn.
const GetProfileResponseCapn_TypeID = 0xef61725d5fccd6e0

func NewGetProfileResponseCapn(s *capnp.Segment) (GetProfileResponseCapn, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return GetProfileResponseCapn{st}, err
}

func NewRootGetProfileResponseCapn(s *capnp.Segment) (GetProfileResponseCapn, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return GetProfileResponseCapn{st}, err
}

func ReadRootGetProfileResponseCapn(msg *capnp.Message) (GetProfileResponseCapn, error) {
	root, err := msg.RootPtr()
	return GetProfileResponseCapn{root.Struct()}, err
}

func (s GetProfileResponseCapn) String() string {
	str, _ := text.Marshal(0xef61725d5fccd6e0, s.Struct)
	return str
}

func (s GetProfileResponseCapn) AvatarURL() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s GetProfileResponseCapn) HasAvatarURL() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s GetProfileResponseCapn) AvatarURLBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s GetProfileResponseCapn) SetAvatarURL(v string) error {
	return s.Struct.SetText(0, v)
}

func (s GetProfileResponseCapn) DisplayName() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s GetProfileResponseCapn) HasDisplayName() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s GetProfileResponseCapn) DisplayNameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s GetProfileResponseCapn) SetDisplayName(v string) error {
	return s.Struct.SetText(1, v)
}

// GetProfileResponseCapn_List is a list of GetProfileResponseCapn.
type GetProfileResponseCapn_List struct{ capnp.List }

// NewGetProfileResponseCapn creates a new list of GetProfileResponseCapn.
func NewGetProfileResponseCapn_List(s *capnp.Segment, sz int32) (GetProfileResponseCapn_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return GetProfileResponseCapn_List{l}, err
}

func (s GetProfileResponseCapn_List) At(i int) GetProfileResponseCapn {
	return GetProfileResponseCapn{s.List.Struct(i)}
}

func (s GetProfileResponseCapn_List) Set(i int, v GetProfileResponseCapn) error {
	return s.List.SetStruct(i, v.Struct)
}

func (s GetProfileResponseCapn_List) String() string {
	str, _ := text.MarshalList(0xef61725d5fccd6e0, s.List)
	return str
}

// GetProfileResponseCapn_Promise is a wrapper for a GetProfileResponseCapn promised by a client call.
type GetProfileResponseCapn_Promise struct{ *capnp.Pipeline }

func (p GetProfileResponseCapn_Promise) Struct() (GetProfileResponseCapn, error) {
	s, err := p.Pipeline.Struct()
	return GetProfileResponseCapn{s}, err
}

type PutAvatarURLRequestCapn struct{ capnp.Struct }

// PutAvatarURLRequestCapn_TypeID is the unique identifier for the type PutAvatarURLRequestCapn.
const PutAvatarURLRequestCapn_TypeID = 0x81d3b998d0131f72

func NewPutAvatarURLRequestCapn(s *capnp.Segment) (PutAvatarURLRequestCapn, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return PutAvatarURLRequestCapn{st}, err
}

func NewRootPutAvatarURLRequestCapn(s *capnp.Segment) (PutAvatarURLRequestCapn, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return PutAvatarURLRequestCapn{st}, err
}

func ReadRootPutAvatarURLRequestCapn(msg *capnp.Message) (PutAvatarURLRequestCapn, error) {
	root, err := msg.RootPtr()
	return PutAvatarURLRequestCapn{root.Struct()}, err
}

func (s PutAvatarURLRequestCapn) String() string {
	str, _ := text.Marshal(0x81d3b998d0131f72, s.Struct)
	return str
}

func (s PutAvatarURLRequestCapn) UserID() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s PutAvatarURLRequestCapn) HasUserID() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s PutAvatarURLRequestCapn) UserIDBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s PutAvatarURLRequestCapn) SetUserID(v string) error {
	return s.Struct.SetText(0, v)
}

func (s PutAvatarURLRequestCapn) AvatarURL() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s PutAvatarURLRequestCapn) HasAvatarURL() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s PutAvatarURLRequestCapn) AvatarURLBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s PutAvatarURLRequestCapn) SetAvatarURL(v string) error {
	return s.Struct.SetText(1, v)
}

// PutAvatarURLRequestCapn_List is a list of PutAvatarURLRequestCapn.
type PutAvatarURLRequestCapn_List struct{ capnp.List }

// NewPutAvatarURLRequestCapn creates a new list of PutAvatarURLRequestCapn.
func NewPutAvatarURLRequestCapn_List(s *capnp.Segment, sz int32) (PutAvatarURLRequestCapn_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return PutAvatarURLRequestCapn_List{l}, err
}

func (s PutAvatarURLRequestCapn_List) At(i int) PutAvatarURLRequestCapn {
	return PutAvatarURLRequestCapn{s.List.Struct(i)}
}

func (s PutAvatarURLRequestCapn_List) Set(i int, v PutAvatarURLRequestCapn) error {
	return s.List.SetStruct(i, v.Struct)
}

func (s PutAvatarURLRequestCapn_List) String() string {
	str, _ := text.MarshalList(0x81d3b998d0131f72, s.List)
	return str
}

// PutAvatarURLRequestCapn_Promise is a wrapper for a PutAvatarURLRequestCapn promised by a client call.
type PutAvatarURLRequestCapn_Promise struct{ *capnp.Pipeline }

func (p PutAvatarURLRequestCapn_Promise) Struct() (PutAvatarURLRequestCapn, error) {
	s, err := p.Pipeline.Struct()
	return PutAvatarURLRequestCapn{s}, err
}

type PutDisplayNameRequestCapn struct{ capnp.Struct }

// PutDisplayNameRequestCapn_TypeID is the unique identifier for the type PutDisplayNameRequestCapn.
const PutDisplayNameRequestCapn_TypeID = 0xc2c09e43569e01b4

func NewPutDisplayNameRequestCapn(s *capnp.Segment) (PutDisplayNameRequestCapn, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return PutDisplayNameRequestCapn{st}, err
}

func NewRootPutDisplayNameRequestCapn(s *capnp.Segment) (PutDisplayNameRequestCapn, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return PutDisplayNameRequestCapn{st}, err
}

func ReadRootPutDisplayNameRequestCapn(msg *capnp.Message) (PutDisplayNameRequestCapn, error) {
	root, err := msg.RootPtr()
	return PutDisplayNameRequestCapn{root.Struct()}, err
}

func (s PutDisplayNameRequestCapn) String() string {
	str, _ := text.Marshal(0xc2c09e43569e01b4, s.Struct)
	return str
}

func (s PutDisplayNameRequestCapn) UserID() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s PutDisplayNameRequestCapn) HasUserID() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s PutDisplayNameRequestCapn) UserIDBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s PutDisplayNameRequestCapn) SetUserID(v string) error {
	return s.Struct.SetText(0, v)
}

func (s PutDisplayNameRequestCapn) DisplayName() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s PutDisplayNameRequestCapn) HasDisplayName() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s PutDisplayNameRequestCapn) DisplayNameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s PutDisplayNameRequestCapn) SetDisplayName(v string) error {
	return s.Struct.SetText(1, v)
}

// PutDisplayNameRequestCapn_List is a list of PutDisplayNameRequestCapn.
type PutDisplayNameRequestCapn_List struct{ capnp.List }

// NewPutDisplayNameRequestCapn creates a new list of PutDisplayNameRequestCapn.
func NewPutDisplayNameRequestCapn_List(s *capnp.Segment, sz int32) (PutDisplayNameRequestCapn_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return PutDisplayNameRequestCapn_List{l}, err
}

func (s PutDisplayNameRequestCapn_List) At(i int) PutDisplayNameRequestCapn {
	return PutDisplayNameRequestCapn{s.List.Struct(i)}
}

func (s PutDisplayNameRequestCapn_List) Set(i int, v PutDisplayNameRequestCapn) error {
	return s.List.SetStruct(i, v.Struct)
}

func (s PutDisplayNameRequestCapn_List) String() string {
	str, _ := text.MarshalList(0xc2c09e43569e01b4, s.List)
	return str
}

// PutDisplayNameRequestCapn_Promise is a wrapper for a PutDisplayNameRequestCapn promised by a client call.
type PutDisplayNameRequestCapn_Promise struct{ *capnp.Pipeline }

func (p PutDisplayNameRequestCapn_Promise) Struct() (PutDisplayNameRequestCapn, error) {
	s, err := p.Pipeline.Struct()
	return PutDisplayNameRequestCapn{s}, err
}

const schema_d84e63d12bc2b9b4 = "x\xda2\xf8\xcd\xe1\xc0d\xc8\xba\xdc\x81\x81!\xf8\x03" +
	"\x13+\xdb\xff\"y\xe1\x0b3v^nd\x10\xe4e" +
	"\xfc\xbfe\xe7!\xed\x8b\xc9~7\x18X\x99\xd8\x19\x18" +
	"\x8c-\x15\xb4\x18\x85=\x15\xd8\x19\x18\x84]\x15\xca\x19" +
	"\x18\xff/\xfc\xb0\xea\xed\xef\x1d\xeb\xda\xd1\x143\x82\x14" +
	"/U\xb0b\x14\xde\x0aV\xbcQ\xc1\x9e\x81\xf1\xbf\xca" +
	"\xdeO\x09\xca\xb3\x8fN\xc7\xa6\xf8\xa2\x82\x11\xa3\xf0C" +
	"\xb0\xe2\xbb`\xc5\x87\x8eO~\xf5e\xc9\x93\xb5\xd8\x14" +
	"3*J1\x0a\x0b*\x82\x14\xf3*\x82\x14\xbfy\xd5" +
	"\xe4\xaa\xb6.k\x0b6\xc5\xba\x8aN\x8c\xc2\xb6`\xc5" +
	"\x96`\xc5\xeb\x9f/\xf5\x9a\x9e\x19\xb7\x0f\x9b\xe2XE" +
	"-F\xe1\\\xb0\xe2L\xb0\xe2-\x8c\xf3\xc2\x9c\xe7\x1d" +
	"8\x84-4:\x15\xad\x18\x85g\x82\x15OU\x04\x85" +
	"\xc6\x83kg\xe2c\x8b\x12\xdfcS\xfcUQ\x89Q" +
	"\x98U\x09\xa4\x98Q\xa9\x9c\xe1?\x0a\xdc\xfb\xbf\xa0(" +
	"?-3'U\x8f99\xb1 \xcf*\xa0\xb4\xc4\xb1" +
	",\xb1$\xb1(4\xc8'(\xb5\xb04\xb5\xb8\x84\xdf" +
	"9\xb1 /\x80\x911\x90\x83\x99\x85\x81\x81\x85\x91\x81" +
	"AP\xd3\x8a\x81!P\x85\x991\xd0\x80\x89Q\x90\x91" +
	"Q\x84\x11$\xa8\x1b\xc4\xc0\x10\xa8\xc3\xcc\x18h\xc1\xc4" +
	"h_Z\x9cZ\xe4\xe9\xc2\xc8\xc3\xc0\xc4\xc8\xc3\xc0\xf8" +
	"?\x11j&\x03\xa3\x0f\\\x0c\xd5b\xf7\xd4\x12\x97\xcc" +
	"\xe2\x82\x9c\xc4J\xbf\xc4\xdcT\xb0\xd5\xf6\xc5%0\xab" +
	"Y\xe0V\xf3\x82\xac\xe6`f\x0c\x14\xc1\xb4\x05\xc3D" +
	"$\xaf\x14\x17\xe4\xe7\x15\xcb\xa7b1\x10\xe4l\x1ef" +
	"\xc6@\x09&\"\xdd\x19\x00\x11\x80\x06\x8fs\"39" +
	"nd\xc1\xeak\xb0+\xc1\x8ed@32\x09\xc9\x95" +
	")P-\x0c\xec\x89\xb9\xa9D\xf9\x1e5\"\xc9\x09\xcd" +
	"\x80R|\xf1C(i$\xe1N\x1aDz\x06\x1e\xe8" +
	"\xd0 b\xc7\xb4\x19\x14\x91\x1a\xcc\x8c\x81&H6\x1b" +
	"\x82l6`f\x0c\xb4\xc1\x11\xbbX\xad\x07\x04\x00\x00" +
	"\xff\xffU.9\xfc"

func init() {
	schemas.Register(schema_d84e63d12bc2b9b4,
		0x81d3b998d0131f72,
		0x87aeb8fbedaaf0a1,
		0x97c59b2360f2bd24,
		0xade4a4f4ea93c7c2,
		0xb46aae264582eaec,
		0xbe5e69974aa5e7af,
		0xc2c09e43569e01b4,
		0xef61725d5fccd6e0)
}
