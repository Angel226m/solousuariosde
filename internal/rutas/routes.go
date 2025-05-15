/*package rutas

import (
	"sistema-tours/internal/config"
	"sistema-tours/internal/controladores"
	"sistema-tours/internal/middleware"
	"strconv"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configura todas las rutas de la API
func SetupRoutes(
	router *gin.Engine,
	config *config.Config,
	authController *controladores.AuthController,
	usuarioController *controladores.UsuarioController,
	embarcacionController *controladores.EmbarcacionController,
	// Otros controladores
) {
	// Middleware global
	router.Use(middleware.LoggerMiddleware())
	router.Use(middleware.ErrorMiddleware())
	router.Use(gin.Recovery())

	// Rutas públicas
	public := router.Group("/api/v1")
	{
		// Autenticación
		public.POST("/auth/login", authController.Login)
		public.POST("/auth/refresh", authController.RefreshToken)

		// Registrarse como cliente
		// Para rol CLIENTE únicamente
		// public.POST("/registro", usuarioController.Register)
	}

	// Rutas protegidas (requieren autenticación)
	protected := router.Group("/api/v1")
	protected.Use(middleware.AuthMiddleware(config))
	{
		// Cambiar contraseña (cualquier usuario autenticado)
		protected.POST("/auth/change-password", authController.ChangePassword)

		// Usuarios - Admin
		admin := protected.Group("/admin")
		admin.Use(middleware.RoleMiddleware("ADMIN"))
		{
			// Gestión de usuarios
			admin.POST("/usuarios", usuarioController.Create)
			admin.GET("/usuarios", usuarioController.List)
			admin.GET("/usuarios/:id", usuarioController.GetByID)
			admin.PUT("/usuarios/:id", usuarioController.Update)
			admin.DELETE("/usuarios/:id", usuarioController.Delete)
			admin.GET("/usuarios/rol/:rol", usuarioController.ListByRol)

			// Gestión de embarcaciones
			admin.POST("/embarcaciones", embarcacionController.Create)
			admin.GET("/embarcaciones", embarcacionController.List)
			admin.GET("/embarcaciones/:id", embarcacionController.GetByID)
			admin.PUT("/embarcaciones/:id", embarcacionController.Update)
			admin.DELETE("/embarcaciones/:id", embarcacionController.Delete)
			admin.GET("/embarcaciones/chofer/:idChofer", embarcacionController.ListByChofer)
		}

		// Vendedores
		vendedor := protected.Group("/vendedor")
		vendedor.Use(middleware.RoleMiddleware("ADMIN", "VENDEDOR"))
		{
			// Ver embarcaciones (solo lectura)
			vendedor.GET("/embarcaciones", embarcacionController.List)
			vendedor.GET("/embarcaciones/:id", embarcacionController.GetByID)
		}

		// Choferes
		chofer := protected.Group("/chofer")
		chofer.Use(middleware.RoleMiddleware("ADMIN", "CHOFER"))
		{
			// Ver embarcaciones asignadas
			chofer.GET("/mis-embarcaciones", func(ctx *gin.Context) {
				// Obtener ID del usuario autenticado del contexto
				userID := ctx.GetInt("user_id")
				// Redirigir a la ruta que obtiene embarcaciones por chofer
				ctx.Request.URL.Path = "/api/v1/admin/embarcaciones/chofer/" + strconv.Itoa(userID)
				router.HandleContext(ctx)
			})
		}

		// Clientes
		cliente := protected.Group("/cliente")
		cliente.Use(middleware.RoleMiddleware("ADMIN", "CLIENTE"))
		{
			// En el futuro, rutas específicas para clientes
		}
	}
}
*/

package rutas

