## Реализовать интерфейс Cacher для структруры Cache

```go
type Cacher interface {
	Get(string) (string, bool)
	Set(string, string)
}

type Cache struct {
    ttl time.Duration
}
```

### Исходные данные
- Методы вызываются из разных горутин и должны быть thread-safe
- Чистить кеш не надо

TTL задаётся в конструкторе, в который передается значение ttl:
```go
func NewCache(ttl time.Duration) {
}
```
По истечению TTL поведение такое, как будто значения в кеше не было записано.

### Решение
- Реализация приведена в файле cacher/cacher.go
- В main.go можно увидеть примеры вызова: запись в кеш и чтение.
- Для 
