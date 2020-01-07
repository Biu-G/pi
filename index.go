package main
import "fmt"
import (
    "net/http"
)

var pi string

func PrintPi(pi string) {
	fmt.Println("pi: " + pi)
}

func SayHello(w http.ResponseWriter, req *http.Request) {
	/*if pi==nil {
		w.Write([]byte("pi is nil, nothing to write")) 
		return
	}*/
	pi="hello"
	w.Write([]byte("sent pi hello"))
	PrintPi(pi)
    //pi.Write([]byte("Hello"))
}

func RegisterPi(w http.ResponseWriter, req *http.Request) {
	/*if pi==nil {
		w.Write([]byte("pi is originally nil, assigned it"))
	} else {
		w.Write([]byte("pi switched to a brand new one"))
	}*/
	w.Write([]byte(pi))
	pi=""
	PrintPi(pi)
}

func GreeOn(w http.ResponseWriter, req *http.Request) {
	pi="gree_on"
	w.Write([]byte("sent gree on"))
	PrintPi(pi)
}

func GreeOff(w http.ResponseWriter, req *http.Request) {
    pi="gree_off"
    w.Write([]byte("sent gree off"))
    PrintPi(pi)
}

func GreeHot(w http.ResponseWriter, req *http.Request) {
    pi="gree_hot"
    w.Write([]byte("sent gree hot"))
    PrintPi(pi)
}

func GreeWind(w http.ResponseWriter, req *http.Request) {
    pi="gree_wind"
    w.Write([]byte("sent gree wind"))
    PrintPi(pi)
}


func main() {
	pi = ""
    http.HandleFunc("/hello", SayHello)
	http.HandleFunc("/regi", RegisterPi)
	http.HandleFunc("/gree/on", GreeOn)
    http.HandleFunc("/gree/off", GreeOff)
    http.HandleFunc("/gree/hot", GreeHot)
    http.HandleFunc("/gree/wind", GreeWind)
	PrintPi(pi)
    http.ListenAndServe(":89", nil)
}
