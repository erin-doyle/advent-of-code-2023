# https://gist.github.com/prwhite/8168133
help: ## Show this help
	@ echo 'Usage: make <target>'
	@ echo
	@ echo 'Available targets:'
	@ grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'

# check-aoc-cookie:  ## ensures $AOC_SESSION_COOKIE env var is set
# 	@ test $${AOC_SESSION_COOKIE?env var not set}

skeleton: ## make skeleton main(_test).go files, required: $DAY and $YEAR
	go run cmd/skeleton/main.go -day $(DAY) -year $(YEAR);

input: ## get input, required: $DAY and $YEAR, optional: $COOKIE
	@ if [[ -n $$COOKIE ]]; then \
		go run cmd/input/main.go -day $(DAY) -year $(YEAR) -cookie $(COOKIE); \
	else \
		go run cmd/input/main.go -day $(DAY) -year $(YEAR); \
	fi

prompt: ## get prompt, required: $DAY and $YEAR, optional: $COOKIE
	@ if [[ -n $$COOKIE ]]; then \
		go run cmd/prompt/main.go -day $(DAY) -year $(YEAR) -cookie $(COOKIE); \
	else \
		go run cmd/prompt/main.go -day $(DAY) -year $(YEAR); \
	fi
