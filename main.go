package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber"
	uuid "github.com/jackc/pgx/pgtype/ext/gofrs-uuid"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Task struct {
	ID          uuid.UUID  `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description,omitempty"`
	Status      TaskStatus `json:"status"`
	Deadline    time.Time  `json:"deadline,omitempty"`
	CreatedAt   time.Time  `json:"created_at,omitempty"`
	UpdatedAt   time.Time  `json:"updated_at,omitempty"`
}

type TaskStatus string

const (
	TaskStatusTodo    TaskStatus = "TODO"
	TaskStatusDoing   TaskStatus = "DOING"
	TaskStatusDone    TaskStatus = "DONE"
	TaskStatusDeleted TaskStatus = "DELETED"
)

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/")
	api.Post("/tasks", r.CreateTask)
	api.Get("/tasks", r.GetTasks)
	api.Get("/tasks/:id", r.GetTask)
	api.Patch("/tasks/:id", r.UpdateTask)
	api.Delete("/tasks/:id", r.DeleteTask)
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	db, err := storage.NewConnection(config)
	if err != nil {
		log.Fatal("Error connecting to database", err)
	}

	r := Repository{
		DB: db,
	}

	app := fiber.New()
	r.SetupRoutes(app)
	app.Listen("8000")
}
