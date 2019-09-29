package main

import (
	"html/template"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/taskqueue"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/worker", workerHandler)

	appengine.Main()
}

type Counter struct {
	Name  string
	Count int
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	if name := r.FormValue("name"); name != "" {
		t := taskqueue.NewPOSTTask("/worker", map[string][]string{"name": {name}})
		if _, err := taskqueue.Add(ctx, t, ""); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	q := datastore.NewQuery("Counter")
	var counters []Counter
	if _, err := q.GetAll(ctx, &counters); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := handlerTemplate.Execute(w, counters); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// OK
}

func workerHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	name := r.FormValue("name")
	key := datastore.NewKey(ctx, "Counter", name, 0, nil)
	var counter Counter
	if err := datastore.Get(ctx, key, &counter); err == datastore.ErrNoSuchEntity {
		counter.Name = name
	} else if err != nil {
		log.Errorf(ctx, "%v", err)
		return
	}
	counter.Count++
	if _, err := datastore.Put(ctx, key, &counter); err != nil {
		log.Errorf(ctx, "%v", err)
	}
}

var handlerTemplate = template.Must(template.New("handler").Parse(handlerHTML))

const handlerHTML = `
{{range .}}
<p>{{.Name}}: {{.Count}}</p>
{{end}}
<p>Start a new counter:</p>
<form action="/" method="POST">
<input type="text" name="name">
<input type="submit" value="Add">
</form>
`
