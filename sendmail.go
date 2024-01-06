package main

import (
    "flag"
    "fmt"
    "log"
	"os"
    "github.com/joho/godotenv"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
    
)

func main() {
    file, err := os.OpenFile("/usr/share/sendmail/sendmail/sendmail.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    log.Println("Logging started")
    // read the sendmail.env file
    err2 := godotenv.Load("/usr/share/sendmail/sendmail/sendmail.env")
    if err2 != nil {
        log.Fatal("Error loading .env file")
    }
    // Define command-line flags
    etype := flag.String("etype", "com", "Mandatory, choices are com or esti.")
    msgid := flag.String("mgsid", "", "This is comid or estid")
    name := flag.String("name", "", "The name to greet")
    address := flag.String("address", "", "The customers physical address")
    city := flag.String("city", "", "The customers city")
    phone := flag.String("phone", "", "The customers phone number")
    email := flag.String("email", "", "The customers email address")
    comment := flag.String("comment", "", "The customers comment")
    intake := flag.String("intake", "", "The intake date")
    reqdate := flag.String("reqdate", "", "The requested date")
    

    flag.Parse()

    log.Printf("etype: %s\n", *etype)
    log.Printf("msgid: %s\n", *msgid)
    log.Printf("name: %s\n", *name)
    log.Printf("address: %s\n", *address)
    log.Printf("city: %s\n", *city)
    log.Printf("phone: %s\n", *phone)
    log.Printf("email: %s\n", *email)
    log.Printf("comment: %s\n", *comment)
    log.Printf("intake: %s\n", *intake)
    log.Printf("reqdate: %s\n", *reqdate)


    if *etype == "com" {
        subJECT := "atsbot: " + *name + " has left a new comment"

        h0 := "<div>"
        h1 := "<p>Name: " + *name + "</p>"
        h5 := "<p>Email: " + *email + "</p>"
        h6 := "<p>Comment: " + *comment + "</p>"
        h9 := "</div>"
        h10 := "<a href='http://192.168.0.91:8181/accept/" + *msgid
        h11 := "'><button style='background-color:green;color:white;border-radius:8px;margin:8px;'>Accept</button></a>"
        h12 := "<a href='http://192.168.0.94:8181/reject/" + *msgid
        h13 := "'><button style='background-color:red;color:white;border-radius:8px;margin:8px;'>Reject</button></a>"
        html_str := h0 + h1 + h5 + h6 + h9 + h10 + h11 + h12 + h13

        from := mail.NewEmail("atsbot", "porthose.cjsmo.cjsmo@gmail.com")
        subject := subJECT
        to := mail.NewEmail("Charlie", "porthose.cjsmo.cjsmo@gmail.com")
        plainTextContent := ""
        htmlContent := html_str
        message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
        client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
        response, err := client.Send(message)
        if err != nil {
            log.Println(err)
        } else {
            fmt.Println(response.StatusCode)
            log.Println(response.StatusCode)
            fmt.Println(response.Body)
            log.Println(response.Body)
            fmt.Println(response.Headers)
            log.Println(response.Headers)
        }
    } else {
        subJECT := "atsbot: " + *name + " has left a new comment"
        h0 := "<div>"
        h1 := "<p>Name: " + *name + "</p>"
        h2 := "<p>Address: " + *address + "</p>"
        h3 := "<p>City: " + *city + "</p>"
        h4 := "<p>Phone: " + *phone + "</p>"
        h5 := "<p>Email: " + *email + "</p>"
        h6 := "<p>Comment: " + *comment + "</p>"
        h7 := "<p>Intake: " + *intake + "</p>"
        h8 := "<p>ReqDate: " + *reqdate + "</p>"
        h9 := "</div>"
        h10 := "<a href='http://192.168.0.91:8181/accept/" + *msgid
        h11 := "'><button style='background-color:blue;border-radius:8px;color:white;'>Completed</button></a>"
        html_str := h0 + h1 + h2 + h3 + h4 + h5 + h6 + h7 + h8 + h9 + h10 + h11
        from := mail.NewEmail("atsbot", "porthose.cjsmo.cjsmo@gmail.com")
        subject := subJECT
        to := mail.NewEmail("Charlie", "porthose.cjsmo.cjsmo@gmail.com")
        plainTextContent := ""
        htmlContent := html_str
        message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
        client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
        response, err := client.Send(message)
        if err != nil {
            log.Println(err)
        } else {
            fmt.Println(response.StatusCode)
            log.Println(response.StatusCode)
            fmt.Println(response.Body)
            log.Println(response.Body)
            fmt.Println(response.Headers)
            log.Println(response.Headers)
        }
    }

    
}
