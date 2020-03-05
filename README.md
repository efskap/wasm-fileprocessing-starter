# Go WebAssembly File Processing Starter

### [Demo](https://wasm-fileproc-demo.netlify.com/)

This is a starting point for an imo not-so-uncommon use of WebAssembly: having the user select a file, processing it, and then rendering a template based on the result, all inside the browser. For example, I'm working on [something like that](https://wc.dmitry.lol) for parsing Warcraft 3 replays.

In this case, we print the length of the file and the first line if the contents are a valid UTF-8 string.

It embeds the template inside the binary using [github.com/jerblack/statics](https://github.com/jerblack/statics) (or rather my 1-commit-ahead fork), and defers file processing until the WebAssembly program has initialized.

## Trying it out

Run `make` in the root directory to build, and serve the `public/` directory with something like:

```sh
# install goexec: go get -u github.com/shurcooL/goexec
$ goexec 'http.ListenAndServe(`:8080`, http.FileServer(http.Dir(`.`)))'
```
 
Make sure the server supports the  `application/wasm` Content-Type. Then you can deploy the folder on some static host/cdn like netlify.

Also take a look at this rather helpful overview: https://github.com/golang/go/wiki/WebAssembly
