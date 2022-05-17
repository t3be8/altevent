# Altevent API Project
[![Go Reference](https://pkg.go.dev/badge/golang.org/x/example.svg)](https://pkg.go.dev/golang.org/x/example)
[![Go.Dev reference](https://img.shields.io/badge/gorm-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/gorm.io/gorm?tab=doc)
[![Go.Dev reference](https://img.shields.io/badge/echo-reference-blue?logo=go&logoColor=white)](https://github.com/labstack/echo)
[![Go Reference](https://img.shields.io/badge/midtrans-reference-blue?logo=Midtrans&logoColor=white)](https://github.com/Midtrans/midtrans-go)

# Table of Content
- [Description](#description)
- [How to Use](#how-to-use)
- [Database Schema](#database-schema)
- [Testing Coverage]($testing-coverage)
- [Feature](#feature)
- [Endpoints](#endpoints)
- [Credits](#credits)

# Description
Altevent API merupakan project group ke 2 kloborasi FE & BE pada altera academy. 
project ini merupakan sebuah aplikasi Event Planning yang memiliki fitur sebagai berikut :
1. `User` dapat membuat akun sebagai Users or Organizer
2. `User` melakukan login
3. `Organizer` bisa membuat, mengedit, dan menghapus Event
4. `Users` bisa memberikan komentar pada sebuah Event

# Database Schema
![ERD](https://github.com/t3be8/altevent/blob/main/screenshoot/event-erd.png)

# Testing Coverage
Implement Unit Testing average 


# Feature
List of overall feature in this Project (To get more details see the API Documentation below)
| No.| Feature        | Keterangan                                                             |
| :- | :------------- | :--------------------------------------------------------------------- |
| 1. | Register       | Authentication process                                                 |
| 2. | Login          | Authentication process                                                 |
| 3. | CRUD Event     | Create, Read, Update, and Delete Event                                 |
| 4. | Manage Users   | Read, Update , Delete User                                             |
| 5. |                |                                                                        |


# How to Use
- Clone this repository in your $PATH:
```
$ git clone https://github.com/t3be8/altevent.git
```
- Cp file .env based on this project 
``
cp sample-env .env
``
- Don't forget to create database name as you want in your MySQL
- Run program with command
```
go run main.go
```
# Endpoints
Read the API documentation here [API Endpoint Documentation](https://app.swaggerhub.com/) (Swagger)

# Credits
- [Galang Adi Puranto](https://github.com/adeeplearn) (Author)
- [Alka Prasetya](https://github.com/alkaprasetya) (Author)
- [Rizki Firdaus](https://github.com/marthadinatarf) (Author)

# Spesial Support
- [Jerry Young](https://github.com/jackthepanda96) (Mentor)
- [Pringgo GW] 
- [Fahri]
