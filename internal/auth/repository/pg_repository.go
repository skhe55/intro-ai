package repository

import (
	"context"
	"errors"
	"intro-ai/internal/auth"
	"intro-ai/internal/models"
	"intro-ai/pkg/utils"

	"github.com/jmoiron/sqlx"
)

type authRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) auth.Repository {
	return &authRepository{db: db}
}

func (r *authRepository) Register(ctx context.Context, user *models.User) (*models.User, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	var id uint64

	rowExistedUser := tx.QueryRowContext(
		ctx,
		`SELECT id FROM users WHERE login = $1`,
		user.Login,
	)

	if err := rowExistedUser.Err(); err != nil {
		return nil, utils.RollbackTransaction(tx, err)
	}

	if err := rowExistedUser.Scan(&id); err == nil {
		return nil, utils.RollbackTransaction(tx, errors.New("пользователь уже существует"))
	}

	var newUser models.User

	if err := user.HashPassword(); err != nil {
		return nil, utils.RollbackTransaction(tx, err)
	}

	row := tx.QueryRowContext(
		ctx,
		`INSERT INTO users (username, login, password) VALUES ($1, $2, $3) RETURNING *`,
		user.UserName,
		user.Login,
		user.Password,
	)

	if err := row.Err(); err != nil {
		return nil, utils.RollbackTransaction(tx, err)
	}

	if err := row.Scan(
		&newUser.ID,
		&newUser.UserName,
		&newUser.Login,
		&newUser.Password,
		&newUser.CreatedAt,
		&newUser.DeletedAt,
	); err != nil {
		return nil, utils.RollbackTransaction(tx, err)
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &newUser, nil
}

func (a *authRepository) Login(ctx context.Context, user *models.User) (*models.User, error) {
	tx, err := a.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	var loggedUser models.User

	row := tx.QueryRowContext(
		ctx,
		"SELECT * FROM users WHERE login = $1",
		user.Login,
	)

	if err := row.Err(); err != nil {
		return nil, utils.RollbackTransaction(tx, err)
	}

	if err := row.Scan(
		&loggedUser.ID,
		&loggedUser.UserName,
		&loggedUser.Login,
		&loggedUser.Password,
		&loggedUser.CreatedAt,
		&loggedUser.DeletedAt,
	); err != nil {
		return nil, utils.RollbackTransaction(tx, err)
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &loggedUser, nil
}

func (a *authRepository) GetUserById(ctx context.Context, id uint64) (*models.User, error) {
	conn, err := a.db.Connx(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var user *models.User

	row := conn.QueryRowxContext(
		ctx,
		"SELECT * FROM users WHERE id = $1 RETURNING *",
		id,
	)

	if err := row.Err(); err != nil {
		return nil, err
	}

	if err := row.StructScan(&user); err != nil {
		return nil, err
	}

	return user, nil
}
