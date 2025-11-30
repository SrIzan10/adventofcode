pub mod day01;

pub fn run_day(day: u8) {
    match day {
        1 => day01::run(),
        _ => println!("Day {} not implemented yet", day),
    }
}
