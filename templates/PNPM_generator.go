package templates

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

func Generate(nodeImage string, alpineImage string, repoURL string, commitSHA string, buildDir string, officialTgzURL string, tgzFilename string, outputPath string) {
	bytes, error := os.ReadFile("templates/PNPM_Containerfile")
	if error != nil {
		println("Error reading template file:", error.Error())
		return
	}

	tmpl := string(bytes)

	data := ContainerConfig{
		NODE_IMAGE:       nodeImage,
		ALPINE_IMAGE:     alpineImage,
		REPO_URL:         repoURL,
		COMMIT_SHA:       commitSHA,
		BUILD_DIR:        buildDir,
		OFFICIAL_TGZ_URL: officialTgzURL,
		TGZ_FILENAME:     tgzFilename,
	}

	t := template.Must(template.New("Containerfile").Parse(tmpl))

	containerfile, err := os.Create(outputPath)
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
