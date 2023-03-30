package tagger

import (
	"strings"

	"github.com/hanksudo/tables-to-go/pkg/database"
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
	defaultValue := ""
	isNotNull := ""

	if db.IsAutoIncrement(column) {
		isAutoIncrement = ",autoincrement"
	} else {
		if column.DefaultValue.Valid {
			v := strings.Split(column.DefaultValue.String, "::")[0]
			defaultValue = ",default:" + v
		}
		if column.IsNullable == "NO" {
			isNotNull = ",notnull"
		}
	}

	return `bun:"` + column.Name + isPk + isAutoIncrement + isNotNull + defaultValue + `"`
}
