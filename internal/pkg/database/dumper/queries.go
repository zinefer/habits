package dumper

const getTablesQuery = `SELECT tablename as name FROM pg_catalog.pg_tables WHERE
schemaname != 'pg_catalog' AND schemaname != 'information_schema';`

const getColumnsQuery = `SELECT column_name as name, data_type as type, 
column_default as default, is_nullable as nullable, 
character_maximum_length as limit FROM 
information_schema.columns WHERE table_name = $1;`

const getConstraintsQuery = `SELECT c.conname                                     AS constraint_name,
       c.contype                                     AS constraint_type,
       sch.nspname                                   AS "schema",
       tbl.relname                                   AS "table",
       ARRAY_AGG(col.attname ORDER BY u.attposition) AS columns,
       pg_get_constraintdef(c.oid)                   AS definition
FROM pg_constraint c
       JOIN LATERAL UNNEST(c.conkey) WITH ORDINALITY AS u(attnum, attposition) ON TRUE
       JOIN pg_class tbl ON tbl.oid = c.conrelid
       JOIN pg_namespace sch ON sch.oid = tbl.relnamespace
       JOIN pg_attribute col ON (col.attrelid = tbl.oid AND col.attnum = u.attnum)
WHERE tbl.relname = $1
GROUP BY constraint_name, constraint_type, "schema", "table", definition
ORDER BY "schema", "table";`

const getIndexesQuery = `SELECT indexname, indexdef FROM pg_indexes WHERE tablename = $1`
