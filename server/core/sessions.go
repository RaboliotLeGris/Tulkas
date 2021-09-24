package core

import (
	"net/http"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	log "github.com/sirupsen/logrus"
)

type SessionStore struct {
	store     *sessions.CookieStore
	storeName string
}

func NewSessionStore(storeName string) *SessionStore {
	store := sessions.NewCookieStore(securecookie.GenerateRandomKey(32))
	return &SessionStore{store, storeName}
}

func (s *SessionStore) IsAuth(r *http.Request) bool {
	session, err := s.store.Get(r, "nienna")
	if err != nil {
		log.Debug("Get Session err value: ", err)
		return false
	}
	logged, ok1 := session.Values["logged"]
	_, ok2 := session.Values["username"]
	return ok1 && logged.(bool) && ok2 && !session.IsNew

}

func (s *SessionStore) Get(r *http.Request, key string) interface{} {
	session, err := s.store.Get(r, s.storeName)
	if err != nil {
		return nil
	}
	return session.Values[key]
}

func (s *SessionStore) Set(r *http.Request, w http.ResponseWriter, key string, value interface{}) error {
	session, err := s.store.Get(r, "nienna")
	if err != nil {
		return err
	}
	session.Values[key] = value
	return session.Save(r, w)
}
