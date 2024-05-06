package src

import (
	"github.com/olahol/melody"
	"log"
	"net/http"
)

func Message() {
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
		// Здесь вы можете добавить логику обработки сообщений от клиентов
		// Например, вы можете обрабатывать добавление или удаление символов в документе
		// После обработки сообщения, вы можете транслировать его в соответствующий канал Melody

		err := m.Broadcast(msg)
		if err != nil {
			log.Println("Ошибка при трансляции сообщения: ", err)
		}
	})

	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
