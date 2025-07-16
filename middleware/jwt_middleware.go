package middleware

import (
    "context"
    "diary-app/utils"
    "net/http"
    "strings"
)

type contextKey string

const userKey contextKey = "username"

// AuthMiddleware validates the JWT and adds username to the request context
func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            http.Error(w, "Authorization header missing", http.StatusUnauthorized)
            return
        }

        parts := strings.Split(authHeader, " ")
        if len(parts) != 2 || parts[0] != "Bearer" {
            http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
            return
        }

        token := parts[1]
        username, err := utils.ValidateJWT(token)
        if err != nil {
            http.Error(w, "Invalid token: "+err.Error(), http.StatusUnauthorized)
            return
        }

        // Add username to context
        ctx := context.WithValue(r.Context(), userKey, username)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}

// GetUsernameFromContext retrieves the username from request context
func GetUsernameFromContext(r *http.Request) string {
    if val := r.Context().Value(userKey); val != nil {
        return val.(string)
    }
    return ""
}
