package data

import "log"

type FeatureStorage interface {
	GetAllFeatures(pageSize int, page int) (error, Features)
}

type FeatureStorageInMemory struct{}

func (storage *FeatureStorageInMemory) GetAllFeatures(pageSize int, page int) (error, Features) {
	log.Printf("fetching features by pageSize: %d and page %d\n", pageSize, page)
	return nil, SampleData
}