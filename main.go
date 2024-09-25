package main

import (
	"context"
	"html/template"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"packing-stacking/internal/handlers"

	"github.com/brpaz/echozap"
	"github.com/labstack/echo/v4"
	swag "github.com/swaggo/echo-swagger"
	"go.uber.org/zap"
	_ "packing-stacking/docs" // docs is generated by Swag CLI, you have to import it.
)

//	@title			Packing Stacking API
//	@version		1.0
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

func main() {
	e := echo.New()

	// viper.AutomaticEnv()
	//viper.SetConfigType("env")
	//viper.AddConfigPath(".env")
	//err := viper.ReadInConfig()
	//if err != nil {
	//	return
	//}

	// Logger
	logger, _ := zap.NewProduction()
	e.Use(echozap.ZapLogger(logger))

	t := &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}
	e.Renderer = t
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", nil)
	})

	e.GET("/swagger/*", swag.WrapHandler)
	e.GET("/healthz", func(c echo.Context) error {
		return c.String(http.StatusOK, "i'm up!")
	})

	handler := handlers.NewPacksHandler()

	// Packs router group
	users := e.Group("/packs")
	users.Add(echo.GET, "/add", handler.Add)
	users.Add(echo.GET, "/list", handler.List)
	users.Add(echo.GET, "/remove/:quantity", handler.Remove)
	users.Add(echo.GET, "/calculate/:quantity", handler.Calculate)

	go func() {
		if serverErr := e.Start(":8080"); serverErr != nil && serverErr != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server with a timeout of 10 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
