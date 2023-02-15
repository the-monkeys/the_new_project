package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/89minutes/the_new_project/services/api_gateway/config"
	"github.com/89minutes/the_new_project/services/api_gateway/pkg/article"
	"github.com/89minutes/the_new_project/services/api_gateway/pkg/auth"
	"github.com/89minutes/the_new_project/services/api_gateway/pkg/blogsandposts"
	"github.com/89minutes/the_new_project/services/api_gateway/pkg/user_service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Server struct {
	router *gin.Engine
}

func newServer() *Server {
	return &Server{router: gin.New()}
}

func main() {
	// Load API Gateway configuration
	cfg, err := config.LoadGatewayConfig()
	if err != nil {
		logrus.Fatalf("failed to load the config: %v", err)
	}

	// Set Gin to Release mode
	gin.SetMode(gin.ReleaseMode)

	// Create a gin router and add the Recovery middleware to recover from panics
	server := newServer()
	server.router = gin.New()
	server.router.Use(gin.Recovery())
	server.router.Use(gin.Logger())

	// enable CORS
	server.router.Use(cors.Default())

	// Register REST routes for all the microservice
	authClient := auth.RegisterRouter(server.router, &cfg)
	authClient.Log.SetReportCaller(true)
	authClient.Log.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
		FullTimestamp: false,
	})

	article.RegisterArticleRoutes(server.router, &cfg, authClient)
	user_service.RegisterUserRouter(server.router, &cfg, authClient)
	blogsandposts.RegisterBlogRouter(server.router, &cfg, authClient)

	server.start(context.Background(), cfg)

}

func (s *Server) start(ctx context.Context, addr config.Address) {
	// TLS certificate and key
	var tlsCert, tlsKey string
	if os.Getenv("NO_TLS") != "1" {
		tlsCert = os.Getenv("TLS_CERT")
		if tlsCert == "" {
			tlsCert = "/the_monkeys/vault/certs/cert.pem"
		}
		tlsKey = os.Getenv("TLS_KEY")
		if tlsKey == "" {
			tlsKey = "/the_monkeys/vault/certs/prv_key.pemy"
		}
	}

	// Launch the server (this is a blocking call)
	s.launchServer(ctx, addr, tlsCert, tlsKey)
}

// Start the server
func (s *Server) launchServer(ctx context.Context, addr config.Address, tlsCert, tlsKey string) {
	// If we don't have a TLS certificate, don't enable TLS
	enableTLS := (tlsCert != "" && tlsKey != "")

	// HTTP server (no TLS)
	httpSrv := &http.Server{
		Addr:           addr.APIGatewayHTTP,
		Handler:        s.router,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	// HTTPS server (with TLS)
	httpsSrv := &http.Server{
		Addr:           addr.APIGatewayHTTPS,
		Handler:        s.router,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	// Start the HTTP server in a background goroutine
	go func() {
		logrus.Printf("HTTP server listening at http://%s\n", addr.APIGatewayHTTP)
		// Next call blocks until the server is shut down
		err := httpSrv.ListenAndServe()
		if err != http.ErrServerClosed {
			panic(err)
		}
	}()

	// Start the HTTPS server in a background goroutine
	if enableTLS {
		go func() {
			logrus.Printf("HTTPS server listening at https://%s\n", addr.APIGatewayHTTPS)
			err := httpsSrv.ListenAndServeTLS(tlsCert, tlsKey)
			if err != http.ErrServerClosed {
				panic(err)
			}
		}()
	}

	// Listen to SIGINT and SIGTERM signals
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

	// Block until we either get a termination signal, or until the context is canceled
	select {
	case <-ctx.Done():
	case <-ch:
	}

	// We received an interrupt signal, shut down both servers
	var errHttp, errHttps error
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	errHttp = httpSrv.Shutdown(shutdownCtx)
	if enableTLS {
		errHttps = httpsSrv.Shutdown(shutdownCtx)
	}
	shutdownCancel()
	// Log the errors (could be context canceled)
	if errHttp != nil {
		logrus.Println("HTTP server shutdown error:", errHttp)
	}
	if errHttps != nil {
		logrus.Println("HTTPS server shutdown error:", errHttps)
	}
}
