package main
import (
	"path/filepath"
    "net/http"
    "os"
    "github.com/go-chi/chi"
    "github.com/go-chi/chi/middleware"
)

func main()  {
    r := chi.NewRouter()

    r.Use(middleware.RequestID)
    r.Use(middleware.RealIP)
    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)

    workDir, _ := os.Getwd()
    publicDir := filepath.Join(workDir, "public")

    fs := http.StripPrefix("", http.FileServer(http.Dir(publicDir)))

    r.Handle("/*", fs)

    http.ListenAndServe(":9000", r)
}