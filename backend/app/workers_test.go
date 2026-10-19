package app_test

import (
	"net/http"
	"strings"
	"testing"

	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tests"
)

// addWorker saves a worker account and returns an auth token for it.
func addWorker(t testing.TB, a core.App, name string) (*core.Record, string) {
	workers, err := a.FindCollectionByNameOrId("workers")
	if err != nil {
		t.Fatal(err)
	}
	w := core.NewRecord(workers)
	w.Set("name", name)
	w.SetPassword("password123")
	if err := a.Save(w); err != nil {
		t.Fatal(err)
	}
	token, err := w.NewAuthToken()
	if err != nil {
		t.Fatal(err)
	}
	return w, token
}

func TestWorkers(t *testing.T) {
	cases := []struct {
		name, as, method, url, body string
		status                      int
		content                     string
		after                       func(testing.TB, *tests.TestApp)
	}{
		{
			name: "workers sign in with their name", method: http.MethodPost,
			url:    "/api/collections/workers/auth-with-password",
			body:   `{"identity":"camera","password":"password123"}`,
			status: http.StatusOK, content: `"token":`,
		},
		{
			name: "heartbeats record when a worker was last seen", as: "worker", method: http.MethodPost,
			url: "/api/app/workers/heartbeat", status: http.StatusNoContent,
			after: func(t testing.TB, a *tests.TestApp) {
				w, _ := a.FindFirstRecordByData("workers", "name", "camera")
				if w.GetDateTime("last_seen").IsZero() {
					t.Error("last_seen was not set")
				}
			},
		},
		{
			name: "people can't send worker heartbeats", as: "member", method: http.MethodPost,
			url: "/api/app/workers/heartbeat", status: http.StatusForbidden, content: `"status":403`,
		},
		{
			name: "workers can't list workers", as: "worker", method: http.MethodGet,
			url: "/api/collections/workers/records", status: http.StatusForbidden, content: "Only superusers",
		},
		{
			name: "items record which worker added them", as: "worker", method: http.MethodPost,
			url: "/api/collections/items/records", body: `{"title":"Front door: person","by":"forged"}`,
			status: http.StatusOK, content: `"by":"camera"`,
		},
		{
			name: "items record which person added them", as: "member", method: http.MethodPost,
			url: "/api/collections/items/records", body: `{"title":"Milk"}`,
			status: http.StatusOK, content: `"by":"member"`,
		},
		{
			name: "signed-out visitors can't add items", method: http.MethodPost,
			url: "/api/collections/items/records", body: `{"title":"Spam"}`,
			status: http.StatusBadRequest, content: `"status":400`,
		},
		{
			name: "workers can't delete items", as: "worker", method: http.MethodDelete,
			url: "/api/collections/items/records/{item}", status: http.StatusNotFound, content: `"status":404`,
		},
		{
			name: "people can delete items", as: "member", method: http.MethodDelete,
			url: "/api/collections/items/records/{item}", status: http.StatusNoContent,
		},
	}

	for _, c := range cases {
		a := newApp(t)
		_, memberToken := addUser(t, a, "member@example.com", "member")
		_, workerToken := addWorker(t, a, "camera")
		itemsCol, _ := a.FindCollectionByNameOrId("items")
		item := core.NewRecord(itemsCol)
		item.Set("title", "Existing")
		if err := a.Save(item); err != nil {
			t.Fatal(err)
		}

		headers := map[string]string{}
		switch c.as {
		case "member":
			headers["Authorization"] = memberToken
		case "worker":
			headers["Authorization"] = workerToken
		}
		s := tests.ApiScenario{
			Name:           c.name,
			Method:         c.method,
			URL:            strings.ReplaceAll(c.url, "{item}", item.Id),
			Body:           strings.NewReader(c.body),
			Headers:        headers,
			ExpectedStatus: c.status,
			TestAppFactory: func(testing.TB) *tests.TestApp { return a },
		}
		if c.content != "" {
			s.ExpectedContent = []string{c.content}
		}
		if c.after != nil {
			after := c.after
			s.AfterTestFunc = func(t testing.TB, a *tests.TestApp, _ *http.Response) { after(t, a) }
		}
		s.Test(t)
	}
}
