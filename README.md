<img src="https://i.imgur.com/QYifFFz.png" alt="Screenshot" width="75%" />

# piv

[![Go Report Card](https://goreportcard.com/badge/github.com/hoffa/piv)](https://goreportcard.com/report/github.com/hoffa/piv)

ANSI terminal image viewer.

Only supports the 8 standard ANSI colors. No 24-bit colors, no Unicode tricks.

## Installation

```shell
go get github.com/hoffa/piv
```

## Usage

```shell
piv < image.png
```

piv reads image data is read from stdin. Output width can be specified with `-width`.
