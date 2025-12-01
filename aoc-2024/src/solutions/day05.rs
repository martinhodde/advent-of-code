use super::utils::lines_from_file;
use std::collections::{HashMap, HashSet};

const FILEPATH: &str = "inputs/day05.txt";

/// Reorder the provided update according to the given ordering rules.
fn reorder_update(update: &Vec<u32>, order_rules: &HashMap<u32, HashSet<u32>>) -> Vec<u32> {
    let mut new_update = update.clone();
    for i in 0..new_update.len() {
        let mut j = 0;
        while j < i {
            if order_rules.contains_key(&new_update[i])
                && order_rules[&new_update[i]].contains(&new_update[j])
            {
                // If the page at i incorrectly appears before the page at j,
                // move the page at i directly after index j
                let page = new_update.remove(j);
                new_update.insert(i, page);
            } else {
                j += 1;
            }
        }
    }
    new_update
}

/// Return whether the provided update adheres to the given ordering rules.
fn is_update_valid(update: &Vec<u32>, order_rules: &HashMap<u32, HashSet<u32>>) -> bool {
    for (i, page) in update.iter().enumerate() {
        // For each page in the update, check for an ordering rule violation in any
        // of the preceding pages, and return false if one is found
        if update
            .iter()
            .take(i)
            .any(|p| order_rules.contains_key(page) && order_rules[page].contains(p))
        {
            return false;
        }
    }
    true // No ordering rule violations are detected
}

/// Take the sum of all values at the middle index of each provided update.
fn middle_page_sum(updates: &Vec<Vec<u32>>) -> u32 {
    updates.iter().map(|update| update[update.len() / 2]).sum()
}

fn get_ordering_rules() -> HashMap<u32, HashSet<u32>> {
    let rule_tuples: Vec<(u32, u32)> = lines_from_file(FILEPATH)
        .expect(&format!("Input file {FILEPATH} should exist"))
        .into_iter()
        .take_while(|s| s != "")
        .map(|rule| {
            let (p1, p2) = rule.split_once('|').unwrap();
            (p1.parse().unwrap(), p2.parse().unwrap())
        })
        .collect();

    // To handle rules that have the same preceding page, map each "before" page
    // to a set of all "after" pages present in the input.
    let mut rule_map: HashMap<u32, HashSet<u32>> = HashMap::new();
    for (p1, p2) in rule_tuples {
        rule_map.entry(p1).or_default().insert(p2);
    }
    rule_map
}

fn get_updates() -> Vec<Vec<u32>> {
    lines_from_file(FILEPATH)
        .expect(&format!("Input file {FILEPATH} should exist"))
        .into_iter()
        .skip_while(|s| s != "")
        .skip(1)
        .map(
            // Create sub-vec for individual update
            |update| update.split(',').map(|p| p.parse().unwrap()).collect(),
        )
        .collect()
}

pub fn solve_part_1() {
    let ordering_rules = get_ordering_rules();
    let valid_updates: Vec<Vec<u32>> = get_updates()
        .into_iter()
        .filter(|update| is_update_valid(update, &ordering_rules))
        .collect();
    let sum = middle_page_sum(&valid_updates);
    println!("Middle page number sum of correctly-ordered updates: {sum}")
}

pub fn solve_part_2() {
    let ordering_rules = get_ordering_rules();
    let corrected_updates: Vec<Vec<u32>> = get_updates()
        .into_iter()
        .filter(|update| !is_update_valid(&update, &ordering_rules))
        .map(|update| reorder_update(&update, &ordering_rules))
        .collect();
    let sum = middle_page_sum(&corrected_updates);
    println!("Middle page number sum of corrected formerly out-of-order updates: {sum}")
}
