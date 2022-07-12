package uti

import (
	"github.com/bwmarrin/discordgo"
)

type Embed struct {
	*discordgo.MessageEmbed
}

const (
	EmbedLimitTitle       = 256
	EmbedLimitDescription = 2048
	EmbedLimitFieldValue  = 1024
	EmbedLimitFieldName   = 256
	EmbedLimitField       = 25
	EmbedLimitFooter      = 2048
	EmbedLimit            = 4000
)

func NewEmbed() *Embed {
	return &Embed{&discordgo.MessageEmbed{}}
}

func (e *Embed) SetTitle(name string) *Embed {
	e.Title = name
	return e
}

func (e *Embed) SetDescription(description string) *Embed {
	if len(description) > 2048 {
		description = description[:2048]
	}
	e.Description = description
	return e
}

func (e *Embed) AddField(name, value string) *Embed {
	fields := make([]*discordgo.MessageEmbedField, 0)
	if len(name) > EmbedLimitFieldName {
		name = name[:EmbedLimitFieldName]
	}
	if len(value) > EmbedLimitFieldValue {
		i := EmbedLimitFieldValue
		extended := false
		for i = EmbedLimitFieldValue; i < len(value); {
			if i != EmbedLimitFieldValue && !extended {
				name += " (extended)"
				extended = true
			}
			if value[i] == []byte(" ")[0] || value[i] == []byte("\n")[0] || value[i] == []byte("-")[0] {
				fields = append(fields, &discordgo.MessageEmbedField{
					Name:  name,
					Value: value[i-EmbedLimitFieldValue : i],
				})
			} else {
				fields = append(fields, &discordgo.MessageEmbedField{
					Name:  name,
					Value: value[i-EmbedLimitFieldValue:i-1] + "-",
				})
				i--
			}
			if (i + EmbedLimitFieldValue) < len(value) {
				i += EmbedLimitFieldValue
			} else {
				break
			}
		}
		if i < len(value) {
			name += " (extended)"
			fields = append(fields, &discordgo.MessageEmbedField{
				Name:  name,
				Value: value[i:],
			})
		}
	} else {
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:  name,
			Value: value,
		})
	}
	e.Fields = append(e.Fields, fields...)
	return e
}

func (e *Embed) SetFooter(args ...string) *Embed {
	iconURL := ""
	text := ""
	proxyURL := ""
	switch {
	case len(args) > 2:
		proxyURL = args[2]
		fallthrough
	case len(args) > 1:
		iconURL = args[1]
		fallthrough
	case len(args) > 0:
		text = args[0]
	case len(args) == 0:
		return e
	}
	e.Footer = &discordgo.MessageEmbedFooter{
		IconURL:      iconURL,
		Text:         text,
		ProxyIconURL: proxyURL,
	}
	return e
}

func (e *Embed) SetImage(args ...string) *Embed {
	var (
		URL      string
		proxyURL string
	)
	switch {
	case len(args) == 0:
		return e
	case len(args) > 0:
		URL = args[0]
	case len(args) > 1:
		proxyURL = args[1]
	}
	e.Image = &discordgo.MessageEmbedImage{
		URL:      URL,
		ProxyURL: proxyURL,
	}
	return e
}

func (e *Embed) SetThumbnail(args ...string) *Embed {
	var (
		URL      string
		proxyURL string
	)
	switch {
	case len(args) == 0:
		return e
	case len(args) > 0:
		URL = args[0]
	case len(args) > 1:
		proxyURL = args[1]

	}
	e.Thumbnail = &discordgo.MessageEmbedThumbnail{
		URL:      URL,
		ProxyURL: proxyURL,
	}
	return e
}

func (e *Embed) SetAuthor(args ...string) *Embed {
	var (
		name     string
		iconURL  string
		URL      string
		proxyURL string
	)
	switch {
	case len(args) == 0:
		return e
	case len(args) > 0:
		name = args[0]
	case len(args) > 1:
		iconURL = args[1]
	case len(args) > 2:
		URL = args[2]
	case len(args) > 3:
		proxyURL = args[3]
	}
	e.Author = &discordgo.MessageEmbedAuthor{
		Name:         name,
		IconURL:      iconURL,
		URL:          URL,
		ProxyIconURL: proxyURL,
	}
	return e
}

func (e *Embed) SetURL(URL string) *Embed {
	e.URL = URL
	return e
}

func (e *Embed) SetColor(clr int) *Embed {
	e.Color = clr
	return e
}

func (e *Embed) InlineAllFields() *Embed {
	for _, v := range e.Fields {
		v.Inline = true
	}
	return e
}

func (e *Embed) Truncate() *Embed {
	e.TruncateDescription()
	e.TruncateFields()
	e.TruncateFooter()
	e.TruncateTitle()
	return e
}

func (e *Embed) TruncateFields() *Embed {
	if len(e.Fields) > 25 {
		e.Fields = e.Fields[:EmbedLimitField]
	}
	for _, v := range e.Fields {
		if len(v.Name) > EmbedLimitFieldName {
			v.Name = v.Name[:EmbedLimitFieldName]
		}
		if len(v.Value) > EmbedLimitFieldValue {
			v.Value = v.Value[:EmbedLimitFieldValue]
		}
	}
	return e
}

func (e *Embed) TruncateDescription() *Embed {
	if len(e.Description) > EmbedLimitDescription {
		e.Description = e.Description[:EmbedLimitDescription]
	}
	return e
}

func (e *Embed) TruncateTitle() *Embed {
	if len(e.Title) > EmbedLimitTitle {
		e.Title = e.Title[:EmbedLimitTitle]
	}
	return e
}

func (e *Embed) TruncateFooter() *Embed {
	if e.Footer != nil && len(e.Footer.Text) > EmbedLimitFooter {
		e.Footer.Text = e.Footer.Text[:EmbedLimitFooter]
	}
	return e
}

func (embed *Embed) SendInteractionMessage(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				embed.MessageEmbed,
			},
		},
	})
}
