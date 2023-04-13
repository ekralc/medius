package main

import (
	"crypto/ed25519"
	"encoding/hex"
	"fmt"
	"net/http"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
)

var ApplicationPublicKey ed25519.PublicKey

func init() {
	public_key_hex := os.Getenv("DISCORD_PUBLIC_KEY")
	if len(public_key_hex) > 0 {
		bytes, err := hex.DecodeString(public_key_hex)
		if err != nil {
			panic(fmt.Errorf("couldn't decode public key hex: %v", err))
		}

		ApplicationPublicKey = bytes
	}
}

func VerifyRequestMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !discordgo.VerifyInteraction(ctx.Request, ApplicationPublicKey) {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, "Couldn't verify signature")
			return
		}
	}
}
