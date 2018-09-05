// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// TestDateParsing copied from net/mail message_test.go
// Go standard library package and adjusted to compare results with parseDate in nntp.
package nntp

import (
	"net/mail"
	"testing"
	"time"
)

func TestDateParsing(t *testing.T) {
	tests := []struct {
		dateStr string
		exp     time.Time
	}{
		// RFC 5322, Appendix A.1.1
		{
			"Fri, 21 Nov 1997 09:55:06 -0600",
			time.Date(1997, 11, 21, 9, 55, 6, 0, time.FixedZone("", -6*60*60)),
		},
		// RFC 5322, Appendix A.6.2
		// Obsolete date.
		{
			"21 Nov 97 09:55:06 GMT",
			time.Date(1997, 11, 21, 9, 55, 6, 0, time.FixedZone("GMT", 0)),
		},
		// Commonly found format not specified by RFC 5322.
		{
			"Fri, 21 Nov 1997 09:55:06 -0600 (MDT)",
			time.Date(1997, 11, 21, 9, 55, 6, 0, time.FixedZone("", -6*60*60)),
		},
		// Not a date
		{
			"Lorem ipsum dolor sit amet",
			time.Time{},
		},
	}
	for _, test := range tests {
		hdr := mail.Header{
			"Date": []string{test.dateStr},
		}
		date, err := hdr.Date()
		if (err != nil) && (test.exp != time.Time{}) {
			t.Errorf("Header(Date: %s).Date(): %v", test.dateStr, err)
		} else if !date.Equal(test.exp) {
			t.Errorf("Header(Date: %s).Date() = %+v, want %+v", test.dateStr, date, test.exp)
		}

		date, err = mail.ParseDate(test.dateStr)
		if (err != nil) && (test.exp != time.Time{}) {
			t.Errorf("ParseDate(%s): %v", test.dateStr, err)
		} else if !date.Equal(test.exp) {
			t.Errorf("ParseDate(%s) = %+v, want %+v", test.dateStr, date, test.exp)
		}

		date, err = parseDate(test.dateStr)
		if (err != nil) && (test.exp != time.Time{}) {
			t.Errorf("parseDate(%s): %v", test.dateStr, err)
		} else if !date.Equal(test.exp) {
			t.Errorf("parseDate(%s) = %+v, want %+v", test.dateStr, date, test.exp)
		}

	}
}
