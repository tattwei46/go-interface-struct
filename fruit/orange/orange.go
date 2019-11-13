package orange

type OrangeInterface interface {
	GetOrange() string
	GetFruit() string
}
type Orange struct {
	ID   int
	Name string
}

func (a *Orange) GetOrange() string {
	return "Orange"
}

func (a *Orange) GetFruit() string {
	return a.GetOrange()
}
