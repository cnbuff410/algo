package hashset

type Hashset map[int] bool

func NewSet() Hashset {
        return make(Hashset)
}

func (this Hashset) Add(value int) {
        this[value] = true
}

func (this Hashset) Remove(value int) {
        delete(this, value)
}

func (this Hashset) Contains(value int) bool {
        _, ok := this[value]
        return ok
}

func (this Hashset) Length() int {
        return len(this)
}

func (this Hashset) Union(that Hashset) Hashset{
        ns := NewSet()
        for k, v := range(this) {
                ns[k] = v
        }
        for k, _ := range(that) {
                if _, ok := this[k]; !ok {
                        ns[k] = true
                }
        }
        return ns
}

func (this Hashset) Intersection(that Hashset) Hashset {
        ns := NewSet()
        for k, _ := range(that) {
                if _, ok := this[k]; ok {
                        ns.Add(k)
                }
        }
        for k, _ := range(this) {
                if _, ok := that[k]; ok {
                        ns.Add(k)
                }
        }
        return ns
}
