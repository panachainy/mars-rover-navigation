dev:
	wgo run ./src/main.go --grid_size 5 --obstacles "[(1,2),(3,3)]" --commands "LMLMLMLMM"

start:
	go run ./src/main.go --grid_size 5 --obstacles "[(1,2),(3,3)]" --commands "LMLMLMLMM"

setup:
	go install github.com/bokwoon95/wgo@latest

tidy:
	go mod tidy -v

t: test
test:
	# for clear cache `-count=1`
	@GIN_MODE=test go test -short ./src/...

it: integration.test
integration.test:
	@GIN_MODE=test go test ./tests/integration

# Run both unit and integration tests
test.all:
	@GIN_MODE=test make test
	@GIN_MODE=test make integration.test

# Run tests with verbose output
test.verbose:
	@GIN_MODE=test go test -v ./src/...

# Run integration tests with verbose output
integration.test.verbose:
	@GIN_MODE=test go test -v ./tests/integration

# Run all tests with verbose output
test.all.verbose:
	@GIN_MODE=test go test -v ./src/...
	@GIN_MODE=test go test -v ./tests/integration

# Test with race detection
test.race:
	@GIN_MODE=test go test -race ./src/...
	@GIN_MODE=test go test -race ./tests/integration

tr: test.report
test.report:
	go test -race -covermode=atomic -coverprofile=covprofile.out ./...
	make tc.html

tc: test.cov
test.cov:
	go test -race -covermode=atomic -coverprofile=covprofile.out ./...
	make test.cov.xml

# Generate XML coverage report
tc.xml: test.cov.xml
test.cov.xml:
	gocov convert covprofile.out > covprofile.xml

# Generate HTML coverage report
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
