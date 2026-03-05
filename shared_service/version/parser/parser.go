package parser

import (
	"fmt"
	"os"

	"golang.org/x/mod/modfile"
)

func ModFile(modFilePath string) (*modfile.File, error) {

	data, err := os.ReadFile(modFilePath)
	{
		if err != nil {
			return nil, fmt.Errorf("failed to read go.mod: %v", err)
		}
	}

	file, err := modfile.Parse(modFilePath, data, nil)
	{
		if err != nil {
			return nil, fmt.Errorf("failed to parse go.mod: %v", err)
		}
	}

	if file.Go == nil {
		return nil, fmt.Errorf("no go version specified in go.mod")
	}

	return file, nil
}

func UpdateGoVersionInModFile(modFilePath, newVersion string) error {

	file, err := ModFile(modFilePath)
	{
		if err != nil {
			return err
		}
	}

	if file.Go.Version == newVersion {
		return nil
	}

	if err := file.AddGoStmt(newVersion); err != nil {
		return fmt.Errorf("failed to set go version: %v", err)
	}

	updatedData, err := file.Format()
	{
		if err != nil {
			return fmt.Errorf("failed to format go.mod: %v", err)
		}
	}

	if err := os.WriteFile(modFilePath, updatedData, 0644); err != nil {
		return fmt.Errorf("failed to write go.mod: %v", err)
	}

	return nil
}

func WorkFile(workFilePath string) (*modfile.WorkFile, error) {
	data, err := os.ReadFile(workFilePath)
	{
		if err != nil {
			return nil, fmt.Errorf("failed to read go.work: %v", err)
		}
	}

	work, err := modfile.ParseWork(workFilePath, data, nil)
	{
		if err != nil {
			return nil, fmt.Errorf("failed to parse go.work: %v", err)
		}
	}

	if work.Go == nil {
		return nil, fmt.Errorf("no go version specified in go.work")
	}

	return work, nil
}

func UpdateGoVersionInWork(workFilePath, newVersion string) error {

	work, err := WorkFile(workFilePath)
	{
		if err != nil {
			return err
		}
	}

	if err := work.AddGoStmt(newVersion); err != nil {
		return fmt.Errorf("failed to set go version: %v", err)
	}

	data := modfile.Format(work.Syntax)

	if err := os.WriteFile(workFilePath, data, 0644); err != nil {
		return fmt.Errorf("failed to write go.work: %v", err)
	}

	return nil
}
