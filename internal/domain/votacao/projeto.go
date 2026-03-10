package votacao

type Projeto struct {
	ID string

	Pareceres *[]Parecer

	CreatedAt string
	UpdatedAt string
}
