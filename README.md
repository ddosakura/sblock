# sblock - Static Files Block

Embed files into project (such as a Go executable project).

This package was inspired by the [statik](https://github.com/rakyll/statik) and the need like [#56](https://github.com/rakyll/statik/issues/56).

The package includes:
+ sblocker: the cli of code generater
+ sbi: the interface
+ plugins: the plugin to generate code
+ libs: the library to use the code block

> If you use `golang - origin/zip/default`, the code generated is compatible with [statik](https://github.com/rakyll/statik)

## Usage

Install the command line tool first.

```bash
go get -v -u github.com/ddosakura/sblock/sblocker
```

See the help

```bash
sblocker -h
```

## Official plugin and library (lang/algorithm)

+ golang
    + origin - Do not use compression to shrink the files.
    + zip/default
+ js
    + origin
    + zip/default

+ [The Catalog of Libs](./libs/README.md)

## How to implement Plugin

The plugin needs:
+ Compression algorithm by golang
+ The Code and Example pattern which will be generated

The library needs:
+ Decompression algorithm by your language
+ A FileSystem Abstraction System for you language like [afero](https://github.com/spf13/afero) to save your files

First, you should implement the interface:

```go
// github.com/ddosakura/sblock/sbi/interface.go
type Plugin interface {
    ...
}
```

Then you need to provide a library to use the code generated.

## TODO

+ [x] split library and plugin
+ [x] custom parameters for plugin
+ [ ] the description for custom parameters
+ [ ] add official nodejs/js plugin
    + [ ] sblock4js
    + [ ] sblock4koa
    + [ ] sblock-web-driver
        + custom param `no-dump`
    + [ ] print the explain to use yarn/npm
+ [ ] overwrite should not delete the target folder
    + [x] the `type Plugin interface` need `ForceOverwrite` param
+ [ ] the template to develop sblock-plugin
