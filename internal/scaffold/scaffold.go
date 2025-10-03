package scaffold

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func LogInfo(msg string, args ...interface{}) {
	fmt.Printf("ℹ️ [INFO] %s - %s\n", time.Now().Format("15:04:05"), fmt.Sprintf(msg, args...))
}

func LogSuccess(msg string, args ...interface{}) {
	fmt.Printf("✅ [SUCCESS] %s - %s\n", time.Now().Format("15:04:05"), fmt.Sprintf(msg, args...))
}

func LogWarning(msg string, args ...interface{}) {
	fmt.Printf("⚠️ [WARNING] %s - %s\n", time.Now().Format("15:04:05"), fmt.Sprintf(msg, args...))
}

func LogError(msg string, err error) {
	fmt.Printf("❌ [ERROR] %s - %s: %v\n", time.Now().Format("15:04:05"), msg, err)
}

var projectStructure = map[string][]string{
	"application":    {"dtos", "mappers", "usecases"},
	"domain":         {"entities", "migrate", "repository"},
	"interfaces":     {"http"},
	"infrastructure": {"config", "database", "middleware", "persistence", "redis"},
	"response":       {},
	"util":           {},
}

var templateFiles = map[string]func(string) string{
	"main.go":                              mainTpl,
	"go.mod":                               goModTpl,
	"app.env":                              envTpl,
	"interfaces/http/server.go":            serverTpl,
	"infrastructure/config/config.go":      configTpl,
	"infrastructure/database/postgres.go":  databaseTpl,
	"domain/entities/user.go":              userEntityTpl,
	"domain/repository/user_repository.go": userRepoInterfaceTpl,
	"infrastructure/persistence/user_repository_gorm.go": userRepoGormTpl,
	"application/usecases/user_usecase.go":               userUsecaseTpl,
	"application/dtos/user_dto.go":                       userDtoTpl,
	"application/mappers/user_mapper.go":                 userMapperTpl,
	"interfaces/http/users/user_handler.go":              userHandlerTpl,
	"interfaces/http/users/routes.go":                    userRoutesTpl,
	"response/response.go":                               responseTpl,
}

func CreateProject(projectName string) error {
	start := time.Now()
	moduleName := strings.TrimSpace(projectName)
	if moduleName == "" {
		return fmt.Errorf("project name cannot be empty")
	}

	if _, err := os.Stat(projectName); !os.IsNotExist(err) {
		return fmt.Errorf("directory %s already exists", projectName)
	}

	LogInfo("Starting project generation for module: %s", moduleName)

	if err := os.MkdirAll(projectName, os.ModePerm); err != nil {
		return fmt.Errorf("failed create project root: %w", err)
	}
	LogSuccess("Created project root folder: %s", projectName)

	for folder, subs := range projectStructure {
		base := filepath.Join(projectName, folder)
		if err := os.MkdirAll(base, os.ModePerm); err != nil {
			return fmt.Errorf("failed create dir %s: %w", base, err)
		}
		LogInfo("Created folder: %s", base)
		for _, s := range subs {
			subp := filepath.Join(base, s)
			if err := os.MkdirAll(subp, os.ModePerm); err != nil {
				return fmt.Errorf("failed create dir %s: %w", subp, err)
			}
			LogInfo("Created subfolder: %s", subp)
		}
	}

	for relPath, tplFn := range templateFiles {
		full := filepath.Join(projectName, relPath)
		if err := os.MkdirAll(filepath.Dir(full), os.ModePerm); err != nil {
			return fmt.Errorf("failed create parent dir for %s: %w", full, err)
		}
		content := tplFn(moduleName)
		if err := os.WriteFile(full, []byte(content), 0644); err != nil {
			return fmt.Errorf("failed write file %s: %w", full, err)
		}
		LogSuccess("Created file: %s", full)
	}

	elapsed := time.Since(start)
	LogSuccess("Project %s generated in %s", projectName, elapsed.Truncate(time.Millisecond))
	return nil
}

