package main

import (
  "net/http"
  "fmt"
  "golang.org/x/oauth2"
  "io/ioutil"
  "encoding/json"
  "strconv"
  "time"
  "os/exec"
)

func apiAdminEndpoint(w http.ResponseWriter, r *http.Request){
  fmt.Println(PROJECT_FOLDER +"/front/admin_login.html" )
  http.ServeFile(w, r, PROJECT_FOLDER +  "/front/admin_login.html")  
}

func apiNotificarEndpoint(w http.ResponseWriter, r *http.Request){
  http.ServeFile(w, r, PROJECT_FOLDER +  "/front/notificar_infectat.html")  
}

func apiCercarEstudiantEndpoint(w http.ResponseWriter, r *http.Request){
  alumne:=r.FormValue("usuari")
  fmt.Println(alumne)

  exec.Command("cp", PROJECT_FOLDER + "/front/estat_estudiant.html", PROJECT_FOLDER + "/front/estat_estudiant_" + alumne  + ".html").Run()

  exec.Command("sed", "-i", "s/%%NOM\\.COGNOM%%/" + alumne + "/g", PROJECT_FOLDER + "/front/estat_estudiant_" + alumne + ".html").Run()

  exec.Command("sed", "-i", "s/%%ESTAT%%/" + estatAlumne[alumne] + "/g", PROJECT_FOLDER + "/front/estat_estudiant_" + alumne + ".html").Run()

  http.ServeFile(w, r, PROJECT_FOLDER + "/front/estat_estudiant_" + alumne + ".html")

}

func apiEstatActualEndpoint(w http.ResponseWriter, r *http.Request){
  exec.Command("cp", PROJECT_FOLDER+ "/front/estat_actual.html", PROJECT_FOLDER + "/front/estat_actual_2.html").Run()
  
  st := ""
  a := 0
  for k, v := range estatAlumne {
    if v == "POSITIU" || v == "CONTACTE AMB POSITIU" {
      a++
      st+= "<tr><th> " + k + " </th></tr>"
    }
  }

  exec.Command("sed", "-i", "s/%%NUMERO%%/" + strconv.Itoa(a) + "/g", PROJECT_FOLDER +"/front/estat_actual_2.html").Run()
  exec.Command("sed", "-i", "s|%%LLISTA%%|" + st + "|g", PROJECT_FOLDER +"/front/estat_actual_2.html").Run()


  http.ServeFile(w, r, PROJECT_FOLDER +  "/front/estat_actual_2.html")  
  
}

func apiEnviarNotificarInfectat(w http.ResponseWriter, r *http.Request){
  usuari:=r.FormValue("usuari")
  data:=r.FormValue("data")
  
  fmt.Println(usuari)
  fmt.Println(data)
  t, err := time.Parse(time.RFC3339, data+"T00:00:00Z")
  if err != nil{
    fmt.Println(err) 
  }

  fmt.Println(t)

  cosa,_:=time.ParseDuration("-336h")
  pcrmenyscatorze:=t.Add(cosa)
  for k, v := range contactes[usuari] {
    if v.After(pcrmenyscatorze){
      estatAlumne[k]="CONTACTE AMB POSITIU"
    }
  }

  estatAlumne[usuari]="POSITIU"

  http.ServeFile(w, r, PROJECT_FOLDER + "/front/notificar_infectat.html")
  
}

func apiSentarseEndpoint(w http.ResponseWriter, r *http.Request){

}

// serveLogin gets the OAuth temp credentials and redirects the user to the
// raco's authorization page.
func serveLogin(w http.ResponseWriter, r *http.Request) {
    currentAlumne.Aula=r.FormValue("aula")
    currentAlumne.X, _=strconv.ParseFloat(r.FormValue("x"), 64)
    currentAlumne.Y, _=strconv.ParseFloat(r.FormValue("y"), 64)
//    aula := r.FormValue("aula")
//    x := r.FormValue("x")
//    y := r.FormValue("y")
    fmt.Println(r.FormValue("aula"))
    url := OAUTH_CONF.AuthCodeURL("state", oauth2.AccessTypeOffline)
    http.Redirect(w, r, url, 302)

}

// serveOAuthCallback handles callbacks from the OAuth server.
func serveOAuthCallback(w http.ResponseWriter, r *http.Request) {
    token, err := OAUTH_CONF.Exchange(oauth2.NoContext, r.FormValue("code"))

    if err != nil {
      http.Error(w, "Error exchanging", 500)
      return
    }

//    response, err := http.Get("https://api.fib.upc.edu/v2/jo/classes?access_token="+token.AccessToken)

    cl := &http.Client{}
    req,_ := http.NewRequest("GET","https://api.fib.upc.edu/v2/jo?access_token="+token.AccessToken, nil)

    req.Header.Set("Accept", "application/json")
    res, _ := cl.Do(req)

    content, _ := ioutil.ReadAll(res.Body)
    type AlumneTemp struct {
      Assignatures string
      Avisos string
      Classes string
      Foto string
      Practiques string
      Projectes string
      Username string
      Nom string
      Cognoms string
      Email string
    }

    var al AlumneTemp

    err = json.Unmarshal(content, &al)
    if err != nil {
      fmt.Println(err)
    }
    currentAlumne.Nom=al.Username 
    
    alum := Alumne {
      currentAlumne.Nom,
      currentAlumne.X,
      currentAlumne.Y,
    }

    if contactes[currentAlumne.Nom] == nil {
      contactes[currentAlumne.Nom] = make(map[string]time.Time)
      estatAlumne[currentAlumne.Nom] = "NO CONFINAT"
    }


    for _,v := range aules[currentAlumne.Aula] {
      fmt.Printf("%s  %s  %.2f %.2f %.2f %.2f\n", currentAlumne.Nom, v.Nom,currentAlumne.X, currentAlumne.Y, v.CoordX, v.CoordY)
      if dist(currentAlumne.X, currentAlumne.Y, v.CoordX, v.CoordY) < 2.1  {
        contactes[v.Nom][currentAlumne.Nom] = time.Now()
        contactes[currentAlumne.Nom][v.Nom] = time.Now()
      }
    }

    aules[currentAlumne.Aula]=append(aules[currentAlumne.Aula], alum)


    fmt.Printf("%+v", contactes)

}

// serveLogout clears the token credentials
func serveLogout(w http.ResponseWriter, r *http.Request) {
//	s := session.Get(r)
//	delete(s, tokenCredKey)
//	if err := session.Save(w, r, s); err != nil {
//		http.Error(w, "Error saving session , "+err.Error(), 500)
//		return
//	}
//	http.Redirect(w, r, "/", 302)
}
