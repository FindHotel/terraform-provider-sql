package dialect

import (
	"fmt"
	"log"
	"reflect"
	"strings"

	"gopkg.in/gorp.v1"
)

///////////////////////////////////////////////////////
// Snowflake  //
////////////////

type SnowflakeDialect struct {
	suffix string
}

func (d SnowflakeDialect) QuerySuffix() string { return ";" }

func (d SnowflakeDialect) ToSqlType(val reflect.Type, maxsize int, isAutoIncr bool) string {
	switch val.Kind() {
	case reflect.Ptr:
		return d.ToSqlType(val.Elem(), maxsize, isAutoIncr)
	case reflect.Bool:
		return "BOOLEAN"
	case reflect.Int:
		return "NUMBER"
	case reflect.Int8, reflect.Uint8:
		return "NUMBER"
	case reflect.Int16, reflect.Uint16:
		return "NUMBER"
	case reflect.Int32, reflect.Uint32:
		return "NUMBER"
	case reflect.Int64, reflect.Uint64:
		return "NUMBER"
	case reflect.Float64:
		return "FLOAT"
	case reflect.Float32:
		return "FLOAT"
	case reflect.Slice:
		if val.Elem().Kind() == reflect.Uint8 {
			return "BINARY"
		}
	}

	switch val.Name() {
	case "NullInt64":
		return "NUMBER"
	case "NullFloat64":
		return "FLOAT"
	case "NullBool":
		return "BOOLEAN"
	case "Time":
		return "TIMESTAMP_TZ"
	}

	if maxsize > 0 {
		return fmt.Sprintf("VARCHAR(%d)", maxsize)
	} else {
		return "TEXT"
	}

}

// Returns empty string
func (d SnowflakeDialect) AutoIncrStr() string {
	return "AUTOINCREMENT"
}

func (d SnowflakeDialect) AutoIncrBindValue() string {
	return ""
}

func (d SnowflakeDialect) AutoIncrInsertSuffix(col *gorp.ColumnMap) string {
	return ""
}

// Returns suffix
func (d SnowflakeDialect) CreateTableSuffix() string {
	return d.suffix
}

func (d SnowflakeDialect) TruncateClause() string {
	return "truncate"
}

func (d SnowflakeDialect) BindVar(i int) string {
	return "?"
}

func standardInsertAutoIncr(exec gorp.SqlExecutor, insertSql string, params ...interface{}) (int64, error) {
	log.Printf("EXEC standardInsertAutoIncr: q: %s, params: %+v", insertSql, params)
	res, err := exec.Exec(insertSql, params...)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (d SnowflakeDialect) InsertAutoIncr(exec gorp.SqlExecutor, insertSql string, params ...interface{}) (int64, error) {
	return standardInsertAutoIncr(exec, insertSql, params...)
}

func (d SnowflakeDialect) QuoteField(f string) string {
	return `"` + strings.ToLower(f) + `"`
}

func (d SnowflakeDialect) QuotedTableForQuery(schema string, table string) string {
	if strings.TrimSpace(schema) == "" {
		return d.QuoteField(table)
	}

	return schema + "." + d.QuoteField(table)
}

func (d SnowflakeDialect) IfSchemaNotExists(command, schema string) string {
	return fmt.Sprintf("%s if not exists", command)
}

func (d SnowflakeDialect) IfTableExists(command, schema, table string) string {
	return fmt.Sprintf("%s if exists", command)
}

func (d SnowflakeDialect) IfTableNotExists(command, schema, table string) string {
	return fmt.Sprintf("%s if not exists", command)
}
