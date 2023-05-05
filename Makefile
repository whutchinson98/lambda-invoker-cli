build:
	go build -o lambda-invoker-cli

build-version:
	go build -o lambda-invoker-cli -ldflags="-X 'github.com/whutchinson98/lambda-invoker-cli/cmd.version=$(version)'" .
