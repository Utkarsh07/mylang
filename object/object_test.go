package object

import "testing"

func TestStringHashKey(t *testing.T) {
	str1 := &String{Value: "Hello World"}
	str2 := &String{Value: "Hello World"}

	diff1 := &String{Value: "My name is Utkarsh"}
	diff2 := &String{Value: "My name is Utkarsh"}

	if str1.HashKey() != str2.HashKey() {
		t.Errorf("Strings with same content have different hash keys")
	}

	if diff1.HashKey() != diff2.HashKey() {
		t.Errorf("Strings with same content have different hash keys")
	}

	if str1.HashKey() == diff1.HashKey() {
		t.Errorf("Strings with different content have same hash keys")
	}
}
