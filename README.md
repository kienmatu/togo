## GOLANG TODOS APPLICATION

### Technical stuff

- Architecture: Clean architecture
- Framework: Echo
- ORM: Gorm
- DB: Postgres
- Deployment: Docker

### Overview

- Support JWT
- Limit todos per user in a day
- Unit tests & integration test

### What next ?

- Add role/permission based validation
- Add missing tests
- Implement your new features

### How to run the code locally

Clone the project then:

##### Update .env file

```txt
PORT=8080
JWT_SECRET=B5bJHoI8aVLjAAeV
SIGNING_KEY=ABC
HASH_SALT=SJSHDFDS
TOKEN_TTL=86400
CONNECTION_URL=host=localhost user=postgres password=password1 dbname=todos port=5432
```

```bash
go run cmd/api/main.go
```

##### or by Docker

Update .env file (change host to postgresql)

```txt
#...
CONNECTION_URL=host=postgresql user=postgres password=password1 dbname=todos port=5432
```

RUN COMMAND:

```bash
docker-compose up -d
```

### Then open postman or bash use curl:

#### Register:

```bash
curl --location --request POST 'http://localhost:8080/api/v1/auth/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username": "kiendinh",
    "password": "abc123",
    "limit": 2
}'
```

#### Login:

```bash
curl --location --request POST 'http://localhost:8080/api/v1/auth/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username": "kiendinh",
    "password": "abc123"
}'
```

Take the token from login then.

#### Add todo:

```bash
curl --location --request POST 'http://localhost:8080/api/v1/todos/' \
--header 'Authorization: Bearer YOUR_TOKEN_HERE' \
--header 'Content-Type: application/json' \
--data-raw '{
    "content": "TEST TO DO 2x"
}'
```

#### Get all todos:

Public for all users

```bash
curl --location --request GET 'http://localhost:8080/api/v1/todos/'
```

#### Get user todos:

Get any user's todos with their id

```bash
curl --location --request GET 'http://localhost:8080/api/v1/todos/A_USER_ID' \
--header 'Authorization: Bearer YOUR_TOKEN_HERE'
```

### TESTS

I just implemented some necessary test samples, we can add tests of handlers, usecases and another endpoints ...

#### UNIT TEST

Command:

```bash
go test ./...
```

#### INTEGRATION TEST

Command:

```bash
go test -v ./integration-tests
```

### Conclusion

First time I created a new microservice with Go from scratch, It gave me a challenge but I did it, tried to remove the old mindset in another architecture then go to clean architecture. I love it :heart: :heart: .
