package _9_java常用算法

import (
	"testing"
	"fmt"
)

func TestYosephus(t *testing.T)  {
	//yosephus(2)

	link := &YosefuLink{}
	link.addTailNode(3)
	link.addTailNode(5)
	link.addTailNode(6)
	link.addTailNode(2)
	link.addTailNode(1)
	link.addTailNode(5)
	link.addTailNode(7)
	link.addTailNode(2)

	fmt.Println(link)
	p := link.head
	for {
		fmt.Println(p)
		p = p.next
		if p == link.head{
			break
		}
	}

	err := link.getLastOne(3)
	if err != nil {
		t.Error(err)
	}

}

