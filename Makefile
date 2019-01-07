APPNAME=mock

export GOPATH=/Users/kevinchen/Documents/Golang/koalaone

.PHONY : build
build:
	@echo "GOPATH:"${GOPATH}
	go build -o ${APPNAME}

.PHONY : vgo
vgo:
	vgo install