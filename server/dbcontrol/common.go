package dbcontrol

import (
	"context"

	"github.com/georgysavva/scany/pgxscan"
)

const (
	membership_query = `SELECT COUNT(*) FROM participants WHERE
	course_id = $1 AND
	user_id = (SELECT user_id FROM users WHERE email = $2);`
)

func (c *DBControl) IsMember(email string, course_id int) (int32, error) {
	ctx := context.Background()
	rows, err := c.pool.Query(ctx, membership_query, course_id, email)
	if err != nil {
		return 0, err
	}
	var count int32
	if err = pgxscan.ScanOne(&count, rows); err != nil {
		return 0, err
	}
	return count, nil
}
