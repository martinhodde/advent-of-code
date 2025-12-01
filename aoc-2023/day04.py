import utils

inputs = utils.read_file('inputs/day04.txt')

def build_card_map():
    """
    Construct a map of card index numbers to their winning and owned numbers.
    """
    card_num_map = {}
    for card_idx, card in enumerate(inputs):
        # Determine winning numbers and numbers we own for each card
        winning_nums, owned_nums = map(lambda x : set(x.strip().split()), card.split(':')[1].split('|'))
        card_num_map[card_idx + 1] = (winning_nums, owned_nums)

    return card_num_map

card_map = build_card_map()

memo = {}  # Memoize results
def compute_total_cards(card_nums):
    """
    Given a range of card numbers, compute the number of total scratch cards
    that will be won from them.
    """
    # Filter out card numbers that do not appear in the data
    card_subset = [(num, card_map[num]) for num in card_nums if num in card_map]
    total_scratch_cards = len(card_subset)

    for card_num, (winning, owned) in card_subset:
        # Find how many winning numbers we have for current card
        num_winners_owned = len(owned.intersection(winning))
        # Recursively compute total cards won from new cards if result has not already been cached
        new_cards = range(card_num + 1, card_num + num_winners_owned + 1)
        if new_cards not in memo:
            memo[new_cards] = compute_total_cards(new_cards)
        total_scratch_cards += memo[new_cards]

    return total_scratch_cards

def solve_p1():
    """
    Solve part 1.
    """
    card_point_sum = 0
    for winning_nums, owned_nums in card_map.values():
        # Find how many winning numbers we have and compute card score accordingly
        num_winners_owned = len(owned_nums.intersection(winning_nums))
        card_point_sum += 2 ** (num_winners_owned - 1) if num_winners_owned > 0 else 0

    return card_point_sum

def solve_p2():
    """
    Solve part 2.
    """
    return compute_total_cards(range(1, len(inputs) + 1))

print('Part 1: Card point sum =', solve_p1())
print('Part 2: Number of scratchcards =', solve_p2())
