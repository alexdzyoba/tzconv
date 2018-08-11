.PHONY: build gen
build:
	go get -v ./...
	go install

gen:
	./tzdump.py > tz.go

