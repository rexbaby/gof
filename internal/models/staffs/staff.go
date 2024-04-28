package staffs

import "time"

type Staff struct {
	ID        int64
	Name      string
	Status    int
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt time.Time
	UpdatedBy string
}

func (Staff) TableName() string {
	return "staffs"
}
