package main

import (
	"testing"
)

func TestSortRecordsByDate(t *testing.T) {
	accounts := []AccountingItem{
		{Date: "2021-02-01"},
		{Date: "2020-12-24"},
		{Date: "2021-02-28"},
	}
	expected := []AccountingItem{
		{Date: "2020-12-24"},
		{Date: "2021-02-01"},
		{Date: "2021-02-28"},
	}
	result := SortRecordsByDate(accounts)
	for i, ai := range result {
		if ai != expected[i] {
			t.Errorf("Slice is not sorted as expected %+v\n", result)
		}
	}
}
