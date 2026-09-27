package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	ID    int
	Name  string
	Age   int
	Email string
}

// All students are kept in this slice while the program runs.
var students []Student

// reader lets us read a whole line (so names like "Rahul Kumar" work).
var reader = bufio.NewReader(os.Stdin)

// ---------- Input helpers ----------

// readLine prints a prompt and returns what the user typed (without spaces at the ends).
func readLine(prompt string) string {
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

// readInt keeps asking until the user types a valid number.
func readInt(prompt string) int {
	for {
		text := readLine(prompt)
		num, err := strconv.Atoi(text)
		if err == nil {
			return num
		}
		fmt.Println("Please enter a valid number.")
	}
}

// findIndexByID returns the position of the student in the slice, or -1 if not found.
func findIndexByID(id int) int {
	for i, s := range students {
		if s.ID == id {
			return i
		}
	}
	return -1
}

func printStudent(s Student) {
	fmt.Printf("ID: %d | Name: %-6s | Age: %d | Email: %s\n", s.ID, s.Name, s.Age, s.Email)
}

// ---------- Step 3: Add Student ----------

func addStudent() {
	fmt.Println("\n--- Add Student ---")

	id := readInt("Enter ID: ")
	if findIndexByID(id) != -1 {
		fmt.Println("A student with this ID already exists!")
		return
	}

	name := readLine("Enter Name: ")
	age := readInt("Enter Age: ")
	email := readLine("Enter Email: ")

	s := Student{ID: id, Name: name, Age: age, Email: email}
	students = append(students, s)

	fmt.Println("\nStudent added successfully!")
}

// ---------- Step 4: List Students ----------

func listStudents() {
	fmt.Println("\n--- All Students ---")

	if len(students) == 0 {
		fmt.Println("No students yet.")
		return
	}

	for _, s := range students {
		printStudent(s)
	}
}

// ---------- Step 5: Find Student ----------

func findStudent() {
	fmt.Println("\n--- Find Student ---")

	id := readInt("Enter Student ID: ")
	i := findIndexByID(id)
	if i == -1 {
		fmt.Println("Student not found.")
		return
	}

	s := students[i]
	fmt.Println("\nStudent Found")
	fmt.Println("ID:", s.ID)
	fmt.Println("Name:", s.Name)
	fmt.Println("Age:", s.Age)
	fmt.Println("Email:", s.Email)
}

// ---------- Step 6: Update Student ----------

func updateStudent() {
	fmt.Println("\n--- Update Student ---")

	id := readInt("Enter Student ID: ")
	i := findIndexByID(id)
	if i == -1 {
		fmt.Println("Student not found.")
		return
	}

	fmt.Println("Press Enter to keep the current value.")

	// We change students[i] directly (by index), so the change stays in the slice.
	// Note: `for _, s := range students { s.Name = ... }` would NOT work,
	// because s is only a copy.
	name := readLine(fmt.Sprintf("Name (%s): ", students[i].Name))
	if name != "" {
		students[i].Name = name
	}

	ageText := readLine(fmt.Sprintf("Age (%d): ", students[i].Age))
	if ageText != "" {
		age, err := strconv.Atoi(ageText)
		if err != nil {
			fmt.Println("Invalid age, keeping old value.")
		} else {
			students[i].Age = age
		}
	}

	email := readLine(fmt.Sprintf("Email (%s): ", students[i].Email))
	if email != "" {
		students[i].Email = email
	}

	fmt.Println("\nStudent updated successfully!")
}

// ---------- Step 7: Delete Student ----------

func deleteStudent() {
	fmt.Println("\n--- Delete Student ---")

	id := readInt("Enter Student ID: ")
	i := findIndexByID(id)
	if i == -1 {
		fmt.Println("Student not found.")
		return
	}

	// Take everything before i, and everything after i, and join them.
	students = append(students[:i], students[i+1:]...)

	fmt.Println("Student deleted successfully!")
}

// ---------- Step 8: Menu Loop ----------

func showMenu() {
	fmt.Println("\n===== Student Manager =====")
	fmt.Println("1. Add Student")
	fmt.Println("2. List Students")
	fmt.Println("3. Find Student")
	fmt.Println("4. Update Student")
	fmt.Println("5. Delete Student")
	fmt.Println("6. Exit")
}

func getChoice() int {
	return readInt("Enter your choice: ")
}

func main() {
	for {
		showMenu()
		choice := getChoice()

		switch choice {
		case 1:
			addStudent()
		case 2:
			listStudents()
		case 3:
			findStudent()
		case 4:
			updateStudent()
		case 5:
			deleteStudent()
		case 6:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}
