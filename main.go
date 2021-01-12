package main

import "fmt"

var tmpAns [][][]int = make([][][]int, 0)
var q int

func main() {
	q = 8

	for i := 0; i < q; i++ {
		var m = make([][]int, q)
		for i := 0; i < q; i++ {
			m[i] = make([]int, q)
		}
		//輪流使用第一排
		tryAddQueen(m, i, 0)
	}

	var ans = make([][]string, 0)
	// there are len(tmpAns) answers
	for i := 0; i < len(tmpAns); i++ {
		//each slices([][]int) is a answer
		var singleAns = make([]string, 0)
		for j := 0; j < q; j++ {
			//every slice([]int) means a queen
			var queen = ""
			for k := 0; k < q; k++ {
				var v = tmpAns[i][j][k]
				if v == 0 || v == -1 {
					queen += "."
				} else {
					queen += "Q"
				}
			}
			//fmt.Println(queen)
			singleAns = append(singleAns, queen)
		}
		ans = append(ans, singleAns)
	}
	fmt.Println(len(ans))
}

func tryAddQueen(m [][]int, x, y int) {
	if m[y][x] == -1 {
		return
	}
	m[y][x] = 1
	if y == q-1 {
		tmpAns = append(tmpAns, m)
		return
	}
	//遮蔽下面的位置
	for c := 1; y+c < q; c++ {
		m[y+c][x] = -1
		if x+c < q {
			m[y+c][x+c] = -1
		}
		if x-c >= 0 {
			m[y+c][x-c] = -1
		}
	}
	//測試下一層
	for i := 0; i < q; i++ {
		if y+1 >= q {
			return
		}

		if m[y+1][i] != -1 {
			var nm = deepCopyMatrix(m)
			tryAddQueen(nm, i, y+1)
		}
	}
}

func deepCopyMatrix(m [][]int) [][]int {
	var nm = make([][]int, q)
	for i := 0; i < q; i++ {
		var ns = make([]int, q)
		copy(ns, m[i])
		nm[i] = ns
	}

	return nm
}
