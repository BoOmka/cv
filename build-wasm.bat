go generate
set GOOS=js
set GOARCH=wasm
go build -o dist\main.wasm .\client