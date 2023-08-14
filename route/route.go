package route

import (
	"clean-architecture/config"
	"clean-architecture/handler"
	"clean-architecture/middleware"
	"clean-architecture/repo"
	"clean-architecture/service"
)

type Service struct {
	*config.App
}

func NewService() *Service {
	s := Service{
		config.NewApp(),
	}

	db := s.GetDB()
	repo := repo.NewRepo(db)

	userService := service.NewUser(repo)
	user := handler.NewUser(userService)
	migrate := handler.NewMigration(db)
	health := handler.NewHealth()

	router := s.Router
	v1 := router.Group("/v1")

	// user
	v1.POST("/login", middleware.CheckAdmin(), user.GetUsers)
	router.POST("/check-health", health.CheckHealth)

	// migration
	router.POST("/migrate", middleware.CheckAdmin(), migrate.Migrate)
	return &s
}
