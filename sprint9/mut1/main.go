package main

// пример использования sync.RWMutex
/* type Map struct {
	mu sync.RWMutex
	m  map[string]string
}

func (m *Map) Get(key string) string {
	// RLock() даёт возможность нескольким горутинам читать данные
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.m[key]
}

func (m *Map) Set(key string, v string) {
	// Lock() блокирует все операции
	m.mu.Lock()
	defer m.mu.Unlock()
	m.m[key] = v
}

func main() {
	m := Map{
		m: make(map[string]string),
	}
	for range 10 {
		go func() {
			m.Set("a", ".")
			time.Sleep(50 * time.Millisecond)
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			count := 0
			for {
				v := m.Get("a")
				count++
				if count%10 == 0 {
					fmt.Print(v)
				}
				time.Sleep(20 * time.Millisecond)
			}
		}()
	}
	time.Sleep(1 * time.Second)
} */

// тут состояние гонки
/* func main() {
	var wg sync.WaitGroup
	wg.Add(5)
	counter := 0
	for range 5 {
		go func() {
			counter++
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("OK", counter)
} */

/* type Cache struct {
	m  map[int]int
	mu sync.Mutex
}

func (cache *Cache) Get(i int) int {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	v, ok := cache.m[i]
	if ok {
		return v
	}
	// получаем значение для указанного ключа
	v = 2 * i

	cache.m[i] = v
	return v
}

func main() {
	cache := Cache{
		m: make(map[int]int),
	}
	for i := 0; i < 20; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				cache.Get(j)
			}
		}()
	}
	time.Sleep(1 * time.Second)
	fmt.Println(len(cache.m))
} */

/* type Counter struct {
	mu      sync.Mutex
	counter int
}

func (counter *Counter) Inc() {
	counter.mu.Lock()
	counter.counter++
	counter.mu.Unlock()
}

func main() {
	var c1, c2 Counter
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 100; j++ {
				c1.Inc()
				c2.Inc()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(c1.counter, c2.counter)
}
*/
