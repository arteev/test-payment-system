package api

import (
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http/pprof"
	"test-payment-system/internal/app/payment/database"
)

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
	//subrouter := r.PathPrefix("/api/v1/payment").Subrouter()
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
