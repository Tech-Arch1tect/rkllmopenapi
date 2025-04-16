/*
 * OpenAI API
 *
 * APIs for sampling from and fine-tuning language models
 *
 * API version: 2.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package rkllmopenapi

type DeleteModelResponse struct {

	Id string `json:"id"`

	Object string `json:"object"`

	Deleted bool `json:"deleted"`
}
