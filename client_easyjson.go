// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package jsonrpc2

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonC0e5e3f1DecodeGithubComSpace307Jsonrpc2(in *jlexer.Lexer, out *ResponseError) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "code":
			out.Code = int(in.Int())
		case "message":
			out.Message = string(in.String())
		case "data":
			if m, ok := out.Data.(easyjson.Unmarshaler); ok {
				m.UnmarshalEasyJSON(in)
			} else if m, ok := out.Data.(json.Unmarshaler); ok {
				_ = m.UnmarshalJSON(in.Raw())
			} else {
				out.Data = in.Interface()
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC0e5e3f1EncodeGithubComSpace307Jsonrpc2(out *jwriter.Writer, in ResponseError) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"code\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Code))
	}
	{
		const prefix string = ",\"message\":"
		out.RawString(prefix)
		out.String(string(in.Message))
	}
	{
		const prefix string = ",\"data\":"
		out.RawString(prefix)
		if m, ok := in.Data.(easyjson.Marshaler); ok {
			m.MarshalEasyJSON(out)
		} else if m, ok := in.Data.(json.Marshaler); ok {
			out.Raw(m.MarshalJSON())
		} else {
			out.Raw(json.Marshal(in.Data))
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ResponseError) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC0e5e3f1EncodeGithubComSpace307Jsonrpc2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ResponseError) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC0e5e3f1EncodeGithubComSpace307Jsonrpc2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ResponseError) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC0e5e3f1DecodeGithubComSpace307Jsonrpc2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ResponseError) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC0e5e3f1DecodeGithubComSpace307Jsonrpc2(l, v)
}
func easyjsonC0e5e3f1DecodeGithubComSpace307Jsonrpc21(in *jlexer.Lexer, out *Response) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "jsonrpc":
			out.Version = string(in.String())
		case "result":
			if in.IsNull() {
				in.Skip()
				out.Result = nil
			} else {
				if out.Result == nil {
					out.Result = new(json.RawMessage)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.Result).UnmarshalJSON(data))
				}
			}
		case "error":
			if in.IsNull() {
				in.Skip()
				out.Error = nil
			} else {
				if out.Error == nil {
					out.Error = new(json.RawMessage)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.Error).UnmarshalJSON(data))
				}
			}
		case "id":
			if in.IsNull() {
				in.Skip()
				out.Id = nil
			} else {
				if out.Id == nil {
					out.Id = new(json.RawMessage)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.Id).UnmarshalJSON(data))
				}
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC0e5e3f1EncodeGithubComSpace307Jsonrpc21(out *jwriter.Writer, in Response) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"jsonrpc\":"
		out.RawString(prefix[1:])
		out.String(string(in.Version))
	}
	{
		const prefix string = ",\"result\":"
		out.RawString(prefix)
		if in.Result == nil {
			out.RawString("null")
		} else {
			out.Raw((*in.Result).MarshalJSON())
		}
	}
	{
		const prefix string = ",\"error\":"
		out.RawString(prefix)
		if in.Error == nil {
			out.RawString("null")
		} else {
			out.Raw((*in.Error).MarshalJSON())
		}
	}
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix)
		if in.Id == nil {
			out.RawString("null")
		} else {
			out.Raw((*in.Id).MarshalJSON())
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Response) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC0e5e3f1EncodeGithubComSpace307Jsonrpc21(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Response) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC0e5e3f1EncodeGithubComSpace307Jsonrpc21(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Response) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC0e5e3f1DecodeGithubComSpace307Jsonrpc21(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Response) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC0e5e3f1DecodeGithubComSpace307Jsonrpc21(l, v)
}
func easyjsonC0e5e3f1DecodeGithubComSpace307Jsonrpc22(in *jlexer.Lexer, out *Request) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "jsonrpc":
			out.Version = string(in.String())
		case "method":
			out.Method = string(in.String())
		case "params":
			if m, ok := out.Params.(easyjson.Unmarshaler); ok {
				m.UnmarshalEasyJSON(in)
			} else if m, ok := out.Params.(json.Unmarshaler); ok {
				_ = m.UnmarshalJSON(in.Raw())
			} else {
				out.Params = in.Interface()
			}
		case "id":
			if in.IsNull() {
				in.Skip()
				out.Id = nil
			} else {
				if out.Id == nil {
					out.Id = new(json.RawMessage)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.Id).UnmarshalJSON(data))
				}
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC0e5e3f1EncodeGithubComSpace307Jsonrpc22(out *jwriter.Writer, in Request) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"jsonrpc\":"
		out.RawString(prefix[1:])
		out.String(string(in.Version))
	}
	{
		const prefix string = ",\"method\":"
		out.RawString(prefix)
		out.String(string(in.Method))
	}
	{
		const prefix string = ",\"params\":"
		out.RawString(prefix)
		if m, ok := in.Params.(easyjson.Marshaler); ok {
			m.MarshalEasyJSON(out)
		} else if m, ok := in.Params.(json.Marshaler); ok {
			out.Raw(m.MarshalJSON())
		} else {
			out.Raw(json.Marshal(in.Params))
		}
	}
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix)
		if in.Id == nil {
			out.RawString("null")
		} else {
			out.Raw((*in.Id).MarshalJSON())
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Request) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC0e5e3f1EncodeGithubComSpace307Jsonrpc22(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Request) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC0e5e3f1EncodeGithubComSpace307Jsonrpc22(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Request) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC0e5e3f1DecodeGithubComSpace307Jsonrpc22(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Request) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC0e5e3f1DecodeGithubComSpace307Jsonrpc22(l, v)
}
