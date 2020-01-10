package dumper

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/jmoiron/sqlx"
)

// SQLDumper dumps databases
type SQLDumper struct {
	db *sqlx.DB
}

// New SQLDumper
func New(db *sqlx.DB) *SQLDumper {
	return &SQLDumper{
		db: db,
	}
}

// Dump a database to a file at path
func (d *SQLDumper) Dump(path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}

	tables, err := d.getTables()
	if err != nil {
		return err
	}

	for _, table := range tables {

		tHeader := fmt.Sprintf("CREATE TABLE %v (\n", table.Name)
		_, err := f.WriteString(tHeader)
		if err != nil {
			f.Close()
			return err
		}

		comma := false

		if len(table.Columns) > 0 {
			w := tabwriter.NewWriter(f, 4, 0, 1, ' ', 0)

			for _, column := range table.Columns {
				if !comma {
					comma = true
				} else {
					fmt.Fprintf(w, ",\n")
				}
				fmt.Fprintf(w, "\t%v\t%v", column.Name, getSQLColumnType(column))
			}

			if len(table.Constraints) > 0 {
				for _, constraint := range table.Constraints {
					skippable, _ := columnConstraints[constraint.Type]
					if skippable && len(constraint.Columns) == 1 {
						continue
					}

					if !comma {
						comma = true
					} else {
						fmt.Fprintf(w, ",\n")
					}

					fmt.Fprintf(w, "\t%v", constraint.Definition)
				}
			}

			err := w.Flush()
			if err != nil {
				f.Close()
				return err
			}
		}

		_, err = f.WriteString("\n);\n")
		if err != nil {
			f.Close()
			return err
		}

		if len(table.Indexes) > 0 {
			for _, index := range table.Indexes {
				//CREATE INDEX idx_account_last_logizn ON account(last_login, last_login desc);

				columns := []string{}
				for _, iColumn := range index.Columns {
					column := iColumn.Column.Name
					if len(iColumn.Direction) > 0 {
						column = column + " " + iColumn.Direction
					}

					columns = append(columns, column)
				}

				createIndex := fmt.Sprintf("CREATE INDEX ON %v(%v);\n", table.Name, strings.Join(columns, ", "))
				_, err = f.WriteString(createIndex)
				if err != nil {
					f.Close()
					return err
				}
			}
		}

		_, err = f.WriteString("\n")
		if err != nil {
			f.Close()
			return err
		}

	}

	return f.Close()
}

func (d *SQLDumper) getTables() ([]Table, error) {
	tables := []Table{}
	err := d.db.Select(&tables, getTablesQuery)
	for i := range tables {
		table := &tables[i]

		table.Columns, err = d.getColumns(table)
		if err != nil {
			return tables, err
		}

		table.Indexes, err = d.getIndexes(table)
		if err != nil {
			return tables, err
		}

		table.Constraints, err = d.getConstraints(table)
		if err != nil {
			return tables, err
		}
	}
	return tables, err
}

func (d *SQLDumper) getColumns(table *Table) ([]*Column, error) {
	columns := []*Column{}
	err := d.db.Select(&columns, getColumnsQuery, table.Name)
	return columns, err
}

func (d *SQLDumper) getIndexes(table *Table) ([]*Index, error) {
	indexes := []*Index{}
	rows, err := d.db.Queryx(getIndexesQuery, table.Name)
	if err != nil {
		return indexes, err
	}

	for rows.Next() {
		results := make(map[string]interface{})
		err = rows.MapScan(results)
		if err != nil {
			return indexes, err
		}

		//indexname := string(results["indexname"].([]byte))
		indexdef := results["indexdef"].(string)

		if indexdef[0:12] == "CREATE INDEX" {
			// Parse an index definition like
			// "CREATE INDEX table_col_index ON table USING btree (col_one, col_two DESC)"
			rDef := strings.Split(indexdef, "(")[1]
			def := rDef[:len(rDef)-1]
			cols := strings.Split(def, ",")

			columns := []*IndexColumn{}
			for _, col := range cols {
				cSplit := strings.Split(strings.Trim(col, " "), " ")
				colName := cSplit[0]
				colDir := ""

				if len(cSplit) > 1 {
					colDir = cSplit[1]
				}

				column := findColumnInTable(table, colName)

				columns = append(columns, &IndexColumn{
					Column:    column,
					Direction: colDir,
				})
			}

			indexes = append(indexes, &Index{
				Columns: columns,
			})
		}
	}

	return indexes, err
}

func (d *SQLDumper) getConstraints(table *Table) ([]*Constraint, error) {
	constraints := []*Constraint{}
	rows, err := d.db.Queryx(getConstraintsQuery, table.Name)
	if err != nil {
		return constraints, err
	}

ROW:
	for rows.Next() {
		results := make(map[string]interface{})
		err = rows.MapScan(results)
		if err != nil {
			return constraints, err
		}

		cType := results["constraint_type"].(string)
		def := results["definition"].(string)

		colsArg := string(results["columns"].([]byte))
		cols := strings.Split(colsArg[1:len(colsArg)-1], ",")

		constraint := &Constraint{
			Type:       cType,
			Definition: def,
			Columns:    []*Column{},
		}

		for _, col := range cols {
			column := findColumnInTable(table, col)

			if cType == "p" || cType == "f" {
				column.InferNotNull = true
			}

			skippable, _ := columnConstraints[constraint.Type]
			if skippable && len(cols) == 1 {
				switch cType {
				case "p":
					column.PrimaryKey = true
				case "u":
					column.Unique = true
				case "f":
					// parse "FOREIGN KEY (column) REFERENCES table(column)"
					keyParts := strings.Split(def, "REFERENCES ")
					column.ForeignKey = keyParts[1]
				}

				continue ROW
			}

			constraint.Columns = append(constraint.Columns, column)
		}

		constraints = append(constraints, constraint)
	}

	return constraints, err
}

func findColumnInTable(table *Table, columnName string) *Column {
	for i := range table.Columns {
		if table.Columns[i].Name == columnName {
			return table.Columns[i]
		}
	}
	// Error
	return &Column{}
}
