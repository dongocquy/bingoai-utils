package userstore

type PromptWithSub struct {
	Main         string   `json:"main" dynamodbav:"main"`
	Sub          []string `json:"sub" dynamodbav:"sub"`
	FormatPrompt string   `json:"formatPrompt" dynamodbav:"formatPrompt"`
	RelatedURL   string   `json:"related_url" dynamodbav:"related_url"`
}

type KeyData struct {
	OpenAIKey       string                   `json:"OpenAI_key" dynamodbav:"OpenAI_key"`
	OpenAIModel     string                   `json:"OpenAIModel" dynamodbav:"OpenAIModel"`
	OpenAIEmbedding string                   `json:"OpenAIEmbedding" dynamodbav:"OpenAIEmbedding"`
	ClassifyIntent  string                   `json:"classifyIntent" dynamodbav:"classifyIntent"`
	PostgresDSN     string                   `json:"PostgresDSN" dynamodbav:"PostgresDSN"`
	Prompts         map[string]PromptWithSub `json:"prompts" dynamodbav:"prompts"`
}

type UserConfig struct {
	ID         string  `json:"id" dynamodbav:"id"`
	Keys       KeyData `json:"keys" dynamodbav:"keys"`
	Active     bool    `json:"active" dynamodbav:"active"`
	ClientName string  `json:"client_name" dynamodbav:"client_name"`
	CreatedAt  string  `json:"created_at" dynamodbav:"created_at"`
	ProjectID  string  `json:"project_id" dynamodbav:"project_id"`
}

type ClientUser struct {
	UserEmail   string `json:"user_email" dynamodbav:"user_email"`
	ClientEmail string `json:"client_email" dynamodbav:"client_email"`
	Role        string `json:"role" dynamodbav:"role"`
}
