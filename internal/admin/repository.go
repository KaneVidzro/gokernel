package admin

import (
	"github.com/kanevidzro/gokernel/internal/user"
)

type Repository struct {
    UserRepo *user.Repository
}

// GetAllUsers returns all users.
func (r *Repository) GetAllUsers() ([]*user.User, error) {
    return r.UserRepo.GetAll()
}

// UpdateUserRole sets a user's role.
func (r *Repository) UpdateUserRole(userID, role string) error {
    return r.UserRepo.SetRole(userID, role)
}

// DeactivateUser disables a user account.
func (r *Repository) DeactivateUser(userID string) error {
    return r.UserRepo.SetActive(userID, false)
}
