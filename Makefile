bench-go:
	cd go && go run . -silent ../data/sudoku_hardest.txt

bench-rust:
	cd rust && cargo run --release --quiet ../data/sudoku_hardest.txt

bench-python:
	cd python && python3 main.py ../data/sudoku_hardest.txt

bench-all:
	hyperfine --warmup 1 --export-markdown report.md 'make bench-go' 'make bench-rust' 'make bench-python'


