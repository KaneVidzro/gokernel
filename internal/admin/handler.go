package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
    Service *Service
}

func (h *Handler) Dashboard(c *gin.Context) {
    stats := h.Service.GetSystemStats()
    c.JSON(http.StatusOK, gin.H{"stats": stats})
}

func (h *Handler) ListUsers(c *gin.Context) {
    users, err := h.Service.ListAllUsers()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
        return
    }
    c.JSON(http.StatusOK, users)
}

func (h *Handler) UpdateRole(c *gin.Context) {
    id := c.Param("id")
    var req struct {
        Role string `json:"role"`
    }
    if err := c.ShouldBindJSON(&req); err != nil || req.Role == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role input"})
        return
    }

    if err := h.Service.UpdateUserRole(id, req.Role); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update role"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Role updated"})
}

func (h *Handler) DeactivateUser(c *gin.Context) {
    id := c.Param("id")
    if err := h.Service.DeactivateUser(id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to deactivate user"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "User deactivated"})
}
