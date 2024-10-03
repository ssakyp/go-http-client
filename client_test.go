package client

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestClientCanHitAPI(t *testing.T) {
	t.Run("happy path - can hit the api and return a pokemon", func(t *testing.T) {
		myClient := NewClient()
		poke, err := myClient.GetPokemonByName(context.Background(), "pikachu")
		assert.NoError(t, err)
		assert.Equal(t, "pikachu", poke.Name)
	})

	t.Run("sad path - when the pokemon does not exist", func(t *testing.T) {
		myClient := NewClient()
		_, err := myClient.GetPokemonByName(context.Background(), "non-existing-pokemon")
		assert.Error(t, err)
		assert.Equal(t, PokemonFetchErr{
			Message:    "non-200 status code from the API",
			StatusCode: 404,
		}, err)
	})

	t.Run("happy path - testing the WithAPIURL option function", func(t *testing.T) {
		myClient := NewClient(
			WithAPIURL("my-test-url"),
		)
		assert.Equal(t, "my-test-url", myClient.apiURL)
	})

	t.Run("happy path - tests with httpclient works", func(t *testing.T) {
		myClient := NewClient(
			WithAPIURL("my-test-url"),
			WithHTTPClient(&http.Client{
				Timeout: 1 * time.Second,
			}),
		)
		assert.Equal(t, "my-test-url", myClient.apiURL)
		assert.Equal(t, 1*time.Second, myClient.httpClient.Timeout)

	})

	t.Run("happy test - able to hit locally running test server", func(t *testing.T) {
		ts := httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, `{"name": "pikachu", "height": 10}`)
			}),
		)
		defer ts.Close()

		myClient := NewClient(
			WithAPIURL(ts.URL),
		)

		poke, err := myClient.GetPokemonByName(context.Background(), "pikachu")
		assert.NoError(t, err)
		assert.Equal(t, 10, poke.Height)
	})

	t.Run("sad path test - able to handle 500 status from the API", func(t *testing.T) {
		ts := httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusInternalServerError)
			}),
		)
		defer ts.Close()

		myClient := NewClient(
			WithAPIURL(ts.URL),
		)
		poke, err := myClient.GetPokemonByName(context.Background(), "pikachu")
		assert.Error(t, err)
		assert.Equal(t, 0, poke.Height)
	})
}
