# APIの使用例(エミュレーター環境)
## 1. ユーザー登録APIの使用例
### Windowsコマンドプロンプトでは以下のコマンドを例にuserNameを書き換えて実行。
```
curl -X POST http://localhost:8080/userRegister -H "Content-Type: application/json" -d "{\"name\":\"TestUser\"}"
```
## 2. ログインAPIの使用例
### Windowsコマンドプロンプトでは以下のコマンドを例にUserIDと引継ぎコードを書き換えて実行。
```
curl -X POST http://localhost:8080/login -H "Content-Type: application/json" -d "{\"userID\":6917529027641081856,\"transferCode\":\"edc953df-eb28-4874-aa28-0796619f901b\"}"
```
## 3. テスト認証APIの使用例(削除予定)
### Windowsコマンドプロンプトでは以下のコマンドを例にトークンを書き換えて実行。
```
curl -X GET http://localhost:8080/auth/hello -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDUxNTc4NDUsIm9yaWdfaWF0IjoxNzQ1MTU0MjQ1LCJ1c2VySUQiOjY5MTc1MjkwMjc2NDEwODE4NTZ9.bIxNqQjr6VePSjwdwAV5LYC-T8WzHRMSZsy6qcrB8KU"
```
