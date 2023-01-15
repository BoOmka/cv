# CV

Personal curriculum vitae site built using [Vugu](https://github.com/vugu/vugu) library.

# Run
```shell
go generate
set GOOS=js&&set GOARCH=wasm&&go build -o dist\main.wasm .\client
vgrun -1 ./server/server.go -build
vgrun ./server/static_server.go
```