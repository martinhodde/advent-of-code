import bisect
import utils

inputs = utils.read_file('inputs/day05.txt')

def build_maps():
    """
    Build the map sequence from the input.
    """
    # Parse seed-to-soil map, soil-to-fertilizer map, and so on...
    map_seq, new_map = [], []
    for line in inputs[2:]:
        if line[0].isalpha():  # Heading signifies beginning of new map
            new_map = []
        elif line == '\n':  # New line signifies end of map
            map_seq.append(new_map)
        else:
            # Append mapping range to current map in sorted order by source number
            dst, src, range_len = map(int, line.split())
            bisect.insort(new_map, (src, dst, range_len))

    return map_seq

def seed_to_location(seed, maps):
    """
    Map the given seed to its location number using the provided sequence of maps.
    """
    map_num = seed  # Current number representation in the map chain
    for curr_map in maps:
        # Locate the only possible relevant mapping based on source number
        idx = bisect.bisect(curr_map, (map_num, float('inf'), float('inf')))
        if idx > 0:  # If idx is 0, no mapping exists
            source, dest, length = curr_map[idx - 1]
            if map_num < source + length:
                # Update number representation if mapping exists
                map_num += dest - source

    return map_num

def seed_ranges_to_loc_ranges(seed_ranges, maps):
    """
    Map the given seed ranges to their associated location ranges using the
    provided sequence of maps.
    """
    ranges = seed_ranges
    for curr_map in maps:
        next_ranges = []
        while ranges:
            start, stop = ranges.pop()
            for src, dst, length in curr_map:
                src_upper = src + length - 1  # Upper bound of current source range
                delta = dst - src  # src -> dst diff
                if src <= start <= stop <= src_upper:
                    # Current range is entirely inside of source range,
                    # so map both endpoints according to the same map
                    next_ranges.append((start + delta, stop + delta))
                    break
                if start < src <= stop <= src_upper:
                    # Ranges overlap, starting with current range, so
                    # map the relevant subset of the source range
                    next_ranges.append((src + delta, stop + delta))
                    # Map the remainder in subsequent iteration(s)
                    ranges.append((start, src - 1))
                    break
                if src <= start <= src_upper < stop:
                    # Ranges overlap, starting with source range, so
                    # map the relevant subset of the source range
                    next_ranges.append((start + delta, src_upper + delta))
                    # Map the remainder in subsequent iteration(s)
                    ranges.append((src_upper + 1, stop))
                    break
            else:
                # If the current range does not intersect with any source ranges,
                # simply map the range to itself
                next_ranges.append((start, stop))

        ranges = next_ranges

    return ranges

def solve_p1():
    """
    Solve part 1.
    """
    seeds = map(int, inputs[0].split()[1:])
    return min(map(lambda x : seed_to_location(x, build_maps()), seeds))

def solve_p2():
    """
    Solve part 2.
    """
    seed_range_data = list(map(int, inputs[0].split()[1:]))
    seed_ranges = [(seed_range_data[i], seed_range_data[i] + seed_range_data[i + 1])
                   for i in range(0, len(seed_range_data), 2)]
    return min(seed_ranges_to_loc_ranges(seed_ranges, build_maps()))[0]

print('Part 1: Lowest location number =', solve_p1())
print('Part 2: Lowest location number =', solve_p2())
