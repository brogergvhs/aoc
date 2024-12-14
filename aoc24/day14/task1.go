package main

func getEndRobotsPos(robots []Robot, seconds, width, height int) []Robot {
	movedRobots := make([]Robot, len(robots))
	copy(movedRobots, robots)

	for i := range movedRobots {
		r := &movedRobots[i]
		r.px = (r.px + r.vx*seconds) % width
		if r.px < 0 {
			r.px += width
		}

		r.py = (r.py + r.vy*seconds) % height
		if r.py < 0 {
			r.py += height
		}
	}

	return movedRobots
}

func calculateSafetyFactor(robots []Robot, width, height int) int {
	q1, q2, q3, q4 := 0, 0, 0, 0

	for _, r := range robots {
		if r.px == width/2 || r.py == height/2 {
			continue
		}

		if r.px < width/2 && r.py < height/2 {
			q1++
		} else if r.px >= width/2 && r.py < height/2 {
			q2++
		} else if r.px < width/2 && r.py >= height/2 {
			q3++
		} else if r.px >= width/2 && r.py >= height/2 {
			q4++
		}
	}

	return q1 * q2 * q3 * q4
}
