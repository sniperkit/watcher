NAME=watcher
build:
	@go build -o $(NAME) .

run: build
	@./$(NAME) -e development

run-prod: build
	@./$(NAME) -e production

build-linux:
	export GOOS=linux GOARCH=amd64 
	@go build -v -o watcher_linux .
