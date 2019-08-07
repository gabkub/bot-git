package messageBuilders

import (
	"bot-git/bot/newsSrc/newsAbstract"
	"github.com/mattermost/mattermost-server/model"
)

func Text(text string) *model.Post {
	return &model.Post{
		Message: text,
	}
}

func TitleThumbUrl(title, text, thumbURL string) *model.Post {
	return &model.Post{
		Props: map[string]interface{}{
			"attachments": []model.SlackAttachment{
				{
					Fields: []*model.SlackAttachmentField{
						{
							Short: false,
							Title: title,
							Value: text,
						},
					},
					ThumbURL: thumbURL,
				},
			},
		},
	}
}

func Image(header, imageUrl string) *model.Post {
	return &model.Post{
		Props: map[string]interface{}{
			"attachments": []model.SlackAttachment{
				{
					ImageURL:  imageUrl,
					Title:     header,
					TitleLink: imageUrl,
				},
			},
		},
	}
}

func News(news *newsAbstract.News) *model.Post {
	return &model.Post{
		Props: map[string]interface{}{
			"attachments": []model.SlackAttachment{
				{
					ImageURL:  news.Img.ImageUrl,
					Title:     news.Img.Header,
					TitleLink: news.TitleLink,
				},
			},
		},
	}
}
