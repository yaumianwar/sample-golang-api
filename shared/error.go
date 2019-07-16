package shared

import (
	"github.com/go-errors/errors"
	"net/http"
)

// list of error message handling
var (
	ErrDataNotFound         = errors.New("data tidak ditemukan")
	ErrDeleteSkincareBefore = errors.New("data user tidak dapat dihapus : silahkan hapus data skincare terlebih dahulu")
)

// ErrorResponse model
type ErrorResponse struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
	Errors           error  `json:"errors,omitempty"`
	Status           int    `json:"status"`
}

// NewErrorResponse return error response from API
func NewErrorResponse(err error) ErrorResponse {
	var resp = ErrorResponse{ErrorDescription: err.Error()}
	switch err {
	case ErrDataNotFound:
		resp.Error = "bad_request"
		resp.Status = http.StatusNotFound
	case ErrDeleteSkincareBefore:
		resp.Error = "bad_request"
		resp.Status = http.StatusBadRequest
	default:
		resp.Error = "bad_request"
		resp.Status = http.StatusBadRequest
		resp.ErrorDescription = "unknown error. please contact system administrator"
	}

	return resp
}
