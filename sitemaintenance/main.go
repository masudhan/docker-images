package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/health", healthHandler)
	log.Println("Server listening on :80...")
	http.ListenAndServe(":80", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ipAddress := r.RemoteAddr
	log.Printf("Request from IP: %s\n", ipAddress)

	// Define data to be used in the template
	data := struct {
		Title     string
		Headline  string
		Message   string
		TeamName  string
		LinkColor string
		Theme     string
	}{
		Title:     "Site Maintenance",
		Headline:  "We’ll be back soon!",
		Message:   "Sorry for the inconvenience but we’re performing some maintenance at the moment. We’ll be back online shortly!",
		TeamName:  "Madhus",
		LinkColor: "#0066cc", // You can change this to the desired link color
		Theme:     "Light",    // You can change this to "Light" if needed
	}

	// Parse the HTML template
	tmpl, err := template.New("index").Parse(indexHTML)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Execute the template with the data and write the result to the response writer
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	// Health check always returns a 200 OK response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// The HTML template string
const indexHTML = `<!doctype html>
<meta charset="UTF-8">

<head>
    <title>{{.Title}}</title>
    <style>
        body {
            text-align: center;
            padding: 20px;
            background-color: {{if eq .Theme "Dark"}}#2a2e35{{else}}white{{end}};
            color: {{if eq .Theme "Dark"}}#b2becd{{else}}black{{end}};
        }

        @media (min-width: 768px) {
            body {
                padding-top: 100px;
            }
        }

        h1 {
            font-size: 50px;
        }

        body {
            font: 20px Helvetica, sans-serif;
            color: #333;
        }

        article {
            display: block;
            text-align: left;
            max-width: 650px;
            margin: 0 auto;
        }

        a {
            color: {{.LinkColor}};
            text-decoration: none;
        }

        a:hover {
            color: #333;
            text-decoration: none;
        }
    </style>
    <script type="text/javascript">
        function applyTheme() {
            var themeType = "{{.Theme}}";
            if (themeType === "Dark") {
                document.body.style.backgroundColor = "#2a2e35";
                document.body.style.color = "#b2becd";
            }
            if (themeType === "Light") {
                document.body.style.backgroundColor = "white";
                document.body.style.color = "black";
            }
        }
        window.onload = function () {
            applyTheme();
        };
    </script>
</head>

<body>
    <article>
        <h1>{{.Headline}}</h1>
        <div>
            <p>{{.Message}}</p>
            <p>&mdash;{{.TeamName}}</p>
        </div>
    </article>
</body>
`

