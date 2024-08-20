// Создай сервер с использованием пакета go-chi. Сервер должен иметь несколько основных маршрутов с
// подмаршрутами.
// Первая группа маршрутов должна возвращать статусный код 200 и текст:
// /group1/1 Group 1 Привет, мир 1
// /group1/2 Group 1 Привет, мир 2.
// /group1/3 Group 1 Привет, мир 3.
// Вторая группа маршрутов должна возвращать статусный код 200 и текст
// /group2/1 Group 2 Привет, мир 1.
// /group2/2 Group 2 Привет, мир 2.
// /group2/3 Group 2 Привет, мир 3.
// Третья группа маршрутов должна возвращать статусный код 200 и текст
// /group3/1 Group 3 Привет, мир 1.
// /group3/2 Group 3 Привет, мир 2.
// /group3/3 Group 3 Привет, мир 3.

package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/group1/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Group 1 Привет, мир " + id))
	})
	r.Get("/group2/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Group 2 Привет, мир " + id))
	})
	r.Get("/group3/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Group 3 Привет, мир " + id))
	})
	return r
}

func main() {
	r := NewRouter()
	http.ListenAndServe(":8080", r)
}
