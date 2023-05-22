// Copyright 2022 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build plugins
// +build plugins

package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

//go:generate go run generate.go

func main() {
	data, err := os.ReadFile(filepath.Join("..", "plugins.yml"))
	if err != nil {
		log.Fatal(err)
	}

	var plugins []string
	err = yaml.Unmarshal(data, &plugins)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create("plugins.go")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	_, err = f.WriteString(`// Copyright 2022 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This file is generated by "make plugins".

package plugins

`)
	if err != nil {
		log.Fatal(err)
	}

	if len(plugins) == 0 {
		return
	}

	_, err = f.WriteString("import (\n")
	if err != nil {
		log.Fatal(err)
	}

	for i, plugin := range plugins {
		_, err = f.WriteString(fmt.Sprintf("\t// Register %s plugin.\n", path.Base(plugin)))
		if err != nil {
			log.Fatal(err)
		}
		_, err = f.WriteString(fmt.Sprintf("\t_ \"%s\"\n", plugin))
		if err != nil {
			log.Fatal(err)
		}
		if i < len(plugins)-1 {
			_, err = f.WriteString("\n")
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	_, err = f.WriteString(")\n")
	if err != nil {
		log.Fatal(err)
	}
}
