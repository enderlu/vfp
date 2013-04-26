package vfp

import "time"

type ShortDate struct {
	Year  int
	Month int
	Day   int
}

//Returns the current system date, which is controlled by the
// operating system, or creates a year 2000-compliant Date value
func Date(zargs ...interface{}) string {
	t := time.Now()
	if len(zargs) > 0 {
		t = zargs[0].(time.Time)
	}

	z1, z2, z3 := t.Date()
	return Transform(int(z1)) + "-" +
		Transform(int(z2)) + "-" + Transform(int(z3))
}

func Datetime(zargs ...interface{}) string {
	t := time.Now()
	if len(zargs) > 0 {
		t = zargs[0].(time.Time)
	}

	return Transform(Year(t)) + "-" +
		Transform(Month(t)) + "-" + 
		Transform(Day(t)) + " " + 
		Transform(Hour(t)) + ":"  + 
		Transform(Minute(t)) + ":"  + 
		Transform(Sec(t)) 
	//return Strextract(Transform(t), ".", ".", 0)

}

func getshortdate() (zsd ShortDate) {
	t := time.Now()
	z1, z2, z3 := t.Date()
	zsd.Year, zsd.Month, zsd.Day = int(z1), int(z2), int(z3)
	return zsd

}

func (d ShortDate) String() string {

	return Transform(d.Year) + "-" + Transform(d.Month) + "-" + Transform(d.Day)
}

//Returns the number of seconds that have elapsed since midnight.
func Seconds() float64 {
	zt := time.Now()
	return (float64(Hour(zt))*float64(time.Hour) +
		float64(Minute(zt))*float64(time.Minute) +
		float64(Sec(zt))*float64(time.Second) +
		float64(zt.Nanosecond())) / float64(time.Second)
}

//Returns the year from the specified date or datetime expression.
func Year(zt time.Time) int {
	return int(zt.Year())
}

//Returns the number of the month for a given Date or DateTime expression.
func Month(zt time.Time) int {
	return int(zt.Month())
}

//Returns the numeric day of the month for a given Date or DateTime expression.
func Day(zt time.Time) int {
	return int(zt.Day())
}

//Returns the hour portion from a DateTime expression.
func Hour(zt time.Time) int {
	return int(zt.Hour())
}

//Returns the minute portion from a DateTime expression.
func Minute(zt time.Time) int {
	return int(zt.Minute())
}

//Returns the seconds portion from a DateTime expression.
func Sec(zt time.Time) int {
	return int(zt.Second())
}

func Addyear(zt time.Time, znum int) time.Time {
	return zt.AddDate(znum, 0, 0)
}

func Addmonth(zt time.Time, znum int) time.Time {
	return zt.AddDate(0, znum, 0)
}

func Addday(zt time.Time, znum int) time.Time {
	return zt.AddDate(0, 0, znum)
}

func Addhour(t time.Time, znum int) time.Time {
	year, month, day := t.Date()
	hour, min, sec := t.Clock()
	return time.Date(year, month, day, hour+znum, min, sec, int(t.Nanosecond()), t.Location())
}

func Addminute(t time.Time, znum int) time.Time {
	year, month, day := t.Date()
	hour, min, sec := t.Clock()
	return time.Date(year, month, day, hour, min+znum, sec, int(t.Nanosecond()), t.Location())
}

func Addsecond(t time.Time, znum int) time.Time {
	year, month, day := t.Date()
	hour, min, sec := t.Clock()
	return time.Date(year, month, day, hour, min, sec+znum, int(t.Nanosecond()), t.Location())
}
