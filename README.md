# Sudoku-solving based language benchmark

This is a benchmark of a Sudoku solver written in various languages. The goal is two-fold:
- to compare the performance of different languages
- it's my `Hello, world!` when I learn a new language


## Results

The results are for a 9x9 Sudoku, with 3x3 blocks. The benchmark is run on a 2022 MacBook Pro, with a M1 processor and 16GB of RAM.

### For an easy Sudoku

| Command             |    Mean [ms] | Min [ms] | Max [ms] |    Relative |
| :------------------ | -----------: | -------: | -------: | ----------: |
| `make bench-go`     | 173.8 ± 11.2 |    164.1 |    204.5 | 1.79 ± 0.25 |
| `make bench-rust`   |  97.3 ± 12.1 |     89.9 |    136.2 |        1.00 |
| `make bench-python` | 439.3 ± 26.1 |    385.9 |    475.4 | 4.52 ± 0.62 |

### For a hard Sudoku

| Command             |     Mean [ms] | Min [ms] | Max [ms] |     Relative |
| :------------------ | ------------: | -------: | -------: | -----------: |
| `make bench-go`     |   212.1 ± 9.4 |    198.9 |    229.1 |         1.00 |
| `make bench-rust`   |   317.7 ± 7.4 |    302.6 |    326.6 |  1.50 ± 0.07 |
| `make bench-python` | 7177.4 ± 70.5 |   7103.8 |   7329.6 | 33.84 ± 1.53 |

### For a very hard Sudoku, made especially against backtracking

| Command             |         Mean [ms] | Min [ms] | Max [ms] |      Relative |
| :------------------ | ----------------: | -------: | -------: | ------------: |
| `make bench-go`     |       680.2 ± 4.5 |    671.4 |    687.2 |          1.00 |
| `make bench-rust`   |   12738.4 ± 104.5 |  12621.5 |  12956.7 |  18.73 ± 0.20 |
| `make bench-python` | 439984.2 ± 2345.9 | 436827.3 | 443363.3 | 646.87 ± 5.53 |

## Reproduce

Every benchmark must have a reproduce section, along with the usual License and Contributing sections. Not doing so is **wrong**. *Hi, [TurboPack](https://turbo.build/pack)*

### Prerequisites

- [Hyperfine](https://github.com/sharkdp/hyperfine)
- [Rust](https://www.rust-lang.org/tools/install)
- [Go](https://golang.org/doc/install)
- [Python](https://www.python.org/downloads/)
  - `pip install numpy rich matplotlib argcomplete`

### Run

```bash
make bench-all
```

## FAQ

As these are my `Hello World`s, the code is probably not idiomatic. I'm open to suggestions and PRs. 

> "The xxx implementation is not optimized for performance!"

This is your language's fault, not mine: I would be happy to code a faster algorithm, but if the language doesn't make it easy, I'm not going to do it.

> "But you include compilation time!"

Yes, I do. If your language is slow compiling, too bad. Not biased in favor of Go at all. Also, I don't want to lose my time compiling the same thing over and over again.
