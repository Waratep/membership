package gin_server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/Waratep/membership/src/use_case"
	"github.com/gin-gonic/gin"
)

type GinServer struct {
	useCase *use_case.UseCase
	config  *ServerConfig
	server  *http.Server
}

type ServerConfig struct {
	AppVersion    string
	RequestLog    bool
	ListenAddress string
	Debug         bool
}

func New(u *use_case.UseCase, sc *ServerConfig) *GinServer {
	router := gin.Default()

	if sc.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router.GET("/", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	server := &http.Server{
		Addr:    sc.ListenAddress,
		Handler: router,
	}
	gServer := &GinServer{
		useCase: u,
		config:  sc,
		server:  server,
	}

	gServer.addRouteMember(router)

	return gServer
}

func (g GinServer) Start(wg *sync.WaitGroup) {
	wg.Add(2)
	defer wg.Done()

	go func() {
		// service connections
		defer wg.Done()

		if err := g.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Println("listen", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := g.server.Shutdown(ctx); err != nil {
		defer wg.Done()

		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Println("timeout of 1 seconds.")
	}
	log.Println("Server exiting")
}
