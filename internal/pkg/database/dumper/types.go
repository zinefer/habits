package dumper

import (
	"database/sql"
	"strconv"
)

// Table describes a sql table
type Table struct {
	Name    string
    Columns []Column
    Indexes []Index
}

// Column describes a sql column
type Column struct {
	Name     string
	Type     string
	Default  sql.NullString
	Nullable string
	Limit    sql.NullInt64
}

// Index describes a sql index
type Index struct {
    Type string
    ColumnNames []string
}

var columnTypeToSQL = map[string]string{
	"text":                        "TEXT",
	"character varying":           "VARCHAR",
	"character":                   "CHAR",
	"timestamp without time zone": "TIMESTAMP",
	"timestamp with time zone":    "TIMESTAMP WITH TIME ZONE",
	"integer":                     "INTEGER",
	"boolean":                     "BOOLEAN",
}

func getSQLColumnType(column *Column) string {
	if cType, ok := columnTypeToSQL[column.Type]; ok {
		switch cType {
		case "INTEGER":
			if column.Default.String[0:8] == "nextval(" {
				cType = "SERIAL"
			}

		}

		if column.Limit.Valid {
			cType = cType + "(" + strconv.FormatInt(column.Limit.Int64, 10) + ")"
        }
        
        options := getSQLColumnOptions(column)
        if len(options) > 0 {
            cType = cType + " " + options
        }

		return cType
	}

	return "ERROR"
}

func getSQLColumnOptions(column *Column) string {
	if column.Nullable == "NO" {
		return "NOT NULL"
	}

	return ""
}
