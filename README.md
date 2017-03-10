# Openpay Go API Wrapper

[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/mazingstudio/openpay)

Este package es un wrapper de la API de Openpay.

## Instrucciones de Uso

Para interactuar con la API de Openpay se debe crear una instancia de `Merchant` (la cual es segura para uso concurrente). Posteriormente, todas las instancias de `Customer` obtenidas a través de peticiones a la API tendrán una referencia al `Merchant` al que pertenecen. Esta referencia es necesaria para hacer operaciones en cuentas de clientes; si una instancia de `Customer` no se obtiene a través de la API, es necesario proveer una referencia a un `Merchant` manualmente. Para más información, ver la documentación de GoDoc.

```go
import (
	"github.com/mazingstudio/openpay"
)

func main() {
	merchant := openpay.NewMerchant("myMerchantID", "myPrivateKey")

	customer, err := merchant.AddCustomer(&openpay.CustomerArgs{
		Name:     "Juan",
		LastName: "Perez",
		Email:    "juan@email.com",
		Address:  openpay.Address{
			// etc.
		},
	})
	if err != nil {
		panic(err)
	}

	card, err := customer.AddCardWithToken(&openpay.CardTokenArgs{
		TokenID:         "aCardToken",
		DeviceSessionID: "aSessionID",
	})
	if err != nil {
		panic(err)
	}
}
```
