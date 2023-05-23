export LDFLAGS='-s -w '

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="$LDFLAGS" -trimpath -o scf_scanner_linux_amd64 main.go
CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -ldflags="$LDFLAGS" -trimpath -o scf_scanner_windows_386.exe  main.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags="$LDFLAGS" -trimpath -o scf_scanner_windows_amd64.exe  main.go
CGO_ENABLED=0 GOOS=windows GOARCH=arm64 go build -ldflags="$LDFLAGS" -trimpath -o scf_scanner_windows_arm64.exe  main.go
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags="$LDFLAGS" -trimpath -o scf_scanner_darwin_amd64 main.go
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags="$LDFLAGS" -trimpath -o scf_scanner_darwin_arm64 main.go

upx -9 scf_scanner_linux_amd64
upx -9 scf_scanner_windows_386.exe
upx -9 scf_scanner_windows_amd64.exe
upx -9 scf_scanner_windows_arm64.exe
upx -9 scf_scanner_darwin_amd64
upx -9 scf_scanner_darwin_arm64

zip scf_scanner_linux_amd64.zip scf_scanner_linux_amd64 config.yaml
zip scf_scanner_windows_386.zip scf_scanner_windows_386.exe config.yaml
zip scf_scanner_windows_amd64.zip scf_scanner_windows_amd64.exe config.yaml
zip scf_scanner_windows_arm64.zip scf_scanner_windows_arm64.exe config.yaml
zip scf_scanner_darwin_amd64.zip scf_scanner_darwin_amd64 config.yaml
zip scf_scanner_darwin_arm64.zip scf_scanner_darwin_arm64 config.yaml

rm -f scf_scanner_linux_amd64
rm -f scf_scanner_windows_386.exe
rm -f scf_scanner_windows_amd64.exe
rm -f scf_scanner_windows_arm64.exe
rm -f scf_scanner_darwin_amd64
rm -f scf_scanner_darwin_arm64