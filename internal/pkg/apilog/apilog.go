package apilog

import (
	"context"
	"go.uber.org/zap"
	"net/http"
	"strings"
	"test-payment-system/pkg/requestid"
)

func LogStart(ctx context.Context, log *zap.SugaredLogger, keysAndValues ...interface{}) *zap.SugaredLogger {
	requestID := requestid.RequestIDFromContext(ctx)
	log = log.Desugar().WithOptions(zap.AddCallerSkip(1)).Sugar().
		With(zap.String("requestID", requestID.String()))
	log.Infow("started", keysAndValues...)
	return log
}

func LogFinish(ctx context.Context, log *zap.SugaredLogger, response interface{}, err error, keysAndValues ...interface{}) {
	if err == nil {
		log.Infow("finished", keysAndValues...)
		return
	}

	if are, ok := err.(interface{ ExtraFields() map[string]interface{} }); ok {
		extras := are.ExtraFields()
		arrExtras := make([]interface{}, 0, len(extras))
		for key, value := range are.ExtraFields() {
			arrExtras = append(arrExtras, zap.Any(key, value))
		}
		log = log.With(arrExtras...)
	}
	if len(keysAndValues) > 0 {
		log.Errorw(err.Error(), append([]interface{}{zap.Error(err), keysAndValues}))
	} else {
		log.Errorw(err.Error(), zap.Error(err))
	}
}

func ValuesFromRequest(r *http.Request) []interface{} {
	var values []interface{}
	r.ParseForm()
	for k, v := range r.Form {
		values = append(values, k, strings.Join(v, ";"))
	}
	return values
}
