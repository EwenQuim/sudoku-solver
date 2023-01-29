mod file;
mod solver;

use solver::solve;
use std::time::Instant;

fn main() {
    // Get input from cli
    let args: Vec<String> = std::env::args().collect();
    let board = file::read_file(&args[1]);

    print_board(board);

    let now = Instant::now();

    let (solved, _stats) = solve(board);

    let elapsed = now.elapsed();
    println!("Elapsed: {elapsed:.2?}");

    print_board(solved);
}

fn print_board(board: [[u8; 9]; 9]) {
    for line in board {
        for cell in line {
            print!("{cell} ");
        }
        println!();
    }
}
