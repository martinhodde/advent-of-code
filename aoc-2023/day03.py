import utils

schematic = utils.read_file('inputs/day03.txt')

def build_bboxes():
    """
    Compute bounding boxes for all part numbers in the schematic.
    """
    part_num_bboxes = {}
    curr_part_num = ''
    top, left = 0, 0
    for i, line in enumerate(schematic):
        for j, character in enumerate(line):
            if character.isdigit():
                if not curr_part_num:
                    # In this case, we have encountered the start of a number,
                    # so we note the top-left coordinates of the bounding box
                    top, left = i - 1, j - 1
                curr_part_num += character

            elif curr_part_num:
                # Here, we have reached the end of the number so we close the bounding box
                # and map it to its part number
                bottom, right = i + 1, j
                part_num_bboxes[(top, left, bottom, right)] = int(curr_part_num)
                curr_part_num = ''  # Reset current part number

    return part_num_bboxes

bboxes = build_bboxes()

def solve_p1():
    """
    Solve part 1.
    """
    # Check all the symbols for intersections with the bounding boxes
    part_num_sum = 0
    for i, line in enumerate(schematic):
        for j, character in enumerate(line.strip()):
            if character != '.' and not character.isdigit():
                # Check symbol location against all part number bounding boxes
                for (top, left, bottom, right), part_num in bboxes.items():
                    if i in range(top, bottom + 1) and j in range(left, right + 1):
                        # Symbol is adjacent to a part number
                        part_num_sum += part_num

    return part_num_sum

def solve_p2():
    """
    Solve part 2.
    """
    # Check all gear symbols for intersections with two bounding boxes
    gear_ratio_sum = 0
    for i, line in enumerate(schematic):
        for j, character in enumerate(line.strip()):
            if character == '*':
                # Check gear location against all part number bounding boxes
                part_num_count, gear_ratio = 0, 1
                for (top, left, bottom, right), part_num in bboxes.items():
                    if i in range(top, bottom + 1) and j in range(left, right + 1):
                        gear_ratio *= part_num
                        part_num_count += 1

                # Only add gear ratio if the gear is adjacent to exactly two part numbers
                if part_num_count == 2:
                    gear_ratio_sum += gear_ratio

    return gear_ratio_sum

print('Part 1: Part number sum =', solve_p1())
print('Part 2: Gear ratio sum =', solve_p2())
