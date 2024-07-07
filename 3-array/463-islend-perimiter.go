package __array

import "strconv"

func islandPerimeter(grid [][]int) int {
	hashSet := make(map[string]struct{})
	perimeter := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				perimeter += dfs(i, j, &grid, &hashSet)
			}
		}
	}

	return perimeter
}

func dfs(i int, j int, grid *[][]int, hashSet *map[string]struct{}) int {
	coordinates := strconv.Itoa(i) + "," + strconv.Itoa(j)

	if i >= len(*grid) || j >= len((*grid)[0]) || i < 0 || j < 0 || (*grid)[i][j] == 0 {
		return 1
	}
	if _, ok := (*hashSet)[coordinates]; ok {
		return 0
	}

	(*hashSet)[coordinates] = struct{}{}

	perimeter := 0
	perimeter += dfs(i, j+1, grid, hashSet) // right
	perimeter += dfs(i+1, j, grid, hashSet) // down
	perimeter += dfs(i, j-1, grid, hashSet) // left
	perimeter += dfs(i-1, j, grid, hashSet) // up

	return perimeter
}
