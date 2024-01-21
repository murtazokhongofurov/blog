# Blog

## 1. Clone Project
```
https://github.com/murtazokhongofurov/blog.git
```

## 2. Add Environment Variables file using example.env file
Create an environment variables file using the provided example.env file and configure the following variables: <br>

HTTP_PORT=8181 <br>
DB_USER=username <br>
DB_PASS=pass <br>
DB_HOST=localhost <br>
DB_PORT=5432 <br>
DB_NAME=dbname <br>
AUTH_FILE_PATH=./config/auth.conf <br>
CSV_FILE_PATH=./config/roles.csv <br>
SIGNING_KEY=dfkljadlkvnwe834riekkj3489uf39ri3 <br>
MAIL_USERNAME=username <br>
MAIL_PASSWORD=password <br>
SMTP_HOST=host <br>
LOG_LEVEL=debug <br>

## 3. Run migrations by this following command
Execute the following command to run migrations: <br>
```
    make migrate_up
```

## 4. Run project
You can run the project using either of the following commands: <br>
```
make run
```

or

```
go run cmd/main.go
```

## 5. Open Swagger Api docs and check apis
Visit the Swagger API documentation at: <br>
http://localhost:8181/v1/docs/index.html

## 6. Run test by command bellow
Execute the following command to run tests and generate coverage: <br>
```
go test ./... -cover
```

## Test Result:
![alt text](images/test.png)


