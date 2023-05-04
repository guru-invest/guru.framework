package options

type GCPCredentialsOption struct {
	Type                    string `json:"type"`
	ProjectID               string `json:"project_id"`
	PrivateKeyID            string `json:"private_key_id"`
	PrivateKey              string `json:"private_key"`
	ClientEmail             string `json:"client_email"`
	ClientID                string `json:"client_id"`
	AuthURI                 string `json:"auth_uri"`
	TokenURI                string `json:"token_uri"`
	AuthProviderX509CertURL string `json:"auth_provider_x509_cert_url"`
	ClientX509CertURL       string `json:"client_x509_cert_url"`
	Bucket                  string `json:"bucket"`
}

type AWSCredentialsOption struct {
	AccessKey string `json:"access-key"`
	SecretKey string `json:"secret-key"`
	UrlS3     string `json:"url-s3"`
	ARN       string `json:"arn"`
	Region    string `json:"region"`
	Bucket    string `json:"bucket"`
}
