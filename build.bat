go build -ldflags "-s -w"
upx -9  -o js.exe gojs.exe