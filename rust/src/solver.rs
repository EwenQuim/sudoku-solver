// Only used at init time, perf not important
fn digits_possible(s: [[u8; 9]; 9], i: usize, j: usize) -> Option<Vec<u8>> {
    if s[i][j] != 0 {
        // cell already set
        return None;
    }

    let mut digits = Vec::new();
    for n in 1..10 {
        if is_available(&s, i, j, n) {
            digits.push(n);
        }
    }

    Some(digits)
}

// Only used at init time, perf not important
fn matrix_possibilities(s: [[u8; 9]; 9]) -> [[Vec<u8>; 9]; 9] {
    let mut tab: [[Vec<u8>; 9]; 9] = Default::default();

    for (i, line) in tab.iter_mut().enumerate() {
        for (j, cell) in line.iter_mut().enumerate() {
            match digits_possible(s, i, j) {
                Some(v) => {
                    if !v.is_empty() {
                        *cell = v;
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
    for (i, line) in possibilities.iter().enumerate() {
        for (j, cell) in line.iter().enumerate() {
            let score =
                100 * cell.len() as i32 - aligned_neighbors(&s, i, j) - square_neighbors(&s, i, j);
            if score <= 0 {
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
fn is_available_in_line(s: &[[u8; 9]; 9], i: usize, j: usize, n: u8) -> bool {
    for k in 0..9 {
        if (s[i][k] == n && k != j) || (s[k][j] == n && k != i) {
            return false;
        }
    }
    true
}

// Used at each step, perf is important
fn is_available_in_bloc(s: &[[u8; 9]; 9], i: usize, j: usize, n: u8) -> bool {
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
fn is_available(s: &[[u8; 9]; 9], i: usize, j: usize, n: u8) -> bool {
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

        if index_of_current_digit_for[i][j] < possibilities[i][j].len() {
            let client = possibilities[i][j][index_of_current_digit_for[i][j]];

            if is_available(&s, i, j, client) {
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

fn aligned_neighbors(s: &[[u8; 9]; 9], i: usize, j: usize) -> i32 {
    let mut count = 0;
    for k in 0..9 {
        if j != k && s[i][k] != 0 {
            count += 1;
        }
        if i != k && s[k][j] != 0 {
            count += 1;
        }
    }

    count
}

fn square_neighbors(s: &[[u8; 9]; 9], i: usize, j: usize) -> i32 {
    let mut count = 0;
    let i_start = (i / 3) * 3;
    let j_start = (j / 3) * 3;
    for (k, _) in s.iter().enumerate().skip(i_start).take(3) {
        for l in j_start..j_start + 3 {
            if s[k][l] != 0 && (k != i || l != j) {
                count += 1;
            }
        }
    }

    count
}
