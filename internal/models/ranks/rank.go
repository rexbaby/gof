package ranks

type Rank struct {
	ID      int64
	StaffID int64
	DeptID  int64
	Level   int
}

func (Rank) TableName() string {
	return "ranks"
}
