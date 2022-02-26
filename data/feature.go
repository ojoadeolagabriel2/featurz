package data

import "github.com/google/uuid"

const (
	FeatureDisabled   = 0
	FeatureRollingOut = 1
	FeatureRollout    = 2
)

type Feature struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Slug        string `json:"slug"`
	Group       string `json:"group"`
	State       int    `json:"state"`
}

type TestFeatures []Feature

var SampleData TestFeatures = []Feature{
	{
		Id:          uuid.New().String(),
		Name:        "view_card_details",
		Description: "view card details",
		Slug:        "view_card_details",
		Group:       "payments",
		State: FeatureRollout,
	},
	{
		Id:          uuid.New().String(),
		Name:        "cache_issuer",
		Description: "mapped to issuer functionality",
		Slug:        "cache_issuer",
		Group:       "infrastructure",
		State: FeatureRollingOut,
	},
}
