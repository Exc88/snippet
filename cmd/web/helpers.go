package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s\n", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)
	http.Error(w, http.StatusText(500), http.StatusInternalServerError)
}
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}
func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
	fmt.Println()
}
func (app *application) render(w http.ResponseWriter, r *http.Request, name string, td *templateData){
ts,ok := app.templateCache[name]
if !ok {
	app.serverError(w,fmt.Errorf("Шаблон %s не существует!",name))
	return
}
err := ts.Execute(w,td)
if err != nil{
	app.serverError(w,err)
}
}