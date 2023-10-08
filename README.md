# go.reading.list

This is an example of Go web service and application that uses a Postgres database.

Taken from the Pluralsight course 'Building Go Web Services and Applications' by Josh Duffney


## Postgres Database ##

Run Postgres in Docker: docker run --name reading-list-db-container -e POSTGRES_PASSWORD=secretpassword -d -p 5432:5432 postgres

Shell into Postgres container: docker exec -it reading-list-db-container /bin/bash

Connect to Postgress instance (likely not prompted for password inside container): psql -h localhost -p 5432 -U postgres

Create database: CREATE DATABASE readinglist;

Create new user so don't need superuser password: CREATE ROLE readinglist WITH LOGIN PASSWORD 'pa55w0rd';

Change to the 'readinglist' database: \c readinglist

Create table: CREATE TABLE IF NOT EXISTS books (
                id bigserial PRIMARY KEY,
                create_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
                title text NOT NULL,
                published integer NOT NULL,
                pages integer NOT NULL,
                genres text[] NOT NULL,
                rating real NOT NULL,
                version integer NOT NULL DEFAULT 1);

Grant the 'readinglist' user access to the table: GRANT SELECT, INSERT, UPDATE, DELETE ON TABLE books TO readinglist;

As using 'bigserial' type  for 'id' need sequence permissions: GRANT USAGE, SELECT ON SEQUENCE books_id_seq TO readinglist;

The variable with dsn: export READINGLIST_DB_DSN 'postgres://readinglist:pa55w0rd@localhost:5432/readinglist?sslmode=disable'

Also remember to import the 'database/sql' and 'github.com/lib/pq' (will need go get github.com/lib/pq) packages.

## Routes ##

Health check: curl localhost:4000/v1/healthcheck

Create a post body: BODY='{"title":"Lord of the Rings","published":1954,"pages":1170,"genres":["Fiction","Fantasy"],"rating":4.7}'
Create a post body: BODY='{"title":"The Hitchhiker'\''s Guide to the Galaxy","published":1979,"pages":224,"genres":["Fiction","Comedy","Science Fiction"],"rating":4.5}'
Create a post body: BODY='{"title":"A Game of Thrones","published":1996,"pages":694,"genres":["Fiction","Fantasy"],"rating":4.4}'

Add book to list: curl -i -d "$BODY" localhost:4000/v1/books

Get books on list: curl localhost:4000/v1/books

Get details of book with ID 2: curl localhost:4000/v1/books/2

Update rating of a book: BODY='{"title":"The Hitchhiker'\''s Guide to the Galaxy","published":1979,"pages":224,"genres":["Fiction","Comedy","Science Fiction"],"rating":4.6}'
Update details of book with ID 2: curl -X PUT -d "$BODY" localhost:4000/v1/books/2

Delete book with ID 2: curl -X DELETE localhost:4000/v1/books/2