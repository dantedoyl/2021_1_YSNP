// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package models

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

func easyjson9e1087fdDecodeGithubComGoParkMailRu20211YSNPInternalAppModels(in *jlexer.Lexer, out *UserOAuthRequest) {
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
		case "id":
			out.ID = uint64(in.Uint64())
		case "last_name":
			out.LastName = string(in.String())
		case "first_name":
			out.FirstName = string(in.String())
		case "photo_max":
			out.Photo = string(in.String())
		case "user_oauth_id":
			out.UserOAuthID = float64(in.Float64())
		case "user_oauth_type":
			out.UserOAuthType = string(in.String())
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
func easyjson9e1087fdEncodeGithubComGoParkMailRu20211YSNPInternalAppModels(out *jwriter.Writer, in UserOAuthRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.ID))
	}
	{
		const prefix string = ",\"last_name\":"
		out.RawString(prefix)
		out.String(string(in.LastName))
	}
	{
		const prefix string = ",\"first_name\":"
		out.RawString(prefix)
		out.String(string(in.FirstName))
	}
	{
		const prefix string = ",\"photo_max\":"
		out.RawString(prefix)
		out.String(string(in.Photo))
	}
	{
		const prefix string = ",\"user_oauth_id\":"
		out.RawString(prefix)
		out.Float64(float64(in.UserOAuthID))
	}
	{
		const prefix string = ",\"user_oauth_type\":"
		out.RawString(prefix)
		out.String(string(in.UserOAuthType))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserOAuthRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9e1087fdEncodeGithubComGoParkMailRu20211YSNPInternalAppModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserOAuthRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9e1087fdEncodeGithubComGoParkMailRu20211YSNPInternalAppModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserOAuthRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9e1087fdDecodeGithubComGoParkMailRu20211YSNPInternalAppModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserOAuthRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9e1087fdDecodeGithubComGoParkMailRu20211YSNPInternalAppModels(l, v)
}
func easyjson9e1087fdDecodeGithubComGoParkMailRu20211YSNPInternalAppModels1(in *jlexer.Lexer, out *UserData) {
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
		case "id":
			out.ID = uint64(in.Uint64())
		case "name":
			out.Name = string(in.String())
		case "surname":
			out.Surname = string(in.String())
		case "sex":
			out.Sex = string(in.String())
		case "email":
			out.Email = string(in.String())
		case "telephone":
			out.Telephone = string(in.String())
		case "password":
			out.Password = string(in.String())
		case "dateBirth":
			out.DateBirth = string(in.String())
		case "latitude":
			out.Latitude = float64(in.Float64())
		case "longitude":
			out.Longitude = float64(in.Float64())
		case "radius":
			out.Radius = uint64(in.Uint64())
		case "address":
			out.Address = string(in.String())
		case "linkImages":
			out.LinkImages = string(in.String())
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
func easyjson9e1087fdEncodeGithubComGoParkMailRu20211YSNPInternalAppModels1(out *jwriter.Writer, in UserData) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.ID))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"surname\":"
		out.RawString(prefix)
		out.String(string(in.Surname))
	}
	{
		const prefix string = ",\"sex\":"
		out.RawString(prefix)
		out.String(string(in.Sex))
	}
	{
		const prefix string = ",\"email\":"
		out.RawString(prefix)
		out.String(string(in.Email))
	}
	{
		const prefix string = ",\"telephone\":"
		out.RawString(prefix)
		out.String(string(in.Telephone))
	}
	if in.Password != "" {
		const prefix string = ",\"password\":"
		out.RawString(prefix)
		out.String(string(in.Password))
	}
	{
		const prefix string = ",\"dateBirth\":"
		out.RawString(prefix)
		out.String(string(in.DateBirth))
	}
	{
		const prefix string = ",\"latitude\":"
		out.RawString(prefix)
		out.Float64(float64(in.Latitude))
	}
	{
		const prefix string = ",\"longitude\":"
		out.RawString(prefix)
		out.Float64(float64(in.Longitude))
	}
	{
		const prefix string = ",\"radius\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.Radius))
	}
	{
		const prefix string = ",\"address\":"
		out.RawString(prefix)
		out.String(string(in.Address))
	}
	{
		const prefix string = ",\"linkImages\":"
		out.RawString(prefix)
		out.String(string(in.LinkImages))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserData) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9e1087fdEncodeGithubComGoParkMailRu20211YSNPInternalAppModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserData) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9e1087fdEncodeGithubComGoParkMailRu20211YSNPInternalAppModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserData) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9e1087fdDecodeGithubComGoParkMailRu20211YSNPInternalAppModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserData) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9e1087fdDecodeGithubComGoParkMailRu20211YSNPInternalAppModels1(l, v)
}
func easyjson9e1087fdDecodeGithubComGoParkMailRu20211YSNPInternalAppModels2(in *jlexer.Lexer, out *SignUpRequest) {
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
		case "name":
			out.Name = string(in.String())
		case "surname":
			out.Surname = string(in.String())
		case "sex":
			out.Sex = string(in.String())
		case "email":
			out.Email = string(in.String())
		case "telephone":
			out.Telephone = string(in.String())
		case "password1":
			out.Password1 = string(in.String())
		case "password2":
			out.Password2 = string(in.String())
		case "dateBirth":
			out.DateBirth = string(in.String())
		case "linkImages":
			out.LinkImages = string(in.String())
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
func easyjson9e1087fdEncodeGithubComGoParkMailRu20211YSNPInternalAppModels2(out *jwriter.Writer, in SignUpRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix[1:])
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"surname\":"
		out.RawString(prefix)
		out.String(string(in.Surname))
	}
	{
		const prefix string = ",\"sex\":"
		out.RawString(prefix)
		out.String(string(in.Sex))
	}
	{
		const prefix string = ",\"email\":"
		out.RawString(prefix)
		out.String(string(in.Email))
	}
	{
		const prefix string = ",\"telephone\":"
		out.RawString(prefix)
		out.String(string(in.Telephone))
	}
	{
		const prefix string = ",\"password1\":"
		out.RawString(prefix)
		out.String(string(in.Password1))
	}
	{
		const prefix string = ",\"password2\":"
		out.RawString(prefix)
		out.String(string(in.Password2))
	}
	{
		const prefix string = ",\"dateBirth\":"
		out.RawString(prefix)
		out.String(string(in.DateBirth))
	}
	{
		const prefix string = ",\"linkImages\":"
		out.RawString(prefix)
		out.String(string(in.LinkImages))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SignUpRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9e1087fdEncodeGithubComGoParkMailRu20211YSNPInternalAppModels2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SignUpRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9e1087fdEncodeGithubComGoParkMailRu20211YSNPInternalAppModels2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SignUpRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9e1087fdDecodeGithubComGoParkMailRu20211YSNPInternalAppModels2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SignUpRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9e1087fdDecodeGithubComGoParkMailRu20211YSNPInternalAppModels2(l, v)
}
func easyjson9e1087fdDecodeGithubComGoParkMailRu20211YSNPInternalAppModels3(in *jlexer.Lexer, out *SellerData) {
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
		case "id":
			out.ID = uint64(in.Uint64())
		case "name":
			out.Name = string(in.String())
		case "surname":
			out.Surname = string(in.String())
		case "telephone":
			out.Telephone = string(in.String())
		case "linkImages":
			out.LinkImages = string(in.String())
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
func easyjson9e1087fdEncodeGithubComGoParkMailRu20211YSNPInternalAppModels3(out *jwriter.Writer, in SellerData) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.ID))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"surname\":"
		out.RawString(prefix)
		out.String(string(in.Surname))
	}
	{
		const prefix string = ",\"telephone\":"
		out.RawString(prefix)
		out.String(string(in.Telephone))
	}
	{
		const prefix string = ",\"linkImages\":"
		out.RawString(prefix)
		out.String(string(in.LinkImages))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SellerData) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9e1087fdEncodeGithubComGoParkMailRu20211YSNPInternalAppModels3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SellerData) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9e1087fdEncodeGithubComGoParkMailRu20211YSNPInternalAppModels3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SellerData) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9e1087fdDecodeGithubComGoParkMailRu20211YSNPInternalAppModels3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SellerData) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9e1087fdDecodeGithubComGoParkMailRu20211YSNPInternalAppModels3(l, v)
}
func easyjson9e1087fdDecodeGithubComGoParkMailRu20211YSNPInternalAppModels4(in *jlexer.Lexer, out *Response) {
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
		case "response":
			if in.IsNull() {
				in.Skip()
				out.Response = nil
			} else {
				in.Delim('[')
				if out.Response == nil {
					if !in.IsDelim(']') {
						out.Response = make([]struct {
							LastName  string `json:"last_name"`
							FirstName string `json:"first_name"`
							Photo     string `json:"photo_max"`
						}, 0, 1)
					} else {
						out.Response = []struct {
							LastName  string `json:"last_name"`
							FirstName string `json:"first_name"`
							Photo     string `json:"photo_max"`
						}{}
					}
				} else {
					out.Response = (out.Response)[:0]
				}
				for !in.IsDelim(']') {
					var v1 struct {
						LastName  string `json:"last_name"`
						FirstName string `json:"first_name"`
						Photo     string `json:"photo_max"`
					}
					easyjson9e1087fdDecode(in, &v1)
					out.Response = append(out.Response, v1)
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
func easyjson9e1087fdEncodeGithubComGoParkMailRu20211YSNPInternalAppModels4(out *jwriter.Writer, in Response) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"response\":"
		out.RawString(prefix[1:])
		if in.Response == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Response {
				if v2 > 0 {
					out.RawByte(',')
				}
				easyjson9e1087fdEncode(out, v3)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Response) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9e1087fdEncodeGithubComGoParkMailRu20211YSNPInternalAppModels4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Response) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9e1087fdEncodeGithubComGoParkMailRu20211YSNPInternalAppModels4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Response) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9e1087fdDecodeGithubComGoParkMailRu20211YSNPInternalAppModels4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Response) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9e1087fdDecodeGithubComGoParkMailRu20211YSNPInternalAppModels4(l, v)
}
func easyjson9e1087fdDecode(in *jlexer.Lexer, out *struct {
	LastName  string `json:"last_name"`
	FirstName string `json:"first_name"`
	Photo     string `json:"photo_max"`
}) {
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
		case "last_name":
			out.LastName = string(in.String())
		case "first_name":
			out.FirstName = string(in.String())
		case "photo_max":
			out.Photo = string(in.String())
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
func easyjson9e1087fdEncode(out *jwriter.Writer, in struct {
	LastName  string `json:"last_name"`
	FirstName string `json:"first_name"`
	Photo     string `json:"photo_max"`
}) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"last_name\":"
		out.RawString(prefix[1:])
		out.String(string(in.LastName))
	}
	{
		const prefix string = ",\"first_name\":"
		out.RawString(prefix)
		out.String(string(in.FirstName))
	}
	{
		const prefix string = ",\"photo_max\":"
		out.RawString(prefix)
		out.String(string(in.Photo))
	}
	out.RawByte('}')
}
func easyjson9e1087fdDecodeGithubComGoParkMailRu20211YSNPInternalAppModels5(in *jlexer.Lexer, out *ProfileData) {
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
		case "id":
			out.ID = uint64(in.Uint64())
		case "name":
			out.Name = string(in.String())
		case "surname":
			out.Surname = string(in.String())
		case "sex":
			out.Sex = string(in.String())
		case "email":
			out.Email = string(in.String())
		case "telephone":
			out.Telephone = string(in.String())
		case "dateBirth":
			out.DateBirth = string(in.String())
		case "latitude":
			out.Latitude = float64(in.Float64())
		case "longitude":
			out.Longitude = float64(in.Float64())
		case "radius":
			out.Radius = uint64(in.Uint64())
		case "address":
			out.Address = string(in.String())
		case "linkImages":
			out.LinkImages = string(in.String())
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
func easyjson9e1087fdEncodeGithubComGoParkMailRu20211YSNPInternalAppModels5(out *jwriter.Writer, in ProfileData) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.ID))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"surname\":"
		out.RawString(prefix)
		out.String(string(in.Surname))
	}
	{
		const prefix string = ",\"sex\":"
		out.RawString(prefix)
		out.String(string(in.Sex))
	}
	{
		const prefix string = ",\"email\":"
		out.RawString(prefix)
		out.String(string(in.Email))
	}
	{
		const prefix string = ",\"telephone\":"
		out.RawString(prefix)
		out.String(string(in.Telephone))
	}
	{
		const prefix string = ",\"dateBirth\":"
		out.RawString(prefix)
		out.String(string(in.DateBirth))
	}
	{
		const prefix string = ",\"latitude\":"
		out.RawString(prefix)
		out.Float64(float64(in.Latitude))
	}
	{
		const prefix string = ",\"longitude\":"
		out.RawString(prefix)
		out.Float64(float64(in.Longitude))
	}
	{
		const prefix string = ",\"radius\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.Radius))
	}
	{
		const prefix string = ",\"address\":"
		out.RawString(prefix)
		out.String(string(in.Address))
	}
	{
		const prefix string = ",\"linkImages\":"
		out.RawString(prefix)
		out.String(string(in.LinkImages))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ProfileData) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9e1087fdEncodeGithubComGoParkMailRu20211YSNPInternalAppModels5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ProfileData) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9e1087fdEncodeGithubComGoParkMailRu20211YSNPInternalAppModels5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ProfileData) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9e1087fdDecodeGithubComGoParkMailRu20211YSNPInternalAppModels5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ProfileData) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9e1087fdDecodeGithubComGoParkMailRu20211YSNPInternalAppModels5(l, v)
}
func easyjson9e1087fdDecodeGithubComGoParkMailRu20211YSNPInternalAppModels6(in *jlexer.Lexer, out *PasswordChangeRequest) {
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
		case "oldPassword":
			out.OldPassword = string(in.String())
		case "newPassword1":
			out.NewPassword1 = string(in.String())
		case "newPassword2":
			out.NewPassword2 = string(in.String())
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
func easyjson9e1087fdEncodeGithubComGoParkMailRu20211YSNPInternalAppModels6(out *jwriter.Writer, in PasswordChangeRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"oldPassword\":"
		out.RawString(prefix[1:])
		out.String(string(in.OldPassword))
	}
	{
		const prefix string = ",\"newPassword1\":"
		out.RawString(prefix)
		out.String(string(in.NewPassword1))
	}
	{
		const prefix string = ",\"newPassword2\":"
		out.RawString(prefix)
		out.String(string(in.NewPassword2))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PasswordChangeRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9e1087fdEncodeGithubComGoParkMailRu20211YSNPInternalAppModels6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PasswordChangeRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9e1087fdEncodeGithubComGoParkMailRu20211YSNPInternalAppModels6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PasswordChangeRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9e1087fdDecodeGithubComGoParkMailRu20211YSNPInternalAppModels6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PasswordChangeRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9e1087fdDecodeGithubComGoParkMailRu20211YSNPInternalAppModels6(l, v)
}
func easyjson9e1087fdDecodeGithubComGoParkMailRu20211YSNPInternalAppModels7(in *jlexer.Lexer, out *LocationRequest) {
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
		case "latitude":
			out.Latitude = float64(in.Float64())
		case "longitude":
			out.Longitude = float64(in.Float64())
		case "radius":
			out.Radius = uint64(in.Uint64())
		case "address":
			out.Address = string(in.String())
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
func easyjson9e1087fdEncodeGithubComGoParkMailRu20211YSNPInternalAppModels7(out *jwriter.Writer, in LocationRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"latitude\":"
		out.RawString(prefix[1:])
		out.Float64(float64(in.Latitude))
	}
	{
		const prefix string = ",\"longitude\":"
		out.RawString(prefix)
		out.Float64(float64(in.Longitude))
	}
	{
		const prefix string = ",\"radius\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.Radius))
	}
	{
		const prefix string = ",\"address\":"
		out.RawString(prefix)
		out.String(string(in.Address))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v LocationRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9e1087fdEncodeGithubComGoParkMailRu20211YSNPInternalAppModels7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v LocationRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9e1087fdEncodeGithubComGoParkMailRu20211YSNPInternalAppModels7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *LocationRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9e1087fdDecodeGithubComGoParkMailRu20211YSNPInternalAppModels7(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *LocationRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9e1087fdDecodeGithubComGoParkMailRu20211YSNPInternalAppModels7(l, v)
}
func easyjson9e1087fdDecodeGithubComGoParkMailRu20211YSNPInternalAppModels8(in *jlexer.Lexer, out *Achievement) {
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
		case "title":
			out.Titie = string(in.String())
		case "description":
			out.Description = string(in.String())
		case "date":
			out.Date = string(in.String())
		case "link_pic":
			out.LinkPic = string(in.String())
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
func easyjson9e1087fdEncodeGithubComGoParkMailRu20211YSNPInternalAppModels8(out *jwriter.Writer, in Achievement) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix[1:])
		out.String(string(in.Titie))
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"date\":"
		out.RawString(prefix)
		out.String(string(in.Date))
	}
	{
		const prefix string = ",\"link_pic\":"
		out.RawString(prefix)
		out.String(string(in.LinkPic))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Achievement) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9e1087fdEncodeGithubComGoParkMailRu20211YSNPInternalAppModels8(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Achievement) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9e1087fdEncodeGithubComGoParkMailRu20211YSNPInternalAppModels8(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Achievement) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9e1087fdDecodeGithubComGoParkMailRu20211YSNPInternalAppModels8(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Achievement) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9e1087fdDecodeGithubComGoParkMailRu20211YSNPInternalAppModels8(l, v)
}
