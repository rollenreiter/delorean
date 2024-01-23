build:
	go build -o delorean .

run:
	go build -o delorean .
	./delorean

clean:
	go clean
	rm delorean

install:
	go install .
