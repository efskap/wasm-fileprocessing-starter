OUTDIR='public'

all:
	go generate
	GOOS=js GOARCH=wasm go build -o $(OUTDIR)/main.wasm .
	cp "$(shell go env GOROOT)/misc/wasm/wasm_exec.js" $(OUTDIR)
	cp -r static/* $(OUTDIR)