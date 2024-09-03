# Weather API

## Описание

Этот API предоставляет информацию о текущей погоде и средней температуре и влажности для различных городов.

## Отчет о покрытии кода

Вы можете просмотреть отчет о покрытии кода по следующей ссылке:

[Открыть отчет о покрытии](coverage.html)

## Запуск

### 1. Docker

Запустить докер и написать комманду находясь в директории проекта `docker-compose up -d`

### 2. Локально

Иметь установленный GO. Запустить проект находясь в директории проекта `go run cmd/main.go` 

## Эндпоинты

### 1. Получить среднюю температуру и влажность по городам

**URL:** `/weather/average`  
**Метод:** `GET`  
**Описание:** Возвращает среднюю температуру и влажность для каждого города.

**Ответ:**

```json
[
    {
        "city": "minsk",
        "temp": 19,
        "humidity": 40
    },
    {
        "city": "kyiv",
        "temp": 23,
        "humidity": 22
    },
    {
        "city": "vilnius",
        "temp": 19,
        "humidity": 46
    },
    {
        "city": "riga",
        "temp": 20,
        "humidity": 54
    },
    {
        "city": "tallinn",
        "temp": 18,
        "humidity": 66
    },
    {
        "city": "sofia",
        "temp": 23,
        "humidity": 47
    }
]
```

**URL:** `/weather`  
**Метод:** `GET`  
**Описание:** Возвращает данные о погоде по всем городам.

**Ответ:**

```json
[
    {
        "city": "Rīga",
        "temp": 20.34,
        "feels_like": 20.75,
        "temp_min": 20.04,
        "temp_max": 21.62,
        "wind_speed": 2.06,
        "wind_deg": 350
    },
    {
        "city": "Tallinn",
        "temp": 18.24,
        "feels_like": 18.47,
        "temp_min": 18.06,
        "temp_max": 18.62,
        "wind_speed": 1.54,
        "wind_deg": 60
    },
    {
        "city": "Sofia",
        "temp": 24.04,
        "feels_like": 23.78,
        "temp_min": 24.04,
        "temp_max": 25.83,
        "wind_speed": 1.54,
        "wind_deg": 50
    },
    {
        "city": "Minsk",
        "temp": 21.86,
        "feels_like": 21.17,
        "temp_min": 21.86,
        "temp_max": 21.86,
        "wind_speed": 3.15,
        "wind_deg": 60
    },
    {
        "city": "Kyiv",
        "temp": 24.23,
        "feels_like": 23.78,
        "temp_min": 24.23,
        "temp_max": 26.48,
        "wind_speed": 3.6,
        "wind_deg": 70
    },
    {
        "city": "Vilnius",
        "temp": 24.49,
        "feels_like": 24.12,
        "temp_min": 24.49,
        "temp_max": 24.49,
        "wind_speed": 2.57,
        "wind_deg": 80
    }
]
```

**URL:** `/weather/:city`  
**Метод:** `GET`  
**Описание:** Возвращает данные о погоде по определенному городу. Из доступных городов: `minsk`, `kyiv`, `vilnius`, `riga`, `tallinn`, `sofia`.

**Ответ:**

```json
{
    "city": "Vilnius",
    "temp": 24.49,
    "feels_like": 24.12,
    "temp_min": 24.49,
    "temp_max": 24.49,
    "wind_speed": 2.57,
    "wind_deg": 80
}
```