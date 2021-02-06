package main

import "testing"

func TestGetLastFreeItemId(t *testing.T) {
	accounts := []AccountingItem{
		{ItemId: 2},
		{ItemId: 5},
	}
	if i := GetLastFreeItemId(&accounts); i != 6 {
		t.Errorf("LastID isn't %d\n", i)
	}

	accounts = []AccountingItem{} 
	if i := GetLastFreeItemId(&accounts); i != 1 {
		t.Errorf("LastID isn't %d\n", i)
	}
}
