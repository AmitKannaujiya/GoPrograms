package dictionarysystem

import (
	"fmt"
	"sync"
)

var once sync.Once
var singleDatasetManager *DictionaryManager

type DictionaryManager struct {
	datasets map[string]*Dataset
	mutex     sync.RWMutex
}

func GetDictionaryManager() *DictionaryManager {
	if singleDatasetManager == nil {
		once.Do(func() {
			singleDatasetManager = &DictionaryManager{
				datasets: make(map[string]*Dataset),
			}
		})
	}
	return singleDatasetManager
}

// create dataset
func (dm *DictionaryManager) CreateDataset(datasetName string) error {
	dm.mutex.Lock()
	defer dm.mutex.Unlock()
	if _, exists := dm.datasets[datasetName]; exists {
		return fmt.Errorf("Dataset : %s already exists", datasetName)
	}
	dm.datasets[datasetName] = NewDataset(datasetName)
	return nil
}

// delete dataset
func (dm *DictionaryManager) DeleteDataset(datasetName string) error {
	dm.mutex.Lock()
	defer dm.mutex.Unlock()
	if _, exists := dm.datasets[datasetName]; !exists {
		return fmt.Errorf("Dataset : %s does not exists", datasetName)
	}
	delete(dm.datasets, datasetName)
	return nil
}

// add word to dataset

func (dm *DictionaryManager) AddWordToDataset(datasetName, word string) error {
	dm.mutex.Lock()
	defer dm.mutex.Unlock()
	dataset, exists := dm.datasets[datasetName];
	if !exists {
		return fmt.Errorf("Dataset : %s does not exists", datasetName)
	}
	dataset.AddWord(word)
	return nil
}

// Search Word 
func (dm *DictionaryManager) SearchWord(word string) ([]string, error) {
	dm.mutex.RLock()
	defer dm.mutex.RUnlock()
	var result []string

	for _, dataset := range dm.datasets {
		if dataset.IsPresent(word) {
			result = append(result, word)
		}
	}
	if len(result) == 0 {
		return nil, fmt.Errorf("Word : %s not present in dictionary", word)
	}
	return result, nil
}

func (dm *DictionaryManager) ListDatasets() []string {
	dm.mutex.RLock()
	defer dm.mutex.RUnlock()
	var result []string
	for datasetName, _ := range dm.datasets {
		result = append(result, datasetName)
	}
	return result
}
