import utils

inputs = utils.read_file('inputs/day09.txt')
sequences = [list(map(int, line.split())) for line in inputs]

def extrapolate(num_seq, backward=0):
    """
    Return the number that follows (or precedes if backward) the given number sequence.
    """
    if not any(num_seq):
        # Base case: return zero when the number sequence contains only zeros
        return 0

    # Alternate sign of coefficient with each successive recursive call for backward case
    delta = extrapolate([num_seq[i + 1] - num_seq[i] for i in range(len(num_seq) - 1)], -backward)
    return backward * num_seq[0] + delta if backward else num_seq[-1] + delta

def solve_p1():
    """
    Solve part 1.
    """
    return sum(map(extrapolate, sequences))

def solve_p2():
    """
    Solve part 2.
    """
    return sum(map(lambda seq : extrapolate(seq, backward=1), sequences))

print('Part 1: Sum of extrapolated values =', solve_p1())
print('Part 2: Sum of extrapolated values =', solve_p2())
