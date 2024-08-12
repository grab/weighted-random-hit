// Copyright 2024 Grabtaxi Holdings Pte Ltd (GRAB), All rights reserved.

package weightedrandomhit

import (
	"math/rand"
)

type Options struct {
	defaultWeight    int
	chanceMultiplier int
}

type OptionsFunc func(*Options)

// IsCategoryHit checks if a category hit/miss based on category weights
func IsCategoryHit(
	category string,
	categoryWeights map[string]int,
	targetHit int,
	maxAllowedHitsForEachCategory int,
	opts ...OptionsFunc) bool {

	// set default options
	options := Options{
		defaultWeight:    0,
		chanceMultiplier: 1,
	}

	// override default options
	for _, opt := range opts {
		opt(&options)
	}

	var weight, maxWeight, totalWeight int

	for _, weight = range categoryWeights {
		totalWeight += weight
		if weight > maxWeight {
			maxWeight = weight
		}
	}

	if val, exists := categoryWeights[category]; exists {
		weight = val
	} else {
		if options.defaultWeight == 0 {
			// if no default weight is set, unknown categories will be ignored
			return false
		}
		weight = options.defaultWeight
	}

	adjustedTarget := targetHit * (maxWeight / weight)
	if adjustedTarget > maxAllowedHitsForEachCategory {
		adjustedTarget = maxAllowedHitsForEachCategory
	}

	if rand.Intn(totalWeight)%totalWeight <= adjustedTarget*options.chanceMultiplier {
		return true
	}

	return false
}

// WithChanceMultiplier use this to multiply the chance of category hit. Useful for testing
func WithChanceMultiplier(multiplier int) OptionsFunc {
	return func(options *Options) {
		options.chanceMultiplier = multiplier
	}
}

// WithDefaultWeight use this to have a default weight for unknown categories
func WithDefaultWeight(defaultWeight int) OptionsFunc {
	return func(options *Options) {
		options.defaultWeight = defaultWeight
	}
}
