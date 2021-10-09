package stringUtil

import (
	"fmt"
	"testing"
)

func TestSplit(t *testing.T) {
	s := "1,2;3|4||5,6;7|8||9,"
	// seps := []string{",", ";", "|", "||"}
	retList := Split(s, nil)
	if retList[0] != "1" || retList[1] != "2" || retList[2] != "3" || retList[3] != "4" || retList[4] != "5" || retList[5] != "6" || retList[6] != "7" || retList[7] != "8" || retList[8] != "9" {
		t.Errorf("ecptected:123456789, but now got:%v", retList)
	}
}

func TestSplitToIntSlice(t *testing.T) {
	s := "1, 2, 3, 4, 5, a"
	if _, err := SplitToIntSlice(s, ","); err == nil {
		t.Errorf("Expected got err, but got nil")
	}

	s = "1, 5, 39,"
	if intSlice, err := SplitToIntSlice(s, ","); err != nil {
		t.Errorf("Expected got nil, but got error:%s", err)
	} else {
		// fmt.Printf("intSlice:%v\n", intSlice)
		if intSlice[0] != 1 || intSlice[1] != 5 || intSlice[2] != 39 {
			t.Errorf("Expected got %s, but got %v", s, intSlice)
		}
	}
}

func TestSplitToFloat64(t *testing.T) {
	result, err := SplitToFloat64Slice("1.11,2.22", ",")
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Printf("%v\n", result)
}
