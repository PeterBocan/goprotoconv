package goprotoconv

import (
	"go/ast"
	"go/parser"
	"go/token"
)

// Type is an interface which coallesces different types
type Type interface {
	isProtobufType()
}

// ProtobufField specifies a field of a protobuf message
type ProtobufField struct {
	Name string
	Type string
}

// ProtobufStruct represents
type ProtobufStruct struct {
	Name   string
	Fields []ProtobufField
}

func (t ProtobufStruct) isProtobufType() {}

// ProtobufFile represents a list of all symbols generated
// by Proto compiler.
type ProtobufFile struct {
	Filepath        string
	PackageName     string
	TypeDefinitions []Type
	Imports         map[string]string

	syntaxTree   *ast.File
	tokenFileSet *token.FileSet
}

func (file *ProtobufFile) loadImports() {
	for _, importSpec := range file.syntaxTree.Imports {
		path := importSpec.Path.Value
		identifier := importSpec.Name.Name
		file.Imports[identifier] = path
	}
}

func (file ProtobufFile) loadStructFields(s *ProtobufStruct, fieldList *ast.FieldList) {
	// for _, field := range fieldList.List {
	// }
}

func (file *ProtobufFile) loadStructDefinitions() {
	file.PrintAST()
	for _, declaration := range file.syntaxTree.Decls {
		switch genDecl := declaration.(type) {
		case *ast.GenDecl:
			// we are interested in type declarations only i.e. type Something struct {}
			// not sure if len(genDecl.Specs) == 1 is sufficient
			if genDecl.Tok == token.TYPE && len(genDecl.Specs) == 1 {
				for _, spec := range genDecl.Specs {
					// we are looking for specific types
					if typeSpec, ok := spec.(*ast.TypeSpec); ok {
						if typeStruct, ok := typeSpec.Type.(*ast.StructType); ok {
							protobufMessage := ProtobufStruct{
								Name:   typeSpec.Name.Name,
								Fields: []ProtobufField{},
							}
							// pull out all fields in the struct
							file.loadStructFields(&protobufMessage, typeStruct.Fields)
							file.TypeDefinitions = append(file.TypeDefinitions, protobufMessage)
						}
					}
				}
			}
			break
		}
	}
}

func (file *ProtobufFile) loadAllSymbols() {
	file.loadImports()
	file.loadStructDefinitions()
}

// LoadProtobufGoFile loads generated Protobuf Go file into the AST.
func LoadProtobufGoFile(filepath string) (*ProtobufFile, error) {
	tokenFileSet := token.NewFileSet()
	syntaxTreeRepresentation, err := parser.ParseFile(tokenFileSet, filepath, nil, parser.SpuriousErrors)
	if err != nil {
		return nil, err
	}

	file := &ProtobufFile{
		Filepath:        filepath,
		syntaxTree:      syntaxTreeRepresentation,
		tokenFileSet:    tokenFileSet,
		Imports:         map[string]string{},
		TypeDefinitions: []Type{},
	}

	file.PackageName = syntaxTreeRepresentation.Name.Name
	file.loadAllSymbols()
	return file, nil
}

// PrintAST prints an AST of the loaded file
func (file ProtobufFile) PrintAST() {
	ast.Print(file.tokenFileSet, file.syntaxTree)
}
