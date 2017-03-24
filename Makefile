
native:
	go build

linux:
	GOOS=linux GOARCH=amd64 go build -o myna_linux

win:
	GOOS=windows GOARCH=amd64 go build -o myna.exe

osx:
	GOOS=darwin GOARCH=386 go build -o myna

clean:
	rm -rf myna myna.exe

get-deps:
	go get -u -v github.com/urfave/cli
	go get -u -v github.com/howeyc/gopass
	go get -u -v github.com/ebfe/scard
	go get -u -v github.com/ianmcmahon/encoding_ssh
	go get -u -v github.com/fullsailor/pkcs7
