package auth

type APIKey struct {
	Name      string `json:"name"`
	Key       string `json:"key"`
	ExpiredAt uint64 `json:"expiredAt"`
}

type Token struct {
	// access token
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VybmFtZSI6InN0cmFuZ2VyIiwidXNlcklkIjoiMDA1YjQzOTctOTRjMi00YWZjLWIzNTEtNWIzY2VkNzI4MDkzIiwiaXNzIjoid2lsc29uIiwiZXhwIjoxNjk3MzI0ODY4LCJpYXQiOjE2OTcyODg4NjgsImp0aSI6IjdkZDI4NDZlLTFkZjEtNDBkMS04YmZlLTA3ZGI3ZmE4NmFhYiJ9.HnAaz-WOmugqfdz_oXphsJY_zQl3FCzrtYCm90WJgGU"`
	// refresh token
	Refresh string `json:"refresh,omitempty" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VybmFtZSI6InN0cmFuZ2VyIiwidXNlcklkIjoiMDA1YjQzOTctOTRjMi00YWZjLWIzNTEtNWIzY2VkNzI4MDkzIiwiaXNzIjoid2lsc29uIiwiZXhwIjoxNjk3MzI0ODY4LCJpYXQiOjE2OTcyODg4NjgsImp0aSI6IjdkZDI4NDZlLTFkZjEtNDBkMS04YmZlLTA3ZGI3ZmE4NmFhYiJ9.HnAaz-WOmugqfdz_oXphsJY_zQl3FCzrtYCm90WJgGU"`
}
