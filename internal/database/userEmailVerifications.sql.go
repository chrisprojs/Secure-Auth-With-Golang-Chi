// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: userEmailVerifications.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const checkEmailVerification = `-- name: CheckEmailVerification :one
SELECT 
  res_message::TEXT,
  emailverify_id::UUID,
  user_id::VARCHAR,
  username::VARCHAR,
  email::VARCHAR,
  used_for::VARCHAR
FROM check_email_verification($1, $2)
AS result(res_message, emailverify_id, user_id, username, email, used_for)
`

type CheckEmailVerificationParams struct {
	VerifCode     string    `json:"_verif_code"`
	EmailverifyID uuid.UUID `json:"_emailverify_id"`
}

type CheckEmailVerificationRow struct {
	ResMessage    string    `json:"res_message"`
	EmailverifyID uuid.UUID `json:"emailverify_id"`
	UserID        string    `json:"user_id"`
	Username      string    `json:"username"`
	Email         string    `json:"email"`
	UsedFor       string    `json:"used_for"`
}

func (q *Queries) CheckEmailVerification(ctx context.Context, arg CheckEmailVerificationParams) (CheckEmailVerificationRow, error) {
	row := q.db.QueryRowContext(ctx, checkEmailVerification, arg.VerifCode, arg.EmailverifyID)
	var i CheckEmailVerificationRow
	err := row.Scan(
		&i.ResMessage,
		&i.EmailverifyID,
		&i.UserID,
		&i.Username,
		&i.Email,
		&i.UsedFor,
	)
	return i, err
}

const createEmailVerification = `-- name: CreateEmailVerification :one
SELECT
user_id::VARCHAR,
emailverify_id::UUID
FROM generate_email_verification($1,$2,$3)
AS result(user_id,emailverify_id)
`

type CreateEmailVerificationParams struct {
	UserID    string `json:"_user_id"`
	VerifCode string `json:"_verif_code"`
	UsedFor   string `json:"_used_for"`
}

type CreateEmailVerificationRow struct {
	UserID        string    `json:"user_id"`
	EmailverifyID uuid.UUID `json:"emailverify_id"`
}

func (q *Queries) CreateEmailVerification(ctx context.Context, arg CreateEmailVerificationParams) (CreateEmailVerificationRow, error) {
	row := q.db.QueryRowContext(ctx, createEmailVerification, arg.UserID, arg.VerifCode, arg.UsedFor)
	var i CreateEmailVerificationRow
	err := row.Scan(&i.UserID, &i.EmailverifyID)
	return i, err
}

const fetchEmailVerification = `-- name: FetchEmailVerification :one
SELECT
  EXTRACT(EPOCH FROM (expires_at - CURRENT_TIMESTAMP))::INTEGER as time_left
FROM useremailverifications
WHERE emailverify_id = $1
`

func (q *Queries) FetchEmailVerification(ctx context.Context, emailverifyID uuid.UUID) (int32, error) {
	row := q.db.QueryRowContext(ctx, fetchEmailVerification, emailverifyID)
	var time_left int32
	err := row.Scan(&time_left)
	return time_left, err
}

const resendEmailVerificationCode = `-- name: ResendEmailVerificationCode :one
UPDATE UserEmailVerifications
SET verif_code = $1, expires_at = CURRENT_TIMESTAMP + INTERVAL '60 seconds'
FROM users
WHERE UserEmailVerifications.emailverify_id = $2
  AND users.user_id = UserEmailVerifications.user_id
RETURNING
UserEmailVerifications.emailverify_id,
EXTRACT(EPOCH FROM (UserEmailVerifications.expires_at - CURRENT_TIMESTAMP))::INTEGER as time_left,
users.email
`

type ResendEmailVerificationCodeParams struct {
	VerifCode     string    `json:"verif_code"`
	EmailverifyID uuid.UUID `json:"emailverify_id"`
}

type ResendEmailVerificationCodeRow struct {
	EmailverifyID uuid.UUID `json:"emailverify_id"`
	TimeLeft      int32     `json:"time_left"`
	Email         string    `json:"email"`
}

func (q *Queries) ResendEmailVerificationCode(ctx context.Context, arg ResendEmailVerificationCodeParams) (ResendEmailVerificationCodeRow, error) {
	row := q.db.QueryRowContext(ctx, resendEmailVerificationCode, arg.VerifCode, arg.EmailverifyID)
	var i ResendEmailVerificationCodeRow
	err := row.Scan(&i.EmailverifyID, &i.TimeLeft, &i.Email)
	return i, err
}

const updateUserEmail = `-- name: UpdateUserEmail :one
SELECT
  message::VARCHAR,
  user_id::VARCHAR,
  username::VARCHAR,
  email::VARCHAR
FROM update_user_email($1, $2)
AS result(message,user_id,username,email)
`

type UpdateUserEmailParams struct {
	UserID        string    `json:"_user_id"`
	EmailverifyID uuid.UUID `json:"_emailverify_id"`
}

type UpdateUserEmailRow struct {
	Message  string `json:"message"`
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (q *Queries) UpdateUserEmail(ctx context.Context, arg UpdateUserEmailParams) (UpdateUserEmailRow, error) {
	row := q.db.QueryRowContext(ctx, updateUserEmail, arg.UserID, arg.EmailverifyID)
	var i UpdateUserEmailRow
	err := row.Scan(
		&i.Message,
		&i.UserID,
		&i.Username,
		&i.Email,
	)
	return i, err
}
