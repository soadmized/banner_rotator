# Banners Rotator - сервис ротации баннеров

### Запуск сервиса
1. Создать в корне проекта `.env` файл, переменные и значение описаны в файле `.env.dist`
2. Запустить MongoDB `docker run -d --name mongo_rotator -p 27017:27017  mongo`
3. Запустить сервис `go run main.go`

### Описание API
1. `/remove_banner` - удаляет указанный баннер из ротации в указанном слоте.  
Body:
```
{
   "slot_id": string,
   "banner_id": string
}
```
2. `/add_banner` - добавляет указанный баннер в ротацию в указанном слоте.  
Body:
```
{
   "slot_id": string,
   "banner_id": string
}
```
3. `/add_click` - добавляет один клик для указанных слота, баннера и демогруппы.  
Body:
```
{
   "slot_id": string,
   "banner_id": string,
   "group_id": string
}
```
4. `/pick_banner` - выбирает баннера для показа указанной демогруппе в указанном слоте. Используется алгоритм epsilon-greedy.  
Body:
```
{
   "slot_id": string,
   "group_id": string
}
```

### Запуск тестов
Тесты можно запустить командой `make test`
