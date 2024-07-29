package memory

import (
	"chat/repository/cache/memory"
)

func ChatroomLocalMemoryInit() {
	memory.NewRoomMap()
}
