package scripts

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func EjecuteScripts(db *sql.DB) {

	createTable_usuarios := `
		CREATE TABLE IF NOT EXISTS usuarios (
			id INTEGER PRIMARY KEY ,
			nombre TEXT NOT NULL,
			apellidos TEXT NOT NULL,
			email TEXT NOT NULL UNIQUE,
			usuario TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL
		);
	`
	createTable_clases := `
		CREATE TABLE IF NOT EXISTS clases (
			id INTEGER PRIMARY KEY ,
			name TEXT NOT NULL,
			titulo TEXT NOT NULL,
			descripcion TEXT NOT NULL,
			teacher_id INTEGER NOT NULL,
			color TEXT NOT NULL,
			FOREIGN KEY (teacher_id) REFERENCES usuarios(id)
		);
	`
	createTable_recursos := `
		CREATE TABLE IF NOT EXISTS recursos (
			id INTEGER PRIMARY KEY ,
			name TEXT NOT NULL,
			url TEXT NOT NULL
		);
	`
	createTable_clases_recursos := `
		CREATE TABLE IF NOT EXISTS clases_recursos (
			clase_id INTEGER NOT NULL,
			recurso_id INTEGER NOT NULL,
			FOREIGN KEY (clase_id) REFERENCES clases(id),
			FOREIGN KEY (recurso_id) REFERENCES recursos(id),
			PRIMARY KEY (clase_id, recurso_id)
		);
	`
	createTable_ejercicios := `
		CREATE TABLE IF NOT EXISTS ejercicios (
			id INTEGER PRIMARY KEY ,
			name TEXT NOT NULL,
			descripcion TEXT NOT NULL,
			type TEXT NOT NULL
		);
	`

	createTable_ejercicios_clases := `
		CREATE TABLE IF NOT EXISTS ejercicios_clases (
			ejercicio_id INTEGER NOT NULL,
			clase_id INTEGER NOT NULL,
			FOREIGN KEY (ejercicio_id) REFERENCES ejercicios(id),
			FOREIGN KEY (clase_id) REFERENCES clases(id),
			PRIMARY KEY (ejercicio_id, clase_id)
		);
	`
	createTable_clases_tareas := `
		CREATE TABLE IF NOT EXISTS clases_tareas (
			clase_id INTEGER NOT NULL,
			tarea_id INTEGER NOT NULL,
			FOREIGN KEY (clase_id) REFERENCES clases(id),
			FOREIGN KEY (tarea_id) REFERENCES tareas(id),
			PRIMARY KEY (clase_id, tarea_id)
		);
	`
	createTable_tareas := `
		CREATE TABLE IF NOT EXISTS tareas (
			id INTEGER PRIMARY KEY ,
			name TEXT NOT NULL,
			descripcion TEXT NOT NULL
		);
	`
	// crearUsuario := `
	// 	INSERT INTO usuarios (id, nombre, apellidos, email, usuario, password)
	// 	VALUES (1, 'Adiel', 'González', 'adiel@gmail.com', 'adiel', 'Adiel123');
	// `
	// crearClase := `
	// 	INSERT INTO clases (id,name, titulo, descripcion, teacher_id, color)
	// 	VALUES ('1', 'Clase de Ejemplo', 'Título de Ejemplo', 'Descripción de Ejemplo', 1, '#FF5733');
	// `
	// crearRecurso := `
	// 	INSERT INTO recursos (id,name, url)
	// 	VALUES ('1', 'Recurso de Ejemplo', 'CP 1_Tema 1_Modelación de Problemas de PL.pdf');
	// 	INSERT INTO recursos (id,name, url)
	// 	VALUES ('2', 'Recurso de Ejemplo', 'CP 2. Tipos de arquitectura_act.pdf');
	// `
	// asignarRecurso := `
	// 	INSERT INTO  clases_recursos (clase_id,recurso_id)
	// 	VALUES ('1','2'),
	// 	VALUES ('1','1');
	// `

	// crearEjercicio := `
	// 	INSERT INTO ejercicios (id,name, descripcion, type)
	// 	VALUES ('1','Ejercicio de Ejemplo', 'Descripción del ejercicio de ejemplo', 'tipo1');
	// `
	// crearTarea := `
	// 	INSERT INTO tareas (id,name, descripcion)
	// 	VALUES ('1','Tarea de Ejemplo', 'Descripción de la tarea de ejemplo');
	// `

	_, err := db.Exec(createTable_usuarios)
	if err != nil {
		log.Fatalf("Error creando tabla usuarios: %v", err)
	}
	_, err = db.Exec(createTable_clases)
	if err != nil {
		log.Fatalf("Error creando tabla clases: %v", err)
	}
	_, err = db.Exec(createTable_recursos)
	if err != nil {
		log.Fatalf("Error creando tabla recursos: %v", err)
	}
	_, err = db.Exec(createTable_clases_recursos)
	if err != nil {
		log.Fatalf("Error creando tabla clases_recursos: %v", err)
	}
	_, err = db.Exec(createTable_ejercicios)
	if err != nil {
		log.Fatalf("Error creando tabla ejercicios: %v", err)
	}
	_, err = db.Exec(createTable_ejercicios_clases)
	if err != nil {
		log.Fatalf("Error creando tabla ejercicios_clases: %v", err)
	}
	_, err = db.Exec(createTable_tareas)
	if err != nil {
		log.Fatalf("Error creando tabla tareas: %v", err)
	}
	_, err = db.Exec(createTable_clases_tareas)
	if err != nil {
		log.Fatalf("Error creando tabla clases_tareas: %v", err)
	}

	// _, err = db.Exec(crearUsuario)
	// if err != nil {
	// 	log.Fatalf("Error creando usuario de ejemplo: %v", err)
	// }
	// _, err = db.Exec(crearClase)
	// if err != nil {
	// 	log.Fatalf("Error creando clase de ejemplo: %v", err)
	// }
	// _, err = db.Exec(crearRecurso)
	// if err != nil {
	// 	log.Fatalf("Error creando recurso de ejemplo: %v", err)
	// }
	// _, err = db.Exec(asignarRecurso)

	// _, err = db.Exec(crearEjercicio)
	// if err != nil {
	// 	log.Fatalf("Error creando ejercicio de ejemplo: %v", err)
	// }
	// _, err = db.Exec(crearTarea)
	// if err != nil {
	// 	log.Fatalf("Error creando tarea de ejemplo: %v", err)
	// }

}
