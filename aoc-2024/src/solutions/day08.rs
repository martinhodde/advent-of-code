use super::utils::{lines_from_file, try_step};
use itertools::Itertools;
use std::collections::{HashMap, HashSet};

const FILEPATH: &str = "inputs/day08.txt";

/// Given a map of frequencies to their antenna locations, return the distinct locations
/// of all antinodes across all frequencies according to the provided antinode location function.
fn antinode_locations(
    antenna_locs: &HashMap<char, HashSet<(usize, usize)>>,
    grid: &Vec<Vec<char>>,
    antinode_fn: fn((usize, usize), (usize, usize), &Vec<Vec<char>>) -> HashSet<(usize, usize)>,
) -> HashSet<(usize, usize)> {
    antenna_locs
        .iter()
        .map(|(_, locs)| {
            // For each pairwise combination of antenna locations of the same frequency,
            // determine the possible antinode locations
            locs.iter()
                .combinations(2)
                .map(|loc| antinode_fn(*loc[0], *loc[1], grid))
                .flatten() // Flatten antinodes across all pairwise combos for the current frequency
        })
        .flatten() // Flatten antinodes across all frequencies
        .collect()
}

/// Return a map of each frequency to the set of locations of the associated antennas.
fn antenna_locations(grid: &Vec<Vec<char>>) -> HashMap<char, HashSet<(usize, usize)>> {
    let mut locations = HashMap::new();
    for (i, row) in grid.iter().enumerate() {
        for (j, &c) in row.iter().enumerate() {
            if c != '.' {
                locations
                    .entry(c)
                    .or_insert_with(HashSet::new)
                    .insert((i, j));
            }
        }
    }
    locations
}

/// Given two antenna locations of the same frequency, compute the (up to) two possible antinode
/// locations, subject to the bounds of the provided grid.
fn get_antinode_pts(
    loc1: (usize, usize),
    loc2: (usize, usize),
    grid: &Vec<Vec<char>>,
) -> HashSet<(usize, usize)> {
    let (i_1, j_1) = (loc1.0 as isize, loc1.1 as isize);
    let (i_2, j_2) = (loc2.0 as isize, loc2.1 as isize);

    // For each of the two endpoints, trace a vector from the other point to itself, then add
    // this vector to the current point to determine potential antinode locations, filtering
    // out those that are out of bounds
    [
        (loc1, (i_1 - i_2, j_1 - j_2)),
        (loc2, (i_2 - i_1, j_2 - j_1)),
    ]
    .into_iter()
    .filter_map(|(loc, step)| try_step(loc, step, grid))
    .collect()
}

/// Given two antenna locations of the same frequency, compute all possible antinode
/// locations, subject to the bounds of the provided grid and taking into account the
/// effects of resonant harmonics.
fn get_antinode_pts_with_resonance(
    loc1: (usize, usize),
    loc2: (usize, usize),
    grid: &Vec<Vec<char>>,
) -> HashSet<(usize, usize)> {
    let (i_1, j_1) = (loc1.0 as isize, loc1.1 as isize);
    let (i_2, j_2) = (loc2.0 as isize, loc2.1 as isize);

    // For each of the two endpoints, trace a vector from the other point to itself, then add
    // this vector to the current point repeatedly to determine potential antinode locations,
    // terminating when the vector addition results in stepping out of bounds
    let mut antinode_locs: HashSet<(usize, usize)> = HashSet::from([loc1, loc2]);
    let loc_dirs = [
        (loc1, (i_1 - i_2, j_1 - j_2)),
        (loc2, (i_2 - i_1, j_2 - j_1)),
    ];

    for (start_loc, step) in loc_dirs {
        let mut loc = start_loc;
        while let Some(new_loc) = try_step(loc, step, grid) {
            antinode_locs.insert(new_loc);
            loc = new_loc;
        }
    }

    antinode_locs
}

fn get_grid() -> Vec<Vec<char>> {
    lines_from_file(FILEPATH)
        .expect(&format!("Input file {FILEPATH} should exist"))
        .into_iter()
        .map(|line| line.chars().collect())
        .collect()
}

pub fn solve_part_1() {
    let grid = get_grid();
    let antenna_locs = antenna_locations(&grid);
    let antinode_locs = antinode_locations(&antenna_locs, &grid, get_antinode_pts);
    let num_antinode_locs = antinode_locs.len();
    println!("Number of unique locations that contain an antinode: {num_antinode_locs}")
}

pub fn solve_part_2() {
    let grid = get_grid();
    let antenna_locs = antenna_locations(&grid);
    let antinode_locs = antinode_locations(&antenna_locs, &grid, get_antinode_pts_with_resonance);
    let num_antinode_locs = antinode_locs.len();
    println!("Number of unique locations that contain an antinode, considering resonant harmonics: {num_antinode_locs}")
}
