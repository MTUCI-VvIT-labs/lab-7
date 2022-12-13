package entities

import "time"

type Week struct {
	WeekNumber int
	WeekType   string
}

func NewWeek() Week {
	var week Week
	week.calculateWeek()
	week.calculateWeekType()
	return week
}

func (w *Week) calculateWeek() {
	firstSeptember := time.Date(2022, 9, 1, 0, 0, 0, 0, time.UTC)
	today := time.Now().AddDate(0, 0, 7)

	_, weekFromYear := today.ISOWeek()
	_, weekFromFirstSeptember := firstSeptember.ISOWeek()

	w.WeekNumber = weekFromYear - weekFromFirstSeptember
}

func (w *Week) calculateWeekType() {
	if w.WeekNumber%2 == 0 {
		w.WeekType = "чётная"
	} else {
		w.WeekType = "нечётная"
	}
}
