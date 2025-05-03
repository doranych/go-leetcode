package main

import (
	"fmt"
	"os"
	"path"
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

	// Fetch the daily challenge
	question, err := cli.GetDailyChallenge()
	if err != nil {
		fmt.Println("Failed to fetch daily challenge:", err)
		os.Exit(1)
	}

	// Create directory with snake case name
	dirName := toSnakeCase(question.TitleSlug)
	err = os.MkdirAll(dirName, os.ModePerm)
	if err != nil {
		fmt.Printf("Failed to create directory %s: %v\n", dirName, err)
		os.Exit(1)
	}

	// Create main.go
	mainGoContent := fmt.Sprintf("package %s\n\n// TODO: Implement solution\n", dirName)
	err = os.WriteFile(path.Join(dirName, "main.go"), []byte(mainGoContent), os.ModePerm)
	if err != nil {
		fmt.Printf("Failed to create main.go: %v\n", err)
		os.Exit(1)
	}

	// Create main_test.go
	testGoContent := fmt.Sprintf("package %s\n\nimport \"testing\"\n\n// TODO: Add tests\n", dirName)
	err = os.WriteFile(path.Join(dirName, "main_test.go"), []byte(testGoContent), os.ModePerm)
	if err != nil {
		fmt.Printf("Failed to create main_test.go: %v\n", err)
		os.Exit(1)
	}

	// Create problem.md
	problemUrl := fmt.Sprintf("https://leetcode.com/problems/%s/", question.TitleSlug)
	heading := fmt.Sprintf("[%s](%s)\n===\n", question.Title, problemUrl)

	problemMdContent := heading + question.Content
	problemMdContent = strings.ReplaceAll(problemMdContent, "\r\n", "\n")
	err = os.WriteFile(path.Join(dirName, "problem.md"), []byte(problemMdContent), os.ModePerm)
	if err != nil {
		fmt.Printf("Failed to create problem.md: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully created files for today's challenge: %s\n", question.Title)
}

func toSnakeCase(str string) string {
	re := regexp.MustCompile(`([a-z0-9])([A-Z])`)
	snake := strings.ToLower(re.ReplaceAllString(str, "${1}_${2}"))
	return strings.ReplaceAll(snake, "-", "_")
}
