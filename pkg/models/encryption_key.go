package models

type EncryptionKey struct {
	KeyId   string `json:"keyId"`
	KeyName string `json:"keyName,omitempty"`
	Status  string `json:"status,omitempty"`
}
