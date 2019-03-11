package main
import (
"net/http"
"html/template"
"path"
)


type Book struct {

Title string
Author string

}




func main () {
 http.HandleFunc("/",ShowBooks);
 http.ListenAndServe(":8080",nil);
}


func ShowBooks(w http.ResponseWriter, r *http.Request){

	  book:=Book{"Krall","Rozmowy z katem"}

	  fp:=path.Join("templates","index.html")
	  tmpl,err:=template.ParseFiles(fp)



	  if err!=nil {
		  http.Error(w,err.Error(),http.StatusInternalServerError);
          return
	  }


	  if err:=tmpl.Execute(w,book);err!=nil{
		  http.Error(w,err.Error(),http.StatusInternalServerError)
	  }

}
