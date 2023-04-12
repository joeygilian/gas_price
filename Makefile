backend:
	go run app/main.go

db: 
	docker run --name investree_test_db -p 5432:5432 -e POSTGRES_PASSWORD=password -d postgres

init-db:
	cat ./initdb/query.sql | docker exec -i investree_test_db psql -U postgres -d postgres