package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/instinctG/Rest-API/internal/comment"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

// Handler - stores pointer to our comment service
type Handler struct {
	Router  *mux.Router
	Service *coment.Service
}

type Response struct {
	Message string
	Error   string
}

// NewHandler - returns a pointer to a Handler
func NewHandler(service *coment.Service) *Handler {
	return &Handler{Service: service}
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(log.Fields{
			"Method": r.Method,
			"Path":   r.URL.Path,
		}).Info("handled request!")
		next.ServeHTTP(w, r)
	})
}

// BasicAuth - a handy middleware function that logs out incoming requests
func BasicAuth(origin func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if user == "admin" && pass == "secret" && ok {
			origin(w, r)
		} else {
			sendErrorResponse(w, "not authorized", errors.New("not authorized"))
		}

	}
}

// ValidateToken - validates an incoming jwt token
func validateToken(accessToken string) bool {
	var mySigningKey = []byte("restapi")
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}
		return mySigningKey, nil
	})

	if err != nil {
		return false
	}

	return token.Valid
}

// JWTAuth - a decorator function for authentication
func JWTAuth(origin func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info("jwt auth endpoint hit")
		authHeader := r.Header["Authorization"]
		if authHeader == nil {
			sendErrorResponse(w, "not authorized", errors.New("not authorized"))
		}

		authHeaderParts := strings.Split(authHeader[0], " ")
		if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
			sendErrorResponse(w, "not authorized", errors.New("not authorized"))
		}

		if validateToken(authHeaderParts[1]) {
			origin(w, r)
		} else {
			sendErrorResponse(w, "not authorized", errors.New("not authorized"))
		}

	}
}

// SetupRoutes - sets up all the routes for our application
func (h *Handler) SetupRoutes() {
	log.Info("Setting Up Routes")
	h.Router = mux.NewRouter()
	h.Router.Use(LoggingMiddleware)

	h.Router.HandleFunc("/api/comment/{id}", h.GetComment).Methods("GET")
	h.Router.HandleFunc("/api/comments", h.GetAllComments).Methods("GET")
	h.Router.HandleFunc("/api/comment", JWTAuth(h.PostComment)).Methods("POST")
	h.Router.HandleFunc("/api/comment/{id}", JWTAuth(h.UpdateComment)).Methods("PUT")
	h.Router.HandleFunc("/api/comment/{id}", JWTAuth(h.DeleteComment)).Methods("DELETE")
	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(Response{Message: "I Am Alive"}); err != nil {
			panic(err)
		}
	})
}

func sendErrorResponse(w http.ResponseWriter, message string, err error) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.NewEncoder(w).Encode(Response{
		Message: message,
		Error:   err.Error(),
	}); err != nil {
		panic(err)
	}
}

func sendOkResponse(w http.ResponseWriter, resp interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	return json.NewEncoder(w).Encode(resp)
}
