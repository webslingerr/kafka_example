package helper

import (
	"database/sql"
	"time"
)

func NullString(s string) (ns sql.NullString) {
	if s != "" {
		ns.String = s
		ns.Valid = true
	}
	return ns
}

func NullTime(s string) (ns sql.NullTime) {
	var err error
	if s != "" {
		ns.Time, err = time.Parse("2006-01-02T15:04:05.999Z", s)
		ns.Valid = err == nil
	}

	return ns
}

func NullDouble(s float64) (ns sql.NullFloat64) {
	if s != 0 {
		ns.Float64 = s
		ns.Valid = true
	}
	return ns
}

func StringValue(ns sql.NullString) string {
	if ns.Valid {
		s := ns.String
		return s
	}
	return ""
}

func DoubleValue(ns sql.NullFloat64) float64 {
	if ns.Valid {
		s := ns.Float64
		return s
	}
	return 0
}

func Int64Value(ns sql.NullInt64) int64 {
	if ns.Valid {
		s := ns.Int64
		return s
	}
	return 0
}

func FormatNullTime(nt sql.NullTime, format string) string {
	if nt.Valid {
		return nt.Time.Format(format)
	}
	return ""
}
