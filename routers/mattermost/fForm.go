package mattermost

import (
	"fmt"
	"net/http"

	"github.com/mattermost/mattermost-app-test/constants"
	"github.com/mattermost/mattermost-app-test/utils"
	"github.com/mattermost/mattermost-plugin-apps/apps"
)

const fullMarkdown = "## Markdown title" +
	"\nHello world" +
	"\nText styles: _italics_ **bold** **_bold-italic_** ~~strikethrough~~ `code`" +
	"\nUsers and channels: @sysadmin ~town-square" +
	"\n```" +
	"\nCode block" +
	"\n```" +
	"\n:+1: :banana_dance:" +
	"\n***" +
	"\n> Quote\n" +
	"\nLink: [here](www.google.com)" +
	"\nImage: ![img](https://gdm-catalog-fmapi-prod.imgix.net/ProductLogo/4acbc64f-552d-4944-8474-b44a13a7bd3e.png?auto=format&q=50&fit=fill)" +
	"\nList:" +
	"\n- this" +
	"\n- is" +
	"\n- a" +
	"\n- list" +
	"\nNumbered list" +
	"\n1. this" +
	"\n2. is" +
	"\n3. a" +
	"\n4. list" +
	"\nItems" +
	"\n- [ ] Item one" +
	"\n- [ ] Item two" +
	"\n- [x] Completed item"

const markdownTable = "\n| Left-Aligned  | Center Aligned  | Right Aligned |" +
	"\n| :------------ |:---------------:| -----:|" +
	"\n| Left column 1 | this text       |  $100 |" +
	"\n| Left column 2 | is              |   $10 |" +
	"\n| Left column 3 | centered        |    $1 |"

func fFormOK(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
	resp := apps.CallResponse{
		Type: apps.CallResponseTypeForm,
		Form: &apps.Form{
			Title:  "Test",
			Header: "Test header",
			Call: &apps.Call{
				Path: constants.BindingPathOK,
			},
			Fields: []*apps.Field{
				{
					Name:       "text",
					Type:       apps.FieldTypeText,
					Label:      "text",
					ModalLabel: "text",
				},
			},
		},
	}
	utils.WriteCallResponse(w, resp)
}

func fFullFormOK(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
	resp := apps.CallResponse{
		Type: apps.CallResponseTypeForm,
		Form: &apps.Form{
			Title:  "Test Full Form",
			Header: "Test header",
			Call: &apps.Call{
				Path: constants.BindingPathLookupOK,
			},
			Fields: []*apps.Field{
				{
					Name:  "lookup",
					Type:  apps.FieldTypeDynamicSelect,
					Label: "lookup",
				},
				{
					Name:  "text",
					Type:  apps.FieldTypeText,
					Label: "text",
				},
				{
					Type:        "markdown",
					Description: "***\n## User information\nRemember to fill all these fields with the **user** information, not the general information.",
				},
				// {
				// 	Name: "mk1",
				// 	Type: "markdown",
				// 	Description: "## Markdown title" +
				// 		"\nHello world" +
				// 		"\nText styles: _italics_ **bold** **_bold-italic_** ~~strikethrough~~ `code`" +
				// 		"\nUsers and channels: @sysadmin ~town-square" +
				// 		"\n```" +
				// 		"\nCode block" +
				// 		"\n```" +
				// 		"\n:+1: :banana_dance:" +
				// 		"\n***" +
				// 		"\n> Quote\n" +
				// 		"\nLink: [here](www.google.com)" +
				// 		"\nImage: ![img](https://gdm-catalog-fmapi-prod.imgix.net/ProductLogo/4acbc64f-552d-4944-8474-b44a13a7bd3e.png?auto=format&q=50&fit=fill)" +
				// 		"\nList:" +
				// 		"\n- this" +
				// 		"\n- is" +
				// 		"\n- a" +
				// 		"\n- list" +
				// 		"\nNumbered list" +
				// 		"\n1. this" +
				// 		"\n2. is" +
				// 		"\n3. a" +
				// 		"\n4. list" +
				// 		"\nItems" +
				// 		"\n- [ ] Item one" +
				// 		"\n- [ ] Item two" +
				// 		"\n- [x] Completed item",
				// },
				// {
				// 	Name: "mk2",
				// 	Type: "markdown",
				// 	Description: "\n| Left-Aligned  | Center Aligned  | Right Aligned |" +
				// 		"\n| :------------ |:---------------:| -----:|" +
				// 		"\n| Left column 1 | this text       |  $100 |" +
				// 		"\n| Left column 2 | is              |   $10 |" +
				// 		"\n| Left column 3 | centered        |    $1 |",
				// },
				{
					Name:  "boolean",
					Type:  apps.FieldTypeBool,
					Label: "boolean",
				},
				{
					Name:  "channel",
					Type:  apps.FieldTypeChannel,
					Label: "channel",
				},
				{
					Name:  "user",
					Type:  apps.FieldTypeUser,
					Label: "user",
				},
				{
					Name:  "static",
					Type:  apps.FieldTypeStaticSelect,
					Label: "static",
					SelectStaticOptions: []apps.SelectOption{
						{
							Label: "static value 1",
							Value: "sv1",
						},
						{
							Label: "static value 2",
							Value: "sv2",
						},
					},
				},
			},
		},
	}
	utils.WriteCallResponse(w, resp)
}

