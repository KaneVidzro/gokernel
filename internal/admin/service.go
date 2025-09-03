package admin

import (
	"github.com/kanevidzro/gokernel/internal/user"
)

type Service struct {
    UserRepo *user.Repository
}

// GetSystemStats returns basic system metrics for the admin dashboard.
func (s *Service) GetSystemStats() map[string]interface{} {
    totalUsers, _ := s.UserRepo.CountUsers()
    activeAdmins, _ := s.UserRepo.CountByRole("admin")

    return map[string]interface{}{
        "total_users":   totalUsers,
        "active_admins": activeAdmins,
    }
}

// ListAllUsers returns all users in the system.
func (s *Service) ListAllUsers() ([]*user.User, error) {
    return s.UserRepo.GetAll()
}

// UpdateUserRole changes the role of a user.
func (s *Service) UpdateUserRole(userID, role string) error {
    return s.UserRepo.SetRole(userID, role)
}

// DeactivateUser disables a user account.
func (s *Service) DeactivateUser(userID string) error {
    return s.UserRepo.SetActive(userID, false)
}
