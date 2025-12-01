use crate::solutions;
use std::{
    fs::File,
    io::{self, BufRead, BufReader},
    path::Path,
};

pub fn get_solver_fn(day: u32, part: u32) -> Result<fn(), &'static str> {
    match (day, part) {
        (1, 1) => Ok(solutions::day01::solve_part_1),
        (1, 2) => Ok(solutions::day01::solve_part_2),
        (2, 1) => Ok(solutions::day02::solve_part_1),
        (2, 2) => Ok(solutions::day02::solve_part_2),
        (3, 1) => Ok(solutions::day03::solve_part_1),
        (3, 2) => Ok(solutions::day03::solve_part_2),
        (4, 1) => Ok(solutions::day04::solve_part_1),
        (4, 2) => Ok(solutions::day04::solve_part_2),
        (5, 1) => Ok(solutions::day05::solve_part_1),
        (5, 2) => Ok(solutions::day05::solve_part_2),
        (6, 1) => Ok(solutions::day06::solve_part_1),
        (6, 2) => Ok(solutions::day06::solve_part_2),
        (7, 1) => Ok(solutions::day07::solve_part_1),
        (7, 2) => Ok(solutions::day07::solve_part_2),
        (8, 1) => Ok(solutions::day08::solve_part_1),
        (8, 2) => Ok(solutions::day08::solve_part_2),
        (9, 1) => Ok(solutions::day09::solve_part_1),
        (9, 2) => Ok(solutions::day09::solve_part_2),
        (10, 1) => Ok(solutions::day10::solve_part_1),
        (10, 2) => Ok(solutions::day10::solve_part_2),
        (11, 1) => Ok(solutions::day11::solve_part_1),
        (11, 2) => Ok(solutions::day11::solve_part_2),
        _ => todo!("no solver for day {day} part {part}"),
    }
}

pub fn lines_from_file(filename: impl AsRef<Path>) -> io::Result<Vec<String>> {
    BufReader::new(File::open(filename)?).lines().collect()
}

/// Return the coordinates of the hypothetical result of taking the given step from
/// the provided starting point. Return None if the step would be out of bounds.
pub fn try_step<T>(
    start: (usize, usize),
    step: (isize, isize),
    grid: &Vec<Vec<T>>,
) -> Option<(usize, usize)> {
    match (
        TryInto::<usize>::try_into(start.0 as isize + step.0),
        TryInto::<usize>::try_into(start.1 as isize + step.1),
    ) {
        (Ok(i), Ok(j)) => {
            if i < grid.len() && j < grid[0].len() {
                Some((i, j))
            } else {
                None // At least one index is too high, out of bounds
            }
        }
        _ => None, // At least one index is negative, out of bounds
    }
}
