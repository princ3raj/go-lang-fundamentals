package main
import(
	"net/http"
	"fmt"
	"log"
)

func main(){
	http.HandleFunc("/", handler) //each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//handler echoes the path component of the request url
func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "URL.path = %q\n", r.URL.Path)
}