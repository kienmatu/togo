### Notes
This is a simple backend for a todo service, right now this service can handle login/list/create simple tasks, to make it run:
- `go run main.go`
- Import Postman collection from docs to check example

Candidates are invited to implement the below requirements but the point is **NOT to resolve everything perfectly** but selective about what you can do best in a limited time.  
Thus, **there is no correct-or-perfect answer**, your solutions are a way for us to continue the discussion and collaboration.  

We're using **Golang** but candidates can use any language (NodeJS/Java/PHP/Python...) **as long as**:  
- You show us how to run **reliably** - many of us use Ubuntu, some use Mac
- Your solution is **compatible with our REST interface** and we can use our Postman collection for verifying

---

### Functional requirement:
Right now a user can add as many tasks as they want, we want the ability to **limit N tasks per day**.  
For example, users are limited to create only 5 tasks only per day, if the daily limit is reached, return 4xx code to the client and ignore the create request.

### Non-functional requirements:
- [ ] **A nice README on how to run, what is missing, what else you want to improve but don't have enough time**
- [ ] **Consistency is a MUST**
- [ ] Fork this repo and show us your development progress by a PR
- [ ] Write integration tests for this project
- [ ] Make this code DRY
- [ ] Write unit test for the services layer
- [ ] Change from using SQLite to Postgres with docker-compose
- [ ] This project includes many issues from code to DB structure, feel free to optimize them
- [ ] Write unit test for storages layer
- [ ] Split services layer to use case and transport layer


#### DB Schema
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

#### Sequence diagram
![auth and create tasks request](https://github.com/manabie-com/togo/blob/master/docs/sequence.svg)
