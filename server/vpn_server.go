package main

import (
	"fmt"
	"golang.zx2c4.com/wireguard/wgctrl"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
	"log"
)

func main() {
	// Create a WireGuard controller
	client, err := wgctrl.New()
	if err != nil {
		log.Fatalf("Failed to create WireGuard controller: %v", err)
	}
	defer client.Close()

	// Generate a new private key for the server
	serverKey, err := wgtypes.GeneratePrivateKey()
	if err != nil {
		log.Fatalf("Failed to generate server private key: %v", err)
	}
	fmt.Printf("Server Private Key: %s\n", serverKey.String())

	// Define server configuration
	// Define server configuration
	port := 51820
	serverConfig := wgtypes.Config{
		PrivateKey: &serverKey,
		ListenPort: &port,
	}

	// Apply the configuration to a WireGuard interface
	if err := client.ConfigureDevice("wg0", serverConfig); err != nil {
		log.Fatalf("Failed to configure WireGuard interface: %v", err)
	}

	fmt.Println("WireGuard VPN server is up and running on interface wg0.")
}
