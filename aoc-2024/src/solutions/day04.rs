use super::utils::{lines_from_file, try_step};
use std::collections::HashMap;

const FILEPATH: &str = "inputs/day04.txt";

/// Part 1
struct XmasSearch {
    start: char,              // The first letter of the searched word
    end: char,                // The final letter of the searched word
    seq: HashMap<char, char>, // Sequence for crossword search
}

impl XmasSearch {
    // Define each direction in which the word can be spelled
    pub const DIRECTIONS: [(isize, isize); 8] = [
        (-1, -1),
        (-1, 0),
        (-1, 1),
        (0, -1),
        (0, 1),
        (1, -1),
        (1, 0),
        (1, 1),
    ];
}

/// Part 2
struct XMASSearch {
    center: char,     // The middle character of an X
    wings: Vec<char>, // The two characters that comprise the wings of an X
}

impl XMASSearch {
    // Define each direction in the X pattern
    pub const TOP_LEFT: (isize, isize) = (-1, -1);
    pub const TOP_RIGHT: (isize, isize) = (-1, 1);
    pub const BOTTOM_LEFT: (isize, isize) = (1, -1);
    pub const BOTTOM_RIGHT: (isize, isize) = (1, 1);

    pub const DIRECTIONS: [(isize, isize); 4] = [
        XMASSearch::TOP_LEFT,
        XMASSearch::TOP_RIGHT,
        XMASSearch::BOTTOM_LEFT,
        XMASSearch::BOTTOM_RIGHT,
    ];
}

trait Search {
    fn num_matches_from_pt(&self, coords: (usize, usize), grid: &Vec<Vec<char>>) -> u32;
}

impl Search for XmasSearch {
    /// If a start character is detected at the provided coordinates in the grid, draw a line outward
    /// in every possible direction and check for the correct sequence of characters.
    fn num_matches_from_pt(&self, coords: (usize, usize), grid: &Vec<Vec<char>>) -> u32 {
        if grid[coords.0][coords.1] != self.start {
            return 0;
        }

        let mut num_matches: u32 = 0;
        for dir in XmasSearch::DIRECTIONS.into_iter() {
            let (mut i, mut j) = coords;
            loop {
                if grid[i][j] == self.end {
                    // We have found the desired word if the final character is reached
                    num_matches += 1;
                    break;
                } else if let Some((i_next, j_next)) = try_step((i, j), dir, grid) {
                    if grid[i_next][j_next] == self.seq[&grid[i][j]] {
                        // Update position if the character is the next in the sequence
                        (i, j) = (i_next, j_next);
                    } else {
                        break;
                    }
                } else {
                    break; // Step would be out of bounds
                }
            }
        }

        num_matches
    }
}

impl Search for XMASSearch {
    /// If a center character is detected at the provided coordinates in the grid, draw an X outward
    /// and check for the correct distribution of wing characters.
    fn num_matches_from_pt(&self, coords: (usize, usize), grid: &Vec<Vec<char>>) -> u32 {
        let (i, j) = coords;
        if grid[i][j] != self.center {
            return 0;
        }

        XMASSearch::DIRECTIONS.into_iter().all(|dir| {
            if let Some((i_next, j_next)) = try_step(coords, dir, grid) {
                if !self.wings.contains(&grid[i_next][j_next]) {
                    return false;
                } else {
                    // Ensure the opposite wing character is not equal to that of the current wing
                    match dir {
                        XMASSearch::TOP_LEFT => {
                            if i + 1 >= grid.len() || j + 1 >= grid[0].len() {
                                false
                            } else {
                                // TOP_LEFT should not be equal to BOTTOM_RIGHT
                                grid[i_next][j_next] != grid[i + 1][j + 1]
                            }
                        }
                        XMASSearch::TOP_RIGHT => {
                            if i + 1 >= grid.len() {
                                false
                            } else {
                                // TOP_RIGHT should not be equal to BOTTOM_LEFT
                                grid[i_next][j_next] != grid[i + 1][j - 1]
                            }
                        }
                        // The TOP_LEFT and TOP_RIGHT match arms already compare against the
                        // BOTTOM_LEFT and BOTTOM_RIGHT characters, so we default to true here
                        _ => true,
                    }
                }
            } else {
                return false; // Step would be out of bounds
            }
        }) as u32
    }
}

fn num_matches_in_grid(match_fn: impl Fn((usize, usize), &Vec<Vec<char>>) -> u32) -> u32 {
    let grid = get_grid();

    // Take the cross product of the grid index ranges
    let grid_range: Vec<(usize, usize)> = (0..grid.len())
        .flat_map(|i| (0..grid[0].len()).map(move |j| (i, j)))
        .collect();

    // Sum matches over all grid indices
    grid_range
        .into_iter()
        .map(|coords| match_fn(coords, &grid))
        .sum()
}

fn get_grid() -> Vec<Vec<char>> {
    lines_from_file(FILEPATH)
        .expect(&format!("Input file {FILEPATH} should exist"))
        .into_iter()
        .map(|line| line.chars().collect())
        .collect()
}

pub fn solve_part_1() {
    let search = XmasSearch {
        start: 'X',
        end: 'S',
        seq: HashMap::from([('X', 'M'), ('M', 'A'), ('A', 'S')]),
    };

    let num_matches = num_matches_in_grid(|coords, grid| search.num_matches_from_pt(coords, grid));
    println!("Number of times XMAS appears: {num_matches}")
}

pub fn solve_part_2() {
    let search = XMASSearch {
        center: 'A',
        wings: vec!['M', 'S'],
    };

    let num_matches = num_matches_in_grid(|coords, grid| search.num_matches_from_pt(coords, grid));
    println!("Number of times X-MAS appears: {num_matches}")
}
