package entities

type Column struct {
	Name     string
	Position int
}

func (c Column) String() string {
	return c.Name
}
func (c Column) Equals(other Column) bool {
	return c.Name == other.Name && c.Position == other.Position
}
