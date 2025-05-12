package list

import (
	pg "github.com/ShkrutDenis/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateUserTable struct{}

func (m *CreateUserTable) GetName() string {
	return "CreateUserTable"
}

func (m *CreateUserTable) Up(con *sqlx.DB) {
	table := pg.NewTable("users", con)
	table.Column("id").Type("int unsigned").Autoincrement()
	table.PrimaryKey("id")
	table.String("name", 500).Nullable()
	table.String("password", 1000).Nullable()
	table.Column("deleted_at").Type("timestamptz").Nullable()
	table.WithTimestamps()

	table.MustExec()
}

func (m *CreateUserTable) Down(con *sqlx.DB) {
	pg.DropTable("users", con).MustExec()
}
