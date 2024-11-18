package inmemoryrdbms

import (
	"fmt"
	"sync"
)

var once sync.Once
var singleDatabaseManager *DataBaseManager
type DataBaseManager struct {
	databases map[string]*Database
	activeDB  *Database
	mutex     sync.RWMutex
}

func GetDatabaseManager() *DataBaseManager {
	if singleDatabaseManager == nil {
		once.Do(func() {
			singleDatabaseManager = &DataBaseManager{
				databases: make(map[string]*Database),
			}
		})
	}
	return singleDatabaseManager
}

// create database
func (dm *DataBaseManager) CreateDatabase(databaseName string) error {
	dm.mutex.Lock()
	defer dm.mutex.Unlock()
	if _, exists := dm.databases[databaseName]; exists {
		return fmt.Errorf("database : %s already exists", databaseName)
	}
	dm.databases[databaseName] = NewDatabase(databaseName)
	return nil
}
// delete database 
func (dm *DataBaseManager) DeleteDatabase(databaseName string) error {
	dm.mutex.Lock()
	defer dm.mutex.Unlock()
	if _, exists := dm.databases[databaseName]; !exists {
		return fmt.Errorf("database : %s does not exists", databaseName)
	}
	delete(dm.databases, databaseName)
	return nil
}
// switch database
func (dm *DataBaseManager) SwitchDatabase(databaseName string) error {
	dm.mutex.Lock()
	defer dm.mutex.Unlock()
	database, exists := dm.databases[databaseName]
	if !exists {
		return fmt.Errorf("database : %s does not exists", databaseName)
	}
	dm.activeDB = database
	return nil
}
