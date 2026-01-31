package ue

import "gopkg.in/ini.v1"

const DefaultGameIniFileName = "DefaultGame.ini"

const (
	GeneralProjectSettings IniSection = "/Script/EngineSettings.GeneralProjectSettings"
)

const (
	ProjectId   IniKey = "ProjectID"
	ProjectName IniKey = "ProjectName"
)

type DefaultGame struct {
	FilePath      string
	IniFile       *ini.File
	EngineVersion EngineVersion
}

func LoadDefaultGame(filePath string, engineVersion EngineVersion) (*DefaultGame, error) {
	var (
		dg  DefaultGame
		err error
	)
	dg.IniFile, err = ini.Load(filePath)
	if err != nil {
		return nil, err
	}
	dg.EngineVersion = engineVersion
	return &dg, nil
}

func (dg *DefaultGame) GetProjectID() (string, bool) {
	var value string
	section, err := dg.IniFile.GetSection(GeneralProjectSettings.String())
	if err != nil {
		return value, false
	}
	key := section.Key(ProjectId.String())
	if key == nil {
		return value, false
	}
	return key.Value(), true
}

func (dg *DefaultGame) GetProjectName() (string, bool) {
	var value string
	section, err := dg.IniFile.GetSection(GeneralProjectSettings.String())
	if err != nil {
		return value, false
	}
	key := section.Key(ProjectName.String())
	if key == nil {
		return value, false
	}
	return key.Value(), true
}
