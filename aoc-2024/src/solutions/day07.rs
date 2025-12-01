use super::utils::lines_from_file;

const FILEPATH: &str = "inputs/day07.txt";

/// Compute the sum of all calibration test values whose operands can be combined
/// satisfy the provided equation validation function.
fn total_calibration_result(
    calibration_eqs: &Vec<(u64, Vec<u64>)>,
    validation_fn: fn(&Vec<u64>, u64, u64) -> bool,
) -> u64 {
    calibration_eqs
        .into_iter()
        .filter(|(tst, ops)| validation_fn(&ops[1..].to_vec(), *tst, ops[0]))
        .map(|(tst, _)| tst)
        .sum()
}

/// Return whether some combination of + and * operators on the operands, evaluated
/// from left to right, will result in the provided test value.
fn is_valid_eq(operands: &Vec<u64>, test_val: u64, running_result: u64) -> bool {
    if operands.is_empty() {
        return running_result == test_val;
    } else if running_result > test_val {
        return false; // Early exit if running result has exceeded test value
    }

    // Recurse on both operators for remaining operands
    is_valid_eq(
        &operands[1..].to_vec(),
        test_val,
        running_result + operands[0],
    ) || is_valid_eq(
        &operands[1..].to_vec(),
        test_val,
        running_result * operands[0],
    )
}

/// Return whether some combination of +, *, and || operators on the operands, evaluated
/// from left to right, will result in the provided test value.
fn is_valid_eq_with_concat(operands: &Vec<u64>, test_val: u64, running_result: u64) -> bool {
    if operands.is_empty() {
        return running_result == test_val;
    } else if running_result > test_val {
        return false; // Early exit if running result has exceeded test value
    }

    // Recurse on all 3 operators for remaining operands
    is_valid_eq_with_concat(
        &operands[1..].to_vec(),
        test_val,
        running_result + operands[0],
    ) || is_valid_eq_with_concat(
        &operands[1..].to_vec(),
        test_val,
        running_result * operands[0],
    ) || is_valid_eq_with_concat(
        &operands[1..].to_vec(),
        test_val,
        running_result * 10u64.pow(operands[0].to_string().len() as u32) + operands[0],
    )
}

fn get_calibration_eqs() -> Vec<(u64, Vec<u64>)> {
    lines_from_file(FILEPATH)
        .expect(&format!("Input file {FILEPATH} should exist"))
        .into_iter()
        .map(|line| {
            let (calibration, operands) = line.split_once(':').unwrap();
            (
                calibration.parse().unwrap(),
                operands
                    .split_whitespace()
                    .map(|op| op.parse().unwrap())
                    .collect(),
            )
        })
        .collect()
}

pub fn solve_part_1() {
    let equations = get_calibration_eqs();
    let sum = total_calibration_result(&equations, is_valid_eq);
    println!("Total calibration result: {sum}")
}

pub fn solve_part_2() {
    let equations = get_calibration_eqs();
    let sum = total_calibration_result(&equations, is_valid_eq_with_concat);
    println!("Total calibration result with concatenation: {sum}")
}
