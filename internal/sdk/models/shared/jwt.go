// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type JWTAlgorithm string

const (
	JWTAlgorithmHs256 JWTAlgorithm = "HS256"
	JWTAlgorithmHs384 JWTAlgorithm = "HS384"
	JWTAlgorithmHs512 JWTAlgorithm = "HS512"
	JWTAlgorithmRs256 JWTAlgorithm = "RS256"
	JWTAlgorithmRs384 JWTAlgorithm = "RS384"
	JWTAlgorithmRs512 JWTAlgorithm = "RS512"
	JWTAlgorithmEs256 JWTAlgorithm = "ES256"
	JWTAlgorithmEs384 JWTAlgorithm = "ES384"
	JWTAlgorithmEs512 JWTAlgorithm = "ES512"
	JWTAlgorithmPs256 JWTAlgorithm = "PS256"
	JWTAlgorithmPs384 JWTAlgorithm = "PS384"
	JWTAlgorithmPs512 JWTAlgorithm = "PS512"
	JWTAlgorithmEdDsa JWTAlgorithm = "EdDSA"
)

func (e JWTAlgorithm) ToPointer() *JWTAlgorithm {
	return &e
}
func (e *JWTAlgorithm) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "HS256":
		fallthrough
	case "HS384":
		fallthrough
	case "HS512":
		fallthrough
	case "RS256":
		fallthrough
	case "RS384":
		fallthrough
	case "RS512":
		fallthrough
	case "ES256":
		fallthrough
	case "ES384":
		fallthrough
	case "ES512":
		fallthrough
	case "PS256":
		fallthrough
	case "PS384":
		fallthrough
	case "PS512":
		fallthrough
	case "EdDSA":
		*e = JWTAlgorithm(v)
		return nil
	default:
		return fmt.Errorf("invalid value for JWTAlgorithm: %v", v)
	}
}

type JWTConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *JWTConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type Jwt struct {
	Algorithm *JWTAlgorithm `json:"algorithm,omitempty"`
	Consumer  *JWTConsumer  `json:"consumer,omitempty"`
	// Unix epoch when the resource was created.
	CreatedAt    *int64   `json:"created_at,omitempty"`
	ID           *string  `json:"id,omitempty"`
	Key          *string  `json:"key,omitempty"`
	RsaPublicKey *string  `json:"rsa_public_key,omitempty"`
	Secret       *string  `json:"secret,omitempty"`
	Tags         []string `json:"tags,omitempty"`
}

func (o *Jwt) GetAlgorithm() *JWTAlgorithm {
	if o == nil {
		return nil
	}
	return o.Algorithm
}

func (o *Jwt) GetConsumer() *JWTConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *Jwt) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *Jwt) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Jwt) GetKey() *string {
	if o == nil {
		return nil
	}
	return o.Key
}

func (o *Jwt) GetRsaPublicKey() *string {
	if o == nil {
		return nil
	}
	return o.RsaPublicKey
}

func (o *Jwt) GetSecret() *string {
	if o == nil {
		return nil
	}
	return o.Secret
}

func (o *Jwt) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}
