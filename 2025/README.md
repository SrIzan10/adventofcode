# Advent of Code 2025 - Rust

## Daily Workflow

To add a new day's solution (e.g., Day 2):

1.  **Create the Source File**:
    - Copy `src/days/day01.rs` to `src/days/day02.rs`.
    - Update the logic in `part1` and `part2`.

2.  **Register the Module**:
    - Open `src/days/mod.rs`.
    - Add `pub mod day02;` at the top.
    - Add a match arm for the new day:
      ```rust
      2 => day02::run(),
      ```

3.  **Add Input**:
    - Create a new file `inputs/day02.txt`.
    - Paste your puzzle input into it.

4.  **Run**:
    - Execute the solution with:
      ```bash
      cargo run 2
      ```
