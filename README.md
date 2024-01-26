## Go(Gin) Todo REST API

ユーザー登録機能とタスク管理機能を実装した簡単なREST APIです。

## 認証方式

このAPIは認証にJSON Web Token (JWT)を使用します。リクエストを認証するには、HTTPリクエストのAuthorizationヘッダーにJWTトークンを含めます。

## 仕様

### ユーザー登録

| メソッド | endpoint | 仕様 |
| --- | --- | --- |
| POST | /register | ユーザー新規登録 |
| POST | /login | ログイン |
| POST | /logout | ログアウト |

### タスク管理

| メソッド | endpoint | 仕様 |
| --- | --- | --- |
| GET | /tasks | タスクの一覧取得 |
| POST | /tasks | タスクの新規登録 |
| DELETE | /tasks/{id} | 指定IDのタスクの削除 |
| PUT | /tasks/{id} | 指定IDのタスクの更新 |
