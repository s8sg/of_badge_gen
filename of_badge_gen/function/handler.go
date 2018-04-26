package function

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const (
	ImagePath = "/image/%s_%s.svg"
)

var (
	validateCustomers = "false"
	customersURL      = ""
	user              = ""
	repo              = ""
)

// Handle a serverless request
func Handle(req []byte) string {
	m, perr := url.ParseQuery(os.Getenv("Http_Query"))
	if perr != nil {
		log.Fatalf("failed to query, error : %s", perr.Error())
	}
	badge, err := getBadge(m)
	if err != nil {
		log.Fatalf(err.Error())
	}
	return string(badge)
}

func init() {
	validateCustomers = os.Getenv("validate_customers")
	if validateCustomers == "" {
		validateCustomers = "false"
	}
	customersURL = os.Getenv("customers_url")
}

func getBadge(query url.Values) ([]byte, error) {
	// https://0341c281.ngrok.io/function/s8sg-of_badge_gen?user=s8sg&repo=regex_go&branch=master

	user := query.Get("user")
	repo := query.Get("repo")
	branch := query.Get("branch")
	if branch == "" {
		branch = "master"
	}

	if validateCustomers == "true" && validateUser(user) == false {
		return nil, fmt.Errorf("failed to valicate customer: %s", user)
	}

	commitStatus, cerr := getCommitStatus(user, repo, branch)
	if cerr != nil {
		return nil, fmt.Errorf("failed to get commit status, error: %s", cerr.Error())
	}

	imagePath := fmt.Sprintf(ImagePath, commitStatus.State, commitStatus.Statuses[0].Context)
	imageContent, ierr := ioutil.ReadFile(imagePath)
	if ierr != nil {
		return nil, fmt.Errorf("failed to read image file %s, error: %s", imagePath, ierr.Error())
	}
	return imageContent, nil
}

func validateUser(user string) bool {
	found := false
	customers, getErr := getCustomers(customersURL)
	if getErr != nil {
		fmt.Println("failed to get customer list, error :", getErr.Error())
		return false
	}

	for _, customer := range customers {
		if customer == user {
			found = true
		}
	}
	return found
}

// getCustomers reads a list of customers separated by new lines
// who are valid users of OpenFaaS cloud
func getCustomers(customerURL string) ([]string, error) {
	customers := []string{}

	c := http.Client{}

	httpReq, _ := http.NewRequest(http.MethodGet, customerURL, nil)
	res, reqErr := c.Do(httpReq)

	if reqErr != nil {
		return customers, reqErr
	}

	if res.Body != nil {
		defer res.Body.Close()

		pageBody, _ := ioutil.ReadAll(res.Body)
		customers = strings.Split(string(pageBody), "\n")
	}

	return customers, nil
}
