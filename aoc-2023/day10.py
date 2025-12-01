import utils

grid = utils.read_file('inputs/day10.txt')

def find_start():
    """
    Return the coordinates of the starting tile in the grid.
    """
    for i, line in enumerate(grid):
        for j, tile in enumerate(line):
            if tile == 'S':
                start = (i, j)
                replace_start(start)  # Swap 'S' with actual pipe tile
                return start

    raise ValueError('Grid must contain start tile!')

def replace_start(start):
    """
    Replace the tile at the provided starting coordinates with the appropriate pipe.
    """
    i, j = start
    # Check for connections to left, right, top, and bottom tiles
    start_connections = (
        grid[i][j - 1] in '-LF',
        grid[i][j + 1] in '-J7',
        grid[i - 1][j] in '|7F',
        grid[i + 1][j] in '|LJ'
    )

    # Determine the start pipe according to adjacent connections
    match start_connections:
        case (False, False, True, True):
            pipe = '|'
        case (True, True, False, False):
            pipe = '-'
        case (False, True, True, False):
            pipe = 'L'
        case (True, False, True, False):
            pipe = 'J'
        case (True, False, False, True):
            pipe = '7'
        case (False, True, False, True):
            pipe = 'F'

    # Swap start tile with the correct pipe 
    grid[i] = grid[i][:j] + pipe + grid[i][j + 1:]

def get_adjacent(pipe):
    """
    Return the coordinates of both tiles that are connected to the provided pipe tile.
    """
    i, j = pipe
    match grid[i][j]:
        case '|':  # Connected to top and bottom
            return [(i - 1, j), (i + 1, j)]
        case '-':  # Connected to left and right
            return [(i, j - 1), (i, j + 1)]
        case 'L':  # Connected to top and right
            return [(i - 1, j), (i, j + 1)]
        case 'J':  # Connected to top and left
            return [(i - 1, j), (i, j - 1)]
        case '7':  # Connected to bottom and left
            return [(i + 1, j), (i, j - 1)]
        case 'F':  # Connected to bottom and right
            return [(i + 1, j), (i, j + 1)]

def tiles_in_loop(start):
    """
    Return the coordinates of all pipes in the loop containing the given starting tile.
    """
    queue, loop = [start], {start}
    # Perform a breadth first search
    while queue:
        pipe = queue.pop(0)
        for adj_pipe in get_adjacent(pipe):
            if adj_pipe not in loop:
                queue.append(adj_pipe)
                loop.add(adj_pipe)

    return loop

def num_tiles_enclosed_by_loop(loop):
    """
    Compute the number of tiles enclosed by the provided loop coordinate set.
    """
    num_enclosed = 0
    for i, line in enumerate(grid):
        # Maintain the latest "opening" pipe per row (e.g. 'F', 'L')
        open_pipe, in_loop = [], False
        for j, pipe in enumerate(line):
            if (i, j) in loop:
                if pipe == '|':
                    # The '|' pipe signifies a vertical partition, so we change the
                    # in-loop status of subsequent pipes
                    in_loop ^= True

                elif pipe in 'LF':
                    # The opening pipes do not change in-loop status alone, so we save
                    # them for future reference upon reaching a closing pipe
                    open_pipe = pipe

                elif pipe in 'J7':
                    # When a closing pipe is reached, we match it with the most recent
                    # opening pipe. The in-loop status only changes if the combined
                    # opening and closing pipes form a vertical barricade:
                    # ↳↲ and ↱↰ do not change in-loop status but ↳↰ and ↱↲ do
                    if open_pipe + pipe in ['FJ', 'L7']:
                        in_loop ^= True

            elif in_loop:
                # If the current tile is not within the loop itself, but enclosed by it,
                # we increment the counter
                num_enclosed += 1

    return num_enclosed


pipe_loop = tiles_in_loop(find_start())

def solve_p1():
    """
    Solve part 1.
    """
    # The distance to the farthest tile in the loop is half the length of the entire loop
    return len(pipe_loop) // 2

def solve_p2():
    """
    Solve part 2.
    """
    return num_tiles_enclosed_by_loop(pipe_loop)

print('Part 1: Distance to farthest position from start in the loop =', solve_p1())
print('Part 2: Number of tiles enclosed by the loop =', solve_p2())
