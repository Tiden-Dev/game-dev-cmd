package ue

import "testing"

func TestReadUProject(t *testing.T) {
	const (
		fp                     = "testdata/57/Default.uproject"
		expectedEnabledPlugins = 2
	)
	project, err := ReadUProject(fp)
	if err != nil {
		t.Fatal(err)
	}

	if project.EngineAssociation != V57.String() {
		t.Error("engine association is not V57")
	}

	var enabledPlugins int
	for _, plugin := range project.Plugins {
		if plugin.Enabled {
			enabledPlugins++
		}
	}

	if enabledPlugins != expectedEnabledPlugins {
		t.Error("expected enabled plugins")
	}
}
