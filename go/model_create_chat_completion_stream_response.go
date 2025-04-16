/*
 * OpenAI API
 *
 * APIs for sampling from and fine-tuning language models
 *
 * API version: 2.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package rkllmopenapi

type CreateChatCompletionStreamResponse struct {

	Id string `json:"id"`

	Object string `json:"object"`

	Created int32 `json:"created"`

	Model string `json:"model"`

	Choices []CreateChatCompletionStreamResponseChoicesInner `json:"choices"`
}
