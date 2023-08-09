# FinalTaskFromMediaSoft
Финальное задание по курсу "Разработка на Go" от MediaSoft, три микросервиса: Ресторан, Работник и Сервер статистики.
Реализовано три базы данных на PostgreSQL:
1: dbname=Customer user=postgres password=159753 sslmode=disable port=5432
2: dbname=Restourant user=postgres password=159753 sslmode=disable port=5432
3: dbname=Statistics user=postgres password=159753 sslmode=disable port=5432
Их необходимо создать на своем компьютере
Для работ баз данных с uuid, необходимо после создания баз данных прописать команду CREATE EXTENSION IF NOT EXISTS "uuid-ossp"; в каждой базе данных
С базами данных Restourant и Statistics реализована работа с помощью GORM, и поэтому включена автомиграция при инициализации баз данных 
С базой данных Customer реализована работа с помощью нативного SQL, поэтому в проекте прописан файл миграций и dbconfig, для работы с миграция используется sql-migrate
Сервер kafka на своем компьютере необходимо настроить на работу по адресу: localhost:9092
