package api

import (
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"net/http/pprof"
	"test-payment-system/internal/app/payment/database"
	"test-payment-system/internal/app/payment/dto"
	"test-payment-system/internal/pkg/config"
	"test-payment-system/internal/pkg/service"
	"test-payment-system/pkg/version"
)

const PathAPIPrefix = "/api/v1/payment"

type API struct {
	log *zap.SugaredLogger
	db  database.Database
}

func New(log *zap.SugaredLogger, db database.Database) *API {
	newAPI := &API{
		log: log,
		db:  db,
	}
	return newAPI
}

func (a *API) AddDebugHandler(r *mux.Router, prefix string) {
	subrouter := r.PathPrefix(prefix).Subrouter()
	subrouter.HandleFunc("/debug/pprof/*", pprof.Index)
	subrouter.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	subrouter.HandleFunc("/debug/pprof/profile", pprof.Profile)
	subrouter.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	subrouter.HandleFunc("/debug/pprof/trace", pprof.Trace)
}

func (a *API) GetRoutes(r *mux.Router) *mux.Router {
	if config.CurrentMode == config.ModeDevelopment {
		a.AddDebugHandler(r, PathAPIPrefix)
	}
	routerInternal := r.PathPrefix("/api/v1/internal/payment").Subrouter()
	routerInternal.HandleFunc("/version", service.ToJSONResponse(version.GetVersionHandler)).
		Methods(http.MethodGet, http.MethodOptions)

	subrouter := r.PathPrefix(PathAPIPrefix).Subrouter()
	subrouter.HandleFunc("/wallet", service.ToJSONDataObjectRequestResponse(
		a.NewWallet, &dto.NewWalletRequestMeta{},
	)).Methods(http.MethodPost, http.MethodOptions)

	subrouter.HandleFunc("/wallet", service.ToJSONResponse(a.GetWallet)).
		Methods(http.MethodGet, http.MethodOptions)
	return r
}

func (a *API) Close() error {
	return nil
}

func checkAudiences(_ string, audiences []string) bool {
	if len(audiences) == 0 {
		return true
	}
	return false
}
