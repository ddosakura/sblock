# sblock - Static Files Block

Embed files into project (such as a Go executable project).

This package was inspired by the [statik](https://github.com/rakyll/statik) and the need like [#56](https://github.com/rakyll/statik/issues/56).

The package include:
+ sblocker: the cli of code generater
+ libs: the plugin to generate code & the library to use the code block

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

## TODO

+ [ ] custom parameters for plugin
