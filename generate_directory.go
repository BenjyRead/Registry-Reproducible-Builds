package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Simple CLI tool to generate a .gitkeep file for a given npm package and version.
// Usage: `go run generate_directory.go -dep "@sveltejs/adapter-auto:^7.0.0"`
// Or: `go run generate_directory.go -name "@sveltejs/adapter-auto" -version "^7.0.0"`
// NOTE: This is GPT coded, please verify and test
func main() {
	dep := flag.String("dep", "", `Dependency string e.g. "@sveltejs/adapter-auto:^7.0.0"`)
	name := flag.String("name", "", "Package name e.g. @sveltejs/adapter-auto")
	version := flag.String("version", "", "Version e.g. ^7.0.0")
	root := flag.String("root", "npmjs", "Root directory (default: npmjs)")
	flag.Parse()

	var pkgName, pkgVersion string

	if *dep != "" {
		parts := strings.Split(*dep, ":")
		if len(parts) != 2 {
			exitErr(errors.New("invalid -dep format, expected name:version"))
		}
		pkgName = strings.TrimSpace(parts[0])
		pkgVersion = strings.TrimSpace(parts[1])
	} else {
		if *name == "" || *version == "" {
			exitErr(errors.New("provide either -dep or both -name and -version"))
		}
		pkgName = *name
		pkgVersion = *version
	}

	pkgVersion = cleanVersion(pkgVersion)

	scope, pkg := splitPackage(pkgName)

	var path string
	if scope != "" {
		path = filepath.Join(*root, scope, pkg, pkgVersion)
	} else {
		path = filepath.Join(*root, pkg, pkgVersion)
	}

	err := os.MkdirAll(path, 0o755)
	if err != nil {
		exitErr(err)
	}

	gitkeep := filepath.Join(path, ".gitkeep")

	if _, err := os.Stat(gitkeep); os.IsNotExist(err) {
		f, err := os.Create(gitkeep)
		if err != nil {
			exitErr(err)
		}
		if err := f.Close(); err != nil {
			exitErr(err)
		}
		fmt.Println("Created:", gitkeep)
	} else {
		fmt.Println("Gitkeep already exists:", gitkeep)
	}
}

func splitPackage(name string) (scope, pkg string) {
	name = strings.TrimSpace(name)

	if strings.HasPrefix(name, "@") {
		name = strings.TrimPrefix(name, "@")
		parts := strings.Split(name, "/")
		if len(parts) != 2 {
			exitErr(errors.New("invalid scoped package format"))
		}
		return parts[0], parts[1]
	}

	return "", name
}

func cleanVersion(v string) string {
	v = strings.TrimSpace(v)
	v = strings.TrimLeft(v, "^~>=<")
	return v
}

func exitErr(err error) {
	fmt.Fprintln(os.Stderr, "Error:", err)
	os.Exit(1)
}
