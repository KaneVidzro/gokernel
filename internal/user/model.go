package user

import "time"

type User struct {
    ID        string    `db:"id"`
    Name      string    `db:"name"`
    Email     string    `db:"email"`
    Password  string    `db:"password"` // This stores the hashed password
    Role      string    `db:"role"`
    CreatedAt time.Time `db:"created_at"`
}
