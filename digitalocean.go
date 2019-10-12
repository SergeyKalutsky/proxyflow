package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/digitalocean/godo"
	"golang.org/x/oauth2"
)

func readToken() string {
	file, err := os.Open("token.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	tocken := string(b)
	return tocken
}

// TokenSource is a fucking bullshit
type TokenSource struct {
	AccessToken string
}

// Token is a bullshit.
func (t *TokenSource) Token() (*oauth2.Token, error) {
	token := &oauth2.Token{
		AccessToken: t.AccessToken,
	}
	return token, nil
}

func main() {
	token := readToken()
	tokenSource := &TokenSource{
		AccessToken: token,
	}

	oauthClient := oauth2.NewClient(context.Background(), tokenSource)
	client := godo.NewClient(oauthClient)
	fmt.Println(client)

}
