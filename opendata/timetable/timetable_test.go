package timetable

import (
  "testing"
  "time"
)

var ofWdayTests = []struct {
  timetable, out Timetable
  in             time.Weekday
}{
  {
    timetable: Timetable{
      Common{Institution: "WdayTest", Term: "前期", Year: 2013},
      []Lecture{
        Lecture{
          Name:       "施設管理工学I",
          Grade:      2,
          Department: Civil,
          Location:   "2C教室",
          Wday:       time.Monday,
        },
        Lecture{
          Name:       "データ構造とアルゴリズム",
          Grade:      2,
          Department: Civil,
          Location:   "2C教室",
          Wday:       time.Tuesday,
        },
      },
    },
    out: Timetable{
      Common{Institution: "WdayTest", Term: "前期", Year: 2013},
      []Lecture{
        Lecture{
          Name:       "施設管理工学I",
          Grade:      2,
          Department: Civil,
          Location:   "2C教室",
          Wday:       time.Monday,
        },
      },
    },
    in: time.Monday,
  },
}

var ofClassTests = []struct {
  timetable, out Timetable
  grade          int
  dep            Department
}{
  {
    timetable: Timetable{
      Common{Institution: "ClassTest", Term: "前期", Year: 2013},
      []Lecture{
        Lecture{
          Name:       "施設管理工学I",
          Grade:      2,
          Department: Civil,
          Location:   "2C教室",
          Wday:       time.Monday,
        },
        Lecture{
          Name:       "データ構造とアルゴリズム",
          Grade:      1,
          Department: Electrical,
          Location:   "1E教室",
          Wday:       time.Tuesday,
        },
      },
    },
    out: Timetable{
      Common{Institution: "ClassTest", Term: "前期", Year: 2013},
      []Lecture{
        Lecture{
          Name:       "施設管理工学I",
          Grade:      2,
          Department: Civil,
          Location:   "2C教室",
          Wday:       time.Monday,
        },
      },
    },
    grade: 2,
    dep:   Civil,
  },
}

// テストに合格すればtrueを返す。
// NOTE: 実際には、aとbが等しければtrueを返す。
//       しかし、あくまでもテストに合格したかどうかを確認する目的で使うために作ったため、Equal関数ではない。
// NOTE: 計算量はO(mn)です。ここでは、nはlen(a), mはlen(a.Lectures)です。
//       O(1)ではないことに注意。
// TODO: 即席で適当に書いたので、汚い。
//       時間があればもう少し綺麗に書きなおす。
func isPassedTest(a, b Timetable) bool {
  if a.Common != b.Common {
    return false
  }
  if len(a.Lectures) != len(b.Lectures) {
    return false
  } else {
    for i, _ := range a.Lectures {
      al := a.Lectures[i]
      bl := b.Lectures[i]
      if al.Name != bl.Name ||
        al.Grade != bl.Grade ||
        al.Department != bl.Department ||
        al.Location != bl.Location ||
        al.Wday != bl.Wday ||
        al.StartTime != bl.StartTime ||
        al.EndTime != bl.EndTime ||
        al.URI != bl.URI {
        return false
      }
      if len(al.Lecturers) != len(bl.Lecturers) {
        return false
      }
      for j, _ := range al.Lecturers {
        if al.Lecturers[j] != bl.Lecturers[j] {
          return false
        }
      }
    }
  }
  return true
}

func TestOfWday(t *testing.T) {
  for _, test := range ofWdayTests {
    out := test.timetable.OfWday(test.in)
    if !isPassedTest(out, test.out) {
      t.Error(out)
    }
  }
}

func TestOfClass(t *testing.T) {
  for _, test := range ofClassTests {
    out := test.timetable.OfClass(test.grade, test.dep)
    if !isPassedTest(out, test.out) {
      t.Error(out)
    }
  }
}
