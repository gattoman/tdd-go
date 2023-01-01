package money

type Dollar struct {
    amount int
}

func NewDollar(amount int) *Dollar {
    return &Dollar{amount: amount}
}

func (d Dollar) times (multiplier int) *Dollar {
    return NewDollar( d.amount * multiplier)
}
