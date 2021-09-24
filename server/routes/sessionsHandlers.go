package routes

import (
	"encoding/json"
	"net/http"

	"github.com/raboliotlegris/Tulkas/core"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type LoginUserBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type PostLoginHandler struct {
	Cfg          *core.Config
	SessionStore *core.SessionStore
}

func (h PostLoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Debug("Request POST /api/users/login")
	var body LoginUserBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.Error("Fail to deserialize user login struct")
		w.WriteHeader(400)
		return
	}

	if body.Username == "" || body.Password == "" {
		log.Error("Empty username or password")
		w.WriteHeader(400)
		return
	}

	if body.Username == h.Cfg.UserName && bcrypt.CompareHashAndPassword([]byte(h.Cfg.UserHashPassword), []byte(body.Password)) == nil {
		if err := h.SessionStore.Set(r, w, "username", body.Username); err != nil {
			w.WriteHeader(403)
			return
		}
		if err := h.SessionStore.Set(r, w, "logged", true); err != nil {
			w.WriteHeader(403)
			return
		}
		w.WriteHeader(200)
		return
	}
	w.WriteHeader(401)
}

type GetCheckSessionHandler struct {
	SessionStore *core.SessionStore
}

func (h GetCheckSessionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Debug("Request GET /api/users/check")
	if h.SessionStore.IsAuth(r) {
		w.WriteHeader(200)
		return
	}
	w.WriteHeader(403)
}
