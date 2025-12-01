import utils
import math

inputs = utils.read_file('inputs/day08.txt')
instructions = list(inputs[0].strip())
graph = {line[:3] : (line[7:10], line[12:15]) for line in inputs[2:]}

def compute_step_count(node, instrs, stop_criterion):
    """
    Calculate the number of steps it takes to reach the stopping criterion
    when executing the given instructions beginning at the provided start node.
    """
    step_count = 0
    while not stop_criterion(node):
        # Get next node based on next instruction
        instr = instrs.pop(0)
        node = graph[node][0 if instr == 'L' else 1]
        # Put instruction back in queue and increase step count
        instrs.append(instr)
        step_count += 1

    return step_count

def solve_p1():
    """
    Solve part 1.
    """
    return compute_step_count('AAA', instructions, lambda x : x == 'ZZZ')

def solve_p2():
    """
    Solve part 2.
    """
    # Starting nodes all end with 'A'
    start_nodes = [node for node in graph.keys() if node[2] == 'A']
    # Find the LCM of the step counts for all valid paths
    return math.lcm(*[compute_step_count(node, instructions, lambda x : x[2] == 'Z') for node in start_nodes])

print('Part 1: Number of steps required =', solve_p1())
print('Part 2: Number of steps required =', solve_p2())
