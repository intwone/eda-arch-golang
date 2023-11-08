package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	config "github.com/intwone/eda-arch-golang/internal/private_config"
	cryptography "github.com/intwone/eda-arch-golang/internal/private_cryptography/services"
	"github.com/intwone/eda-arch-golang/internal/private_database/gorm"
	"github.com/intwone/eda-arch-golang/internal/private_database/gorm/repositories"
	hasher "github.com/intwone/eda-arch-golang/internal/private_hasher/services"
	messengerHandlers "github.com/intwone/eda-arch-golang/internal/private_messenger/handlers"
	useCase "github.com/intwone/eda-arch-golang/internal/public_auth/application/use_cases"
	authDomainEvents "github.com/intwone/eda-arch-golang/internal/public_auth/events"
	controllers "github.com/intwone/eda-arch-golang/internal/public_auth/presentation/controllers"
	"github.com/intwone/eda-arch-golang/internal/public_auth/presentation/routes"
	contactHandlers "github.com/intwone/eda-arch-golang/internal/public_contact/handlers"
	passwordDomainEvents "github.com/intwone/eda-arch-golang/internal/public_password/events"
	userHandlers "github.com/intwone/eda-arch-golang/internal/public_user/handlers"
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

	// Repositories
	contactRepository := repositories.NewGORMContactRepository(db)
	passwordRepository := repositories.NewGORMPasswordRepository(db)
	userRepository := repositories.NewGORMUserRepository(db)
	bcryptHasher := hasher.NewBcryptHasher()
	jwtCryptography := cryptography.NewJWTCryptography(env.JWT_SECRET)

	// UseCases
	authCreateUseCase := useCase.NewAuthCreateUseCase(eventDispatcher, contactRepository, passwordRepository, bcryptHasher)
	authenticateUseCase := useCase.NewAuthenticateUseCase(eventDispatcher, contactRepository, passwordRepository, userRepository, jwtCryptography, bcryptHasher)

	// Events
	passwordCreatedEmailDispatchHandler := messengerHandlers.NewPasswordCreatedEmailDispatchHandler()
	eventDispatcher.Register(passwordDomainEvents.PasswordCreatedEventName, passwordCreatedEmailDispatchHandler)

	contactStatusVerifiedHandler := contactHandlers.NewContactVerifiedHandler(contactRepository)
	eventDispatcher.Register(authDomainEvents.AuthenticatedEventName, contactStatusVerifiedHandler)

	userStatusVerifiedHandler := userHandlers.NewUserStatusVerifiedHandler(userRepository)
	eventDispatcher.Register(authDomainEvents.AuthenticatedEventName, userStatusVerifiedHandler)

	// Controllers
	authCreateController := controllers.NewAuthCreateController(authCreateUseCase)
	authenticateController := controllers.NewAuthenticateController(authenticateUseCase)

	authControllers := controllers.AuthControllers{
		AuthCreateController:   authCreateController,
		AuthenticateController: authenticateController,
	}

	app := fiber.New()

	routes.SetupRoutes(app, authControllers)

	log.Fatal(app.Listen(":8000"))
}
