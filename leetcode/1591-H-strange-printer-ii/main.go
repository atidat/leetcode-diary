package main

import "fmt"

func main() {

}
func isPrintable(ta [][]int) bool {
	usedColor := make(map[int]bool, 0)


	for i := 0; i < len(ta[0]); i++ {
		for j := 0; j < len(ta); j++ {
			if i == 0 && j == 0 {
				continue
			}
			if j - 1 >= 0 && ta[i][i] == ta[i][j-1] {
				continue
			}
			if i != 0 && j == 0 && ta[i][j] == ta[i-1][j] {
				continue
			}
			// 当前位置[i][j]与[i-1][j] 或者 [i][j-1]不同，且是新颜色
			// 		则进行染色
			//		否则退出
			if _, ok := usedColor[ta[i][j]]; ok {
				return false
			} else {
				if j == 0 {
					ta[i][j] = ta[i-1][j]
				} else {
					ta[i][j] =ta[i][j-1]
				}
				usedColor[ta[i][j]] = true
			}
			fmt.Println()
			fmt.Println(ta)
		}
	}
	return true

}

func nearby(ta [][]int, tarColor int, posX, posY int) {
	navi := [][]int{
		{0, 1},
		//{-1, 0},
		//{-1, 0},
		{0, 1},
	}
	// 自身染色
	if posY == 0 {
		ta[posX][posY] = ta[posX-1][posY]
	} else {
		ta[posX][posY] = ta[posX][posY-1]
	}
	//周边染色
	var tmpColor int
	for i := 0; i < 4; i++ {
		if ta[posX+navi[i][0]][posY+navi[i][1]] == tarColor {
			continue
		} else {
			tmpColor = ta[posX+navi[i][0]][posY+navi[i][1]]
			runX, runY := navi[i][0], navi[i][1]
			for {
				if ta[posX+runX][posY+runY] != tmpColor {
					break
				}

			}
		}
	}
}