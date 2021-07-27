package main

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
	"github.com/slack-go/slack"

	"github.com/nicolas-martin/hive/api/bot/client"
	"github.com/nicolas-martin/hive/api/config"
	"github.com/nicolas-martin/hive/api/model"
	"github.com/nicolas-martin/hive/api/repo"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	s := client.NewSlack(cfg)
	r := repo.NewRepo(cfg)
	r.AddUser(&model.User{DisplayName: "martinni39", SlackUserID: ""})

	userID, err := s.GetUserID("martinni39")
	if err != nil {
		log.Fatal(err)
	}

	// TODO: try out more complex messages
	// err = s.PostMessage(userID, "Hello world")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// btest := &slack.Blocks{}
	// err = btest.UnmarshalJSON(blocksStr)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	blocksStr := `{
    "replace_original": false,
    "delete_original": false,
    "blocks": [
        {
            "type": "section",
            "text": {
                "type": "mrkdwn",
                "text": "You have a new request:\n*\u003cfakeLink.toEmployeeProfile.com|Fred Enriquez - New device request\u003e*"
            }
        },
        {
            "type": "section",
            "fields": [
                {
                    "type": "mrkdwn",
                    "text": "*Type:*\nComputer (laptop)"
                },
                {
                    "type": "mrkdwn",
                    "text": "*When:*\nSubmitted Aut 10"
                },
                {
                    "type": "mrkdwn",
                    "text": "*Last Update:*\nMar 10, 2015 (3 years, 5 months)"
                },
                {
                    "type": "mrkdwn",
                    "text": "*Reason:*\nAll vowel keys aren't working."
                },
                {
                    "type": "mrkdwn",
                    "text": "*Specs:*\n\"Cheetah Pro 15\" - Fast, really fast\""
                }
            ]
        },
        {
            "type": "actions",
            "elements": [
                {
                    "type": "button",
                    "text": {
                        "type": "plain_text",
                        "text": "Approve"
                    },
                    "value": "click_me_123"
                },
                {
                    "type": "button",
                    "text": {
                        "type": "plain_text",
                        "text": "Deny"
                    },
                    "value": "click_me_123"
                }
            ]
        }
    ]
}`
	attachment := new(slack.Attachment)
	err = json.Unmarshal([]byte(blocksStr), attachment)
	if err != nil {
		log.Fatal(err)
	}
	err = s.PostAttachement(userID, *attachment)
	if err != nil {
		log.Fatal(err)
	}

}

func postAttachement() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	s := client.NewSlack(cfg)
	r := repo.NewRepo(cfg)
	r.AddUser(&model.User{DisplayName: "martinni39", SlackUserID: ""})

	userID, err := s.GetUserID("martinni39")
	if err != nil {
		log.Fatal(err)
	}

	// TODO: try out more complex messages
	// err = s.PostMessage(userID, "Hello world")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// btest := &slack.Blocks{}
	// err = btest.UnmarshalJSON(blocksStr)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	originalAttachmentJson := `{
		"id": 1,
		"blocks": [
		  {
			"type": "section",
			"block_id": "xxxx",
			"text": {
			  "type": "mrkdwn",
			  "text": "Pick something:",
			  "verbatim": true
			},
			"accessory": {
			  "type": "static_select",
			  "action_id": "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
			  "placeholder": {
				"type": "plain_text",
				"text": "Select one item",
				"emoji": true
			  },
			  "options": [
				{
				  "text": {
					"type": "plain_text",
					"text": "ghi",
					"emoji": true
				  },
				  "value": "ghi"
				}
			  ]
			}
		  }
		],
		"color": "#13A554",
		"fallback": "[no preview available]"
	  }`

	attachment := new(slack.Attachment)
	err = json.Unmarshal([]byte(originalAttachmentJson), attachment)
	if err != nil {
		log.Fatal(err)
	}
	err = s.PostAttachement(userID, *attachment)
	if err != nil {
		log.Fatal(err)
	}

}

func blocks() {

}
