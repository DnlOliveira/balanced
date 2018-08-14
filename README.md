
### Pre-req
`npm install`

### Dev 
`npm run dev`
or
`go run main.go`

### Deploy
```
npm run cleanup (if previously deployed)
npm run build
GOOS=linux GOARCH=amd64 go build
nohup ./balanced &
kill pid
```