/*
 * NamSor API v2
 *
 * NamSor API v2 : enpoints to process personal names (gender, cultural origin or ethnicity) in all alphabets or languages. Use GET methods for small tests, but prefer POST methods for higher throughput (batch processing of up to 100 names at a time). Need something you can't find here? We have many more features coming soon. Let us know, we'll do our best to add it! 
 *
 * API version: 2.0.13
 * Contact: contact@namsor.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package namsorapi

// Simple metrics on a classifier
type ClassifierMetricsOut struct {
	SoftwareVersion string `json:"softwareVersion,omitempty"`
	HostAddress string `json:"hostAddress,omitempty"`
	LearnQueueSize int32 `json:"learnQueueSize,omitempty"`
	BufferSize int32 `json:"bufferSize,omitempty"`
	PreClassifyQueueSize int32 `json:"preClassifyQueueSize,omitempty"`
	FactKeysSize int32 `json:"factKeysSize,omitempty"`
	FactsLearned int64 `json:"factsLearned,omitempty"`
	ClassifyDurationsCurrent float64 `json:"classifyDurationsCurrent,omitempty"`
	ClassifyDurationsSummary float64 `json:"classifyDurationsSummary,omitempty"`
	LearnDurationsCurrent float64 `json:"learnDurationsCurrent,omitempty"`
	LearnDurationsSummary float64 `json:"learnDurationsSummary,omitempty"`
	ClassifierName string `json:"classifierName,omitempty"`
	FeaturesSize int64 `json:"featuresSize,omitempty"`
	AiVettedEstimateTotal int64 `json:"aiVettedEstimateTotal,omitempty"`
	AiVettedEstimatePrecision float64 `json:"aiVettedEstimatePrecision,omitempty"`
	AiVettedEstimateRecall float64 `json:"aiVettedEstimateRecall,omitempty"`
	AiVettedLearnTotal int64 `json:"aiVettedLearnTotal,omitempty"`
	AiNonVettedEstimateTotal int64 `json:"aiNonVettedEstimateTotal,omitempty"`
	AiNonVettedEstimatePrecision float64 `json:"aiNonVettedEstimatePrecision,omitempty"`
	AiNonVettedEstimateRecall float64 `json:"aiNonVettedEstimateRecall,omitempty"`
	AiNonVettedLearnTotal int64 `json:"aiNonVettedLearnTotal,omitempty"`
	MetricTimeStamp int64 `json:"metricTimeStamp,omitempty"`
	AiStartTime int64 `json:"aiStartTime,omitempty"`
	AiVettedExpectedClassMetrics []ExpectedClassMetricsOut `json:"aiVettedExpectedClassMetrics,omitempty"`
	AiNonVettedExpectedClassMetrics []ExpectedClassMetricsOut `json:"aiNonVettedExpectedClassMetrics,omitempty"`
}
