package inmemoryrdbms

import (
	"fmt"
	"sync"
)

type Database struct {
	databaseName string
	tables       map[string]*Table
	mutex        sync.RWMutex
}

func NewDatabase(databaseName string) *Database {
	return &Database{
		databaseName: databaseName,
		tables:       make(map[string]*Table),
	}
}

// create table

func (d *Database) CreateTable(tableName string, schema *Schema) error {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	if _, exists := d.tables[tableName]; exists {
		return fmt.Errorf("table : %s already exists", tableName)
	}
	d.tables[tableName] = NewTable(tableName, schema)
	return nil
}

// drop table

func (d *Database) DropTable(tableName string) error {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	if _, exists := d.tables[tableName]; !exists {
		return fmt.Errorf("table : %s does not exists", tableName)
	}
	delete(d.tables, tableName)
	return nil
}

// truncate table

func (d *Database) TruncateTable(tableName string) error {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	table, exists := d.tables[tableName]
	if !exists {
		return fmt.Errorf("table : %s does not exists", tableName)
	}
	table.Truncate()
	return nil
}
