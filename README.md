# **Task Manager CLI**

![Go Version](https://img.shields.io/github/go-mod/go-version/github/ms-alejo/tasks)  
A simple and intuitive Command-Line Interface (CLI) for managing tasks efficiently, built using the [Cobra](https://github.com/spf13/cobra) library in Go. This tool allows you to add, list, complete, and delete tasks from the terminal.

---

## **Features**

- **Add Tasks**: Quickly add tasks with a description.
- **List Tasks**: View all uncompleted tasks or include completed ones with a flag.
- **Complete Tasks**: Mark tasks as complete using their unique ID.
- **Delete Tasks**: Remove tasks from the list permanently.

---

## **Installation**

### **1. Clone the Repository**
```bash
git clone https://github.com/ms-alejo/tasks.git
cd tasks
```

### **2. Install Dependencies**
Ensure you have Go installed (version 1.18 or later).

```bash
go mod tidy
```

### **3. Build the Application**
```bash
go build -o tasks
```

### **4. Run the Application**
```bash
./tasks
```

---

## **Usage**

This CLI tool supports the following commands:

### **1. Add a Task**
Add a new task with a description:
```bash
./tasks add "Tidy my desk"
```

### **2. List Tasks**
List all uncompleted tasks:
```bash
./tasks list
```

Include completed tasks with the `--all` flag:
```bash
./tasks list --all
```

### **3. Mark a Task as Complete**
Mark a task as complete by its ID:
```bash
./tasks complete 1
```

### **4. Delete a Task**
Delete a task by its ID:
```bash
./tasks delete 1
```

---

## **Examples**

### **Adding Tasks**
```bash
./tasks add "Organize my files"
./tasks add "Write project documentation"
```

### **Listing Tasks**
```bash
./tasks list
```
Output:
```
ID      Task                          Created
1       Organize my files             2 minutes ago
2       Write project documentation   1 minute ago
```

### **Marking a Task as Complete**
```bash
./tasks complete 1
```

### **Deleting a Task**
```bash
./tasks delete 2
```
