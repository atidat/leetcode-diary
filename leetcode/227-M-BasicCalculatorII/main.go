package main

import "fmt"
/*
func deleteSpace(s string) string {
	b := make([]byte, 0)
	for i := 0; i <len(s); i++ {
		if s[i] != ' ' {
			b = append(b, s[i])
		}
	}
	return string(b)
}

func readExpr(s string, ind int) (int, int) {
	lenS := len(s)
	if ind > lenS - 1 {
		return -1, -1
	}

	val := 0
	for ; ind < lenS && ((s[ind] >= '0' && s[ind] <= '9') || s[ind] == ' '); ind++ {
		if s[ind] == ' ' {
			continue
		}
		val *= 10
		val += int(s[ind] - '0')
	}

	if ind == lenS {
		return val, -1
	}

	for ind < lenS {
		if s[ind] == ' ' {
			ind++
		} else {
			break
		}
	}

	if ind == lenS {
		ind = -1
	}
	// fmt.Println(val, ind, s)
	return val, ind
}

func except(s string) int {
	i, val := 0, 0
	for ; i < len(s); i++ {
		if !(s[i] >= '0' && s[i] <= '9') {
			break
		}
		val *= 10
		val += int(s[i] - '0')
	}
	if i == len(s) {
		return val
	}
	return -1
}
*/
/*
func calculate(str string) int {

// 基于栈实现
	// 把表达式的第一个数字先缓存，再把运算符塞到栈，从缓存中塞回栈
	// 再读取表达式，把数字放缓存，再把运算符塞到栈，再从缓存中塞回栈
	// 反复以往
	// 读取最后一个数字后，不再有运算符，停止读表达式，
	// 开始每两个一组，取栈顶记录

	// 栈：通过切片去实现，不自己实现栈结构
	// str := deleteSpace(s)
	// tmp := except(str)
	// if tmp != -1 {
	//	fmt.Printf("result is %d\n", tmp)
	//	return tmp
	// }

	numRepo := make([]int, 0)
	opeRepo := make([]byte, 0)
	index := 0
	for {
		num, operator := readExpr(str, index)
		if num == -1 {
			break
		}
		if operator == -1 {
			numRepo = append(numRepo, num)
			break
		}

		if len(opeRepo) > 0 &&
			(opeRepo[len(opeRepo) - 1] == '/' || opeRepo[len(opeRepo) - 1] == '*') {
			numRepo[len(numRepo) - 1] = calc(numRepo[len(numRepo) - 1], num, opeRepo[len(opeRepo) - 1])
			opeRepo = opeRepo[:len(opeRepo) - 1]
		} else {
			numRepo = append(numRepo, num)
		}

		opeRepo = append(opeRepo, str[operator])
		index = operator + 1
	}

	if len(opeRepo) == 0 {
		fmt.Printf("result is %d\n", numRepo[0])
		return numRepo[0]
	}
	if opeRepo[len(opeRepo) - 1] == '/' || opeRepo[len(opeRepo) - 1] == '*' {
		numRepo[len(numRepo) - 2] = calc(numRepo[len(numRepo) - 2], numRepo[len(numRepo) - 1],
			opeRepo[len(opeRepo) - 1])
		numRepo = numRepo[0:len(numRepo) - 1]
		opeRepo = opeRepo[:len(opeRepo) - 1]
	}

	return finalCalc(numRepo, opeRepo)
}

func finalCalc(numRepo []int, opeRepo []byte) int {
	// fmt.Printf("numRepo is %v\n", numRepo)
	// fmt.Printf("opeRepo is %v\n", opeRepo)
	for len(opeRepo) > 0 {
		numRepo[0] = calc(numRepo[0], numRepo[1], opeRepo[0])
		numRepo = append(numRepo[0:1], numRepo[2:]...)
		opeRepo = opeRepo[1:]
	}
	fmt.Printf("result is %d\n", numRepo[0])
	return numRepo[0]
}

func calc(a, b int, oper byte) int {
	switch oper {
	case '+': return a + b
	case '-': return a - b
	case '*': return a * b
	case '/': return a / b
	}
	return -1
}
*/

func calculate(str string) int {
	operMap := map[rune]bool {
		'+': true,
		'-': true,
		'*': true,
		'/': true,
	}

	preOp := '+'
	num := 0
	for k, v := range str {
		if isNum(v) {
			num = num * 10 + int(v - '0')
		}

		if operMap[v] {
			switch preOp {
			case '+':
				preNum += num
			case '-':
				preNum -= num
			case '*':
				preNum *= num
			case '/':
				preNum /= num
			}
			preOp = v
			num = 0
		}

	}
}

func isNum(b rune) bool {
	return b >= '0' && b <= '9'
}

func main() {
	_ = calculate("3+2*2")
	_ = calculate(" 3/2 ")
	_ = calculate(" 3+5 / 2")
	_ = calculate("1-1+1")
	_ = calculate("0")
	_ = calculate("123")
}

















