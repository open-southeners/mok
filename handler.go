package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/gorilla/mux"
)

type HttpResource struct {
	FilePath string
	Params   []string
}

func mapWorkingDirectory(path string, resources map[string]HttpResource, parent string) {
	entries, err := os.ReadDir(path)

	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		if e.Type().IsDir() {
			mapWorkingDirectory(path+"/"+e.Name(), resources, strings.Join([]string{parent, e.Name()}, "/"))
			return
		}

		var fileName = e.Name()

		if !strings.HasSuffix(fileName, ".json") {
			return
		}

		filePath := fileName
		filePath = strings.Replace(filePath, ".json", "", len(filePath)-5)

		var matches []string
		var params []string

		if strings.Contains(filePath, "[") && strings.Contains(filePath, "]") {
			filePath = strings.ReplaceAll(filePath, "[", "{")
			filePath = strings.ReplaceAll(filePath, "]", "}")

			compile, _ := regexp.Compile(`(?i){([a-z0-9]+)}`)

			matches = compile.FindAllString(filePath, -1)

			for _, m := range matches {
				matchedParam := m
				matchedParam = strings.ReplaceAll(matchedParam, "{", "")
				matchedParam = strings.ReplaceAll(matchedParam, "}", "")

				params = append(params, matchedParam)
			}
		}

		resources[strings.Join([]string{parent, filePath}, "/")] = HttpResource{
			FilePath: path + "/" + fileName,
			Params:   params,
		}
	}
}

func (h *HttpResource) handleRequest(w http.ResponseWriter, req *http.Request) {
	start := time.Now()
	fileContent, err := os.ReadFile(h.FilePath)

	if err != nil {
		log.Fatal(err)
	}

	w.Header().Add("Content-Type", "application/json")

	ctx := mux.Vars(req)

	gofakeit.Seed(0)
	content, err := gofakeit.Template(string(fileContent), &gofakeit.TemplateOptions{Data: ctx})

	if err != nil {
		log.Print(err)
	}

	fmt.Fprint(w, content)

	fmt.Printf("%s ........................................... %d ms\n", req.URL.Path, time.Since(start).Milliseconds())
}

func HandlesRequests(address string, path string) {
	r := mux.NewRouter()
	resources := make(map[string]HttpResource)

	mapWorkingDirectory(path, resources, "")

	for k, v := range resources {
		r.HandleFunc(k, v.handleRequest)
	}

	fmt.Printf("Listening to requests at %s\n", address)

	log.Fatal(http.ListenAndServe(address, r))
}
