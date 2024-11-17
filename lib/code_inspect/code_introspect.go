package code_inspect

import (
	"fmt"
	"strconv"

	"bufio"
	"encoding/json"
	"os"
	"regexp"
	"strings"
)

func ExtractJSONComment(filePath string) (map[string]interface{}, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var content string
	var inCommentBlock bool

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "/*") {
			inCommentBlock = true
		}
		if inCommentBlock {
			content += line + "\n"
		}
		if strings.Contains(line, "*/") {
			inCommentBlock = false
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	re := regexp.MustCompile(`(?s)\{.*\}`)
	match := re.FindString(content)
	if match == "" {
		return nil, fmt.Errorf("no JSON comment found in the file")
	}

	var jsonComment map[string]interface{}
	if err := json.Unmarshal([]byte(match), &jsonComment); err != nil {
		return nil, err
	}

	return jsonComment, nil
}

func DisplayAllFields(jsonComment map[string]interface{}) {
	jsonData, _ := json.MarshalIndent(jsonComment, "", "  ")
	fmt.Println(string(jsonData))
}

func DisplaySpecificFields(jsonComment map[string]interface{}) {
	fields := make([]string, 0, len(jsonComment))
	for field := range jsonComment {
		fields = append(fields, field)
	}

	fmt.Println("Available fields to display:")
	for idx, field := range fields {
		fmt.Printf("%d. %s\n", idx+1, field)
	}

	fmt.Print("Enter the numbers of the fields you want to display, separated by commas (e.g., 1,2): ")
	var choices string
	fmt.Scanln(&choices)

	selectedFields := strings.Split(choices, ",")
	selectedData := make(map[string]interface{})
	for _, choice := range selectedFields {
		idx := strings.TrimSpace(choice)
		if i, err := strconv.Atoi(idx); err == nil && i > 0 && i <= len(fields) {
			selectedData[fields[i-1]] = jsonComment[fields[i-1]]
		}
	}

	jsonData, _ := json.MarshalIndent(selectedData, "", "  ")
	fmt.Println(string(jsonData))
}

func MainMenu(jsonComment map[string]interface{}) {
	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Display all fields")
		fmt.Println("2. Choose specific fields to display")
		fmt.Println("3. Exit")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			DisplayAllFields(jsonComment)
		case 2:
			DisplaySpecificFields(jsonComment)
		case 3:
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
