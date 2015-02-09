# Thanks for inspiration
# - https://gist.github.com/yanatan16/2951128
# - https://github.com/mackerelio/mackerel-agent/blob/master/Makefile

# Followings are ommited from 'deps' because of need of sudo 
# sudo gox -build-toolchain
# sudo gem install fpm

FPM = /usr/bin/fpm

SOURCE_LIST := gorandom.go 
TEST_LIST := tests
DEPENDS_LIST := 

ALL_LIST = $(SOURCE_LIST) $(TEST_LIST)

all: build test raspi edison

build: deps $(SOURCE_LIST)
	go fmt ./...
	golint ./...
	go build github.com/kgbu/gorandom/cmd/gorandom

raspi: deps build
	rm -rf ./packages
	mkdir -p ./packages/usr/local/gorandom/{bin,etc}
	gox -os="linux" -arch="arm" "github.com/kgbu/gorandom/cmd/gorandom"
	cp gorandom_linux_arm packages/usr/local/gorandom/bin/gorandom
	cd packages; $(FPM) -s dir -t deb -a armhf -n gorandom -v 0.0.`date +%s` .

edison: deps build
	gox -os="linux" -arch="386" "github.com/kgbu/gorandom/cmd/gorandom"
	echo 'gorandom_386 is made for Intel Edison'

test: deps $(ALL_LIST)
	go build
	go fmt
	go test -coverpkg github.com/kgbu/gorandom ./tests/*_test.go

deps:
	go get -d -v -t ./...
	go get -u github.com/kr/pty
	for dep in $(DEPENDS_LIST) ; do \
		go get -u $$dep ; \
		godep update $$dep ; done
	godep save
