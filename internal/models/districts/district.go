package districts

type District struct {
	ID     int64
	Name   string
	Status int
}

func (District) TableName() string {
	return "districts"
}
