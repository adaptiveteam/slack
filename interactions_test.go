package slack

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	dialogSubmissionCallback = `{
  "type": "dialog_submission",
  "submission": {
      "name": "Sigourney Dreamweaver",
      "email": "sigdre@example.com",
      "phone": "+1 800-555-1212",
      "meal": "burrito",
      "comment": "No sour cream please",
      "team_channel": "C0LFFBKPB",
      "who_should_sing": "U0MJRG1AL"
  },
  "callback_id": "employee_offsite_1138b",
  "team": {
      "id": "T1ABCD2E12",
      "domain": "coverbands"
  },
  "user": {
      "id": "W12A3BCDEF",
      "name": "dreamweaver"
  },
  "channel": {
      "id": "C1AB2C3DE",
      "name": "coverthon-1999"
  },
  "action_ts": "936893340.702759",
  "token": "M1AqUUw3FqayAbqNtsGMch72",
  "response_url": "https://hooks.slack.com/app/T012AB0A1/123456789/JpmK0yzoZDeRiqfeduTBYXWQ"
}`
	actionCallback = `{}`
)

func assertInteractionCallback(t *testing.T, callback InteractionCallback, encoded string) {
	var decoded InteractionCallback
	assert.Nil(t, json.Unmarshal([]byte(encoded), &decoded))
	assert.Equal(t, decoded, callback)
}

func TestDialogCallback(t *testing.T) {
	expected := InteractionCallback{
		Type:        InteractionTypeDialogSubmission,
		Token:       "M1AqUUw3FqayAbqNtsGMch72",
		CallbackID:  "employee_offsite_1138b",
		ResponseURL: "https://hooks.slack.com/app/T012AB0A1/123456789/JpmK0yzoZDeRiqfeduTBYXWQ",
		ActionTs:    "936893340.702759",
		Team:        Team{ID: "T1ABCD2E12", Name: "", Domain: "coverbands"},
		Channel: Channel{
			GroupConversation: GroupConversation{
				Conversation: Conversation{
					ID: "C1AB2C3DE",
				},
				Name: "coverthon-1999",
			},
		},
		User: User{
			ID:   "W12A3BCDEF",
			Name: "dreamweaver",
		},
		DialogSubmissionCallback: DialogSubmissionCallback{
			Submission: map[string]string{
				"team_channel":    "C0LFFBKPB",
				"who_should_sing": "U0MJRG1AL",
				"name":            "Sigourney Dreamweaver",
				"email":           "sigdre@example.com",
				"phone":           "+1 800-555-1212",
				"meal":            "burrito",
				"comment":         "No sour cream please",
			},
		},
	}
	assertInteractionCallback(t, expected, dialogSubmissionCallback)
}

func TestActionCallback(t *testing.T) {
	assertInteractionCallback(t, InteractionCallback{}, actionCallback)
}

func TestActionCallback2(t *testing.T) {
	actionCallback := "[{\"block_id\":\"1\"}]"
	//interactionCallback := "{\"type\":\"\",\"token\":\"\",\"callback_id\":\"\",\"response_url\":\"\",\"trigger_id\":\"\",\"action_ts\":\"\",\"team\":{\"id\":\"\",\"name\":\"\",\"domain\":\"\"},\"channel\":{\"id\":\"\",\"created\":0,\"is_open\":false,\"is_group\":false,\"is_shared\":false,\"is_im\":false,\"is_ext_shared\":false,\"is_org_shared\":false,\"is_pending_ext_shared\":false,\"is_private\":false,\"is_mpim\":false,\"unlinked\":0,\"name_normalized\":\"\",\"num_members\":0,\"priority\":0,\"user\":\"\",\"name\":\"\",\"creator\":\"\",\"is_archived\":false,\"members\":null,\"topic\":{\"value\":\"\",\"creator\":\"\",\"last_set\":0},\"purpose\":{\"value\":\"\",\"creator\":\"\",\"last_set\":0},\"is_channel\":false,\"is_general\":false,\"is_member\":false,\"locale\":\"\"},\"user\":{\"id\":\"\",\"team_id\":\"\",\"name\":\"\",\"deleted\":false,\"color\":\"\",\"real_name\":\"\",\"tz_label\":\"\",\"tz_offset\":0,\"profile\":{\"first_name\":\"\",\"last_name\":\"\",\"real_name\":\"\",\"real_name_normalized\":\"\",\"display_name\":\"\",\"display_name_normalized\":\"\",\"email\":\"\",\"skype\":\"\",\"phone\":\"\",\"image_24\":\"\",\"image_32\":\"\",\"image_48\":\"\",\"image_72\":\"\",\"image_192\":\"\",\"image_original\":\"\",\"title\":\"\",\"status_expiration\":0,\"team\":\"\",\"fields\":[]},\"is_bot\":false,\"is_admin\":false,\"is_owner\":false,\"is_primary_owner\":false,\"is_restricted\":false,\"is_ultra_restricted\":false,\"is_stranger\":false,\"is_app_user\":false,\"is_invited_user\":false,\"has_2fa\":false,\"has_files\":false,\"presence\":\"\",\"locale\":\"\",\"updated\":0,\"enterprise_user\":{\"id\":\"\",\"enterprise_id\":\"\",\"enterprise_name\":\"\",\"is_admin\":false,\"is_owner\":false,\"teams\":null}},\"original_message\":{\"replace_original\":false,\"delete_original\":false,\"blocks\":null},\"message\":{\"replace_original\":false,\"delete_original\":false,\"blocks\":null},\"name\":\"\",\"value\":\"\",\"message_ts\":\"\",\"attachment_id\":\"\",\"actions\":{\"AttachmentActions\":null,\"BlockActions\":null},\"submission\":null}"
	var ac ActionCallbacks
	err := json.Unmarshal([]byte(actionCallback), &ac)
	if err != nil {
		t.Errorf("Could not parse to ActionCallbacks: %v\n", err)
	}
	assert.Equal(t, 1, len(ac.BlockActions))
}

func TestActionCallback3(t *testing.T) {
	action := AttachmentAction{
		Text: "text",
	}
	ac := ActionCallbacks{AttachmentActions: []*AttachmentAction{&action}}
	bytes, err := json.Marshal(ac)
	if err != nil {
		t.Errorf("Could not marshal ActionCallbacks: %v\n", err)
	}
	err = ac.UnmarshalJSON(bytes)
	actionCallback := "[{\"name\":\"\",\"text\":\"text\",\"type\":\"\"}]"
	assert.Equal(t, actionCallback, string(bytes))
	var ac2 ActionCallbacks
	err = json.Unmarshal([]byte(actionCallback), &ac2)
	if err != nil {
		t.Errorf("Could not parse to ActionCallbacks: %v\n", err)
	}
	assert.Equal(t, 1, len(ac2.AttachmentActions))
	assert.Equal(t, "text", ac2.AttachmentActions[0].Text)
}
