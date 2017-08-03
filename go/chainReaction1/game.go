package chainReaction1

func chainReaction1(n int, moves [][]int) (gameMap [][]int) {
	gameMap = createEmptyArray(n)
	for _, move := range moves {
		gameMap = addPoint(gameMap, move, n)
	}
	return
}

func addPoint(currentMap [][]int, move []int, n int) [][]int {
	dx := move[0]
	dy := move[1]
	currentMap[dx][dy] = currentMap[dx][dy] + 1
	if currentMap[dx][dy] == getMaxDimenssionInPosition(dx, dy, n) {
		currentMap[dx][dy] = 0
		if dx > 0 {
			currentMap = addPoint(currentMap, []int{dx - 1, dy}, n)
		}
		if dx < n-1 {
			currentMap = addPoint(currentMap, []int{dx + 1, dy}, n)
		}
		if dy > 0 {
			currentMap = addPoint(currentMap, []int{dx, dy - 1}, n)
		}
		if dy < n-1 {
			currentMap = addPoint(currentMap, []int{dx, dy + 1}, n)
		}
	}
	return currentMap
}

func getMaxDimenssionInPosition(dx int, dy int, n int) (dimension int) {
	dimension = 4
	if dx == 0 {
		dimension = dimension - 1
	}
	if dy == 0 {
		dimension = dimension - 1
	}
	if dx == n-1 {
		dimension = dimension - 1
	}
	if dy == n-1 {
		dimension = dimension - 1
	}
	return dimension
}

func createEmptyArray(n int) (arr [][]int) {
	arr = make([][]int, n)
	for i := range arr {
		arr[i] = make([]int, n)
	}
	return
}
