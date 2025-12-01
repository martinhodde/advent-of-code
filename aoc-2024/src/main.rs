mod solutions;

use clap::Parser;
use solutions::utils::get_solver_fn;

/// The AoC problem defined by the day and part
#[derive(Parser)]
struct AoCProblem {
    #[arg(short = 'd', long = "day", value_parser = clap::value_parser!(u32).range(1..26))]
    day: u32,

    #[arg(short = 'p', long = "part", value_parser = clap::value_parser!(u32).range(1..3))]
    part: u32,
}

fn main() {
    let args = AoCProblem::parse();
    println!("Solving day: {:?}, part: {:?}", args.day, args.part);

    // Call the solver for requested day and part
    get_solver_fn(args.day, args.part).expect("Solver function should be implemented")()
}
