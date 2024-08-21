// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"context"
)

type Querier interface {
	CreateTgBotUsers(ctx context.Context, arg CreateTgBotUsersParams) error
	GetActiveTgBotUsers(ctx context.Context) ([]int64, error)
	GetLichessData(ctx context.Context) ([]string, error)
	InsertLichessData(ctx context.Context, arg InsertLichessDataParams) error
	UpdateTgBotUsers(ctx context.Context, arg UpdateTgBotUsersParams) error
}

var _ Querier = (*Queries)(nil)