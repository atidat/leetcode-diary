package main

type mixKV struct {
	key int
	val int
}

type LRUCache struct {
	mixStack []*mixKV
	kvps	 map[int]*mixKV
	capacity int
}

/*
	mixStack：存储指针，指向{k, v} 键-值对
	kvps：存储{k, {k, v}} 键-键值对
	capacity：设定容量
*/
func Constructor(capacity int) LRUCache {
	return LRUCache{
		mixStack: make([]*mixKV, 0, 2),
		kvps: make(map[int]*mixKV, 0),
		capacity: capacity,
	}
}

func (this *LRUCache) updateOrder(key int) {
	//fmt.Println()
	//fmt.Println("to be update", this.kvps, this.mixStack)
	for ki, vi := range this.mixStack {
		if vi.key == key {
			if ki < len(this.kvps) - 1 {
				tmpvi := vi
				//fmt.Println("update 1.1", this.kvps, this.mixStack)
				this.mixStack = append(this.mixStack[:ki], this.mixStack[ki+1:]...)
				//fmt.Println("update 1.2", this.kvps, this.mixStack)
				this.mixStack = append(this.mixStack, tmpvi)
				//fmt.Println("update 1.3", this.kvps, this.mixStack)
			}
		}
	}
	//fmt.Println("updated", this.kvps, this.mixStack)
	//fmt.Println()
}

func (this *LRUCache) Get(key int) int {
	//fmt.Println(this.kvps, this.mixStack)
	kv, ok := this.kvps[key]
	if !ok {
		return -1
	}
	this.updateOrder(key)
	//fmt.Println("Get", this.kvps, this.mixStack)
	return kv.val
}

func (this *LRUCache) Put(key int, value int) {
	// 1、key已存在
	kv, ok := this.kvps[key]
	if ok {
		kv.val = value
		this.updateOrder(key)
		//fmt.Println("Put1", this.kvps, this.mixStack)
		return
	}

	// fmt.Println(len(this.mixStack), this.capacity)
	// 2、key不存在，则需要添加key，若添加前已满，还需要剔除一个
	if len(this.kvps) < this.capacity {
		this.kvps[key] = &mixKV{key: key, val: value}
		tmp, _ := this.kvps[key]
		this.mixStack = append(this.mixStack, tmp)
		//fmt.Println("Put1", this.kvps, this.mixStack)
	} else {
		//fmt.Println("Put2.1", this.kvps, this.mixStack)
		delete(this.kvps, this.mixStack[0].key)
		// fmt.Println("Put2.2", this.kvps, this.mixStack)
		this.kvps[key] = &mixKV{key: key, val: value}
		tmp, _ := this.kvps[key]
		this.mixStack = this.mixStack[1:]
		this.mixStack = append(this.mixStack, tmp)
		//fmt.Println("Put2.3", this.kvps, this.mixStack)
	}
}
