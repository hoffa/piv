<img src="https://i.imgur.com/Jp9nccS.png" alt="Screenshot" width="75%" />

# piv

[![Go Report Card](https://goreportcard.com/badge/github.com/hoffa/piv)](https://goreportcard.com/report/github.com/hoffa/piv)

ANSI terminal image viewer.

Only supports the 8 standard ANSI colors. No 24-bit colors, no Unicode tricks.

## Installation

```shell
go get github.com/hoffa/piv
```

## Usage

```
Usage of piv:
  -width int
    	output image width (original size if 0) (default 80)
```

piv reads image data from stdin.

## Example

```shell
curl https://placekitten.com/600/600 | piv
```

Output width can be specified with `-width`. Setting it to `0` will display the image in its native resolution:

```shell
curl https://placekitten.com/600/600 | piv -width 0 | less -RS
```
