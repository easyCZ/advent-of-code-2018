package main

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
	"time"
)

const (
	shiftStart = iota
	fallAsleep
	wakeUp
)

var (
	logRegex        = regexp.MustCompile(`\[(?P<Date>\d{4}-\d{2}\-\d{2} \d{2}:\d{2})\] (?P<Message>.*)`)
	shiftStartRegex = regexp.MustCompile(`Guard #(?P<Id>\d+) begins shift`)
	fallAsleepRegex = regexp.MustCompile(`falls asleep`)
	wakeUpRegex     = regexp.MustCompile(`wakes up`)

	dateLayout = "2006-01-02 15:04"
)

func max(vals []int) int {
	m := math.MinInt32
	for i := range vals {
		if m < vals[i] {
			m = vals[i]
		}
	}
	return m
}

type Log struct {
	ts      time.Time
	message string
}

func (l *Log) String() string {
	return fmt.Sprintf("%v %s", l.ts, l.message)
}

type Guard struct {
	id     int64
	sleeps []*Interval
	total  int
}

func (g *Guard) totalSlept() int {
	c := 0
	for _, interval := range g.sleeps {
		from := interval.from.Minute()
		to := interval.to.Minute()
		if to < from {
			to = 60 - from + to
		}
		for i := from; i < to; i++ {
			c += 1
		}
	}
	g.total = c
	return c
}

func (g *Guard) mostSleptMinute() (value int, minute int) {
	var minutes [60]int
	for _, interval := range g.sleeps {
		from := interval.from.Minute()
		to := interval.to.Minute()
		if to < from {
			to = 60 - from + to
		}
		for i := from; i < to; i++ {
			minutes[i%60] += 1
		}
	}

	maxIndex := 0
	for i, val := range minutes {
		if minutes[maxIndex] < val {
			maxIndex = i
		}
	}

	return minutes[maxIndex], maxIndex
}

type Interval struct {
	from time.Time
	to   time.Time
}

func parseShiftStart(s string) (int64, bool) {
	if !shiftStartRegex.Match([]byte(s)) {
		return 0, false
	}
	matches := shiftStartRegex.FindStringSubmatch(s)
	id, err := strconv.ParseInt(matches[1], 10, 32)
	if err != nil {
		return 0, false
	}
	return id, true
}

func buildSleepSchedule(logs []*Log) ([]*Guard, error) {
	i := 0
	guards := make(map[int64]*Guard)

	var guard *Guard
	for i < len(logs) {
		if id, ok := parseShiftStart(logs[i].message); ok {
			guard, ok = guards[id]
			if !ok {
				guard = &Guard{
					id:     id,
					sleeps: make([]*Interval, 0),
				}
			}
		}

		if fallAsleepRegex.Match([]byte(logs[i].message)) {
			from := logs[i].ts
			to := logs[i+1].ts
			interval := &Interval{from, to}
			guard.sleeps = append(guard.sleeps, interval)
			guards[guard.id] = guard
			i += 1
		}

		i += 1
	}

	grds := make([]*Guard, 0)
	for _, g := range guards {
		grds = append(grds, g)
	}

	return grds, nil
}

func findSleepyGuard(guards []*Guard) *Guard {
	var g *Guard
	var max int

	for _, guard := range guards {
		slept := guard.totalSlept()
		if max < slept {
			max = slept
			g = guard
		}
	}

	return g
}

func guardWithMaxMinutesSlept(guards []*Guard) (g *Guard, maxSleptMinute int, maxSleptValue int) {
	maxSleptValue = 0
	maxSleptMinute = 0
	for _, guard := range guards {
		value, minute := guard.mostSleptMinute()
		if maxSleptValue < value {
			g = guard
			maxSleptValue = value
			maxSleptMinute = minute
		}
	}

	return g, maxSleptMinute, maxSleptValue
}

func parse(r io.Reader) ([]*Log, error) {
	scanner := bufio.NewScanner(r)

	logs := make([]*Log, 0)
	for scanner.Scan() {
		t := scanner.Text()
		matches := logRegex.FindStringSubmatch(t)
		ts, err := time.Parse(dateLayout, matches[1])
		if err != nil {
			return nil, errors.Wrapf(err, "failed to parse date from %s", t)
		}
		logs = append(logs, &Log{
			message: matches[2],
			ts:      ts,
		})
	}

	sort.Slice(logs, func(i, j int) bool {
		return logs[i].ts.Before(logs[j].ts)
	})

	return logs, nil
}

func main() {
	logs, err := parse(os.Stdin)
	if err != nil {
		panic(err)
	}

	guards, err := buildSleepSchedule(logs)
	if err != nil {
		panic(err)
	}

	sleeper := findSleepyGuard(guards)
	_, favouriteMinute := sleeper.mostSleptMinute()
	fmt.Println(fmt.Sprintf("Part one: #%v, minute %v => %v", sleeper.id, favouriteMinute, sleeper.id*int64(favouriteMinute)))

	g, minute, _ := guardWithMaxMinutesSlept(guards)
	fmt.Println(fmt.Sprintf("Part two: #%v, minute %v => %v", g.id, minute, g.id*int64(minute)))
}
