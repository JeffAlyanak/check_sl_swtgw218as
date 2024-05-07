package main

import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/jeffalyanak/check_sl_swtgw218as/check"
)

func main() {
	host := flag.String("h", "", "IP or Fully-qualified domain name to check.")
	username := flag.String("u", "", "Username")
	password := flag.String("p", "", "Password")

	flag.Parse()
	if *host == "" {
		fmt.Println("Please provide an IP or fully-qualified domain name.")
		os.Exit(3)
	}
	if *username == "" {
		fmt.Println("Please provide a username.")
		os.Exit(3)
	}
	if *password == "" {
		fmt.Println("Please provide a password.")
		os.Exit(3)
	}

	// Create mew Performance Data struct and start the timer
	var perfData check.PerfData
	perfData.StartTimer(time.Now())

	// URL, form, and cookie params
	baseURL := "http://" + *host + "/port.cgi?page=stats"
	params := url.Values{}
	params.Set("page", "stats")
	formParams := url.Values{}
	formParams.Set("username", *username)
	formParams.Set("password", *password)
	formParams.Set("language", "EN")
	formParams.Set("Response", getMD5Hash(*username+*password))
	cookie := &http.Cookie{Name: "admin", Value: getMD5Hash(*username + *password)}

	// Create a new GET request with the parameters
	req, err := http.NewRequest("GET", baseURL, strings.NewReader(formParams.Encode()))
	if err != nil {
		log.Println("Error creating request:", err)
	}
	req.AddCookie(cookie)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.URL.RawQuery = params.Encode()

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error:", err)
		os.Exit(3)
	}
	document, err := goquery.NewDocumentFromReader(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Println("Error:", err)
		os.Exit(3)
	}

	// Parse the data
	totalBadPackets := 0
	document.Find("table tr").Each(func(i int, s *goquery.Selection) {
		if i != 0 {
			s.Find("td").Each(func(j int, td *goquery.Selection) {
				switch j {
				case 3:
					perfData.Add("TX Good Port "+fmt.Sprint(i), td.Text(), "c")
				case 4:
					packets, err := strconv.Atoi(td.Text())
					if err != nil {
						fmt.Println("Error:", err.Error())
						os.Exit(3)
					}
					totalBadPackets += packets
					perfData.Add("TX Errors Port "+fmt.Sprint(i), td.Text(), "c")
				case 5:
					perfData.Add("RX Good Port "+fmt.Sprint(i), td.Text(), "c")
				case 6:
					packets, err := strconv.Atoi(td.Text())
					if err != nil {
						fmt.Println("Error:", err.Error())
						os.Exit(3)
					}
					totalBadPackets += packets
					perfData.Add("RX Errors Port "+fmt.Sprint(i), td.Text(), "c")
				}
			})
		}
	})

	if totalBadPackets > 0 {
		printIntro("WARNING", fmt.Sprint((totalBadPackets)))
	} else {
		printIntro("OK", fmt.Sprint((totalBadPackets)))
	}
	fmt.Println(fmt.Sprint(perfData.Get()))
}

func printIntro(issue string, url string) {
	fmt.Print(issue + " - total packet errors: " + url)
}

// Helper function to calculate MD5 hash
func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
