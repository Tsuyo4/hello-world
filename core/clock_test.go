package core

import "testing"

type performance struct {
	deltaTime int64
}

func (p *performance) Next( delta int64) {
	p.deltaTime += delta
}

func (p *performance) Now() int64 {
	return p.deltaTime;
}


func TestClockWithPerformance( t *testing.T) {
	selfPerformance := &performance{};

	callDateNow = func() int64 {
		return selfPerformance.Now()
	}

	clock := &Clock{}

	clock.Start()

	selfPerformance.Next(123)
	if clock.GetElapsedTime() != 0.123 {
		t.Error(`getElapsedTime Error 0.123`)
	}

	selfPerformance.Next(100)
	if clock.GetElapsedTime() != 0.223 {
		t.Error(`getElapsedTime Error 0.223`)
	}
	
	clock.Stop()

	selfPerformance.Next(1000)
	if clock.GetElapsedTime() != 0.223 {
		t.Error(`don't update time if the clock was stopped`)
	}	
	
}
