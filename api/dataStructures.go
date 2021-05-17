package main

type root struct {
	Records []recordType `json:"records"`
}

type recordType struct {
	Event eventType `json:"event"`
	Setup setupType `json:"setup"`
	Time  int       `json:"time"`
}

type eventType struct {
	T string `json:"type"`
}

type setupType struct {
	URL      string `json:"url"`
	NodeName string `json:"nodeName"`
}

type toReturn struct {
	Records []item `json:"records"`
}

type item struct {
	ID        int    `json:"id"`
	EventType string `json:"eventType"`
	Time      int    `json:"time"`
	HTMLtag   string `json:"htmlTag"`
}

type idStruct struct {
	ID int `json:"id"`
}

type statistics struct {
	CountDifferentInteractions       int     `json:"countDifferentInteractions"`
	CalculateTotalTimeOfInteractions int     `json:"calculateTotalTimeOfInteractions"`
	CalculateLongestSequenceOfInput  int     `json:"calculateLongestSequenceOfInput"`
	CalculateLongestSequenceType     string  `json:"calculateLongestSequenceType"`
	CalculateLongestSequenceCounter  int     `json:"calculateLongestSequenceCounter"`
	CounterTotalInteractions         int     `json:"counterTotalInteractions"`
	MinDelayBetweenInteractions      float32 `json:"minDelayBetweenInteractions"`
	MaxDelayBetweenInteractions      float32 `json:"maxDelayBetweenInteractions"`
	MeanDelayBetweenInteractions     float32 `json:"meanDelayBetweenInteractions"`
}
