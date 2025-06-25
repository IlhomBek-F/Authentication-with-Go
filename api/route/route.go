package route

import (
	"auth/api/controller"
	"auth/model"
	"net/http"
	"os"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
)

type Server struct {
	model.Server
}

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(rate.Limit(10))))
	e.Use(configureCORS())

	server := &controller.Controller{
		Port: s.Server.Port,
		Db:   s.Server.Db,
	}

	secretKey := os.Getenv("ACCESS_TOKEN_SECRET")

	publicRoute := e.Group("/api")
	protectedRoute := publicRoute.Group("")

	publicRoute.POST("/signup", server.SignUp)
	publicRoute.POST("/login", server.Login)

	protectedRoute.Use(echojwt.JWT([]byte(secretKey)))
	protectedRoute.GET("/users", server.GetUsers)

	return e
}

func configureCORS() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"https://*", "http://*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           30,
	})
}
