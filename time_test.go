// Copyright (c) 2012-2016 Eli Janssen
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package tai64

import (
	"sort"
	"testing"
	"time"
)

var tests = map[string]struct {
	t string
	o string
}{
	"a while ago":           {"1920-01-01T00:00:00Z", "@3fffffffa1f2cd8a00000000"},
	"before tai swap":       {"1969-12-31T23:59:49Z", "@3fffffffffffffff00000000"},
	"at tai swap":           {"1969-12-31T23:59:50Z", "@400000000000000000000000"},
	"after tai swap":        {"1969-12-31T23:59:51Z", "@400000000000000100000000"},
	"before unix epoch":     {"1969-12-31T23:59:59Z", "@400000000000000900000000"},
	"at unix epoch":         {"1970-01-01T00:00:00Z", "@400000000000000a00000000"},
	"after unix epoch":      {"1970-01-01T00:00:01Z", "@400000000000000b00000000"},
	"before tai-utc epoch":  {"1970-01-01T00:00:09Z", "@400000000000001300000000"},
	"at tai-utc epoch":      {"1970-01-01T00:00:10Z", "@400000000000001400000000"},
	"after tai-utc epoch":   {"1970-01-01T00:00:11Z", "@400000000000001500000000"},
	"right before adjust 1": {"1972-06-30T23:59:59Z", "@4000000004b2580900000000"},
	"right at adjust 1":     {"1972-07-01T00:00:00Z", "@4000000004b2580b00000000"},
	"right after adjust 1":  {"1972-07-01T00:00:01Z", "@4000000004b2580c00000000"},
	"right before adjust 2": {"2016-12-31T23:59:59Z", "@40000000586846a300000000"},
	"right at adjust 2":     {"2017-01-01T00:00:00Z", "@40000000586846a500000000"},
	"right after adjust 2":  {"2017-01-01T00:00:01Z", "@40000000586846a600000000"},
	"nanoseconds":           {"2015-06-30T23:59:59.908626131Z", "@4000000055932da2362888d3"},
}

func getOrderedTestNames() []string {
	var keys []string
	for k := range tests {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func TestRoundTripNano(t *testing.T) {
	for _, name := range getOrderedTestNames() {
		tt := tests[name]
		tm, err := time.Parse(time.RFC3339Nano, tt.t)
		if err != nil {
			t.Fatalf("%s: test failed parsing time.Time", name)
		}

		o := FormatNano(tm)
		p, err := Parse(o)
		if err != nil {
			t.Fatalf("%s: test failed parsing", name)
		}

		if tm != p {
			t.Fatalf("%s: test failed date compare:\n %s != %s", name, tm, p)
		}
	}
}

func TestRoundTrip(t *testing.T) {
	for _, name := range getOrderedTestNames() {
		tt := tests[name]
		tm, err := time.Parse(time.RFC3339Nano, tt.t)
		if err != nil {
			t.Fatalf("%s: test failed parsing time.Time", name)
		}

		o := Format(tm)
		p, err := Parse(o[:17])
		if err != nil {
			t.Fatalf("%s: test failed parsing", name)
		}

		ts := tm.Truncate(time.Second)
		if ts != p {
			t.Fatalf("%s: test failed date compare:\n %s != %s", name, ts, p)
		}
	}
}

func TestFormat(t *testing.T) {
	for _, name := range getOrderedTestNames() {
		tt := tests[name]
		tm, err := time.Parse(time.RFC3339Nano, tt.t)
		if err != nil {
			t.Fatalf("%s: test failed parsing time.Time", name)
		}

		o := FormatNano(tm)
		if tt.o != o {
			t.Fatalf("%s: test failed date compare:\n %s != %s", name, tt.o, o)
		}
	}
}

func TestParse(t *testing.T) {
	for _, name := range getOrderedTestNames() {
		tt := tests[name]
		tm, err := time.Parse(time.RFC3339Nano, tt.t)
		if err != nil {
			t.Fatalf("%s: test failed parsing time.Time", name)
		}

		p, err := Parse(tt.o)
		if err != nil {
			t.Fatalf("%s: test failed parsing", name)
		}

		if tm != p {
			t.Fatalf("%s: test failed date compare:\n %s != %s", name, tm, p)
		}
	}
}

func BenchmarkFormat(b *testing.B) {
	t := time.Date(2016, 12, 31, 23, 59, 59, 00, time.UTC)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Format(t)
	}
}

func BenchmarkParse(b *testing.B) {
	s := "@40000000586846a300000000"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Parse(s)
	}
}
