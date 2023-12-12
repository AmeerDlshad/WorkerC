.PHONY: dev
dev:
	wrangler dev

.PHONY: build
build:
	go run github.com/AmeerDlshad/WorkerC/cmd/workers-assets-gen@v0.18.0 -mode=go
	GOOS=js GOARCH=wasm go build -o ./build/app.wasm .

.PHONY: deploy
deploy:
	wrangler deploy
