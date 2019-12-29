# piv

[![Go Report Card](https://goreportcard.com/badge/github.com/hoffa/piv)](https://goreportcard.com/report/github.com/hoffa/piv)

ANSI terminal picture viewer.

Only supports the 8 standard ANSI colors. No 24-bit colors, no Unicode tricks.

## Installation

```shell
go get github.com/hoffa/piv
```

## Example

```shell
piv < image.png
```

Image data is read from stdin. Output width can optionally be specified with `-width`.
