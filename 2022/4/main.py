input = open('input.txt', 'r').read().strip().split('\n')

one_to_ninetynine = list(range(1,100))

def part1():
    is_second_pair = False
    summed_first_pair = False
    sum = 0
    for pair in input:
        divided_pair = pair.split(',')
        pair1, pair2 = parse_number_ranges(divided_pair[0]), parse_number_ranges(divided_pair[1])
        not_matches = returnNotMatches(pair1, pair2)
        # remove duplicates from array (there are a few which totally overlap and cause two empty arrays) https://stackoverflow.com/a/3819368
        not_matches = [list(t) for t in {tuple(item) for item in not_matches}]
        for i in not_matches:
            if not i:
                sum += 1
    
    return sum


def part2():
    sum = 0
    for pair in input:
        range1, range2 = pair.split(',')
        start1, end1 = map(int, range1.split('-'))
        start2, end2 = map(int, range2.split('-'))
        
        if max(start1, start2) <= min(end1, end2):
            sum += 1
    
    return sum

def parse_number_ranges(n):
    numbers = n.split('-')
    numbers = [int(i) for i in numbers]
    return one_to_ninetynine[numbers[0] - 1:numbers[1]]

# https://stackoverflow.com/a/35713174
def returnNotMatches(a, b):
    return [[x for x in a if x not in b], [x for x in b if x not in a]]

print('Part1:', part1())
print('Part2:', part2())