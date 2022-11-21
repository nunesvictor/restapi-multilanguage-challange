package schemas

type AutorInput struct {
	Nome      string
	Sobrenome string
}

type EditoraInput struct {
	Nome       string
	Localidade string
}

type GeneroInput struct {
	Descricao string
}

type SingleResourceUri struct {
	ID uint `uri:"id" binding:"required"`
}
