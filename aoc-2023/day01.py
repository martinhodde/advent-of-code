import utils

inputs = utils.read_file('inputs/day01.txt')

digit_words = {
    # Overlapping digit word edge cases appearing in inputs
    'twone' : '21',
    'oneight' : '18',
    'eightwo' : '82',
    # Regular old digits
    'one' : '1',
    'two' : '2',
    'three' : '3',
    'four' : '4',
    'five' : '5',
    'six' : '6',
    'seven' : '7',
    'eight' : '8',
    'nine' : '9'
}

def get_calibration_value(calibration_line):
    """
    Extract the calibration value from the first and last digits of the
    provided calibration line.
    """
    first_digit, last_digit = None, None
    for character in calibration_line:
        if character.isdigit():
            if first_digit is None:
                # Only set first digit if we have not yet encountered
                # any digits in the current input line
                first_digit = character
            last_digit = character  # Update last digit unconditionally

    return int(first_digit + last_digit)

def preprocess_calibration_line(calibration_line):
    """
    Perform a preliminary pass over the input line to replace digit words
    with their corresonding digits according to the digit map.
    """
    for digit_word, digit in digit_words.items():
        if digit_word in calibration_line:
            calibration_line = calibration_line.replace(digit_word, digit)

    return calibration_line

def solve_p1():
    """
    Solve part 1.
    """
    return sum(map(get_calibration_value, inputs))

def solve_p2():
    """
    Solve part 2.
    """
    return sum(map(get_calibration_value, map(preprocess_calibration_line, inputs)))

print('Part 1: Calibration sum =', solve_p1())
print('Part 2: Calibration sum =', solve_p2())
