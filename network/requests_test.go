package network

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/qdm12/golibs/helpers"
	"github.com/qdm12/golibs/security"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_NewClient(t *testing.T) {
	t.Parallel()
	c := NewClient(time.Second)
	assert.NotNil(t, c)
}

func Test_DoHTTPRequest(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		response  *http.Response
		clientErr error
		readBody  func(r io.Reader) ([]byte, error)
		status    int
		content   []byte
		err       error
	}{
		"no error": {
			&http.Response{
				Body:       ioutil.NopCloser(bytes.NewBufferString("body")),
				StatusCode: 200}, nil, ioutil.ReadAll,
			200, []byte("body"), nil},
		"http error": {
			nil, fmt.Errorf("error"), ioutil.ReadAll,
			0, nil, fmt.Errorf("error")},
		"body read error": {
			&http.Response{
				Body:       ioutil.NopCloser(bytes.NewBufferString("body")),
				StatusCode: 200},
			nil,
			func(r io.Reader) ([]byte, error) {
				return nil, fmt.Errorf("error")
			},
			0, nil, fmt.Errorf("error")},
	}
	for name, tc := range tests {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			mockHTTPClient := &mockHttpClient{}
			mockHTTPClient.On("Do", mock.Anything).
				Return(tc.response, tc.clientErr).Once()
			c := &ClientImpl{
				httpClient: mockHTTPClient,
				readBody:   tc.readBody,
			}
			status, content, err := c.DoHTTPRequest(nil)
			helpers.AssertErrorsEqual(t, tc.err, err)
			assert.Equal(t, tc.status, status)
			assert.Equal(t, tc.content, content)
			mockHTTPClient.AssertExpectations(t)
		})
	}
}

func Test_GetContent(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		URL       string
		setters   []GetContentSetter
		response  *http.Response
		clientErr error
		content   []byte
		status    int
		err       error
	}{
		"no error": {
			"https://domain.com",
			nil,
			&http.Response{
				Body:       ioutil.NopCloser(bytes.NewBufferString("body")),
				StatusCode: 200},
			nil,
			[]byte("body"), 200, nil},
		"http error": {
			"https://domain.com",
			nil,
			nil,
			fmt.Errorf("error"),
			nil, 0, fmt.Errorf("cannot GET content of URL https://domain.com: error")},
		"bad URL error": {
			"\n",
			nil,
			nil,
			nil,
			nil, 0, fmt.Errorf("cannot GET content of URL \n: parse \n: net/url: invalid control character in URL")},
		"set random user agent": {
			"https://domain.com",
			[]GetContentSetter{UseRandomUserAgent()},
			&http.Response{
				Body:       ioutil.NopCloser(bytes.NewBufferString("body")),
				StatusCode: 200},
			nil,
			[]byte("body"), 200, nil},
	}
	for name, tc := range tests {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			mockHTTPClient := &mockHttpClient{}
			if tc.clientErr != nil || (tc.response != nil && tc.response.Body != nil) {
				mockHTTPClient.On("Do", mock.Anything).
					Return(tc.response, tc.clientErr).Once()
			}
			c := &ClientImpl{
				httpClient: mockHTTPClient,
				readBody:   ioutil.ReadAll,
				userAgents: []string{"abc"},
				random:     security.NewRandom(),
			}
			content, status, err := c.GetContent(tc.URL, tc.setters...)
			helpers.AssertErrorsEqual(t, tc.err, err)
			assert.Equal(t, tc.status, status)
			assert.Equal(t, tc.content, content)
			mockHTTPClient.AssertExpectations(t)
		})
	}
}