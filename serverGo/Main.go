package main

import (
    "encoding/json"
    "fmt"
	"os"
    "net/http"
	
)

type Message struct {
    Text string `json:"text"`
}

func handler(w http.ResponseWriter, r *http.Request) {
    // Permitir solicitudes desde cualquier origen
    w.Header().Set("Access-Control-Allow-Origin", "*")

    if r.Method == "OPTIONS" {
        // Preflight request
        w.Header().Set("Access-Control-Allow-Methods", "POST")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
        return
    }

    // Decodificar el JSON recibido
    var msg Message
    err := json.NewDecoder(r.Body).Decode(&msg)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Imprimir el mensaje recibido
    fmt.Printf("Mensaje recibido:\n%s\n", msg.Text)


	//Guardar el mensaje en un archivo .txt

	file, err := os.Create("mensaje.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	file.WriteString(msg.Text)


	
    // Responder al cliente
    response := Message{Text: "¡Mensaje recibido correctamente!"}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func main() {
    // Ruta para manejar las solicitudes JSON
    http.HandleFunc("/json", handler)

    // Iniciar el servidor en el puerto 8080
    fmt.Println("Servidor en ejecución en http://localhost:8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Printf("Error al iniciar el servidor: %s\n", err)
    }
}
