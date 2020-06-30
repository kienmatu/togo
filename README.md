### Overview
This is a simple backend for good old todo service, right now this service can handle login/list/create simple tasks.  
To make it run:
- `go run main.go`
- Import Postman collection from `docs` to check example

Candidates are invited to implement below requirements but the point is not to resolve everything in a perfect way but selective what you can do best in a limited time.  
Thus, there is no correct-or-perfect answer, your solutions are way for us to continue the discussion and collaborating.
 
### Requirements
Right now a user can add many task as they want, we want ability to limit N task per day.

Example: users are limited to create only 5 task only per day, if imit reached, return 4xx code to client and ignore the create request.
#### Backend requirements
- Write integration tests for this project
- Make this code DRY
- Write unit test for `services` layer
- Change from using `SQLite` to `Postgres` with `docker-compose`
#### Frontend requirements
- Login interface
- Token management state using `Redux`
- UI for list/create tasks
#### Optional requirements
- Write unit test for `storages` layer
- Split `services` layer to `use case` and `transport` layer

### DB Schema
```sql
-- users definition

CREATE TABLE users (
	id TEXT NOT NULL,
	password TEXT NOT NULL,
	max_todo INTEGER DEFAULT 5 NOT NULL,
	CONSTRAINT users_PK PRIMARY KEY (id)
);

INSERT INTO users (id, password, max_todo) VALUES('firstUser', 'example', 5);

-- tasks definition

CREATE TABLE tasks (
	id TEXT NOT NULL,
	content TEXT NOT NULL,
	user_id TEXT NOT NULL,
    created_date TEXT NOT NULL,
	CONSTRAINT tasks_PK PRIMARY KEY (id),
	CONSTRAINT tasks_FK FOREIGN KEY (user_id) REFERENCES users(id)
);
```

### Sequence diagram
![auth and create tasks request](https://github.com/manabie-com/togo/blob/master/docs/sequence.svg)