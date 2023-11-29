package main

import (
	"log"
	"net/http"
	"projetotccgo/config"
	"projetotccgo/controllers/classroomcontroller"
	"projetotccgo/controllers/homecontroller"
	"projetotccgo/controllers/studentcontroller"
	"projetotccgo/controllers/teachercontroller"
)

func main() {
	config.ConnectDB()

	// 1. Homepage
	http.HandleFunc("/", homecontroller.Welcome)

	// 2. Teachers
	http.HandleFunc("/teacher", teachercontroller.IndexTeacher)
	http.HandleFunc("/teacher/addTeacher", teachercontroller.AddTeacher)
	http.HandleFunc("/teacher/editTeacher", teachercontroller.EditTeacher)
	http.HandleFunc("/teacher/deleteTeacher", teachercontroller.DeleteTeacher)

	// 3. Student
	http.HandleFunc("/student", studentcontroller.IndexStudent)
	http.HandleFunc("/student/addStudent", studentcontroller.AddStudent)
	http.HandleFunc("/student/editStudent", studentcontroller.EditStudent)
	http.HandleFunc("/student/deleteStudent", studentcontroller.DeleteStudent)

	// 3. Classroom
	http.HandleFunc("/classroom", classroomcontroller.IndexClassroom)
	http.HandleFunc("/classroom/addClassroom", classroomcontroller.AddClassroom)
	http.HandleFunc("/classroom/editClassroom", classroomcontroller.EditClassroom)
	http.HandleFunc("/classroom/deleteClassroom", classroomcontroller.DeleteClassroom)

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
