.PHONY: dev
dev:
	wrangler dev

.PHONY: build
build:
	go run ../../cmd/workers-assets-gen
	#tinygo build -o ./build/app.wasm -target wasm -no-debug ./...
	tinygo build -o ./build/app.wasm -target wasm -no-debug ./...

.PHONY: deploy
deploy:
	wrangler deploy
