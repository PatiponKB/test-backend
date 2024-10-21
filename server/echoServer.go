package server

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/PatiponKB/backend-test/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type echoServer struct {
	app  *echo.Echo
	db 	 *gorm.DB
	conf *config.Config
}

var (
	server *echoServer
	once    sync.Once
)

func NewechoServer(conf *config.Config, db *gorm.DB) *echoServer{
		echoApp := echo.New()
		echoApp.Logger.SetLevel(log.DEBUG)

		once.Do(func() {
			server = &echoServer{
				app : echoApp,
				db: db,
				conf: conf,
			}
		})
		return server
	}

func (s *echoServer) Start() {

	s.app.GET("/v1/health", s.healthcheck)
	s.initBeerRouter()

	s.httpListening()
}

func (s *echoServer) httpListening() {
	url := fmt.Sprintf(":%d", s.conf.Server.Port)

	if err := s.app.Start(url);err != nil && err != http.ErrServerClosed{
		s.app.Logger.Fatalf("Error: %s", err.Error())
	}
}

func (s *echoServer) healthcheck(c echo.Context) error{
	return c.String(http.StatusOK,"Status OK")
}

