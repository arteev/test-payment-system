package service

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"net/http"
	"strings"
	"test-payment-system/pkg/requestid"
)

type EmptyType struct{}

var Empty = &EmptyType{}

// Response for swagger generator
type Response struct {
	Data      interface{} `json:"data"`
	Success   bool        `json:"success"`
	Error     string      `json:"error,omitempty"`
	RequestID string      `json:"request_id,omitempty" example:"948b9acf-36c0-452d-af21-66b362778fa3"` //The X-Request-ID from request header. The request ID represented in the HTTP header X-Request-ID let you to link all the log lines which are common to a single web request.

	ExtraDetail struct {
		ID          string            `json:"id,omitempty" example:"1dQqPlQgJuPPJJfAd7pjmfBWMoP"` // Error Id in current request
		Code        string            `json:"code,omitempty" example:"1.1.1"`                     // Group error code
		ErrorOrigin string            `json:"error_origin,omitempty" example:"invalid parameter"` // Origin of error (group)
		Extra       map[string]string `json:"extra,omitempty"`                                    // Extra fields
	} `json:"error_detail,omitempty"`
}

type DataObject interface {
	Validate() error
}

type DataObjectMeta interface {
	New() DataObject
}

type DataObjectMetaNewFunc func() DataObject

func (f DataObjectMetaNewFunc) New() DataObject {
	return f()
}

type JSONResponderFunc func(ctx context.Context, r *http.Request) (interface{}, error)
type JSONDataRequestResponderFunc func(ctx context.Context, dataMeta DataObject) (interface{}, error)

func ToJSONResponse(handler JSONResponderFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		data, err := handler(ctx, r)
		RespondJSON(w, r, data, err)
	}
}

// Respond responds data or error as a response
func RespondJSON(w http.ResponseWriter, r *http.Request, data interface{}, err error) {
	const (
		keyRequestID = "request_id"
	)
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()
	requestID := requestid.RequestIDFromContext(ctx)

	if err != nil {
		data := make(map[string]interface{}, 1)
		data["error"] = processErrorMessage(err)
		data["success"] = false
		if !requestID.IsEmpty() {
			data[keyRequestID] = requestID.String()
		}
		if extra := putExtraFromError(err); extra != nil {
			data["error_detail"] = extra
		}
		//TODO: extend error: code, cause, etc.
		buf := bytes.NewBuffer(nil)
		json.NewEncoder(buf).Encode(data)
		w.WriteHeader(http.StatusBadRequest)
		buf.WriteTo(w)
		return
	}
	if data == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	commonData := map[string]interface{}{
		"success": true,
	}
	if !requestID.IsEmpty() {
		commonData[keyRequestID] = requestID.String()
	}

	if _, ok := data.(*EmptyType); !ok {
		keyData := "data"
		//if are, ok := data.(interface{ KeyDataResponse() string }); ok {
		//	keyData = are.KeyDataResponse()
		//}
		commonData[keyData] = data
	}

	if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
		w.Header().Set("Content-Encoding", "gzip")
		writer := gzip.NewWriter(w)
		defer writer.Close()
		json.NewEncoder(writer).Encode(commonData)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(commonData)
}

var ErrCommonInvalidParameter func(err error) error

func ToJSONDataObjectRequestResponse(handler JSONDataRequestResponderFunc, dataMeta DataObjectMeta) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		contentType := r.Header.Get("Content-type")
		if !strings.HasPrefix(contentType, "application/json") {
			http.Error(w, "Unexpected Content-Type.Expect application/json", 400)
			return
		}
		object := dataMeta.New()
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(object)
		if err != nil {
			http.Error(w, "error decoding", 500)
			return
		}
		if err = object.Validate(); err != nil {
			if ErrCommonInvalidParameter != nil {
				err = ErrCommonInvalidParameter(err)
			}
			RespondJSON(w, r, nil, err)
			return
		}
		ctx := r.Context()
		data, err := handler(ctx, object)
		RespondJSON(w, r, data, err)
	}
}

func processErrorMessage(err error) string {
	if err == nil {
		return ""
	}
	errCause := errors.Cause(err)
	var msg string
	switch value := errCause.(type) {
	case interface{ GetMessage() string }: //external error
		msg = value.GetMessage()
	default:
		msg = errCause.Error()
	}
	return msg
}

func putExtraFromError(err error) map[string]interface{} {
	if err == nil {
		return nil
	}
	if are, ok := err.(interface{ ExtraFields() map[string]interface{} }); ok {
		return are.ExtraFields()
	}
	return nil
}
