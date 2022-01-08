DIR?=./build/_output
fmt:
	@go fmt ./...
vet:
	@go vet ./...
gen-cert:
	@openssl req -x509 -nodes -newkey rsa:2048 -keyout ${DIR}/server.key -out ${DIR}/server.crt -days 3650 -subj "/C=CN/ST=Shaanxi/L=Xian/O=Global Security/OU=IT Department/CN=*"

build: fmt vet
	@CGO_ENABLED=0 go build -v -ldflags "-w" -o ${DIR}/webhook cmd/main.go 
run: build gen-cert
	${DIR}/webhook --key=${DIR}/server.key --cert=${DIR}/server.crt
clean:
	@rm -rf ./build/_output

.PHONY: build vet clean
