package handlers

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/gorilla/mux"
	"github.com/vermavashish/try/data"
)

// swagger:route GET /products products listProducts
// Return a list of products from the database
// responses:
//	200: productsResponse

// getProducts returns all the products from the data source
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")
	rw.Header().Add("Content-Type", "application/json")

	// fetch products drom the data source
	lp := data.GetProducts()

	// serialize the list to JSON
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

// swagger:route GET /products/{id} products listSingle
// Return a list of products from the database
// responses:
//	200: productsResponse
//	404: errorResponse

// ListSingle handles GET requests
func (p *Products) ListSingle(rw http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to covert id", http.StatusBadRequest)
		return
	}

	p.l.Println("[DEBUG] get record id", id)

	prod, _, err := data.FindProduct(id)

	switch err {
	case nil:

	case data.ErrProductNotFound:
		p.l.Println("[ERROR] fetching product", err)

		rw.WriteHeader(http.StatusNotFound)
		return
	default:
		p.l.Println("[ERROR] fetching product", err)

		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = prod.ToJSON(rw)
	if err != nil {
		// we should never be here but log the error just incase
		p.l.Println("[ERROR] serializing product", err)
	}
}

func (p *Products) Tryaws(rw http.ResponseWriter, r *http.Request) {

	emailIDPtr := flag.String("e", "testthree@gmail.com", "The email address of the user")
	userPoolIDPtr := flag.String("p", "us-east-2_S346G935R", "The ID of the user pool")
	userNamePtr := flag.String("n", "testthree", "The name of the user")
	// passwordPtr := flag.String("w", "Pass345#", "The password of the user")

	flag.Parse()

	if *emailIDPtr == "" || *userPoolIDPtr == "" || *userNamePtr == "" {
		fmt.Println("You must supply an email address, user pool ID, and user name")
		fmt.Println("Usage: go run CreateUser.go -e EMAIL-ADDRESS -p USER-POOL-ID -n USER-NAME")
		os.Exit(1)
	}

	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	cognitoClient := cognitoidentityprovider.New(sess)

	newUserData := &cognitoidentityprovider.AdminCreateUserInput{
		DesiredDeliveryMediums: []*string{
			aws.String("EMAIL"),
		},
		UserAttributes: []*cognitoidentityprovider.AttributeType{
			{
				Name:  aws.String("email"),
				Value: aws.String(*emailIDPtr),
			},
		},
	}

	newUserData.SetUserPoolId(*userPoolIDPtr)
	newUserData.SetUsername(*userNamePtr)

	newUser, err := cognitoClient.AdminCreateUser(newUserData)
	if err != nil {
		fmt.Println("Got error creating user:", err)
	}
	p.l.Println(newUser)
}

func (p *Products) AuthAws(rw http.ResponseWriter, r *http.Request) {

	emailIDPtr := flag.String("e", "testtwo@gmail.com", "The email address of the user")
	userPoolIDPtr := flag.String("p", "us-east-2_S346G935R", "The ID of the user pool")
	userNamePtr := flag.String("n", "testtwo", "The name of the user")

	flag.Parse()

	if *emailIDPtr == "" || *userPoolIDPtr == "" || *userNamePtr == "" {
		fmt.Println("You must supply an email address, user pool ID, and user name")
		fmt.Println("Usage: go run CreateUser.go -e EMAIL-ADDRESS -p USER-POOL-ID -n USER-NAME")
		os.Exit(1)
	}

	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	cognitoClient := cognitoidentityprovider.New(sess)

	newUserData := &cognitoidentityprovider.AdminCreateUserInput{
		DesiredDeliveryMediums: []*string{
			aws.String("EMAIL"),
		},
		UserAttributes: []*cognitoidentityprovider.AttributeType{
			{
				Name:  aws.String("email"),
				Value: aws.String(*emailIDPtr),
			},
		},
	}

	newUserData.SetUserPoolId(*userPoolIDPtr)
	newUserData.SetUsername(*userNamePtr)

	newUser, err := cognitoClient.AdminCreateUser(newUserData)
	if err != nil {
		fmt.Println("Got error creating user:", err)
	}

	p.l.Println(newUser)

}
