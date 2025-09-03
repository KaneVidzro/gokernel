package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerV1 struct {
    Repository *Repository
}

func (h *HandlerV1) GetUser(c *gin.Context) {
    id := c.Param("id")

    user, err := h.Repository.GetByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "id":         user.ID,
        "email":      user.Email,
        "role":       user.Role,
        "created_at": user.CreatedAt,
    })
}
