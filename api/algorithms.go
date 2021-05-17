package main

import (
	"fmt"
	"log"
	"strconv"
)

func countDifferentInteractions(c chan bool, stats *statistics) {

	differentInteractions := make(map[string]bool)
	counterDifferents := 0
	stats.CounterTotalInteractions = len(toRet.Records)

	//Linear in time algorithm O(n) <====
	for _, v := range toRet.Records {
		if _, ok := differentInteractions[v.EventType]; ok {
		} else {
			differentInteractions[v.EventType] = true
			counterDifferents++
		}
	}
	stats.CountDifferentInteractions = counterDifferents
	log.Println("[countDifferentInteractions] number of different interactions is: " + strconv.Itoa(counterDifferents))
	c <- true

	/*
		*********ALTERNATIVE WITH DB INTERACTION**********

		SELECT COUNT (DISTINCT event_type) AS countDifferents
		FROM public.records;

	*/

}
func calculateTotalTimeOfInteractions(c chan bool, stats *statistics) {
	const MaxUint = ^uint(0)
	const MinUint = 0
	const MaxInt = int(MaxUint >> 1)
	const MinInt = -MaxInt - 1

	minTstamp := MaxInt
	maxTstamp := 0

	//Linear in time algorithm O(n) <====
	for _, v := range toRet.Records {
		if v.Time > maxTstamp {
			maxTstamp = v.Time
		}
		if v.Time < minTstamp {
			minTstamp = v.Time
		}
	}

	stats.CalculateTotalTimeOfInteractions = (maxTstamp - minTstamp) / 1000 //Here I'm returning the value in seconds
	log.Println("[calculateTotalTimeOfInteractions] total time of interaction is " + strconv.Itoa(stats.CalculateTotalTimeOfInteractions) + " seconds")
	c <- true

	/*
		*********ALTERNATIVE WITH DB INTERACTION**********

		SELECT MAX(time) - MIN(time) AS totalTimeOfInteraction
		FROM public.records;

	*/
}

func calculateLongestSequences(c chan bool, stats *statistics) {

	//I filter focus event as asked in the requirements
	tmp := make([]item, 0)
	for _, v := range toRet.Records {
		if v.EventType != "focus" {
			tmp = append(tmp, v)
		}
	}
	maxEventCounter := 0
	maxEventType := ""
	eventCounter := 0
	eventType := ""
	maxInputeventCounter := 0

	//Linear in time algorithm O(n) <====
	for _, v := range tmp {
		if v.EventType == eventType {
			//increment the sequence counter
			eventCounter++
		} else {
			//stop sequence, check max
			if eventCounter > maxEventCounter {
				maxEventCounter = eventCounter
				maxEventType = eventType
			}
			//In addition I check and save the longest input sequence
			if eventType == "input" && eventCounter > maxInputeventCounter {
				maxInputeventCounter = eventCounter
			}
			eventCounter = 1
			eventType = v.EventType
		}
	}
	log.Println("[calculateLongestSequences] max event type is '" + maxEventType + "' with length equal to  " + strconv.Itoa(maxEventCounter))
	log.Println("[calculateLongestSequences] longest INPUT sequence is: " + strconv.Itoa(maxInputeventCounter))

	stats.CalculateLongestSequenceOfInput = maxInputeventCounter
	stats.CalculateLongestSequenceCounter = maxEventCounter
	stats.CalculateLongestSequenceType = maxEventType
	c <- true
}

func calculateMinMaxMeanTimeBetweenInteractions(c chan bool, stats *statistics) {
	const MaxUint = ^uint(0)
	const MinUint = 0
	const MaxInt = int(MaxUint >> 1)
	const MinInt = -MaxInt - 1

	minDelay := MaxInt
	maxDelay := 0
	currentDelay := 0
	totDelayTMP := 0
	counter := 0

	//Linear in time algorithm O(n) <====
	for i := 0; i < len(toRet.Records)-1; i++ {
		currentDelay = toRet.Records[i+1].Time - toRet.Records[i].Time //Since they are ordered by time this delay is the right value I'm looking for
		if currentDelay < minDelay {
			minDelay = currentDelay
		}
		if currentDelay > maxDelay {
			maxDelay = currentDelay
		}
		totDelayTMP += currentDelay
		counter++
	}

	stats.MaxDelayBetweenInteractions = float32(maxDelay) / 1000
	stats.MinDelayBetweenInteractions = float32(minDelay) / 1000
	stats.MeanDelayBetweenInteractions = (float32(totDelayTMP) / float32(counter) / 1000) //Here I'm calculating the value in seconds
	log.Println("[calculateMinMaxMeanTimeBetweenInteractions] min time between interactions is: " + fmt.Sprintf("%f", stats.MaxDelayBetweenInteractions) + " seconds")
	log.Println("[calculateMinMaxMeanTimeBetweenInteractions] max time between interactions is: " + fmt.Sprintf("%f", stats.MinDelayBetweenInteractions) + " seconds")
	log.Println("[calculateMinMaxMeanTimeBetweenInteractions] mean time between interactions is: " + fmt.Sprintf("%f", stats.MeanDelayBetweenInteractions) + " seconds")
	c <- true

}
