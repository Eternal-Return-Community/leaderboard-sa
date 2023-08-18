package commands

import (
	"fmt"
	"image/color"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/fogleman/gg"
)

const (
	template = "./src/assets/template.png"
	final    = "./src/assets/final.png"
)

func Dodge(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {

	if len(args) < 2 {
		s.ChannelMessageSendReply(m.ChannelID, "Adicione um nickname.", m.Reference())
		return
	}

	nickname := args[1]
	if len(nickname) >= 16 {
		s.ChannelMessageSendReply(m.ChannelID, "Nickname precisa ter até 16 digitos.", m.Reference())
		return
	}

	generateImage(nickname, s, m)
}

func generateImage(nickname string, s *discordgo.Session, m *discordgo.MessageCreate) {

	message, _ := s.ChannelMessageSendReply(m.ChannelID, "Imagem está sendo gerada. Aguarde alguns segundos.", m.Reference())

	img, err := gg.LoadImage(template)
	if err != nil {
		log.Fatal(err)
	}

	dc := gg.NewContextForImage(img)
	dc.SetColor(color.White)

	if err := dc.LoadFontFace("/usr/share/fonts/truetype/dejavu/DejaVuSans-Bold.ttf", fontSize(len(nickname))); err != nil {
		log.Fatal(err)
	}

	dc.DrawString(nickname, 200.0, 500.0)

	err = dc.SavePNG(final)
	if err != nil {
		log.Fatal(err)
	}

	s.ChannelMessageSendComplex(m.ChannelID, loadImage(m.Author.ID))
	s.ChannelMessageDelete(m.ChannelID, message.ID)

}

func fontSize(nickname int) float64 {

	size := 0.0
	if nickname > 1 && nickname <= 10 {
		size = 35.0
		return size
	} else if nickname >= 10 && nickname <= 16 {
		size = 25.0
		return size
	}

	return size

}

func loadImage(author string) *discordgo.MessageSend {
	file, err := os.Open(final)
	if err != nil {
		log.Fatal(err)
	}

	image := discordgo.MessageSend{
		Content: fmt.Sprintf("<@%s>", author),
		Files: []*discordgo.File{
			{
				Name:   "final.png",
				Reader: file,
			},
		},
	}

	return &image
}
