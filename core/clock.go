package core

import "time"

type Clock struct {
	autoStart   bool
	startTime   int64
	oldTime     int64
	elapsedTime float64
	running     bool
}

func ClockNew(autoStart bool) *Clock {
	clock := &Clock{}
	clock.autoStart = autoStart

	return clock
}

var callDateNow = func() int64 {
	return time.Now().Unix()
}

func (clock *Clock) Start() {

	clock.startTime = callDateNow()
	clock.oldTime = clock.startTime
	clock.running  = true	
	
}

func (clock *Clock) Stop() {
	clock.GetElapsedTime()
	clock.running = false
}


func (clock *Clock) GetElapsedTime() float64 {

	clock.GetDelta()
	return clock.elapsedTime
	
}

func (clock *Clock) GetDelta() float64 {
	diff := 0.0

	if clock.autoStart && !clock.running {
		clock.Start()
	}

	if clock.running {

		newTime := callDateNow()

		diff = (float64)( newTime - clock.oldTime ) / 1000.0;

		clock.oldTime = newTime

		clock.elapsedTime += diff;
	}

	return diff
}
