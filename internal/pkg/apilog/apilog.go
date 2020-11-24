package apilog

import (
	"context"
	"net/http"
	"strings"
	"test-payment-system/pkg/requestid"

	"go.uber.org/zap"
)

const skipCaller = 1

func LogStart(ctx context.Context, log *zap.SugaredLogger, keysAndValues ...interface{}) *zap.SugaredLogger {
	requestID := requestid.FromContext(ctx)
	log = log.Desugar().WithOptions(zap.AddCallerSkip(skipCaller)).Sugar().
		With(zap.String("requestID", requestID.String()))
	log.Infow("started", keysAndValues...)
	return log
}

func LogFinish(ctx context.Context, log *zap.SugaredLogger, response interface{},
	err error, keysAndValues ...interface{}) {
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
		log.Errorw(err.Error(), append(keysAndValues, zap.Error(err)))
	} else {
		log.Errorw(err.Error(), zap.Error(err))
	}
}

func ValuesFromRequest(r *http.Request) []interface{} {
	values := make([]interface{}, 0, len(r.Form))
	if err := r.ParseForm(); err != nil {
		panic(err)
	}
	for k, v := range r.Form {
		values = append(values, k, strings.Join(v, ";"))
	}
	return values
}
