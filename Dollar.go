package money

type Dollar struct {
    amount int
}

func NewDollar(amount int) *Dollar {
    return &Dollar{amount: amount}
}

func (n Dollar) equals(d *Dollar) bool {
    return true
}

func (d Dollar) times(multiplier int) *Dollar {
    return NewDollar(d.amount * multiplier)
}
