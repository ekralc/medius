package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
)

// getVideoFromAttachments returns the URL of the first attachment containing a WebM
// or nil if none exists.
func getVideoFromMessage(messages map[string]*discordgo.Message) string {
	for _, msg := range messages {
		for _, attachment := range msg.Attachments {
			if attachment.ContentType == "video/webm" || strings.HasSuffix(".webm", attachment.Filename) {
				return attachment.URL
			}
		}
	}

	return ""
}

func interactionHandler(ctx *gin.Context) {
	var interaction discordgo.Interaction
	if ctx.BindJSON(&interaction) != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if interaction.Type == discordgo.InteractionPing {
		ctx.JSON(http.StatusOK, discordgo.InteractionResponse{
			Type: discordgo.InteractionResponsePong,
		})
		return
	}

	if interaction.Type == discordgo.InteractionApplicationCommand {
		data := interaction.ApplicationCommandData()
		switch data.Name {
		case "ping":
			ctx.JSON(http.StatusOK, discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Pong!",
				},
			})
			return

		case "Caption":
			fmt.Println(data.Resolved)
			ctx.JSON(http.StatusOK, discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseModal,
				Data: &discordgo.InteractionResponseData{
					Title:    "Image Caption",
					CustomID: "input_modal",
					Components: []discordgo.MessageComponent{
						discordgo.ActionsRow{
							Components: []discordgo.MessageComponent{
								discordgo.TextInput{
									CustomID:    "caption",
									Label:       "Caption",
									Style:       discordgo.TextInputShort,
									Placeholder: "An image caption",
									Required:    true,
								},
							},
						},
					},
				},
			})
			return

		case "Convert WebM to MP4":
			data := interaction.ApplicationCommandData().Resolved.Messages
			vidURL := getVideoFromMessage(data)

			if vidURL == "" {
				log.Println("Received request to convert but contained no vids")
				ctx.JSON(http.StatusBadRequest, gin.H{})
				return
			}

			publishVideoNotification(&VideoNotification{
				VideoURL:         vidURL,
				InteractionToken: interaction.Token,
			})

			log.Println("Sent video publish notification")

			ctx.JSON(http.StatusOK, discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Flags: discordgo.MessageFlagsEphemeral,
				},
			})
			return
		}
	}

	ctx.AbortWithStatus(http.StatusNotImplemented)
}
