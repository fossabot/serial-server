bin/serial-server: cmd/server/main.go
	go build -o $@ $^

clean:
	rm -f bin/serial-server
