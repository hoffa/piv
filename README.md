<img src="https://i.imgur.com/Jp9nccS.png" alt="Screenshot" width="75%" />

# piv

[![Build Status](https://travis-ci.com/hoffa/piv.svg?branch=master)](https://travis-ci.com/hoffa/piv)
[![Go Report Card](https://goreportcard.com/badge/github.com/hoffa/piv)](https://goreportcard.com/report/github.com/hoffa/piv)

Fast ANSI terminal image viewer.

Only supports the 8 standard ANSI colors. No 24-bit colors, no Unicode tricks.

## Installation

```shell
go get github.com/hoffa/piv
```

## Usage

Image data is read from stdin.

```
Usage of piv:
  -ratio float
        character width-to-height ratio (default 0.5)
  -width int
        output image width (native width if 0) (default 80)
```

## Example

```shell
curl https://placekitten.com/600/600 | piv
```

Setting `-width` to `0` will display the image in its native resolution:

```shell
curl https://placekitten.com/600/600 | piv -width 0 | less -RS
```
