## Running Locally
### Requirements
Go 1.18+ is required because [embed](https://golang.org/pkg/embed/) is used for input files and generics are being used.

Use `go run main.go -part <1 or 2>` will be usable to run the actual inputs for that day.

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

Note that skeletons use [embed][embed] and __will not compile__ without an `input.txt` file located in the same folder. Input files can be made via `make input`.
```sh
make skeleton DAY=5 YEAR=2023
make input DAY=5 YEAR=2023 AOC_SESSION_COOKIE=your_cookie
```

### Fetch inputs and write to input.txt files
Requires passing your cookie from AOC from either `-cookie` flag, or `AOC_SESSION_COOKIE` env variable.
```sh
make input DAY=1 YEAR=2023
```

## Attribution
Most of the organization of this project as well as the utilities for fetching the daily inputs and generating each day's boilerplate was cribbed from [https://github.com/alexchao26/advent-of-code-go](https://github.com/alexchao26/advent-of-code-go.).