func fFullFormDisabledOK(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
	resp := apps.CallResponse{
		Type: apps.CallResponseTypeForm,
		Form: &apps.Form{
			Title:  "Test Full Form",
			Header: "Test header",
			Call: &apps.Call{
				Path: constants.BindingPathLookupOK,
			},
			Fields: []*apps.Field{
				{
					Name:     "lookup",
					Type:     apps.FieldTypeDynamicSelect,
					Label:    "lookup",
					ReadOnly: true,
				},
				{
					Name:     "text",
					Type:     apps.FieldTypeText,
					Label:    "text",
					ReadOnly: true,
					Value:    "Hello world",
				},
				{
					Type:  "markdown",
					Value: "Hello ~~world~~",
				},
				{
					Name:     "boolean",
					Type:     apps.FieldTypeBool,
					Label:    "boolean",
					ReadOnly: true,
				},
				{
					Name:     "channel",
					Type:     apps.FieldTypeChannel,
					Label:    "channel",
					ReadOnly: true,
				},
				{
					Name:     "user",
					Type:     apps.FieldTypeUser,
					Label:    "user",
					ReadOnly: true,
				},
				{
					Name:  "static",
					Type:  apps.FieldTypeStaticSelect,
					Label: "static",
					SelectStaticOptions: []apps.SelectOption{
						{
							Label: "static value 1",
							Value: "sv1",
						},
						{
							Label: "static value 2",
							Value: "sv2",
						},
					},
					ReadOnly: true,
				},
			},
		},
	}
	utils.WriteCallResponse(w, resp)
}

func fDynamicFormOK(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
	numFields := len(c.Values)
	fields := []*apps.Field{}

	for i := 0; i < numFields+5; i++ {
		fields = append(fields, &apps.Field{
			Name:          fmt.Sprintf("static%v", i),
			Type:          apps.FieldTypeStaticSelect,
			Label:         fmt.Sprintf("static%v", i),
			SelectRefresh: true,
			SelectStaticOptions: []apps.SelectOption{
				{
					Label: "static value 1",
					Value: "sv1",
				},
				{
					Label: "static value 2",
					Value: "sv2",
				},
			},
		})
	}

	resp := apps.CallResponse{
		Type: apps.CallResponseTypeForm,
		Form: &apps.Form{
			Title:  "Test Dynamic Form",
			Header: "Test header",
			Call: &apps.Call{
				Path: constants.BindingPathDynamicFormOK,
			},
			Fields: fields,
		},
	}
	utils.WriteCallResponse(w, resp)
}

func fFormInvalid(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
	resp := apps.CallResponse{
		Type: apps.CallResponseTypeForm,
	}
	utils.WriteCallResponse(w, resp)
}

func fFormMultiselect(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
	resp := apps.CallResponse{
		Type: apps.CallResponseTypeForm,
		Form: &apps.Form{
			Title:  "Test Multiselect Form",
			Header: "Test header",
			Call: &apps.Call{
				Path: constants.BindingPathOK,
			},
			Fields: []*apps.Field{
				{
					Name:          "static",
					Type:          apps.FieldTypeStaticSelect,
					Label:         "static",
					SelectIsMulti: true,
					SelectStaticOptions: []apps.SelectOption{
						{
							Label: "static value 1",
							Value: "sv1",
						},
						{
							Label: "static value 2",
							Value: "sv2",
						},
						{
							Label: "static value 3",
							Value: "sv3",
						},
						{
							Label: "static value 4",
							Value: "sv4",
						},
					},
				},
				{
					Name:          "user",
					Type:          apps.FieldTypeUser,
					Label:         "user",
					SelectIsMulti: true,
				},
				{
					Name:          "channel",
					Type:          apps.FieldTypeChannel,
					Label:         "channel",
					SelectIsMulti: true,
				},
			},
		},
	}
	utils.WriteCallResponse(w, resp)
}

