import string
alphabet = list(string.ascii_lowercase)
uppercaseAlphabet = list(string.ascii_uppercase)

input = open('input.txt', 'r').read().strip().split('\n')

def part1():
    sum = 0
    for i in input:
        # i can use google right? https://stackoverflow.com/a/4789617
        firstPart, secondPart = i[:len(i)//2], i[len(i)//2:]
        firstPartArr, secondPartArr = [*firstPart], [*secondPart]
        sameChar = [x for x in firstPartArr if x in secondPartArr][:1]
        sameChar = ''.join(sameChar)
        
        sum += getPriority(sameChar)

    return sum

def part2():
    sum = 0
    i = [input[i:i+3] for i in range(0, len(input), 3)]
    for el in i:
        firstPart, secondPart, thirdPart = el
        firstPartArr, secondPartArr, thirdPartArr = [*firstPart], [*secondPart], [*thirdPart]
        sameChar = [x for x in firstPartArr if x in secondPartArr and x in thirdPartArr][:1]
        sameChar = ''.join(sameChar)

        sum += getPriority(sameChar)

    return sum

def getPriority(char):
    if char.isupper():
        # UPPERCASE
        return uppercaseAlphabet.index(char) + 27
    else:
        # lowercase
        # adding one because we start from zero
        return alphabet.index(char) + 1

print('Part1:', part1())
print('Part2:', part2())