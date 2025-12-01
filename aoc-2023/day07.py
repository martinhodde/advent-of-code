import utils
from collections import Counter
from functools import cmp_to_key

inputs = utils.read_file('inputs/day07.txt')
hands = [line.split()[0] for line in inputs]
bids = [int(line.split()[1]) for line in inputs]

cards = {
    'A' : 14,
    'K' : 13,
    'Q' : 12,
    'J' : 11,
    'T' : 10,
    '9' : 9,
    '8' : 8,
    '7' : 7,
    '6' : 6,
    '5' : 5,
    '4' : 4,
    '3' : 3,
    '2' : 2
}

# Map joker card to lowest value in new dictionary
joker_cards = {k : 1 if k == 'J' else v for k, v in cards.items()}

def compare_hands(hand1, hand2, joker=False):
    """
    Return 1 if hand1 beats hand2 and return -1 if hand2 beats hand1.
    """
    # Keep track of the count of each card that appears in the hand
    hand1_counts = sorted(Counter(hand1).values(), reverse=True)
    hand2_counts = sorted(Counter(hand2).values(), reverse=True)

    if joker:
        # Replace the joker(s) with optimal cards
        hand1_counts, hand2_counts = replace_jokers(hand1), replace_jokers(hand2)

    if hand1_counts[0] == hand2_counts[0]:
        # Distinguish between full house and three of a kind as well as
        # between one pair and two pair
        if hand1_counts[0] in [2, 3] and hand1_counts[1] != hand2_counts[1]:
            return compare(hand1_counts[1], hand2_counts[1])

        # If we reach this point, the hands are of the same type and we must break the tie
        return break_tie(hand1, hand2, joker=joker)

    # Compare max counts since hands are different types
    return compare(hand1_counts[0], hand2_counts[0])

def break_tie(hand1, hand2, joker=False):
    """
    When hand1 and hand2 are of the same type, break the tie with the values
    of the first card that differs between the hands.
    """
    card_map = joker_cards if joker else cards
    return compare(tuple(map(card_map.get, hand1)), tuple(map(card_map.get, hand2)))

def compare(x, y):
    """
    This function is a simple comparator between two values that cannot be equal.
    """
    return 1 if x > y else -1

def replace_jokers(hand):
    """
    Replace the jokers in the given hand to maximize its strength and
    return the card counts of the new hand in descending order.
    """
    num_jokers = Counter(hand)['J']
    card_counts = sorted(Counter(hand.replace('J', '')).values(), reverse=True)

    if card_counts:
        # Put jokers toward the highest count card in hand
        card_counts[0] += num_jokers
    else:
        # Original hand is all jokers, consider this five of a kind
        card_counts = [5]

    return card_counts

def compute_winnings(joker=False):
    """
    Compute the total winnings given whether the 'J' card is considered a joker.
    """
    sorted_hands = sorted(zip(hands, bids), key=cmp_to_key(lambda x, y : compare_hands(x[0], y[0], joker=joker)))
    return sum(rank * bid for rank, (_, bid) in enumerate(sorted_hands, 1))

def solve_p1():
    """
    Solve part 1.
    """
    return compute_winnings()

def solve_p2():
    """
    Solve part 2.
    """
    return compute_winnings(joker=True)

print('Part 1: Total winnings =', solve_p1())
print('Part 2: Total winnings =', solve_p2())
