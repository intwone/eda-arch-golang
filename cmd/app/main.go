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
	accountHandlers "github.com/intwone/eda-arch-golang/internal/public_account/handlers"
	authUseCase "github.com/intwone/eda-arch-golang/internal/public_auth/application/use_cases"
	authDomainEvents "github.com/intwone/eda-arch-golang/internal/public_auth/events"
	authControllers "github.com/intwone/eda-arch-golang/internal/public_auth/presentation/controllers"
	authRoute "github.com/intwone/eda-arch-golang/internal/public_auth/presentation/routes"
	contactHandlers "github.com/intwone/eda-arch-golang/internal/public_contact/handlers"
	passwordDomainEvents "github.com/intwone/eda-arch-golang/internal/public_password/events"
	permissionHandlers "github.com/intwone/eda-arch-golang/internal/public_permission/handlers"
	personHandlers "github.com/intwone/eda-arch-golang/internal/public_person/handlers"
	userUseCase "github.com/intwone/eda-arch-golang/internal/public_user/application/use_cases"
	userDomainEvents "github.com/intwone/eda-arch-golang/internal/public_user/events"
	userHandlers "github.com/intwone/eda-arch-golang/internal/public_user/handlers"
	userControllers "github.com/intwone/eda-arch-golang/internal/public_user/presentation/controllers"
	userRoute "github.com/intwone/eda-arch-golang/internal/public_user/presentation/routes"
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
	permissionRepository := repositories.NewGORMPermissionRepository(db)
	personRepository := repositories.NewGORMPersonRepository(db)
	accountRepository := repositories.NewGORMAccountRepository(db)
	jwtCryptography := cryptography.NewJWTCryptography(env.JWT_SECRET)
	bcryptHasher := hasher.NewBcryptHasher()

	// UseCases
	authCreateUseCase := authUseCase.NewAuthCreateUseCase(eventDispatcher, contactRepository, passwordRepository, bcryptHasher)
	authenticateUseCase := authUseCase.NewAuthenticateUseCase(eventDispatcher, contactRepository, passwordRepository, userRepository, jwtCryptography, bcryptHasher)
	userCreateUseCase := userUseCase.NewCreateUserUseCase(eventDispatcher, userRepository, contactRepository, personRepository)

	// Events
	passwordCreatedEmailDispatchHandler := messengerHandlers.NewPasswordCreatedEmailDispatchHandler()
	eventDispatcher.Register(passwordDomainEvents.PasswordCreatedEventName, passwordCreatedEmailDispatchHandler)

	contactStatusVerifiedHandler := contactHandlers.NewContactVerifiedHandler(contactRepository)
	eventDispatcher.Register(authDomainEvents.AuthenticatedEventName, contactStatusVerifiedHandler)

	userStatusVerifiedHandler := userHandlers.NewUserStatusVerifiedHandler(userRepository)
	eventDispatcher.Register(authDomainEvents.AuthenticatedEventName, userStatusVerifiedHandler)

	personCreatedHandler := personHandlers.NewPersonCreateHandler(personRepository)
	eventDispatcher.Register(userDomainEvents.UserCreatedEventName, personCreatedHandler)

	permissionCreateHandler := permissionHandlers.NewPermissionCreateHandler(permissionRepository)
	eventDispatcher.Register(userDomainEvents.UserCreatedEventName, permissionCreateHandler)

	accountCreateHandler := accountHandlers.NewAccountCreateHandler(accountRepository)
	eventDispatcher.Register(userDomainEvents.UserCreatedEventName, accountCreateHandler)

	// Controllers
	authCreateController := authControllers.NewAuthCreateController(authCreateUseCase)
	authenticateController := authControllers.NewAuthenticateController(authenticateUseCase)
	userCreateController := userControllers.NewUserCreateController(userCreateUseCase)

	authControllers := authControllers.AuthControllers{
		AuthCreateController:   authCreateController,
		AuthenticateController: authenticateController,
	}

	userControllers := userControllers.UserControllers{
		UserCreateController: userCreateController,
	}

	app := fiber.New()

	authRoute.SetupAuthRoutes(app, authControllers)
	userRoute.SetupUserRoutes(app, userControllers)

	log.Fatal(app.Listen(":8000"))
}
