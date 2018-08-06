# Template

Template is an web application template. It contains many things needed for
painless start when creating web application.

### Template contains:
 - [x] [logrus](https://godoc.org/github.com/sirupsen/logrus) logger
 - [x] PostgreSQL database
 - [x] [SQL migrations](https://github.com/golang-migrate/migrate)
 - [x] [TOML](https://github.com/BurntSushi/toml) configuration file
 - [x] [systemd socket activation](https://vincent.bernat.im/en/blog/2018-systemd-golang-socket-activation) - not tested yet
 - [x] [structs validator](https://github.com/kerak19/template/tree/master/internal/validate)
 - [x] Docker compose file with PostgreSQL database
 - [x] [higher](https://github.com/kerak19/template/tree/master/internal/model) and [lower](https://github.com/kerak19/template/tree/master/internal/repo/usersdb) level models acting as database interface
 - [x] [functions](https://github.com/kerak19/template/blob/master/internal/request/responses.go) providing nice(i guess) API request's response format 
 - [x] [Controller](https://github.com/kerak19/template/blob/master/internal/controller/controller.go) acting as main backend router using [httprouter](https://github.com/julienschmidt/httprouter) - router will probably change because of [this issue](https://github.com/julienschmidt/httprouter/issues/73), or at least until v2 is out
 - [x] [Basic handlers](https://github.com/kerak19/template/tree/master/internal/controller/users) for new users creation and login(login has output only for now)

### Todo:
  - [ ] middlewares:
    - [ ] authentication
    - [ ] authorization using ACL
    - [ ] logger for extended logging capabilities
  - [ ] real login
  - [ ] add logs into handlers 
  - [ ] add query debugger to models
  - [ ] whatever else comes to my mind

# Running template:
In the root of project:
```
    docker-compose up
```
and in other terminal window:
```
    go run main.go
```
#### Creating new migration
In the root of project:
```
    ./migrate name_of_migration
```