use super::utils::lines_from_file;
use std::{cell::RefCell, collections::HashMap};

const FILEPATH: &str = "inputs/day11.txt";

thread_local! {
    // Global memoization cache
    static MEMO: RefCell<HashMap<(u64, u64), u64>> = RefCell::new(HashMap::new());
}

fn get_from_memo(key: (u64, u64)) -> Option<u64> {
    MEMO.with(|map| {
        let map = map.borrow();
        map.get(&key).copied()
    })
}

fn insert_into_memo(key: (u64, u64), value: u64) {
    MEMO.with(|map| {
        map.borrow_mut().insert(key, value);
    });
}

/// Recursively compute the number of stones that will ultimately result from the provided
/// stone after the given number of blinks occur.
///
/// Intermediate results are cached in a memoization hashmap for tractability.
fn num_stones_after_blinks(stone: u64, num_blinks: u64) -> u64 {
    if let Some(num_stones) = get_from_memo((stone, num_blinks)) {
        return num_stones;
    }

    let num_stones = if num_blinks == 0 {
        1 // Base case: no blinks remaining
    } else {
        if stone == 0 {
            // If the stone is engraved with a 0, replace it with a 1
            num_stones_after_blinks(1, num_blinks - 1)
        } else if digit_count(stone) % 2 == 0 {
            // If the stone is engraved with a number that has an even number of digits,
            // replace it with two stones. The left half of the digits are engraved on the
            // new left stone, and the right half of the digits are engraved on the new
            // right stone (the new numbers do not keep extra leading zeroes)
            let divisor = 10u64.pow(digit_count(stone) / 2);
            num_stones_after_blinks(stone / divisor, num_blinks - 1)
                + num_stones_after_blinks(stone % divisor, num_blinks - 1)
        } else {
            // If none of the other rules apply, the stone is replaced by a new stone:
            // the old stone's number multiplied by 2024
            num_stones_after_blinks(stone * 2024, num_blinks - 1)
        }
    };

    insert_into_memo((stone, num_blinks), num_stones);
    num_stones
}

/// Produce the stone engravings that result from performing the specified number of blinks,
/// given the intial vector of stone engravings.
fn simulate_blinks(stones: &Vec<u64>, num_blinks: u64) -> Vec<u64> {
    let mut new_stones: Vec<u64> = stones.clone();

    for _ in 0..num_blinks {
        let mut curr_stones: Vec<u64> = Vec::new();
        for stone in new_stones {
            if stone == 0 {
                // If the stone is engraved with a 0, replace it with a 1
                curr_stones.push(1);
            } else if digit_count(stone) % 2 == 0 {
                // If the stone is engraved with a number that has an even number of digits,
                // replace it with two stones. The left half of the digits are engraved on the
                // new left stone, and the right half of the digits are engraved on the new
                // right stone (the new numbers do not keep extra leading zeroes)
                let divisor = 10u64.pow(digit_count(stone) / 2);
                curr_stones.push(stone / divisor);
                curr_stones.push(stone % divisor);
            } else {
                // If none of the other rules apply, the stone is replaced by a new stone:
                // the old stone's number multiplied by 2024
                curr_stones.push(stone * 2024);
            }
        }
        new_stones = curr_stones;
    }

    new_stones
}

/// Compute the number of digits in the provided stone engraving
fn digit_count(stone: u64) -> u32 {
    (stone as f64).log10().floor() as u32 + 1
}

fn get_stones() -> Vec<u64> {
    lines_from_file(FILEPATH)
        .expect(&format!("Input file {FILEPATH} should exist"))
        .get(0)
        .expect(&format!("Input file {FILEPATH} should have contents"))
        .split(" ")
        .map(|s| s.parse().expect("Cannot parse string to u64"))
        .collect()
}

pub fn solve_part_1() {
    let stones = get_stones();
    let stones_after_blinking = simulate_blinks(&stones, 25).len();
    println!("Number of stones after blinking 25 times: {stones_after_blinking}");
}

pub fn solve_part_2() {
    let stones = get_stones();
    let stones_after_blinking: u64 = stones
        .into_iter()
        .map(|stone| num_stones_after_blinks(stone, 75))
        .sum();
    println!("Number of stones after blinking 75 times: {stones_after_blinking}");
}
