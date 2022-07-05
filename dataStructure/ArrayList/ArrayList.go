package ArrayList

import (
	"errors"
	"fmt"
)

//问题
//接口实现前边的括号追加内容,为方法内使用的外部对象
//interface{}是什么 []interface{}又是什么
//数组[:]里的:是怎么用的

type List interface {
	Size() int
	Get(index int) (interface{}, error)
	Set(index int, newval interface{}) error
	Insert(index int, newval interface{}) error
	Appand(newval interface{})
	Clear()
	Delete(index int) error
	String() string
}

type ArrayList struct {
	//数组结构
	dataStore []interface{} //数组的存储
	TheSize   int           //数组的大小
}

func NewArrayList() *ArrayList {
	//初始化数组链表
	list := new(ArrayList)                      //初始化结构
	list.dataStore = make([]interface{}, 0, 10) //开辟空间
	list.TheSize = 0
	return list
}

func (list *ArrayList) checkisFull() {
	//检查数组是否已满
	if list.TheSize == cap(list.dataStore) { //cap 判断空间
		newdataStore := make([]interface{}, 2*list.TheSize, 2*list.TheSize)
		copy(newdataStore, list.dataStore) //把list.datastore 拷贝到newdataStore
		list.dataStore = newdataStore      //赋值
	}
}

func (list *ArrayList) Size() int {
	//数组大小
	return list.TheSize //返回数组大小
}

func (list *ArrayList) Get(index int) (interface{}, error) {
	//获取数据
	if index < 0 || index >= list.TheSize {
		return nil, errors.New("索引越界")
	}
	return list.dataStore[index], nil
}

func (list *ArrayList) Set(index int, newval interface{}) error {
	//设置数据
	if index < 0 || index >= list.TheSize {
		return errors.New("索引越界")
	}
	list.dataStore[index] = newval
	return nil

}

func (list *ArrayList) Insert(index int, newval interface{}) error {
	//插入数据
	if index < 0 || index >= list.TheSize {
		return errors.New("索引越界")
	}
	list.checkisFull()                               //检测内存,如果满了,自动追加
	list.dataStore = list.dataStore[:list.TheSize+1] //插入数据需要内存移位
	for i := list.TheSize; i > index; i-- {
		list.dataStore[i] = list.dataStore[i-1]
	}
	list.dataStore[index] = newval
	list.TheSize++
	return nil
}

func (list *ArrayList) Appand(newval interface{}) {
	//追加数据
	list.dataStore = append(list.dataStore, newval)
	list.TheSize++
}

func (list *ArrayList) Clear() {
	//清空数组
	list.dataStore = make([]interface{}, 0, 10) //重新开辟空间
	list.TheSize = 0
}

func (list *ArrayList) Delete(index int) error {
	//删除数组
	list.dataStore = append(list.dataStore[:index], list.dataStore[index+1:]) //重新叠加跳过index
	list.TheSize--
	return nil
}

func (list *ArrayList) String() string {
	//字符串打印数组
	return fmt.Sprint(list.dataStore)
}
