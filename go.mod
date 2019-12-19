module github.com/zinefer/habits

go 1.12

require (
	github.com/go-chi/chi v4.0.2+incompatible
	github.com/gorilla/sessions v1.1.1
	github.com/jmoiron/sqlx v1.2.0
	github.com/lib/pq v1.0.0
	github.com/markbates/goth v1.59.0
	golang.org/x/net v0.0.0-20191209160850-c0dbc17a3553 // indirect
	golang.org/x/oauth2 v0.0.0-20191202225959-858c2ad4c8b6 // indirect
)

replace github.com/markbates/goth v1.59.0 => github.com/zinefer/goth v1.59.1-0.20191216233856-3e2a3b141469
