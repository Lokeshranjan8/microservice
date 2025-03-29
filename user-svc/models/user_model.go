package models

type User struct {
    ID    uint   `gorm:"primaryKey" json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

type Response struct {
    Message     string      `json:"message"`
    Explanation string      `json:"explanation"`
    Data        any `json:"data"`
}
