go build -ldflags "-s -w"
upx -9  -o jsc.exe js.exe