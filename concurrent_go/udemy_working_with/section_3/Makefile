
bin/race_demo: race_demo/main.go
	go build -race -o bin/race_demo race_demo/main.go
	@ /usr/bin/printf \
		"Notice, go build with %b-race%b option!  (execute with; $$ bin/race_demo ) \n" \
		"\033[31m" "\033[0m"

bin/fix_race_use_mutex: fix_race_use_mutex/main.go
	go build -race -o bin/fix_race_use_mutex fix_race_use_mutex/main.go
	@ /usr/bin/printf \
		"Notice, go build with %b-race%b option!  (execute with; $$ bin/fix_race_use_mutex ) \n" \
		"\033[31m" "\033[0m"

.PHONY: clean compile_race_demo test_race compile_fix

compile_race_demo: race_demo/main.go
	go build -o bin/race_demo race_demo/main.go

test_race: race_demo/main_test.go
	go test -race ./race_demo/

compile_fix: fix_race_use_mutex/main.go
	go build -o bin/fix_race_use_mutex fix_race_use_mutex/main.go

test_fix: fix_race_use_mutex/main_test.go
	go test -v -race ./fix_race_use_mutex/

clean:
	rm -f bin/*
