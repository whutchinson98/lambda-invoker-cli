package config

type Config struct {
	LambdaName  string `json:"lambdaName"`
	RequestPath string `json:"requestPath"`
	Region      string `json:"region"`
}
