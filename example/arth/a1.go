//求字符串中出现字符最多的两个字符
package main

import (
	"fmt"
)

func main() {
	//d := "list box choose x"
	d := "我是中国人的，从来都是中"
	s := []rune(d)
	zl := len(s)

	var tt []rune
	var ti []int

	for i := 0; i < zl; i++ {
		for j := i + 1; j < zl; j++ {
			if s[i] < s[j] {
				s[i], s[j] = s[j], s[i]
			}
		}

	}

	w := rune(-1)
	zi := -1
	x := 0
	for _, v := range s {
		if w != v {
			tt = append(tt, v)
			zi++

			x = 1
			ti = append(ti, x)
		} else {
			x++
			ti[zi] = x
		}
		w = v
	}

	zl = len(ti)
	for i := 0; i < zl; i++ {
		for j := i + 1; j < zl; j++ {
			if ti[i] < ti[j] {
				ti[i], ti[j] = ti[j], ti[i]
				tt[i], tt[j] = tt[j], tt[i]

			}
		}

	}
	fmt.Println(string(tt[0]), ti[0], string(tt[1]), ti[1])
	fmt.Println(string(tt))
	fmt.Println(ti)
	fmt.Println(d)
}
