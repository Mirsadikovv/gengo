package internal

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/fobus1289/ufa_shared/gengo/service"
	"github.com/fobus1289/ufa_shared/gengo/stuble"
	"github.com/iancoleman/strcase"
)

// NewProject creates a full new project from scratch following the primer architecture.
// It generates: root files, src/main.go, auth_service module, and the first module with its first entity.
func NewProject(projectName, moduleName, entityName, modPath string) {
	projectDir := strcase.ToSnake(projectName)

	data := map[string]interface{}{
		"ProjectName": projectName,
		"ModuleName":  moduleName,
		"EntityName":  entityName,
		"ModPath":     modPath,
	}

	// --- Root directory ---
	mustMkdir(projectDir)

	// Root main.go
	mustRender(stuble.MainEntry, data, projectDir, "main.go")

	// dev.go / prod.go
	mustRender(stuble.Dev, data, projectDir, "dev.go")
	mustRender(stuble.Prod, data, projectDir, "prod.go")

	// .env.example
	mustRender(stuble.Env, data, projectDir, ".env.example")

	// .gitignore
	mustRenderRaw(stuble.Gitignore, projectDir, ".gitignore")

	// Makefile
	mustRender(stuble.Makefile, data, projectDir, "Makefile")

	// --- src/ ---
	mustMkdir(projectDir, "src")
	mustRender(stuble.SrcMain, data, projectDir, "src", "main.go")

	// --- auth_service module (bootstrap) ---
	authDtoDir := filepath.Join(projectDir, "src", "module", "auth_service", "dto")
	authMiddlewareDir := filepath.Join(projectDir, "src", "module", "auth_service", "middleware")
	mustMkdir(authDtoDir)
	mustMkdir(authMiddlewareDir)
	mustRenderRaw(stuble.AuthDto, authDtoDir, "auth_dto.go")
	mustRender(stuble.AuthMiddleware, data, authMiddlewareDir, "auth_middleware.go")

	// --- First module ---
	moduleSnake := strcase.ToSnake(moduleName)
	entitySnake := strcase.ToSnake(entityName)

	moduleDir := filepath.Join(projectDir, "src", "module", moduleSnake)
	dtoDir := filepath.Join(moduleDir, "dto")
	modelDir := filepath.Join(moduleDir, "model")
	handlerDir := filepath.Join(moduleDir, "handler")
	serviceDir := filepath.Join(moduleDir, "service")

	for _, d := range []string{dtoDir, modelDir, handlerDir, serviceDir} {
		mustMkdir(d)
	}

	// cmd.go at module root
	mustRender(stuble.Cmd, data, moduleDir, "cmd.go")

	// dto
	mustRender(stuble.Dto, data, dtoDir, entitySnake+"_dto.go")

	// model
	mustRender(stuble.Model, data, modelDir, entitySnake+"_model.go")

	// handler
	mustRender(stuble.Handler, data, handlerDir, entitySnake+"_handler.go")

	// service
	mustRender(stuble.Service, data, serviceDir, entitySnake+"_service.go")

	// --- go mod init + tidy ---
	runCmd(projectDir, "go", "mod", "init", modPath)
	runCmd(projectDir, "go", "mod", "tidy")

	// --- goimports ---
	runGoimports(projectDir)

	fmt.Printf("\n✓ Project '%s' created successfully!\n", projectDir)
	fmt.Printf("  cd %s && go run -tags dev .\n", projectDir)
}

// --- helpers ---

func mustMkdir(parts ...string) {
	if err := service.MkdirAll(parts...); err != nil {
		log.Fatalf("mkdir %v: %v", parts, err)
	}
}

func mustRender(tmpl string, data map[string]interface{}, parts ...string) {
	if err := service.RenderToFile(tmpl, data, parts...); err != nil {
		log.Fatalf("render %v: %v", parts, err)
	}
}

func mustRenderRaw(content string, parts ...string) {
	filePath := filepath.Join(parts...)
	if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
		log.Fatalf("write %v: %v", filePath, err)
	}
}

func runCmd(dir string, name string, args ...string) {
	cmd := exec.Command(name, args...)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Printf("warning: %s %v: %v", name, args, err)
	}
}

func runGoimports(dir string) {
	// Install goimports if not present (best effort)
	installCmd := exec.Command("go", "install", "golang.org/x/tools/cmd/goimports@latest")
	installCmd.Stdout = os.Stdout
	installCmd.Stderr = os.Stderr
	_ = installCmd.Run()

	cmd := exec.Command("goimports", "-w", ".")
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Printf("warning: goimports: %v", err)
	}
}
