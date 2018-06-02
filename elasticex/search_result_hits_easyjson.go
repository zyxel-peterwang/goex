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

func easyjsonD77b37e1DecodeGithubComChouandyGoexElasticex(in *jlexer.Lexer, out *SearchHits) {
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
		case "search_after":
			if m, ok := out.SearchAfter.(easyjson.Unmarshaler); ok {
				m.UnmarshalEasyJSON(in)
			} else if m, ok := out.SearchAfter.(json.Unmarshaler); ok {
				_ = m.UnmarshalJSON(in.Raw())
			} else {
				out.SearchAfter = in.Interface()
			}
		case "hits":
			if in.IsNull() {
				in.Skip()
				out.Hits = nil
			} else {
				in.Delim('[')
				if out.Hits == nil {
					if !in.IsDelim(']') {
						out.Hits = make([]json.RawMessage, 0, 2)
					} else {
						out.Hits = []json.RawMessage{}
					}
				} else {
					out.Hits = (out.Hits)[:0]
				}
				for !in.IsDelim(']') {
					var v1 json.RawMessage
					if data := in.Raw(); in.Ok() {
						in.AddError((v1).UnmarshalJSON(data))
					}
					out.Hits = append(out.Hits, v1)
					in.WantComma()
				}
				in.Delim(']')
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
func easyjsonD77b37e1EncodeGithubComChouandyGoexElasticex(out *jwriter.Writer, in SearchHits) {
	out.RawByte('{')
	first := true
	_ = first
	if in.SearchAfter != nil {
		const prefix string = ",\"search_after\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if m, ok := in.SearchAfter.(easyjson.Marshaler); ok {
			m.MarshalEasyJSON(out)
		} else if m, ok := in.SearchAfter.(json.Marshaler); ok {
			out.Raw(m.MarshalJSON())
		} else {
			out.Raw(json.Marshal(in.SearchAfter))
		}
	}
	{
		const prefix string = ",\"hits\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Hits == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Hits {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.Raw((v3).MarshalJSON())
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SearchHits) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD77b37e1EncodeGithubComChouandyGoexElasticex(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SearchHits) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD77b37e1EncodeGithubComChouandyGoexElasticex(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SearchHits) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD77b37e1DecodeGithubComChouandyGoexElasticex(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SearchHits) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD77b37e1DecodeGithubComChouandyGoexElasticex(l, v)
}