import (
	"sistema-tours/internal/config"
	"sistema-tours/internal/controladores"
	"sistema-tours/internal/middleware"
	"strconv"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configura todas las rutas de la API
func SetupRoutes(
	router *gin.Engine,
	config *config.Config,
	authController *controladores.AuthController,
	usuarioController *controladores.UsuarioController,
	embarcacionController *controladores.EmbarcacionController,
	tourController *controladores.TourController,
	// Otros controladores
) {
	// Middleware global
	router.Use(middleware.LoggerMiddleware())
	router.Use(middleware.ErrorMiddleware())
	router.Use(gin.Recovery())

	// Rutas públicas
	public := router.Group("/api/v1")
	{
		// Autenticación
		public.POST("/auth/login", authController.Login)
		public.POST("/auth/refresh", authController.RefreshToken)

		// Registrarse como cliente
		// Para rol CLIENTE únicamente
		// public.POST("/registro", usuarioController.Register)
	}

	// Rutas protegidas (requieren autenticación)
	protected := router.Group("/api/v1")
	protected.Use(middleware.AuthMiddleware(config))
	{
		// Cambiar contraseña (cualquier usuario autenticado)
		protected.POST("/auth/change-password", authController.ChangePassword)

		// Usuarios - Admin
		admin := protected.Group("/admin")
		admin.Use(middleware.RoleMiddleware("ADMIN"))
		{
			// Gestión de usuarios
			admin.POST("/usuarios", usuarioController.Create)
			admin.GET("/usuarios", usuarioController.List)
			admin.GET("/usuarios/:id", usuarioController.GetByID)
			admin.PUT("/usuarios/:id", usuarioController.Update)
			admin.DELETE("/usuarios/:id", usuarioController.Delete)
			admin.GET("/usuarios/rol/:rol", usuarioController.ListByRol)

			// Gestión de embarcaciones
			admin.POST("/embarcaciones", embarcacionController.Create)
			admin.GET("/embarcaciones", embarcacionController.List)
			admin.GET("/embarcaciones/:id", embarcacionController.GetByID)
			admin.PUT("/embarcaciones/:id", embarcacionController.Update)
			admin.DELETE("/embarcaciones/:id", embarcacionController.Delete)
			admin.GET("/embarcaciones/chofer/:idChofer", embarcacionController.ListByChofer)

			// Gestión de tipos de tour
			admin.POST("/tipos-tour", tourController.CreateTipoTour)
			admin.GET("/tipos-tour", tourController.ListTiposTour)
			admin.GET("/tipos-tour/:id", tourController.GetTipoTourByID)
			admin.PUT("/tipos-tour/:id", tourController.UpdateTipoTour)
			admin.DELETE("/tipos-tour/:id", tourController.DeleteTipoTour)

			// Gestión de horarios de tour
			admin.POST("/horarios", tourController.CreateHorario)
			admin.GET("/horarios", tourController.ListHorarios)
			admin.GET("/horarios/:id", tourController.GetHorarioByID)
			admin.PUT("/horarios/:id", tourController.UpdateHorario)
			admin.DELETE("/horarios/:id", tourController.DeleteHorario)

			// Esta ruta usaba :idTipoTour que causaba el conflicto, cambiamos a :id para mantener consistencia
			admin.GET("/tipos-tour/:id/horarios", tourController.GetHorariosByTipoTourID)
		}

		// Vendedores
		vendedor := protected.Group("/vendedor")
		vendedor.Use(middleware.RoleMiddleware("ADMIN", "VENDEDOR"))
		{
			// Ver embarcaciones (solo lectura)
			vendedor.GET("/embarcaciones", embarcacionController.List)
			vendedor.GET("/embarcaciones/:id", embarcacionController.GetByID)

			// Ver tipos de tour (solo lectura)
			vendedor.GET("/tipos-tour", tourController.ListTiposTour)
			vendedor.GET("/tipos-tour/:id", tourController.GetTipoTourByID)

			// Ver horarios (solo lectura)
			vendedor.GET("/horarios", tourController.ListHorarios)
			vendedor.GET("/horarios/:id", tourController.GetHorarioByID)

			// Aquí también cambiamos :idTipoTour a :id para mantener consistencia
			vendedor.GET("/tipos-tour/:id/horarios", tourController.GetHorariosByTipoTourID)
		}

		// Choferes
		chofer := protected.Group("/chofer")
		chofer.Use(middleware.RoleMiddleware("ADMIN", "CHOFER"))
		{
			// Ver embarcaciones asignadas
			chofer.GET("/mis-embarcaciones", func(ctx *gin.Context) {
				// Obtener ID del usuario autenticado del contexto
				userID := ctx.GetInt("user_id")
				// Redirigir a la ruta que obtiene embarcaciones por chofer
				ctx.Request.URL.Path = "/api/v1/admin/embarcaciones/chofer/" + strconv.Itoa(userID)
				router.HandleContext(ctx)
			})

			// Ver tipos de tour (solo lectura)
			chofer.GET("/tipos-tour", tourController.ListTiposTour)
			chofer.GET("/tipos-tour/:id", tourController.GetTipoTourByID)
		}

		// Clientes
		cliente := protected.Group("/cliente")
		cliente.Use(middleware.RoleMiddleware("ADMIN", "CLIENTE"))
		{
			// Ver tipos de tour disponibles (solo lectura)
			cliente.GET("/tipos-tour", tourController.ListTiposTour)
			cliente.GET("/tipos-tour/:id", tourController.GetTipoTourByID)

			// Aquí también cambiamos :idTipoTour a :id para mantener consistencia
			cliente.GET("/tipos-tour/:id/horarios", tourController.GetHorariosByTipoTourID)
		}
	}
}
