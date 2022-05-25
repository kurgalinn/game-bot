package service

import "testing"

func TestGenerateID(t *testing.T) {
	id1 := GenerateID()
	id2 := GenerateID()
	if id1 == id2 {
		t.Errorf("First ID %v is not different from second %v", id1, id2)
	}
}
