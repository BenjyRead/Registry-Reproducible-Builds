package main

import (
	"reproducible-builds/templates"
)

func main() {
	templates.Generate(
		"node:23-alpine",
		"alpine:3.23",
		"https://github.com/sveltejs/svelte.git",
		"e1427aa0c8b5abe88aea972ea397ea499cb8f2db",
		"packages/svelte",
		"https://registry.npmjs.org/svelte/-/svelte-5.48.2.tgz",
		"svelte-5.48.2.tgz",
		"npmjs/svelte/5.48.2/Containerfile",
	)
}
