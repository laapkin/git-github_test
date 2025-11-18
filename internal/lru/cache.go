package lru

type KeyT = int
type ValueT = string

type Cache struct {
	data     map[KeyT]ValueT
	order    []KeyT
	indexKey int // куда записывать новый элемент
	capacity int
}

func New(capacity int) *Cache {
	return &Cache{
		data:     make(map[KeyT]ValueT, capacity),
		order:    make([]KeyT, capacity),
		indexKey: 0,
		capacity: capacity,
	}
}

func (ca *Cache) Get(key KeyT) (ValueT, bool) { // получает
	val, ok := ca.data[key]

	return val, ok
}

// кладет значение по ключу в кеш
func (ca *Cache) Put(key KeyT, value ValueT) {
	// очищать/заменять данные только при забитой емкости
	if len(ca.data) == ca.capacity {
		// узнать самый старый ключ
		old_key := ca.order[ca.indexKey]

		// удалить старые данные из мапы
		delete(ca.data, old_key)
	}

	// добавить в мапу новое значение по новому ключу
	ca.data[key] = value
	// добавить новую пару в срез порядка
	ca.order[ca.indexKey] = key

	// обновить следующий индекс для записи
	ca.indexKey = (ca.indexKey + 1) % ca.capacity
}

func (ca *Cache) Size() int {
	return len(ca.data)
}

func (ca *Cache) String() string {
	return ""
}
