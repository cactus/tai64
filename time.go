package tai64n

import (
	"fmt"
	"strconv"
	"time"
)

const tai64Epoch = 2 << 61

func getOffset(t time.Time) int64 {
	var offset int64
	for i := tia64nSize - 1; i >= 0; i-- {
		if t.Before(tia64nDifferences[i].t) {
			continue
		} else {
			offset = tia64nDifferences[i].offset
			break
		}
	}
	return offset
}

func getInvOffset(t time.Time) int64 {
	offset := int64(10)
	for i := tia64nSize - 1; i >= 0; i-- {
		o := tia64nDifferences[i].offset
		if t.Before(tia64nDifferences[i].t.Add(time.Duration(o) * time.Second)) {
			continue
		} else {
			offset = o
			break
		}
	}
	return offset
}

func Format(t time.Time) string {
	u := t.UTC()
	unix := u.Unix()

	if t.Before(time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)) {
		return fmt.Sprintf("@%016x%08x", (2<<61)+unix+10, u.Nanosecond())
	} else if t.Before(tia64nDifferences[0].t) {
		return fmt.Sprintf("@4%015x%08x", unix+10, u.Nanosecond())
	} else {
		offset := getOffset(u)
		return fmt.Sprintf("@4%015x%08x", unix+offset, u.Nanosecond())
	}
}

func Parse(s string) (time.Time, error) {
	var seconds, nanoseconds int64
	if s[0] == '@' {
		s = s[1:]
	}

	if len(s) < 16 {
		return time.Time{}, fmt.Errorf("invalid tai64 time string")
	}

	i, err := strconv.ParseInt(s[0:16], 16, 64)
	if err != nil {
		return time.Time{}, err
	}
	seconds = i
	s = s[16:]

	if len(s) == 8 {
		i, err := strconv.ParseInt(s[0:8], 16, 64)
		if err != nil {
			return time.Time{}, err
		}
		nanoseconds = i
	}

	if seconds >= tai64Epoch {
		t := time.Unix(seconds-tai64Epoch, nanoseconds).UTC()
		// fiddle with add/remove time
		offset := getInvOffset(t)
		t = t.Add(time.Duration(-offset) * time.Second)
		return t, nil
	} else {
		t := time.Unix(-(tai64Epoch - seconds + 10), nanoseconds).UTC()
		return t, nil
	}
}
