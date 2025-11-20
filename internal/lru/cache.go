package lru

type KeyT = int
type ValueT = string

type Cache struct {
	data  map[KeyT]ValueT
	order []KeyT
	size  int
}

func New(capacity int) *Cache {
	return &Cache{
		data:  make(map[KeyT]ValueT, capacity),
		order: make([]KeyT, capacity),
		size:  capacity,
	}
}

func (ca *Cache) Get(key KeyT) (ValueT, bool) { // получает
	val, ok := ca.data[key]
	if !ok {
		return val, ok
	}

	var idx int

	for i, el := range ca.order {
		if el == key {
			idx = i
			break
		}
	}

	if idx < len(ca.order)-1 {
		ca.order = append(
			append(ca.order[:idx], ca.order[idx+1:]...),
			key)
	}

	return val, ok
}

// кладет значение по ключу в кеш
func (ca *Cache) Put(key KeyT, value ValueT) {
	if _, ok := ca.data[key]; !ok {
		delete(ca.data, ca.order[0]) // ИСПРАВИТЬ БАГ, 2 НЕ ВЫВОДИТСЯ
		ca.order = append(ca.order[1:], key)
	}
	ca.data[key] = value
}

func (ca *Cache) Size() int {
	return len(ca.data)
}

// func (ca *Cache) String() string {
// 	return ""
// }