func GenerateUnitTest(name, folder, structName, module, testType string) error {
	LogInfo("Generating unit test for %s in %s", structName, folder)

	if err := os.MkdirAll(folder, os.ModePerm); err != nil {
		return fmt.Errorf("failed create folder %s: %w", folder, err)
	}

	tmplFile := filepath.Join(TemplatePath(), "templates", testType+".go.tmpl")
	tmpl, err := template.ParseFiles(tmplFile)
	if err != nil {
		return fmt.Errorf("failed parse template: %w", err)
	}

	testFile := filepath.Join(folder, strings.ToLower(name)+"_test.go")
	f, err := os.Create(testFile)
	if err != nil {
		return fmt.Errorf("failed create test file: %w", err)
	}
	defer f.Close()

	data := map[string]string{
		"Package":    filepath.Base(folder),
		"StructName": structName,
		"Module":     module,
	}

	if err := tmpl.Execute(f, data); err != nil {
		return fmt.Errorf("failed execute template: %w", err)
	}

	LogSuccess("Unit test generated: %s", testFile)
	return nil
}

// ----------------------- TEMPLATES -----------------------

func mainTpl(module string) string {
	return fmt.Sprintf(`package main

import (
	"log"

	"%s/interfaces/http"
)

func main() {
	srv := http.NewServer()
	if err := srv.Start(":8080"); err != nil {
		log.Fatalf("failed to start server: %%v", err)
	}
}
`, module)
}

func goModTpl(module string) string {
	return fmt.Sprintf(`module %s

go 1.21

require (
	github.com/gin-gonic/gin v1.10.0
	gorm.io/gorm v1.25.5
	gorm.io/driver/postgres v1.4.0
	github.com/spf13/viper v1.16.0
)
`, module)
}

func envTpl(module string) string {
	return `ENVIRONMENT=development
BASE_URL=http://localhost:8080
ALLOWED_ORIGINS=*
ALLOWED_METHODS=GET,POST,PUT,DELETE
AUTHORIZATION_HEADERS=Authorization

DB_USER=postgres
DB_PASSWORD=postgres
DB_HOST=localhost
DB_PORT=5432
DB_NAME=wellsdb
DB_SOURCE=postgresql://postgres:postgres@localhost:5432/wellsdb?sslmode=disable
HTTP_SERVER_ADDRESS=:8080

REDIS_ADDR=localhost:6379
REDIS_PASSWORD=

JWT_SECRET=supersecret
JWT_ISSUER=wells
ACCESS_TOKEN_TTL=15m
REFRESH_TOKEN_TTL=24h
`
}

func serverTpl(module string) string {
	return fmt.Sprintf(`package http

import (
	"fmt"
	"log"
	"os"
	"time"

	"%s/application/usecases"
	"%s/infrastructure/config"
	"%s/infrastructure/database"
	"%s/infrastructure/persistence"
	users "%s/interfaces/http/users"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Engine *gin.Engine
}

func NewServer() *Server {
	r := gin.Default()

	// Logging
	if err := os.MkdirAll("logs", os.ModePerm); err != nil {
		log.Println("warning: failed create logs dir:", err)
	}
	logFile := fmt.Sprintf("logs/%%s.log", time.Now().Format("2006-01-02"))
	if f, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666); err == nil {
		log.SetOutput(f)
	} else {
		log.Println("warning: log file not created, using stdout")
	}

	cfg := config.GetConfig()
	if err := database.InitializeDatabase(cfg.DBSource); err != nil {
		log.Fatalf("failed open db: %%v", err)
	}

	repo := persistence.NewUserRepositoryGorm(database.GetDB())
	uc := usecases.NewUserUsecase(repo)
	handler := users.NewUserHandler(uc)

	api := r.Group("/api/v1")
	users.RegisterRoutes(api, handler)

	return &Server{Engine: r}
}

func (s *Server) Start(addr string) error {
	fmt.Println("🚀 starting server at", addr)
	return s.Engine.Run(addr)
}
`, module, module, module, module, module)
}

