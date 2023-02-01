# Compilation
compile-rust:
	cd rust && cargo build --release

compile-go:
	cd go && go build .

compile-all:
	make compile-rust
	make compile-go

# Benchmarks
bench-go:
	./go/sudoku-solver-go -silent ./data/sudoku_hard.txt

bench-rust:
	./rust/target/release/solver ./data/sudoku_hard.txt

bench-python:
	python3 ./python/main.py ./data/sudoku_hard.txt

bench-all: compile-all
	hyperfine --warmup 1 --export-markdown report.md --export-json report.json 'make bench-go' 'make bench-rust' 'make bench-python'

bench-all-except-python: compile-all
	hyperfine --warmup 1 --export-markdown report.md --export-json report.json 'make bench-go' 'make bench-rust'

bar-chart:
	python3 ./barchart.py --bins 100 -o bar_chart.png --type bar report.json || echo "Please \n- install python3, matplotlib, numpy \n- run a benchmark first with 'make bench-all'"
