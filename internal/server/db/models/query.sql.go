// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: query.sql

package db

import (
	"context"
	"time"
)

const createGlobalSettings = `-- name: CreateGlobalSettings :one
INSERT INTO
    global_settings (
        smtp_enabled,
        add_member_email_subject,
        add_member_email_template
    )
VALUES
    (?, ?, ?) RETURNING id, smtp_enabled, smtp_host, smtp_port, smtp_username, smtp_password, from_address, add_member_email_subject, add_member_email_template
`

type CreateGlobalSettingsParams struct {
	SmtpEnabled            bool
	AddMemberEmailSubject  interface{}
	AddMemberEmailTemplate interface{}
}

func (q *Queries) CreateGlobalSettings(ctx context.Context, arg CreateGlobalSettingsParams) (GlobalSetting, error) {
	row := q.db.QueryRowContext(ctx, createGlobalSettings, arg.SmtpEnabled, arg.AddMemberEmailSubject, arg.AddMemberEmailTemplate)
	var i GlobalSetting
	err := row.Scan(
		&i.ID,
		&i.SmtpEnabled,
		&i.SmtpHost,
		&i.SmtpPort,
		&i.SmtpUsername,
		&i.SmtpPassword,
		&i.FromAddress,
		&i.AddMemberEmailSubject,
		&i.AddMemberEmailTemplate,
	)
	return i, err
}

const createNewConnection = `-- name: CreateNewConnection :one
INSERT INTO
    connections (subdomain, team_member_id)
VALUES
    (?, ?) RETURNING id, subdomain, team_member_id, created_at, closed_at
`

type CreateNewConnectionParams struct {
	Subdomain    string
	TeamMemberID int64
}

func (q *Queries) CreateNewConnection(ctx context.Context, arg CreateNewConnectionParams) (Connection, error) {
	row := q.db.QueryRowContext(ctx, createNewConnection, arg.Subdomain, arg.TeamMemberID)
	var i Connection
	err := row.Scan(
		&i.ID,
		&i.Subdomain,
		&i.TeamMemberID,
		&i.CreatedAt,
		&i.ClosedAt,
	)
	return i, err
}

const createSession = `-- name: CreateSession :one
INSERT INTO
    sessions (token, user_id)
VALUES
    (?, ?) RETURNING id, user_id, token, created_at
`

type CreateSessionParams struct {
	Token  string
	UserID int64
}

func (q *Queries) CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error) {
	row := q.db.QueryRowContext(ctx, createSession, arg.Token, arg.UserID)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Token,
		&i.CreatedAt,
	)
	return i, err
}

const createTeam = `-- name: CreateTeam :one
INSERT INTO
    teams (name, slug)
VALUES
    (?, ?) RETURNING id, name, slug, created_at
`

type CreateTeamParams struct {
	Name string
	Slug string
}

func (q *Queries) CreateTeam(ctx context.Context, arg CreateTeamParams) (Team, error) {
	row := q.db.QueryRowContext(ctx, createTeam, arg.Name, arg.Slug)
	var i Team
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Slug,
		&i.CreatedAt,
	)
	return i, err
}

const createTeamMember = `-- name: CreateTeamMember :one
INSERT INTO
    team_members (user_id, team_id, role, secret_key)
VALUES
    (?, ?, ?, ?) RETURNING id, user_id, team_id, secret_key, role, added_by_user_id, created_at
`

type CreateTeamMemberParams struct {
	UserID    int64
	TeamID    int64
	Role      string
	SecretKey string
}

func (q *Queries) CreateTeamMember(ctx context.Context, arg CreateTeamMemberParams) (TeamMember, error) {
	row := q.db.QueryRowContext(ctx, createTeamMember,
		arg.UserID,
		arg.TeamID,
		arg.Role,
		arg.SecretKey,
	)
	var i TeamMember
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.TeamID,
		&i.SecretKey,
		&i.Role,
		&i.AddedByUserID,
		&i.CreatedAt,
	)
	return i, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO
    users (
        email,
        first_name,
        last_name,
        is_super_user,
        github_access_token,
        github_avatar_url
    )
VALUES
    (?, ?, ?, ?, ?, ?) RETURNING id, email, first_name, last_name, is_super_user, github_access_token, github_avatar_url, created_at
