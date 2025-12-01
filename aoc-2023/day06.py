import utils
import math

inputs = utils.read_file('inputs/day06.txt')

def num_ways_to_win(T, D):
    """
    Distance traveled for a single race is given by the expression -t^2 + Tt
    where t is the charge time and T is the total race duration. This quantity
    must be greater than the current distance record D to win: -t^2 + Tt - D > 0.
    This function computes the numbe of ways to win a race by finding the difference
    between the roots of this quadratic.
    """

    disc = math.sqrt(T ** 2 - 4 * D)
    # If the discriminant is a perfect square, set higher delta for rounding
    delta = 1.0 if disc.is_integer() else 0.5
    lower, upper = round((T - disc) / 2 + delta), round((T + disc) / 2 - delta)
    return upper - lower + 1


# Interpretation of input for part 1
times, distances = (list(map(int, line.split()[1:])) for line in inputs)

# Interpretation of input for part 2
time, distance = (int(''.join(line.split()[1:])) for line in inputs)

def solve_p1():
    """
    Solve part 1.
    """
    return math.prod(num_ways_to_win(T, D) for T, D in zip(times, distances))

def solve_p2():
    """
    Solve part 2.
    """
    return num_ways_to_win(time, distance)

print('Part 1: Product of number of ways to break record =', solve_p1())
print('Part 2: Number of ways to break single race record =', solve_p2())
