
bin/stingy_spendy: stingy_spendy/main.go
	go build -o bin/stingy_spendy stingy_spendy/main.go

bin/stingy_spendy2: stingy_spendy2/main.go
	go build -o bin/stingy_spendy2 stingy_spendy2/main.go

.PHONY: all clean

all: bin/stingy_spendy bin/stingy_spendy2

clean:
	rm -f bin/*

