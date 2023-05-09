// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourcePokeapiPokeapiEnum string

const (
	SourcePokeapiPokeapiEnumPokeapi SourcePokeapiPokeapiEnum = "pokeapi"
)

func (e SourcePokeapiPokeapiEnum) ToPointer() *SourcePokeapiPokeapiEnum {
	return &e
}

func (e *SourcePokeapiPokeapiEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "pokeapi":
		*e = SourcePokeapiPokeapiEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourcePokeapiPokeapiEnum: %v", v)
	}
}

// SourcePokeapi - The values required to configure the source.
type SourcePokeapi struct {
	// Pokemon requested from the API.
	PokemonName string                   `json:"pokemon_name"`
	SourceType  SourcePokeapiPokeapiEnum `json:"sourceType"`
}
