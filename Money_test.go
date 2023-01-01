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
