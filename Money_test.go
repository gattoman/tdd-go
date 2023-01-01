package money

import "testing"

func TestMultiplication(t *testing.T) {
	five := NewDollar(5)
	product := five.times(2)
	if product.amount != 10 {
		t.Errorf("times(2) is not 10")
	}
	product = five.times(3)
	if product.amount != 15 {
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
