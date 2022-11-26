package main
import (
 "encoding/json"
 "fmt"
 "io/ioutil"
 "net/http"
 "github.com/twilio/twilio-go"
 "github.com/twilio/twilio-go/rest/api/v2010"
 "os"
)

type Response struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
   }

   func main() {
	fmt.Println("Calling API...")
   client := &http.Client{}
	req, err := http.NewRequest("GET", "https://icanhazdadjoke.com/", nil)
	if err != nil {
	 fmt.Print(err.Error())
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
	 fmt.Print(err.Error())
	}
   defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	 fmt.Print(err.Error())
	}
   var responseObject Response
	json.Unmarshal(bodyBytes, &responseObject)
	fmt.Printf("API Response as struct %+v\n", responseObject)
	fmt.Printf(responseObject.Joke)

    
   // call SMS for jason
   sms_jason(responseObject.Joke)
   // call SMS for carson
   sms_carson(responseObject.Joke)
   // call SMS for heather
   sms_heather(responseObject.Joke)
   }

   func sms_jason(jokemsg string) {
	// SMS stuff
	client := twilio.NewRestClient()
    params := &openapi.CreateMessageParams{}
    params.SetTo(os.Getenv("TO_PHONE_NUMBER_JASON"))
    params.SetFrom(os.Getenv("TWILIO_PHONE_NUMBER"))
    params.SetBody(jokemsg)
    _, err := client.Api.CreateMessage(params)
    if err != nil {
        fmt.Println(err.Error())
    } else {
        fmt.Println("SMS sent successfully to JASON!")
    }
    }

	func sms_carson(jokemsg string) {
	// SMS stuff
	  client := twilio.NewRestClient()
	  params := &openapi.CreateMessageParams{}
	  params.SetTo(os.Getenv("TO_PHONE_NUMBER_CARSON"))
	  params.SetFrom(os.Getenv("TWILIO_PHONE_NUMBER"))
	  params.SetBody(jokemsg)
	
		_, err := client.Api.CreateMessage(params)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("SMS sent successfully to CARSON!")
		}
	}
   
	func sms_heather(jokemsg string) {
		// SMS stuff
		  client := twilio.NewRestClient()
		  params := &openapi.CreateMessageParams{}
		  params.SetTo(os.Getenv("TO_PHONE_NUMBER_HEATHER"))
		  params.SetFrom(os.Getenv("TWILIO_PHONE_NUMBER"))
		  params.SetBody(jokemsg)
		
			_, err := client.Api.CreateMessage(params)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("SMS sent successfully to HEATHER!")
			}
		}