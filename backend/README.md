## Backend Server

### ディレクトリ構成
```
/backend
|-- /config               // 設定ファイルの読み込みや環境変数の処理
|   |-- config.yml
|-- /internal
|   |-- /api              // API層, エンドポイントの定義
|   |   |-- /v1
|   |       |-- user.go
|   |       |-- room.go
|   |       |-- match.go
|   |-- /models           // スキーマを定義
|   |   |-- user.go
|   |   |-- room.go
|   |-- /services         // ビジネスロジックの実装, API層から呼び出される
|   |   |-- user.go
|   |   |-- match.go
|   |   |-- websocket.go
|   |-- /repositories     // データアクセス層, データベース関連の操作を集中管理する
|   |   |-- user.go
|   |   |-- room.go
|-- /middleware           // ミドルウェア層
|   |-- jwt.go
|-- /pkg
|   |-- /db               // データベースの初期化と設定
|   |   |-- postgres.go
|   |   |-- redis.go
|-- /routes
|   |-- routes.go
|-- /utils              // アプリケーション全体で使用されるユーティリティ関数やヘルパー関数
|   |-- jwt.go
|   |-- password.go
|   |-- response.go
|-- go.mod
|-- go.sum
|-- main.go

```