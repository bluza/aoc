package main

type Field struct {
	rows [][]string
}

func (f Field) String() string {
	var repr string
	for _, row := range f.rows {
		for _, c := range row {
			repr += c
		}
		repr += "\n"
	}
	return repr
}
