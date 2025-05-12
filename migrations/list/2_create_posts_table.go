package list

import (
	pg "github.com/ShkrutDenis/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreatePostTable struct{}

func (m *CreatePostTable) GetName() string {
	return "CreatePostTable"
}

func (m *CreatePostTable) Up(con *sqlx.DB) {
	table := pg.NewTable("posts", con)
	table.Column("id").Type("integer").Autoincrement()
	table.PrimaryKey("id")
	table.String("title", 500).Nullable()
	table.String("content", 1000).Nullable()
	table.Column("deleted_at").Type("timestamptz").Nullable()
	table.Column("user_id").Type("integer")
	table.ForeignKey("user_id").
		Reference("users").
		On("id").
		OnDelete("cascade").
		OnUpdate("cascade")
	table.WithTimestamps()

	table.MustExec()
}

func (m *CreatePostTable) Down(con *sqlx.DB) {
	pg.DropTable("posts", con).MustExec()
}
