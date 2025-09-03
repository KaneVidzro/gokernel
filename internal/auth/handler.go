package auth

import (
	"context"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/kanevidzro/gokernel/internal/user"
	"github.com/redis/go-redis/v9"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

type AuthService struct {
    UserRepo  *user.Repository
    Redis     *redis.Client
    JWTSecret []byte
}

func (s *AuthService) RevokeToken(ctx context.Context, jti string, expiry time.Duration) error {
    return s.Redis.Set(ctx, "revoked:"+jti, "true", expiry).Err()
}

type AuthHandler struct {
    Service *AuthService
}

func (h *AuthHandler) Register(c *gin.Context) {
    var req struct {
        Name     string `json:"name"`
        Email    string `json:"email"`
        Password string `json:"password"`
    }
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    existingUser, err := h.Service.UserRepo.GetByEmail(req.Email)
    if err == nil && existingUser != nil {
        c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
        return
    }

    hashed, err := HashPassword(req.Password)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
        return
    }

    newUser := &user.User{
        Name:     req.Name,
        Email:    req.Email,
        Password: hashed,
        Role:     "user",
    }

    if err := h.Service.UserRepo.CreateUser(newUser); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func (h *AuthHandler) Login(c *gin.Context) {
    var req struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    userRecord, err := h.Service.UserRepo.GetByEmail(req.Email)
    if err != nil || userRecord == nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
        return
    }

    if !CheckPasswordHash(req.Password, userRecord.Password) {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    token, jti, err := GenerateToken(userRecord.ID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Token generation failed"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "token": token,
        "user": gin.H{
            "id":    userRecord.ID,
            "email": userRecord.Email,
            "role":  userRecord.Role,
        },
        "jti": jti,
    })
}

func (h *AuthHandler) Logout(c *gin.Context) {
    tokenStr := c.GetHeader("Authorization")
    if tokenStr == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Missing token"})
        return
    }

    tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")
    token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
        return h.Service.JWTSecret, nil
    })
    if err != nil || !token.Valid {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
        return
    }

    claims := token.Claims.(jwt.MapClaims)
    jti := claims["jti"].(string)
    exp := time.Until(time.Unix(int64(claims["exp"].(float64)), 0))

    if err := h.Service.RevokeToken(c.Request.Context(), jti, exp); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to revoke token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}
