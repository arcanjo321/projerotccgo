package classroomcontroller

import (
	"fmt"
	"html/template"
	"net/http"
	"projetotccgo/entities"
	"projetotccgo/models/classroommodel"
	"strconv"
)

func IndexClassroom(w http.ResponseWriter, r *http.Request) {
	classrooms := classroommodel.GetAll()
	data := map[string]any{
		"classrooms": classrooms,
	}

	temp, err := template.ParseFiles("views/classroom/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func AddClassroom(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/classroom/create.html")
		if err != nil {
			panic(err)
		}

		temp.Execute(w, nil)
	}

	if r.Method == "POST" {
		var classroom entities.Classroom

		classroom.Codigo_sala = r.FormValue("Codigo_sala")
		classroom.Hora_inicio_reserva = r.FormValue("Hora_inicio_reserva")
		classroom.Hora_fim_reserva = r.FormValue("Hora_fim_reserva")
		classroom.Local_sala = r.FormValue("Local_sala")
		classroom.Nome_sala = r.FormValue("Nome_sala")
		classroom.Reservado = r.FormValue("Reservado") == "true"

		if ok := classroommodel.Create(classroom); !ok {
			temp, _ := template.ParseFiles("views/classroom/create.html")
			temp.Execute(w, nil)
		}

		http.Redirect(w, r, "classroom", http.StatusSeeOther)
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

func EditClassroom(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		temp, err := template.ParseFiles("views/classroom/edit.html")
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

		classroom := classroommodel.Detail(id)
		if classroom.Nome_sala == "" {
			http.Error(w, "Professor não encontrado", http.StatusNotFound)
			return
		}

		data := map[string]interface{}{
			"classroom": classroom,
		}

		if err := temp.Execute(w, data); err != nil {
			http.Error(w, fmt.Sprintf("Erro ao executar o modelo: %v", err), http.StatusInternalServerError)
		}
	}

	if r.Method == http.MethodPost {
		var classroom entities.Classroom

		idString := r.FormValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(w, fmt.Sprintf("Erro ao converter ID para inteiro: %v", err), http.StatusBadRequest)
			return
		}

		// Aqui, use o ID diretamente, pois você já o converteu para int
		classroom.Id = uint(id)

		classroom.Codigo_sala = r.FormValue("Codigo_sala")
		classroom.Hora_inicio_reserva = r.FormValue("Hora_inicio_reserva")
		classroom.Hora_fim_reserva = r.FormValue("Hora_fim_reserva")
		classroom.Local_sala = r.FormValue("Local_sala")
		classroom.Nome_sala = r.FormValue("Nome_sala")
		classroom.Reservado = r.FormValue("Reservado") == "true"

		if ok, err := classroommodel.Update(id, classroom); err != nil {
			http.Error(w, fmt.Sprintf("Erro ao atualizar o professor: %v", err), http.StatusInternalServerError)
			return
		} else if !ok {
			http.Error(w, "Erro ao atualizar o professor", http.StatusInternalServerError)
			return
		}

		// Redirecione apenas após uma atualização bem-sucedida
		http.Redirect(w, r, "/classroom", http.StatusSeeOther)
	}
}

func DeleteClassroom(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	if err := classroommodel.Delete(id); err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/classroom", http.StatusSeeOther)
}
