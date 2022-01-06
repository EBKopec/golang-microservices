package restclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var (
	enabledMocks = false
	mocks        = make(map[string]*Mock)
)

type Mock struct {
	Url        string
	HttpMethod string
	Response   *http.Response
	BodyText   string
	Err        error
}

func (m *Mock) GetResponse() *http.Response {
	m.Response.Body = ioutil.NopCloser(strings.NewReader(m.BodyText))
	return m.Response
}

func GetMockId(httpMethod string, url string) string {
	return fmt.Sprintf("%s_%s", httpMethod, url)
}

func StartMockups() {
	enabledMocks = true
}

func FlushMocks() {
	mocks = make(map[string]*Mock)
}

func StopMockups() {
	enabledMocks = false
}

func AddMockups(mock Mock) {
	mocks[GetMockId(mock.HttpMethod, mock.Url)] = &mock
}

func Post(url string, body interface{}, headers http.Header) (*http.Response, error) {
	if enabledMocks {
		mock := mocks[GetMockId(http.MethodPost, url)]
		if mock == nil {
			return nil, errors.New("no mockup found for give request")
		}
		return mock.GetResponse(), mock.Err
	}

	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonBytes))
	request.Header = headers

	client := http.Client{}
	return client.Do(request)
}
