package elasticex

import (
	"github.com/olivere/elastic"
)

// TermsAggregation terms aggregation struct
type TermsAggregation struct {
	Fields         []string
	RankingSize    int
	SumAggregation *SumAggregation
}

// Build build query
func (c *TermsAggregation) Build() (string, *elastic.TermsAggregation) {
	aggs := make([]*elastic.TermsAggregation, 0)
	// Iterate fields
	for i, field := range c.Fields {
		aggs = append(aggs, elastic.NewTermsAggregation())
		// Set field
		aggs[i] = aggs[i].Field(field)
		// Set ranking size
		if c.RankingSize > 0 {
			aggs[i] = aggs[i].Size(c.RankingSize)
		}
		// Set order
		if c.SumAggregation == nil {
			aggs[i] = aggs[i].OrderByCount(false)
		} else {
			aggs[i] = aggs[i].OrderByAggregation(c.SumAggregation.Field, false)
		}
		// Set sum aggsregation
		if c.SumAggregation != nil {
			aggs[i] = aggs[i].SubAggregation(c.SumAggregation.Field, c.SumAggregation.Build())
		}
		if i > 0 {
			aggs[i-1] = aggs[i-1].SubAggregation(field, aggs[i])
		}
	}

	return c.Fields[0], aggs[0]
}
