package commands

import (
	"fmt"
	"image/color"
	"log"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/fogleman/gg"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const (
	pathCharacters = "./src/assets/characters"
	template       = "./src/assets/template.png"
	final          = "./src/assets/final.png"
)

func Dodge(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {

	if len(args) < 3 {
		s.ChannelMessageSendReply(m.ChannelID, "Digite o nome de um personagem e um nickname. \nExemplo: **!dodge aya skye**", m.Reference())
		return
	}

	character := strings.ToLower(args[1])
	nickname := args[2]

	if len(nickname) >= 16 {
		s.ChannelMessageSendReply(m.ChannelID, "Nickname precisa ter até 16 digitos.", m.Reference())
		return
	}

	generateImage(character, nickname, s, m)
}

func generateImage(character string, nickname string, s *discordgo.Session, m *discordgo.MessageCreate) {

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

	characters, err := loadCharacter(character, s, m)
	if err != nil {
		if len(characters) == 0 {
			s.ChannelMessageEdit(m.ChannelID, message.ID, "Atualmente não possui nenhuma imagem de personagem disponível. \nEnvie uma mensagem para <@811913211737014322> avisando sobre a situação.")
			return
		}

		s.ChannelMessageEdit(m.ChannelID, message.ID, fmt.Sprintf("**%s** não foi encontrado. Digite o nome de outro personagem. \nPersonagens: **%s**", character, characters))
		return
	}

	characterImage, err := gg.LoadImage(pathCharacters + characters)
	if err != nil {
		log.Fatal(err)
	}

	dc.DrawString(nickname, 200.0, 500.0)
	dc.DrawImage(characterImage, -10.0, 320.0)

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

func loadCharacter(c string, s *discordgo.Session, m *discordgo.MessageCreate) (string, error) {
	list, err := os.ReadDir(pathCharacters)
	if err != nil {
		s.ChannelMessageSendReply(m.ChannelID, "Ocorreu um erro ao carregar a lista de personagens.", m.Reference())
	}

	var png string
	var characterList []string

	for _, character := range list {
		if strings.Replace(character.Name(), ".png", "", 1) == c {
			png = character.Name()
			break
		}
		characterList = append(characterList, cases.Title(language.Und, cases.NoLower).String(strings.Replace(character.Name(), ".png", "", 1)))
	}

	if len(png) == 0 {
		return strings.Join(characterList, ", "), fmt.Errorf(c)
	}

	return "/" + png, nil
}
