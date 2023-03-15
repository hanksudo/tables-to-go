package tagger

import (
	"strings"

	"github.com/fraenky8/tables-to-go/pkg/database"
)

// Bun represents the Bun "bun"-tag.
type Bun struct{}

// GenerateTag for Bun to satisfy the Tagger interface.
func (t Bun) GenerateTag(db database.Database, column database.Column) string {

	isPk := ""
	if db.IsPrimaryKey(column) {
		isPk = ",pk"
	}

	isAutoIncrement := ""
	if db.IsAutoIncrement(column) {
		isAutoIncrement = ",autoincrement"
	}

	isNotNull := ""
	if column.IsNullable == "NO" {
		isNotNull = ",notnull"
	}

	defaultValue := ""
	if column.DefaultValue.Valid {
		v := strings.Split(column.DefaultValue.String, "::")[0]
		defaultValue = ",default:" + v
	}

	return `bun:"` + column.Name + isPk + isAutoIncrement + isNotNull + defaultValue + `"`
}
