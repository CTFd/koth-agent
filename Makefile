all:
	make docs
	make linux
	make darwin
	make unix
	make windows

binaries:
	make linux
	make darwin
	make unix
	make windows

docs:
	cd src && swag init
	echo "$$(echo "<script>SWAGGER_JSON=$$(cat src/docs/swagger.json);</script>"; cat src/docs/index.orig.html)" > src/docs/index.html
	mkdir -p dist/docs
	cp src/docs/index.html dist/docs
	cp src/docs/swagger.* dist/docs

linux:
	GOOS=linux GOARCH=amd64 \
	go build -ldflags "${LDFLAGS}" -o dist/agent.linux src/main.go

darwin:
	GOOS=darwin GOARCH=amd64 \
	go build -ldflags "${LDFLAGS}" -o dist/agent.darwin src/main.go

unix:
	GOOS=freebsd GOARCH=amd64 \
	go build -ldflags "${LDFLAGS}" -o dist/agent.unix src/main.go

windows:
	GOOS=windows GOARCH=amd64 \
	go build -ldflags "${LDFLAGS}" -o dist/agent.windows.exe src/main.go

clean:
	rm -rf dist/*