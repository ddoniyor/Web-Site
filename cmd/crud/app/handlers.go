package app

import (

	"crud/pkg/crud/models"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
)

func (receiver *server) handleBurgersList() func(http.ResponseWriter, *http.Request) {
	tpl, err := template.ParseFiles(filepath.Join(receiver.templatesPath, "index.gohtml"))
	if err != nil {
		panic(err)
	}

	return func(writer http.ResponseWriter, request *http.Request) {
		list, err := receiver.burgersSvc.BurgersList()
		if err != nil {
			log.Print(err)
			http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		data := struct {
			Title string
			Burgers []models.Burger
		}{
			Title: "McDonalds",
			Burgers: list,
		}

		err = tpl.Execute(writer, data)
		if err != nil {
			log.Print(err)
			http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}
}

func (receiver *server) handleBurgersSave() func(responseWriter http.ResponseWriter, request *http.Request) {
	_, err := template.ParseFiles(filepath.Join(receiver.templatesPath, "index.gohtml"))
	if err != nil {
		panic(err)
	}
	return func(writer http.ResponseWriter, request *http.Request) {
		if err != nil {
			log.Print(err)
			http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		err := request.ParseForm()
		if err != nil {
			panic(err)
		}
		name := request.PostForm.Get("name")
		price := request.PostForm.Get("price")
		parsedPrice, err := strconv.Atoi(price)
		if err != nil {
			http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest) // 400 - Bad Request
			return
		}
		receiver.burgersSvc.Save(models.Burger{Name: name, Price: parsedPrice})
		http.Redirect(writer, request, "/", http.StatusPermanentRedirect)
	}
}

func (receiver *server) handleBurgersRemove() func(responseWriter http.ResponseWriter, request *http.Request) {
	_, err := template.ParseFiles(filepath.Join(receiver.templatesPath, "index.gohtml"))
	if err != nil {
		panic(err)
	}
	return func(writer http.ResponseWriter, request *http.Request) {
		if err != nil {
			log.Print(err)
			http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		err := request.ParseForm()
		if err != nil {
			panic(err)
		}
		id := request.PostForm.Get("id")
		log.Print(id)
		parsedId, err := strconv.Atoi(id)
		if err != nil {
			http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest) // 400 - Bad Request
			return
		}
		receiver.burgersSvc.RemoveById(parsedId)

		http.Redirect(writer, request, "/", http.StatusPermanentRedirect)
		return
	}
}

func (receiver *server) handleFavicon() func(http.ResponseWriter, *http.Request) {

	file, err := ioutil.ReadFile(filepath.Join(receiver.assetsPath, "favicon.ico"))
	if err != nil {
		panic(err)
	}
	return func(writer http.ResponseWriter, request *http.Request) {
		_, err := writer.Write(file)
		if err != nil {
			log.Print(err)
		}
	}
}
