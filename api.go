package main

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type AuthState struct {
	sync.Mutex
	Password       string
	FailedAttempts int
	LockoutUntil   time.Time
}

type LoginRequest struct {
	Password string `json:"password"`
}

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func generatePassword(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func main() {
	// Generate random 8-character hex password
	password, err := generatePassword(4)
	if err != nil {
		log.Fatalf("Failed to generate password: %v", err)
	}

	state := &AuthState{
		Password: password,
	}

	fmt.Printf("\n====================================\n")
	fmt.Printf(" GENERATED PASSWORD: %s\n", state.Password)
	fmt.Printf("====================================\n\n")

	// Serve API endpoint
	http.HandleFunc("/api/login", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != http.MethodPost {
			http.Error(w, `{"success":false,"message":"Only POST allowed"}`, http.StatusMethodNotAllowed)
			return
		}

		var req LoginRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(Response{Success: false, Message: "Invalid JSON"})
			return
		}

		state.Lock()
		defer state.Unlock()

		now := time.Now()

		// Check if penalty box is active
		if now.Before(state.LockoutUntil) {
			remaining := time.Until(state.LockoutUntil).Truncate(time.Second)
			w.WriteHeader(http.StatusTooManyRequests)
			json.NewEncoder(w).Encode(Response{
				Success: false,
				Message: fmt.Sprintf("Locked out. Try again in %s", remaining),
			})
			return
		}

		// Verify password
		if req.Password == state.Password {
			state.FailedAttempts = 0
			json.NewEncoder(w).Encode(Response{Success: true, Message: "Authenticated successfully!"})
			return
		}

		// Handle failed attempt
		state.FailedAttempts++
		if state.FailedAttempts >= 3 {
			state.LockoutUntil = now.Add(1 * time.Hour)
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(Response{
				Success: false,
				Message: "Maximum attempts reached. Locked out for 1 hour.",
			})
		} else {
			remaining := 3 - state.FailedAttempts
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(Response{
				Success: false,
				Message: fmt.Sprintf("Incorrect password. %d attempts remaining.", remaining),
			})
		}
	})

	// Serve static files from './public' directory if it exists
	// fileServer := http.FileServer(http.Dir("./public"))

	index := func() {
		return ""
	}

	http.Handle("/", index)
 
	fmt.Println("Server running on http://localhost:8080")
	fmt.Println("Send a POST request to?")
	log.Fatal(http.ListenAndServe(":8080", nil))
}