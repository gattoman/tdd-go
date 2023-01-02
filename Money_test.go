package money

import (
	"reflect"
	"testing"
)

func TestMultiplication(t *testing.T) {
	five := NewDollar(5)

	if !reflect.DeepEqual(NewDollar(10), five.times(2)) {
		t.Errorf("times(2) is not 10")
	}
	if !reflect.DeepEqual(NewDollar(15), five.times(3)) {
		t.Errorf("times(3) is not 15")
	}
}

func TestEquality(t *testing.T) {
	if !(NewDollar(5).equals(NewDollar(5))) {
		t.Errorf("two of instances are not same")
	}
	if NewDollar(5).equals(NewDollar(6)) {
		t.Errorf("two of instances are not same")
	}
}
