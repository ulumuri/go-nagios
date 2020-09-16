package integration

import (
	"net/http"
	"testing"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/ulumuri/go-nagios/nagios"
)

const (
	success      = "Success"
	dumpResponse = true
)

func TestArchive(t *testing.T) {
	if len(testInstanceAddress) == 0 {
		t.Skip()
	}

	c, err := nagios.NewClient(http.DefaultClient, testInstanceAddress)
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}

	now := time.Now()
	then := now.Add(-24 * time.Hour)

	t.Run("blank alert count", func(t *testing.T) {
		req := nagios.AlertCountRequest{
			GeneralAlertRequest: nagios.GeneralAlertRequest{
				FormatOptions: nagios.FormatOptions{
					Enumerate: true,
				},
				StartTime: then.Unix(),
				EndTime:   now.Unix(),
			},
		}

		var count nagios.AlertCount

		if err := c.Query(req, &count); err != nil {
			t.Errorf("Query: %v", err)
		}

		if count.Result.TypeText != success {
			t.Errorf("TypeText != %s", success)
		}

		if dumpResponse {
			spew.Dump(count)
		}
	})

	t.Run("blank alert count with options switched", func(t *testing.T) {
		req := nagios.AlertCountRequest{
			GeneralAlertRequest: nagios.GeneralAlertRequest{
				FormatOptions: nagios.FormatOptions{
					Whitespace: true,
					Enumerate:  true,
					Bitmask:    true,
					Duration:   true,
				},
				ObjectTypes: nagios.ObjectTypes{
					Host:    true,
					Service: true,
				},
				StateTypes: nagios.StateTypes{
					Soft: true,
					Hard: true,
				},
				HostStates: nagios.HostStates{
					Up:          true,
					Down:        true,
					Unreachable: true,
				},
				ServiceStates: nagios.ServiceStates{
					Ok:       true,
					Warning:  true,
					Critical: true,
					Unknown:  true,
				},
				StartTime: then.Unix(),
				EndTime:   now.Unix(),
			},
		}

		var count nagios.AlertCount

		if err := c.Query(req, &count); err != nil {
			t.Errorf("Query: %v", err)
		}

		if count.Result.TypeText != success {
			t.Errorf("TypeText != %s", success)
		}

		if dumpResponse {
			spew.Dump(count)
		}
	})

	t.Run("blank alert list", func(t *testing.T) {
		req := nagios.AlertListRequest{
			GeneralAlertRequest: nagios.GeneralAlertRequest{
				FormatOptions: nagios.FormatOptions{
					Enumerate: true,
				},
				StartTime: then.Unix(),
				EndTime:   now.Unix(),
			},
		}

		var count nagios.AlertList

		if err := c.Query(req, &count); err != nil {
			t.Errorf("Query: %v", err)
		}

		if count.Result.TypeText != success {
			t.Errorf("TypeText != %s", success)
		}

		if dumpResponse {
			spew.Dump(count)
		}
	})

	t.Run("blank alert list with options switched", func(t *testing.T) {
		req := nagios.AlertListRequest{
			nagios.GeneralAlertRequest{
				FormatOptions: nagios.FormatOptions{
					Whitespace: true,
					Enumerate:  true,
					Bitmask:    true,
					Duration:   true,
				},
				ObjectTypes: nagios.ObjectTypes{
					Host:    true,
					Service: true,
				},
				StateTypes: nagios.StateTypes{
					Soft: true,
					Hard: true,
				},
				HostStates: nagios.HostStates{
					Up:          true,
					Down:        true,
					Unreachable: true,
				},
				ServiceStates: nagios.ServiceStates{
					Ok:       true,
					Warning:  true,
					Critical: true,
					Unknown:  true,
				},
				StartTime: then.Unix(),
				EndTime:   now.Unix(),
			},
		}

		var count nagios.AlertList

		if err := c.Query(req, &count); err != nil {
			t.Errorf("Query: %v", err)
		}

		if count.Result.TypeText != success {
			t.Errorf("TypeText != %s", success)
		}

		if dumpResponse {
			spew.Dump(count)
		}
	})
}