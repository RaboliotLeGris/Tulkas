package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"github.com/raboliotlegris/Tulkas/core"
)

func Create_router(cfg *core.Config, sessionStore *core.SessionStore) *mux.Router {
	log.Info("Creating routers")
	// Routes order creation matter.
	r := mux.NewRouter()

	r.PathPrefix("/api/users/login").Handler(PostLoginHandler{Cfg: cfg, SessionStore: sessionStore}).Methods("POST")
	r.PathPrefix("/api/users/check").Handler(GetCheckSessionHandler{SessionStore: sessionStore}).Methods("GET")

	r.PathPrefix("/api/live").Handler(GetLiveImage{}).Methods("GET")
	r.PathPrefix("/api/live").Handler(PostLiveImage{}).Methods("POST")

	r.PathPrefix("/api/toggle/light").Handler(PostToggleLight{}).Methods("POST")
	r.PathPrefix("/api/save").Handler(PostSaveImages{}).Methods("POST")

	r.PathPrefix("/").Handler(StaticHandler{StaticPath: "static", IndexPath: "index.html"})

	return r
}

func Launch(cfg *core.Config, router *mux.Router) error {
	log.Info("Launching HTTP server")

	srv := &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf("0.0.0.0:%v", cfg.Port_HTTP),
	}

	return srv.ListenAndServe()
}
