package parser_test

import (
	"fmt"
	"io"
	"os"
	"parser"
	"parser/hcl"
	"parser/tokens"
	"parser/xml"
	"parser/yaml"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestParser(t *testing.T) {
	root := "./test_files"

	hclToYaml := parser.NewParser(hcl.HCL{}, yaml.YAML{})
	hclToXML := parser.NewParser(hcl.HCL{}, xml.XML{})

	err := filepath.WalkDir(root, func(path string, entry os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if entry.IsDir() || strings.ToLower(filepath.Ext(path)) != ".hcl" {
			return nil
		}

		yamlPath := strings.TrimSuffix(path, filepath.Ext(path)) + ".yaml"
		xmlPath := strings.TrimSuffix(path, filepath.Ext(path)) + ".xml"

		hclFile, err := os.Open(path)
		if err != nil {
			return err
		}
		defer hclFile.Close()

		yamlFile, err := os.OpenFile(yamlPath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		defer yamlFile.Close()

		xmlFile, err := os.OpenFile(xmlPath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		defer xmlFile.Close()

		if err := hclToYaml.Parse(hclFile, yamlFile); err != nil {
			return err
		}

		_, err = hclFile.Seek(0, io.SeekStart)
		if err != nil {
			panic(err)
		}

		if err := hclToXML.Parse(hclFile, xmlFile); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		t.Fatal(err)
	}
}

func TestAST(t *testing.T) {
	f, err := os.Open("./test_files/test2.hcl")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	lexer := hcl.NewLexer(f)

	tk := make([]*tokens.Token, 0)
	for token, err := range lexer.Tokens() {
		if err != nil {
			t.Fatal(err)
		}

		if token.Type() != tokens.TokenTypeComment {
			tk = append(tk, token)
		}
	}

	parser := hcl.NewParser(tk)
	file, err := parser.Parse()
	if err != nil {
		t.Fatal(err)
	}

	printAST(file)
}

func printAST(file *hcl.File) {
	var printNode func(node hcl.Node, depth int)

	printNode = func(node hcl.Node, depth int) {
		indent := strings.Repeat(" ", depth*2)

		switch n := node.(type) {
		case *hcl.Block:
			fmt.Printf("%sBlock: %s\n", indent, n.Name)

			for _, label := range n.Labels {
				fmt.Printf("%s  Label: %q\n", indent, label)
			}

			for _, child := range n.Nodes {
				printNode(child, depth+1)
			}

		case *hcl.Attribute:
			fmt.Printf("%sAttribute: %s\n", indent, n.Name)
			printNode(n.Value, depth+1)

		case *hcl.String:
			fmt.Printf("%sString: %q\n", indent, string(*n))

		case *hcl.Number:
			fmt.Printf("%sNumber: %s\n", indent, string(*n))

		case *hcl.Boolean:
			fmt.Printf("%sBoolean: %t\n", indent, bool(*n))

		case *hcl.Null:
			fmt.Printf("%sNull\n", indent)

		case *hcl.Expression:
			fmt.Printf("%sExpression: %s\n", indent, string(*n))

		case *hcl.Array:
			fmt.Printf("%sArray\n", indent)
			for _, element := range *n {
				printNode(element, depth+1)
			}

		case *hcl.Object:
			fmt.Printf("%sObject\n", indent)
			for _, attribute := range *n {
				printNode(attribute, depth+1)
			}
		}
	}

	fmt.Println("File")

	for _, node := range file.Nodes {
		printNode(node, 1)
	}
}

func TestParse100(t *testing.T) {
	hclToYaml := parser.NewParser(hcl.HCL{}, xml.XML{})
	hclFile, err := os.Open("./test_files/schedule.hcl")
	if err != nil {
		t.Fatal(err)
	}
	defer hclFile.Close()

	start := time.Now()

	for range 100 {
		_, err = hclFile.Seek(0, io.SeekStart)
		if err != nil {
			panic(err)
		}

		w := io.Discard
		if err := hclToYaml.Parse(hclFile, w); err != nil {
			t.Fatal(err)
		}
	}

	t.Log(time.Since(start))
}