func configTpl(module string) string {
	return `package config

import (
	"fmt"
	"sync"

	"github.com/spf13/viper"
)

var configInstance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		config, err := LoadConfig(".")
		if err != nil {
			panic("Error loading config: " + err.Error())
		}
		configInstance = &config
	})
	return configInstance
}

type Config struct {
	Environment          string ` + "`mapstructure:\"ENVIRONMENT\"`" + `
	BaseURL              string ` + "`mapstructure:\"BASE_URL\"`" + `
	AllowedOrigins       string ` + "`mapstructure:\"ALLOWED_ORIGINS\"`" + `
	AllowedMethods       string ` + "`mapstructure:\"ALLOWED_METHODS\"`" + `
	AuthorizationHeaders string ` + "`mapstructure:\"AUTHORIZATION_HEADERS\"`" + `

	DbUser     string ` + "`mapstructure:\"DB_USER\"`" + `
	DbPassword string ` + "`mapstructure:\"DB_PASSWORD\"`" + `
	DbHost     string ` + "`mapstructure:\"DB_HOST\"`" + `
	DbPort     string ` + "`mapstructure:\"DB_PORT\"`" + `
	DbName     string ` + "`mapstructure:\"DB_NAME\"`" + `
	DBSource   string ` + "`mapstructure:\"DB_SOURCE\"`" + `

	ServerPort        string ` + "`mapstructure:\"SERVER_PORT\"`" + `
	HTTPServerAddress string ` + "`mapstructure:\"HTTP_SERVER_ADDRESS\"`" + `

	REDIS_ADDR     string ` + "`mapstructure:\"REDIS_ADDR\"`" + `
	REDIS_PASSWORD string ` + "`mapstructure:\"REDIS_PASSWORD\"`" + `

	JWTSecret       string ` + "`mapstructure:\"JWT_SECRET\"`" + `
	JWTIssuer       string ` + "`mapstructure:\"JWT_ISSUER\"`" + `
	AccessTokenTTL  string ` + "`mapstructure:\"ACCESS_TOKEN_TTL\"`" + `
	RefreshTokenTTL string ` + "`mapstructure:\"REFRESH_TOKEN_TTL\"`" + `
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	if err = viper.ReadInConfig(); err != nil {
		// fallback ke env var saja kalau file tidak ada
	}
	err = viper.Unmarshal(&config)

	// kalau DBSource kosong, generate dari variabel yang lain
	if config.DBSource == "" {
		config.DBSource = fmt.Sprintf(
			"postgresql://%s:%s@%s:%s/%s?sslmode=disable",
			config.DbUser,
			config.DbPassword,
			config.DbHost,
			config.DbPort,
			config.DbName,
		)
	}

	return
}
`
}

func databaseTpl(module string) string {
	return `package database

import (
	"fmt"
	"` + module + `/domain/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitializeDatabase(dsn string) error {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to PostgreSQL: %w", err)
	}

	DB = db
	if err := DB.AutoMigrate(&entities.User{}); err != nil {
		return fmt.Errorf("failed to migrate: %w", err)
	}

	fmt.Println("✅ PostgreSQL connected and migrated")
	return nil
}

func GetDB() *gorm.DB {
	return DB
}
`
}

func userEntityTpl(module string) string {
	_ = module
	return `package entities

import "time"

type User struct {
	ID        uint      ` + "`gorm:\"primaryKey\" json:\"id\"`" + `
	Name      string    ` + "`json:\"name\"`" + `
	Email     string    ` + "`json:\"email\" gorm:\"unique\"`" + `
	CreatedAt time.Time ` + "`json:\"created_at\"`" + `
	UpdatedAt time.Time ` + "`json:\"updated_at\"`" + `
}
`
}

func userRepoInterfaceTpl(module string) string {
	return `package repository

import "` + module + `/domain/entities"

type UserRepository interface {
	Create(user *entities.User) error
	FindAll() ([]entities.User, error)
	FindByID(id uint) (*entities.User, error)
	Update(user *entities.User) error
	Delete(id uint) error
}
`
}

func userRepoGormTpl(module string) string {
	return `package persistence

import (
	"` + module + `/domain/entities"
	"` + module + `/domain/repository"

	"gorm.io/gorm"
)

type UserRepositoryGorm struct {
	db *gorm.DB
}

func NewUserRepositoryGorm(db *gorm.DB) repository.UserRepository {
	return &UserRepositoryGorm{db: db}
}

func (r *UserRepositoryGorm) Create(user *entities.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepositoryGorm) FindAll() ([]entities.User, error) {
	var users []entities.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepositoryGorm) FindByID(id uint) (*entities.User, error) {
	var user entities.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryGorm) Update(user *entities.User) error {
	return r.db.Save(user).Error
}

func (r *UserRepositoryGorm) Delete(id uint) error {
	return r.db.Delete(&entities.User{}, id).Error
}
`
}

