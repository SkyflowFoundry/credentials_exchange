package main

import (
	"01-Login/middleware"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/joho/godotenv"
	saUtil "github.com/skyflowapi/skyflow-go/serviceaccount/util"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading the .env file: %v", err)
	}

	router := http.NewServeMux()

	//endpoint that validates auth0 bearer token and validates skyflow bearer token
	router.Handle("/api/generate-bearer-token", middleware.EnsureValidToken()(
		http.HandlerFunc(generateBearerToken),
	))

	log.Print("Server listening on http://localhost:8080")
	if err := http.ListenAndServe("0.0.0.0:8080", router); err != nil {
		log.Fatalf("There was an error with the http server: %v", err)
	}
}

func generateBearerToken(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Cannot do this method, must be POST", http.StatusMethodNotAllowed)
		return
	}

	token := r.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)
	if token == nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}
	//enter path to skyflow credentials file
	credentialsPath := "[path]"
	bearerToken, err := saUtil.GenerateBearerToken(credentialsPath)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error generating Skyflow bearer token: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": bearerToken.AccessToken})
}
