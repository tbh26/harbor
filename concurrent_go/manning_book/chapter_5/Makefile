
bin/stingy_spendy: stingy_spendy/main.go
	go build -o bin/stingy_spendy stingy_spendy/main.go

bin/stingy_spendy2: stingy_spendy2/main.go
	go build -o bin/stingy_spendy2 stingy_spendy2/main.go

bin/wait_signal_broadcast: wait_signal_broadcast/main.go
	go build -o bin/wait_signal_broadcast wait_signal_broadcast/main.go

bin/write_prefer_mutex: write_prefer_mutex/main.go
	go build -o bin/write_prefer_mutex write_prefer_mutex/main.go

bin/semaphore_demo: semaphore_demo/main.go
	go build -o bin/semaphore_demo semaphore_demo/main.go

.PHONY: all clean

all: bin/stingy_spendy bin/stingy_spendy2 bin/wait_signal_broadcast bin/write_prefer_mutex bin/semaphore_demo

clean:
	rm -f bin/*

