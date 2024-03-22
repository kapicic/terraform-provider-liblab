// Generated by LIBLAB | https://liblab.com

package token

type GetTokenResponse struct {
	Id        *float64 `json:"id,omitempty" required:"true"`
	Name      *string  `json:"name,omitempty" required:"true"`
	ExpiresAt *string  `json:"expiresAt,omitempty" required:"true"`
	Scope     []string `json:"scope,omitempty" required:"true"`
}

func (g *GetTokenResponse) SetId(id float64) {
	g.Id = &id
}

func (g *GetTokenResponse) GetId() *float64 {
	if g == nil {
		return nil
	}
	return g.Id
}

func (g *GetTokenResponse) SetName(name string) {
	g.Name = &name
}

func (g *GetTokenResponse) GetName() *string {
	if g == nil {
		return nil
	}
	return g.Name
}

func (g *GetTokenResponse) SetExpiresAt(expiresAt string) {
	g.ExpiresAt = &expiresAt
}

func (g *GetTokenResponse) GetExpiresAt() *string {
	if g == nil {
		return nil
	}
	return g.ExpiresAt
}

func (g *GetTokenResponse) SetScope(scope []string) {
	g.Scope = scope
}

func (g *GetTokenResponse) GetScope() []string {
	if g == nil {
		return nil
	}
	return g.Scope
}