package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-faster/errors"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	api "github.com/vmedinskiy/highload/api/generated"
	"github.com/vmedinskiy/highload/internal/api/handlers"
	"github.com/vmedinskiy/highload/internal/api/middlewares"
	"github.com/vmedinskiy/highload/internal/pkg/config"
	"github.com/vmedinskiy/highload/internal/pkg/database"
	"github.com/vmedinskiy/highload/internal/pkg/dbutils"
	"github.com/vmedinskiy/highload/internal/pkg/log"
)

type Server struct {
	server *http.Server
	config *config.Config
	db     *sqlx.DB
}

func RunCmd(cmd *cobra.Command, _ []string) error {
	configFile, err := cmd.Flags().GetString("config")
	if err != nil {
		return errors.Wrap(err, "getting config flag")
	}
	cfg := &config.Config{}
	viper.SetConfigFile(configFile)

	if err := viper.ReadInConfig(); err != nil {
		return errors.Wrap(err, "read config file")
	}
	if err := viper.Unmarshal(cfg); err != nil {
		return errors.Wrap(err, "unmarshal config")
	}
	if err := log.Init(cfg.LogLevel); err != nil {
		return errors.Wrap(err, "init logger")
	}
	defer log.Sync()
	log.L().Infof("loaded Config %#v", cfg)
	srv := newServer(cfg)
	srv.Run()
	return nil
}

func (s *Server) Run() {
	serverCtx, serverStopCtx := context.WithCancel(context.Background())
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig
		shutdownCtx, cancel := context.WithTimeout(serverCtx, 10*time.Second)
		defer cancel()
		go func() {
			<-shutdownCtx.Done()
			if errors.Is(shutdownCtx.Err(), context.DeadlineExceeded) {
				log.L().Fatal("graceful shutdown timed out.. forcing exit.")
			}
		}()

		err := s.Shutdown(shutdownCtx)
		if err != nil {
			log.L().Fatal(err)
		}
		log.L().Info("graceful shutdown complete.")
		serverStopCtx()
	}()
	log.L().Infof("starting server at...%s", s.config.Port)
	err := s.server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.L().Fatal(err)
	}
	<-serverCtx.Done()
}

func (s *Server) initDB() {
	db, err := database.New(s.config, dbutils.NewLogger())
	if err != nil {
		log.L().Fatal(err)
	}
	s.db = db
	err = s.checkDBStructure()
	if err != nil {
		log.L().Fatal(err)
	}
}

func (s *Server) checkDBStructure() error {
	/*q := ``
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	args := make(map[string]any)
	return dbutils.NamedExec(ctx, s.db, q, args)*/
	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	if s.db != nil {
		err := s.db.Close()
		if err != nil {
			log.C(ctx).Fatalf("failed closing db: %s", err.Error())
		}
		log.C(ctx).Debug("db closed")
	}
	return s.server.Shutdown(ctx) //nolint:wrapcheck
}

func newServer(cfg *config.Config) *Server {
	r := chi.NewRouter()
	returningServer := &Server{
		server: &http.Server{
			Addr:              ":" + cfg.Port,
			Handler:           r,
			ReadHeaderTimeout: 3 * time.Second,
		},
		config: cfg,
	}

	// returningServer.initToken()
	// returningServer.setupGlobalMiddleware(r)
	returningServer.setupHandlers(r)
	returningServer.initDB()
	return returningServer
}

func (s *Server) setupHandlers(r *chi.Mux) {
	apiHandler, err := api.NewServer(handlers.NewServerHandler(), nil)
	if err != nil {
		log.L().Fatal(err)
	}
	r.Route("/api", func(r chi.Router) {
		r.Use(middlewares.UrlHijack)
		r.Mount("/", apiHandler)
	})
}
