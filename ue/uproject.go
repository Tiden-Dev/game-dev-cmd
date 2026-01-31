package ue

import (
	"encoding/json"
	"os"
)

// UProject represents the structure of an Unreal Engine .uproject file
type UProject struct {
	FileVersion        int                `json:"FileVersion"`        // Usually 3
	EngineAssociation  string             `json:"EngineAssociation"`  // "5.4", "5.5", or GUID for source build
	Category           string             `json:"Category,omitempty"` // Optional, e.g. "Samples", "Games"
	Description        string             `json:"Description,omitempty"`
	Modules            []ModuleDescriptor `json:"Modules,omitempty"` // Present in code projects
	TargetPlatforms    []string           `json:"TargetPlatforms,omitempty"`
	Plugins            []PluginDescriptor `json:"Plugins,omitempty"` // Enabled/disabled plugins
	WhitelistPlatforms []string           `json:"WhitelistPlatforms,omitempty"`
	BlacklistPlatforms []string           `json:"BlacklistPlatforms,omitempty"`
}

// ModuleDescriptor matches entries in the "Modules" array
type ModuleDescriptor struct {
	Name                   string   `json:"Name"`                   // e.g. "MyGame", "MyGameEditor"
	Type                   string   `json:"Type"`                   // "Runtime", "Developer", "Editor", "UncookedOnly", etc.
	LoadingPhase           string   `json:"LoadingPhase,omitempty"` // "Default", "PreDefault", "PostEngineInit", etc.
	AdditionalDependencies []string `json:"AdditionalDependencies,omitempty"`
	// Rarely used extra fields
	WhitelistPlatforms []string `json:"WhitelistPlatforms,omitempty"`
	BlacklistPlatforms []string `json:"BlacklistPlatforms,omitempty"`
}

// PluginDescriptor matches entries in the "Plugins" array
type PluginDescriptor struct {
	Name    string `json:"Name"`    // e.g. "OnlineSubsystemSteam", "Nanite", "MovieRenderQueue"
	Enabled bool   `json:"Enabled"` // true/false
	// Optional target filters (rare)
	TargetAllowList []string `json:"TargetAllowList,omitempty"`
	TargetDenyList  []string `json:"TargetDenyList,omitempty"`
}

func ReadUProject(filePath string) (*UProject, error) {
	var project = &UProject{}
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &project); err != nil {
		return nil, err
	}
	return project, nil
}
