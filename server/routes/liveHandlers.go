package routes

import "net/http"

type GetLiveImage struct{}

func (l GetLiveImage) ServeHTTP(w http.ResponseWriter, r *http.Request) {}

type PostLiveImage struct{}

func (l PostLiveImage) ServeHTTP(w http.ResponseWriter, r *http.Request) {}
