use super::utils::{lines_from_file, try_step};
use std::collections::{HashMap, HashSet};

const FILEPATH: &str = "inputs/day06.txt";

/// Compute the set of distinct coordinates at which a single obstruction can be added
/// to induce a loop in the path of the guard, given the starting position and direction.
fn obstruction_positions(
    start_pos: (usize, usize),
    start_step: (isize, isize),
    grid: &Vec<Vec<char>>,
) -> HashSet<(usize, usize)> {
    let mut obstructions = HashSet::new();
    let mut visited = HashMap::from([(start_pos, HashSet::from([start_step]))]);
    let mut pos = start_pos;
    let mut step = start_step;

    while let Some((i_next, j_next)) = try_step(pos, step, grid) {
        // Simulate obstacle in front of guard, then trace path in search of a loop
        let new_grid = sim_obstacle_in_front(pos, step, &visited, grid);
        let new_path = walk_path(pos, step, &new_grid);
        if has_loop(&new_path, &new_grid) {
            obstructions.insert(try_step(pos, step, &new_grid).unwrap());
        }

        // Continue traversing the original path
        if grid[i_next][j_next] == '#' {
            step = (step.1, -step.0); // Take a 90-degree clockwise turn at obstacle
        } else {
            pos = (i_next, j_next); // Continue traveling in the same direction otherwise
        }
        visited.entry(pos).or_insert_with(HashSet::new).insert(step);
    }

    obstructions
}

/// Return a copy of the provided grid, but with an obstacle inserted in front of the given position
/// according to the step direction. If the proposed obstacle location is at the original starting point,
/// within the path taken to reach the current location, or out of bounds, do not insert an obstacle.
fn sim_obstacle_in_front(
    pos: (usize, usize),
    step: (isize, isize),
    visited: &HashMap<(usize, usize), HashSet<(isize, isize)>>,
    grid: &Vec<Vec<char>>,
) -> Vec<Vec<char>> {
    let mut new_grid = grid.clone();
    if let Some((i_front, j_front)) = try_step(pos, step, grid) {
        if !visited.contains_key(&(i_front, j_front)) && grid[i_front][j_front] != '^' {
            new_grid[i_front][j_front] = '#';
        }
    }
    new_grid
}

/// Return whether the provided set of visited positions forms a loop, as indicated
/// by the presence of the designated "loop key" of (grid.len(), grid[0].len())
fn has_loop(
    path: &HashMap<(usize, usize), HashSet<(isize, isize)>>,
    grid: &Vec<Vec<char>>,
) -> bool {
    path.contains_key(&(grid.len(), grid[0].len()))
}

/// Return the set of all grid coordinates visited by the guard, starting at the given position
/// and direction, mapped to the guard's direction(s) of travel while visiting each coordinate.
/// This assumes that the guard takes a right turn each time she encounters an obstacle.
///
/// For bookkeeping purposes, if a loop is encountered, the returned set will have a designated
/// "loop key" of (grid.len(), grid[0].len()) mapped to an empty set.
fn walk_path(
    start_pos: (usize, usize),
    start_step: (isize, isize),
    grid: &Vec<Vec<char>>,
) -> HashMap<(usize, usize), HashSet<(isize, isize)>> {
    let mut visited = HashMap::from([(start_pos, HashSet::from([start_step]))]);
    let mut pos = start_pos;
    let mut step = start_step;

    while let Some((i_next, j_next)) = try_step(pos, step, grid) {
        if grid[i_next][j_next] == '#' {
            step = (step.1, -step.0); // Take a 90-degree clockwise turn at obstacle
        } else {
            pos = (i_next, j_next); // Continue traveling in the same direction otherwise

            if visited.contains_key(&pos) && visited[&pos].contains(&step) {
                // Position already visited in the same orientation, loop detected...
                // Insert special loop key and exit
                visited.insert((grid.len(), grid[0].len()), HashSet::default());
                break;
            }
        }

        // Create or update position entry with new orientation
        visited.entry(pos).or_insert_with(HashSet::new).insert(step);
    }

    visited
}

/// Return the grid coordinates of the starting position, at which point
/// the guard is facing up. Error if no starting point is found.
fn find_start_pt(grid: &Vec<Vec<char>>) -> Result<(usize, usize), &str> {
    for i in 0..grid.len() {
        for j in 0..grid[0].len() {
            if grid[i][j] == '^' {
                return Ok((i, j));
            }
        }
    }
    Err("Starting point not found in grid")
}

fn get_grid() -> Vec<Vec<char>> {
    lines_from_file(FILEPATH)
        .expect(&format!("Input file {FILEPATH} should exist"))
        .into_iter()
        .map(|line| line.chars().collect())
        .collect()
}

fn show_grid(
    grid: &Vec<Vec<char>>,
    visited: &HashMap<(usize, usize), HashSet<(isize, isize)>>,
    obstacles: &HashSet<(usize, usize)>,
) {
    let mut new_grid = grid.clone();
    for i in 0..grid.len() {
        for j in 0..grid[0].len() {
            if obstacles.contains(&(i, j)) {
                new_grid[i][j] = 'O';
            } else if visited.contains_key(&(i, j)) {
                if grid[i][j] != '^' {
                    new_grid[i][j] = 'X';
                }
            }
        }
    }

    for line in new_grid {
        let chars: String = line.iter().collect();
        println!("{chars}");
    }
}

pub fn solve_part_1() {
    let grid = get_grid();
    let start_pos = find_start_pt(&grid).unwrap();
    let start_step = (-1, 0); // Guard starts facing up
    let path = walk_path(start_pos, start_step, &grid);
    show_grid(&grid, &path, &HashSet::default());

    let num_visited = path.len();
    println!("Number of distinct positions visited by guard: {num_visited}")
}

pub fn solve_part_2() {
    let grid = get_grid();
    let start_pos = find_start_pt(&grid).unwrap();
    let start_step = (-1, 0); // Guard starts facing up
    let path = walk_path(start_pos, start_step, &grid);
    let obstacles = obstruction_positions(start_pos, start_step, &grid);
    show_grid(&grid, &path, &obstacles);

    let num_obstructions = obstacles.len();
    println!("Number of possible obstruction positions that create a loop: {num_obstructions}")
}
