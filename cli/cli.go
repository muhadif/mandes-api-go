package cli

import (
	"context"
	"fmt"
	"github.com/muhadif/mandes/app"
	"github.com/muhadif/mandes/route"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run() {
	deps := app.InitDependency()
	apps := app.NewApp(deps)

	routes := route.NewRouter(apps)

	srv := &http.Server{
		Handler: routes,
		Addr:    fmt.Sprintf(":%s", deps.Cfg.HttpPort),
	}

	go runServer(srv)
	waitForShutdown(deps, srv)
}

func runServer(srv *http.Server) {
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalln(err.Error())
	}
}

func waitForShutdown(deps *app.Dependency, srv *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit,
		syscall.SIGINT,
		syscall.SIGTERM,
		os.Interrupt)

	<-quit

	log.Println("shutting down")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf(`error shutdown server. err: %v`, err)
	}

	deps.Destroy()
	log.Println("shutdown complete")
}
