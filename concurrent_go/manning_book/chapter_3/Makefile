
bin/countdown: countdown/main.go
	go build -race -o bin/countdown countdown/main.go

.PHONY: clean compile_countdown

clean:
	rm -f bin/*

compile_countdown:
	go build -x -o bin/countdown -gcflags=-m countdown/main.go
