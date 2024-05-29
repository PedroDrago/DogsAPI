air:
	
run:
	go run ./cmd/api

build:
	go build -o /bin/api ./cmd/api

clean:
	rm -f ./bin/api

fclean: clean

re: fclean build
