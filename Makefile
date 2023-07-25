migration_up: 
	migrate -path database/migration/ -database ${PSQL_CONNECTION_STRING} -verbose up

migration_down: 
	migrate -path database/migration/ -database ${PSQL_CONNECTION_STRING} -verbose down

migration_fix: 
	migrate -path database/migration/ -database ${PSQL_CONNECTION_STRING} force VERSION

regen_gql_schema:
	go run github.com/99designs/gqlgen generate