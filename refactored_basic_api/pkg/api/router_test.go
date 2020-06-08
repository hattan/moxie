package api

import (
	"encoding/json"
	"testing"

	"github.com/google/uuid"
	m "github.com/hattan/moxie/pkg/api/models"
	"github.com/stretchr/testify/assert"
)

var testArticles = m.Articles{
	{Id: uuid.New(), Title: "Hello", Desc: "Fake Article", Content: "woow"},
}

var apiTests = []APITestCase{
	{
		Tag:              "t1 - HomePage",
		Method:           "GET",
		URL:              "/",
		Status:           200,
		ExpectedResponse: "Welcome to the HomePage!",
	},
	{
		Tag:          "t2 - Get All Articles",
		Method:       "GET",
		URL:          "/articles",
		Status:       200,
		ExpectedJSON: getStringFromJson(testArticles),
	},
	{
		Tag:          "t3 - Get Article By Id",
		Method:       "GET",
		URL:          "/article/" + testArticles[0].Id.String(),
		Status:       200,
		ExpectedJSON: getStringFromJson(testArticles[0]),
	},
	{
		Tag:    "t3 - Get Article By Invalid UUID throws 500 error",
		Method: "GET",
		URL:    "/article/badid",
		Status: 500,
	},
	{
		Tag:    "t5 - Get Article By new UUID retuns 404",
		Method: "GET",
		URL:    "/article/" + uuid.New().String(),
		Status: 404,
	},
}

func TestRouter(t *testing.T) {
	SetUpTestFixtures(testArticles)
	router := NewRouter()

	for _, test := range apiTests {
		res := testAPI(router, test.Method, test.URL, test.Tag)
		assert.Equal(t, test.Status, res.Code, test.Tag)
		if test.ExpectedResponse != "" {
			assert.Equal(t, test.ExpectedResponse, res.Body.String(), test.Tag)
		}
		if test.ExpectedJSON != "" {
			body := res.Body.String()
			assert.JSONEq(t, test.ExpectedJSON, body, test.Tag)
		}
	}
}

func getStringFromJson(input interface{}) string {
	expected, _ := json.Marshal(input)
	return string(expected)
}
