package main

import (
	m "kata/model"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Rover moving", func() {
	It("move forward one position", func() {
		robot := m.Robot{PositionX: 0, PositionY: 0, Direction: m.North}
		robotPosition := robot.Recieve("f")
		Expect(robotPosition).To(Equal(m.Robot{PositionX: 0, PositionY: 1, Direction: m.North}))
	})

	It("move backward one position", func() {
		robot := m.Robot{PositionX: 0, PositionY: 0, Direction: m.North}
		robotPosition := robot.Recieve("b")
		Expect(robotPosition).To(Equal(m.Robot{PositionX: 0, PositionY: -1, Direction: m.North}))
	})

	It("move left one position, change its direction", func() {
		robot := m.Robot{PositionX: 0, PositionY: 0, Direction: m.North}
		robotPosition := robot.Recieve("l")
		Expect(robotPosition).To(Equal(m.Robot{PositionX: -1, PositionY: 0, Direction: m.West}))
	})

	It("move right one position, change its direction", func() {
		robot := m.Robot{PositionX: 0, PositionY: 0, Direction: m.North}
		robotPosition := robot.Recieve("r")
		Expect(robotPosition).To(Equal(m.Robot{PositionX: 1, PositionY: 0, Direction: m.East}))
	})

	It("move forward two position", func() {
		robot := m.Robot{PositionX: 0, PositionY: 0, Direction: m.North}
		robotPosition := robot.Recieve("ff")
		Expect(robotPosition).To(Equal(m.Robot{PositionX: 0, PositionY: 2, Direction: m.North}))
	})
	/* 	It("passing", func() {
		Expect(42).To(Equal(42))
	}) */
})
