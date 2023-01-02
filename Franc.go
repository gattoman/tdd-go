package money

type Flanc struct {
	amount int
}

func NewFlanc(amount int) *Flanc {
	return &Flanc{amount: amount}
}

func (self Flanc) equals(Flanc *Flanc) bool {
	return self.amount == Flanc.amount
}

func (self Flanc) times(multiplier int) *Flanc {
	return NewFlanc(self.amount * multiplier)
}
