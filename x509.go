package mailgun

import (
	"context"
	"fmt"
	"regexp"
)

type GetX509Response struct {
	Status      string `json:"status"`
	Error       string `json:"error"`
	Certificate string `json:"certificate"`
}

type RegenerateX509Response struct {
	Message  string `json:"message"`
	Location string `json:"location"`
}

type InitiateX509X509Response struct {
	Message  string `json:"message"`
	Location string `json:"location"`
}

var (
	apiBaseVersion = regexp.MustCompile(`/v[0-9]+$`)
)

func getVersionBase(base, version string) string {
	return apiBaseVersion.ReplaceAllString(base, version)
}
func generateVersionPublicApiUrl(mg Mailgun, version, endpoint string) string {
	return fmt.Sprintf("%s/%s", getVersionBase(mg.APIBase(), "/"+version), endpoint)
}

func (mg *MailgunImpl) GetX509Status(ctx context.Context, domain string) (*GetX509Response, error) {
	r := newHTTPRequest(generateVersionPublicApiUrl(mg, "v2", x509Endpoint) + "/" + domain + "/status")
	r.setClient(mg.Client())
	r.setBasicAuth(basicAuthUser, mg.APIKey())
	var resp = &GetX509Response{}
	err := getResponseFromJSON(ctx, r, resp)
	return resp, err
}
func (mg *MailgunImpl) RegenerateX509(ctx context.Context, domain string) (*RegenerateX509Response, error) {
	r := newHTTPRequest(generateVersionPublicApiUrl(mg, "v2", x509Endpoint) + "/" + domain)
	r.setClient(mg.Client())
	r.setBasicAuth(basicAuthUser, mg.APIKey())
	var resp = &RegenerateX509Response{}
	err := putResponseFromJSON(ctx, r, nil, resp)
	return resp, err
}
func (mg *MailgunImpl) InitiateX509(ctx context.Context, domain string) (*InitiateX509X509Response, error) {
	r := newHTTPRequest(generateVersionPublicApiUrl(mg, "v2", x509Endpoint) + "/" + domain)
	r.setClient(mg.Client())
	r.setBasicAuth(basicAuthUser, mg.APIKey())
	var resp = &InitiateX509X509Response{}
	err := postResponseFromJSON(ctx, r, nil, resp)
	return resp, err
}
