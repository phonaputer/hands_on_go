package server

import (
	"context"
	"github.com/phonaputer/hands_on_go/internal/config"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func Run() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	conf, err := config.LoadConfig()
	if err != nil {
		logrus.WithError(err).Fatal("failed to load config")
	}

	app, err := initializeUserServerApplication(conf)
	if err != nil {
		logrus.WithError(err).Fatal("failed to initialize application")
	}

	s := &http.Server{Addr: ":8080", Handler: app.userServerHTTPHandler}

	serverStopCh := make(chan struct{}) // channel to detect server stop
	signalCh := make(chan os.Signal, 1) // channel to detect OS signals
	signal.Notify(signalCh, syscall.SIGTERM, syscall.SIGINT)

	logrus.Info("listening on port 8080")

	go listenAndServe(s, serverStopCh)

	select { // wait for the first event on either of the two channels
	case <-signalCh:
		err = s.Shutdown(context.Background())
		if err != nil {
			logrus.WithError(err).Error("error gracefully shutting down server")
		}
	case <-serverStopCh:
	}

	logrus.Info("exiting")
}

func listenAndServe(s *http.Server, serverStopCh chan struct{}) {
	defer close(serverStopCh)

	err := s.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		logrus.WithError(err).Error("error running HTTP server")
	}
}
