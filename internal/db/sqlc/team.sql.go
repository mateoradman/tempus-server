// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: team.sql

package db

import (
	"context"
)

const createTeam = `-- name: CreateTeam :one
INSERT INTO teams (
    name,
    manager_id
) VALUES (
    $1, $2
)
RETURNING id, name, manager_id, created_at, updated_at
`

type CreateTeamParams struct {
	Name      string `json:"name"`
	ManagerID *int64 `json:"manager_id"`
}

func (q *Queries) CreateTeam(ctx context.Context, arg CreateTeamParams) (Team, error) {
	row := q.db.QueryRow(ctx, createTeam, arg.Name, arg.ManagerID)
	var i Team
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.ManagerID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteTeam = `-- name: DeleteTeam :one
DELETE
FROM teams
WHERE id = $1
RETURNING id, name, manager_id, created_at, updated_at
`

func (q *Queries) DeleteTeam(ctx context.Context, id int64) (Team, error) {
	row := q.db.QueryRow(ctx, deleteTeam, id)
	var i Team
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.ManagerID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getTeam = `-- name: GetTeam :one
SELECT id, name, manager_id, created_at, updated_at 
FROM teams
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetTeam(ctx context.Context, id int64) (Team, error) {
	row := q.db.QueryRow(ctx, getTeam, id)
	var i Team
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.ManagerID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listTeamMembers = `-- name: ListTeamMembers :many
SELECT id, username, email, name, surname, company_id, password, gender, birth_date, created_at, updated_at, language, country, timezone, manager_id, team_id
FROM users
WHERE team_id = $1
ORDER BY id
LIMIT $2
OFFSET $3
`

type ListTeamMembersParams struct {
	TeamID *int64 `json:"team_id"`
	Limit  int32  `json:"limit"`
	Offset int32  `json:"offset"`
}

func (q *Queries) ListTeamMembers(ctx context.Context, arg ListTeamMembersParams) ([]User, error) {
	rows, err := q.db.Query(ctx, listTeamMembers, arg.TeamID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.Email,
			&i.Name,
			&i.Surname,
			&i.CompanyID,
			&i.Password,
			&i.Gender,
			&i.BirthDate,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Language,
			&i.Country,
			&i.Timezone,
			&i.ManagerID,
			&i.TeamID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listTeams = `-- name: ListTeams :many
SELECT id, name, manager_id, created_at, updated_at
FROM teams
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListTeamsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListTeams(ctx context.Context, arg ListTeamsParams) ([]Team, error) {
	rows, err := q.db.Query(ctx, listTeams, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Team{}
	for rows.Next() {
		var i Team
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.ManagerID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateTeam = `-- name: UpdateTeam :one
UPDATE teams
SET name = $2, manager_id = $3
WHERE id = $1
RETURNING id, name, manager_id, created_at, updated_at
`

type UpdateTeamParams struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	ManagerID *int64 `json:"manager_id"`
}

func (q *Queries) UpdateTeam(ctx context.Context, arg UpdateTeamParams) (Team, error) {
	row := q.db.QueryRow(ctx, updateTeam, arg.ID, arg.Name, arg.ManagerID)
	var i Team
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.ManagerID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
