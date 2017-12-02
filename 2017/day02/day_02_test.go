package main

import "testing"

func Test_minMaxDiff(t *testing.T) {
	cases := []struct {
		r    Row
		want float64
	}{
		// Official cases
		{Row{5, 1, 9, 5}, 8},
		{Row{7, 5, 3}, 4},
		{Row{2, 4, 6, 8}, 6},
	}

	for _, c := range cases {
		got := minMaxDiff(c.r)

		if got != c.want {
			t.Errorf("minMaxDiff(%+v) wants %f; got %f", c.r, c.want, got)
		}
	}
}

func Test_findEvenDivend(t *testing.T) {
	cases := []struct {
		r    Row
		want float64
	}{
		// Official cases
		{Row{5, 9, 2, 8}, 4},
		{Row{9, 4, 7, 3}, 3},
		{Row{3, 8, 6, 5}, 2},

		// Unofficial cases
		{Row{}, 0},
	}

	for _, c := range cases {
		got := findEvenDividend(c.r)

		if got != c.want {
			t.Errorf("findEvenDivend(%+v) wants %f; got %f", c.r, c.want, got)
		}
	}
}

func Test_spreadsheetChecksum(t *testing.T) {
	cases := []struct {
		spreadsheet string
		want        Checksum
		wantErr     bool
	}{
		// Official result
		{puzzleInput, Checksum{48357, 351}, false},

		// Error cases
		{"happy holidays", Checksum{0, 0}, true},
	}

	for _, c := range cases {
		got, err := spreadsheetChecksum(c.spreadsheet)

		if c.wantErr == (err == nil) {
			t.Errorf("wantErr %t; got err: %s", c.wantErr, err)
		}

		if got != c.want {
			t.Errorf("spreadsheetChecksum wanted %+v; got %+v", c.want, got)
		}
	}
}
