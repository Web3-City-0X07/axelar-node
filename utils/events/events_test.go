package events_test

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path"
	"reflect"
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/axelarnetwork/utils/funcs"
	"github.com/axelarnetwork/utils/slices"
)

// This test is dependent on type packages for all modules being loaded in event_imports.go.
// It will panic if an event is found but the respective module is not loaded.
func TestEventsAreMarshallable(t *testing.T) {
	eventFiles := findEventFiles()
	for _, eventFile := range eventFiles {
		names, err := parseEventNames(eventFile)
		assert.NoError(t, err)
		assert.NotEmpty(t, names, eventFile)

		for _, name := range names {
			msgType := proto.MessageType(name)
			require.NotNilf(t, msgType, "event %s not registered, make sure the corresponding module is imported in event_imports.go", name)
			event := reflect.New(msgType.Elem()).Interface().(proto.Message)

			abciEvent, err := sdk.TypedEventToEvent(event)
			assert.NoError(t, err)

			event2, err := sdk.ParseTypedEvent(abci.Event(abciEvent))
			assert.NoError(t, err)
			assert.Equal(t, funcs.Must(json.Marshal(event)), funcs.Must(json.Marshal(event2)))
		}
	}
}

// by convention events will be generated by the protobuf compiler in x/{module}/types/events.pb.go,
// so this checks all potential event file locations and returns all existing ones
func findEventFiles() []string {
	moduleDir := "../../x"
	modules := slices.Filter(funcs.Must(os.ReadDir(moduleDir)), os.DirEntry.IsDir)

	potentialFiles := slices.Map(modules, func(module os.DirEntry) string {
		return path.Join(moduleDir, module.Name(), "types/events.pb.go")
	})

	return slices.Filter(potentialFiles, func(file string) bool {
		_, err := os.Stat(file)
		if os.IsNotExist(err) {
			fmt.Printf("no events found in %s\n", file)
		}
		if err != nil {
			return false
		}
		return true
	})
}

// events are registered with the proto package by their fully qualified names in the init() function
// of the generated events.pb.go file. Therefore, we need to parse the qualified names from the generated file
// to later retrieve the actual event instances from the proto package. To do that we traverse the file's abstract syntax tree (AST).
func parseEventNames(file string) ([]string, error) {
	root, err := findRootNode(file)
	if err != nil {
		return nil, err
	}

	inits := findInitFuncs(root)
	registerTypeCalls := slices.FlatMap(inits, findRegisterTypeCalls)
	eventNames := slices.FlatMap(registerTypeCalls, parseEventNamesFromArgs)

	return eventNames, nil
}

func findRootNode(file string) (*ast.File, error) {
	fs := token.NewFileSet()
	f, err := parser.ParseFile(fs, file, nil, parser.AllErrors)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func findInitFuncs(f *ast.File) []*ast.FuncDecl {
	fns := slices.TryCast[ast.Decl, *ast.FuncDecl](f.Decls)
	return slices.Filter(fns, func(decl *ast.FuncDecl) bool { return decl.Name.Name == "init" })
}

func findRegisterTypeCalls(init *ast.FuncDecl) []*ast.CallExpr {
	initStatements := slices.TryCast[ast.Stmt, *ast.ExprStmt](init.Body.List)
	initExprStatements := slices.Map(initStatements, func(stmt *ast.ExprStmt) ast.Expr { return stmt.X })
	initExpressions := slices.TryCast[ast.Expr, *ast.CallExpr](initExprStatements)

	registerTypeCalls := slices.Filter(initExpressions, func(expr *ast.CallExpr) bool {
		switch f := expr.Fun.(type) {
		case *ast.SelectorExpr:
			return f.Sel.Name == "RegisterType"
		default:
			return false
		}
	})
	return registerTypeCalls
}

func parseEventNamesFromArgs(registerTypeCalls *ast.CallExpr) []string {
	eventNameArgs := slices.TryCast[ast.Expr, *ast.BasicLit](registerTypeCalls.Args)

	eventNames := slices.Map(eventNameArgs, func(arg *ast.BasicLit) string {
		eventName := funcs.Must(strconv.Unquote(arg.Value))
		fmt.Printf("found event %s\n", eventName)
		return eventName
	})
	return eventNames
}
