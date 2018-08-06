package main

import (
	"context"
	"database/sql"
	"errors"
	"flag"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/coreos/go-systemd/activation"
	"github.com/kerak19/template/internal/config"
	"github.com/kerak19/template/internal/controller"
	"github.com/kerak19/template/internal/flags"
	_ "github.com/lib/pq"
	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/postgres"
	_ "github.com/mattes/migrate/source/file"
	"github.com/onrik/logrus/filename"
	"github.com/sirupsen/logrus"
)

var (
	configsPath = []string{"configs/dev.toml"}
	log         = logrus.StandardLogger()
)

func init() {
	flag.Var((*flags.StringList)(&configsPath), "config", "Path to TOML configuration file")
	flag.Parse()
}

func debugLogger(debug bool) {
	if debug {
		log.AddHook(filename.NewHook())
		log.SetLevel(logrus.DebugLevel)
	}
}

func runMigrations(migrationsPath string, db *sql.DB) error {
	if migrationsPath == "" {
		return errors.New("missing migrations path")
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithDatabaseInstance("file://"+migrationsPath,
		"postgres", driver)
	if err != nil {
		return err
	}
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}

func connectDatabase(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return db, err
	}
	return db, db.Ping()
}

func run() int {

	config, err := config.ReadFromFiles(configsPath...)
	if err != nil {
		log.WithError(err).Error("Error while readin TOML configuration file")
		return -1
	}

	debugLogger(config.DebugMode)

	log.WithField("config", config).Debug("Configuration file loaded")

	db, err := connectDatabase(config.Database.URL)
	if err != nil {
		log.WithError(err).Error("Error while opening database connection")
		return -1
	}
	defer db.Close()

	err = runMigrations(config.Database.Migrations, db)
	if err != nil {
		log.WithError(err).Error("Error while running database migrations")
		return -1
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	controller := controller.Controller(ctx, db, config, log)

	listener, err := initListener(config.Addr)
	if err != nil {
		log.WithError(err).Error("Error while obtaining listener")
		return -1
	}

	server := &http.Server{
		Handler: controller,
	}

	go serve(server, listener)

	closeOnSignal(cancel, server, config.ServerShutdownTimeout)

	return 0
}

func initListener(addr string) (net.Listener, error) {
	if addr != "systemd" {
		return net.Listen("tcp", addr)
	}
	listeners, err := activation.Listeners()
	if err != nil {
		return nil, err
	}
	if len(listeners) < 1 {
		return nil, errors.New("couldn't get listeners")
	}
	return listeners[0], nil
}

func serve(s *http.Server, l net.Listener) {
	log.WithField("addr", l.Addr().String()).Info("App is listening")
	s.Serve(l)
}

var sigCh = make(chan os.Signal, 1)

func closeOnSignal(cancel context.CancelFunc, server *http.Server, timeout time.Duration) {
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh
	cancel()

	log.Info("Shutting server down")
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	err := server.Shutdown(ctx)
	if err != nil {
		log.WithError(err).Error("Error while shutting server down")
	}
}

func main() {
	os.Exit(run())
}
