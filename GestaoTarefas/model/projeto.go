package model

type Projeto struct {
	Nome      string
	Id        string
	Descricao string
	Idusuario int    //ID do usuário responsável pelo projeto
	tarefas   string //tarefas associadas ao projeto
}
