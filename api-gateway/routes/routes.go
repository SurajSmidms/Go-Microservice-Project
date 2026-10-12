package routes

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	gwMiddleware "github.com/SurajSmidms/GOLANG_MICROSERVICE_PROJECT/api-gateway/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func SetupRouter() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Fetch service URLs from environment variables (Docker container names)
	authURL := getEnv("AUTH_SERVICE_URL", "http://localhost:8081")
	productURL := getEnv("PRODUCT_SERVICE_URL", "http://localhost:8082")
	orderURL := getEnv("ORDER_SERVICE_URL", "http://localhost:8083")
	paymentURL := getEnv("PAYMENT_SERVICE_URL", "http://localhost:8084")

	// Public Routes (Auth Service proxy)
	r.Post("/api/auth/register", proxyRequest(authURL))
	r.Post("/api/auth/login", proxyRequest(authURL))

	// Protected Routes Group (Requires Auth Middleware)
	r.Group(func(r chi.Router) {
		r.Use(gwMiddleware.AuthMiddleware)

		r.HandleFunc("/api/products*", proxyRequest(productURL))
		r.HandleFunc("/api/orders*", proxyRequest(orderURL))
		r.HandleFunc("/api/payments*", proxyRequest(paymentURL))
	})

	return r
}

func proxyRequest(target string) http.HandlerFunc {
	url, _ := url.Parse(target)
	proxy := httputil.NewSingleHostReverseProxy(url)
	return func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	}
}