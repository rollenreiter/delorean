NAME=delorean

build:
	go build -o delorean main.go

run:
	go build -o delorean main.go
	./delorean

clean:
	go clean
	rm delorean
