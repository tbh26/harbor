
bin/select_intro: select_intro/main.go
	go build -o bin/select_intro select_intro/main.go

bin/another_select: another_select/main.go
	go build -o bin/another_select another_select/main.go

bin/character_frequency: character_frequency/main.go
	go build -o bin/character_frequency character_frequency/main.go

.PHONY: all clean

all: bin/select_intro bin/another_select bin/character_frequency

clean:
	rm -f bin/*

