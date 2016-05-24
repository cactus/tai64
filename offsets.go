// Copyright (c) 2012-2016 Eli Janssen
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

// THIS FILE IS AUTOGENERATED. DO NOT EDIT!

package tai64n

// http://maia.usno.navy.mil/ser7/tai-utc.dat
// http://www.stjarnhimlen.se/comp/time.html
var tia64nDifferences = []struct {
	utime  int64
	offset int64
}{
	{63072000, 10}, // 1972-01-01T00:00:00Z
	{78796800, 11}, // 1972-07-01T00:00:00Z
	{94694400, 12}, // 1973-01-01T00:00:00Z
	{126230400, 13}, // 1974-01-01T00:00:00Z
	{157766400, 14}, // 1975-01-01T00:00:00Z
	{189302400, 15}, // 1976-01-01T00:00:00Z
	{220924800, 16}, // 1977-01-01T00:00:00Z
	{252460800, 17}, // 1978-01-01T00:00:00Z
	{283996800, 18}, // 1979-01-01T00:00:00Z
	{315532800, 19}, // 1980-01-01T00:00:00Z
	{362793600, 20}, // 1981-07-01T00:00:00Z
	{394329600, 21}, // 1982-07-01T00:00:00Z
	{425865600, 22}, // 1983-07-01T00:00:00Z
	{489024000, 23}, // 1985-07-01T00:00:00Z
	{567993600, 24}, // 1988-01-01T00:00:00Z
	{631152000, 25}, // 1990-01-01T00:00:00Z
	{662688000, 26}, // 1991-01-01T00:00:00Z
	{709948800, 27}, // 1992-07-01T00:00:00Z
	{741484800, 28}, // 1993-07-01T00:00:00Z
	{773020800, 29}, // 1994-07-01T00:00:00Z
	{820454400, 30}, // 1996-01-01T00:00:00Z
	{867715200, 31}, // 1997-07-01T00:00:00Z
	{915148800, 32}, // 1999-01-01T00:00:00Z
	{1136073600, 33}, // 2006-01-01T00:00:00Z
	{1230768000, 34}, // 2009-01-01T00:00:00Z
	{1341100800, 35}, // 2012-07-01T00:00:00Z
	{1435708800, 36}, // 2015-07-01T00:00:00Z
}

var tia64nSize = len(tia64nDifferences)
