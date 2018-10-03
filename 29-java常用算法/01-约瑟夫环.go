package _9_java常用算法

import (
	"errors"
	"fmt"
)

// 41 人 2 人存活，问存活人的原始编号

//解题思路
// 数组容量41初始化0 ，轮到3的标记当前自杀序号，知道循环结束

// alive 存活的人数
const numAll = 41
const killInter = 3
func yosephus(alive uint8) error {

	if alive <= 0 {
		return errors.New("存活人没意义")
	}

	people := [numAll]uint8{}
	var count uint8 = 0
	var pos = -1
	var i = 0
	// 第二步 标记要自杀的人，并打印
	for count < numAll {

		// 第一步 找到要自杀的人
		for {
			pos = (pos + 1) % numAll
			if people[pos] == 0 {
				i++
			}
			if i == killInter {
				i =0
				break
			}
		}
		count++
		people[pos] = count
		fmt.Printf("第 %2d 个人自杀，约瑟夫环编号为 %2d \t", count, pos+1)
		if count % 3 == 0{
			fmt.Println("\n")
		}
	}

	fmt.Println("\n")
	// 第三步 输出结果
	alive = numAll - alive
	for i := 0; i < numAll; i++ {

		if people[i] > alive {
			fmt.Printf(" 初识编号为 %2d, 约瑟夫环编号 : %d \n", i+1, people[i])
		}
	}

	return nil
}



