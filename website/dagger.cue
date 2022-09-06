package main

import (
	"dagger.io/dagger"
	"universe.dagger.io/bash"
	"universe.dagger.io/docker"
  "universe.dagger.io/netlify"
)

dagger.#Plan & {
  client: {
    filesystem: ".": read: contents: dagger.#FS
    env: {
      NETLIFY_TOKEN: dagger.#Secret
      REACT_APP_API_URL: string
    }
  }
	actions: {
    pull: docker.#Pull & {
      source: "node:lts"
    }
    copy: docker.#Copy & {
      input:    pull.output
      contents: client.filesystem.".".read.contents
    }
    install: bash.#Run & {
      input: copy.output
      script: contents: """
        npm install
        """
    }
    build: bash.#Run & {
      input: install.output
      script: contents: """
        REACT_APP_API_URL=\(client.env.REACT_APP_API_URL) npm run build
        """
      export: directories: "/build": dagger.#FS
    }

    deploy: netlify.#Deploy & {
			contents: build.export.directories."/build"
			site:     string | *"kpenfound-desk"
      token: client.env.NETLIFY_TOKEN
		}
	}
}
