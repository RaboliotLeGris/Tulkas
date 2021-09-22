package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/sessions"
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
	SessionStore *sessions.CookieStore
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
		session, err := h.SessionStore.Get(r, "tulkas")
		if err != nil {
			log.Error("/api/login: ", err)
			w.WriteHeader(401)
			return
		}
		session.Values["username"] = body.Username
		session.Save(r, w)
		w.WriteHeader(200)
		return
	}
	w.WriteHeader(401)
}

type GetCheckSessionHandler struct {
	SessionStore *sessions.CookieStore
}

func (h GetCheckSessionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Debug("Request GET /api/users/check")
	session, err := h.SessionStore.Get(r, "tulkas")
	if err != nil {
		log.Debug("Get Session err value: ", err)
		w.WriteHeader(403)
		return
	}
	_, ok := session.Values["username"]
	if ok && !session.IsNew {
		w.WriteHeader(200)
		return
	}
	w.WriteHeader(403)
}
