package server

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/transport/http"
	httpStatus "github.com/go-kratos/kratos/v2/transport/http/status"
	"google.golang.org/grpc/status"
	http1 "net/http"
)

// HTTPError is an HTTP error.
type HTTPError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("HTTPError code: %d message: %s", e.Code, e.Message)
}

// FromError try to convert an error to *HTTPError.
func FromError(err error) *HTTPError {
	if err == nil {
		return nil
	}
	if se := new(HTTPError); errors.As(err, &se) {

		return se
	}
	fmt.Println(err)
	gs, ok := status.FromError(err)
	if !ok {
		fmt.Println("2")
		return &HTTPError{Code: 500, Message: err.Error()}
	}
	fmt.Println(gs)
	return &HTTPError{Code: httpStatus.FromGRPCCode(gs.Code()), Message: gs.Message()}
}

// ErrorEncoder 自定义错误返回
func ErrorEncoder(w http.ResponseWriter, r *http.Request, err error) {
	se := FromError(err)
	codec, _ := http.CodecForRequest(r, "Accept")
	body, err := codec.Marshal(se)
	if err != nil {
		w.WriteHeader(http1.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", ContentType(codec.Name()))
	w.WriteHeader(se.Code)
	_, _ = w.Write(body)
}
