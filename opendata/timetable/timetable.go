package timetable

import (
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
  Department Department
  Wday       time.Weekday
  StartTime  time.Time
  EndTime    time.Time
  Location   string
  Lecturers  []string
  URI        string
}

type Department string

const (
  Mechanical              Department = "機械工学科"
  Electrical                         = "電気情報工学科"
  Civil                              = "都市システム工学科"
  Architecture                       = "建築学科"
  MechanicalAndElectronic            = "機械・電子システム工学専攻"
  ArchitectureAndCivil               = "建築・都市システム工学専攻"
)

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
func (t *Timetable) OfClass(grade int, dep Department) (wt Timetable) {
  wt.Common = t.Common
  for _, l := range t.Lectures {
    if l.Grade == grade && l.Department == dep {
      wt.Lectures = append(wt.Lectures, l)
    }
  }
  return
}
