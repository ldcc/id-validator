package idvalidator

import (
	"testing"
)

// go test -v -cover -coverprofile=cover.out
// go tool cover -func=cover.out
// go tool cover -html=cover.out
func TestIsValid(t *testing.T) {
	ids := [4]string{
		"440308199901101512",
		"610104620927690",
		"810000199408230021",
		"830000199201300022",
	}
	for _, id := range ids {
		if !IsValid(id) {
			t.Errorf("%s must be true.", id)
		}
	}

	errIds := [6]string{
		"440308199901101513",
		"4403081999011015133",
		"510104621927691",
		"61010462092769",
		"810000199408230022",
		"830000199201300023",
	}
	for _, id := range errIds {
		if IsValid(id) {
			t.Errorf("%s must be false.", id)
		}
	}
}

func TestGetInfo(t *testing.T) {
	info, err := GetInfo("440202197406035316")
	if err != nil {
		t.Fatal("Errors must be nil.", err)
	}
	t.Log(info)
	//info, err = GetInfo("440308199901101513")
	//if err != nil {
	//    t.Fatal("Errors must not be nil.", err)
	//}
	//t.Log(info)
}

func TestUpgradeId(t *testing.T) {
	_, err := UpgradeId("610104620927690")
	if err != nil {
		t.Errorf("Errors must be nil.")
	}

	_, e := UpgradeId("61010462092769")
	if e == nil {
		t.Errorf("Errors must not be nil.")
	}
}

func TestFakeId(t *testing.T) {
	id := FakeId()
	if len(id) != 18 {
		t.Errorf("String length must be 18. : %s", id)
	}
	if !IsValid(id) {
		t.Errorf("%s must be true.", id)
	}
}

func TestFakeRequireId(t *testing.T) {
	id := FakeRequireId(false, "", "", 0)
	if len(id) != 15 {
		t.Errorf("String length must be 15. : %s", id)
	}
	if !IsValid(id) {
		t.Errorf("%s must be true.", id)
	}

	info, _ := GetInfo(id)
	if info.Sex != 0 {
		t.Errorf("%s must be 0.", "0")
	}
}
