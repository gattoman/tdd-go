package money

type Dollar struct {
	amount int
}

func NewDollar(amount int) *Dollar {
	return &Dollar{amount: amount}
}

func (self Dollar) equals(dollar *Dollar) bool {
	return self.amount == dollar.amount
}

func (self Dollar) times(multiplier int) *Dollar {
	return NewDollar(self.amount * multiplier)
}
