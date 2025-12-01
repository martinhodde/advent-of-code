package registry

type Solver func([]string) string

var table = map[int]map[int]Solver{}

func Register(day, part int, fn Solver) {
	if table[day] == nil {
		table[day] = map[int]Solver{}
	}
	table[day][part] = fn
}

func Lookup(day, part int) Solver {
	if table[day] == nil {
		return nil
	}
	return table[day][part]
}
