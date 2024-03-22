// Generated by LIBLAB | https://liblab.com

package token

type CreateTokenResponse struct {
	Id   *float64 `json:"id,omitempty" required:"true"`
	Name *string  `json:"name,omitempty" required:"true"`
	// Warning: only shown once, please save securely
	Token *string `json:"token,omitempty" required:"true"`
	// Defaults to 1 year from creation date
	ExpiresAt *string `json:"expiresAt,omitempty" required:"true"`
}

func (c *CreateTokenResponse) SetId(id float64) {
	c.Id = &id
}

func (c *CreateTokenResponse) GetId() *float64 {
	if c == nil {
		return nil
	}
	return c.Id
}

func (c *CreateTokenResponse) SetName(name string) {
	c.Name = &name
}

func (c *CreateTokenResponse) GetName() *string {
	if c == nil {
		return nil
	}
	return c.Name
}

func (c *CreateTokenResponse) SetToken(token string) {
	c.Token = &token
}

func (c *CreateTokenResponse) GetToken() *string {
	if c == nil {
		return nil
	}
	return c.Token
}

func (c *CreateTokenResponse) SetExpiresAt(expiresAt string) {
	c.ExpiresAt = &expiresAt
}

func (c *CreateTokenResponse) GetExpiresAt() *string {
	if c == nil {
		return nil
	}
	return c.ExpiresAt
}