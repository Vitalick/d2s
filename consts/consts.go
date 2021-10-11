package consts

import "encoding/binary"

//ActId type for quests.Act and waypoints.Act and for indexing Act in quests.Difficulty and waypoints.Difficulty
type ActId byte

var (
	BinaryEndian = binary.LittleEndian
)

const (
	Act1 ActId = iota
	Act2
	Act3
	Act4
	Act5
)

const ActsCount = 5
