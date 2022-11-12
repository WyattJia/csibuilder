/*
 Copyright 2022 CSIBuilder

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License
*/

package csi

import (
	"fmt"
	"path/filepath"

	"csibuilder/pkg/machinery"
)

var _ machinery.Template = &Node{}

// Node scaffolds the file that defines the node server of csi
// nolint:maligned
type Node struct {
	machinery.TemplateMixin
	machinery.ResourceMixin
	machinery.RepositoryMixin
	machinery.BoilerplateMixin

	Force bool
}

func (c *Node) SetTemplateDefaults() error {
	if c.Path == "" {
		c.Path = filepath.Join("pkg/csi", "node.go")
	}
	c.Path = c.Resource.Replacer().Replace(c.Path)
	fmt.Println(c.Path)

	body, err := tplFS.ReadFile("templates/node.go.tpl")
	if err != nil {
		return err
	}
	c.TemplateBody = string(body)

	if c.Force {
		c.IfExistsAction = machinery.OverwriteFile
	} else {
		c.IfExistsAction = machinery.Error
	}
	return nil
}
