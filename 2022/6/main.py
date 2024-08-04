import math

input = open('input.txt', 'r').read().strip()

def get_substring(group,subst):
    return input[group:group+subst]

# https://stackoverflow.com/a/74312161
def has_duplicates(string):
    return len(set(string)) < len(string)

def find_nd_substring(subst):
    for g in range(math.trunc(len(input))):
        if not has_duplicates(get_substring(g, subst)):
            return g + subst

def part1():
    return find_nd_substring(4)

def part2():
    return find_nd_substring(14)

print('Part1:', part1())
print('Part2:', part2())