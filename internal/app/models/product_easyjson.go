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

func easyjsonCf3f67efDecodeGithubComGoParkMailRu20211YSNPInternalAppModels(in *jlexer.Lexer, out *WaitingReview) {
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
		case "product_id":
			out.ProductID = uint64(in.Uint64())
		case "product_name":
			out.ProductName = string(in.String())
		case "product_image":
			out.ProductImage = string(in.String())
		case "target_id":
			out.TargetID = uint64(in.Uint64())
		case "target_name":
			out.TargetName = string(in.String())
		case "target_avatar":
			out.TargetAvatar = string(in.String())
		case "type":
			out.Type = string(in.String())
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
func easyjsonCf3f67efEncodeGithubComGoParkMailRu20211YSNPInternalAppModels(out *jwriter.Writer, in WaitingReview) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"product_id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.ProductID))
	}
	{
		const prefix string = ",\"product_name\":"
		out.RawString(prefix)
		out.String(string(in.ProductName))
	}
	{
		const prefix string = ",\"product_image\":"
		out.RawString(prefix)
		out.String(string(in.ProductImage))
	}
	{
		const prefix string = ",\"target_id\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.TargetID))
	}
	{
		const prefix string = ",\"target_name\":"
		out.RawString(prefix)
		out.String(string(in.TargetName))
	}
	{
		const prefix string = ",\"target_avatar\":"
		out.RawString(prefix)
		out.String(string(in.TargetAvatar))
	}
	{
		const prefix string = ",\"type\":"
		out.RawString(prefix)
		out.String(string(in.Type))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v WaitingReview) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCf3f67efEncodeGithubComGoParkMailRu20211YSNPInternalAppModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v WaitingReview) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCf3f67efEncodeGithubComGoParkMailRu20211YSNPInternalAppModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *WaitingReview) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCf3f67efDecodeGithubComGoParkMailRu20211YSNPInternalAppModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *WaitingReview) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCf3f67efDecodeGithubComGoParkMailRu20211YSNPInternalAppModels(l, v)
}
func easyjsonCf3f67efDecodeGithubComGoParkMailRu20211YSNPInternalAppModels1(in *jlexer.Lexer, out *Review) {
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
		case "content":
			out.Content = string(in.String())
		case "rating":
			out.Rating = float32(in.Float32())
		case "reviewer_id":
			out.ReviewerID = uint64(in.Uint64())
		case "reviewer_name":
			out.ReviewerName = string(in.String())
		case "reviewer_avatar":
			out.ReviewerAvatar = string(in.String())
		case "product_id":
			out.ProductID = uint64(in.Uint64())
		case "product_name":
			out.ProductName = string(in.String())
		case "product_image":
			out.ProductImage = string(in.String())
		case "target_id":
			out.TargetID = uint64(in.Uint64())
		case "type":
			out.Type = string(in.String())
		case "creation_time":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.CreationTime).UnmarshalJSON(data))
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
func easyjsonCf3f67efEncodeGithubComGoParkMailRu20211YSNPInternalAppModels1(out *jwriter.Writer, in Review) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.ID))
	}
	{
		const prefix string = ",\"content\":"
		out.RawString(prefix)
		out.String(string(in.Content))
	}
	{
		const prefix string = ",\"rating\":"
		out.RawString(prefix)
		out.Float32(float32(in.Rating))
	}
	{
		const prefix string = ",\"reviewer_id\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.ReviewerID))
	}
	{
		const prefix string = ",\"reviewer_name\":"
		out.RawString(prefix)
		out.String(string(in.ReviewerName))
	}
	{
		const prefix string = ",\"reviewer_avatar\":"
		out.RawString(prefix)
		out.String(string(in.ReviewerAvatar))
	}
	{
		const prefix string = ",\"product_id\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.ProductID))
	}
	{
		const prefix string = ",\"product_name\":"
		out.RawString(prefix)
		out.String(string(in.ProductName))
	}
	{
		const prefix string = ",\"product_image\":"
		out.RawString(prefix)
		out.String(string(in.ProductImage))
	}
	{
		const prefix string = ",\"target_id\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.TargetID))
	}
	{
		const prefix string = ",\"type\":"
		out.RawString(prefix)
		out.String(string(in.Type))
	}
	{
		const prefix string = ",\"creation_time\":"
		out.RawString(prefix)
		out.Raw((in.CreationTime).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Review) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCf3f67efEncodeGithubComGoParkMailRu20211YSNPInternalAppModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Review) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCf3f67efEncodeGithubComGoParkMailRu20211YSNPInternalAppModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Review) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCf3f67efDecodeGithubComGoParkMailRu20211YSNPInternalAppModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Review) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCf3f67efDecodeGithubComGoParkMailRu20211YSNPInternalAppModels1(l, v)
}
func easyjsonCf3f67efDecodeGithubComGoParkMailRu20211YSNPInternalAppModels2(in *jlexer.Lexer, out *ProductListData) {
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
		case "date":
			out.Date = string(in.String())
		case "amount":
			out.Amount = int(in.Int())
		case "linkImages":
			if in.IsNull() {
				in.Skip()
				out.LinkImages = nil
			} else {
				in.Delim('[')
				if out.LinkImages == nil {
					if !in.IsDelim(']') {
						out.LinkImages = make([]string, 0, 4)
					} else {
						out.LinkImages = []string{}
					}
				} else {
					out.LinkImages = (out.LinkImages)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.LinkImages = append(out.LinkImages, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "userLiked":
			out.UserLiked = bool(in.Bool())
		case "tariff":
			out.Tariff = int(in.Int())
		case "close":
			out.Close = bool(in.Bool())
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
func easyjsonCf3f67efEncodeGithubComGoParkMailRu20211YSNPInternalAppModels2(out *jwriter.Writer, in ProductListData) {
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
		const prefix string = ",\"date\":"
		out.RawString(prefix)
		out.String(string(in.Date))
	}
	{
		const prefix string = ",\"amount\":"
		out.RawString(prefix)
		out.Int(int(in.Amount))
	}
	{
		const prefix string = ",\"linkImages\":"
		out.RawString(prefix)
		if in.LinkImages == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.LinkImages {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.String(string(v3))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"userLiked\":"
		out.RawString(prefix)
		out.Bool(bool(in.UserLiked))
	}
	{
		const prefix string = ",\"tariff\":"
		out.RawString(prefix)
		out.Int(int(in.Tariff))
	}
	{
		const prefix string = ",\"close\":"
		out.RawString(prefix)
		out.Bool(bool(in.Close))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ProductListData) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCf3f67efEncodeGithubComGoParkMailRu20211YSNPInternalAppModels2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ProductListData) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCf3f67efEncodeGithubComGoParkMailRu20211YSNPInternalAppModels2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ProductListData) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCf3f67efDecodeGithubComGoParkMailRu20211YSNPInternalAppModels2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ProductListData) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCf3f67efDecodeGithubComGoParkMailRu20211YSNPInternalAppModels2(l, v)
}
func easyjsonCf3f67efDecodeGithubComGoParkMailRu20211YSNPInternalAppModels3(in *jlexer.Lexer, out *ProductData) {
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
		case "date":
			out.Date = string(in.String())
		case "amount":
			out.Amount = int(in.Int())
		case "linkImages":
			if in.IsNull() {
				in.Skip()
				out.LinkImages = nil
			} else {
				in.Delim('[')
				if out.LinkImages == nil {
					if !in.IsDelim(']') {
						out.LinkImages = make([]string, 0, 4)
					} else {
						out.LinkImages = []string{}
					}
				} else {
					out.LinkImages = (out.LinkImages)[:0]
				}
				for !in.IsDelim(']') {
					var v4 string
					v4 = string(in.String())
					out.LinkImages = append(out.LinkImages, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "description":
			out.Description = string(in.String())
		case "category":
			out.Category = string(in.String())
		case "address":
			out.Address = string(in.String())
		case "longitude":
			out.Longitude = float64(in.Float64())
		case "latitude":
			out.Latitude = float64(in.Float64())
		case "views":
			out.Views = int(in.Int())
		case "likes":
			out.Likes = int(in.Int())
		case "tariff":
			out.Tariff = int(in.Int())
		case "ownerId":
			out.OwnerID = uint64(in.Uint64())
		case "ownerName":
			out.OwnerName = string(in.String())
		case "ownerSurname":
			out.OwnerSurname = string(in.String())
		case "ownerLinkImages":
			out.OwnerLinkImages = string(in.String())
		case "owner_rating":
			out.OwnerRating = float64(in.Float64())
		case "close":
			out.Close = bool(in.Bool())
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
func easyjsonCf3f67efEncodeGithubComGoParkMailRu20211YSNPInternalAppModels3(out *jwriter.Writer, in ProductData) {
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
		const prefix string = ",\"date\":"
		out.RawString(prefix)
		out.String(string(in.Date))
	}
	{
		const prefix string = ",\"amount\":"
		out.RawString(prefix)
		out.Int(int(in.Amount))
	}
	{
		const prefix string = ",\"linkImages\":"
		out.RawString(prefix)
		if in.LinkImages == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.LinkImages {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.String(string(v6))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"category\":"
		out.RawString(prefix)
		out.String(string(in.Category))
	}
	{
		const prefix string = ",\"address\":"
		out.RawString(prefix)
		out.String(string(in.Address))
	}
	{
		const prefix string = ",\"longitude\":"
		out.RawString(prefix)
		out.Float64(float64(in.Longitude))
	}
	{
		const prefix string = ",\"latitude\":"
		out.RawString(prefix)
		out.Float64(float64(in.Latitude))
	}
	{
		const prefix string = ",\"views\":"
		out.RawString(prefix)
		out.Int(int(in.Views))
	}
	{
		const prefix string = ",\"likes\":"
		out.RawString(prefix)
		out.Int(int(in.Likes))
	}
	{
		const prefix string = ",\"tariff\":"
		out.RawString(prefix)
		out.Int(int(in.Tariff))
	}
	{
		const prefix string = ",\"ownerId\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.OwnerID))
	}
	{
		const prefix string = ",\"ownerName\":"
		out.RawString(prefix)
		out.String(string(in.OwnerName))
	}
	{
		const prefix string = ",\"ownerSurname\":"
		out.RawString(prefix)
		out.String(string(in.OwnerSurname))
	}
	{
		const prefix string = ",\"ownerLinkImages\":"
		out.RawString(prefix)
		out.String(string(in.OwnerLinkImages))
	}
	{
		const prefix string = ",\"owner_rating\":"
		out.RawString(prefix)
		out.Float64(float64(in.OwnerRating))
	}
	{
		const prefix string = ",\"close\":"
		out.RawString(prefix)
		out.Bool(bool(in.Close))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ProductData) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCf3f67efEncodeGithubComGoParkMailRu20211YSNPInternalAppModels3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ProductData) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCf3f67efEncodeGithubComGoParkMailRu20211YSNPInternalAppModels3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ProductData) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCf3f67efDecodeGithubComGoParkMailRu20211YSNPInternalAppModels3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ProductData) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCf3f67efDecodeGithubComGoParkMailRu20211YSNPInternalAppModels3(l, v)
}
func easyjsonCf3f67efDecodeGithubComGoParkMailRu20211YSNPInternalAppModels4(in *jlexer.Lexer, out *Category) {
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
			out.Title = string(in.String())
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
func easyjsonCf3f67efEncodeGithubComGoParkMailRu20211YSNPInternalAppModels4(out *jwriter.Writer, in Category) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix[1:])
		out.String(string(in.Title))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Category) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCf3f67efEncodeGithubComGoParkMailRu20211YSNPInternalAppModels4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Category) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCf3f67efEncodeGithubComGoParkMailRu20211YSNPInternalAppModels4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Category) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCf3f67efDecodeGithubComGoParkMailRu20211YSNPInternalAppModels4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Category) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCf3f67efDecodeGithubComGoParkMailRu20211YSNPInternalAppModels4(l, v)
}
