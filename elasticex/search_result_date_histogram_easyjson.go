// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package elasticex

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

func easyjson87940f94DecodeGithubComChouandyGoexElasticex(in *jlexer.Lexer, out *SearchDateHistogramSumItem) {
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
		case "value":
			out.Value = float64(in.Float64())
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
func easyjson87940f94EncodeGithubComChouandyGoexElasticex(out *jwriter.Writer, in SearchDateHistogramSumItem) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"value\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.Value))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SearchDateHistogramSumItem) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson87940f94EncodeGithubComChouandyGoexElasticex(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SearchDateHistogramSumItem) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson87940f94EncodeGithubComChouandyGoexElasticex(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SearchDateHistogramSumItem) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson87940f94DecodeGithubComChouandyGoexElasticex(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SearchDateHistogramSumItem) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson87940f94DecodeGithubComChouandyGoexElasticex(l, v)
}
func easyjson87940f94DecodeGithubComChouandyGoexElasticex1(in *jlexer.Lexer, out *SearchDateHistogramItem) {
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
			out.Timestamp = int64(in.Int64())
		case "datetime":
			out.Datetime = string(in.String())
		case "value":
			out.Value = int64(in.Int64())
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
func easyjson87940f94EncodeGithubComChouandyGoexElasticex1(out *jwriter.Writer, in SearchDateHistogramItem) {
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
		out.Int64(int64(in.Timestamp))
	}
	{
		const prefix string = ",\"datetime\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Datetime))
	}
	{
		const prefix string = ",\"value\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.Value))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SearchDateHistogramItem) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson87940f94EncodeGithubComChouandyGoexElasticex1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SearchDateHistogramItem) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson87940f94EncodeGithubComChouandyGoexElasticex1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SearchDateHistogramItem) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson87940f94DecodeGithubComChouandyGoexElasticex1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SearchDateHistogramItem) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson87940f94DecodeGithubComChouandyGoexElasticex1(l, v)
}