`

type CreateUserParams struct {
	Email             string
	FirstName         interface{}
	LastName          interface{}
	IsSuperUser       bool
	GithubAccessToken interface{}
	GithubAvatarUrl   interface{}
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Email,
		arg.FirstName,
		arg.LastName,
		arg.IsSuperUser,
		arg.GithubAccessToken,
		arg.GithubAvatarUrl,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.FirstName,
		&i.LastName,
		&i.IsSuperUser,
		&i.GithubAccessToken,
		&i.GithubAvatarUrl,
		&i.CreatedAt,
	)
	return i, err
}

const deleteSession = `-- name: DeleteSession :exec
DELETE FROM sessions
WHERE
    token = ?
`

func (q *Queries) DeleteSession(ctx context.Context, token string) error {
	_, err := q.db.ExecContext(ctx, deleteSession, token)
	return err
}

const getActiveConnectionsForTeam = `-- name: GetActiveConnectionsForTeam :many
SELECT
    connections.id,
    connections.subdomain,
    connections.created_at,
    connections.closed_at,
    users.email,
    users.first_name,
    users.last_name,
    users.github_avatar_url
FROM
    connections
    JOIN team_members ON team_members.id = connections.team_member_id
    JOIN users ON users.id = team_members.user_id
WHERE
    team_id = ?
    AND closed_at IS NULL
ORDER BY
    connections.id DESC
LIMIT
    20
`

type GetActiveConnectionsForTeamRow struct {
	ID              int64
	Subdomain       string
	CreatedAt       time.Time
	ClosedAt        interface{}
	Email           string
	FirstName       interface{}
	LastName        interface{}
	GithubAvatarUrl interface{}
}

func (q *Queries) GetActiveConnectionsForTeam(ctx context.Context, teamID int64) ([]GetActiveConnectionsForTeamRow, error) {
	rows, err := q.db.QueryContext(ctx, getActiveConnectionsForTeam, teamID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetActiveConnectionsForTeamRow
	for rows.Next() {
		var i GetActiveConnectionsForTeamRow
		if err := rows.Scan(
			&i.ID,
			&i.Subdomain,
			&i.CreatedAt,
			&i.ClosedAt,
			&i.Email,
			&i.FirstName,
			&i.LastName,
			&i.GithubAvatarUrl,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getGlobalSettings = `-- name: GetGlobalSettings :one
SELECT
    id, smtp_enabled, smtp_host, smtp_port, smtp_username, smtp_password, from_address, add_member_email_subject, add_member_email_template
FROM
    global_settings
LIMIT
    1
`

func (q *Queries) GetGlobalSettings(ctx context.Context) (GlobalSetting, error) {
	row := q.db.QueryRowContext(ctx, getGlobalSettings)
	var i GlobalSetting
	err := row.Scan(
		&i.ID,
		&i.SmtpEnabled,
		&i.SmtpHost,
		&i.SmtpPort,
		&i.SmtpUsername,
		&i.SmtpPassword,
		&i.FromAddress,
		&i.AddMemberEmailSubject,
		&i.AddMemberEmailTemplate,
	)
	return i, err
}

const getRecentConnectionsForTeam = `-- name: GetRecentConnectionsForTeam :many
SELECT
    connections.id,
    connections.subdomain,
    connections.created_at,
    connections.closed_at,
    users.email,
    users.first_name,
    users.last_name,
    users.github_avatar_url
FROM
    connections
    JOIN team_members ON team_members.id = connections.team_member_id
    JOIN users ON users.id = team_members.user_id
WHERE
    team_id = ?
ORDER BY
    connections.id DESC
LIMIT
    20
`

type GetRecentConnectionsForTeamRow struct {
	ID              int64
	Subdomain       string
	CreatedAt       time.Time
	ClosedAt        interface{}
	Email           string
	FirstName       interface{}
	LastName        interface{}
	GithubAvatarUrl interface{}
}

func (q *Queries) GetRecentConnectionsForTeam(ctx context.Context, teamID int64) ([]GetRecentConnectionsForTeamRow, error) {
	rows, err := q.db.QueryContext(ctx, getRecentConnectionsForTeam, teamID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetRecentConnectionsForTeamRow
	for rows.Next() {
		var i GetRecentConnectionsForTeamRow
		if err := rows.Scan(
			&i.ID,
			&i.Subdomain,
			&i.CreatedAt,
			&i.ClosedAt,
			&i.Email,
			&i.FirstName,
			&i.LastName,
			&i.GithubAvatarUrl,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTeamById = `-- name: GetTeamById :one
