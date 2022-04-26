package scrape

// scrapeapi.go HAS TEN TODOS - TODO_5-TODO_14 and an OPTIONAL "ADVANCED" ASK

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"

	"github.com/gorilla/mux"
)

//==========================================================================\\

// Helper function walk function, modfied from Chap 7 BHG to enable passing in of
// additional parameter http responsewriter; also appends items to global Files and
// if responsewriter is passed, outputs to http

var files_found = 0

func walkFn(w http.ResponseWriter) filepath.WalkFunc {
	return func(path string, f os.FileInfo, err error) error {
		w.Header().Set("Content-Type", "application/json")

		for _, r := range regexes {
			if r.MatchString(path) {
				//TODO_5: As it currently stands the same file can be added to the array more than once
				//TODO_5: Prevent this from happening by checking if the file AND location already exist as a single record
				dir, filename := filepath.Split(path)
				toAdd := true
				for _, file := range Files {
					if file.Filename == string(filename) && file.Location == string(dir) {
						toAdd = false
					}
				}
				if toAdd {
					var tfile FileInfo

					tfile.Filename = string(filename)
					tfile.Location = string(dir)
					tfile.Number = files_found
					files_found += 1
					Files = append(Files, tfile)
					if w != nil && len(Files) > 0 {

						//TODO_6: The current key value is the LEN of Files (this terrible);
						//TODO_6: Create some variable to track how many files have been added
						w.Write([]byte(`"` + (strconv.FormatInt(int64(tfile.Number), 10) + `":  `)))
						json.NewEncoder(w).Encode(tfile)
						w.Write([]byte(`,`))

					}
				}
				if Log_lvl > 1 {
					log.Printf("[+] HIT: %s\n", path)
				}

			}

		}
		return nil
	}

}

//TODO_7: One of the options for the API is a query command
//TODO_7: Create a walkFn2 function based on the walkFn function,
//TODO_7: Instead of using the regexes array, define a single regex
//TODO_7: Hint look at the logic in scrape.go to see how to do that;
//TODO_7: You won't have to itterate through the regexes for loop in this func!

func walkFn2(w http.ResponseWriter, query string) filepath.WalkFunc {
	return func(path string, f os.FileInfo, err error) error {
		q := regexp.MustCompile(`(?i)` + query)
		if q.MatchString(path) {
			//TODO_5: As it currently stands the same file can be added to the array more than once
			//TODO_5: Prevent this from happening by checking if the file AND location already exist as a single record
			dir, filename := filepath.Split(path)
			toAdd := true
			for _, file := range Files {
				if file.Filename == string(filename) && file.Location == string(dir) {
					toAdd = false
				}
			}
			if toAdd {
				var tfile FileInfo

				tfile.Filename = string(filename)
				tfile.Location = string(dir)
				tfile.Number = files_found
				files_found += 1
				Files = append(Files, tfile)
				if w != nil && len(Files) > 0 {

					//TODO_6: The current key value is the LEN of Files (this terrible);
					//TODO_6: Create some variable to track how many files have been added
					w.Write([]byte(`"` + (strconv.FormatInt(int64(tfile.Number), 10) + `":  `)))
					json.NewEncoder(w).Encode(tfile)
					w.Write([]byte(`,`))

				}
			}
			if Log_lvl > 1 {
				log.Printf("[+] HIT: %s\n", path)
			}

		}
		return nil

	}
}

//==========================================================================\\

