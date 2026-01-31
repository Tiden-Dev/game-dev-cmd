package ue

type IniSection string

func (i IniSection) String() string {
	return string(i)
}

type IniKey string

func (i IniKey) String() string {
	return string(i)
}
