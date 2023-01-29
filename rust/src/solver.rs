use std::collections::HashMap;

// Only used at init time, perf not important
fn digits_possible(s: [[u8; 9]; 9], i: usize, j: usize) -> Option<Vec<u8>> {
    if s[i][j] != 0 {
        // cell already set
        return None;
    }

    let mut digits = Vec::new();
    for n in 1..10 {
        if is_available(s, i, j, n) {
            digits.push(n);
        }
    }

    Some(digits)
}

// Only used at init time, perf not important
fn matrix_possibilities(s: [[u8; 9]; 9]) -> HashMap<String, Vec<u8>> {
    let mut tab: HashMap<String, Vec<u8>> = HashMap::new();

    for i in 0..9 {
        for j in 0..9 {
            match digits_possible(s, i, j) {
                Some(v) => {
                    if v.len() != 0 {
                        tab.insert(format!("{i}{j}"), v);
                    }
                }
                None => continue,
            }
        }
    }
    tab
}

#[derive(Debug, PartialEq, Eq, PartialOrd, Ord, Clone, Copy)]
struct Pos {
    i: usize,
    j: usize,
}

// Only used at init time, perf not important
fn tableau_order(s: [[u8; 9]; 9]) -> Vec<Pos> {
    #[derive(Debug, PartialEq, Eq, PartialOrd, Ord)]
    struct PosWithScores {
        pos: Pos,
        score: i32,
    }

    let possibilities = matrix_possibilities(s);

    let mut liste_scores = vec![];
    for i in 0..9 {
        for j in 0..9 {
            let key = format!("{i}{j}");
            let score = match possibilities.get(&key) {
                Some(v) => v.len() as i32,
                _ => 0,
            };
            if score == 0 {
                continue;
            }
            liste_scores.push(PosWithScores {
                pos: Pos { i, j },
                score,
            });
        }
    }

    liste_scores.sort_by(|a, b| a.score.cmp(&b.score));
    liste_scores.into_iter().map(|p| p.pos).collect()
}

// Used at each step, perf is important
fn is_available_in_line(s: [[u8; 9]; 9], i: usize, j: usize, n: u8) -> bool {
    for k in 0..9 {
        if (s[i][k] == n && k != j) || (s[k][j] == n && k != i) {
            return false;
        }
    }
    true
}

// Used at each step, perf is important
fn is_available_in_bloc(s: [[u8; 9]; 9], i: usize, j: usize, n: u8) -> bool {
    for (k, line) in s.iter().enumerate().skip((i / 3) * 3).take(3) {
        if k == i {
            continue;
        }
        for (l, case) in line.iter().enumerate().skip((j / 3) * 3).take(3) {
            if l != j && *case == n {
                return false;
            }
        }
    }

    true
}

// Used at each step, perf is important
fn is_available(s: [[u8; 9]; 9], i: usize, j: usize, n: u8) -> bool {
    is_available_in_line(s, i, j, n) && is_available_in_bloc(s, i, j, n)
}

pub struct Stats {
    tries: i32,
    going_back: i32,
}

// main function
pub fn solve(si: [[u8; 9]; 9]) -> ([[u8; 9]; 9], Stats) {
    let mut s = si;
    let possibilities = matrix_possibilities(s);
    let slice_order = tableau_order(s);
    let max_digit_to_find = slice_order.len();
    let mut index_of_current_digit_for: [[usize; 9]; 9] = [[0; 9]; 9];

    let mut rank: usize = 0;
    let mut stats = Stats {
        tries: 0,
        going_back: 0,
    };

    while rank < max_digit_to_find {
        stats.tries += 1;

        let n = slice_order[rank];
        let i = n.i;
        let j = n.j;
        let key = format!("{i}{j}");

        let possibility = possibilities.get(&key).unwrap();

        if index_of_current_digit_for[i][j] < possibility.len() {
            let client = possibility[index_of_current_digit_for[i][j]];

            if is_available(s, i, j, client) {
                s[i][j] = client;
                rank += 1;
            } else {
                index_of_current_digit_for[i][j] += 1;
            }
        } else {
            s[i][j] = 0;
            index_of_current_digit_for[i][j] = 0;
            rank -= 1;
            stats.going_back += 1;
            let n = slice_order[rank];
            let i = n.i;
            let j = n.j;
            index_of_current_digit_for[i][j] += 1;
        }
    }

    (s, stats)
}
