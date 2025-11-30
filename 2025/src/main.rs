mod days;
mod utils;

use std::env;

fn main() {
    let args: Vec<String> = env::args().collect();
    
    if args.len() < 2 {
        println!("Usage: cargo run <day>");
        return;
    }

    let day: u8 = match args[1].parse() {
        Ok(d) => d,
        Err(_) => {
            println!("Invalid day number");
            return;
        }
    };

    days::run_day(day);
}
