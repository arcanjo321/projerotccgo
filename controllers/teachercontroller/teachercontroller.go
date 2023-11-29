package teachercontroller

import (
	"fmt"
	"html/template"
	"net/http"
	"projetotccgo/entities"
	"projetotccgo/models/teachermodel"
	"strconv"
)

func IndexTeacher(w http.ResponseWriter, r *http.Request) {
	teachers := teachermodel.GetAll()
	data := map[string]any{
		"teachers": teachers,
	}

	temp, err := template.ParseFiles("views/teacher/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func AddTeacher(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/teacher/create.html")
		if err != nil {
			panic(err)
		}

		temp.Execute(w, nil)
	}

	if r.Method == "POST" {
		var teacher entities.Teacher

		teacher.Nome_professor = r.FormValue("Nome_professor")
		teacher.Cpf = r.FormValue("Cpf")
		teacher.Externo = r.FormValue("Externo") == "true"
		teacher.Instituicao = r.FormValue("Instituicao")

		if ok := teachermodel.Create(teacher); !ok {
			temp, _ := template.ParseFiles("views/teacher/create.html")
			temp.Execute(w, nil)
		}

		http.Redirect(w, r, "teacher", http.StatusSeeOther)
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

func EditTeacher(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		temp, err := template.ParseFiles("views/teacher/edit.html")
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

		teacher := teachermodel.Detail(id)
		if teacher.Nome_professor == "" {
			http.Error(w, "Professor não encontrado", http.StatusNotFound)
			return
		}

		data := map[string]interface{}{
			"teacher": teacher,
		}

		if err := temp.Execute(w, data); err != nil {
			http.Error(w, fmt.Sprintf("Erro ao executar o modelo: %v", err), http.StatusInternalServerError)
		}
	}

	if r.Method == http.MethodPost {
		var teacher entities.Teacher

		idString := r.FormValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(w, fmt.Sprintf("Erro ao converter ID para inteiro: %v", err), http.StatusBadRequest)
			return
		}

		// Aqui, use o ID diretamente, pois você já o converteu para int
		teacher.Id = uint(id)

		teacher.Cpf = r.FormValue("Cpf")
		teacher.Nome_professor = r.FormValue("Nome_professor")
		teacher.Externo, _ = strconv.ParseBool(r.FormValue("Externo"))
		teacher.Instituicao = r.FormValue("Instituicao")

		if ok, err := teachermodel.Update(id, teacher); err != nil {
			http.Error(w, fmt.Sprintf("Erro ao atualizar o professor: %v", err), http.StatusInternalServerError)
			return
		} else if !ok {
			http.Error(w, "Erro ao atualizar o professor", http.StatusInternalServerError)
			return
		}

		// Redirecione apenas após uma atualização bem-sucedida
		http.Redirect(w, r, "/teacher", http.StatusSeeOther)
	}
}

func DeleteTeacher(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	if err := teachermodel.Delete(id); err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/teacher", http.StatusSeeOther)
}
