package auth

import (
	"net/http"
	"testing"
)



func TestGetApiKey(t *testing.T){
	t.Run("ValidAuthorizationHeader", func(t *testing.T){
		headers := http.Header{}
		headers.Set("Authorization", "ApiKey 12345")
        apiKey, err:=GetAPIKey(headers)
		if err != nil{
			t.Errorf("Expected no error, but got %v", err)
		}
		if apiKey!="12345"{
			t.Errorf("Expected apiKey to be '12345', but got %s", apiKey)
		}
	})
	//Test case with no auth header
	t.Run("NoAuthorizationHeader", func (t *testing.T)  {
		headers := http.Header{}
		_, err := GetAPIKey(headers)
		if err==nil {
			t.Error("Expected an error, but got none")
		}
		
	})
	// Test case with malformed auth header
	t.Run("MalformedAuthorizationHeader", func (t *testing.T)  {
		headers := http.Header{}
		headers.Set("Authorization", "InvlaidFormat")
		_, err:=GetAPIKey(headers)
		if err == nil {
			t.Error("Expected an error, but got none")
		}
	})
}
