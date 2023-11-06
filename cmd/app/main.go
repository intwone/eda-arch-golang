package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	useCase "github.com/intwone/eda-arch-golang/internal/domain/modules/auth/use_cases"
	messenger "github.com/intwone/eda-arch-golang/internal/domain/modules/messenger/handlers"
	domainEvents "github.com/intwone/eda-arch-golang/internal/domain/modules/password/events"
	"github.com/intwone/eda-arch-golang/internal/infra/database/gorm"
	"github.com/intwone/eda-arch-golang/internal/infra/database/gorm/repositories"
	"github.com/intwone/eda-arch-golang/internal/infra/hasher"
	"github.com/intwone/eda-arch-golang/internal/main/config"
	"github.com/intwone/eda-arch-golang/internal/main/routes"
	controllers "github.com/intwone/eda-arch-golang/internal/presentation"
	authControllers "github.com/intwone/eda-arch-golang/internal/presentation/auth/controllers"
	"github.com/intwone/eda-arch-golang/pkg/events"
)

func main() {
	env := config.Env()

	config := &gorm.Config{
		Host:    env.DATABASE_HOST,
		Port:    env.DATABASE_PORT,
		User:    env.DATABASE_USER,
		Pass:    env.DATABASE_PASSWORD,
		Name:    env.DATABASE_NAME,
		SSLMode: env.DATABASE_SSL_MODE,
	}

	db, err := gorm.NewConn(config)

	if err != nil {
		log.Fatal(err)
	}

	gorm.Migrate(db)

	eventDispatcher := events.NewEventDispatcher()

	passwordCreatedEmailDispatchHandler := messenger.NewPasswordCreatedEmailDispatchHandler()
	eventDispatcher.Register(domainEvents.PasswordCreatedEventName, passwordCreatedEmailDispatchHandler)

	contactRepository := repositories.NewGORMContactRepository(db)
	passwordRepository := repositories.NewGORMPasswordRepository(db)
	bcryptHasher := hasher.NewBcryptHasher()

	authCreateUseCase := useCase.NewAuthCreateUseCase(eventDispatcher, contactRepository, passwordRepository, bcryptHasher)
	authCreateController := authControllers.NewAuthCreateController(authCreateUseCase)

	authControllers := controllers.AuthControllers{
		AuthCreateController: authCreateController,
	}

	app := fiber.New()

	routes.SetupRoutes(app, authControllers)

	log.Fatal(app.Listen(":3000"))
}
