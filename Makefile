get:
	go get -d -v ./...

build:
	go build main.go

populate_wallet:
	./main -mode populate -entity wallet

populate_asset:
	./main -mode populate -entity asset

populate_stock:
	./main -mode populate -entity stock

clean:
	rm main
