/*package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sistema-tours/internal/config"
	"sistema-tours/internal/controladores"
	"sistema-tours/internal/repositorios"
	"sistema-tours/internal/rutas"
	"sistema-tours/internal/servicios"
	"sistema-tours/internal/utils"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	// Cargar configuración
	cfg := config.LoadConfig()

	// Configurar modo de Gin según entorno
	if cfg.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Inicializar router
	router := gin.Default()

	// Inicializar validador
	utils.InitValidator()

	// Conectar a la base de datos
	db, err := connectDB(cfg)
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}
	defer db.Close()

	// Inicializar repositorios
	usuarioRepo := repositorios.NewUsuarioRepository(db)
	embarcacionRepo := repositorios.NewEmbarcacionRepository(db)
	// Otros repositorios...

	// Inicializar servicios
	authService := servicios.NewAuthService(usuarioRepo, cfg)
	usuarioService := servicios.NewUsuarioService(usuarioRepo)
	embarcacionService := servicios.NewEmbarcacionService(embarcacionRepo, usuarioRepo)
	// Otros servicios...

	// Inicializar controladores
	authController := controladores.NewAuthController(authService)
	usuarioController := controladores.NewUsuarioController(usuarioService)
	embarcacionController := controladores.NewEmbarcacionController(embarcacionService)
	// Otros controladores...

	// Configurar rutas
	rutas.SetupRoutes(
		router,
		cfg,
		authController,
		usuarioController,
		embarcacionController,
		// Otros controladores...
	)

	// Iniciar servidor
	serverAddr := fmt.Sprintf("%s:%s", cfg.ServerHost, cfg.ServerPort)
	log.Printf("Servidor iniciado en %s", serverAddr)
	if err := router.Run(serverAddr); err != nil {
		log.Fatalf("Error al iniciar servidor: %v", err)
	}
}

// connectDB establece conexión con la base de datos PostgreSQL
func connectDB(cfg *config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBSSLMode,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	// Verificar conexión
	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Conexión exitosa a la base de datos")
	return db, nil
}

func runMigrations(db *sql.DB) error {
	migrationFile, err := os.ReadFile("./migrations/crear_tablas.sql")
	if err != nil {
		return err
	}

	_, err = db.Exec(string(migrationFile))
	if err != nil {
		return err
	}

	log.Println("Migraciones ejecutadas exitosamente")
	return nil
}
*/

package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sistema-tours/internal/config"
	"sistema-tours/internal/controladores"
	"sistema-tours/internal/repositorios"
	"sistema-tours/internal/rutas"
	"sistema-tours/internal/servicios"
	"sistema-tours/internal/utils"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	// Cargar configuración
	cfg := config.LoadConfig()

	// Configurar modo de Gin según entorno
	if cfg.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Inicializar router
	router := gin.Default()

	// Inicializar validador
	utils.InitValidator()

	// Conectar a la base de datos
	db, err := connectDB(cfg)
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}
	defer db.Close()

	// Inicializar repositorios
	usuarioRepo := repositorios.NewUsuarioRepository(db)
	embarcacionRepo := repositorios.NewEmbarcacionRepository(db)
	tipoTourRepo := repositorios.NewTipoTourRepository(db)
	// Otros repositorios...

	// Inicializar servicios
	authService := servicios.NewAuthService(usuarioRepo, cfg)
	usuarioService := servicios.NewUsuarioService(usuarioRepo)
	embarcacionService := servicios.NewEmbarcacionService(embarcacionRepo, usuarioRepo)
	tourService := servicios.NewTourService(tipoTourRepo)
	// Otros servicios...

	// Inicializar controladores
	authController := controladores.NewAuthController(authService)
	usuarioController := controladores.NewUsuarioController(usuarioService)
	embarcacionController := controladores.NewEmbarcacionController(embarcacionService)
	tourController := controladores.NewTourController(tourService)
	// Otros controladores...

	// Configurar rutas
	rutas.SetupRoutes(
		router,
		cfg,
		authController,
		usuarioController,
		embarcacionController,
		tourController,
		// Otros controladores...
	)

	// Iniciar servidor
	serverAddr := fmt.Sprintf("%s:%s", cfg.ServerHost, cfg.ServerPort)
	log.Printf("Servidor iniciado en %s", serverAddr)
	if err := router.Run(serverAddr); err != nil {
		log.Fatalf("Error al iniciar servidor: %v", err)
	}
}

// connectDB establece conexión con la base de datos PostgreSQL
func connectDB(cfg *config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBSSLMode,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	// Verificar conexión
	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Conexión exitosa a la base de datos")
	return db, nil
}

func runMigrations(db *sql.DB) error {
	migrationFile, err := os.ReadFile("./migrations/crear_tablas.sql")
	if err != nil {
		return err
	}

	_, err = db.Exec(string(migrationFile))
	if err != nil {
		return err
	}

	log.Println("Migraciones ejecutadas exitosamente")
	return nil
}
