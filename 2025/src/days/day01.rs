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
        let nums: i32 = imp[1..].parse().unwrap();

        for cur in 0..nums {
            if c == 'L' {
                number -= 1;
            } else {
                number += 1;
            }

            // tomfoolery
            if number < 0 {
                number = 99;
            } else if number == 100 {
                number = 0;
            }

            // nums - 1 so jank lmao
            if cur == nums - 1 && number == 0 {
                solution += 1;
            }
        }
    }
    return solution.to_string();
}

fn part2(input: &str) -> String {
    let mut number = 50;
    let mut solution = 0;
    for imp in input.lines() {
        let c = imp.chars().next().unwrap_or(' ');
        let nums: i32 = imp[1..].parse().unwrap();

        for _ in 0..nums {
            if c == 'L' {
                number -= 1;
            } else {
                number += 1;
            }

            // tomfoolery
            if number < 0 {
                number = 99;
            } else if number == 100 {
                number = 0;
            }

            // did we land
            if number == 0 {
                solution += 1;
            }
        }
    }
    return solution.to_string();
}
