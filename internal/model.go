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

type Man interface {
	/* Create(string)
	Update(string)
	Delete() */
}

func NewMan(m *man) Man {
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
