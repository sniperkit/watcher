NAME=watcher
build:
	@go build -o $(NAME) .

run: build
	@./$(NAME) -e development

run-prod: build
	@./$(NAME) -e development
