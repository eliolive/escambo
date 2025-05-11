package main

import (
	"database/sql"
	"escambo/internal/postagem/postagemhandler"
	"escambo/internal/postagem/postagemrepo"
	"escambo/internal/postagem/postagemsvc"
	"escambo/internal/proposta/propostahandler"
	"escambo/internal/proposta/propostarepo"
	"escambo/internal/proposta/propostasvc"
	"escambo/internal/usuario/usuariohandler"
	"escambo/internal/usuario/usuariorepo"
	"escambo/internal/usuario/usuariosvc"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"github.com/rs/cors"

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

	postRepo := postagemrepo.NewRepository(db)
	postService := postagemsvc.NewService(postRepo)
	postHandler := postagemhandler.NewHandler(postService)

	r := mux.NewRouter()
	r.HandleFunc("/postagens/{id}", postHandler.GetPost).Methods("GET")
	r.HandleFunc("/postagens", postHandler.UpsavePost).Methods("PUT")

	userRepo := usuariorepo.NewRepository(db)
	usuarioService := usuariosvc.NewService(userRepo)
	usuarioHandler := usuariohandler.NewHandler(usuarioService)

	r.HandleFunc("/usuarios", usuarioHandler.UpsertUsuario).Methods("PUT")

	propostaRepo := propostarepo.NewRepository(db)
	propostaService := propostasvc.NewService(propostaRepo)
	propostaHandler := propostahandler.NewHandler(&propostaService)

	r.HandleFunc("/propostas", propostaHandler.UpsaveProposta).Methods("PUT")
	r.HandleFunc("/propostas", propostaHandler.GetPropostas).Methods("GET")

	corsHandler := cors.New(cors.Options{AllowedOrigins: []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	}).Handler(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Servidor rodando na porta " + port + "...")
	log.Fatal(http.ListenAndServe(":"+port, corsHandler))

}
