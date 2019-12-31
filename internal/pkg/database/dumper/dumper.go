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
			w := tabwriter.NewWriter(f, 8, 8, 0, '\t', 0)

			for _, column := range table.Columns {
				if !comma {
					comma = true
				} else {
					fmt.Fprintf(w, ",\n")
				}
				fmt.Fprintf(w, "\t%v\t%v", column.Name, getSQLColumnType(&column))
			}

			if len(table.Indexes) > 0 {
				for _, index := range table.Indexes {
					if !comma {
						comma = true
					} else {
						fmt.Fprintf(w, ",\n")
					}

					fmt.Fprintf(w, "\t%v (%v)", index.Type, strings.Join(index.ColumnNames, ","))
				}
			}

			err := w.Flush()
			if err != nil {
				f.Close()
				return err
			}
		}

		_, err = f.WriteString("\n);\n\n")
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
		tables[i].Columns, err = d.getColumns(&tables[i])
		if err != nil {
			return tables, err
		}

		tables[i].Indexes, err = d.getIndexes(&tables[i])
		if err != nil {
			return tables, err
		}
	}
	return tables, err
}

func (d *SQLDumper) getColumns(table *Table) ([]Column, error) {
	columns := []Column{}
	err := d.db.Select(&columns, getColumnsQuery, table.Name)
	return columns, err
}

func (d *SQLDumper) getIndexes(table *Table) ([]Index, error) {
	indexes := []Index{}
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
		
		indexname := string(results["indexname"].([]byte))
		indexdef := results["indexdef"].(string)

		iType := "Error"

		defCreateUniqueIndex := indexdef[0:19] == "CREATE UNIQUE INDEX"
		defCreateIndex := indexdef[0:12] == "CREATE INDEX"

		typeIdentifier := indexname[len(indexname)-4:]

		if defCreateIndex {
			iType = "INDEX"
		} else if defCreateUniqueIndex {
			switch typeIdentifier {
			case "pkey":
				iType = "PRIMARY KEY"
			case "_key":
				iType = "UNIQUE"
			}
		}

		// Parse an index definition like
		// "CREATE UNIQUE INDEX table_pkey ON users USING btree (provider, name)"
		def := strings.Split(indexdef, "(")[1]
		columns := strings.Split(def[:len(def)-1], ",")

		indexes = append(indexes, Index{
			Type:        iType,
			ColumnNames: columns,
		})
	}

	return indexes, err
}
