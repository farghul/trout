# Trout

Trout is a WordPress plugin **release to production** tool. Named after a fish that often gets *released*.

![Trout](trout.webp)

## Prerequisites

Login information to download the update package. -- ***premium content only*** --

Googles' [Go language](https://go.dev) installed to enable building executables from source code.

## Build

From the root folder containing the `go` files, use the command that matches your environment:

### Windows & Mac:

``` console
go build -o [name] .
```

### Linux:

``` console
GOOS=linux GOARCH=amd64 go build -o [name] .
```

## Run

Ensure the folder containing your ***composer.json*** file is predefined as variable and run:

``` console
[program] [flag] [vendor/plugin]:[version] [ticket#]
```

## Example

### Release (In-house Production ready content):

``` console
trout -r bcgov-plugin/bcgov-inline-comments:1.9.0 820
```

Flags `-r` can accept multiple paired arguments, chain together as many as you like!

## License

Code is distributed under [The Unlicense](https://github.com/farghul/trout/blob/main/LICENSE.md) and is part of the Public Domain.
