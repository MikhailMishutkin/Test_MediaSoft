// в процессе - непонятно, будет ли на выходе из конструктора интерфейс или ссылка на структуру

package model

type man struct {
	Name        string
	Surname     string
	YearOfBirth string
	GroupName   string
}

type group struct {
	Name     string
	Members  []string
	Subgroup *group
}

type Group interface{}

// ??????????? надо ли использовать???
type Man interface {
	/* Create(string)
	Update(string)
	Delete() */
}

// man - структура или интерфейс???
func NewMan(m *man) *man {
	return &man{
		Name:        m.Name,
		Surname:     m.Surname,
		YearOfBirth: m.YearOfBirth,
		GroupName:   m.GroupName,
	}
}

func NewGroup(g *group) Group {
	return &group{
		Name:     g.Name,
		Members:  g.Members,
		Subgroup: g.Subgroup,
	}
}
