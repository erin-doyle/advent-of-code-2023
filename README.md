## Running Locally

Use `go run main.go -part <1 or 2>` to run the actual inputs for that day.

Use `go test -run RegExpToMatchFunctionNames .` to run examples and unit tests via the `main_test.go` files.

## Commands

Makefile should be fairly self-documenting. Alternatively you can run the binaries yourself via `go run` or `go build`.

`make help` prints a help message.

### Make skeleton files

```sh
for ((i=1; i<26; i++)); do
make skeleton DAY=$i YEAR=2023
done
```

Note that skeletons use [embed][embed] and __will not compile__ without an `input.txt` file located in the same folder. Input files can be made via `make input` (see details below).

```sh
make skeleton DAY=5 YEAR=2023
```

### Fetch inputs and write to input.txt files

You can fetch the day's input from AOC using the `make input` command with the `DAY` and `YEAR`.
To fetch the input for each day you will need your AOC session cookie.  You can retrieve this from your browser where you have logged in to AOC.  This can then be provided in 1 of 3 ways:

1. Create an `.env` file and add the variable: `AOC_SESSION_COOKIE` with your cookie as the value.

```sh
make input DAY=1 YEAR=2023
```

2. Set an `AOC_SESSION_COOKIE` env variable:

```sh
export AOC_SESSION_COOKIE=your_cookie
make input DAY=5 YEAR=2023
```

3. Pass the cookie using the `COOKIE` flag:

```sh
make input DAY=5 YEAR=2023 COOKIE=your_cookie
```

Precedence will be given in the following order:

1. `COOKIE` flag
2. `.env` `AOC_SESSION_COOKIE` variable
3. `AOC_SESSION_COOKIE` environment variable

## Attribution

Most of the organization of this project as well as the utilities for fetching the daily inputs and generating each day's boilerplate was cribbed from [https://github.com/alexchao26/advent-of-code-go](https://github.com/alexchao26/advent-of-code-go.).

The solutions and any modifications made after the first commit are 100% my own.
