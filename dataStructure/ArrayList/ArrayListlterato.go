package ArrayList

/*
数组迭代器
*/
type Iterator interface {
	HasNext()                   //是否有下一个
	Next() (interface{}, error) //下一个元素
	Remove()                    //删除
	GetIndex() int              //得到索引
}



func () HasNext() {
	//TODO implement me
	panic("implement me")
}

func () Next() (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func () Remove() {
	//TODO implement me
	panic("implement me")
}

func () GetIndex() int {
	//TODO implement me
	panic("implement me")
}

type Iterable interface {
	Iterator() Iterator //构造初始化接口
}

//构造指针访问数组
type ArraylistIterator struct {
	list         *ArrayList //数组指针 相当于一个类包含了另一个类
	currentindex int        //当前索引
}

func (list *ArrayList) Iterator() Iterator {
	it := new(ArraylistIterator)
	it.currentindex = 0
	it.list = list
	return it
}
