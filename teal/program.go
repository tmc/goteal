package teal

import "strings"

type Program struct {
	Lines []string
}

func (p Program) String() string {
	return strings.Join(p.Lines, "\n")
}

func (p *Program) AppendLine(l string) {
	p.Lines = append(p.Lines, l)
}
