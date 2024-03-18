package restful

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestBindRequestBody_success(t *testing.T) {
	setup()
	v := NewValidator()

	type SomeReq struct {
		Foo string `json:"foo"`
		Bar int8   `json:"bar"`
	}
	e.Post("/test", func(c *fiber.Ctx) error {
		foobar := SomeReq{}
		err := v.Bind(c, &foobar)
		if err != nil {
			return c.SendStatus(400)
		}
		return c.SendString("ok")
	})

	ic := SomeReq{
		Foo: "foo",
		Bar: 1,
	}
	icm, err := json.Marshal(ic)
	assert.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/test", bytes.NewReader(icm))
	req.Header.Set("Content-Type", "application/json")

	resp, err := e.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestValidateRequestBody_badReq(t *testing.T) {
	setup()
	v := NewValidator()

	type SomeReq struct {
		Foo string `json:"foo" validate:"min=5"`
		Bar int8   `json:"bar" validate:"required"`
	}
	e.Post("/test", func(c *fiber.Ctx) error {
		foobar := SomeReq{}
		err := v.Bind(c, &foobar)
		if err != nil {
			return c.SendStatus(400)
		}
		return c.SendString("ok")
	})

	ic := SomeReq{
		Foo: "foo",
	}
	icm, err := json.Marshal(ic)
	assert.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/test", bytes.NewReader(icm))
	req.Header.Set("Content-Type", "application/json")

	resp, err := e.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestBindRequestBody_fail(t *testing.T) {
	setup()
	v := NewValidator()

	type SomeReq struct {
		Foo string `json:"foo"`
		Bar int8   `json:"bar"`
	}
	e.Post("/test", func(c *fiber.Ctx) error {
		foobar := SomeReq{}
		err := v.Bind(c, &foobar)
		if err != nil {
			return c.SendStatus(400)
		}
		return c.SendString("ok")
	})

	type Unk struct {
		Foo int8 `json:"foo"`
	}
	ic := Unk{Foo: 2}
	icm, err := json.Marshal(ic)
	assert.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/test", bytes.NewReader(icm))
	req.Header.Set("Content-Type", "application/json")

	resp, err := e.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}
