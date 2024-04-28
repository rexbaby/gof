package departments

import (
	"time"
)

type Department struct {
	ID        int64
	Name      string
	Status    int
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt time.Time
	UpdatedBy string
}

func (Department) TableName() string {
	return "departments"
}
