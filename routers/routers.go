package routers

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/austineyoyogie/go-hardware-store/api-jwt-auth/auth_controllers"
	"github.com/austineyoyogie/go-hardware-store/api-jwt-auth/auth_repository"
	"github.com/austineyoyogie/go-hardware-store/api-jwt-auth/auth_routes"

	"github.com/austineyoyogie/go-hardware-store/api-products/product_controllers"
	"github.com/austineyoyogie/go-hardware-store/api-products/product_repository"
	"github.com/austineyoyogie/go-hardware-store/api-products/product_routes"

	"github.com/austineyoyogie/go-hardware-store/database"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func RouterHandler() {
	l := log.New(os.Stdout, "[go-hardware-store-api-server]: ", log.LstdFlags)
	r := mux.NewRouter().StrictSlash(true)
	f := http.FileServer(http.Dir("./static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", f))

	db := database.Connect()
	if db != nil {
		defer db.Close()
	}

	permissionsRepository := auth_repository.NewPermissionsRepository(db)
	permissionsController := auth_controllers.NewPermissionsController(permissionsRepository)
	permissionRoutes := auth_routes.NewPermissionRoutes(permissionsController)

	usersRepository := auth_repository.NewUsersRepository(db)
	usersController := auth_controllers.NewUsersController(usersRepository)
	userRoutes := auth_routes.NewUserRoutes(usersController)

	loginRepository := auth_repository.UserLoginRepository(db)
	loginController := auth_controllers.UserLoginController(loginRepository)
	loginRoutes := auth_routes.UserLoginRoutes(loginController)

	auth_routes.AuthInstall(r, permissionRoutes, userRoutes, loginRoutes)

	categoriesRepository := product_repository.NewCategoriesRepository(db)
	categoriesController := product_controllers.NewCategoriesController(categoriesRepository)
	categoryRoutes := product_routes.NewCategoryRoutes(categoriesController)

	productsRepository := product_repository.NewProductsRepository(db)
	productsController := product_controllers.NewProductsController(productsRepository)
	productRoutes := product_routes.NewProductRoutes(productsController)

	product_routes.ProductInstall(r, categoryRoutes, productRoutes)

	s := &http.Server{
		Addr:         ":8000",
		Handler:      loadCors(r),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	go func() {
		l.Println("Starting [SERVER] on Port 8000")
		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Interrupt received, %s\n", err)
			os.Exit(1)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}

func loadCors(r http.Handler) http.Handler {
	headers := handlers.AllowedHeaders([]string{"Content-Type", "X-Requested-With", "Location", "Entity", "Accept", "Authorization"})
	methods := handlers.AllowedMethods([]string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions})
	origins := handlers.AllowedOrigins([]string{"*"})
	return handlers.CORS(headers, methods, origins)(r)
}
