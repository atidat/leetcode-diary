package main

func main() {

}

/*1
var alivePeople map[string]bool
var orders []string*/

type ThroneInheritance struct {
	/*1	name string
	alive bool
	childs []*ThroneInheritance*/
	root string
	name string
	childs map[string][]string
	dead map[string]bool
}


func Constructor(kingName string) ThroneInheritance {
	/*1	alivePeople = make(map[string]bool, 0)
	alivePeople[kingName] = true
	orders = []string{}
	orders = append(orders, kingName)


	return ThroneInheritance{
		name: kingName,
		alive: true,
		childs: []*ThroneInheritance{},
	}*/
	c := make(map[string][]string)
	c[kingName] = []string{}
	d := make(map[string]bool)
	d[kingName] = false

	return ThroneInheritance{
		root: kingName,
		name: kingName,
		childs: c,
		dead: d,
	}
}


func (this *ThroneInheritance) Birth(parentName string, childName string)  {
	/*1	alivePeople[childName] = true

	type functype func(*ThroneInheritance, string, string) int
	var dfs functype
	dfs = func(t *ThroneInheritance, pn, cn string) int {
		if t.name == pn {
			// 统计同父节点下，其他兄弟节点及他们子孙节点的总数
			type futy func(*ThroneInheritance) int
			var calc futy
			calc = func(t *ThroneInheritance) int {
				var re int
				for j := 0; j < len(t.childs); j++ {
					if t.childs[j].alive {
						re += 1
					}

					re += calc(t.childs[j])
				}
				return re
			}
			sonl := calc(t)

			child := ThroneInheritance{
				childName,
				true,
				[]*ThroneInheritance{},
			}
			t.childs = append(t.childs, &child)
			return sonl
		}


		for _, c := range t.childs {
			t := dfs(c, pn, cn)
			if t != -1 {
				return t
			}
		}
		return -1
	}
	sonl := dfs(this, parentName, childName)
	if sonl == -1 {
		panic("corrupted")
	}
	var k int
	for ; k < len(orders); k++ {
		if orders[k] == parentName {
			break
		}
	}

	tmp := make([]string, len(orders[k+sonl+1:]))
	copy(tmp, orders[k+sonl+1:])
	orders = append(orders[:k+sonl+1], childName)
	orders = append(orders, tmp...)*/
	if v, ok := this.childs[parentName]; ok {
		this.childs[parentName] = append(v, childName)
	} else {
		this.childs[parentName] = []string{childName}
	}
	this.dead[parentName] = false
}


func (this *ThroneInheritance) Death(name string)  {
	this.dead[name] = true
}


func (this *ThroneInheritance) GetInheritanceOrder() []string {
	/*1 return orders*/
	var res []string
	type functype func(*ThroneInheritance, string)
	var dfs functype
	dfs = func(t *ThroneInheritance, n string) {
		if !this.dead[t.name] {
			res = append(res, t.name)
		}
		for _, v := range this.childs[t.name] {
			dfs(this, v)
		}
	}
	dfs(this, this.root)
	return res
}

/**
 * Your ThroneInheritance object will be instantiated and called as such:
 * obj := Constructor(kingName);
 * obj.Birth(parentName,childName);
 * obj.Death(name);
 * param_3 := obj.GetInheritanceOrder();
 */