// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package apigwex

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

func easyjson22b64118DecodeGithubComChouandyGoexAwsexApigwex(in *jlexer.Lexer, out *Logger) {
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
		case "timestamp":
			out.Timestamp = string(in.String())
		case "request_id":
			out.RequestID = string(in.String())
		case "level":
			out.Level = string(in.String())
		case "status":
			out.Status = int(in.Int())
		case "method":
			out.Method = string(in.String())
		case "path":
			out.Path = string(in.String())
		case "latency":
			out.Latency = string(in.String())
		case "identity":
			if in.IsNull() {
				in.Skip()
				out.Identity = nil
			} else {
				if out.Identity == nil {
					out.Identity = new(Identity)
				}
				(*out.Identity).UnmarshalEasyJSON(in)
			}
		case "query_string_parameters":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.QueryStringParameters = make(map[string]string)
				} else {
					out.QueryStringParameters = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v1 string
					v1 = string(in.String())
					(out.QueryStringParameters)[key] = v1
					in.WantComma()
				}
				in.Delim('}')
			}
		case "path_parameters":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.PathParameters = make(map[string]string)
				} else {
					out.PathParameters = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v2 string
					v2 = string(in.String())
					(out.PathParameters)[key] = v2
					in.WantComma()
				}
				in.Delim('}')
			}
		case "body":
			out.Body = string(in.String())
		case "error":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Error).UnmarshalJSON(data))
			}
		case "metadata":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Metadata).UnmarshalJSON(data))
			}
		case "location":
			out.Location = string(in.String())
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
func easyjson22b64118EncodeGithubComChouandyGoexAwsexApigwex(out *jwriter.Writer, in Logger) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"timestamp\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Timestamp))
	}
	{
		const prefix string = ",\"request_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.RequestID))
	}
	{
		const prefix string = ",\"level\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Level))
	}
	{
		const prefix string = ",\"status\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Status))
	}
	{
		const prefix string = ",\"method\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Method))
	}
	{
		const prefix string = ",\"path\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Path))
	}
	{
		const prefix string = ",\"latency\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Latency))
	}
	if in.Identity != nil {
		const prefix string = ",\"identity\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Identity).MarshalEasyJSON(out)
	}
	if len(in.QueryStringParameters) != 0 {
		const prefix string = ",\"query_string_parameters\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('{')
			v3First := true
			for v3Name, v3Value := range in.QueryStringParameters {
				if v3First {
					v3First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v3Name))
				out.RawByte(':')
				out.String(string(v3Value))
			}
			out.RawByte('}')
		}
	}
	if len(in.PathParameters) != 0 {
		const prefix string = ",\"path_parameters\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('{')
			v4First := true
			for v4Name, v4Value := range in.PathParameters {
				if v4First {
					v4First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v4Name))
				out.RawByte(':')
				out.String(string(v4Value))
			}
			out.RawByte('}')
		}
	}
	if in.Body != "" {
		const prefix string = ",\"body\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Body))
	}
	if len(in.Error) != 0 {
		const prefix string = ",\"error\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.Error).MarshalJSON())
	}
	if len(in.Metadata) != 0 {
		const prefix string = ",\"metadata\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.Metadata).MarshalJSON())
	}
	if in.Location != "" {
		const prefix string = ",\"location\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Location))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Logger) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson22b64118EncodeGithubComChouandyGoexAwsexApigwex(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Logger) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson22b64118EncodeGithubComChouandyGoexAwsexApigwex(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Logger) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson22b64118DecodeGithubComChouandyGoexAwsexApigwex(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Logger) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson22b64118DecodeGithubComChouandyGoexAwsexApigwex(l, v)
}
func easyjson22b64118DecodeGithubComChouandyGoexAwsexApigwex1(in *jlexer.Lexer, out *Identity) {
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
		case "account_id":
			out.AccountID = string(in.String())
		case "source_ip":
			out.SourceIP = string(in.String())
		case "user_arn":
			out.UserArn = string(in.String())
		case "user_agent":
			out.UserAgent = string(in.String())
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
func easyjson22b64118EncodeGithubComChouandyGoexAwsexApigwex1(out *jwriter.Writer, in Identity) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"account_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.AccountID))
	}
	{
		const prefix string = ",\"source_ip\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.SourceIP))
	}
	{
		const prefix string = ",\"user_arn\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.UserArn))
	}
	{
		const prefix string = ",\"user_agent\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.UserAgent))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Identity) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson22b64118EncodeGithubComChouandyGoexAwsexApigwex1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Identity) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson22b64118EncodeGithubComChouandyGoexAwsexApigwex1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Identity) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson22b64118DecodeGithubComChouandyGoexAwsexApigwex1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Identity) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson22b64118DecodeGithubComChouandyGoexAwsexApigwex1(l, v)
}
