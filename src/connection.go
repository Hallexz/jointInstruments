package src

import (
	"github.com/olahol/melody"
	"log"
	"net/http"
	"strings"
)

func ConnectToMelody() {
	m := melody.New()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		err := m.HandleRequest(w, r)
		if err != nil {
			log.Println("Ошибка при обработке запроса: ", err)
		}
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		message := string(msg)

		if strings.Contains(message, ":") {
			operation := strings.Split(message, ":")[0]
			data := strings.Split(message, ":")[1]

			switch operation {
			case "add":
				log.Printf("Добавление символа: %s", data)
			case "delete":
				log.Printf("Удаление символа в позиции: %s", data)
			case "subscribe":
				log.Printf("Подписка: %s", data)
			case "unsubscribe":
				log.Printf("Отписка: %s", data)
			default:
				log.Printf("Неизвестная операция: %s", operation)
			}
		} else {
			log.Printf("Сообщение не содержит двоеточие: %s", message)
		}
	})

	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