SELECT
    id, name, slug, created_at
FROM
    teams
WHERE
    id = ?
LIMIT
    1
`

func (q *Queries) GetTeamById(ctx context.Context, id int64) (Team, error) {
	row := q.db.QueryRowContext(ctx, getTeamById, id)
	var i Team
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Slug,
		&i.CreatedAt,
	)
	return i, err
}

const getTeamMemberByEmail = `-- name: GetTeamMemberByEmail :one
SELECT
    team_members.id, user_id, team_id, secret_key, role, added_by_user_id, team_members.created_at, users.id, email, first_name, last_name, is_super_user, github_access_token, github_avatar_url, users.created_at
FROM
    team_members
    JOIN users ON users.id = team_members.user_id
WHERE
    users.email = ?
LIMIT
    1
`

type GetTeamMemberByEmailRow struct {
	ID                int64
	UserID            int64
	TeamID            int64
	SecretKey         string
	Role              string
	AddedByUserID     interface{}
	CreatedAt         time.Time
	ID_2              int64
	Email             string
	FirstName         interface{}
	LastName          interface{}
	IsSuperUser       bool
	GithubAccessToken interface{}
	GithubAvatarUrl   interface{}
	CreatedAt_2       time.Time
}

func (q *Queries) GetTeamMemberByEmail(ctx context.Context, email string) (GetTeamMemberByEmailRow, error) {
	row := q.db.QueryRowContext(ctx, getTeamMemberByEmail, email)
	var i GetTeamMemberByEmailRow
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.TeamID,
		&i.SecretKey,
		&i.Role,
		&i.AddedByUserID,
		&i.CreatedAt,
		&i.ID_2,
		&i.Email,
		&i.FirstName,
		&i.LastName,
		&i.IsSuperUser,
		&i.GithubAccessToken,
		&i.GithubAvatarUrl,
		&i.CreatedAt_2,
	)
	return i, err
}

const getTeamMemberById = `-- name: GetTeamMemberById :one
SELECT
    team_members.id, team_members.user_id, team_members.team_id, team_members.secret_key, team_members.role, team_members.added_by_user_id, team_members.created_at,
    users.id, users.email, users.first_name, users.last_name, users.is_super_user, users.github_access_token, users.github_avatar_url, users.created_at
FROM
    team_members
    JOIN users ON users.id = team_members.user_id
WHERE
    team_members.id = ?
LIMIT
    1
`

type GetTeamMemberByIdRow struct {
	ID                int64
	UserID            int64
	TeamID            int64
	SecretKey         string
	Role              string
	AddedByUserID     interface{}
	CreatedAt         time.Time
	ID_2              int64
	Email             string
	FirstName         interface{}
	LastName          interface{}
	IsSuperUser       bool
	GithubAccessToken interface{}
	GithubAvatarUrl   interface{}
	CreatedAt_2       time.Time
}

func (q *Queries) GetTeamMemberById(ctx context.Context, id int64) (GetTeamMemberByIdRow, error) {
	row := q.db.QueryRowContext(ctx, getTeamMemberById, id)
	var i GetTeamMemberByIdRow
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.TeamID,
		&i.SecretKey,
		&i.Role,
		&i.AddedByUserID,
		&i.CreatedAt,
		&i.ID_2,
		&i.Email,
		&i.FirstName,
		&i.LastName,
		&i.IsSuperUser,
		&i.GithubAccessToken,
		&i.GithubAvatarUrl,
		&i.CreatedAt_2,
	)
	return i, err
}

const getTeamMemberByUserIdAndTeamSlug = `-- name: GetTeamMemberByUserIdAndTeamSlug :one
SELECT
    team_members.id, team_members.user_id, team_members.team_id, team_members.secret_key, team_members.role, team_members.added_by_user_id, team_members.created_at,
    users.id, users.email, users.first_name, users.last_name, users.is_super_user, users.github_access_token, users.github_avatar_url, users.created_at
FROM
    team_members
    JOIN users ON users.id = team_members.user_id
    JOIN teams ON teams.id = team_members.team_id
