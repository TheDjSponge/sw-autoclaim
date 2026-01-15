package database

import (
	"context"
)

func (q *Queries) DeleteUserAndCount(ctx context.Context, arg DeleteUserParams) (int, error) {
	tag, err := q.DeleteUser(ctx, arg)
    return int(tag.RowsAffected()), err
}