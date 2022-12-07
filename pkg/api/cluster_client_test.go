package api

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models"
	"github.com/h2non/gock"
	. "github.com/onsi/gomega"
)

const (
	testAPIURL = "https://localhost"
)

func init() {

}
func TestConnectionString(t *testing.T) {
	RegisterTestingT(t)
	defer gock.Off()
	client := NewClusterClient(API{BaseURL: testAPIURL, Token: "TOKEN"})

	var cases = []struct {
		id       string
		connInfo *models.ClusterConnection
		code     int
		err      error
	}{
		{
			id: "some-id",
			connInfo: &models.ClusterConnection{
				PgUri: "postgresql://something",
			},
			code: 200,
			err:  nil,
		},
		{
			id:       "some-id",
			connInfo: &models.ClusterConnection{},
			code:     404,
			err:      errors.New("resource Not Found"),
		},
	}

	for _, test_case := range cases {
		gock.New(testAPIURL).
			Get(fmt.Sprintf("/clusters/%s/connection", test_case.id)).
			Reply(test_case.code).
			JSON(struct {
				Data models.ClusterConnection
			}{
				Data: *test_case.connInfo,
			})

		info, err := client.ConnectionString(context.Background(), test_case.id)
		if test_case.err == nil {
			Expect(err).To(BeNil())
		} else {
			Expect(err).To(BeEquivalentTo(err))
		}

		Expect(info).To(BeEquivalentTo(test_case.connInfo))
		gock.CleanUnmatchedRequest()
	}

}
