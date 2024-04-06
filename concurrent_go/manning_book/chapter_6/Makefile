
bin/waitgroup_intro: waitgroup_intro/main.go
	go build -o bin/waitgroup_intro waitgroup_intro/main.go

bin/character_frequency: character_frequency/main.go
	go build -o bin/character_frequency character_frequency/main.go

bin/sp_wg_demo: sp_wg_demo/main.go
	go build -o bin/sp_wg_demo sp_wg_demo/main.go


.PHONY: all clean

all: bin/waitgroup_intro bin/character_frequency bin/sp_wg_demo

clean:
	rm -f bin/*
