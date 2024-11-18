package inmemoryrdbms

import (
	"errors"
	"sync"
)

type Table struct {
	autoIncreamentedId int
	tableName string
	schema *Schema
	rows []Row
	mutex     sync.RWMutex
}

func NewTable(tableName string, schema *Schema) *Table {
	return &Table{
		autoIncreamentedId: 1,
		tableName: tableName,
		schema: schema,
		rows: make([]Row, 0),
	}
}

func(t *Table) Truncate() {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.rows = make([]Row, 0)
}
// insert
func(t *Table) Insert(row Row) error {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	// check constrains
	if err := t.schema.ValidateRow(row); err != nil {
		return err
	}

	// check unique 

	for _ , currentRow := range t.rows {
		if t.schema.IsDuplicate(currentRow, row) {
			return errors.New("Duplicate row, unique constraints violated")
		}
	}
	row.values[0] = t.autoIncreamentedId // 0 column should be id
	t.autoIncreamentedId++
	t.rows = append(t.rows, row)
	return nil
}
// delete
func(t *Table) Delete(condition func(Row) bool) int {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	deleteCount := 0
	rows := make([]Row, 0)
	for _ , currentRow := range t.rows {
		if condition(currentRow) {
			deleteCount++
		} else {
			rows = append(rows, currentRow)
		}
	}
	t.rows = rows
	return deleteCount
}
// select 
func(t *Table) Select(condition func(Row) bool) []Row {
	t.mutex.RLock()
	defer t.mutex.RUnlock()

	var rows []Row

	for _ , currentRow := range t.rows {
		if condition(currentRow) {
			rows = append(rows, currentRow)
		}
	}
	return rows
}

// update 

func(t *Table) Update(condition func(Row) bool, updateFunc func(*Row)) int {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	updateCount :=0 
	for i := range t.rows {
		if condition(t.rows[i]) {
			updateFunc(&t.rows[i])
			updateCount++
		}
	}
	return updateCount
}