package calc

import (
	"testing"
)

var tables = []struct {
	input         string
	expected      int
	errorMessage  string
	expectedError string
}{
	{"", 0, "Empty String retuns 0", ""},
	{"1", 1, "Single number returns same digit", ""},
	{"4,10", 14, "2 numbers with comma retuns total", ""},
	{"10,20,30", 60, "3 numbers with comman returns total", ""},
	{"10,10,10,10,10,10,10,10,10,10", 100, "10 numbers with commans returns total", ""},
	{"1\n2", 3, "2 numbers with newline returns total", ""},
	{"1\n2,3", 6, "3 numbers , newline and comma returns total", ""},
	{"//;\n1;2", 3, "Support custom delimiters", ""},
	{"-1,2", 0, "Negative Numbers throws Exception", "negatives not allowed -1"},
	{"-1,-2,-5", 0, "Multiple negative numbers show all", "negatives not allowed -1,-2,-5"},
	{"1001,10,2", 12, "Numbers bigger than 1000 should be ignored", ""},
}

func TestSum(t *testing.T) {
	for _, table := range tables {
		total, err := Sum(table.input)
		if err != nil {
			if table.expectedError != err.Error() {
				t.Errorf("%s - Expected Error not found Expected:%s Go:%s", table.errorMessage, table.expectedError, err.Error())
			}
		} else if total != table.expected {
			t.Errorf("%s - Input:%s Expected:%d got %d", table.errorMessage, table.input, table.expected, total)
		}
	}
}

func TestGetCalledCount(t *testing.T) {
	Reset()
	count := GetCalledCount()
	if count != 0 {
		t.Errorf("GetCalledCount - should be zero if not calls")
	}
	for i := 1; i < 7; i++ {
		Reset()
		for k := 0; k < i; k++ {
			Sum("1,1")
		}
		count = GetCalledCount()
		if count != i {
			t.Errorf("GetCalled Count - Expected %d call count. Got %d", i, count)
		}
	}

}
