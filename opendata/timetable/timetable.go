package timetable

import (
  "../../../anctgo"
  "time"
)

type Timetable struct {
  Common
  Lectures []Lecture
}

type Common struct {
  Institution string
  Year        int
  Term        string
}

type Lecture struct {
  Name       string
  Grade      int
  Department anctgo.Department
  Wday       time.Weekday
  StartTime  time.Time
  EndTime    time.Time
  Location   string
  Lecturers  []string
  URI        string
}

// 指定した週のTimetableを取り出す。
// Note: このメソッドはO(n)である。n = len(t.Lectures)
func (t *Timetable) OfWday(wday time.Weekday) (wt Timetable) {
  wt.Common = t.Common
  for _, l := range t.Lectures {
    if l.Wday == wday {
      wt.Lectures = append(wt.Lectures, l)
    }
  }
  return
}

// 指定したクラスのTimetableを取り出す。
// NOTE: このメソッドはO(n)である。 n = len(t.Lectures)
func (t *Timetable) OfClass(grade int, dep anctgo.Department) (wt Timetable) {
  wt.Common = t.Common
  for _, l := range t.Lectures {
    if l.Grade == grade && l.Department == dep {
      wt.Lectures = append(wt.Lectures, l)
    }
  }
  return
}
