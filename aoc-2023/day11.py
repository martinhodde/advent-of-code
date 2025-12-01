import utils
from itertools import combinations

image = utils.read_file('inputs/day11.txt')

def galaxy_locs():
    """
    Return the coordinates of all galaxies in the input image. 
    """
    return {(i, j) for i, row in enumerate(image) for j, char in enumerate(row) if char == '#'}

def indices_to_expand():
    """
    Return the indices of the rows and columns in the image that contain no galaxies
    and will therefore expand.
    """
    row_idxs = {i for i, row in enumerate(image) if not any(c == '#' for c in row)}
    col_idxs = {j for j in range(len(image[0])) if not any(row[j] == '#' for row in image)}

    return row_idxs, col_idxs

def dist(p1, p2, row_idxs, col_idxs, exp_fac=2):
    """
    Compute the distance between coordinates p1 and p2, accounting for cosmic expansion.
    """
    # Distances across both axes before expansion
    v_dist, h_dist = abs(p2[0] - p1[0]), abs(p2[1] - p1[1])

    # Add expansion across both axes
    v_range = set(range(min(p1[0], p2[0]), max(p1[0], p2[0])))
    h_range = set(range(min(p1[1], p2[1]), max(p1[1], p2[1])))

    v_dist += (exp_fac - 1) * len(v_range.intersection(row_idxs))
    h_dist += (exp_fac - 1) * len(h_range.intersection(col_idxs))

    # Return new Manhattan distance between points
    return v_dist + h_dist


rows, cols = indices_to_expand()

def solve_p1():
    """
    Solve part 1.
    """
    return sum(dist(p1, p2, rows, cols) for p1, p2 in combinations(galaxy_locs(), 2))

def solve_p2():
    """
    Solve part 2.
    """
    expansion = 1000000  # Much larger expansion factor due to "older" age of galaxies
    return sum(dist(p1, p2, rows, cols, expansion) for p1, p2 in combinations(galaxy_locs(), 2))

print('Part 1: Sum of shortest paths between galaxy pairs =', solve_p1())
print('Part 2: Sum of shortest paths between older galaxy pairs =', solve_p2())
