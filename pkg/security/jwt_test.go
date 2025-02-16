package security

import (
	"fmt"
	"testing"
)

func TestSecurityJWT(t *testing.T) {

	pubKey := `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAqVz/6yls0/czWWR4H6zS
1+0IHmJ7Cq162rrnaOvFa3CpHbnLeWoDfJ/lRiKtHIUcDbkj9B3qCF6JW0beC5zf
fQp3vZ8lqh6L8P7yfoitpym1hx4pUAbdGO85j5ut3yTBkmuleyZ5syEgTeVytc2h
1swHjP4N+3ainKZbHCq4RkpNPjYvI1GnaY3fcazFjwPPPlQS1LmPafG6kZD3V8Cv
LdmJOnxLorCQFpbrSq5Y3xP/ir80bUQriXjO8VIh2zgvrrFW/jljOywV4GVBuQNi
U7PHfawneyiVrJAFiS7tt1GUNYaTzUlF9HEfAzvUj7TEg3oaL5xCM89lORjKdKum
XwIDAQAB
-----END PUBLIC KEY-----
`

	// Create a new JWT object
	j := "eyJhbGciOiJSU0EtT0FFUC0yNTYiLCJlbmMiOiJBMjU2Q0JDLUhTNTEyIiwiY3R5IjoiSldUIn0.YdKaZLCzCMvCyXAR8_gzWLY5iqllNZYZpJ8oOpcKOmMsM89s0Z6VYV3gY61IEgwfDocwqDTXvm5I-NZuyrXV1ePVRR6fHGd2_kTOjJ9gCVxjuzA2D5IN--jvWdkqQWrWImp68VkwtCoF57r7L1QkzYt6WP6wsHuChOvojPyy_j5jW6ix__7oPWLCj4NvN12_miQqXAci60uXpTK0dh_NqN_KhCaBr0Nhx-Y-uEQCkWU_-3OXwOkxj65eTuuufGfQCw9aJl9w4oeQ8acrPdnzxKnWMygIBaTYd4Y0gjqLX1v8Zw46O0C0_ZfEVlv1QPYclO6Vb1dZwPtMzdh4P5pbJg.IKtp3153puhslWj6vOON4w.dKFD_JHoM9a0RFBwB-0-ZVlrunLlFEzp1isZDedUXrnpX-kaJ9ujONVY6xmgpPYZYlSSGJ8chL2DdRsL-zDE-CXgezdOsfZTDoZVnR-4-olNQmI3uMnfKsbmiPFXdqwnuHrKc_q2tEQQ1DvZ-hKbLz_nFl8YVea0TNIkNTWHgWTgQhegUBoh5VFXZGukYM6Px-dFTKF6hGoiL4CrYKO9_UVp9K0rzfBZ1jBim5AoEBolJ-RalbK2RDYgw84qqvpKv5OlDWfaflwgW3YirnYDoTNDcKHrZGhEmodnyk3yK43KeVpVI4NwI1LBjJ5o_kbRLob1Q9bEJLwSCgVyIsHG3KQ7Iej1mIdQqejpxHpUO-cb_ipdsCFpk-YcHzlAetePPkg9sIN9ncnZy6ofsMyskVMwAJnQldLqorFtP8ISdVSR0MPK0y7k9sRQaKq1QMcZr464Xj6IVYbUn8TI561slfW0RCEPhPaYy2u1FjFFcOosNb4XKVqVxjvYyeMV98rt2AC2FGaDK4aPhkTI8qSCDxj70vmONQ0PF6RHWh218nDPSNHx2DuJPkg_0RK8lgvsVVMh9jB7GiRjUp9ElnGldD26QDQtojgggvi6zIN2KNRhN9OOXOe4QxhufnBbUv-dfNGzF2xr37ZyPTkXcL5O86fYXP-06zg5sPw5FEPmhEtMofwiYPBmJzA40yY37XuleaxdBPYWdwkhqCnn8NhKSJoITXQ4ZYaGLi7vbjp5L3IvKqMlgTlv_mpnhriGUxASdckzPY2Wl7smmuOc_ybeKQ.etfmoj0-eFrMqCrzOZlwLfj1eXw-ibqNwzOmJPixEEc"



    // Load private key from a file
    privateKey, err := LoadRSAPrivateKey("private.pem")
    if err != nil {
        t.Fatalf("Error loading private key: %v", err)
    }

    // Decrypt the JWE
    decryptedJWT, err := DecryptJWE(encryptedJWT, privateKey)
    if err != nil {
        t.Fatalf("Error decrypting JWT: %v", err)
    }

    fmt.Println("Decrypted JWT:", decryptedJWT)
    // Add assertions to verify the decrypted JWT if needed
}

	// Load private key from a file (replace with the actual path to your key)
	privateKey, err := LoadRSAPrivateKey("private.pem")
	if err != nil {
		fmt.Println("Error loading private key:", err)
		return
	}

	// Decrypt the JWE
	decryptedJWT, err := DecryptJWE(encryptedJWT, privateKey)
	if err != nil {
		fmt.Println("Error decrypting JWT:", err)
		return
	}

	fmt.Println("Decrypted JWT:", decryptedJWT)
}
