package ue

import "testing"

const (
	testDefaultGameFilePath = "testdata/57/DefaultGame.ini"
)

func TestDefaultGame_GetProjectID(t *testing.T) {
	const expectedProjectId = "B66FCFC2EF4EBFD6288AB7B6C208B882"
	dg, err := LoadDefaultGame(testDefaultGameFilePath, V57)
	if err != nil {
		t.Fatal(err)
	}
	projectId, found := dg.GetProjectID()
	if !found {
		t.Fatal("projectId not found")
	}

	if projectId != expectedProjectId {
		t.Errorf("projectId is %s, expected %s", projectId, expectedProjectId)
	}
}

func TestDefaultGame_GetProjectName(t *testing.T) {
	const expectedProjectName = "Third Person BP Game Template"
	dg, err := LoadDefaultGame(testDefaultGameFilePath, V57)
	if err != nil {
		t.Fatal(err)
	}
	projectName, found := dg.GetProjectName()
	if !found {
		t.Fatal("projectName not found")
	}

	if projectName != expectedProjectName {
		t.Errorf("projectName is %s, expected %s", projectName, expectedProjectName)
	}
}
