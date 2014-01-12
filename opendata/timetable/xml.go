package timetable

import (
  "encoding/xml"
  "fmt"
  "io"
  "net/http"
  "strconv"
  "time"
)

// TODO:Unmarshal関数があるにも関わらず、Marshal関数は存在しない。
// 需要はそれほど高くないが、やはり作成しておくべき。

// サーバーからXMLファイルをダウンロードする。
// FIXME:このコードでは、2013年度後期の時間割しかダウンロードしない。このため、最新ではないデータを返す可能性がある。
// 常に最新のデータを返すか、あるいはいつのデータをダウンロードするかを選択できるようにするべき。
func DownloadXML() (io.Reader, error) {
  res, err := http.Get("http://www.akashi.ac.jp/data/timetable/timetable201310.xml")
  return res.Body, err
}

// XMLから[]Timetableを作成する。
// ただし、バリテーションは行わないため、正しいデータであることは保証されない。
// また、不正な値は無視され、それによってエラーが返されることはない。
// (errorが常にnilであるということを保証するものではない。例えば、XMLとして正しくない場合、当然errorはそれを表す値になる。
func UnmarshalXML(reader io.Reader) (tables []Timetable, err error) {
  dec := xml.NewDecoder(reader)
  for {
    t, err := dec.Token()
    if err == io.EOF {
      break
    } else if err != nil {
      fmt.Println(err)
      break
    }
    switch t.(type) {
    case xml.StartElement:
      if t.(xml.StartElement).Name.Local == "Timetable" {
        table, err := parseTimetable(dec)
        tables = append(tables, table)
        if err != nil {
          return tables, err
        }
      }
    }
  }
  return
}

func parseTimetable(dec *xml.Decoder) (table Timetable, err error) {
  for {
    t, err := dec.Token()
    if err != nil {
      return table, err
    }

    switch t.(type) {
    case xml.StartElement:
      switch t.(xml.StartElement).Name.Local {
      case "Common":
        table.Common, err = parseCommon(dec)
      case "Lectures":
        table.Lectures, err = parseLectures(dec)
      }
    case xml.EndElement:
      if t.(xml.EndElement).Name.Local == "Timetable" {
        return table, err
      }
    }
  }
}

func parseCommon(dec *xml.Decoder) (c Common, err error) {
  fmt.Println("In common")
  var element string
  for {
    t, err := dec.Token()
    if err != nil {
      return c, err
    }

    switch t.(type) {
    case xml.StartElement:
      element = t.(xml.StartElement).Name.Local
    case xml.EndElement:
      if t.(xml.EndElement).Name.Local == "Common" {
        return c, nil
      }
      element = ""
    case xml.CharData:
      text := string(t.(xml.CharData))
      switch element {
      case "Institution":
        c.Institution = text
      case "AcademicYear":
        c.Year, _ = strconv.Atoi(text)
      case "AcademicTerm":
        c.Term = text
      }
    }
  }
}

func parseLectures(dec *xml.Decoder) (lecs []Lecture, err error) {
  fmt.Println("In Lectures")
  var element string
  for {
    t, err := dec.Token()
    if err != nil {
      return lecs, err
    }

    switch t.(type) {
    case xml.StartElement:
      element = t.(xml.StartElement).Name.Local
      if element == "Lecture" {
        tmpL, err := parseLecture(dec)
        lecs = append(lecs, tmpL)
        if err != nil {
          return lecs, err
        }
        element = ""
      }
    case xml.EndElement:
      if t.(xml.EndElement).Name.Local == "Lectures" {
        fmt.Println("Out Lectures")
        return lecs, nil
      }

    }
  }
}

func parseLecture(dec *xml.Decoder) (lec Lecture, err error) {
  fmt.Println("In Lecture")
  var element string
  for {
    t, err := dec.Token()
    if err != nil {
      return lec, err
    }

    switch t.(type) {
    case xml.StartElement:
      element = t.(xml.StartElement).Name.Local
      if element == "Lecturers" {
        lec.Lecturers, err = parseLecturers(dec)
        if err != nil {
          return lec, err
        }
      }
    case xml.EndElement:
      if t.(xml.EndElement).Name.Local == "Lecture" {
        fmt.Println("Out Lecture")
        return lec, nil
      }
      element = ""
    case xml.CharData:
      text := string(t.(xml.CharData))
      switch element {
      case "Name":
        lec.Name = text
      case "Grade":
        lec.Grade, _ = strconv.Atoi(text)
      case "Department":
        lec.Department = Department(text)
      case "Wday":
        wdayI, _ := strconv.Atoi(text)
        lec.Wday = time.Weekday(wdayI)
      case "StartTime":
        rfc3339_text := "0000-01-01T" + text
        errr := lec.StartTime.UnmarshalText([]byte(rfc3339_text))

        fmt.Println(errr, "\n\n\n\n")
      case "EndTime":
        rfc3339_text := "0000-01-01T" + text
        lec.EndTime.UnmarshalText([]byte(rfc3339_text))
      case "Lacation":
        lec.Location = text
      }
    }
  }
}

func parseLecturers(dec *xml.Decoder) (lecs []string, err error) {
  var element string
  for {
    t, err := dec.Token()
    if err != nil {
      return lecs, err
    }

    switch t.(type) {
    case xml.StartElement:
      element = t.(xml.StartElement).Name.Local
    case xml.EndElement:
      if t.(xml.EndElement).Name.Local == "Lecturers" {
        return lecs, nil
      }
      element = ""
    case xml.CharData:
      text := string(t.(xml.CharData))
      if element == "Lecturer" {
        lecs = append(lecs, text)
      }
    }
  }
}
