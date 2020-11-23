# shit

`shit` (**Sh**eets **I**n **T**erminal) is a simple csv / xlsx reader
you can use from the comforts of your terminal.

> This is still a work in progress. Use it at your own discretion. There
> will be no official releases until I deem it as ready.

## Screenshot

![screenshot](./img/screenshot.jpg)

## Installation

I have plans of making `shit` available on Homebrew, as well as
publishing precompiled binaries for easier installation. However, for
now, please make sure you have `go` on your system before installing
`shit`.

### Automatic

- The fastest way is to use `go get github.com/chunkhang/shit`

### Manual

- Clone this repository into the right path: `$GOPATH/src/github.com/chunkhang/shit`
- Build with `go build`
- Move the built binary `shit` to somewhere like `/usr/local/bin`

## Usage

```
shit [file]
```

## FAQ

### How do I pronounce `shit`?

Pronounce `shit` like how you pronounce "sheet".

### Do I need `shit`?

Maybe you just prefer reading csv / xlsx files from the terminal. Or
perhaps like me, your licence for Microsoft Excel expired, and you still
miss it sometimes.

## Inspirations

- [scim](https://github.com/andmarti1424/sc-im)
- [wk](https://github.com/SheetJS/wk)
- [x_x](https://github.com/kristianperkins/x_x)
