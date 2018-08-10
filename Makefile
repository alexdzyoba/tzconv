.PHONY: build gen
build:
	go install

gen:
	./tzdump.py > tz.go

