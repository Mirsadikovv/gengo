package service

import (
	"fmt"
	"os"
	"strings"
)

const (
	modulesMarker   = "// @gengo:modules"
	migrationMarker = "// @gengo:migration"
)

// InjectModule adds a module Cmd call and model migration into src/main.go.
// It looks for @gengo:modules and @gengo:migration markers and inserts after them.
func InjectModule(mainGoPath, moduleSnake, entitySnake, entityCamel, modPath string) error {
	content, err := os.ReadFile(mainGoPath)
	if err != nil {
		return fmt.Errorf("read %s: %w", mainGoPath, err)
	}

	src := string(content)

	// 1. Inject cmd call after @gengo:modules marker
	cmdCall := fmt.Sprintf(
		"\t\t%s_cmd.Cmd(router, db, log, authMiddleware)",
		moduleSnake,
	)
	src, err = insertAfterMarker(src, modulesMarker, cmdCall)
	if err != nil {
		return fmt.Errorf("inject modules: %w", err)
	}

	// 2. Inject model migration after @gengo:migration marker
	modelLine := fmt.Sprintf(
		"\t\t&%s_model.%s{},",
		entitySnake, entityCamel,
	)
	src, err = insertAfterMarker(src, migrationMarker, modelLine)
	if err != nil {
		return fmt.Errorf("inject migration: %w", err)
	}

	// 3. Add imports for new module cmd and model
	cmdImport := fmt.Sprintf(
		`%s_cmd "%s/src/module/%s"`,
		moduleSnake, modPath, moduleSnake,
	)
	modelImport := fmt.Sprintf(
		`%s_model "%s/src/module/%s/model"`,
		entitySnake, modPath, moduleSnake,
	)
	src = addImports(src, cmdImport, modelImport)

	return os.WriteFile(mainGoPath, []byte(src), 0644)
}

// insertAfterMarker inserts a new line immediately after the line containing the marker.
func insertAfterMarker(src, marker, newLine string) (string, error) {
	lines := strings.Split(src, "\n")
	for i, line := range lines {
		if strings.Contains(line, marker) {
			// Insert after this line
			updated := make([]string, 0, len(lines)+1)
			updated = append(updated, lines[:i+1]...)
			updated = append(updated, newLine)
			updated = append(updated, lines[i+1:]...)
			return strings.Join(updated, "\n"), nil
		}
	}
	return src, fmt.Errorf("marker %q not found in file", marker)
}

// addImports adds import lines into the existing import block.
// Finds the last import line and appends before the closing paren.
func addImports(src string, imports ...string) string {
	// Find the import block closing paren
	lines := strings.Split(src, "\n")
	importBlockEnd := -1
	inImport := false

	for i, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "import (" {
			inImport = true
			continue
		}
		if inImport && trimmed == ")" {
			importBlockEnd = i
			inImport = false
			break
		}
	}

	if importBlockEnd == -1 {
		return src
	}

	// Build import lines with tab indent, skip already present ones
	existingBlock := strings.Join(lines[:importBlockEnd], "\n")
	var newImports []string
	for _, imp := range imports {
		if !strings.Contains(existingBlock, imp) {
			newImports = append(newImports, "\t"+imp)
		}
	}

	if len(newImports) == 0 {
		return src
	}

	updated := make([]string, 0, len(lines)+len(newImports))
	updated = append(updated, lines[:importBlockEnd]...)
	updated = append(updated, newImports...)
	updated = append(updated, lines[importBlockEnd:]...)
	return strings.Join(updated, "\n")
}
