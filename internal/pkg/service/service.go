package service

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"test-payment-system/internal/pkg/config"
	"time"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

const (
	shutdownTimeout = time.Second * 2
)

type API interface {
	GetRoutes(r *mux.Router) *mux.Router
	Close() error
}

// Service basic service for real service implementation
// Starts http service, expects completion signals
type Service struct {
	config  *config.API
	log     *zap.SugaredLogger
	api     API
	rootCtx context.Context
}

// New creates and returns base service
func New(cfg *config.API, log *zap.SugaredLogger, api API) *Service {
	newService := &Service{
		config: cfg,
		log:    log,
		api:    api,
	}
	newService.rootCtx = newService.handleSignals(context.Background())
	return newService
}

func (s *Service) GetRouters() *mux.Router {
	router := mux.NewRouter()

	return s.addHandlers(s.api.GetRoutes(router))
}

// Signal Handling
func (s *Service) handleSignals(ctx context.Context) context.Context {
	signals := make(chan os.Signal, 1)
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
		sig := <-signals
		defer signal.Stop(signals)
		s.log.Debug("signal:", sig)
		cancel()
	}()
	return ctx
}

func (s *Service) GetRootContext() context.Context {
	return s.rootCtx
}

// Start launch of the service and all its dependencies
func (s *Service) Start() error {
	defer s.done()
	if are, ok := s.api.(interface {
		Start(ctx context.Context) error
	}); ok {
		err := are.Start(s.GetRootContext())
		if err != nil {
			return err
		}
	}
	routers := s.GetRouters()
	routers.Use(RequestScope)
	loggerHandler := s.getLoggerHander(routers)

	server := &http.Server{
		Handler: loggerHandler,
		Addr:    s.config.Host + ":" + strconv.Itoa(s.config.Port),
	}
	s.log.Infow("listen", "host", s.config.Host, "port", s.config.Port)
	errs := make(chan error, 1)
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errs <- err
		}
	}()
	// Wait for interrupt signal or server error
	select {
	case <-s.rootCtx.Done():
		// Shutdown the server gracefully
		ctx, cancel := context.WithTimeout(s.rootCtx, shutdownTimeout)
		defer cancel()
		return server.Shutdown(ctx)

	case err := <-errs:
		s.log.Errorw("serve listener is stopped", zap.Error(err))
		return err
	}
}

func (s *Service) getLoggerHander(routers *mux.Router) http.Handler {
	loggerHandler := s.log.Desugar().WithOptions(
		zap.WithCaller(false),
	).Sugar()
	loggedRouter := loggingHandler(loggerHandler, routers)
	return loggedRouter
}

func (s *Service) done() {
	if err := s.api.Close(); err != nil {
		s.log.Errorw("failed to close api handlers", zap.Error(err))
	}
}

func (s *Service) addHandlers(r *mux.Router) *mux.Router {
	return r
}
