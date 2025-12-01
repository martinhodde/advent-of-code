use super::utils::lines_from_file;
use counter::Counter;
use std::collections::BinaryHeap;

const FILEPATH: &str = "inputs/day01.txt";

fn get_sorted_lists(lines: &Vec<String>) -> (Vec<u32>, Vec<u32>) {
    // Aggregate both input columns in binary heaps to maintain sorted order
    let mut heap1: BinaryHeap<u32> = BinaryHeap::new();
    let mut heap2: BinaryHeap<u32> = BinaryHeap::new();
    for line in lines {
        let vals: Vec<u32> = line
            .split_whitespace()
            .map(|s| s.parse().unwrap())
            .collect();
        heap1.push(vals[0]);
        heap2.push(vals[1]);
    }

    // Convert to sorted vectors
    (heap1.into_sorted_vec(), heap2.into_sorted_vec())
}

pub fn solve_part_1() {
    let lines = lines_from_file(FILEPATH).expect(&format!("Input file {FILEPATH} should exist"));
    // Compute the element-wise absolute difference between the two lists, then sum over the result
    let (list1, list2) = get_sorted_lists(&lines);
    let sum: u32 = list1
        .iter()
        .zip(list2.iter())
        .map(|(&v1, &v2)| v1.abs_diff(v2))
        .sum();

    println!("Total distance between lists: {sum}")
}

pub fn solve_part_2() {
    let lines = lines_from_file(FILEPATH).expect(&format!("Input file {FILEPATH} should exist"));
    let (list1, list2) = get_sorted_lists(&lines);
    let list2_ctr = list2.iter().collect::<Counter<_, u32>>();
    let sum: u32 = list1.iter().map(|n| n * list2_ctr[n]).sum();

    println!("Similarity score between lists: {sum}")
}
