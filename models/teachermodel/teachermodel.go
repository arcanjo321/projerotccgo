package teachermodel

import (
	"fmt"
	"projetotccgo/config"
	"projetotccgo/entities"
)

func GetAll() []entities.Teacher {
	rows, err := config.DB.Query(`SELECT * FROM professores`)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var teachers []entities.Teacher

	for rows.Next() {
		var teacher entities.Teacher
		if err := rows.Scan(&teacher.Id, &teacher.Cpf, &teacher.Externo, &teacher.Nome_professor, &teacher.Instituicao); err != nil {
			panic(err)
		}

		teachers = append(teachers, teacher)
	}

	return teachers
}

func Create(teacher entities.Teacher) bool {
	result, err := config.DB.Exec(`
		INSERT INTO professores (cpf, externo, nome_professor, instituicao) 
		VALUES (?, ?, ?, ?)`,
		teacher.Cpf, teacher.Externo, teacher.Nome_professor, teacher.Instituicao,
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

func Detail(id int) entities.Teacher {
	row := config.DB.QueryRow(`SELECT cpf, nome_professor, instituicao, externo FROM professores WHERE id = ?`, id)

	var teacher entities.Teacher
	if err := row.Scan(&teacher.Cpf, &teacher.Nome_professor, &teacher.Instituicao, &teacher.Externo); err != nil {
		panic(err.Error())
	}

	return teacher
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

func Update(id int, teacher entities.Teacher) (bool, error) {
	// Usando placeholders para prevenir SQL injection
	query := `UPDATE professores SET nome_professor = ?, cpf = ?, externo = ?, instituicao = ? WHERE id = ?`
	_, err := config.DB.Exec(query, teacher.Nome_professor, teacher.Cpf, teacher.Externo, teacher.Instituicao, teacher.Id)
	if err != nil {
		// Retornando o erro para o chamador da função
		return false, fmt.Errorf("falha ao executar a query UPDATE: %v", err)
	}

	return true, nil
}

func Delete(id int) error {
	_, err := config.DB.Exec("DELETE FROM professores WHERE id = ?", id)
	return err
}
