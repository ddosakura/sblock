build:
	go build -o ./build/sblocker .

testgo-build:
	go run . -f -e ./example/go/example -t ./example/go/block -p block
testgo:
	go run ./example/go/example

testjs-pre:
	cd ./example/js; yarn add file:/home/moyinzi/projects/gKit/sblock-project/libs/js/sblock4js; yarn add file:/home/moyinzi/projects/gKit/sblock-project/libs/js/sblock4koa
testjs-build:
	go run . -f -e ./example/js/src -t ./example/js/src -p block -l js
testjs:
	node ./example/js