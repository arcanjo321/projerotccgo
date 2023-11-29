package classroommodel

import (
	"fmt"
	"projetotccgo/config"
	"projetotccgo/entities"
)

func GetAll() []entities.Classroom {
	rows, err := config.DB.Query(`SELECT * FROM sala`)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var classrooms []entities.Classroom

	for rows.Next() {
		var classroom entities.Classroom
		if err := rows.Scan(&classroom.Id, &classroom.Codigo_sala, &classroom.Hora_inicio_reserva, &classroom.Hora_fim_reserva, &classroom.Local_sala, &classroom.Nome_sala, &classroom.Reservado); err != nil {
			panic(err)
		}

		classrooms = append(classrooms, classroom)
	}

	return classrooms
}

func Create(classroom entities.Classroom) bool {
	result, err := config.DB.Exec(`
		INSERT INTO sala (codigo_sala, nome_sala, local_sala, hora_inicio_reserva, hora_fim_reserva, reservado) 
		VALUES (?, ?, ?, ?, ?, ?)`,
		classroom.Codigo_sala, classroom.Nome_sala, classroom.Local_sala, classroom.Hora_inicio_reserva, classroom.Hora_fim_reserva, classroom.Reservado,
	)

	if err != nil {
		panic(err)
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	return lastInsertId > 0
}

func Detail(id int) entities.Classroom {
	row := config.DB.QueryRow(`SELECT codigo_sala, nome_sala, local_sala, hora_inicio_reserva, hora_fim_reserva, reservado FROM sala WHERE id = ?`, id)

	var classroom entities.Classroom
	if err := row.Scan(
		&classroom.Codigo_sala,
		&classroom.Nome_sala,
		&classroom.Local_sala,
		&classroom.Hora_inicio_reserva,
		&classroom.Hora_fim_reserva,
		&classroom.Reservado); err != nil {
		panic(err.Error())
	}

	return classroom
}

// func Update(id int, teacher entities.Teacher) bool {
// 	query, err := config.DB.Exec(`UPDATE professores SET nome_professor = ?, cpf = ?, externo = ?, instituicao = ? WHERE id = ?`, teacher.Nome_professor, teacher.Cpf, teacher.Externo, teacher.Instituicao, id)
// 	if err != nil {
// 		panic(err)
// 	}

// 	result, err := query.RowsAffected()
// 	if err != nil {
// 		panic(err)
// 	}

// 	return result > 0
// }

func Update(id int, classroom entities.Classroom) (bool, error) {
	// Usando placeholders para prevenir SQL injection
	query := `UPDATE sala SET codigo_sala = ?, local_sala = ?, nome_sala = ?, hora_inicio_reserva = ?, hora_fim_reserva = ?, reservado = ? WHERE id = ?`
	_, err := config.DB.Exec(query, classroom.Codigo_sala, classroom.Local_sala, classroom.Nome_sala, classroom.Hora_inicio_reserva, classroom.Hora_fim_reserva, classroom.Reservado, classroom.Id)
	if err != nil {
		// Retornando o erro para o chamador da função
		return false, fmt.Errorf("falha ao executar a query UPDATE: %v", err)
	}

	return true, nil
}

func Delete(id int) error {
	_, err := config.DB.Exec("DELETE FROM sala WHERE id = ?", id)
	return err
}
