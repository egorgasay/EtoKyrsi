# ЭтоКурсы  ![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/egorgasay/EtoKyrsi) ![GitHub issues](https://img.shields.io/github/issues/egorgasay/EtoKyrsi) ![License](https://img.shields.io/badge/license-MIT-green)

Веб-приложение написаное на Go (с использованием фреймворка Gin), для преподавателей, которые хотят разместить свои курсы в интернет пространстве.

# Превью

## Создание урока в пару кликов!
  
![image](https://user-images.githubusercontent.com/128689324/227157135-c0539b6d-9570-4398-92b0-837fd6f81c3b.png)  
![image](https://user-images.githubusercontent.com/128689324/227157557-efaec6b5-2cbf-4db7-97f0-964911675272.png)

## В конце каждого урока располагается поле для получения ответа студента    
  
![image](https://user-images.githubusercontent.com/128689324/227157895-37544bdb-14e1-46f2-843a-7e2e7a2006dc.png)

## Удобная панель для проверки работ!  
  
![image](https://user-images.githubusercontent.com/128689324/227160435-0542bd82-935a-4646-bc2c-7ed58c37baa2.png)

# MUP
## Для создания уроков был разработан язык разметки MUP (MessageUP) и его конвертация в HTML
  
# Документация
## Правила использования
Все @ операторы должны быть отделены пустой строкой сверху и *снизу   
*операторы которые не модифицируют текст

## Поддержка html
MUP можно использовать вместе с html кодом

## 1. Стандартный блок сообщения
```python3
@msg
```
![image](https://user-images.githubusercontent.com/128689324/227161716-fb85d675-3f53-42c5-88b6-41ffde65dad4.png)
## 1. Стандартный блок сообщения
```python3
@msg
```
![image](https://user-images.githubusercontent.com/128689324/227161716-fb85d675-3f53-42c5-88b6-41ffde65dad4.png)

## 2. Предупреждающий блок сообщения
```python3
@msg-warn
```
![image](https://user-images.githubusercontent.com/102957432/230735207-8a74be14-e8ac-4f22-a033-e265b1d1b96c.png)


## 3. Заголовок в виде блока сообщения
```python3
@msg-header
```
![image](https://user-images.githubusercontent.com/102957432/230735218-a5ded23a-8495-49a4-8df5-82ca70506484.png)


## 4. Перенос строки
```python3
@n
```
![image](https://user-images.githubusercontent.com/102957432/230735226-d56f7e8e-8970-4ee2-a6ca-8bfc4ced4dcc.png)


## 5. Заголовок
```python3
@header
```
![image](https://user-images.githubusercontent.com/102957432/230735238-418653cc-6e86-4e6b-a97d-5f3408851c67.png)

## 6. Кнопка-ссылка
```python3
@btn-link  
Text!Test  
Link!https://github.com/egorgasay
```
![image](https://user-images.githubusercontent.com/102957432/230735289-68825fe0-61e0-41dd-91ad-97abfb117b47.png)

## 7. Жирный текст
```python3
@text
```
![image](https://user-images.githubusercontent.com/102957432/230735294-a7bf203a-ab02-44ff-a4fe-ed810a3b4cb6.png)


## 7. Обычный текст (по умолчанию)
```python3
@st
```
![image](https://user-images.githubusercontent.com/102957432/230735304-ce0e1802-9e83-4715-8907-dd0618acebed.png)
