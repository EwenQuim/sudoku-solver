# Sudoku-solving based language benchmark

This is a benchmark of a Sudoku solver written in various languages. The goal is two-fold:
- to compare the performance of different languages
- it's my `Hello, world!` when I learn a new language


## Results

The results are for a 9x9 Sudoku, with 3x3 blocks. The benchmark is run on a 2022 MacBook Pro, with a M1 processor and 16GB of RAM.

### For an easy Sudoku

| Command             |    Mean [ms] | Min [ms] | Max [ms] |      Relative |
| :------------------ | -----------: | -------: | -------: | ------------: |
| `make bench-go`     |    7.7 ± 1.2 |      5.7 |     14.3 |   1.35 ± 0.28 |
| `make bench-rust`   |    5.7 ± 0.8 |      4.8 |     10.8 |          1.00 |
| `make bench-python` | 410.5 ± 12.2 |    392.0 |    429.8 | 72.12 ± 10.03 |

### For a hard Sudoku

![Graph displaying the table below](bar_chart_all.png)

![Graph displaying the table below, without Python that is too much of an outlier](bar_chart_except_python.png)


| Command             |     Mean [ms] | Min [ms] | Max [ms] |       Relative |
| :------------------ | ------------: | -------: | -------: | -------------: |
| `make bench-go`     |    32.2 ± 1.2 |     30.5 |     36.8 |    1.42 ± 0.11 |
| `make bench-rust`   |    22.6 ± 1.6 |     20.5 |     31.9 |           1.00 |
| `make bench-python` | 7162.1 ± 69.2 |   7059.1 |   7305.0 | 316.46 ± 22.27 |

### For a very hard Sudoku, made especially against backtracking

| Command             |         Mean [ms] | Min [ms] | Max [ms] |      Relative |
| :------------------ | ----------------: | -------: | -------: | ------------: |
| `make bench-go`     |       485.7 ± 8.5 |    474.7 |    498.9 |          1.00 |
| `make bench-rust`   |      1017.7 ± 7.6 |   1012.5 |   1035.6 |   2.10 ± 0.04 |
| `make bench-python` | 439984.2 ± 2345.9 | 436827.3 | 443363.3 | 646.87 ± 5.53 |

*Yes, the Python version takes 7 minutes to solve this Sudoku, while the other languages answer in a second.*

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
