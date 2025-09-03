package user

import "database/sql"

type Repository struct {
    DB *sql.DB
}


func (r *Repository) GetByEmail(email string) (*User, error) {
    var u User
    err := r.DB.QueryRow(`
        SELECT id, email, password, role, created_at
        FROM users WHERE email = $1
    `, email).Scan(&u.ID, &u.Email, &u.Password, &u.Role, &u.CreatedAt)
    if err != nil {
        return nil, err
    }
    return &u, nil
}


func (r *Repository) ExistsByID(id string) (bool, error) {
    var exists bool
    err := r.DB.QueryRow(`
        SELECT EXISTS(SELECT 1 FROM users WHERE id = $1)
    `, id).Scan(&exists)
    return exists, err
}

func (r *Repository) GetByID(id string) (*User, error) {
    var u User
    err := r.DB.QueryRow(`
        SELECT id, email, password, role, created_at
        FROM users WHERE id = $1
    `, id).Scan(&u.ID, &u.Email, &u.Password, &u.Role, &u.CreatedAt)
    if err != nil {
        return nil, err
    }
    return &u, nil
}



func (r *Repository) CreateUser(u *User) error {
    _, err := r.DB.Exec(`
        INSERT INTO users (email, password_hash, role)
        VALUES ($1, $2, $3)
    `, u.Email, u.Password, u.Role)
    return err
}

func (r *Repository) GetAll() ([]*User, error) {
    rows, err := r.DB.Query("SELECT id, email, role FROM users")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var users []*User
    for rows.Next() {
        var u User
        if err := rows.Scan(&u.ID, &u.Email, &u.Role); err != nil {
            return nil, err
        }
        users = append(users, &u)
    }
    return users, nil
}



func (r *Repository) SetRole(userID, role string) error {
    _, err := r.DB.Exec("UPDATE users SET role = ? WHERE id = ?", role, userID)
    return err
}

func (r *Repository) SetActive(userID string, active bool) error {
    _, err := r.DB.Exec("UPDATE users SET active = ? WHERE id = ?", active, userID)
    return err
}

func (r *Repository) CountUsers() (int, error) {
    var count int
    err := r.DB.QueryRow("SELECT COUNT(*) FROM users").Scan(&count)
    return count, err
}

func (r *Repository) CountByRole(role string) (int, error) {
    var count int
    err := r.DB.QueryRow("SELECT COUNT(*) FROM users WHERE role = ?", role).Scan(&count)
    return count, err
}

