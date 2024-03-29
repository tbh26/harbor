
bin/race_demo: race_demo/main.go
	go build -race -o bin/race_demo race_demo/main.go
	@ /usr/bin/printf \
		"Notice, go build with %b-race%b option!  (execute with; $$ bin/race_demo ) \n" \
		"\033[31m" "\033[0m"

.PHONY: clean compile_race_demo

compile_race_demo: race_demo/main.go
	go build -o bin/race_demo race_demo/main.go

clean:
	rm -f bin/*
