.PHONY: build
build:
	go build -o ./build/sblocker .

.PHONY: testgo-build
testgo-build:
	go run . -f -e ./example/go/example -t ./example/go/block -p block
.PHONY: testgo
testgo:
	go run ./example/go/example

.PHONY: testjs-pre
testjs-pre:
	cd ./example/js; yarn add file:/home/moyinzi/projects/gKit/sblock-project/libs/js/sblock4js; yarn add file:/home/moyinzi/projects/gKit/sblock-project/libs/js/sblock4koa
.PHONY: testjs-build
testjs-build:
	go run . -f -e ./example/js/src -t ./example/js/src -p block -l js
.PHONY: testjs
testjs:
	node ./example/js