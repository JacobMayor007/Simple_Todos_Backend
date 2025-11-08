To get started in Go

go mod init {name of your package}
go get github.com/gofiber/fiber/v2

I use fiber framework because of the similarities of it with Express.js
I love javascript, but sometimes you need to try other languages to grow.

PORT = 3201

```

ENDPOINT "/todos"

```

```
Routes:
(PATCH "/todos")
(DELETE "/todos")
(POST "/todos")
(GET "/todos")
```

You might ask where is the data, and how would you update because there is no data in the parameter?

```
For security reasons:

All the data, is in the body.

```

Ex:

```
GET Todos:
("/todos")

Response:
id
title
completed
createdAt

```

```
POST Todos:
("/todos")

Body:
id: 1,
title: "Example",
completed: false,
createdAt: "now" <--- ("This is for testing only, for production better to use time built in package of Go")

Response:
id
title
completed
createdAt

```

```
DELETE
("/todos")
Body:
id: 1,

Response:
id,
message
```

```
PUT
("/todos")
Body:
id: 1,
title: "Example",
completed: false,

Response:
id
title
completed
createdAt

```
