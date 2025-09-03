package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/kanevidzro/gokernel/internal/user"
	"golang.org/x/crypto/bcrypt"
)


func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

func GenerateToken(u *user.User) (string, string, error) {
    jti := uuid.NewString() // unique token ID

 claims := jwt.MapClaims{
    "user_id": u.ID,
    "role":    u.Role,
    "jti":     jti,
    "exp":     time.Now().Add(time.Hour * 72).Unix(),
}



    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    signed, err := token.SignedString(jwtSecret)
    return signed, jti, err
}