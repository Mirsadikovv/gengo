package internal

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/fobus1289/ufa_shared/gengo/service"
	"github.com/fobus1289/ufa_shared/gengo/stuble"
	"github.com/iancoleman/strcase"
)

// AddModule adds a new module with its first entity to an existing project.
// It generates module files under src/module/{moduleName}/ and updates src/main.go.
func AddModule(moduleName, entityName string) {
	// Read mod path from go.mod
	modPath, err := readModPath("go.mod")
	if err != nil {
		log.Fatalf("could not read go.mod: %v\nMake sure you run gengo --add from the project root directory.", err)
	}

	moduleSnake := strcase.ToSnake(moduleName)
	entitySnake := strcase.ToSnake(entityName)
	entityCamel := strcase.ToCamel(entityName)

	data := map[string]interface{}{
		"ModuleName": moduleName,
		"EntityName": entityName,
		"ModPath":    modPath,
	}

	// --- Create module directory structure ---
	moduleDir := filepath.Join("src", "module", moduleSnake)
	dtoDir := filepath.Join(moduleDir, "dto")
	modelDir := filepath.Join(moduleDir, "model")
	handlerDir := filepath.Join(moduleDir, "handler")
	serviceDir := filepath.Join(moduleDir, "service")

	for _, d := range []string{dtoDir, modelDir, handlerDir, serviceDir} {
		if err := service.MkdirAll(d); err != nil {
			log.Fatalf("mkdir %s: %v", d, err)
		}
	}

	// cmd.go at module root
	if err := service.RenderToFile(stuble.Cmd, data, moduleDir, "cmd.go"); err != nil {
		log.Fatalf("render cmd.go: %v", err)
	}

	// dto
	if err := service.RenderToFile(stuble.Dto, data, dtoDir, entitySnake+"_dto.go"); err != nil {
		log.Fatalf("render dto: %v", err)
	}

	// model
	if err := service.RenderToFile(stuble.Model, data, modelDir, entitySnake+"_model.go"); err != nil {
		log.Fatalf("render model: %v", err)
	}

	// handler
	if err := service.RenderToFile(stuble.Handler, data, handlerDir, entitySnake+"_handler.go"); err != nil {
		log.Fatalf("render handler: %v", err)
	}

	// service
	if err := service.RenderToFile(stuble.Service, data, serviceDir, entitySnake+"_service.go"); err != nil {
		log.Fatalf("render service: %v", err)
	}

	// --- Update src/main.go ---
	mainGoPath := filepath.Join("src", "main.go")
	if err := service.InjectModule(mainGoPath, moduleSnake, entitySnake, entityCamel, modPath); err != nil {
		log.Fatalf("update src/main.go: %v\nTip: make sure src/main.go contains @gengo:modules and @gengo:migration markers.", err)
	}

	// --- go mod tidy + goimports ---
	runCmd(".", "go", "mod", "tidy")
	runGoimports(".")

	fmt.Printf("\n✓ Module '%s' with entity '%s' added successfully!\n", moduleSnake, entityCamel)
	fmt.Printf("  Don't forget to check src/main.go for the injected imports and calls.\n")
}

// readModPath reads the module path from go.mod file.
func readModPath(goModPath string) (string, error) {
	f, err := os.Open(goModPath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "module ") {
			return strings.TrimPrefix(line, "module "), nil
		}
	}
	return "", fmt.Errorf("module directive not found in go.mod")
}
