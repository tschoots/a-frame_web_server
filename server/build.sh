rm -rf a-frame_web_server

CGO_ENABLED=0 go build -a --installsuffix cgo --ldflags="-s" -o a-frame_web_server