WHERE
    users.id = ?
    AND teams.slug = ?
LIMIT
    1
`

type GetTeamMemberByUserIdAndTeamSlugParams struct {
	ID   int64
	Slug string
}

type GetTeamMemberByUserIdAndTeamSlugRow struct {
	ID                int64
	UserID            int64
	TeamID            int64
	SecretKey         string
	Role              string
	AddedByUserID     interface{}
	CreatedAt         time.Time
	ID_2              int64
	Email             string
	FirstName         interface{}
	LastName          interface{}
	IsSuperUser       bool
	GithubAccessToken interface{}
	GithubAvatarUrl   interface{}
	CreatedAt_2       time.Time
}

func (q *Queries) GetTeamMemberByUserIdAndTeamSlug(ctx context.Context, arg GetTeamMemberByUserIdAndTeamSlugParams) (GetTeamMemberByUserIdAndTeamSlugRow, error) {
	row := q.db.QueryRowContext(ctx, getTeamMemberByUserIdAndTeamSlug, arg.ID, arg.Slug)
	var i GetTeamMemberByUserIdAndTeamSlugRow
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.TeamID,
		&i.SecretKey,
		&i.Role,
		&i.AddedByUserID,
		&i.CreatedAt,
		&i.ID_2,
		&i.Email,
		&i.FirstName,
		&i.LastName,
		&i.IsSuperUser,
		&i.GithubAccessToken,
		&i.GithubAvatarUrl,
		&i.CreatedAt_2,
	)
	return i, err
}

const getTeamMembers = `-- name: GetTeamMembers :many
SELECT
    users.email,
    team_members.role,
    users.github_avatar_url
FROM
    team_members
    JOIN users ON users.id = team_members.user_id
WHERE
    team_id = ?
`

type GetTeamMembersRow struct {
	Email           string
	Role            string
	GithubAvatarUrl interface{}
}

func (q *Queries) GetTeamMembers(ctx context.Context, teamID int64) ([]GetTeamMembersRow, error) {
	rows, err := q.db.QueryContext(ctx, getTeamMembers, teamID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetTeamMembersRow
	for rows.Next() {
		var i GetTeamMembersRow
		if err := rows.Scan(&i.Email, &i.Role, &i.GithubAvatarUrl); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTeamUserBySecretKey = `-- name: GetTeamUserBySecretKey :one
SELECT
    id, user_id, team_id, secret_key, role, added_by_user_id, created_at
FROM
    team_members
WHERE
    secret_key = ?
LIMIT
    1
`

func (q *Queries) GetTeamUserBySecretKey(ctx context.Context, secretKey string) (TeamMember, error) {
	row := q.db.QueryRowContext(ctx, getTeamUserBySecretKey, secretKey)
	var i TeamMember
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.TeamID,
		&i.SecretKey,
		&i.Role,
		&i.AddedByUserID,
		&i.CreatedAt,
	)
	return i, err
}

const getTeamsOfUser = `-- name: GetTeamsOfUser :many
SELECT
    teams.id, teams.name, teams.slug, teams.created_at
FROM
    team_members
    JOIN teams ON teams.id = team_members.team_id
WHERE
    team_members.user_id = ?
`

func (q *Queries) GetTeamsOfUser(ctx context.Context, userID int64) ([]Team, error) {
	rows, err := q.db.QueryContext(ctx, getTeamsOfUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Team
	for rows.Next() {
		var i Team
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Slug,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT
    id, email, first_name, last_name, is_super_user, github_access_token, github_avatar_url, created_at
FROM
    users
WHERE
    email = ?
LIMIT
    1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.FirstName,
		&i.LastName,
		&i.IsSuperUser,
		&i.GithubAccessToken,
		&i.GithubAvatarUrl,
		&i.CreatedAt,
	)
	return i, err
}

const getUserById = `-- name: GetUserById :one
SELECT
    users.id,
    users.email,
    users.created_at,
    users.first_name,
    users.last_name,
    users.github_avatar_url,
    users.is_super_user
FROM
    users
WHERE
    id = ?
LIMIT
    1
