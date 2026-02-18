package main

import (
	"os"
	"text/template"
)

type ContainerConfig struct {
	NODE_IMAGE       string
	REPO_URL         string
	COMMIT_SHA       string
	BUILD_DIR        string
	OFFICIAL_TGZ_URL string
	TGZ_FILENAME     string
	ALPINE_IMAGE     string
}

func main() {
	bytes, error := os.ReadFile("cmd/rep-build/templates/PNPM_Containerfile")
	if error != nil {
		println("Error reading template file:", error.Error())
		return
	}

	tmpl := string(bytes)

	data := ContainerConfig{
		NODE_IMAGE:       "docker.io/library/node:23-alpine",
		ALPINE_IMAGE:     "docker.io/library/alpine:3.23",
		REPO_URL:         "https://github.com/sveltejs/svelte.git",
		COMMIT_SHA:       "e1427aa0c8b5abe88aea972ea397ea499cb8f2db",
		BUILD_DIR:        "packages/svelte",
		OFFICIAL_TGZ_URL: "https://registry.npmjs.org/svelte/-/svelte-5.48.2.tgz",
		TGZ_FILENAME:     "svelte-5.48.2.tgz",
	}

	t := template.Must(template.New("Containerfile").Parse(tmpl))

	containerfile, err := os.Create("cmd/rep-build/npmjs/svelte/5.48.2/Containerfile")
	if err != nil {
		println("Error creating Containerfile:", err.Error())
		return
	}

	defer func() {
		if cerr := containerfile.Close(); cerr != nil {
			println("Error closing Containerfile:", cerr.Error())
		}
	}()

	if err := t.Execute(containerfile, data); err != nil {
		println("Error executing template:", err.Error())
		return
	}
}
