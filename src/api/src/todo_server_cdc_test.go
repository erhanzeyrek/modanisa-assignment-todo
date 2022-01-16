package main

import (
	"testing"
	"todo-api/app"

	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
)

func TestProvider(t *testing.T) {

	pact := &dsl.Pact{
	  Provider: "TodoProvider",
	}
	
	go app.StartApp()
  
	pact.VerifyProvider(t, types.VerifyRequest{
	  ProviderBaseURL:        "http://localhost:9090",
	  PactURLs:        []string{"../../client/pacts/todo-consumer-todo-provider.json"},
	})
  }
