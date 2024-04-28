
bin/intro: intro/main.go
	go build -o bin/intro intro/main.go

bin/channel_patterns: channel_patterns/main.go
	go build -o bin/channel_patterns channel_patterns/main.go

.PHONY: all clean

all: bin/intro bin/channel_patterns

clean:
	rm -f bin/*
