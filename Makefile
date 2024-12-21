TAG=
PACKAGE=github.com/Mabernetes/nc

all: run

run:
	go run $(PACKAGE)

build.linux:
	GOOS=linux GOARCH=amd64 go build -o m8s-nc-$(TAG) $(PACKAGE)

build.windows:
	GOOS=windows GOARCH=amd64 go build -o m8s-nc-$(TAG).exe $(PACKAGE)

build.all:
	build.linux
	build.windows

clean:
	rm m8s-nc-$(TAG).exe
	rm m8s-nc-$(TAG)
