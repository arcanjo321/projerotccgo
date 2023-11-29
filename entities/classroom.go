package entities

type Classroom struct {
	Id                  uint
	Codigo_sala         string
	Nome_sala           string
	Local_sala          string
	Hora_inicio_reserva string
	Hora_fim_reserva    string
	Reservado           bool
}
