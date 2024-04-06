
bin/channels_intro: channels_intro/main.go
	go build -o bin/channels_intro channels_intro/main.go

bin/channel_buffer_demo_slow_receiver: channel_buffer_demo_slow_receiver/main.go
	go build -o bin/channel_buffer_demo_slow_receiver channel_buffer_demo_slow_receiver/main.go

bin/channel_direction_demo: channel_direction_demo/main.go
	go build -o bin/channel_direction_demo channel_direction_demo/main.go

.PHONY: all clean

all: bin/channels_intro bin/channel_buffer_demo_slow_receiver bin/channel_direction_demo

clean:
	rm -f bin/*

