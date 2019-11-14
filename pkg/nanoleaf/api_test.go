package nanoleaf

import (
	"context"
	"net/http"
	"testing"
)

const IP = "DEVICE_ID"
const TOKEN = "TOKEN"

func createClient() (Client, error) {
	return NewClient(http.DefaultClient, IP)
}

func createAuthenticatedClient() (Client, error) {
	return NewClientWithToken(http.DefaultClient, IP, TOKEN)
}
func TestCreateClient(t *testing.T) {
	_, err := createClient()
	if err != nil {
		t.Error(err)
	}
}

func TestAuthorize(t *testing.T) {
	client, _ := createClient()

	err := client.Authorize(context.Background())

	if err != nil {
		t.Error(err)
	}

}

func TestGetInfo(t *testing.T) {
	client, _ := createAuthenticatedClient()

	_, err := client.GetInfo(context.Background())

	if err != nil {
		t.Error(err)
	}

}

func TestSetState(t *testing.T) {
	client, _ := createAuthenticatedClient()
	state := &State{
		On: &OnValue{Value: Bool(true)},
	}
	err := client.SetState(context.Background(), state)
	if err != nil {
		t.Error(err)
	}
}

func TestGetPower(t *testing.T) {
	client, _ := createAuthenticatedClient()
	_, err := client.GetPower(context.Background())
	if err != nil {
		t.Error(err)
	}
}

func TestSetPower(t *testing.T) {
	client, _ := createAuthenticatedClient()
	err := client.SetPower(context.Background(), true)
	if err != nil {
		t.Error(err)
	}
}

func TestGetBrightness(t *testing.T) {
	client, _ := createAuthenticatedClient()
	_, err := client.GetBrightness(context.Background())
	if err != nil {
		t.Error(err)
	}
}

func TestSetBrightness(t *testing.T) {
	client, _ := createAuthenticatedClient()
	err := client.SetBrightness(context.Background(), 2)
	if err != nil {
		t.Error(err)
	}
}

func TestSetBrightnessWithDuration(t *testing.T) {
	client, _ := createAuthenticatedClient()
	err := client.SetBrightnessWithDuration(context.Background(), 3, 4)
	if err != nil {
		t.Error(err)
	}
}

func TestGetHue(t *testing.T) {
	client, _ := createAuthenticatedClient()
	_, err := client.GetHue(context.Background())
	if err != nil {
		t.Error(err)
	}
}

func TestSetHue(t *testing.T) {
	client, _ := createAuthenticatedClient()
	err := client.SetHue(context.Background(), 2)
	if err != nil {
		t.Error(err)
	}
}

func TestGetSaturation(t *testing.T) {
	client, _ := createAuthenticatedClient()
	_, err := client.GetSaturation(context.Background())
	if err != nil {
		t.Error(err)
	}
}

func TestSetSaturation(t *testing.T) {
	client, _ := createAuthenticatedClient()
	err := client.SetSaturation(context.Background(), 2)
	if err != nil {
		t.Error(err)
	}
}

func TestGetColorTemperature(t *testing.T) {
	client, _ := createAuthenticatedClient()
	_, err := client.GetColorTemperature(context.Background())
	if err != nil {
		t.Error(err)
	}
}

func TestGetColorMode(t *testing.T) {
	client, _ := createAuthenticatedClient()
	_, err := client.GetColorMode(context.Background())
	if err != nil {
		t.Error(err)
	}
}

func TestGetCurrentEffect(t *testing.T) {
	client, _ := createAuthenticatedClient()
	_, err := client.GetCurrentEffect(context.Background())
	if err != nil {
		t.Error(err)
	}
}

func TestGetAllEffects(t *testing.T) {
	client, _ := createAuthenticatedClient()
	_, err := client.GetAllEffects(context.Background())
	if err != nil {
		t.Error(err)
	}
}

func TestSelectEffect(t *testing.T) {
	client, _ := createAuthenticatedClient()
	err := client.SelectEffect(context.Background(), "Snowfall")
	if err != nil {
		t.Error(err)
	}
}
