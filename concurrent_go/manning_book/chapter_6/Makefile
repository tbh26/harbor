
bin/waitgroup_intro: waitgroup_intro/main.go
	go build -o bin/waitgroup_intro waitgroup_intro/main.go

bin/character_frequency: character_frequency/main.go
	go build -o bin/character_frequency character_frequency/main.go

.PHONY: all clean

all: bin/waitgroup_intro bin/character_frequency

clean:
	rm -f bin/*

