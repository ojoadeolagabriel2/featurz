package data

import "github.com/google/uuid"

const (
	FeatureDisabled   = 0
	FeatureRollingOut = 1
	FeatureRollout    = 2
)

const (
	Percent = "Percent"
	Count   = "Count"
)

type Feature struct {
	Id           string   `json:"id"`
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	Slug         string   `json:"slug"`
	RolloutType  string   `json:"rollout_type"`
	RolloutValue float64  `json:"rollout_value"`
	Users        []string `json:"users"`
	Groups       []string `json:"groups"`
	State        int      `json:"state"`
	Tags         []string `json:"tags"`
}

type Features []Feature

var SampleData Features = []Feature{
	{
		Id:           uuid.New().String(),
		Name:         "view_card_details",
		Description:  "view card details",
		Slug:         "view_card_details",
		RolloutType:  Percent,
		RolloutValue: 45.5,
		Groups:       []string{"payments", "transfers"},
		State:        FeatureRollout,
		Users:        []string{"10001", "10002"},
		Tags:         []string{"cards", "payments"},
	},
	{
		Id:           uuid.New().String(),
		Name:         "cache_issuer",
		Description:  "mapped to issuer functionality",
		Slug:         "cache_issuer",
		RolloutType:  Count,
		RolloutValue: 90,
		Groups:       []string{"cache", "performance"},
		State:        FeatureRollingOut,
		Users:        []string{"10003"},
		Tags:         []string{"infrastructure", "platform"},
	},
	{
		Id:           uuid.New().String(),
		Name:         "async_mailing",
		Description:  "manage async mailing notification",
		Slug:         "async_mailing",
		RolloutType:  Percent,
		RolloutValue: 10,
		Groups:       []string{"mail", "async"},
		State:        FeatureRollingOut,
		Users:        []string{"10013"},
		Tags:         []string{"platform"},
	},
	{
		Id:           uuid.New().String(),
		Name:         "acquirer_switching",
		Description:  "mapped to issuer functionality",
		Slug:         "cache_issuer",
		RolloutType:  Count,
		RolloutValue: 90,
		Groups:       []string{"cache", "performance"},
		State:        FeatureRollingOut,
		Users:        []string{"10003"},
		Tags:         []string{"infrastructure", "platform"},
	},
}
