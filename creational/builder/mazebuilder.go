package main

type MazeBuilderInterface interface {
	buildMaze()
	buildRoom(int)
	buildDoor(int, int)
	getMaze() *Maze
}

type StandardMazeBuilder struct {
	currentMaze *Maze
}

func (mazeBuilder *StandardMazeBuilder) buildMaze(){
	mazeBuilder.currentMaze = new(Maze)
	mazeBuilder.currentMaze.rooms = make(map[int]Spacer)
}

func (mazeBuilder *StandardMazeBuilder) buildRoom(roomNo int){
	_, ok := mazeBuilder.currentMaze.roomNo(roomNo)
	if(!ok){
		newRoom := Room{
			roomNo: roomNo,
		}
		mazeBuilder.currentMaze.addRoom(&newRoom)

		newRoom.setSide(north, new(Wall))
		newRoom.setSide(south, new(Wall))
		newRoom.setSide(west, new(Wall))
		newRoom.setSide(east, new(Wall))
	}
}

func (mazeBuilder *StandardMazeBuilder) buildDoor(roomFrom int, roomTo int){
	r1, _ := mazeBuilder.currentMaze.roomNo(roomFrom)
	r2, _ := mazeBuilder.currentMaze.roomNo(roomTo)

	room1 := r1.(*Room);
	room2 := r2.(*Room);

	door := &Door{
		room1: room1,
		room2: room2,
	}
	
	room1.setSide(mazeBuilder.commonWall(room1, room2), door)
	room1.setSide(mazeBuilder.commonWall(room2, room1), door)
}

func (mazeBuilder *StandardMazeBuilder) getMaze() *Maze {
	return mazeBuilder.currentMaze
}

func (mazebuilder *StandardMazeBuilder) commonWall(room1 Spacer, room2 Spacer) int {
	if(room1.getRoomNo() > room2.getRoomNo()){
		return east
	} else{
		return west
	}
}

type CountingMazeBuilder struct {
	doors int
	rooms int
	currentMaze *Maze
}

func (mazeBuilder *CountingMazeBuilder) buildMaze(){
	mazeBuilder.currentMaze = new(Maze)
	mazeBuilder.currentMaze.rooms = make(map[int]Spacer)
}

func (mazeBuilder *CountingMazeBuilder) buildRoom(roomNo int){
	mazeBuilder.rooms += 1
}

func (mazeBuilder *CountingMazeBuilder) buildDoor(roomFrom int, roomTo int){
	mazeBuilder.doors += 1
}

func (mazeBuilder *CountingMazeBuilder) getMaze() *Maze {
	return mazeBuilder.currentMaze
}

func (mazeBuilder *CountingMazeBuilder) getCounts() (int, int) {
	return mazeBuilder.rooms, mazeBuilder.doors
}

