<img src="https://i.imgur.com/TX8aCOa.png" alt="Screenshot" width="75%" />

# piv

[![Build Status](https://github.com/hoffa/piv/workflows/.github/workflows/workflow.yml/badge.svg)](https://github.com/hoffa/piv/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/hoffa/piv)](https://goreportcard.com/report/github.com/hoffa/piv)

Fast ANSI terminal image viewer.

- Supports PNG, GIF, JPEG, BMP, TIFF and WEBP formats.
- Only uses the 8 standard ANSI colors.
- No 24-bit colors or Unicode tricks.

## Installation

```shell
go get github.com/hoffa/piv
```

## Usage

```
Usage of piv:
  -r float
    	character width-to-height ratio (default 0.5)
  -w int
    	output image width (native width if 0) (default 80)

Image data is read from standard input.
```

## Example

```shell
curl https://placekitten.com/600/600 | piv
```

Setting `-width` to `0` will display the image in its native resolution:

```shell
curl https://placekitten.com/600/600 | piv -width 0 | less -RS
```
