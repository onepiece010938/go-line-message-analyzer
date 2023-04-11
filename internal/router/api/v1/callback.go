package v1

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/onepiece010938/go-line-message-analyzer/internal/app"
)

type LineHandler struct {
	bot *linebot.Client
}

func Callback(app *app.Application) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("In Callback2")
		log.Println(c.Request)
		ctx := c.Request.Context()
		events, err := app.LineBotClient.ParseRequest(c.Request)
		if err != nil {
			if err == linebot.ErrInvalidSignature {

				c.JSON(http.StatusBadRequest, nil)
			} else {
				c.JSON(http.StatusInternalServerError, nil)
			}
			return
		}
		lineHandler := LineHandler{bot: app.LineBotClient}
		for _, event := range events {
			log.Printf("Got event %v", event)
			switch event.Type {
			case linebot.EventTypeMessage:
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					if err := lineHandler.handleText(ctx, app, message, event.ReplyToken, event.Source); err != nil {
						log.Print(err)
					}

				case *linebot.FileMessage:
					if err := lineHandler.handleFile(ctx, app, message, event.ReplyToken); err != nil {
						log.Print(err)
					}
				default:
					log.Printf("Unknown message: %v", message)
				}
			case linebot.EventTypeFollow:
				if err := lineHandler.replyText(event.ReplyToken, "Got followed event"); err != nil {
					log.Print(err)
				}
			case linebot.EventTypeUnfollow:
				log.Printf("Unfollowed this bot: %v", event)
			case linebot.EventTypeJoin:
				if err := lineHandler.replyText(event.ReplyToken, "Joined "+string(event.Source.Type)); err != nil {
					log.Print(err)
				}
			case linebot.EventTypeLeave:
				log.Printf("Left: %v", event)
			case linebot.EventTypePostback:
				data := event.Postback.Data
				if data == "DATE" || data == "TIME" || data == "DATETIME" {
					data += fmt.Sprintf("(%v)", *event.Postback.Params)
				}
				if err := lineHandler.replyText(event.ReplyToken, "Got postback: "+data); err != nil {
					log.Print(err)
				}
			case linebot.EventTypeBeacon:
				if err := lineHandler.replyText(event.ReplyToken, "Got beacon: "+event.Beacon.Hwid); err != nil {
					log.Print(err)
				}
			default:
				log.Printf("Unknown event: %v", event)
			}
		}

	}

}

func (l *LineHandler) handleText(ctx context.Context, app *app.Application, message *linebot.TextMessage, replyToken string, source *linebot.EventSource) error {
	switch message.Text {
	case "sample":
		test, err := app.AnalyzeService.AnalyzeTest(ctx)
		if err != nil {
			return l.replyText(replyToken, "ERROR")
		}
		return l.replyText(replyToken, test)
	case "profile":
		if source.UserID != "" {
			profile, err := l.bot.GetProfile(source.UserID).Do()
			if err != nil {
				return l.replyText(replyToken, err.Error())
			}
			if _, err := l.bot.ReplyMessage(
				replyToken,
				linebot.NewTextMessage("Display name: "+profile.DisplayName),
				linebot.NewTextMessage("Status message: "+profile.StatusMessage),
			).Do(); err != nil {
				return err
			}
		} else {
			return l.replyText(replyToken, "Bot can't use profile API without user ID")
		}

	case "confirm":
		template := linebot.NewConfirmTemplate(
			"Do it?",
			linebot.NewMessageAction("Yes", "Yes!"),
			linebot.NewMessageAction("No", "No!"),
		)
		if _, err := l.bot.ReplyMessage(
			replyToken,
			linebot.NewTemplateMessage("Confirm alt text", template),
		).Do(); err != nil {
			return err
		}

	case "datetime":
		template := linebot.NewButtonsTemplate(
			"", "", "Select date / time !",
			linebot.NewDatetimePickerAction("date", "DATE", "date", "", "", ""),
			linebot.NewDatetimePickerAction("time", "TIME", "time", "", "", ""),
			linebot.NewDatetimePickerAction("datetime", "DATETIME", "datetime", "", "", ""),
		)
		if _, err := l.bot.ReplyMessage(
			replyToken,
			linebot.NewTemplateMessage("Datetime pickers alt text", template),
		).Do(); err != nil {
			return err
		}

	default:
		log.Printf("Echo message to %s: %s", replyToken, message.Text)
		if _, err := l.bot.ReplyMessage(
			replyToken,
			linebot.NewTextMessage(message.Text),
		).Do(); err != nil {
			return err
		}
	}
	return nil
}

func (l *LineHandler) handleFile(ctx context.Context, app *app.Application, message *linebot.FileMessage, replyToken string) error {
	// content, err := l.bot.GetMessageContent(message.ID).Do()
	// if err != nil {
	// 	return l.replyText(replyToken, err.Error())
	// }
	content, err := l.bot.GetMessageContent(message.ID).Do()
	if err != nil {
		return l.replyText(replyToken, err.Error())
	}
	last, err := app.AnalyzeService.StartAnalyze(content.Content)
	if err != nil {
		return l.replyText(replyToken, err.Error())
	}

	return l.replyText(replyToken, fmt.Sprintf("File `%s` (%d bytes) received. %s", message.FileName, message.FileSize, last))
}

func (l *LineHandler) replyText(replyToken, text string) error {
	if _, err := l.bot.ReplyMessage(
		replyToken,
		linebot.NewTextMessage(text),
	).Do(); err != nil {
		return err
	}
	return nil
}

/*
func (l *LineHandler) handleHeavyContent(messageID string, callback func(*os.File) error) error {
	content, err := l.bot.GetMessageContent(messageID).Do()
	if err != nil {
		return err
	}
	defer content.Content.Close()
	log.Printf("Got file: %s", content.ContentType)
	originalContent, err := saveContent(content.Content)
	if err != nil {
		return err
	}
	return callback(originalContent)
}

func saveContent(content io.ReadCloser) (*os.File, error) {
	file, err := ioutil.TempFile("downloadDir", "")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	_, err = io.Copy(file, content)
	if err != nil {
		return nil, err
	}
	log.Printf("Saved %s", file.Name())
	return file, nil
}
*/
