package nanoleaf

import (
	"context"
	"errors"
	"io/ioutil"
	"net/url"
)

var NOT_IMPLEMENTED error = errors.New("Not Implemented")

func (c *client) Authorize(ctx context.Context) error {
	authResp := &AddUserResponse{}
	_, err := c.makeRequest(ctx, "POST", "new", nil, authResp)
	if err != nil {
		return err
	}

	c.Token = authResp.AuthToken
	if c.Token != nil {
		tokenPath := &url.URL{Path: *c.Token}
		c.BaseURL = c.BaseURL.ResolveReference(tokenPath)
		return nil
	}

	return errors.New("Unable to authenticate")
}

func (c *client) GetInfo(ctx context.Context) (*DeviceInfo, error) {
	info := &DeviceInfo{}
	_, err := c.makeRequest(ctx, "GET", "", nil, info)
	if err != nil {
		return nil, err
	}
	return info, nil
}

func (c *client) SetState(ctx context.Context, state *State) error {
	_, err := c.makeRequest(ctx, "PUT", "state", state, nil)
	return err
}

func (c *client) GetPower(ctx context.Context) (*OnValue, error) {
	value := &OnValue{}
	_, err := c.makeRequest(ctx, "GET", "state/on", nil, value)
	if err != nil {
		return nil, err
	}
	return value, nil
}

func (c *client) SetPower(ctx context.Context, value bool) error {
	state := &State{
		On: &OnValue{Value: &value},
	}
	return c.SetState(ctx, state)
}

func (c *client) GetBrightness(ctx context.Context) (*RangedValue, error) {
	value := &RangedValue{}
	_, err := c.makeRequest(ctx, "GET", "state/brightness", nil, value)
	if err != nil {
		return nil, err
	}
	return value, err
}

func (c *client) SetBrightness(ctx context.Context, value int) error {
	state := &State{
		Brightness: &RangedValue{Value: &value},
	}
	return c.SetState(ctx, state)
}

func (c *client) SetBrightnessWithDuration(ctx context.Context, value int, duration int) error {
	state := &State{
		Brightness: &RangedValue{Value: &value, Duration: &duration},
	}
	return c.SetState(ctx, state)
}

func (c *client) GetHue(ctx context.Context) (*RangedValue, error) {
	value := &RangedValue{}
	_, err := c.makeRequest(ctx, "GET", "state/hue", nil, value)
	if err != nil {
		return nil, err
	}
	return value, nil
}

func (c *client) SetHue(ctx context.Context, value int) error {
	state := &State{
		Hue: &RangedValue{Value: &value},
	}
	return c.SetState(ctx, state)
}

func (c *client) GetSaturation(ctx context.Context) (*RangedValue, error) {
	value := &RangedValue{}
	_, err := c.makeRequest(ctx, "GET", "state/sat", nil, value)
	if err != nil {
		return nil, err
	}
	return value, nil
}

func (c *client) SetSaturation(ctx context.Context, value int) error {
	state := &State{
		Sat: &RangedValue{Value: &value},
	}
	return c.SetState(ctx, state)
}

func (c *client) GetColorTemperature(ctx context.Context) (*RangedValue, error) {
	value := &RangedValue{}
	_, err := c.makeRequest(ctx, "GET", "state/ct", nil, value)
	if err != nil {
		return nil, err
	}
	return value, nil
}

func (c *client) SetColorTemperature(ctx context.Context, value int) error {
	state := &State{
		CT: &RangedValue{Value: &value},
	}
	return c.SetState(ctx, state)
}

func (c *client) GetColorMode(ctx context.Context) (*string, error) {
	resp, err := c.makeRequestRaw(ctx, "GET", "state/colorMode", nil)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	value := string(body)

	return &value, err
}

func (c *client) GetCurrentEffect(ctx context.Context) (*string, error) {
	var value *string
	_, err := c.makeRequestRaw(ctx, "GET", "effects/select", nil)

	return value, err
}

func (c *client) GetAllEffects(ctx context.Context) ([]*string, error) {
	value := &[]*string{}
	_, err := c.makeRequest(ctx, "GET", "effects/effectsList", nil, value)
	return *value, err
}

func (c *client) SelectEffect(ctx context.Context, effect string) error {
	value := &Effects{Select: &effect}
	_, err := c.makeRequest(ctx, "PUT", "effects", value, nil)
	return err
}
