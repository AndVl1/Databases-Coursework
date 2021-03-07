// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package model

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

func easyjsonC80ae7adDecodeGithubComAndVl1BugTrackerBackendModel(in *jlexer.Lexer, out *Users) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(Users, 0, 8)
			} else {
				*out = Users{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 *User
			if in.IsNull() {
				in.Skip()
				v1 = nil
			} else {
				if v1 == nil {
					v1 = new(User)
				}
				(*v1).UnmarshalEasyJSON(in)
			}
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC80ae7adEncodeGithubComAndVl1BugTrackerBackendModel(out *jwriter.Writer, in Users) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
			if v2 > 0 {
				out.RawByte(',')
			}
			if v3 == nil {
				out.RawString("null")
			} else {
				(*v3).MarshalEasyJSON(out)
			}
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v Users) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC80ae7adEncodeGithubComAndVl1BugTrackerBackendModel(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Users) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC80ae7adEncodeGithubComAndVl1BugTrackerBackendModel(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Users) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC80ae7adDecodeGithubComAndVl1BugTrackerBackendModel(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Users) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC80ae7adDecodeGithubComAndVl1BugTrackerBackendModel(l, v)
}
func easyjsonC80ae7adDecodeGithubComAndVl1BugTrackerBackendModel1(in *jlexer.Lexer, out *User) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "userId":
			out.Id = uint64(in.Uint64())
		case "login":
			out.Login = string(in.String())
		case "password":
			out.Password = string(in.String())
		case "name":
			out.Name = string(in.String())
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
func easyjsonC80ae7adEncodeGithubComAndVl1BugTrackerBackendModel1(out *jwriter.Writer, in User) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"userId\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.Id))
	}
	{
		const prefix string = ",\"login\":"
		out.RawString(prefix)
		out.String(string(in.Login))
	}
	{
		const prefix string = ",\"password\":"
		out.RawString(prefix)
		out.String(string(in.Password))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v User) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC80ae7adEncodeGithubComAndVl1BugTrackerBackendModel1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v User) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC80ae7adEncodeGithubComAndVl1BugTrackerBackendModel1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *User) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC80ae7adDecodeGithubComAndVl1BugTrackerBackendModel1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *User) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC80ae7adDecodeGithubComAndVl1BugTrackerBackendModel1(l, v)
}
func easyjsonC80ae7adDecodeGithubComAndVl1BugTrackerBackendModel2(in *jlexer.Lexer, out *Labels) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(Labels, 0, 8)
			} else {
				*out = Labels{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v4 *Label
			if in.IsNull() {
				in.Skip()
				v4 = nil
			} else {
				if v4 == nil {
					v4 = new(Label)
				}
				(*v4).UnmarshalEasyJSON(in)
			}
			*out = append(*out, v4)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC80ae7adEncodeGithubComAndVl1BugTrackerBackendModel2(out *jwriter.Writer, in Labels) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v5, v6 := range in {
			if v5 > 0 {
				out.RawByte(',')
			}
			if v6 == nil {
				out.RawString("null")
			} else {
				(*v6).MarshalEasyJSON(out)
			}
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v Labels) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC80ae7adEncodeGithubComAndVl1BugTrackerBackendModel2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Labels) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC80ae7adEncodeGithubComAndVl1BugTrackerBackendModel2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Labels) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC80ae7adDecodeGithubComAndVl1BugTrackerBackendModel2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Labels) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC80ae7adDecodeGithubComAndVl1BugTrackerBackendModel2(l, v)
}
func easyjsonC80ae7adDecodeGithubComAndVl1BugTrackerBackendModel3(in *jlexer.Lexer, out *Label) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "labelId":
			out.LabelId = int(in.Int())
		case "labelName":
			out.LabelName = int(in.Int())
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
func easyjsonC80ae7adEncodeGithubComAndVl1BugTrackerBackendModel3(out *jwriter.Writer, in Label) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"labelId\":"
		out.RawString(prefix[1:])
		out.Int(int(in.LabelId))
	}
	{
		const prefix string = ",\"labelName\":"
		out.RawString(prefix)
		out.Int(int(in.LabelName))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Label) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC80ae7adEncodeGithubComAndVl1BugTrackerBackendModel3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Label) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC80ae7adEncodeGithubComAndVl1BugTrackerBackendModel3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Label) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC80ae7adDecodeGithubComAndVl1BugTrackerBackendModel3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Label) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC80ae7adDecodeGithubComAndVl1BugTrackerBackendModel3(l, v)
}
func easyjsonC80ae7adDecodeGithubComAndVl1BugTrackerBackendModel4(in *jlexer.Lexer, out *Comments) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(Comments, 0, 8)
			} else {
				*out = Comments{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v7 *Comment
			if in.IsNull() {
				in.Skip()
				v7 = nil
			} else {
				if v7 == nil {
					v7 = new(Comment)
				}
				(*v7).UnmarshalEasyJSON(in)
			}
			*out = append(*out, v7)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC80ae7adEncodeGithubComAndVl1BugTrackerBackendModel4(out *jwriter.Writer, in Comments) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v8, v9 := range in {
			if v8 > 0 {
				out.RawByte(',')
			}
			if v9 == nil {
				out.RawString("null")
			} else {
				(*v9).MarshalEasyJSON(out)
			}
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v Comments) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC80ae7adEncodeGithubComAndVl1BugTrackerBackendModel4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Comments) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC80ae7adEncodeGithubComAndVl1BugTrackerBackendModel4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Comments) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC80ae7adDecodeGithubComAndVl1BugTrackerBackendModel4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Comments) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC80ae7adDecodeGithubComAndVl1BugTrackerBackendModel4(l, v)
}
func easyjsonC80ae7adDecodeGithubComAndVl1BugTrackerBackendModel5(in *jlexer.Lexer, out *Comment) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "commentId":
			out.Id = int(in.Int())
		case "text":
			out.Text = string(in.String())
		case "attachmentPath":
			out.AttachmentPath = string(in.String())
		case "authorId":
			out.AuthorId = int(in.Int())
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
func easyjsonC80ae7adEncodeGithubComAndVl1BugTrackerBackendModel5(out *jwriter.Writer, in Comment) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"commentId\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Id))
	}
	{
		const prefix string = ",\"text\":"
		out.RawString(prefix)
		out.String(string(in.Text))
	}
	{
		const prefix string = ",\"attachmentPath\":"
		out.RawString(prefix)
		out.String(string(in.AttachmentPath))
	}
	{
		const prefix string = ",\"authorId\":"
		out.RawString(prefix)
		out.Int(int(in.AuthorId))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Comment) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC80ae7adEncodeGithubComAndVl1BugTrackerBackendModel5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Comment) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC80ae7adEncodeGithubComAndVl1BugTrackerBackendModel5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Comment) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC80ae7adDecodeGithubComAndVl1BugTrackerBackendModel5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Comment) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC80ae7adDecodeGithubComAndVl1BugTrackerBackendModel5(l, v)
}
func easyjsonC80ae7adDecodeGithubComAndVl1BugTrackerBackendModel6(in *jlexer.Lexer, out *Bugs) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(Bugs, 0, 8)
			} else {
				*out = Bugs{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v10 *Bug
			if in.IsNull() {
				in.Skip()
				v10 = nil
			} else {
				if v10 == nil {
					v10 = new(Bug)
				}
				(*v10).UnmarshalEasyJSON(in)
			}
			*out = append(*out, v10)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC80ae7adEncodeGithubComAndVl1BugTrackerBackendModel6(out *jwriter.Writer, in Bugs) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v11, v12 := range in {
			if v11 > 0 {
				out.RawByte(',')
			}
			if v12 == nil {
				out.RawString("null")
			} else {
				(*v12).MarshalEasyJSON(out)
			}
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v Bugs) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC80ae7adEncodeGithubComAndVl1BugTrackerBackendModel6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Bugs) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC80ae7adEncodeGithubComAndVl1BugTrackerBackendModel6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Bugs) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC80ae7adDecodeGithubComAndVl1BugTrackerBackendModel6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Bugs) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC80ae7adDecodeGithubComAndVl1BugTrackerBackendModel6(l, v)
}
func easyjsonC80ae7adDecodeGithubComAndVl1BugTrackerBackendModel7(in *jlexer.Lexer, out *Bug) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "bugId":
			out.Id = int(in.Int())
		case "name":
			out.Name = string(in.String())
		case "description":
			out.Description = string(in.String())
		case "status":
			out.Status = string(in.String())
		case "authorId":
			out.AuthorId = int(in.Int())
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
func easyjsonC80ae7adEncodeGithubComAndVl1BugTrackerBackendModel7(out *jwriter.Writer, in Bug) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"bugId\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Id))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"status\":"
		out.RawString(prefix)
		out.String(string(in.Status))
	}
	{
		const prefix string = ",\"authorId\":"
		out.RawString(prefix)
		out.Int(int(in.AuthorId))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Bug) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC80ae7adEncodeGithubComAndVl1BugTrackerBackendModel7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Bug) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC80ae7adEncodeGithubComAndVl1BugTrackerBackendModel7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Bug) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC80ae7adDecodeGithubComAndVl1BugTrackerBackendModel7(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Bug) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC80ae7adDecodeGithubComAndVl1BugTrackerBackendModel7(l, v)
}
func easyjsonC80ae7adDecodeGithubComAndVl1BugTrackerBackendModel8(in *jlexer.Lexer, out *Attachments) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(Attachments, 0, 8)
			} else {
				*out = Attachments{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v13 *Attachment
			if in.IsNull() {
				in.Skip()
				v13 = nil
			} else {
				if v13 == nil {
					v13 = new(Attachment)
				}
				(*v13).UnmarshalEasyJSON(in)
			}
			*out = append(*out, v13)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC80ae7adEncodeGithubComAndVl1BugTrackerBackendModel8(out *jwriter.Writer, in Attachments) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v14, v15 := range in {
			if v14 > 0 {
				out.RawByte(',')
			}
			if v15 == nil {
				out.RawString("null")
			} else {
				(*v15).MarshalEasyJSON(out)
			}
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v Attachments) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC80ae7adEncodeGithubComAndVl1BugTrackerBackendModel8(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Attachments) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC80ae7adEncodeGithubComAndVl1BugTrackerBackendModel8(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Attachments) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC80ae7adDecodeGithubComAndVl1BugTrackerBackendModel8(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Attachments) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC80ae7adDecodeGithubComAndVl1BugTrackerBackendModel8(l, v)
}
func easyjsonC80ae7adDecodeGithubComAndVl1BugTrackerBackendModel9(in *jlexer.Lexer, out *Attachment) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "attachmentId":
			out.Id = int(in.Int())
		case "authorId":
			out.AuthorId = int(in.Int())
		case "attachmentPath":
			out.AttachmentPath = string(in.String())
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
func easyjsonC80ae7adEncodeGithubComAndVl1BugTrackerBackendModel9(out *jwriter.Writer, in Attachment) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"attachmentId\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Id))
	}
	{
		const prefix string = ",\"authorId\":"
		out.RawString(prefix)
		out.Int(int(in.AuthorId))
	}
	{
		const prefix string = ",\"attachmentPath\":"
		out.RawString(prefix)
		out.String(string(in.AttachmentPath))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Attachment) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC80ae7adEncodeGithubComAndVl1BugTrackerBackendModel9(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Attachment) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC80ae7adEncodeGithubComAndVl1BugTrackerBackendModel9(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Attachment) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC80ae7adDecodeGithubComAndVl1BugTrackerBackendModel9(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Attachment) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC80ae7adDecodeGithubComAndVl1BugTrackerBackendModel9(l, v)
}
