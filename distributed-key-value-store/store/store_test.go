package store

import (
	"errors"
	"testing"
)

func TestStore(t *testing.T) {
	store := NewStore()
	store.AddRecord("12", "qwerty")

	val, err := store.GetRecord("key")
	if !errors.Is(err, ErrorNoSuchKey) {
		t.Errorf("failed; expected %v, get %v", ErrorNoSuchKey, err)
	}
	if val != "" {
		t.Errorf("failed; expected %s, get %s", "", val)
	}


	val, err = store.GetRecord("12")
	if err != nil {
		if !errors.Is(err, ErrorNoSuchKey) {
			t.Errorf("failed; expected %v, get %v", ErrorNoSuchKey, err)
		}		
	}
	if val != "qwerty" {
		t.Errorf("failed; expected %s, get %s", "qwerty", val)
	}
	
	err = store.DeleteRecord("12")
	if err != nil {
		if !errors.Is(err, ErrorNoSuchKey) {
			t.Errorf("failed; expected %v, get %v", ErrorNoSuchKey, err)
		}
	}

	err = store.DeleteRecord("qwe")
	if err != nil {
		if !errors.Is(err, ErrorNoSuchKey) {
			t.Errorf("failed; expected %v, get %v", ErrorNoSuchKey, err)
		}
	}

}
