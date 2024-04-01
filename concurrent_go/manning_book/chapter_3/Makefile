
bin/countdown: countdown/main.go
	go build -race -o bin/countdown countdown/main.go

bin/character_frequency: character_frequency/main.go
	go build -race -o bin/character_frequency character_frequency/main.go

bin/stingy_spendy: stingy_spendy/main.go
	go build -race -o bin/stingy_spendy stingy_spendy/main.go

.PHONY: clean compile_countdown

clean:
	rm -f bin/*

compile_countdown:
	go build -x -o bin/countdown -gcflags=-m countdown/main.go
