// Copyright (c) 2012-2016 Eli Janssen
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.
package tai64n

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRoundTrip(t *testing.T) {
	var tests = map[string]time.Time{
		"a while ago":                time.Date(1920, 01, 01, 00, 00, 00, 00, time.UTC),
		"second before tai swap":     time.Date(1969, 12, 31, 23, 59, 49, 00, time.UTC),
		"second after tai swap":      time.Date(1969, 12, 31, 23, 59, 51, 00, time.UTC),
		"second before unix epoc":    time.Date(1969, 12, 31, 23, 59, 59, 00, time.UTC),
		"second after unix epoc":     time.Date(1970, 01, 01, 00, 00, 01, 00, time.UTC),
		"second before tai-utc epoc": time.Date(1970, 01, 01, 00, 00, 8, 00, time.UTC),
		"second after tai-utc epoc":  time.Date(1970, 01, 01, 00, 00, 11, 00, time.UTC),
		"a more current date":        time.Date(2016, 12, 31, 23, 59, 59, 00, time.UTC),
	}

	for name, tt := range tests {
		o := Format(tt)
		p, err := Parse(o)
		assert.Nil(t, err, "%s: test failed parsing", name)
		assert.Equal(t, tt, p, "%s: test failed date compare: %s", name)
	}
}
