package cache

// Cache интерфейс с нициализированной мапой
type Cache struct {
	myMap map[string]interface{}
}

// New создан метод для обращения в Cache
func New() *Cache {
	return &Cache{
		myMap: make(map[string]interface{}),
	}
}

// Get метод возвращает данные по ключу из myMap
func (c *Cache) Get(key string) interface{} {
	return c.myMap[key]
}

// GetMap метод возвращает все данные записанные в myMap
func (c *Cache) GetMap() interface{} {
	return c.myMap
}

// Set записывает значения в myMap
func (c *Cache) Set(key string, value interface{}) {
	c.myMap[key] = value
}

// Delete удаляет данные по ключу из myMap
func (c *Cache) Delete(key string) {
	delete(c.myMap, key)
}
