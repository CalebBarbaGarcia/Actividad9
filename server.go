  
package main

import (
	
	"fmt"
	"strconv"
	"io/ioutil"
	"net/http"
)



type Server struct{
	Materias map[string]map[string]float64

	Alumnos map[string]map[string]float64

	SeInicializo bool
}

var Server1 Server

func (this* Server) Constructor(s string){
	if this.SeInicializo == false{
		this.Materias = make(map[string]map[string]float64)
		this.Alumnos = make(map[string]map[string]float64)
		
		this.SeInicializo = true
	}
}

func (this *Server) AgregarCalificacion(s []string){
	v, err := this.Materias[s[0]]

	if err == false{
		alumno := make(map[string]float64)
		
		f2, _ := strconv.ParseFloat(s[2], 8)
		alumno[s[1]] = f2

		// creacion de una materia
		this.Materias[s[0]] = alumno
	} else {
		_, err2 := v[s[1]]
		if err2 == false{
			alumno := make(map[string]float64)

			f2, _ := strconv.ParseFloat(s[2], 8)

			for auxAlumno, calificacion := range this.Materias[s[0]] {
				alumno[auxAlumno] = calificacion
			}


			alumno[s[1]] = f2
			
			this.Materias[s[0]] = alumno
			
		} else {
			
		}
	}

	v2, err2 := this.Alumnos[s[1]]
	
	if err2 == false{
		clase := make(map[string]float64)
		
		f2, _ := strconv.ParseFloat(s[2], 8)
		clase[s[0]] = f2

		// creacion de una materia
		this.Alumnos[s[1]] = clase
	} else {
		_, err4 := v2[s[0]]
		if err4 == false{
			clase := make(map[string]float64)

			f2, _ := strconv.ParseFloat(s[2], 8)

			for auxClase, calificacion := range this.Alumnos[s[1]] {
				clase[auxClase] = calificacion
			}


			clase[s[0]] = f2
			
			this.Alumnos[s[1]] = clase
			
		} else {
			
		}
	}
}

func (this *Server) PromedioAlumno(nombre string) float64{
	var promedio float64
	var i int64
	promedio = 0
	i = 0

	for _, calificacion := range this.Alumnos[nombre] {
		promedio = promedio + calificacion
		i = i + 1
	}
	promedio = promedio / float64(i)
	return promedio
}

func (this *Server) PromedioMateria(nombre string) float64{
	var promedio float64
	var i int64
	promedio = 0
	i = 0

	for _, calificacion := range this.Materias[nombre] {
		promedio = promedio + calificacion
		i = i + 1
	}
	promedio = promedio / float64(i)
	return promedio
}

func (this *Server) PromedioGeneral() float64{
	var promedio float64
	var promedioGeneral float64
	var i int64
	
	i = 0
	
	promedio = 0
	promedioGeneral = 0

	for nombreAlumno := range this.Alumnos {
		
		promedio = 0
		for _, calificacion := range this.Alumnos[nombreAlumno] {
			promedio = promedio + float64(calificacion)
			i = i + 1
		}
		
		promedioGeneral = promedioGeneral + promedio
		
	}
	promedioGeneral = promedioGeneral / float64(i)
	return promedioGeneral
	
}

func cargarHtml(a string) string {
	html, _ := ioutil.ReadFile(a)

	return string(html)
}


func form(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	fmt.Fprintf(
		res,
		cargarHtml("form.html"),
	)
}

func tareas(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Method)
	switch req.Method {
	case "POST":
		if err := req.ParseForm(); err != nil {
			fmt.Fprintf(res, "ParseForm() error %v", err)
			return
		}
		fmt.Println(req.PostForm)
		
		var s []string

		var f float64

		s = append(s,req.FormValue("materia"))
		s = append(s,req.FormValue("alumno"))
		s = append(s,req.FormValue("calificacion"))

		Server1.AgregarCalificacion(s)

		f = Server1.PromedioGeneral()

		println(f)
	}
}

func form2(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	fmt.Fprintf(
		res,
		cargarHtml("form2.html"),
	)
}

func tareas2(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Method)
	switch req.Method {
	case "POST":
		if err := req.ParseForm(); err != nil {
			fmt.Fprintf(res, "ParseForm() error %v", err)
			return
		}
		fmt.Println(req.PostForm)
		
		
		res.Header().Set(
			"Content-Type",
			"text/html",
		)
		fmt.Fprintf(
			res,
			cargarHtml("respuesta.html"),
			Server1.PromedioAlumno(req.FormValue("alumno")),
		)
	}
}

func form3(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	fmt.Fprintf(
		res,
		cargarHtml("form3.html"),
	)
}

func tareas3(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Method)
	switch req.Method {
	case "POST":
		if err := req.ParseForm(); err != nil {
			fmt.Fprintf(res, "ParseForm() error %v", err)
			return
		}
		fmt.Println(req.PostForm)
		
		
		res.Header().Set(
			"Content-Type",
			"text/html",
		)
		fmt.Fprintf(
			res,
			cargarHtml("respuesta2.html"),
			Server1.PromedioMateria(req.FormValue("materia")),
		)
	}
}

func form4(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	fmt.Fprintf(
		res,
		cargarHtml("form4.html"),
	)
}

func tareas4(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Method)
	switch req.Method {
	case "GET":
		res.Header().Set(
			"Content-Type",
			"text/html",
		)
		fmt.Fprintf(
			res,
			cargarHtml("respuesta3.html"),
			Server1.PromedioGeneral(),
		)

	}
}



func main() {
	Server1.Constructor("hola")
	http.HandleFunc("/form", form)
	http.HandleFunc("/tareas", tareas)
	fmt.Println("Corriendo servirdor de tareas...")

	http.HandleFunc("/form2", form2)
	http.HandleFunc("/tareas2", tareas2)
	fmt.Println("Corriendo servirdor de tareas...")

	http.HandleFunc("/form3", form3)
	http.HandleFunc("/tareas3", tareas3)
	fmt.Println("Corriendo servirdor de tareas...")

	http.HandleFunc("/form4", form4)
	http.HandleFunc("/tareas4", tareas4)
	fmt.Println("Corriendo servirdor de tareas...")

	http.ListenAndServe(":9000", nil)
}