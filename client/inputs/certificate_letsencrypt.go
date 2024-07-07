package inputs

type CertificateLetsEncrypt struct {
	Name           string `json:"name"`
	Certificate    string `json:"certificate"`
	CertificateKey string `json:"certificate_key"`
}
