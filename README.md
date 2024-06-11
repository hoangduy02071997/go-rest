# Restful API Golang

<img align="right" width="159px" src="https://raw.githubusercontent.com/gin-gonic/logo/master/color.png">

Golang Restfull API with modules and Clean Architecture

## What's included?

- [Gin Gonic](https://github.com/gin-gonic/gin) Gin Web Framework <!-- more info? -->
- [Gorm](https://gorm.io/) Object Relation Mapping
- [codegangsta/gin](https://github.com/codegangsta/gin) for live reloading in local development

### Making changes to code

#### Directory structure

```bash
.
├── common             # common func reusable
├── middleware         # all middleware of module
├── modules            # feature as module             
```

#### Edit & Deploy

Edit the code in `index.go`, for example, un-comment the commented lines:

```bash
$ vim index.go
```

### Using database without an ORM

- Create required tables/columns using API Console (Data -> Schema)
- Use Query Builder under API Explorer and create the query
- Replicate the same JSON query in your gin app

<!-- commented until golang codegen is available
- Click on Generate API Code button and select Golang Requests
- Copy and paste the Go code into your Gin app source code
-->

#### Example

To get all entries from `todos` table, the query would look like the following, using [levigross/grequests](https://github.com/levigross/grequests):

```golang
resp, err := grequests.Post("https://localhost:8080/api/v1/items",
    &grequests.RequestOptions{
        JSON: map[string]interface{}{
            "type": "select",
            "args": map[string]interface{}{
                "table":   "todos",
                "columns": []string{"*"},
            },
        },
    },
)
if err != nil {
    fmt.Printf("error: %s", err)
}
if !resp.Ok {
    fmt.Printf("status code: %d, data: %s", resp.StatusCode, string(resp.Bytes())),
}

fmt.Printf("response: %s", string(resp.Bytes()))
```

The output will be something similar to:

```json
[
  {
    "title": "Generated from 99997",
    "description": "Generated from 99997",
    "status": "todo",
    "id": 9890531,
    "created_at": "2024-06-10T23:30:01+07:00",
    "updated_at": "2024-06-10T23:30:01+07:00"
  },
  {
    "title": "Generated from 99997",
    "description": "Generated from 99997",
    "status": "todo",
    "id": 9890530,
    "created_at": "2024-06-10T23:30:01+07:00",
    "updated_at": "2024-06-10T23:30:01+07:00"
  },
  ...
]
```


### Using database with an ORM

Parameters required to connect to PostgreSQL on project are already available as the following environment variables:

- `POSTGRES_HOSTNAME`
- `POSTGRES_PORT`
- `POSTGRES_USERNAME`
- `POSTGRES_PASSWORD`

You can use Go ORMs like [go-pg/pq](https://github.com/go-pg/pg) and [jmoiron/sqlx](https://github.com/jmoiron/sqlx) to connect to Postgres.

Parameters required to connect to MySQL on project are already available as the following environment variables:

- `MYSQL_HOSTNAME`
- `MYSQL_PORT`
- `MYSQL_USERNAME`
- `MYSQL_PASSWORD`

You can use Go ORMs like [go-pg/pq](https://github.com/go-pg/pg) and [jmoiron/sqlx](https://github.com/jmoiron/sqlx) to connect to MySQL.
