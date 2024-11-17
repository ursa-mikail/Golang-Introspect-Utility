# Golang Introspect Utility

## Overview

This utility provides a way to add JSON-style comments in Golang files and introspect them similarly to how Javadoc works in Java. The idea is to insert metadata in JSON format within the Golang file and then use a Golang code inspect lib (./lib/code_inspect/code_introspect.go) to read and display this metadata in an interactive manner.

## Files

1. **example.go**  
   This Golang file contains the code along with JSON-style comments at the top, describing the module, author, version, and functions.

2. **code_introspect.go**  
   A script that reads and introspects the JSON-style comment in a Golang file. It parses the JSON metadata and presents it interactively to the user.

## Features

- **JSON-style Comments**:  
  Allows you to add metadata (e.g., module description, author, version, function details) in a structured JSON format within Golang comments.
  
- **Introspection**:  
  Extracts the JSON-style comment and presents it to the user in an interactive menu format, similar to Javadoc.
  
- **Interactive Menu**:  
  Provides a menu for the user to display all fields or select specific fields to view, making it easy to introspect the metadata.

## Example

### example.go
```Golang
/*
	{
	    "name": "ExampleFunction",
	    "description": "This is an example function.",
	    "parameters": {
	        "param1": "This is the first parameter.",
	        "param2": "This is the second parameter."
	    },
	    "returns": "This function returns nothing."
	}
*/

func ExampleFunction(param1 string, param2 int) {
	fmt.Println("Example function executed.")
}

```

```
"""
Menu:
1. Display all fields
2. Choose specific fields to display
3. Exit
Enter your choice: 2
Available fields to display:
1. description
2. author
3. version
4. functions
Enter the numbers of the fields you want to display, separated by commas (e.g., 1,2): 1
{
    "description": "This module provides an example function."
}

Menu:
1. Display all fields
2. Choose specific fields to display
3. Exit
Enter your choice: 3
"""
```

code_introspect.go is 

1.Highly customizable
2.Interactive                       
3.| Requires manual maintenance for json formatting (or tools to be extended)|


For customized and interactive needs, the code_introspect.go provides a flexible, tailored solution. It can be extended with:

    1.Auto-document with Cross-referencing and Indexing.
    2.Version Control Integration (e.g. linking with git references)
    3.Search Functionality: In-built-in search functionality on both code and generated documentation.
