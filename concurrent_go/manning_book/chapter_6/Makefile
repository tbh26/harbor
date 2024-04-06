
bin/waitgroup_intro: waitgroup_intro/main.go
	go build -o bin/waitgroup_intro waitgroup_intro/main.go

bin/character_frequency: character_frequency/main.go
	go build -o bin/character_frequency character_frequency/main.go

bin/sp_wg_demo: sp_wg_demo/main.go
	go build -o bin/sp_wg_demo sp_wg_demo/main.go

bin/flexible_waitgroup: flexible_waitgroup/main.go
	go build -o bin/flexible_waitgroup flexible_waitgroup/main.go

bin/barrier_intro: barrier_intro/main.go
	go build -o bin/barrier_intro barrier_intro/main.go

.PHONY: all clean

all: bin/waitgroup_intro bin/character_frequency bin/sp_wg_demo bin/flexible_waitgroup bin/barrier_intro

clean:
	rm -f bin/*
