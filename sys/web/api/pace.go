package api
import (
  "fmt"
  "encoding/json"
	"net/http"
  "io/ioutil"

  "github.com/gorilla/mux"
  "github.com/informaticaipsfa/tunel/sys"
  "github.com/informaticaipsfa/tunel/mdl/sssifanb"

)


//ConsultarPACE Militar
func (p *Militar) ConsultarPACE(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
  var M sssifanb.Mensaje
  var cedula = mux.Vars(r)
  url := "http://" + sys.HostIPPace + sys.HostUrlPace + cedula["id"]
  //fmt.Println(url);
  response, err := http.Get(url)
	 if err != nil {
     M.Mensaje = err.Error()
     M.Tipo = 0
     w.WriteHeader(http.StatusForbidden)
     j, _ := json.Marshal(M)
 		 w.Write(j)
		 return
	 } else {
      body, err := ioutil.ReadAll(response.Body)
      if err != nil {
        w.WriteHeader(http.StatusForbidden)
        M.Mensaje = err.Error()
        M.Tipo = 0
        j, _ := json.Marshal(M)
    		w.Write(j)
        return
      }
			 defer response.Body.Close()
			 w.WriteHeader(http.StatusOK)
       w.Write(body)
       return
	 }
}
