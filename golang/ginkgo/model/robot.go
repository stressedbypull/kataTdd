package model

const (
	N = "NORD"
	E = "EAST"
	S = "SOUTH"
	W = "WEST"
)

type Direction string

const (
	North Direction = "N"
	East  Direction = "E"
	South Direction = "S"
	West  Direction = "W"
)

type Robot struct {
	PositionX int
	PositionY int
	Direction Direction
}

type RobotInterface interface {
	Recieve(command string)
}

func (r *Robot) Recieve(command string) Robot {
	robotToreturn := r
	characters := []rune(command)
	for i := 0; i < len(characters); i++ {
		char := string(characters[i])
		switch char {
		case "f":
			robotToreturn.mooveRobot(1)
		case "b":
			robotToreturn.mooveRobot(-1)
		case "l":
			robotToreturn.changeDirection("l")
			robotToreturn.mooveRobot(1)
		case "r":
			robotToreturn.changeDirection("r")
			robotToreturn.mooveRobot(1)
		}
	}
	return *robotToreturn
}

func (r *Robot) changeDirection(command string) {
	if command == "l" {
		switch r.Direction {
		case North:
			r.Direction = West
		case South:
			r.Direction = East
		case East:
			r.Direction = North
		case West:
			r.Direction = South
		}
	} else if command == "r" {
		switch r.Direction {
		case North:
			r.Direction = East
		case South:
			r.Direction = West
		case East:
			r.Direction = South
		case West:
			r.Direction = North
		}
	}
}

/* func (r *Robot) mooveRobot(command string) {
	switch r.Direction {
	case North:
		if command == "b" {
			r.PositionY--
		} else {
			r.PositionY++
		}
	case South:
		if command == "b" {
			r.PositionY++
		} else {
			r.PositionY--
		}
	case East:
		if command == "b" {
			r.PositionX--
		} else {
			r.PositionX++
		}
	case West:
		if command == "b" {
			r.PositionX++
		} else {
			r.PositionX--
		}
	}
} */

func (r *Robot) mooveRobot(distance int) {
	switch r.Direction {
	case North:
		r.PositionY = r.PositionY + distance
	case South:
		r.PositionY = r.PositionY - distance
	case East:
		r.PositionX = r.PositionX + distance
	case West:
		r.PositionX = r.PositionX - distance
	}
}
