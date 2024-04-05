
bin/stingy_spendy: stingy_spendy/main.go
	go build -race -o bin/stingy_spendy stingy_spendy/main.go

bin/character_frequency: character_frequency/main.go
	go build -race -o bin/character_frequency character_frequency/main.go

bin/read_heavy_match_events: read_heavy_match_events/main.go
	go build -race -o bin/read_heavy_match_events read_heavy_match_events/main.go

.PHONY: all clean

all: bin/character_frequency bin/read_heavy_match_events bin/stingy_spendy

clean:
	rm -f bin/*

