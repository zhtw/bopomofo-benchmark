# Bopomofo Benchmark
[![Build Status](https://travis-ci.org/zhtw/bopomofo-benchmark.svg?branch=master)](https://travis-ci.org/zhtw/bopomofo-benchmark)

The purpose of this project is to create a benchmark system for bopomofo based
Traditional Chinese input method

# How to Contribute

First of all, please read [Code organization](http://golang.org/doc/code.html#Organization)
for golang. The layout of workspace shall look likes the following:

    $GOPATH
    ├── src
    │   └── github.com
            └── zhtw
                └── bopomofo-benchmark
                    ├── scripts
                    │   └── build.sh

The project is cloned in `$GOPATH/src/github.com/zhtw/bopomofo-benchmark`, while
`$GOPATH` is where go tool invoked. You can link helper script `scripts/build.sh`
to `$GOPATH/build.sh` so that invoking `$GOPATH/build.sh` will do `go fmt`, `go
test`, `go build` at the same time. You shall use `go fmt` before committing.

To run the benchmark, just run `bopomofo-benchmark` after building. The result
of benchmark will be stored in `report` if no `-reportdir=<path>` is provided.

# Data File Format

The following snippet is data file format used for benchmarking.

    ＃這是註解
    這是輸入|ㄓㄜˋㄕˋㄕㄨㄖㄨˋ

All strings after `#` are treated as comment and ignored. Each line of data file
contains a Chinese string and bopomofo string, separated by `|`. The bopomofo
string is the bopomofo input, and Chinese string is the expected output.

# Bug Report & Feature Request

Please report bug or feature request in [github issue](https://github.com/zhtw/bopomofo-benchmark/issues)

# License
The project is licensed under [MIT license](http://opensource.org/licenses/MIT).
