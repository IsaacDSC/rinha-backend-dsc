package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"testing"

	"github.com/IsaacDSC/rinha-backend-dsc/api"
	"github.com/IsaacDSC/rinha-backend-dsc/shared/dto"

	"github.com/docker/distribution/uuid"
	"github.com/stretchr/testify/assert"
)

type InputDTO struct {
	Apelido    string   `json:"apelido"`
	Nome       string   `json:"nome"`
	Nascimento string   `json:"nascimento"`
	Stack      []string `json:"stack"`
}

var PersonID string

func init() {
	go api.StartAPI()
}
func TestE2E(t *testing.T) {
	t.Run("Should be create an user", func(t *testing.T) {
		url := "http://localhost:3000/pessoas"
		input := InputDTO{
			Apelido:    uuid.Generate().String()[0:32],
			Nome:       "José Roberto21",
			Nascimento: "2000-10-01",
			Stack:      []string{"C#", "Node", "Oracle"},
		}
		payload, err := json.Marshal(input)
		assert.NoError(t, err)
		reader := bytes.NewReader(payload)
		req, _ := http.NewRequest("POST", url, reader)
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("User-Agent", "Insomnia/2023.5.7")
		res, err := http.DefaultClient.Do(req)
		assert.NoError(t, err)
		defer res.Body.Close()
		assert.Equal(t, 201, res.StatusCode)
		assert.NotNil(t, res.Header.Get("Location"))
		PersonID = strings.Split(res.Header.Get("Location"), "/")[2]
	})

	t.Run("Should be create an user", func(t *testing.T) {
		url := "http://localhost:3000/pessoas"
		input := InputDTO{
			Apelido:    uuid.Generate().String()[0:32],
			Nome:       "Ana Barbosa12",
			Nascimento: "1985-09-23",
			Stack:      nil,
		}
		payload, err := json.Marshal(input)
		assert.NoError(t, err)
		reader := bytes.NewReader(payload)
		req, _ := http.NewRequest("POST", url, reader)
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("User-Agent", "Insomnia/2023.5.7")
		res, err := http.DefaultClient.Do(req)
		assert.NoError(t, err)
		defer res.Body.Close()
		assert.Equal(t, 201, res.StatusCode)
		assert.NotNil(t, res.Header.Get("Location"))
	})

	t.Run("Should be retrieve an user by id", func(t *testing.T) {
		url := fmt.Sprintf("http://localhost:3000/pessoas/%s", PersonID)
		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("User-Agent", "Insomnia/2023.5.7")
		res, err := http.DefaultClient.Do(req)
		assert.NoError(t, err)
		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		assert.NoError(t, err)
		assert.Equal(t, 200, res.StatusCode)
		assert.NotNil(t, res.Header.Get("Location"))
		var output dto.Person
		err = json.Unmarshal(body, &output)
		assert.NoError(t, err)
		assert.True(t, len(output.ID) > 0)
		assert.True(t, len(output.LastName) > 0)
		assert.True(t, len(output.Birthday) > 0)
		assert.True(t, len(output.Name) > 0)
		assert.True(t, len(output.Stack) > 0)
	})

	t.Run("Should be get person by stack", func(t *testing.T) {
		url := "http://localhost:3000/pessoas?stack=C%23"
		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("User-Agent", "Insomnia/2023.5.7")
		res, err := http.DefaultClient.Do(req)
		assert.NoError(t, err)
		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		assert.NoError(t, err)
		assert.Equal(t, 200, res.StatusCode)
		fmt.Println(res.Header.Get("Location"))
		assert.NotNil(t, res.Header.Get("Location"))
		var output []dto.Person
		err = json.Unmarshal(body, &output)
		assert.NoError(t, err)
		assert.True(t, len(output) > 0)
		assert.True(t, len(output[0].ID) > 0)
		assert.True(t, len(output[0].LastName) > 0)
		assert.True(t, len(output[0].Birthday) > 0)
		assert.True(t, len(output[0].Name) > 0)
		assert.True(t, len(output[0].Stack) > 0)
	})

	t.Run("Should be counter persons registered", func(t *testing.T) {
		url := "http://localhost:3000/contagem-pessoas"
		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("User-Agent", "Insomnia/2023.5.7")
		res, err := http.DefaultClient.Do(req)
		assert.NoError(t, err)
		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		assert.NoError(t, err)
		assert.Equal(t, 200, res.StatusCode)
		assert.NotNil(t, res.Header.Get("Location"))
		var output map[string]string
		err = json.Unmarshal(body, &output)
		assert.NoError(t, err)
		total, err := strconv.Atoi(output["total"])
		assert.NoError(t, err)
		assert.True(t, total > 1)
	})
}
