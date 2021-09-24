package routes

import "net/http"

type PostToggleLight struct {
}

func (l PostToggleLight) ServeHTTP(w http.ResponseWriter, r *http.Request) {}
