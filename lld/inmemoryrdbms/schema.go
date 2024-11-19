package inmemoryrdbms

import "fmt"

type Schema struct {
	columns       []string
	types         []string
	uniqueColumns map[string]bool
}

func NewSchema(columns []string, types []string, uniqueColumns []string) *Schema {
	uniqueColumn := make(map[string]bool)
	for _, uniqueCol := range uniqueColumns {
		uniqueColumn[uniqueCol] = true
	}
	return &Schema{
		columns:       columns,
		types:         types,
		uniqueColumns: uniqueColumn,
	}
}

func (s *Schema) ValidateRow(row Row) error {
	if len(row.values) != len(s.columns) {
		return fmt.Errorf("No of columns does not match")
	}
	for i, col := range row.values {
		if !s.IsTypeValid(s.types[i], col) {
			return fmt.Errorf("column type %v mismatch %s ", col, s.types[i])
		}
	}
	return nil
}

func (s *Schema) IsTypeValid(types string, val interface{}) bool {
	switch types {
	case "INT":
		_, ok := val.(int)
		return ok
	case "STRING":
		_, ok := val.(string)
		return ok
	default:
		return false
	}
}

func (s *Schema) IsDuplicate(currentRow, row Row) bool {
	for col, uniq := range s.uniqueColumns {
		if uniq {
			colIndex := s.GetColIndex(col)
			if colIndex != -1 && currentRow.values[colIndex] == row.values[colIndex] {
				return true
			}
		}
	}
	return false
}

func (s *Schema) GetColIndex(col string) int {
	for i, column := range s.columns {
		if column == col {
			return i
		}
	}
	return -1
}
