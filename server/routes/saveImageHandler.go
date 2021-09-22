package routes

import "net/http"

type PostSaveImages struct {
}

func (l PostSaveImages) ServeHTTP(w http.ResponseWriter, r *http.Request) {}
