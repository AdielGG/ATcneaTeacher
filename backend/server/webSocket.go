package server

import (
	"app-profesor/backend/server/auth"
	"app-profesor/backend/server/variables"
	"log"
	"net"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/grandcat/zeroconf"
)

func StartWebSocketServer() {
	http.HandleFunc("/ws", handleConnections)
	http.HandleFunc("/auth/login", auth.Login)
	http.HandleFunc("/auth/signup", auth.Register)

	// Obtener IP local
	ifaces, _ := net.Interfaces()
	var hostIP string
	for _, iface := range ifaces {
		addrs, _ := iface.Addrs()
		for _, addr := range addrs {
			ipnet, ok := addr.(*net.IPNet)
			if ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
				hostIP = ipnet.IP.String()
				break
			}
		}
		if hostIP != "" {
			break
		}
	}

	// Publicar el servicio mDNS (visible como "ws-server.local")
	server, err := zeroconf.Register(
		"Go WebSocket Server", // nombre visible
		"_ws._tcp",            // tipo de servicio
		"local.",              // dominio
		8081,                  // puerto
		[]string{"path=/ws"},  // TXT records opcionales
		nil,                   // interfaz de red (nil = todas)
	)
	if err != nil {
		panic(err)
	}
	defer server.Shutdown()

	go registerService(8081)

	go func() {
		log.Println("Servidor WebSocket en :8081/ws")
		err := http.ListenAndServe("0.0.0.0:8081", nil)
		if err != nil {
			log.Fatalf("Error WebSocket: %v", err)
		}
	}()
}

func SendToAll(message []byte) {
	variables.Mutex.Lock()
	defer variables.Mutex.Unlock()
	for client := range variables.Users {
		err := variables.Users[client].Conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Println("Error enviando:", err)
			variables.Users[client].Conn.Close()
			delete(variables.Users, client)
		}
	}
}

func registerService(port int) {
	server, err := zeroconf.Register("MyGoWS", "_ws._tcp", "local.", port, []string{"txtvers=1"}, nil)
	if err != nil {
		log.Fatal("Failed to register mDNS service:", err)
	}
	log.Println("mDNS service registered as 'MyGoWS._ws._tcp.local'")
	defer server.Shutdown()

	select {} // keep alive
}
