package model

type Roles struct {
	ID          int    `json:"iId_RolIsoModulo"`
	Codigo      string `json:"vCodigo"`
	Description string `json:"vDescripcion"`
	EstadoID    int    `json:"iId_Estado"`
}
