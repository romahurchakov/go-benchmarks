package data

// 171 B
var SmallJSON = `
{
	"menu": {
	"id": "file",
	"value": "File",
	"popup": {
	  "menuitem": [
		{"value": "New", "onclick": "CreateNewDoc()"},
		{"value": "Open", "onclick": "OpenDoc()"},
		{"value": "Close", "onclick": "CloseDoc()"}
	  ]
	}
  }
}
`