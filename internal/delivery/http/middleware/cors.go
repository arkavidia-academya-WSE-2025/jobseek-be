package middleware

import "github.com/gofiber/fiber/v2/middleware/cors"

// CORSConfig sets up the CORS middleware
func CORSConfig() cors.Config {
	return cors.Config{
		AllowOrigins: "*", // Change this for production (e.g., "https://example.com")
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}
}
