package data

const (
	FeatureDisabled   = 0
	FeatureRollingOut = 1
	FeatureRollout    = 2
)

type Feature struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Group       string `json:"group"`
	State       int    `json:"state"`
}

type TestFeatures []Feature
