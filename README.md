Requirements:

- Go 1.13^
- Discord Bot credentials obtained from Discord Application Developer API

Installation:

- Create config.json file inside with following content these all are obtained through Discord.

```
{
  "ClientID": BOT_CLIENT_ID,
  "Secret": BOT_SECRET,
  "Permissions": 8, 
  "Token": BOT_TOKEN
}
```

- Inside the folder, `go run main.go` or `go build` then run the compiled executible