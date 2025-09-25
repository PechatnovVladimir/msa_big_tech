# msa_big_tech
Практическая работа на  курсе Микросервисы на Go как в BigTech https://balun.courses/courses/microservice

# Домашнее задание № 1 
Был создан скелет микросервисной архитектуры четырех сервисов `auth`, `users`, `social`, `chat` и реализован доступ к ним через `grpc-gateway`.

# Домашнее задание № 2
## сервис users  
- переход на buf вместо protoc
- разделение proto-файлов
- добавлено (частично) описание и валидацию в proto файлы
- выделены сущности (users/internal/app/models)
- переход на слои
- интерфейс репозитория (users/internal/app/repositories)  
- bynthatqc юзкейса (users/internal/app/usecases) 
- репозиторий inmemory (простенький "на коленке" только для тестирования users/internal/app/repositories/inmemory) 
- кое-какая бизнес-логика в usecase
- интеграция grpc-хендлеров с usecase 



