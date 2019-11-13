package apple

type AppleInterface interface {
	GetApple() string
	GetFruit() string
}
type Apple struct{
	Name string
}

func (a *Apple) GetApple() string {
	return "Apple"
}

func (a *Apple) GetFruit() string {
	return a.GetApple()
}
