package server

import (
	"bytes"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/transport/http"
	"io"
	http1 "net/http"
	"strings"
)

func ResponseEncoder(w http.ResponseWriter, r *http.Request, v interface{}) error {
	if v == nil {
		return nil
	}
	if rd, ok := v.(http.Redirector); ok {
		url, code := rd.Redirect()
		http1.Redirect(w, r, url, code)
		return nil
	}

	codec, _ := http.CodecForRequest(r, "Accept")

	type Aa struct {
		Code    int          `json:"code"`
		Result  *interface{} `json:"result,omitempty"`
		Message string       `json:"Message"`
		Type    string       `json:"type"`
	}

	item := &Aa{
		Result:  &v,
		Message: "ok",
		Type:    "success",
	}

	data, err := codec.Marshal(item)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", ContentType(codec.Name()))
	_, err = w.Write(data)
	if err != nil {
		return err
	}
	return nil
}

const (
	baseContentType = "application"
)

// ContentType returns the content-type with base prefix.
func ContentType(subtype string) string {
	return strings.Join([]string{baseContentType, subtype}, "/")
}

// RequestDecoder 自定义请求格式
func RequestDecoder(r *http.Request, v interface{}) error {
	codec, ok := http.CodecForRequest(r, "Content-Type")
	if !ok {
		return errors.BadRequest("CODEC", fmt.Sprintf("unregister Content-Type: %s", r.Header.Get("Content-Type")))
	}
	data, err := io.ReadAll(r.Body)

	// reset body.
	r.Body = io.NopCloser(bytes.NewBuffer(data))

	if err != nil {
		return errors.BadRequest("CODEC", err.Error())
	}
	if len(data) == 0 {
		return nil
	}
	if err = codec.Unmarshal(data, v); err != nil {
		return errors.BadRequest("CODEC", fmt.Sprintf("body unmarshal %s", err.Error()))
	}
	return nil
}
