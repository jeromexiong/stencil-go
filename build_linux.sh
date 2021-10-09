go-bindata -pkg config_data -o app/bindata/config/config_data.go config/...
go-bindata -fs -pkg static_data -o app/bindata/static/static_data.go public/...

# go env -w GOOS=linux
go env -w GOOS=darwin
# go env -w GOOS=windows
go build -ldflags "-w -s" -o output/stencil-go
