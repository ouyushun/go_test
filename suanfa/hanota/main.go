package main

import "fmt"

func main() {


	A := []int{1, 2, 3}
	B := []int{}
	C := []int{}
	hanota(A, B, C)
	fmt.Println(A)
}

func hanota(A []int, B []int, C []int){
	move(len(A), &A , &B , &C)
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
}

func move(n int, A , B, C *[]int) {
	if n == 0 {
		return
	}
	if n == 1 { //最简问题处理：只有1个盘子，把盘子从A移动到C
		*C = (append(*C, (*A)[len(*A)-1]))
		*A = (*A)[:len((*A))-1]
	} else { //小一级问题处理。
		//把n个盘子从A通过B移动到C =
		//1. 先把n-1个盘子从A通过C移动到B
		//2. 把A的最后一个盘子移动到C
		//3. 把n-1个盘子从B 通过A 移动到C
		move(n-1, A, C, B)
		move(1, A, B, C)
		move(n-1, B, A, C)
	}
}


func move2(n int, A , B, C []int) ([]int, []int, []int) {
	if n == 0 {
		return nil, nil, nil
	}

	if n == 1 { //最简问题处理：只有1个盘子，把盘子从A移动到C
		C = append(C, A[len(A)-1])
		A = (A)[:len(A)-1]
		return A,  B, C
	} else { //小一级问题处理。
		//把n个盘子从A通过B移动到C =
		//1. 先把n-1个盘子从A通过C移动到B
		//2. 把A的最后一个盘子移动到C
		//3. 把n-1个盘子从B 通过A 移动到C
		A, B, C = move2(n-1, A, C, B)
		A, B, C = move2(1, A, B, C)
		A, B, C = move2(n-1, B, A, C)
		return A, B, C
	}
}



