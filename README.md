# Project "Globstark"

This project uses GoLang, WebAssembly (WASM) and containers.

## Requirements

[golang](https://go.dev/doc/install)
[rustup](https://www.rust-lang.org/tools/install)
[wasm-pack](https://rustwasm.github.io/wasm-pack/installer/)

## Documentation

### Initialize the Go Module

```bash
cd globstark
go mod init github.com/antazo/globstark
```

### Build and Run the Containers

```bash
docker-compose up --build
```

<http://localhost:7979/>