func APISTATUS(w http.ResponseWriter, r *http.Request) {

	if Log_lvl > 1 {
		log.Printf("Entering %s end point", r.URL.Path)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{ "status" : "API is up and running ",`))
	var regexstrings []string

	for _, regex := range regexes {
		regexstrings = append(regexstrings, regex.String())
	}

	w.Write([]byte(` "regexs" :`))
	json.NewEncoder(w).Encode(regexstrings)
	w.Write([]byte(`}`))
	if Log_lvl > 1 {
		log.Println(regexes)
	}

}

func MainPage(w http.ResponseWriter, r *http.Request) {
	if Log_lvl > 1 {
		log.Printf("Entering %s end point", r.URL.Path)
	}

	w.Header().Set("Content-Type", "text/html")

	w.WriteHeader(http.StatusOK)
	//TODO_8 - Write out something better than this that describes what this api does

	display := "<html><body><H1>Hi Mike, if you really check this</H1>" +
		"<ul>" +
		"<li><code>localhost:8080/</code> Mainepage, here </li>" +
		"<li><code>localhost:8080/indexer?location=/</code> Get the readouts for the / directory</li>" +
		"<li><code>localhost:8080/search?main.md</code> Search for s pecific file</li>" +
		"<li><code>localhost:8080/reset</code> Reset the regex</li>" +
		"<li><code>localhost:8080/clear</code> Empty the regex array </li>" +
		"<li><code>localhost:8080/addsearch?regex={pattern}</code> Adda regex string to the array</li>" +
		"</ul>" +
		"</body>"
	fmt.Fprintf(w, display)
}

func FindFile(w http.ResponseWriter, r *http.Request) {
	if Log_lvl > 1 {
		log.Printf("Entering %s end point", r.URL.Path)
	}

	q, ok := r.URL.Query()["q"]

	w.WriteHeader(http.StatusOK)
	if ok && len(q[0]) > 0 {
		if Log_lvl > 1 {
			log.Printf("Entering search with query=%s", q[0])
		}

		// ADVANCED: Create a function in scrape.go that returns a list of file locations; call and use the result here
		// e.g., func finder(query string) []string { ... }
		noFilesFound := true
		for _, File := range Files {
			if File.Filename == q[0] {
				noFilesFound = false
				json.NewEncoder(w).Encode(File.Location)
				//consider FOUND = TRUE
			}
		}
		if noFilesFound {
			//https://y.yarn.co/7708c2f6-af8e-4cd4-9fba-6621db1636bc_text.mp4
			videos := []string{"https://y.yarn.co/7708c2f6-af8e-4cd4-9fba-6621db1636bc_text.mp4", "https://y.yarn.co/7708c2f6-af8e-4cd4-9fba-6621db1636bc_text.mp4", "https://y.yarn.co/76b93a31-164b-4519-8ef6-6e30323bc8a7.mp4", "https://y.yarn.co/649f3718-74ee-4188-8ea7-fb58372cf336.mp4"}
			vidurl := videos[rand.Intn(len(videos))]
			display := "<html><body><h1>No Files Found</h1>" +
				"<video controls autoplay>" +
				"<source src=" + vidurl + " type=\"video/mp4\">" +
				"</video></body></html>"
			fmt.Fprintf(w, display)
		}
		//TODO_9: Handle when no matches exist; print a useful json response to the user; hint you might need a "FOUND variable" to check here ...

	} else {
		// didn't pass in a search term, show all that you've found
		w.Write([]byte(`"files":`))
		json.NewEncoder(w).Encode(Files)
	}
}

func IndexFiles(w http.ResponseWriter, r *http.Request) {

	if Log_lvl > 1 {
		log.Printf("Entering %s end point", r.URL.Path)
	}
	w.Header().Set("Content-Type", "application/json")

	location, locOK := r.URL.Query()["location"]
	regex, regOK := r.URL.Query()["regex"]

	//TODO_10: Currently there is a huge risk with this code ... namely, we can search from the root /
	//TODO_10: Assume the location passed starts at /home/ (or in Windows pick some "safe?" location)
	//TODO_10: something like ...  rootDir string := "???"
	//TODO_10: create another variable and append location[0] to rootDir (where appropriate) to patch this hole

	if location[0] == "/" {
		location[0] = "/home" + location[0]
	}
	if locOK && len(location[0]) > 0 {
		w.WriteHeader(http.StatusOK)

	} else {
		w.WriteHeader(http.StatusFailedDependency)
		w.Write([]byte(`{ "parameters" : {"required": "location",`))
		w.Write([]byte(`"optional": "regex"},`))
		w.Write([]byte(`"examples" : { "required": "/indexer?location=/xyz",`))
		w.Write([]byte(`"optional": "/indexer?location=/xyz&regex=(i?).md"}}`))
		return
	}

	//wrapper to make "nice json"
	w.Write([]byte(`{ `))

	// TODO_11: Currently the code DOES NOT do anything with an optionally passed regex parameter
	// Define the logic required here to call the new function walkFn2(w,regex[0])
	// Hint, you need to grab the regex parameter (see how it's done for location above...)

	// if regexOK
	//   call filepath.Walk(location[0], walkFn2(w, `(i?)`+regex[0]))
	// else run code to locate files matching stored regular expression
	if regOK && len(regex[0]) > 0 {
		if err := filepath.Walk(location[0], walkFn2(w, regex[0])); err != nil {
			log.Panicln(err)
		}
	} else {
		if err := filepath.Walk(location[0], walkFn(w)); err != nil {
			log.Panicln(err)
		}
	}

	//wrapper to make "nice json"
	w.Write([]byte(` "status": "completed"} `))

}

//TODO_12 create endpoint that calls resetRegEx AND *** clears the current Files found; ***
//TODO_12 Make sure to connect the name of your function back to the reset endpoint main.go!
func ResetRegEx(w http.ResponseWriter, r *http.Request) {
	resetRegEx()
	emptyFiles()
}

//TODO_13 create endpoint that calls clearRegEx ;
//TODO_12 Make sure to connect the name of your function back to the clear endpoint main.go!
func ClearRegEx(w http.ResponseWriter, r *http.Request) {
	clearRegEx()
}

//TODO_14 create endpoint that calls addRegEx ;
func AddRegEx(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	newGex := "(?i)" + params["regex"]
	addRegEx(newGex)
}

//TODO_14 Make sure to connect the name of your function back to the addsearch endpoint in main.go!
// consider using the mux feature
// params := mux.Vars(r)
// params["regex"] should contain your string that you pass to addRegEx
// If you try to pass in (?i) on the command line you'll likely encounter issues
// Suggestion : prepend (?i) to the search query in this endpoint
