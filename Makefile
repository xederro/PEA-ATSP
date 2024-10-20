SHELL=cmd.exe

manual: prod
	.\cmd.exe

auto: prod run plot

build:
	go build .\cmd

run:
	chcp 65001 & .\cmd.exe -cpu -bf -bab -rep 10 -con 4 5 6 7 8 > data.csv

prod:
	go build -ldflags "-s -w" .\cmd

plot:
	python .\scripts\plottime.py .\data.csv plots

prof:
	go tool pprof -http 127.0.0.1:8080 cpu_profile.prof