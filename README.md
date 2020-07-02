# My first Golang App (HTTP server)

## Availables routes

* `GET /` redirect to `/tasks`
* `GET /tasks` get all tasks
* `GET /tasks/{id}` get one task
* `POST /tasks` create new task
* `PUT /tasks/{id}` update task
* `DELETE /tasks/{id}` delete task

## Interfaces

* ### full task
  * `id` integer
  * `body` string
  * `complete` boolean

* ### create/update task
  * `body` string | required
  * `complete` boolean | optional
