use std::io::{BufRead, BufReader};

// Open file by name, read the content and put it in a  [[u8; 9]; 9] array
pub fn read_file(filename: &str) -> [[u8; 9]; 9] {
    let mut board: [[u8; 9]; 9] = [[0; 9]; 9];

    let mut j = 0;

    let file = std::fs::File::open(filename).expect("file not found");
    let reader = BufReader::new(file);

    for (i, line) in reader.lines().enumerate() {
        let line = line.unwrap();
        for c in line.chars() {
            if c != '.' {
                match c.to_digit(10) {
                    Some(n) => board[i][j] = n as u8,
                    None => panic!("Invalid character {c} in file"),
                }
            }
            j += 1;
        }
        j = 0;
    }

    board
}
