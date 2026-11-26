package grind

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	// save og color
	// if og color == color
	// return img

	// create queue
	// set starting point as first el in queue
	// while len queue > 0
	// if out of bounds or current color (queue[0]
	// else change color

	startingColor := image[sr][sc]
	if startingColor == color {
		return image
	}

	queue := [][2]int{{sr, sc}}

	for len(queue) > 0 {
		cell := queue[0]
		queue = queue[1:]

		row := cell[0]
		coll := cell[1]

		// If startingColor != currentColor we stop, we modify only if color is the same as the startingColor
		// NOTE: I was going to put image[row][coll] in a var, but ran into out of bounds. Validate index before using it to access a slice!
		if row < 0 || row >= len(image) || coll < 0 || coll >= len(image[0]) || startingColor != image[row][coll] {
			continue
		}

		image[row][coll] = color

		queue = append(
			queue,
			[2]int{row + 1, coll},
			[2]int{row - 1, coll},
			[2]int{row, coll + 1},
			[2]int{row, coll - 1},
		)

		continue
	}

	return image
}
