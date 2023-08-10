package utils

import (
	"fmt"
	"go/parser"
	"go/token"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/javascript"
)

// Use pdftotext (based on poppler-utils).
// Other version based on Apache Tika exists in tools
// func TestPdfParsing(t *testing.T) {
// 	res, err := docconv.ConvertPath("test_input/parsing/pdf_test.pdf")
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	fmt.Println(res)
// }

func TestFileRead(t *testing.T) {
	data, err := os.ReadFile("test_input/parsing/text_test.txt")
	if err != nil {
		t.Error(err)
	}
	assert.Contains(t, string(data), "coucou", "The text file should contain coucou")
	t.Log(string(data))
}

func TestGoParse(t *testing.T) {
	fset := token.NewFileSet() // positions are relative to fset

	src := `package foo

import (
	"fmt"
	"time"
)

func bar() {
	fmt.Println(time.Now())
}`

	// Parse src but stop after processing the imports.
	f, err := parser.ParseFile(fset, "", src, parser.ImportsOnly)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the imports from the file's AST.
	for _, s := range f.Imports {
		fmt.Println(s.Path.Value)
	}

}

func TestTreeSitterParse(t *testing.T) {
	parser := sitter.NewParser()
	parser.SetLanguage(javascript.GetLanguage())
	sourceCode := []byte("let a = 1")
	tree := parser.Parse(nil, sourceCode)
	n := tree.RootNode()

	fmt.Println(n) // (program (lexical_declaration (variable_declarator (identifier) (number))))

	child := n.NamedChild(0)
	fmt.Println(child.Type())      // lexical_declaration
	fmt.Println(child.StartByte()) // 0
	fmt.Println(child.EndByte())   // 9

	// change 1 -> true
	//newText := []byte("let a = true")
	tree.Edit(sitter.EditInput{
		StartIndex:  8,
		OldEndIndex: 9,
		NewEndIndex: 12,
		StartPoint: sitter.Point{
			Row:    0,
			Column: 8,
		},
		OldEndPoint: sitter.Point{
			Row:    0,
			Column: 9,
		},
		NewEndPoint: sitter.Point{
			Row:    0,
			Column: 12,
		},
	})

	// check that it changed tree
	assert.True(t, n.HasChanges())
	assert.True(t, n.Child(0).HasChanges())
	assert.False(t, n.Child(0).Child(0).HasChanges()) // left side of the tree didn't change
	assert.True(t, n.Child(0).Child(1).HasChanges())

	// generate new tree
	//newTree := parser.Parse(tree, newText)
}
