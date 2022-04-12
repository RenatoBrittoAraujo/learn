package gof

func getEmptyTable(w, h int) *Table {
	table := make(Table, h)
	for i := range table {
		table[i] = make([]bool, w)
	}
	return &table
}
