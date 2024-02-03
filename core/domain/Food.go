package domain

type Food struct {
	id   int
	name string
	time int
}

func NewFoodWithId(id int, name string, time int) *Food {
	return &Food{
		id:   id,
		name: name,
		time: time,
	}
}

func NewFood(name string, time int) *Food {
	return &Food{
		name: name,
		time: time,
	}
}

func (entity Food) Id() int {
	return entity.id
}
func (entity Food) Name() string {
	return entity.name
}

func (entity Food) Time() int {
	return entity.time
}
