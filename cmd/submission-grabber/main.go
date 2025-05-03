package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/doranych/go-leetcode/pkg/leetcode"
)

func main() {
	token := os.Getenv("LEETCODE_SESSION")
	if token == "" {
		fmt.Println("Please set the LEETCODE_SESSION environment variable.")
		os.Exit(1)
	}
	cli := leetcode.NewClient(token)

	limit := 20
	offset := 0
	for {
		submissionsList, err := cli.GetSubmission(limit, offset)
		if err != nil {
			fmt.Println("network error:", err)
			os.Exit(1)
		}
		offset += limit
		counter := 0
		for _, sub := range submissionsList {
			counter++
			if sub.StatusDisplay != "Accepted" {
				continue
			}
			makeChallengeDir(cli, sub)
		}
		if counter < limit {
			break
		}
	}
}

func makeChallengeDir(cli *leetcode.Client, sub leetcode.Submission) {
	dirName := toSnakeCase(sub.TitleSlug)
	if fi, err := os.Stat(filepath.Join(dirName, "main.go")); err == nil && fi.Size() > 0 {
		fmt.Printf("%+v skipped\n", sub)
		return
	}
	_ = os.Mkdir(dirName, os.ModePerm)

	// Fetch code & question
	// Step 2: Get submission code
	code, err := cli.GetSubmissionCode(sub.Id)
	if code == "" {
		fmt.Printf("Could not fetch code for %s\n", sub.Title)
		fmt.Println(err)
		return
	}

	// Step 3: Fetch question text
	question, err := cli.GetQuestionText(sub.TitleSlug)
	if err != nil {
		fmt.Printf("Could not fetch question for %s\n", sub.Title)
		fmt.Println(err)
		return
	}

	// Step 4: Write files
	// Create main.go
	goFileContent := fmt.Sprintf("package %s\n\n%s", dirName, code)
	_ = os.WriteFile(path.Join(dirName, "main.go"), []byte(goFileContent), os.ModePerm)
	// step 5: Write a test file
	_ = os.WriteFile(path.Join(dirName, "main_test.go"), []byte(fmt.Sprintf("package %s\n\n", dirName)), os.ModePerm)
	// step 6: Write problem.md
	heading := fmt.Sprintf(
		"[%s](%s)\n===\n", sub.Title,
		fmt.Sprintf("https://leetcode.com/problems/%s/", sub.TitleSlug))
	questionMarkdown := question
	_ = os.WriteFile(path.Join(dirName, "problem.md"), []byte((heading + questionMarkdown)), os.ModePerm)

	fmt.Printf("Dumped submission for %s\n", sub.Title)
}

func toSnakeCase(str string) string {
	re := regexp.MustCompile(`([a-z0-9])([A-Z])`)
	snake := strings.ToLower(re.ReplaceAllString(str, "${1}_${2}"))
	return strings.ReplaceAll(snake, "-", "_")
}
