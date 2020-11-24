package requestid

import (
	"context"
	"net/http"
	"test-payment-system/pkg/contextkey"

	uuid "github.com/satori/go.uuid"
)

const maxLenRequestID = 30

// A RequestID represents request ID.
type RequestID string

// FromContext obtains RequestID from given context.
func (rid *RequestID) FromContext(ctx context.Context) (ok bool) {
	*rid, ok = ctx.Value(contextkey.RequestIDKey).(RequestID)
	return ok
}

// ToContext sets the RequestID to given context returning new context.
func (rid RequestID) ToContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, contextkey.RequestIDKey, rid)
}

// String returns the RequestID as string.
func (rid RequestID) String() string {
	return string(rid)
}

// NewRequestID generates new request id
// that is UUID v4.
func NewRequestID() RequestID {
	uid := uuid.NewV4()
	return RequestID(uid.String())
}

func NewRequestIDWarn() RequestID {
	uid := uuid.NewV4()
	return RequestID("warn-" + uid.String())
}

func NewFromString(id string) RequestID {
	return RequestID(id)
}

func AcceptRequestID(ctx context.Context) RequestID {
	var requestID string

	requestRaw := ctx.Value(contextkey.RequestIDKey)
	if requestRaw != nil {
		return requestRaw.(RequestID)
	}

	if requestID != "" && len(requestID) <= maxLenRequestID {
		return RequestID(requestID)
	}
	return NewRequestID()
}

func NewOrFromRequest(r *http.Request) (rid RequestID) {
	value := r.Header.Get("X-Request-ID")
	if value == "" {
		return NewRequestID()
	}
	return NewFromString(value)
}

// FromContext obtains RequestID from given context
// returning it. If the given context has no request id, then
// blank request it will be returned.
func FromContext(ctx context.Context) (rid RequestID) {
	rid.FromContext(ctx)
	return rid
}

func (rid RequestID) IsEmpty() bool {
	return rid == ""
}

func (rid RequestID) ExpandValues(values map[string]interface{}) {
	if !rid.IsEmpty() {
		values["request_id"] = rid
	}
}