`

type GetUserByIdRow struct {
	ID              int64
	Email           string
	CreatedAt       time.Time
	FirstName       interface{}
	LastName        interface{}
	GithubAvatarUrl interface{}
	IsSuperUser     bool
}

func (q *Queries) GetUserById(ctx context.Context, id int64) (GetUserByIdRow, error) {
	row := q.db.QueryRowContext(ctx, getUserById, id)
	var i GetUserByIdRow
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.CreatedAt,
		&i.FirstName,
		&i.LastName,
		&i.GithubAvatarUrl,
		&i.IsSuperUser,
	)
	return i, err
}

const getUserBySession = `-- name: GetUserBySession :one
SELECT
    users.id,
    users.email,
    users.created_at,
    users.first_name,
    users.last_name,
    users.github_avatar_url,
    users.is_super_user
FROM
    users
    JOIN sessions ON sessions.user_id = users.id
WHERE
    sessions.token = ?
LIMIT
    1
`

type GetUserBySessionRow struct {
	ID              int64
	Email           string
	CreatedAt       time.Time
	FirstName       interface{}
	LastName        interface{}
	GithubAvatarUrl interface{}
	IsSuperUser     bool
}

func (q *Queries) GetUserBySession(ctx context.Context, token string) (GetUserBySessionRow, error) {
	row := q.db.QueryRowContext(ctx, getUserBySession, token)
	var i GetUserBySessionRow
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.CreatedAt,
		&i.FirstName,
		&i.LastName,
		&i.GithubAvatarUrl,
		&i.IsSuperUser,
	)
	return i, err
}

const getUsersCount = `-- name: GetUsersCount :one
SELECT
    COUNT(*)
FROM
    users
`

func (q *Queries) GetUsersCount(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, getUsersCount)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const markConnectionAsClosed = `-- name: MarkConnectionAsClosed :exec
UPDATE connections
SET
    closed_at = CURRENT_TIMESTAMP
WHERE
    id = ?
`

func (q *Queries) MarkConnectionAsClosed(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, markConnectionAsClosed, id)
	return err
}

const updateGlobalSettings = `-- name: UpdateGlobalSettings :exec
UPDATE global_settings
SET
    smtp_enabled = ?,
    smtp_host = ?,
    smtp_port = ?,
    smtp_username = ?,
    smtp_password = ?,
    from_address = ?,
    add_member_email_subject = ?,
    add_member_email_template = ?
`

type UpdateGlobalSettingsParams struct {
	SmtpEnabled            bool
	SmtpHost               interface{}
	SmtpPort               interface{}
	SmtpUsername           interface{}
	SmtpPassword           interface{}
	FromAddress            interface{}
	AddMemberEmailSubject  interface{}
	AddMemberEmailTemplate interface{}
}

func (q *Queries) UpdateGlobalSettings(ctx context.Context, arg UpdateGlobalSettingsParams) error {
	_, err := q.db.ExecContext(ctx, updateGlobalSettings,
		arg.SmtpEnabled,
		arg.SmtpHost,
		arg.SmtpPort,
		arg.SmtpUsername,
		arg.SmtpPassword,
		arg.FromAddress,
		arg.AddMemberEmailSubject,
		arg.AddMemberEmailTemplate,
	)
	return err
}

const updateSecretKey = `-- name: UpdateSecretKey :exec
UPDATE team_members
SET
    secret_key = ?
WHERE
    id = ?
`

type UpdateSecretKeyParams struct {
	SecretKey string
	ID        int64
}

func (q *Queries) UpdateSecretKey(ctx context.Context, arg UpdateSecretKeyParams) error {
	_, err := q.db.ExecContext(ctx, updateSecretKey, arg.SecretKey, arg.ID)
	return err
}

const updateUser = `-- name: UpdateUser :exec
UPDATE users
SET
    first_name = COALESCE(?, first_name),
    last_name = COALESCE(?, last_name),
    github_access_token = COALESCE(?, github_access_token),
    github_avatar_url = COALESCE(?, github_avatar_url)
WHERE
    id = ?
`

type UpdateUserParams struct {
	FirstName         interface{}
	LastName          interface{}
	GithubAccessToken interface{}
	GithubAvatarUrl   interface{}
	ID                int64
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.ExecContext(ctx, updateUser,
		arg.FirstName,
		arg.LastName,
		arg.GithubAccessToken,
		arg.GithubAvatarUrl,
		arg.ID,
	)
	return err
}
