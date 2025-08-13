dev:
	wgo run ./src/main.go

start:
	go run ./src/main.go

setup:
	go install go.uber.org/mock/mockgen@latest
	go install github.com/axw/gocov/gocov@latest
	go install github.com/bokwoon95/wgo@latest
	go install golang.org/x/tools/gopls@latest
	make auth.newkey
	brew install golang-migrate

tidy:
	go mod tidy -v

t: test
test:
	# for clear cache `-count=1`
	@GIN_MODE=test go test -short ./...

it: integration.test
integration.test:
	@GIN_MODE=test go test ./...

# Run both unit and integration tests
test.all:
	@GIN_MODE=test go test -short ./...
	@GIN_MODE=test go test ./...

# Run tests with verbose output
test.verbose:
	@GIN_MODE=test go test -v ./...

# Run integration tests with verbose output
integration.test.verbose:
	@GIN_MODE=test go test -tags=integration -v ./...

# Run all tests with verbose output
test.all.verbose:
	@GIN_MODE=test go test -v ./...
	@GIN_MODE=test go test -tags=integration -v ./...

# Test with race detection
test.race:
	@GIN_MODE=test go test -race ./...
	@GIN_MODE=test go test -race -tags=integration ./...

tr: test.html
test.html:
	go test -race -covermode=atomic -coverprofile=covprofile.out ./...
	make tc.html

tc: test.cov
test.cov:
	go test -race -covermode=atomic -coverprofile=covprofile.out ./...
	make test.cov.xml

tc.xml: test.cov.xml
test.cov.xml:
	gocov convert covprofile.out > covprofile.xml

tc.html: test.cov.html
test.cov.html:
	go tool cover -html=covprofile.out -o covprofile.html
	open covprofile.html

c: clean
clean:
	rm -f covprofile.out covprofile.xml covprofile.html
	rm -rf tmp

f: fmt
fmt:
	go fmt ./...

g: generate
generate:
	go generate ./...

b: build
build:
	go build -o apiserver ./api/cmd
