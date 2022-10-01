// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"database/sql"
	"time"
)

type Absence struct {
	ID           int64         `json:"id"`
	UserID       int64         `json:"user_id"`
	Reason       string        `json:"reason"`
	Paid         bool          `json:"paid"`
	Date         time.Time     `json:"date"`
	CreatedAt    time.Time     `json:"created_at"`
	UpdatedAt    sql.NullTime  `json:"updated_at"`
	ApprovedByID sql.NullInt64 `json:"approved_by_id"`
	Length       float32       `json:"length"`
}

type Company struct {
	ID        int64        `json:"id"`
	Name      string       `json:"name"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}

type Entry struct {
	ID        int64        `json:"id"`
	UserID    int64        `json:"user_id"`
	StartTime time.Time    `json:"start_time"`
	EndTime   time.Time    `json:"end_time"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
	Date      time.Time    `json:"date"`
}

type Permission struct {
	ID        int64        `json:"id"`
	Name      string       `json:"name"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}

type Role struct {
	ID          int64        `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   sql.NullTime `json:"updated_at"`
}

type RolesPermission struct {
	RolesID       int64 `json:"roles_id"`
	PermissionsID int64 `json:"permissions_id"`
}

type Team struct {
	ID        int64         `json:"id"`
	Name      string        `json:"name"`
	ManagerID sql.NullInt64 `json:"manager_id"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt sql.NullTime  `json:"updated_at"`
}

type User struct {
	ID        int64         `json:"id"`
	Username  string        `json:"username"`
	Email     string        `json:"email"`
	Name      string        `json:"name"`
	Surname   string        `json:"surname"`
	CompanyID sql.NullInt64 `json:"company_id"`
	Password  string        `json:"password"`
	Gender    string        `json:"gender"`
	BirthDate time.Time     `json:"birth_date"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt sql.NullTime  `json:"updated_at"`
	// ISO-2 language code
	Language string `json:"language"`
	// ISO-2 Country code
	Country   string         `json:"country"`
	Timezone  sql.NullString `json:"timezone"`
	ManagerID sql.NullInt64  `json:"manager_id"`
	TeamID    sql.NullInt64  `json:"team_id"`
}

type UsersRole struct {
	UsersID int64 `json:"users_id"`
	RolesID int64 `json:"roles_id"`
}
