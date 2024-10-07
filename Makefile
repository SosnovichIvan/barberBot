PKG=barberbot
BUILD_PATH='./build/barberbot'
BUILD_PATH_DARWIN='./build/barberbot-darwin'
BUILD_PATH_LINUX='./build/barberbot-linux'
RM_PATH='build'

build: 
	GOARCH=amd64 GOOS=darwin go build -o ${BUILD_PATH_DARWIN} ./cmd/main.go
	GOARCH=amd64 GOOS=linux go build -o ${BUILD_PATH_LINUX} ./cmd/main.go

run: build
	${BUILD_PATH_DARWIN}

clean:
	rm -r ${RM_PATH}