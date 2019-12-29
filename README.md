# piv

[![Go Report Card](https://goreportcard.com/badge/github.com/hoffa/piv)](https://goreportcard.com/report/github.com/hoffa/piv)

ANSI terminal picture viewer.

Only supports the 8 standard ANSI colors. No 24-bit colors, no Unicode tricks.

## Installation

```shell
go get github.com/hoffa/piv
```

## Usage

```shell                                                              
% piv -h
Usage of piv:
  -width int
    	output image width (default 80)
```

## Example

```shell
piv < image.png
```
