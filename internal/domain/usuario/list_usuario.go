package usuario

type ListUsuario struct {
	Usuarios []*Usuario
	Total    int64
	Page     int
	Limit    int
}
