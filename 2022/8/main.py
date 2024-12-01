input = open('input.txt', 'r').read().strip().splitlines()

def count_visible_trees(grid):
    rows = len(grid)
    cols = len(grid[0])
    visible_trees = 0

    for row in range(rows):
        for col in range(cols):
            if is_visible(grid, row, col):
                visible_trees += 1

    return visible_trees

def is_visible(grid, row, col):
    height = grid[row][col]
    rows = len(grid)
    cols = len(grid[0])

    # Trees on the edge are always visible
    if row == 0 or col == 0 or row == rows - 1 or col == cols - 1:
        return True

    # Check visibility from left
    if all(grid[row][c] < height for c in range(col)):
        return True

    # Check visibility from right
    if all(grid[row][c] < height for c in range(col + 1, cols)):
        return True

    # Check visibility from top
    if all(grid[r][col] < height for r in range(row)):
        return True

    # Check visibility from bottom
    if all(grid[r][col] < height for r in range(row + 1, rows)):
        return True

    return False

# split hte input like the grid variable
grid = [list(map(int, row.split())) for row in input]
print(grid)

result = count_visible_trees(grid)
print(f"Number of visible trees: {result}")

def part1():
    return

def part2():
    return

print('Part1:', part1())
print('Part2:', part2())