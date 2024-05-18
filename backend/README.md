## Backend Server

### ディレクトリ構成
```
/backend
|-- /config
|   |-- config.yml
|-- /internal
|   |-- /api
|   |   |-- /v1
|   |       |-- user.go
|   |       |-- room.go
|   |       |-- match.go
|   |-- /models
|   |   |-- user.go
|   |   |-- room.go
|   |-- /services
|   |   |-- user.go
|   |   |-- match.go
|   |   |-- websocket.go
|   |-- /repositories
|   |   |-- user.go
|   |   |-- room.go
|-- /middleware
|   |-- jwt.go
|-- /pkg
|   |-- /db
|   |   |-- postgres.go
|   |   |-- redis.go
|-- /routes
|   |-- routes.go
|-- /utils
|   |-- jwt.go
|   |-- password.go
|   |-- response.go
|-- go.mod
|-- go.sum
|-- main.go

```