package routes_test

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	. "github.com/franela/goblin"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"

	"github.com/raboliotlegris/Tulkas/core"
	"github.com/raboliotlegris/Tulkas/routes"
)

func Test_Routes(t *testing.T) {
	g := Goblin(t)
	g.Describe("Routes >", func() {
		g.Before(func() {
			log.SetLevel(log.DebugLevel)
		})

		g.Describe("Sessions >", func() {
			g.Describe("POST /api/users/login >", func() {
				g.It("Login works", func() {
					cfg := core.Config{
						UserName:         "root",
						UserHashPassword: "$2y$10$h9t6WyilHqWryBvNCmsIluF4tDP8Yw3dhtSQY5ZH1b0F.e5MzUvjG", // testpassword
					}
					store := core.NewSessionStore("tulkas")

					w := httptest.NewRecorder()
					r := routes.Create_router(&cfg, store)

					rawBuf, err := json.Marshal(routes.LoginUserBody{Username: "root", Password: "testpassword"})
					require.NoError(t, err)
					buffer := bytes.NewBuffer(rawBuf)

					req := httptest.NewRequest("POST", "/api/users/login", buffer)
					r.ServeHTTP(w, req)
					require.Equal(t, 200, w.Code)
				})
			})
			g.Describe("GET /api/users/check >", func() {
				g.It("Check without login doesn't works", func() {
					cfg := core.Config{}
					store := core.NewSessionStore("tulkas")

					w := httptest.NewRecorder()
					r := routes.Create_router(&cfg, store)

					req := httptest.NewRequest("GET", "/api/users/check", nil)
					r.ServeHTTP(w, req)
					require.Equal(t, 403, w.Code)
				})
				g.It("Check with login works", func() {
					cfg := core.Config{
						UserName:         "root",
						UserHashPassword: "$2y$10$h9t6WyilHqWryBvNCmsIluF4tDP8Yw3dhtSQY5ZH1b0F.e5MzUvjG", // testpassword
					}
					store := core.NewSessionStore("tulkas")

					w := httptest.NewRecorder()
					r := routes.Create_router(&cfg, store)

					rawBuf, err := json.Marshal(routes.LoginUserBody{Username: "root", Password: "testpassword"})
					require.NoError(t, err)
					buffer := bytes.NewBuffer(rawBuf)

					req := httptest.NewRequest("POST", "/api/users/login", buffer)
					r.ServeHTTP(w, req)
					require.Equal(t, 200, w.Code)

					req = httptest.NewRequest("GET", "/api/users/check", nil)
					r.ServeHTTP(w, req)
					require.Equal(t, 200, w.Code)
				})
			})
		})
	})
}
