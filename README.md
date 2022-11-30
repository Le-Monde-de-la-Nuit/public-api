# Public API

Official Public API of [Le Monde de la Nuit](https://discord.gg/9KCJuEYUx2)

## Endpoints

### Members
- [x] /
- [x] /members
- [x] /member?id=:id
- [x] /roles
- [x] /role?id=:id
### Posts
- [x] /posts
- [x] /posts?id=:id
- [ ] /posts?last=:nbr
- [x] /tags
- [x] /tag?id=:id
### Actions
- [ ] /places
- [ ] /place?id=:id
- [ ] /actions
- [ ] /action?id=:id
- [ ] /types
- [ ] /type?id=:id

## Build

To build the project, you need to have Go 1.19 installed.

```bash
go build
```

### Docker

You can also build the project using Docker.

```bash
docker build -t le-monde-de-la-nuit/public-api .
```

or with docker-compose (just rename the file `compose.sample.yaml` into `compose.yaml`)

```bash
docker-compose build
```

### Makefile

The makefile build the project with docker-compose.

```bash
make build
```

## Run

To run the project, you need to run the binary file.

You need to give these args:
- Username of every database first
- Password of every database second

```bash
./public-api user password
```

### Docker

You can also run the project using Docker.

```bash
docker run -p 8080:80 -e USER="user" -e PASSWORD="password" le-monde-de-la-nuit/public-api
```
or with docker-compose (just rename the file `compose.sample.yaml` into `compose.yaml`)
```bash
docker-compose up
```

### Makefile

The makefile run the project with docker-compose.

```bash
make start
```


## Technologies

- Golang 1.19
- Gorilla/mux
