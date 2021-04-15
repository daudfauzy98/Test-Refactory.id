package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

type Body struct {
	Counter int `json:"counter"`
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/log", SendLog)

	fmt.Println("Server listen on : 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func SendLog(w http.ResponseWriter, r *http.Request) {
	var pesanBody Body
	//currentTime := time.Now().Format("2006-01-02T15:04:05Z07:00")

	if r.Method != "POST" {
		WrapAPIError(w, r, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	header := r.Header.Get("X-RANDOM")
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		WrapAPIError(w, r, "Tidak bisa membaca body!", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &pesanBody)
	if err != nil {
		WrapAPIError(w, r, "Error unmarshal : "+err.Error(), http.StatusInternalServerError)
		return
	}

	f, err := os.OpenFile("server.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer f.Close()
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	log.SetOutput(f)
	log.Printf("Success: %v http://%v/ {\"counter\": %v, \"X-RANDOM\": %q}", r.Method, GetIP(r), pesanBody.Counter, header)

	WrapAPIData(w, r, http.StatusOK, "Success!")
}

/*func GetIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}*/

func GetIP(r *http.Request) string {
	//Get IP from the X-REAL-IP header
	ip := r.Header.Get("X-REAL-IP")
	netIP := net.ParseIP(ip)
	if netIP != nil {
		return ip
	}

	//Get IP from X-FORWARDED-FOR header
	ips := r.Header.Get("X-FORWARDED-FOR")
	splitIps := strings.Split(ips, ",")
	for _, ip := range splitIps {
		netIP := net.ParseIP(ip)
		if netIP != nil {
			return ip
		}
	}

	//Get IP from RemoteAddr
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return ""
	}
	netIP = net.ParseIP(ip)
	if netIP != nil {
		return ip
	}
	return ""
}

func WrapAPIError(w http.ResponseWriter, r *http.Request, message string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	result, err := json.Marshal(map[string]interface{}{
		"code":          code,
		"error_type":    http.StatusText(code),
		"error_details": message,
	})
	if err == nil {
		w.Write(result)
	} else {
		log.Println(fmt.Sprintf("can't wrap API error : %s", err))
	}
}

func WrapAPIData(w http.ResponseWriter, r *http.Request, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	result, err := json.Marshal(map[string]interface{}{
		"code":   code,
		"status": message,
	})
	if err == nil {
		//log.Println(message)
		w.Write(result)
	} else {
		log.Println(fmt.Sprintf("can't wrap API data : %s", err))
	}
}
