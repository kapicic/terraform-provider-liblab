// Generated by LIBLAB | https://liblab.com

package token

type FindByUserIdRequestParams struct {
	UserId *float64 `queryParam:"userId" required:"true"`
}

func (params *FindByUserIdRequestParams) SetUserId(userId float64) {
	params.UserId = &userId
}