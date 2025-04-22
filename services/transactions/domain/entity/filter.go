package entity

type GeneralFilter struct {
	Id        string
	IsDeleted bool
	Q         string
	Option    *GeneralFilterOption
	DateRange *DateRange
}

type GeneralFilterOption struct {
	OrderBy string
	Limit   int
	Offset  int
	Sort    string
}

func (f *GeneralFilter) DisableOption() {
	f.Option = nil
}
