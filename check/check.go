package check

import (
	"strconv"
	"time"
)

// PerfData holds the Icinga/Nagios format Performance Data
type PerfData struct {
	timer  time.Time
	String string
}

// StartTimer adds a new key/value Performance Data metric
func (p *PerfData) StartTimer(t time.Time) {
	p.timer = t
}

// Add adds a new key/value Performance Data metric
func (p *PerfData) Add(key string, value string, uom string) {
	if p.String == "" {
		p.String += "|"
	} else {
		p.String += " "
	}
	p.String += "'" + key + "'=" + value + uom
}

// Get returns all the key/value Performance Data metrics
// Stops the timer and adds that to the key/value pairs.
func (p *PerfData) Get() string {
	duration := time.Since(p.timer)
	durationstr := strconv.Itoa(int(duration.Milliseconds()))

	if p.String == "" {
		p.String += "|"
	} else {
		p.String += " "
	}

	return p.String + "'checks_took'=" + durationstr + "ms"
}
