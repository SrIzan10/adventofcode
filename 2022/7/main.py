import json

input = open('input.txt', 'r').read().strip().splitlines()

import json

input_data = open('input.txt', 'r').read().strip().splitlines()
file_structure = {"name": "/", "type": "dir", "contents": [], "size": 0}

def write_file_structure():
    current_path = [file_structure]
    ls_output = []

    for line in input_data:
        if line.startswith('$'):
            if ls_output:
                current_path[-1]["contents"].extend(parse_ls(ls_output))
                ls_output = []

            command = line.strip('$ ').split()
            if command[0] == 'cd':
                if command[1] == '/':
                    current_path = [file_structure]
                elif command[1] == '..':
                    current_path.pop()
                else:
                    for item in current_path[-1]["contents"]:
                        if item["name"] == command[1] and item["type"] == "dir":
                            current_path.append(item)
                            break
                    else:
                        new_dir = {"name": command[1], "type": "dir", "contents": [], "size": 0}
                        current_path[-1]["contents"].append(new_dir)
                        current_path.append(new_dir)
        else:
            ls_output.append(line)

    if ls_output:
        current_path[-1]["contents"].extend(parse_ls(ls_output))

    calculate_directory_sizes(file_structure)

def parse_ls(ls_output):
    content = []
    for line in ls_output:
        parts = line.split()
        if parts[0] == 'dir':
            content.append({"name": parts[1], "type": "dir", "contents": [], "size": 0})
        else:
            content.append({"name": parts[1], "type": "file", "size": int(parts[0])})
    return content

def calculate_directory_sizes(directory):
    total_size = 0
    for item in directory["contents"]:
        if item["type"] == "file":
            total_size += item["size"]
        else:
            total_size += calculate_directory_sizes(item)
    directory["size"] = total_size
    return total_size

write_file_structure()

p1_size = 0
def part1(data=file_structure):
    global p1_size
    
    if data['type'] == 'dir':
        if data['size'] <= 100000:
            p1_size += data['size']
        
        for item in data['contents']:
            part1(item)
    
    return p1_size

# extremely large value to make sure min works
p2_size = 100 << 60
used_size = 70000000 - file_structure['size']
to_free_up = 30000000 - used_size
def part2():
    global p2_size
    
    def dfs(data):
        global p2_size
        
        if data['type'] == 'dir':
            if data['size'] > to_free_up:
                p2_size = min(p2_size, data['size'])
            
            for item in data['contents']:
                dfs(item)
    
    dfs(file_structure)
    
    return p2_size

print('Part1:', part1())
print('Part2:', part2())