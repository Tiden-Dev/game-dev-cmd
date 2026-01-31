package ue

type EngineVersion string

const (
	V50 EngineVersion = "5.0"
	V51 EngineVersion = "5.1"
	V52 EngineVersion = "5.2"
	V53 EngineVersion = "5.3"
	V54 EngineVersion = "5.4"
	V55 EngineVersion = "5.5"
	V56 EngineVersion = "5.6"
	V57 EngineVersion = "5.7"
)

func (ev EngineVersion) String() string {
	return string(ev)
}
