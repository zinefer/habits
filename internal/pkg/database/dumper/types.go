package dumper

import (
	"database/sql"
	"strconv"
	"strings"
)

// Table describes a sql table
type Table struct {
	Name        string
	Columns     []*Column
	Indexes     []*Index
	Constraints []*Constraint
}

// Column describes a sql column
type Column struct {
	Name       string
	Type       string
	Default    sql.NullString
	Nullable   string
	Limit      sql.NullInt64
	PrimaryKey bool
	ForeignKey string
	Unique     bool
	// If a column is contained inside a constraint we do not need to output NOT NULL
	InferNotNull bool
}

// Constraint describes a sql constraint
type Constraint struct {
	Columns    []*Column
	Type       string
	Definition string
}

// Index describes a sql index
type Index struct {
	Columns []*IndexColumn
}

// IndexColumn describes an indexed column
type IndexColumn struct {
	Column    *Column
	Direction string
}

// Column constraints can be marked on the column definition if they only refer to a
// single column
var columnConstraints = map[string]bool{
	"p": true,
	"u": true,
	"f": true,
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
		if columnIsSerial(column) {
			cType = "SERIAL"
		}

		if column.Limit.Valid {
			cType = cType + "(" + strconv.FormatInt(column.Limit.Int64, 10) + ")"
		}

		constraint := getSQLColumnConstraint(column)
		if len(constraint) > 0 {
			cType = cType + " " + constraint
		}

		defaultVar := getSQLColumnDefault(column)
		if len(defaultVar) > 0 {
			cType = cType + " " + defaultVar
		}

		return cType
	}

	return "ERROR"
}

func columnIsSerial(column *Column) bool {
	return column.Type == "integer" &&
		len(column.Default.String) > 8 &&
		column.Default.String[0:8] == "nextval("
}

func getSQLColumnDefault(column *Column) string {
	if columnIsSerial(column) {
		return ""
	}

	defaultVar := column.Default.String
	len := len(defaultVar)
	if len > 0 {
		if len > 6 && defaultVar[len-6:] == "::text" {
			defaultVar = defaultVar[0 : len-6]
		}

		return "DEFAULT " + defaultVar
	}

	return ""
}

func getSQLColumnConstraint(column *Column) string {
	if column.PrimaryKey {
		return "PRIMARY KEY"
	}

	if len(column.ForeignKey) > 0 {
		return "REFERENCES " + column.ForeignKey
	}

	out := []string{}

	if column.Unique {
		out = append(out, "UNIQUE")
	}

	if !column.InferNotNull && column.Nullable == "NO" && !columnIsSerial(column) {
		out = append(out, "NOT NULL")
	}

	return strings.Join(out, " ")
}
