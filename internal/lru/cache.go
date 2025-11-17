package lru

type KeyT = int
type ValueT = string

type Cache struct {
	Data  map[KeyT]ValueT
	Order []KeyT
}

func New(size int) *Cache {
	sizeCache := make(map[KeyT]ValueT, size)
	sizeSlice := make([]KeyT, 0, size)
	s := Cache{Data: sizeCache, Order: sizeSlice}

	return &s
}

func (ca *Cache) Get(key KeyT) (ValueT, bool) { // получает
	val, ok := ca.Data[key]

	return val, ok
}

// кладет значение по ключу в кеш
func (ca *Cache) Put(key KeyT, value ValueT) {
	if len(ca.Order) == cap(ca.Order) {
		delete(ca.Data, ca.Order[0])

		order := make([]KeyT, len(ca.Order)-1, cap(ca.Order))
		copy(order, ca.Order[1:])
		ca.Order = order
	}

	ca.Data[key] = value
	ca.Order = append(ca.Order, key)
}

func (ca *Cache) Size() int {
	return len(ca.Data)
}
