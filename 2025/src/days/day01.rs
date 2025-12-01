use crate::utils;

pub fn run() {
    let input = utils::read_input(1);
    println!("Part 1: {}", part1(&input));
    println!("Part 2: {}", part2(&input));
}

fn part1(input: &str) -> String {
    let mut number = 50;
    let mut solution = 0;
    for imp in input.lines() {
        let c = imp.chars().next().unwrap_or(' ');
        let nums = &imp[1..].parse().unwrap();
        match c {
            'R' => {
                number += nums;
                if number == 99 {
                    println!("holy shitasd f");
                    solution += 1
                } else if number < 0 {
                    number += 99;
                    println!("summing {}", number)
                }

                if number == 0 {
                    println!("holy shit");
                    solution += 1
                }
                println!("{}", number);
            },
            'L' => {
                number -= nums;
                if number == 100 {
                    println!("holy shitasd f");
                    solution += 1
                } else if number < 0 {
                    number += 100;
                    println!("summing {}", number)
                }

                if number == 0 || number == 100 {
                    println!("holy shit");
                    solution += 1;
                    number = 0;
                }
                println!("{}", number); 
            },
            _ => {}
        }
    }
    return solution.to_string();
}

fn part2(_input: &str) -> String {
    // TODO: Solve part 2
    return "Not implemented".to_string()
}
