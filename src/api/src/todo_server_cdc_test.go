package main

import (
	"log"
	"net/http"
	"testing"

	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
)

func TestProvider(t *testing.T) {

	pact := &dsl.Pact{
	  Provider: "TodoProvider",
	}
	log.Fatal(http.ListenAndServe(":9090", http.HandlerFunc(TodoServer)))
  
	pact.VerifyProvider(t, types.VerifyRequest{
	  ProviderBaseURL:        "http://localhost:9090",
	  PactURLs:        []string{"../../client/pacts/todo-consumer-todo-provider.json"},
	})
  }
