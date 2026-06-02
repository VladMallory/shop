package app

import (
	"authTest/internal/config"
	"authTest/internal/handler"
	"authTest/internal/infrastructure/db"
	"authTest/internal/middleware"
	"authTest/internal/repository/postgres"
	"authTest/internal/service"
	"net/http"

	_ "github.com/lib/pq" // Драйвер для Postgres
)

type dbRun interface {
	Start()
}

type App struct {
	db         dbRun
	httpServer *http.Server
}

// func close(body io.Closer, inputErr *error) {
// 	if err := body.Close(); err != nil {
// 		*inputErr = err
// 	}
// }

func New() (*App, error) {
	cfg := config.Config()

	// === DB ===
	dbConn, err := db.InitDB("postgres://myuser:mypassword@localhost:5432/mydb?sslmode=disable")
	if err != nil {
		return nil, err
	}

	postgresRepo := postgres.NewPostgres(dbConn)

	serviceDB := service.NewUserService(postgresRepo)

	authService := service.NewAuthService(postgresRepo, cfg.JWTToken)

	// === HANDLER ===
	authHandler := handler.NewAuthHandler(authService)
	profileHandler := handler.NewProfileHandler(serviceDB)

	// === HTTP ===
	mux := http.NewServeMux()

	// Открытые эндпоинты — без middleware
	mux.HandleFunc("POST /auth/register", authHandler.Register)
	mux.HandleFunc("POST /auth/login", authHandler.Login)

	// Защищённые эндпоинты, обёрнуты в Auth middleware.
	mux.Handle(
		"GET /profile",
		middleware.Auth(authService)(http.HandlerFunc(profileHandler.GetProfile)))

	server := &http.Server{
		Addr:    ":8083",
		Handler: mux,
	}

	return &App{
		db:         serviceDB,
		httpServer: server,
	}, nil
}

func (a *App) Run() error {
	// === DB ===
	a.db.Start()

	a.httpServer.ListenAndServe()

	return nil
}
