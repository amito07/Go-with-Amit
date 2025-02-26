package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"
	"time"
)

type todo struct {
	todoName     string
	todoStatus   string
	todoPriority int
	deadline     string
}

var todoList []todo
var todoName string
var todoStatus string
var todoPriority int
var dateline string
var todoItem todo

func addTask(name string) {
	fmt.Println("Enter task name, status, priority, dateline:")
	fmt.Scan(&todoName, &todoStatus, &todoPriority, &dateline)
	todoItem = todo{todoName, todoStatus, todoPriority, dateline}
	todoList = append(todoList, todoItem)

	filename := "TodoList/" + name + ".csv"

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Check if the file is empty
    fileInfo, err := file.Stat()
    if err != nil {
        fmt.Println(err)
        return
    }

	writer := csv.NewWriter(file)
    defer writer.Flush()

	// If the file is empty, write the headers
    if fileInfo.Size() == 0 {
        headers := []string{"Task Name", "Status", "Priority", "Deadline"}
        err = writer.Write(headers)
        if err != nil {
            fmt.Println(err)
            return
        }
    }

	    // Write the new task record
		newRecord := []string{todoName, todoStatus, strconv.Itoa(todoPriority), dateline}
		err = writer.Write(newRecord)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Task added successfully")
}

func viewTask(name string) {
    filename := "TodoList/" + name + ".csv"

    file, err := os.Open(filename)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()

    reader := csv.NewReader(file)
    records, err := reader.ReadAll()

    if err != nil {
        fmt.Println(err)
        return
    }

    // Create a new tabwriter
    w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.Debug)

    // Print records with days left
    fmt.Fprintf(w, "Task Name\tStatus\tPriority\tDeadline\tDays Left\t\n")
    for i, record := range records {
        // Skip the header row
        if i == 0 {
            continue
        }
        deadline, err := time.Parse("2006-01-02", record[3])
        if err != nil {
            fmt.Println("Error parsing date:", err)
            continue
        }
        daysLeft := int(time.Until(deadline).Hours() / 24)
        fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%d\t\n", record[0], record[1], record[2], record[3], daysLeft)
    }
    // Flush the writer
    w.Flush()
}

func deleteTask(name string) {
	fmt.Println("========================== Delete task ==========================")

	filename := "TodoList/" + name + ".csv"

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Enter task name to delete:")
	var taskName string
	fmt.Scan(&taskName)

	for i, record := range records {
		if record[0] == taskName {
			records = append(records[:i], records[i+1:]...)
			break
		}
	}

	file, err = os.OpenFile(filename, os.O_TRUNC|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, record := range records {
		err = writer.Write(record)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	fmt.Println("Task deleted successfully")
	
}

func updateTask(name string) {
	fmt.Println("========================= Update task ==========================")
	filename := "TodoList/" + name + ".csv"

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Enter task name to update:")
	var taskName string
	fmt.Scan(&taskName)

	for i, record := range records {
		if record[0] == taskName {
			fmt.Println("Enter new task name, status, priority, dateline:")
			fmt.Scan(&todoName, &todoStatus, &todoPriority, &dateline)
			records[i] = []string{todoName, todoStatus, strconv.Itoa(todoPriority), dateline}
			break
		}
	}

	file, err = os.OpenFile(filename, os.O_TRUNC|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, record := range records {
		err = writer.Write(record)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	fmt.Println("Task updated successfully")
}

func TodoApp(username string){
	for {
	fmt.Println("1. Add Task")
	fmt.Println("2. View Task")
	fmt.Println("3. Delete Task")
	fmt.Println("4. Update Task")
	fmt.Println("5. Exit")

	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		addTask(username)
	case 2:
		viewTask(username)
	case 3:
		deleteTask(username)
	case 4:
		updateTask(username)
	case 5:
		break
	default:
		fmt.Println("Invalid choice")
		break
	}
		
	}
}

func signUp() {
	fmt.Println("Enter username")
	var username string
	fmt.Scan(&username)

	file, err := os.OpenFile("user.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	readFile, err := os.Open("user.csv")
	if err != nil {
        fmt.Println(err)
        return
    }
	defer readFile.Close()

	reader := csv.NewReader(readFile)
	records, err := reader.ReadAll()

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, record := range records {
		if record[0] == username {
			fmt.Println("Username already exists")
			return
		}
	}

	writer := csv.NewWriter(file)
	defer writer.Flush()

	newRecord := []string{username}
	err = writer.Write(newRecord)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("New record added successfully")
}

func login() {
	fmt.Println("Enter username")
	var username string
	fmt.Scan(&username)

	readFile, err := os.Open("user.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer readFile.Close()

	reader := csv.NewReader(readFile)
	records, err := reader.ReadAll()

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, record := range records {
		if record[0] == username {
			fmt.Println("Login successful")
			TodoApp(username)
			return
		}
	}

	fmt.Println("Invalid username")
	return
}

func main() {
	fmt.Println("Welcome to Todo App")

	fmt.Println("1. Sign Up")
	fmt.Println("2. Login")
	fmt.Println("3. Exit")

	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		signUp()
	case 2:
		login()
	case 3:
		return
	default:
		fmt.Println("Invalid choice")	
	}
}