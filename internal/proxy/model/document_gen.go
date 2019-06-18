/**
* Copyright 2018 Comcast Cable Communications Management, LLC
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
* http://www.apache.org/licenses/LICENSE-2.0
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

package model

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *HTTPDocument) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "status_code":
			z.StatusCode, err = dc.ReadInt()
			if err != nil {
				return
			}
		case "status":
			z.Status, err = dc.ReadString()
			if err != nil {
				return
			}
		case "headers":
			var zb0002 uint32
			zb0002, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Headers == nil {
				z.Headers = make(map[string][]string, zb0002)
			} else if len(z.Headers) > 0 {
				for key := range z.Headers {
					delete(z.Headers, key)
				}
			}
			for zb0002 > 0 {
				zb0002--
				var za0001 string
				var za0002 []string
				za0001, err = dc.ReadString()
				if err != nil {
					return
				}
				var zb0003 uint32
				zb0003, err = dc.ReadArrayHeader()
				if err != nil {
					return
				}
				if cap(za0002) >= int(zb0003) {
					za0002 = (za0002)[:zb0003]
				} else {
					za0002 = make([]string, zb0003)
				}
				for za0003 := range za0002 {
					za0002[za0003], err = dc.ReadString()
					if err != nil {
						return
					}
				}
				z.Headers[za0001] = za0002
			}
		case "body":
			z.Body, err = dc.ReadBytes(z.Body)
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *HTTPDocument) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 4
	// write "status_code"
	err = en.Append(0x84, 0xab, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65)
	if err != nil {
		return
	}
	err = en.WriteInt(z.StatusCode)
	if err != nil {
		return
	}
	// write "status"
	err = en.Append(0xa6, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73)
	if err != nil {
		return
	}
	err = en.WriteString(z.Status)
	if err != nil {
		return
	}
	// write "headers"
	err = en.Append(0xa7, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73)
	if err != nil {
		return
	}
	err = en.WriteMapHeader(uint32(len(z.Headers)))
	if err != nil {
		return
	}
	for za0001, za0002 := range z.Headers {
		err = en.WriteString(za0001)
		if err != nil {
			return
		}
		err = en.WriteArrayHeader(uint32(len(za0002)))
		if err != nil {
			return
		}
		for za0003 := range za0002 {
			err = en.WriteString(za0002[za0003])
			if err != nil {
				return
			}
		}
	}
	// write "body"
	err = en.Append(0xa4, 0x62, 0x6f, 0x64, 0x79)
	if err != nil {
		return
	}
	err = en.WriteBytes(z.Body)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *HTTPDocument) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	// string "status_code"
	o = append(o, 0x84, 0xab, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65)
	o = msgp.AppendInt(o, z.StatusCode)
	// string "status"
	o = append(o, 0xa6, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73)
	o = msgp.AppendString(o, z.Status)
	// string "headers"
	o = append(o, 0xa7, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73)
	o = msgp.AppendMapHeader(o, uint32(len(z.Headers)))
	for za0001, za0002 := range z.Headers {
		o = msgp.AppendString(o, za0001)
		o = msgp.AppendArrayHeader(o, uint32(len(za0002)))
		for za0003 := range za0002 {
			o = msgp.AppendString(o, za0002[za0003])
		}
	}
	// string "body"
	o = append(o, 0xa4, 0x62, 0x6f, 0x64, 0x79)
	o = msgp.AppendBytes(o, z.Body)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *HTTPDocument) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "status_code":
			z.StatusCode, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "status":
			z.Status, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "headers":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			if z.Headers == nil {
				z.Headers = make(map[string][]string, zb0002)
			} else if len(z.Headers) > 0 {
				for key := range z.Headers {
					delete(z.Headers, key)
				}
			}
			for zb0002 > 0 {
				var za0001 string
				var za0002 []string
				zb0002--
				za0001, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				var zb0003 uint32
				zb0003, bts, err = msgp.ReadArrayHeaderBytes(bts)
				if err != nil {
					return
				}
				if cap(za0002) >= int(zb0003) {
					za0002 = (za0002)[:zb0003]
				} else {
					za0002 = make([]string, zb0003)
				}
				for za0003 := range za0002 {
					za0002[za0003], bts, err = msgp.ReadStringBytes(bts)
					if err != nil {
						return
					}
				}
				z.Headers[za0001] = za0002
			}
		case "body":
			z.Body, bts, err = msgp.ReadBytesBytes(bts, z.Body)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *HTTPDocument) Msgsize() (s int) {
	s = 1 + 12 + msgp.IntSize + 7 + msgp.StringPrefixSize + len(z.Status) + 8 + msgp.MapHeaderSize
	if z.Headers != nil {
		for za0001, za0002 := range z.Headers {
			_ = za0002
			s += msgp.StringPrefixSize + len(za0001) + msgp.ArrayHeaderSize
			for za0003 := range za0002 {
				s += msgp.StringPrefixSize + len(za0002[za0003])
			}
		}
	}
	s += 5 + msgp.BytesPrefixSize + len(z.Body)
	return
}