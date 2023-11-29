package studentcontroller

import (
	"fmt"
	"html/template"
	"net/http"
	"projetotccgo/entities"
	"projetotccgo/models/studentmodel"
	"strconv"
)

func IndexStudent(w http.ResponseWriter, r *http.Request) {
	students := studentmodel.GetAll()
	data := map[string]any{
		"students": students,
	}

	temp, err := template.ParseFiles("views/student/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func AddStudent(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/student/create.html")
		if err != nil {
			panic(err)
		}

		temp.Execute(w, nil)
	}

	if r.Method == "POST" {
		var student entities.Student

		student.Matricula = r.FormValue("Nome_aluno")
		student.Nome_aluno = r.FormValue("Matricula")

		if ok := studentmodel.Create(student); !ok {
			temp, _ := template.ParseFiles("views/student/create.html")
			temp.Execute(w, nil)
		}

		http.Redirect(w, r, "student", http.StatusSeeOther)
	}
}

// func EditTeacher(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "GET" {
// 		temp, err := template.ParseFiles("views/teacher/edit.html")
// 		if err != nil {
// 			panic(err)
// 		}

// 		idString := r.URL.Query().Get("id")
// 		id, err := strconv.Atoi(idString)
// 		if err != nil {
// 			panic(err)
// 		}

// 		teacher := teachermodel.Detail(id)
// 		data := map[string]any{
// 			"teacher": teacher,
// 		}

// 		temp.Execute(w, data)
// 	}

// 	if r.Method == "POST" {
// 		var teacher entities.Teacher

// 		idString := r.FormValue("id")
// 		id, err := strconv.Atoi(idString)
// 		if err != nil {
// 			panic(err)
// 		}

// 		teacher.Cpf = r.FormValue("Cpf")
// 		teacher.Nome_professor = r.FormValue("Nome_professor")
// 		teacher.Externo = r.FormValue("Externo") == "true"
// 		teacher.Instituicao = r.FormValue("Instituicao")

// 		if ok := teachermodel.Update(id, teacher); !ok {
// 			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
// 			return
// 		}

// 		http.Redirect(w, r, "/teacher", http.StatusSeeOther)
// 	}
// }

func EditStudent(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		temp, err := template.ParseFiles("views/student/edit.html")
		if err != nil {
			http.Error(w, fmt.Sprintf("Erro ao analisar o modelo: %v", err), http.StatusInternalServerError)
			return
		}

		idString := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(w, fmt.Sprintf("Erro ao converter ID para inteiro: %v", err), http.StatusBadRequest)
			return
		}

		student := studentmodel.Detail(id)
		if student.Nome_aluno == "" {
			http.Error(w, "Aluno não encontrado", http.StatusNotFound)
			return
		}

		data := map[string]interface{}{
			"student": student,
		}

		if err := temp.Execute(w, data); err != nil {
			http.Error(w, fmt.Sprintf("Erro ao executar o modelo: %v", err), http.StatusInternalServerError)
		}
	}

	if r.Method == http.MethodPost {
		var student entities.Student

		idString := r.FormValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(w, fmt.Sprintf("Erro ao converter ID para inteiro: %v", err), http.StatusBadRequest)
			return
		}

		// Aqui, use o ID diretamente, pois você já o converteu para int
		student.Id = uint(id)

		student.Matricula = r.FormValue("Matricula")
		student.Nome_aluno = r.FormValue("Nome_aluno")

		if ok, err := studentmodel.Update(id, student); err != nil {
			http.Error(w, fmt.Sprintf("Erro ao atualizar o aluno: %v", err), http.StatusInternalServerError)
			return
		} else if !ok {
			http.Error(w, "Erro ao atualizar o aluno", http.StatusInternalServerError)
			return
		}

		// Redirecione apenas após uma atualização bem-sucedida
		http.Redirect(w, r, "/student", http.StatusSeeOther)
	}
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	if err := studentmodel.Delete(id); err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/student", http.StatusSeeOther)
}
