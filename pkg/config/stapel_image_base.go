package config

import (
	"fmt"
)

type StapelImageBase struct {
	Name                  string
	From                  string
	FromLatest            bool
	FromImageName         string
	FromImageArtifactName string
	FromCacheVersion      string
	Git                   *GitManager
	Shell                 *Shell
	Ansible               *Ansible
	Mount                 []*Mount
	Import                []*Import

	raw *rawStapelImage
}

func (c *StapelImageBase) GetName() string {
	return c.Name
}

func (c *StapelImageBase) imports() []*Import {
	return c.Import
}

func (c *StapelImageBase) ImageBaseConfig() *StapelImageBase {
	return c
}

func (c *StapelImageBase) IsArtifact() bool {
	return false
}

func (c *StapelImageBase) exportsAutoExcluding() error {
	for _, exp1 := range c.exports() {
		for _, exp2 := range c.exports() {
			if exp1 == exp2 {
				continue
			}

			if !exp1.AutoExcludeExportAndCheck(exp2) {
				errMsg := fmt.Sprintf("Conflict between imports!\n\n%s\n%s", dumpConfigSection(exp1.GetRaw()), dumpConfigSection(exp2.GetRaw()))
				return newDetailedConfigError(errMsg, nil, c.raw.doc)
			}
		}
	}

	return nil
}

func (c *StapelImageBase) exports() []autoExcludeExport {
	var exports []autoExcludeExport
	if c.Git != nil {
		for _, git := range c.Git.Local {
			exports = append(exports, git)
		}

		for _, git := range c.Git.Remote {
			exports = append(exports, git)
		}
	}

	for _, imp := range c.Import {
		exports = append(exports, imp)
	}

	return exports
}

func (c *StapelImageBase) validate() error {
	if c.From == "" && c.raw.FromImage == "" && c.raw.FromImageArtifact == "" && c.FromImageName == "" && c.FromImageArtifactName == "" {
		return newDetailedConfigError("`from: DOCKER_IMAGE`, `fromImage: IMAGE_NAME`, `fromImageArtifact: IMAGE_ARTIFACT_NAME` required!", nil, c.raw.doc)
	}

	mountByTo := map[string]bool{}
	for _, mount := range c.Mount {
		_, exist := mountByTo[mount.To]
		if exist {
			return newDetailedConfigError("conflict between mounts!", nil, c.raw.doc)
		}

		mountByTo[mount.To] = true
	}

	if !oneOrNone([]bool{c.From != "", c.raw.FromImage != "", c.raw.FromImageArtifact != ""}) {
		return newDetailedConfigError("conflict between `from`, `fromImage` and `fromImageArtifact` directives!", nil, c.raw.doc)
	}

	// TODO: валидацию формата `From`
	// TODO: валидация формата `Name`

	return nil
}
