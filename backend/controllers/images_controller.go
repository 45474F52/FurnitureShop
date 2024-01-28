package controllers

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type ImagesController struct {
	PathToImages string
}

func (controller *ImagesController) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		image := r.URL.Query().Get("image")

		if image == "" {
			image = controller.PathToImages + "\\" + "defaultImage.jpeg"
		}

		file, err := os.Open(image)
		if err != nil {
			http.Error(w, "Ошибка при чтении файла", http.StatusInternalServerError)
			fmt.Printf("Ошибка при чтении файла \"%s\"\n", image)
			return
		}
		defer file.Close()

		fileInfo, err := file.Stat()
		if err != nil {
			http.Error(w, "Ошибка при получении информации о файле", http.StatusInternalServerError)
			fmt.Printf("Ошибка при получении информации о файле \"%s\"\n", image)
			return
		}

		w.Header().Set("Content-Type", "image/jpeg")
		w.Header().Set("Content-Length", fmt.Sprint(fileInfo.Size()))

		_, err = io.Copy(w, file)
		if err != nil {
			http.Error(w, "Ошибка при отправке изображения", http.StatusInternalServerError)
			fmt.Printf("Ошибка при отправке изображения \"%s\"\n", fileInfo.Name())
			return
		}
	}
}
