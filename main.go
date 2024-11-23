package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Хранилище данных
type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*") // Разрешаем запросы с любого источника
    w.Header().Set("Access-Control-Allow-Methods", "POST")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

    if r.Method == http.MethodOptions { 
        return
    }

	var user User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		http.Error(w, "Ошибка при декодировании JSON", http.StatusBadRequest)
		return
	}

	// Печатаем данные пользователя (в реальной жизни их нужно будет сохранить в базе)
	fmt.Printf("Пользователь зарегистрирован: %+v\n", user)

	// Отправляем ответ о успешной регистрации
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Регистрация успешна",
		"username": user.Username,
	})
}

func main() {
	http.HandleFunc("/register", registerHandler)

	// Запускаем сервер
	fmt.Println("Сервер запущен на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
