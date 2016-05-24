// Copyright (c) 2012-2016 Eli Janssen
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.
package tai64n

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var tests = map[string]struct {
	t time.Time
	o string
}{
	"a while ago": {
		time.Date(1920, 01, 01, 00, 00, 00, 00, time.UTC),
		"@3fffffffa1f2cd8a00000000",
	},
	"second before tai swap": {
		time.Date(1969, 12, 31, 23, 59, 49, 00, time.UTC),
		"@3fffffffffffffff00000000",
	},
	"second at tai swap": {
		time.Date(1969, 12, 31, 23, 59, 50, 00, time.UTC),
		"@400000000000000000000000",
	},
	"second after tai swap": {
		time.Date(1969, 12, 31, 23, 59, 51, 00, time.UTC),
		"@400000000000000100000000",
	},
	"second before unix epoc": {
		time.Date(1969, 12, 31, 23, 59, 59, 00, time.UTC),
		"@400000000000000900000000",
	},
	"second at unix epoc": {
		time.Date(1970, 01, 01, 00, 00, 00, 00, time.UTC),
		"@400000000000000a00000000",
	},
	"second after unix epoc": {
		time.Date(1970, 01, 01, 00, 00, 01, 00, time.UTC),
		"@400000000000000b00000000",
	},
	"second before tai-utc epoc": {
		time.Date(1970, 01, 01, 00, 00, 9, 00, time.UTC),
		"@400000000000001300000000",
	},
	"second at tai-utc epoc": {
		time.Date(1970, 01, 01, 00, 00, 10, 00, time.UTC),
		"@400000000000001400000000",
	},
	"second after tai-utc epoc": {
		time.Date(1970, 01, 01, 00, 00, 11, 00, time.UTC),
		"@400000000000001500000000",
	},
	"a more current date": {
		time.Date(2016, 12, 31, 23, 59, 59, 00, time.UTC),
		"@40000000586846a300000000",
	},
}

func TestRoundTrip(t *testing.T) {
	for name, tt := range tests {
		o := Format(tt.t)
		p, err := Parse(o)
		assert.Nil(t, err, "%s: test failed parsing", name)
		assert.Equal(t, tt.t, p, "%s: test failed date compare: %s", name)
	}
}

func TestFormat(t *testing.T) {
	for name, tt := range tests {
		o := Format(tt.t)
		assert.Equal(t, tt.o, o, "%s: test failed date compare", name)
	}
}

func TestParse(t *testing.T) {
	for name, tt := range tests {
		p, err := Parse(tt.o)
		assert.Nil(t, err, "%s: test failed parsing", name)
		assert.Equal(t, tt.t, p, "%s: test failed date compare", name)
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