func fFormWithButtonsOK(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
	resp := apps.CallResponse{
		Type: apps.CallResponseTypeForm,
		Form: &apps.Form{
			Title:         "Test multiple buttons Form",
			Header:        "Test header",
			SubmitButtons: "static",
			Call: &apps.Call{
				Path: constants.BindingPathOK,
			},
			Fields: []*apps.Field{
				{
					Name:          "static",
					Type:          apps.FieldTypeStaticSelect,
					Label:         "static",
					SelectIsMulti: true,
					SelectStaticOptions: []apps.SelectOption{
						{
							Label: "button1",
							Value: "button1",
						},
						{
							Label: "button2",
							Value: "button2",
						},
						{
							Label: "button3",
							Value: "button3",
						},
						{
							Label: "button4",
							Value: "button4",
						},
					},
				},
				{
					Name:  "user",
					Type:  apps.FieldTypeUser,
					Label: "user",
				},
				{
					Name:  "channel",
					Type:  apps.FieldTypeChannel,
					Label: "channel",
				},
			},
		},
	}
	utils.WriteCallResponse(w, resp)
}

func fFormWithMarkdownError(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
	resp := apps.CallResponse{
		Type: apps.CallResponseTypeForm,
		Form: &apps.Form{
			Title:  "Test markdown descriptions and errors",
			Header: "Test header",
			Call: &apps.Call{
				Path: constants.BindingPathMarkdownFormError,
			},
			Fields: []*apps.Field{
				{
					Name:  "static",
					Type:  apps.FieldTypeStaticSelect,
					Label: "static",
					Description: `| Option | Message  | Image |
| :------------ |:---------------:| -----:|
| Opt1 | You are good     |  :smile: |
| Opt2 | You are awesome              | :+1: |
| Opt3| You are great       |    :smirk:  |`,
					SelectIsMulti: true,
					SelectStaticOptions: []apps.SelectOption{
						{
							Label: "button1",
							Value: "button1",
						},
						{
							Label: "button2",
							Value: "button2",
						},
						{
							Label: "button3",
							Value: "button3",
						},
						{
							Label: "button4",
							Value: "button4",
						},
					},
				},
				{
					Name:        "text",
					Type:        apps.FieldTypeText,
					Label:       "text",
					Description: fullMarkdown, //"Go [here](www.google.com) for more information.",
				},
				{
					Name:  "boolean",
					Type:  apps.FieldTypeBool,
					Label: "boolean",
					Description: `Mark this field only if:
					1. You want
					2. You need
					3. You should`,
				},
			},
		},
	}
	utils.WriteCallResponse(w, resp)
}

func fFormWithMarkdownErrorMissingField(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
	resp := apps.CallResponse{
		Type: apps.CallResponseTypeForm,
		Form: &apps.Form{
			Title:  "Test markdown descriptions and errors",
			Header: "Test header",
			Call: &apps.Call{
				Path: constants.BindingPathMarkdownFormErrorMissingField,
			},
			Fields: []*apps.Field{
				{
					Name:  "static",
					Type:  apps.FieldTypeStaticSelect,
					Label: "static",
					Description: `| Option | Message  | Image |
| :------------ |:---------------:| -----:|
| Opt1 | You are good     |  :smile: |
| Opt2 | You are awesome              | :+1: |
| Opt3| You are great       |    :smirk:  |`,
					SelectIsMulti: true,
					SelectStaticOptions: []apps.SelectOption{
						{
							Label: "button1",
							Value: "button1",
						},
						{
							Label: "button2",
							Value: "button2",
						},
						{
							Label: "button3",
							Value: "button3",
						},
						{
							Label: "button4",
							Value: "button4",
						},
					},
				},
				{
					Name:        "text",
					Type:        apps.FieldTypeText,
					Label:       "text",
					Description: fullMarkdown, //"Go [here](www.google.com) for more information.",
				},
				{
					Name:  "boolean",
					Type:  apps.FieldTypeBool,
					Label: "boolean",
					Description: `Mark this field only if:
					1. You want
					2. You need
					3. You should`,
				},
			},
		},
	}
	utils.WriteCallResponse(w, resp)
}
