package interview

import (
	"fmt"
	"testing"
	"time"
	"sync"
	. "github.com/smartystreets/goconvey/convey"
)

func TestStack(t *testing.T) {
	s := NewStack()

	//101个Goroutine 对其进行push
	//101 个Groutine 对其进行pop

	var wg = sync.WaitGroup{}
	var count = 101

	for i:=0 ;i < count ;i++{
		wg.Add(1)
		go func(value int) {
			fmt.Println(s.push(value))
			wg.Done()
		}(i)

	}
	wg.Wait()
	fmt.Println(*s)
	for i:=0 ;i < count ;i++{
		wg.Add(1)
		go func() {
			fmt.Println(s.pop())
			wg.Done()
		}()

	}

	wg.Wait()
	fmt.Println(*s)


}

func TestNewStack_2(t *testing.T) {

	Convey("测试版本2", t, func() {

		Convey("先push 然后 pop测试", func() {
			s2 := NewStack_2()
			var count = 100

			for i := 0; i < count; i++ {

				s2.push_2(i)

			}

			time.Sleep(5 * time.Second)
			fmt.Println(*s2, len(s2.data))

			for i := 0; i < count; i++ {
				s2.pop_2()
			}
			time.Sleep(5 * time.Second)
			fmt.Println(*s2)
			So(len(s2.data), ShouldEqual, 0)
		})
		Convey("同时push 和 pop \n", func() {
			s2 := NewStack_2()
			var count = 100

			for i := 0; i < count; i++ {

				go func(value int) {
					s2.push_2(value)
				}(i)

			}

			for i := 0; i < count; i++ {
				go func() {
					s2.pop_2()
				}()
			}
			time.Sleep(5 * time.Second)
			fmt.Println(*s2)

			So(len(s2.data), ShouldEqual, 0)
		})
	})
}
