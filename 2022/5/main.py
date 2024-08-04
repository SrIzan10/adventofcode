input = open('input.txt', 'r').read().splitlines()

def get_visualization():
    currentLines = []
    for l in input:
        if l != '':
            currentLines.append(l)
        else:
            break
    return currentLines

def get_boxes():
    columns = {}
    visualization = get_visualization()
    
    for line in visualization[:-1]:
        chunks = [line[i:i+4].strip() for i in range(0, len(line), 4)]
        for i, chunk in enumerate(chunks, 1):
            if chunk and chunk != '':
                # Remove brackets
                item = chunk.strip('[]')
                if i not in columns:
                    columns[i] = []
                columns[i].append(item)
    
    return columns

def parse_instructions():
    instructions = input[input.index('')+1:]
    parsed_insts = []
    for inst in instructions:
        inst = inst.split()
        amount = int(inst[1])
        from_stack = int(inst[3])
        to_stack = int(inst[5])
        parsed_insts.append({'amount': amount, 'fromStack': from_stack, 'toStack': to_stack})
    return parsed_insts

def get_code(boxes):
    code = ''
    for key, value in sorted(boxes.items()):
        code += value[0]
    return code

def part1():
    instructions = parse_instructions()
    boxes = get_boxes()

    for inst in instructions:
        amount = inst['amount']
        from_stack = inst['fromStack']
        to_stack = inst['toStack']
        # this just runs the code the amount of times necessary, thus the amt variable never gets used
        for amt in range(amount):
            boxes[to_stack].insert(0, boxes[from_stack].pop(0))
    return get_code(boxes)


def part2():
    instructions = parse_instructions()
    boxes = get_boxes()
    for inst in instructions:
        amount = int(inst['amount'])
        from_stack = inst['fromStack']
        to_stack = inst['toStack']
        
        # heres the sauce
        to_move = boxes[from_stack][:amount]
        boxes[from_stack] = boxes[from_stack][amount:]
        boxes[to_stack] = to_move + boxes[to_stack]
    
    return get_code(boxes)

print('Part1:', part1())
print('Part2:', part2())