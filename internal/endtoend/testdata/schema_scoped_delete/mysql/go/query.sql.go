// Code generated by sqlc. DO NOT EDIT.
// source: query.sql

package querytest

import (
	"context"
)

const schemaScopedDelete = `-- name: SchemaScopedDelete :exec
DELETE FROM foo.bar WHERE id = ?
`

func (q *Queries) SchemaScopedDelete(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, schemaScopedDelete, id)
	return err
}
