package main

import (
	"database/sql"
	"escambo/internal/categoria/categoriarepo"
	"escambo/internal/postagem/posthandler"
	"escambo/internal/postagem/postrepo"
	"escambo/internal/postagem/postsvc"
	"escambo/internal/usuario/usuariohandler"
	"escambo/internal/usuario/usuariorepo"
	"escambo/internal/usuario/usuariosvc"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Panic("Loading envs:", err)
	}

	postgresURI := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", postgresURI)
	if err != nil {
		log.Panic("Openning connection: ", err)
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		log.Panic("Ping DB", err)
	}

	postRepo := postrepo.NewRepository(db)
	categoriaRepo := categoriarepo.NewRepository(db)
	postService := postsvc.NewService(postRepo, categoriaRepo)
	postHandler := posthandler.NewHandler(postService)

	r := mux.NewRouter()
	r.HandleFunc("/postagens/{id}", postHandler.GetPost).Methods("GET")
	r.HandleFunc("/postagens", postHandler.UpsavePost).Methods("PUT")

	userRepo := usuariorepo.NewRepository(db)
	usuarioService := usuariosvc.NewService(userRepo)
	usuarioHandler := usuariohandler.NewHandler(usuarioService)

	r.HandleFunc("/usuarios", usuarioHandler.UpsertUsuario).Methods("PUT")

	log.Println("Servidor rodando na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
