package fruit

type FruitInterface interface {
	GetFruit() string
}

type Fruit struct{}

func (f *Fruit) GetFruitName(fruit FruitInterface) string {
	return fruit.GetFruit()
}
