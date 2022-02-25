package shodan

import (
	"fmt"
	"net/http"
)

//Sorry Mike, was kind of busy this week so this is a very cop-out of an assignment submission
func getFacets() {
	facetsURL, err := http.Get("https://raw.githubusercontent.com/darkoperator/Posh-Shodan/master/en-US/about_Shodan_Host_Search_Facets.help.txt")
	if err != nil {
		panic(err)
	}
	filtersURL, err := http.Get("https://raw.githubusercontent.com/darkoperator/Posh-Shodan/master/en-US/about_Shodan_Host_Search_Filters.help.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(facetsURL)
	fmt.Println(filtersURL)

}
