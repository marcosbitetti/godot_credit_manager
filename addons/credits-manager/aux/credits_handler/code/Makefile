default: cross_compile

cross_compile:
	GOOS=linux GOARCH=amd64 go build -o ../bin/credits-manager-amd64-linux main.go
	upx -9 ../bin/credits-manager-amd64-linux
	#GOOS=linux GOARCH=386 go build -o ../bin/credits-manager-386-linux main.go
	#upx -9 ../bin/credits-manager-386-linux
	GOOS=windows GOARCH=amd64 go build -o ../bin/credits-manager-amd64.exe main.go
	upx -9 ../bin/credits-manager-amd64.exe
	#GOOS=windows GOARCH=386 go build -o ../bin/credits-manager-386.exe main.go
	#upx -9 ../bin/credits-manager-386.exe
	GOOS=darwin GOARCH=amd64 go build -o ../bin/credits-manager-amd64-darwin main.go
	upx -9 ../bin/credits-manager-amd64-darwin
	

