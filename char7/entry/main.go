package main

//语言值得注意的点

/*func main(){
	//使用wg来进行同步
	var wg sync.WaitGroup
	data := [10]int {0,1,2,3,4,5,6,7,8,9}

	//表现出的是range产生的临时变量被共享了引用，然后内部生成的多个goroutine先不进行计算，只是保持引用
	//相当于就是range进行迭代写，内部的十个goroutine进行引用读，最终引的就是最后一个data的索引，也就是9这个值
	for i := range data {
		wg.Add(1)
		//直接引用i时
		go func() {
			fmt.Println(i)//始终输出的是9
			wg.Done()
		}()
	}

	wg.Wait()
}*/

/*func main() {
	var wg sync.WaitGroup

	data := [10]int{0,1,2,3,4,5,6,7,8,9}

	for i := range data {

		wg.Add(1)

		go func(a int) {
			fmt.Println(a)
			wg.Done()
		}(i)//此处就对range临时变量进行了复制并使用。【实质上就跟js中使用that来临时标记this类似】
	}

	wg.Wait()
}*/

//defer的使用【对具名返回值的影响】
/*func f1() (r int) {
	defer func() {
		r++//显然具名返回值实际上相当于在局部重新声明一个局部变量
		//然后
	}()
	return 0//此处相当于先r=0，但是由于defer内引用r变量，进行++之后
	//再真正return【整体流程正因为是这样的，所以表现出的现象如下面的例子所示】
}

func f2() (r int) {
	t := 5

	defer func() {
		t = t + 5
	}()
	return t//此处相当于r=t之后，defer操作引用的都是t变量，对r无影响，最后返回r变量
}

func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1//此处过程为：r赋值为1，然后defer运行，传入r【也就是1】
	//然后在defer内的参数r并不影响返回值的r变量，所以defer操作之后对返回值无影响，最后返回r
}

//当前体现的就是值拷贝、闭包特性、以及defer对具名变量是否影响的判断
//牢记返回的流程为：具名返回值外加defer的闭包实现流程
//1。首先具名变量作为函数内局部变量被初始化
//2。在return处对具名变量进行赋值【直接返回直面量或者变量，都是一个具名变量的赋值过程】
//3。defer内只有直接引用具名返回值变量【且不为闭包内的入参】时，defer内闭包形成对返回值变量的引用【存在影响】
//	否则defer操作无影响
//4。defer执行完成之后，将具名变量的值进行返回
func main() {
	fmt.Println("f1", f1())//1
	fmt.Println("f2", f2())//5
	fmt.Println("f3", f3())//1
}

//所以真正的结论应该是不必要声明具名的返回值
*/

//切片：因多个切片共享底层数组而导致的不稳定的表现。
