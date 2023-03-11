package user

import (
	"database/sql"
)

type User struct {
	ID             int64          `db:"id"`
	FirstName      string         `db:"first_name"`
	SecondName     string         `db:"second_name"`
	Age            sql.NullInt64  `db:"age"`
	Biography      sql.NullString `db:"biography"`
	City           sql.NullString `db:"city"`
	Password       string         `db:"-"`
	HashedPassword []byte         `db:"pwdhsh"`
}