func userUsecaseTpl(module string) string {
	return `package usecases

import (
	"` + module + `/application/dtos"
	"` + module + `/application/mappers"
	"` + module + `/domain/repository"
)

type UserUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) *UserUsecase {
	return &UserUsecase{repo: repo}
}

func (u *UserUsecase) CreateUser(dto *dtos.UserDTO) error {
	user := mappers.ToUserEntity(dto)
	return u.repo.Create(user)
}

func (u *UserUsecase) GetUsers() ([]dtos.UserDTO, error) {
	users, err := u.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return mappers.ToUserDTOs(users), nil
}

func (u *UserUsecase) GetUserByID(id uint) (*dtos.UserDTO, error) {
	user, err := u.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return mappers.ToUserDTO(user), nil
}

func (u *UserUsecase) UpdateUser(dto *dtos.UserDTO) error {
	user := mappers.ToUserEntity(dto)
	return u.repo.Update(user)
}

func (u *UserUsecase) DeleteUser(id uint) error {
	return u.repo.Delete(id)
}
`
}

func userDtoTpl(module string) string {
	_ = module
	return `package dtos

type UserDTO struct {
	ID    uint   ` + "`json:\"id\"`" + `
	Name  string ` + "`json:\"name\"`" + `
	Email string ` + "`json:\"email\"`" + `
}
`
}

func userMapperTpl(module string) string {
	_ = module
	return `package mappers

import (
	"` + module + `/application/dtos"
	"` + module + `/domain/entities"
)

func ToUserDTO(user *entities.User) *dtos.UserDTO {
	return &dtos.UserDTO{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}

func ToUserDTOs(users []entities.User) []dtos.UserDTO {
	var result []dtos.UserDTO
	for _, u := range users {
		result = append(result, dtos.UserDTO{
			ID:    u.ID,
			Name:  u.Name,
			Email: u.Email,
		})
	}
	return result
}

func ToUserEntity(dto *dtos.UserDTO) *entities.User {
	return &entities.User{
		ID:    dto.ID,
		Name:  dto.Name,
		Email: dto.Email,
	}
}
`
}

func userHandlerTpl(module string) string {
	return `package users

import (
	"net/http"
	"strconv"

	"` + module + `/application/dtos"
	"` + module + `/application/usecases"
	"` + module + `/response"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	uc *usecases.UserUsecase
}

func NewUserHandler(uc *usecases.UserUsecase) *UserHandler {
	return &UserHandler{uc: uc}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req dtos.UserDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "invalid request", err.Error())
		return
	}
	if err := h.uc.CreateUser(&req); err != nil {
		response.Error(c, http.StatusInternalServerError, "failed create user", err.Error())
		return
	}
	response.Success(c, "user created", req)
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.uc.GetUsers()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "failed fetch users", err.Error())
		return
	}
	response.Success(c, "users list", users)
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := h.uc.GetUserByID(uint(id))
	if err != nil {
		response.Error(c, http.StatusNotFound, "user not found", err.Error())
		return
	}
	response.Success(c, "user found", user)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req dtos.UserDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "invalid request", err.Error())
		return
	}
	req.ID = uint(id)
	if err := h.uc.UpdateUser(&req); err != nil {
		response.Error(c, http.StatusInternalServerError, "failed update user", err.Error())
		return
	}
	response.Success(c, "user updated", req)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.uc.DeleteUser(uint(id)); err != nil {
		response.Error(c, http.StatusInternalServerError, "failed delete user", err.Error())
		return
	}
	response.Success(c, "user deleted", nil)
}
`
}

func userRoutesTpl(module string) string {
	_ = module
	return `package users

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(rg *gin.RouterGroup, h *UserHandler) {
	r := rg.Group("/users")
	r.POST("/", h.CreateUser)
	r.GET("/", h.GetUsers)
	r.GET("/:id", h.GetUserByID)
	r.PUT("/:id", h.UpdateUser)
	r.DELETE("/:id", h.DeleteUser)
}
`
}

func responseTpl(module string) string {
	_ = module
	return `package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": message,
		"data":    data,
	})
}

func Error(c *gin.Context, status int, message string, err interface{}) {
	c.JSON(status, gin.H{
		"success": false,
		"message": message,
		"error":   err,
	})
}
`
}
