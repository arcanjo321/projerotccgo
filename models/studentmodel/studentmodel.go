package studentmodel

import (
	"fmt"
	"projetotccgo/config"
	"projetotccgo/entities"
)

func GetAll() []entities.Student {
	rows, err := config.DB.Query(`SELECT * FROM alunos`)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var students []entities.Student

	for rows.Next() {
		var student entities.Student
		if err := rows.Scan(&student.Id, &student.Matricula, &student.Nome_aluno); err != nil {
			panic(err)
		}

		students = append(students, student)
	}

	return students
}

func Create(student entities.Student) bool {
	result, err := config.DB.Exec(`
		INSERT INTO alunos (matricula, nome_aluno) 
		VALUES (?, ?)`,
		student.Matricula, student.Nome_aluno,
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

func Detail(id int) entities.Student {
	row := config.DB.QueryRow(`SELECT matricula, nome_aluno FROM alunos WHERE id = ?`, id)

	var student entities.Student
	if err := row.Scan(&student.Matricula, &student.Nome_aluno); err != nil {
		panic(err.Error())
	}

	return student
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

func Update(id int, student entities.Student) (bool, error) {
	// Usando placeholders para prevenir SQL injection
	query := `UPDATE alunos SET matricula = ?, nome_aluno = ? WHERE id = ?`
	_, err := config.DB.Exec(query, student.Matricula, student.Nome_aluno, student.Id)
	if err != nil {
		// Retornando o erro para o chamador da função
		return false, fmt.Errorf("falha ao executar a query UPDATE: %v", err)
	}

	return true, nil
}

func Delete(id int) error {
	_, err := config.DB.Exec("DELETE FROM alunos WHERE id = ?", id)
	return err
}
