# Shopping list web app from Go

A simple Go-based web application that allows users to add and view shopping list items through a clean and easy-to-use interface.

This project uses Go, HTML templates, CSS, and a lightweight built-in web server. It demonstrates routing, form handling, template rendering, and serving static files.

___

## Features

- Add shopping list items
- Show all items in a dedicated page
- Clean, styled user interface

___

## Product Structure 

```
/project
│
├── main.go
├── go.mod
│
├── /template
│   ├── home.html
│   ├── add.html
│   └── showall.html   
│
└── /static
    └── style.css
```

___

# How it works

1. Home Page
- Add  → Go to add item page
- Show All List → Show all stored items

2. Add Page

User can:
- Type a shopping item
- Submit it
- Page updates to show the newly added item
- Option to delete it , return back the home page or go to show all list. 

3. Show All Page

Displays:
- All items in the shopping list
- Message saying “No items yet” if the list is empty