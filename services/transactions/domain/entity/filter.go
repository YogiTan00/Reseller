package entity

type GeneralFilter struct {
	Id        string
	IsDeleted bool
	Q         string
	Option    *GeneralFilterOption
}

type GeneralFilterOption struct {
	OrderBy string
	Limit   int
	Offset  int
}

func (f *GeneralFilter) DisableOption() {
	f.Option = nil
}
