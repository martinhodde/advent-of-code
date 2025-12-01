import math
import utils

inputs = utils.read_file('inputs/day02.txt')
games = [line.replace(',', '').replace(';', '').split() for line in inputs]

# Maximum configured cube number for each color
max_cubes = {'red' : 12, 'green' : 13, 'blue' : 14}

def solve_p1():
    """
    Solve part 1.
    """
    game_id_sum = 0
    for game_idx, game in enumerate(games):
        game_id_sum += game_idx + 1  # Preemptively add id to sum
        for str_idx, game_str in enumerate(game):
            # Check if each number of cubes in input line is possible
            if game_str in ['red', 'green', 'blue']:
                num_cubes = int(game[str_idx - 1])
                if num_cubes > max_cubes[game_str]:
                    # Subtract id from sum if game is impossible
                    game_id_sum -= game_idx + 1
                    break

    return game_id_sum

def solve_p2():
    """
    Solve part 2.
    """
    power_sum = 0
    for game in games:
        min_cubes_req = {'red' : 1, 'green' : 1, 'blue' : 1}
        for str_idx, game_str in enumerate(game):
            if game_str in ['red', 'green', 'blue']:
                num_cubes = int(game[str_idx - 1])
                # Set min required cube number to the max for each color
                min_cubes_req[game_str] = max(num_cubes, min_cubes_req[game_str])

        power_sum += math.prod(min_cubes_req.values())  # Compute power and add to total

    return power_sum

print('Part 1: Game ID sum =', solve_p1())
print('Part 2: Power sum =', solve_p2())
