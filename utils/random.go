package utils

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

/*
Siempre que haya una función init en un paquete
será lo primero que se ejecutará de dicho paquete al
iniciar el programa
*/
func init() {
	rand.Seed(time.Now().UnixNano())
}

// Retorna un numero entero no negativo
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// Retorna una cadena string aleatoria
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomMoney() int64 {
	return RandomInt(0, 1000)
}
