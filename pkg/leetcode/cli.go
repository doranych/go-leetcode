package leetcode

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/hashicorp/go-cleanhttp"
)

const (
	leetcodeGraphqlURL = "https://leetcode.com/graphql"
)

type (
	Client struct {
		headers map[string]string
	}
	SubmissionListResponse struct {
		Data struct {
			SubmissionList SubmissionList `json:"submissionList"`
		} `json:"data"`
	}
	SubmissionList struct {
		Submissions []Submission `json:"submissions"`
	}
	Submission struct {
		Id            string `json:"id"`
		Title         string `json:"title"`
		TitleSlug     string `json:"titleSlug"`
		StatusDisplay string `json:"statusDisplay"`
	}
	SubmissionResponse struct {
		Data struct {
			SubmissionDetails struct {
				Code string `json:"code"`
			} `json:"submissionDetails"`
		} `json:"data"`
	}
	QuestionResponse struct {
		Data struct {
			Question struct {
				Content string `json:"content"`
			} `json:"question"`
		} `json:"data"`
	}

	DailyChallengeResponse struct {
		Data struct {
			ActiveDailyCodingChallengeQuestion struct {
				Date     string   `json:"date"`
				Link     string   `json:"link"`
				Question Question `json:"question"`
			} `json:"activeDailyCodingChallengeQuestion"`
		} `json:"data"`
	}
	Question struct {
		AcRate             float64 `json:"acRate"`
		Difficulty         string  `json:"difficulty"`
		FreqBar            float64 `json:"freqBar"`
		QuestionFrontendId string  `json:"questionFrontendId"`
		IsFavor            bool    `json:"isFavor"`
		IsPaidOnly         bool    `json:"isPaidOnly"`
		Status             string  `json:"status"`
		Title              string  `json:"title"`
		TitleSlug          string  `json:"titleSlug"`
		HasSolution        bool    `json:"hasSolution"`
		TopicTags          []struct {
			Name string `json:"name"`
			ID   string `json:"id"`
			Slug string `json:"slug"`
		} `json:"topicTags"`
		Content string `json:"content"`
	}
)

func NewClient(token string) *Client {
	return &Client{
		headers: map[string]string{
			"Cookie":       "LEETCODE_SESSION=" + token,
			"Referer":      "https://leetcode.com/",
			"Content-Type": "application/json",
		},
	}
}

func (c *Client) GetSubmission(limit, offset int) ([]Submission, error) {
	submissionsResp := SubmissionListResponse{}
	reqBody := map[string]interface{}{
		"query": `query submissionList($offset: Int!, $limit: Int!) {
submissionList(offset: $offset, limit: $limit) {
    submissions {
        id
        title
        titleSlug
        lang
        statusDisplay
    }
  }
}`,
		"variables": map[string]interface{}{
			"offset": offset,
			"limit":  limit,
		},
	}

	respBytes, err := c.doRequest(reqBody)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(respBytes, &submissionsResp)
	if err != nil {
		return nil, err
	}
	return submissionsResp.Data.SubmissionList.Submissions, err
}

func (c *Client) GetSubmissionCode(submissionId string) (string, error) {
	reqBody := map[string]interface{}{
		"query": `query submissionDetails($submissionId: Int!) {
  submissionDetails(submissionId: $submissionId) {
    code
  }
}`,
		"variables": map[string]interface{}{
			"submissionId": submissionId,
		},
	}
	respBytes, err := c.doRequest(reqBody)
	if err != nil {
		return "", err
	}
	var resp SubmissionResponse
	err = json.Unmarshal(respBytes, &resp)
	if err != nil {
		return "", err
	}
	return resp.Data.SubmissionDetails.Code, nil
}

func (c *Client) GetQuestionText(titleSlug string) (string, error) {
	reqBody := map[string]interface{}{
		"query": `query getQuestionDetail($titleSlug: String!) {
  question(titleSlug: $titleSlug) {
    content
  }
}`,
		"variables": map[string]interface{}{
			"titleSlug": titleSlug,
		},
	}
	respBytes, err := c.doRequest(reqBody)
	if err != nil {
		return "", err
	}
	var resp QuestionResponse
	err = json.Unmarshal(respBytes, &resp)
	if err != nil {
		return "", err
	}
	return resp.Data.Question.Content, nil
}

func (c *Client) GetDailyChallenge() (*Question, error) {
	reqBody := map[string]interface{}{
		"query": `
		query questionOfToday {
			activeDailyCodingChallengeQuestion {
				date
				link
				question {
					acRate
					difficulty
					freqBar
					questionFrontendId
					isFavor
					isPaidOnly
					status
					title
					titleSlug
					hasSolution
					topicTags {
						name
						id
						slug
					}
					content
				}
			}
		}`,
	}
	respBytes, err := c.doRequest(reqBody)
	if err != nil {
		return nil, err
	}
	var resp DailyChallengeResponse
	err = json.Unmarshal(respBytes, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Data.ActiveDailyCodingChallengeQuestion.Question, nil
}

func (c *Client) doRequest(query map[string]interface{}) ([]byte, error) {
	bodyBytes, err := json.Marshal(query)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, leetcodeGraphqlURL, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return nil, err
	}

	for key, value := range c.headers {
		req.Header.Set(key, value)
	}

	client := cleanhttp.DefaultClient()
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()
	return io.ReadAll(resp.Body)
}
