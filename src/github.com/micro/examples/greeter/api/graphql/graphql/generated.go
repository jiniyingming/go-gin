// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graphql

import (
	"bytes"
	context "context"
	strconv "strconv"
	sync "sync"

	graphql "github.com/99designs/gqlgen/graphql"
	introspection "github.com/99designs/gqlgen/graphql/introspection"
	go_micro_srv_greeter "github.com/micro/examples/greeter/srv/proto/hello"
	gqlparser "github.com/vektah/gqlparser"
	ast "github.com/vektah/gqlparser/ast"
)

// NewExecutableSchema creates an ExecutableSchema from the ResolverRoot interface.
func NewExecutableSchema(cfg Config) graphql.ExecutableSchema {
	return &executableSchema{
		resolvers:  cfg.Resolvers,
		directives: cfg.Directives,
	}
}

type Config struct {
	Resolvers  ResolverRoot
	Directives DirectiveRoot
}

type ResolverRoot interface {
	Query() QueryResolver
}

type DirectiveRoot struct {
}
type QueryResolver interface {
	Hello(ctx context.Context, name string) (*go_micro_srv_greeter.Response, error)
}

type executableSchema struct {
	resolvers  ResolverRoot
	directives DirectiveRoot
}

func (e *executableSchema) Schema() *ast.Schema {
	return parsedSchema
}

func (e *executableSchema) Query(ctx context.Context, op *ast.OperationDefinition) *graphql.Response {
	ec := executionContext{graphql.GetRequestContext(ctx), e}

	buf := ec.RequestMiddleware(ctx, func(ctx context.Context) []byte {
		data := ec._Query(ctx, op.SelectionSet)
		var buf bytes.Buffer
		data.MarshalGQL(&buf)
		return buf.Bytes()
	})

	return &graphql.Response{
		Data:   buf,
		Errors: ec.Errors,
	}
}

func (e *executableSchema) Mutation(ctx context.Context, op *ast.OperationDefinition) *graphql.Response {
	return graphql.ErrorResponse(ctx, "mutations are not supported")
}

func (e *executableSchema) Subscription(ctx context.Context, op *ast.OperationDefinition) func() *graphql.Response {
	return graphql.OneShot(graphql.ErrorResponse(ctx, "subscriptions are not supported"))
}

type executionContext struct {
	*graphql.RequestContext
	*executableSchema
}

var helloResponseImplementors = []string{"HelloResponse"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) _HelloResponse(ctx context.Context, sel ast.SelectionSet, obj *go_micro_srv_greeter.Response) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, helloResponseImplementors)

	out := graphql.NewOrderedMap(len(fields))
	invalid := false
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("HelloResponse")
		case "msg":
			out.Values[i] = ec._HelloResponse_msg(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	if invalid {
		return graphql.Null
	}
	return out
}

// nolint: vetshadow
func (ec *executionContext) _HelloResponse_msg(ctx context.Context, field graphql.CollectedField, obj *go_micro_srv_greeter.Response) graphql.Marshaler {
	rctx := &graphql.ResolverContext{
		Object: "HelloResponse",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(ctx context.Context) (interface{}, error) {
		return obj.Msg, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	return graphql.MarshalString(res)
}

var queryImplementors = []string{"Query"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) _Query(ctx context.Context, sel ast.SelectionSet) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, queryImplementors)

	ctx = graphql.WithResolverContext(ctx, &graphql.ResolverContext{
		Object: "Query",
	})

	var wg sync.WaitGroup
	out := graphql.NewOrderedMap(len(fields))
	invalid := false
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Query")
		case "hello":
			wg.Add(1)
			go func(i int, field graphql.CollectedField) {
				out.Values[i] = ec._Query_hello(ctx, field)
				wg.Done()
			}(i, field)
		case "__type":
			out.Values[i] = ec._Query___type(ctx, field)
		case "__schema":
			out.Values[i] = ec._Query___schema(ctx, field)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	wg.Wait()
	if invalid {
		return graphql.Null
	}
	return out
}

// nolint: vetshadow
func (ec *executionContext) _Query_hello(ctx context.Context, field graphql.CollectedField) graphql.Marshaler {
	rawArgs := field.ArgumentMap(ec.Variables)
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["name"]; ok {
		var err error
		arg0, err = graphql.UnmarshalString(tmp)
		if err != nil {
			ec.Error(ctx, err)
			return graphql.Null
		}
	}
	args["name"] = arg0
	rctx := &graphql.ResolverContext{
		Object: "Query",
		Args:   args,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, nil, func(ctx context.Context) (interface{}, error) {
		return ec.resolvers.Query().Hello(ctx, args["name"].(string))
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*go_micro_srv_greeter.Response)
	rctx.Result = res

	if res == nil {
		return graphql.Null
	}

	return ec._HelloResponse(ctx, field.Selections, res)
}

// nolint: vetshadow
func (ec *executionContext) _Query___type(ctx context.Context, field graphql.CollectedField) graphql.Marshaler {
	rawArgs := field.ArgumentMap(ec.Variables)
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["name"]; ok {
		var err error
		arg0, err = graphql.UnmarshalString(tmp)
		if err != nil {
			ec.Error(ctx, err)
			return graphql.Null
		}
	}
	args["name"] = arg0
	rctx := &graphql.ResolverContext{
		Object: "Query",
		Args:   args,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, nil, func(ctx context.Context) (interface{}, error) {
		return ec.introspectType(args["name"].(string)), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	rctx.Result = res

	if res == nil {
		return graphql.Null
	}

	return ec.___Type(ctx, field.Selections, res)
}

// nolint: vetshadow
func (ec *executionContext) _Query___schema(ctx context.Context, field graphql.CollectedField) graphql.Marshaler {
	rctx := &graphql.ResolverContext{
		Object: "Query",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, nil, func(ctx context.Context) (interface{}, error) {
		return ec.introspectSchema(), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection.Schema)
	rctx.Result = res

	if res == nil {
		return graphql.Null
	}

	return ec.___Schema(ctx, field.Selections, res)
}

var __DirectiveImplementors = []string{"__Directive"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___Directive(ctx context.Context, sel ast.SelectionSet, obj *introspection.Directive) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, __DirectiveImplementors)

	out := graphql.NewOrderedMap(len(fields))
	invalid := false
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__Directive")
		case "name":
			out.Values[i] = ec.___Directive_name(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "description":
			out.Values[i] = ec.___Directive_description(ctx, field, obj)
		case "locations":
			out.Values[i] = ec.___Directive_locations(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "args":
			out.Values[i] = ec.___Directive_args(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	if invalid {
		return graphql.Null
	}
	return out
}

// nolint: vetshadow
func (ec *executionContext) ___Directive_name(ctx context.Context, field graphql.CollectedField, obj *introspection.Directive) graphql.Marshaler {
	rctx := &graphql.ResolverContext{
		Object: "__Directive",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(ctx context.Context) (interface{}, error) {
		return obj.Name, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	return graphql.MarshalString(res)
}

// nolint: vetshadow
func (ec *executionContext) ___Directive_description(ctx context.Context, field graphql.CollectedField, obj *introspection.Directive) graphql.Marshaler {
	rctx := &graphql.ResolverContext{
		Object: "__Directive",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(ctx context.Context) (interface{}, error) {
		return obj.Description, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	return graphql.MarshalString(res)
}

// nolint: vetshadow
func (ec *executionContext) ___Directive_locations(ctx context.Context, field graphql.CollectedField, obj *introspection.Directive) graphql.Marshaler {
	rctx := &graphql.ResolverContext{
		Object: "__Directive",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(ctx context.Context) (interface{}, error) {
		return obj.Locations, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]string)
	rctx.Result = res
	arr1 := graphql.Array{}
	for idx1 := range res {
		arr1 = append(arr1, func() graphql.Marshaler {
			return graphql.MarshalString(res[idx1])
		}())
	}
	return arr1
}

// nolint: vetshadow
func (ec *executionContext) ___Directive_args(ctx context.Context, field graphql.CollectedField, obj *introspection.Directive) graphql.Marshaler {
	rctx := &graphql.ResolverContext{
		Object: "__Directive",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(ctx context.Context) (interface{}, error) {
		return obj.Args, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]introspection.InputValue)
	rctx.Result = res
	arr1 := graphql.Array{}
	for idx1 := range res {
		rctx := &graphql.ResolverContext{
			Index:  &idx1,
			Result: &res[idx1],
		}
		ctx := graphql.WithResolverContext(ctx, rctx)
		arr1 = append(arr1, func() graphql.Marshaler {

			return ec.___InputValue(ctx, field.Selections, &res[idx1])
		}())
	}
	return arr1
}

var __EnumValueImplementors = []string{"__EnumValue"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___EnumValue(ctx context.Context, sel ast.SelectionSet, obj *introspection.EnumValue) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, __EnumValueImplementors)

	out := graphql.NewOrderedMap(len(fields))
	invalid := false
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__EnumValue")
		case "name":
			out.Values[i] = ec.___EnumValue_name(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "description":
			out.Values[i] = ec.___EnumValue_description(ctx, field, obj)
		case "isDeprecated":
			out.Values[i] = ec.___EnumValue_isDeprecated(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "deprecationReason":
			out.Values[i] = ec.___EnumValue_deprecationReason(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	if invalid {
		return graphql.Null
	}
	return out
}

// nolint: vetshadow
func (ec *executionContext) ___EnumValue_name(ctx context.Context, field graphql.CollectedField, obj *introspection.EnumValue) graphql.Marshaler {
	rctx := &graphql.ResolverContext{
		Object: "__EnumValue",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(ctx context.Context) (interface{}, error) {
		return obj.Name, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	return graphql.MarshalString(res)
}

// nolint: vetshadow
func (ec *executionContext) ___EnumValue_description(ctx context.Context, field graphql.CollectedField, obj *introspection.EnumValue) graphql.Marshaler {
	rctx := &graphql.ResolverContext{
		Object: "__EnumValue",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(ctx context.Context) (interface{}, error) {
		return obj.Description, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	return graphql.MarshalString(res)
}

// nolint: vetshadow
func (ec *executionContext) ___EnumValue_isDeprecated(ctx context.Context, field graphql.CollectedField, obj *introspection.EnumValue) graphql.Marshaler {
	rctx := &graphql.ResolverContext{
		Object: "__EnumValue",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(ctx context.Context) (interface{}, error) {
		return obj.IsDeprecated, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(bool)
	rctx.Result = res
	return graphql.MarshalBoolean(res)
}

// nolint: vetshadow
func (ec *executionContext) ___EnumValue_deprecationReason(ctx context.Context, field graphql.CollectedField, obj *introspection.EnumValue) graphql.Marshaler {
	rctx := &graphql.ResolverContext{
		Object: "__EnumValue",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(ctx context.Context) (interface{}, error) {
		return obj.DeprecationReason, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	return graphql.MarshalString(res)
}

var __FieldImplementors = []string{"__Field"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___Field(ctx context.Context, sel ast.SelectionSet, obj *introspection.Field) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, __FieldImplementors)

	out := graphql.NewOrderedMap(len(fields))
	invalid := false
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__Field")
		case "name":
			out.Values[i] = ec.___Field_name(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "description":
			out.Values[i] = ec.___Field_description(ctx, field, obj)
		case "args":
			out.Values[i] = ec.___Field_args(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "type":
			out.Values[i] = ec.___Field_type(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "isDeprecated":
			out.Values[i] = ec.___Field_isDeprecated(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "deprecationReason":
			out.Values[i] = ec.___Field_deprecationReason(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	if invalid {
		return graphql.Null
	}
	return out
}

// nolint: vetshadow
func (ec *executionContext) ___Field_name(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) graphql.Marshaler {
	rctx := &graphql.ResolverContext{
		Object: "__Field",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(ctx context.Context) (interface{}, error) {
		return obj.Name, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	return graphql.MarshalString(res)
}

// nolint: vetshadow
func (ec *executionContext) ___Field_description(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) graphql.Marshaler {
	rctx := &graphql.ResolverContext{
		Object: "__Field",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(ctx context.Context) (interface{}, error) {
		return obj.Description, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	return graphql.MarshalString(res)
}

// nolint: vetshadow
func (ec *executionContext) ___Field_args(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) graphql.Marshaler {
	rctx := &graphql.ResolverContext{
		Object: "__Field",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(ctx context.Context) (interface{}, error) {
		return obj.Args, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]introspection.InputValue)
	rctx.Result = res
	arr1 := graphql.Array{}
	for idx1 := range res {
		rctx := &graphql.ResolverContext{
			Index:  &idx1,
			Result: &res[idx1],
		}
		ctx := graphql.WithResolverContext(ctx, rctx)
		arr1 = append(arr1, func() graphql.Marshaler {

			return ec.___InputValue(ctx, field.Selections, &res[idx1])
		}())
	}
	return arr1
}

// nolint: vetshadow
func (ec *executionContext) ___Field_type(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) graphql.Marshaler {
	rctx := &graphql.ResolverContext{
		Object: "__Field",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(ctx context.Context) (interface{}, error) {
		return obj.Type, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	rctx.Result = res

	if res == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}

	return ec.___Type(ctx, field.Selections, res)
}

// nolint: vetshadow
func (ec *executionContext) ___Field_isDeprecated(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) graphql.Marshaler {
	rctx := &graphql.ResolverContext{
		Object: "__Field",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(ctx context.Context) (interface{}, error) {
		return obj.IsDeprecated, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(bool)
	rctx.Result = res
	return graphql.MarshalBoolean(res)
}

// nolint: vetshadow
func (ec *executionContext) ___Field_deprecationReason(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) graphql.Marshaler {
	rctx := &graphql.ResolverContext{
		Object: "__Field",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(ctx context.Context) (interface{}, error) {
		return obj.DeprecationReason, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	return graphql.MarshalString(res)
}

var __InputValueImplementors = []string{"__InputValue"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___InputValue(ctx context.Context, sel ast.SelectionSet, obj *introspection.InputValue) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, __InputValueImplementors)

	out := graphql.NewOrderedMap(len(fields))
	invalid := false
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__InputValue")
		case "name":
			out.Values[i] = ec.___InputValue_name(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "description":
			out.Values[i] = ec.___InputValue_description(ctx, field, obj)
		case "type":
			out.Values[i] = ec.___InputValue_type(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "defaultValue":
			out.Values[i] = ec.___InputValue_defaultValue(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	if invalid {
		return graphql.Null
	}
	return out
}

// nolint: vetshadow
func (ec *executionContext) ___InputValue_name(ctx context.Context, field graphql.CollectedField, obj *introspection.InputValue) graphql.Marshaler {
	rctx := &graphql.ResolverContext{
		Object: "__InputValue",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(ctx context.Context) (interface{}, error) {
		return obj.Name, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	return graphql.MarshalString(res)
}

// nolint: vetshadow
func (ec *executionContext) ___InputValue_description(ctx context.Context, field graphql.CollectedField, obj *introspection.InputValue) graphql.Marshaler {
	rctx := &graphql.ResolverContext{
		Object: "__InputValue",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(ctx context.Context) (interface{}, error) {
		return obj.Description, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	return graphql.MarshalString(res)
}

// nolint: vetshadow
func (ec *executionContext) ___InputValue_type(ctx context.Context, field graphql.CollectedField, obj *introspection.InputValue) graphql.Marshaler {
	rctx := &graphql.ResolverContext{
		Object: "__InputValue",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(ctx context.Context) (interface{}, error) {
		return obj.Type, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	rctx.Result = res

	if res == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}

	return ec.___Type(ctx, field.Selections, res)
}

// nolint: vetshadow
func (ec *executionContext) ___InputValue_defaultValue(ctx context.Context, field graphql.CollectedField, obj *introspection.InputValue) graphql.Marshaler {
	rctx := &graphql.ResolverContext{
		Object: "__InputValue",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(ctx context.Context) (interface{}, error) {
		return obj.DefaultValue, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	rctx.Result = res

	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

var __SchemaImplementors = []string{"__Schema"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___Schema(ctx context.Context, sel ast.SelectionSet, obj *introspection.Schema) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, __SchemaImplementors)

	out := graphql.NewOrderedMap(len(fields))
	invalid := false
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__Schema")
		case "types":
			out.Values[i] = ec.___Schema_types(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "queryType":
			out.Values[i] = ec.___Schema_queryType(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "mutationType":
			out.Values[i] = ec.___Schema_mutationType(ctx, field, obj)
		case "subscriptionType":
			out.Values[i] = ec.___Schema_subscriptionType(ctx, field, obj)
		case "directives":
			out.Values[i] = ec.___Schema_directives(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	if invalid {
		return graphql.Null
	}
	return out
}

// nolint: vetshadow
func (ec *executionContext) ___Schema_types(ctx context.Context, field graphql.CollectedField, obj *introspection.Schema) graphql.Marshaler {
	rctx := &graphql.ResolverContext{
		Object: "__Schema",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(ctx context.Context) (interface{}, error) {
		return obj.Types(), nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]introspection.Type)
	rctx.Result = res
	arr1 := graphql.Array{}
	for idx1 := range res {
		rctx := &graphql.ResolverContext{
			Index:  &idx1,
			Result: &res[idx1],
		}
		ctx := graphql.WithResolverContext(ctx, rctx)
		arr1 = append(arr1, func() graphql.Marshaler {

			return ec.___Type(ctx, field.Selections, &res[idx1])
		}())
	}
	return arr1
}

// nolint: vetshadow
func (ec *executionContext) ___Schema_queryType(ctx context.Context, field graphql.CollectedField, obj *introspection.Schema) graphql.Marshaler {
	rctx := &graphql.ResolverContext{
		Object: "__Schema",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(ctx context.Context) (interface{}, error) {
		return obj.QueryType(), nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	rctx.Result = res

	if res == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}

	return ec.___Type(ctx, field.Selections, res)
}

// nolint: vetshadow
func (ec *executionContext) ___Schema_mutationType(ctx context.Context, field graphql.CollectedField, obj *introspection.Schema) graphql.Marshaler {
	rctx := &graphql.ResolverContext{
		Object: "__Schema",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(ctx context.Context) (interface{}, error) {
		return obj.MutationType(), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	rctx.Result = res

	if res == nil {
		return graphql.Null
	}

	return ec.___Type(ctx, field.Selections, res)
}

// nolint: vetshadow
func (ec *executionContext) ___Schema_subscriptionType(ctx context.Context, field graphql.CollectedField, obj *introspection.Schema) graphql.Marshaler {
	rctx := &graphql.ResolverContext{
		Object: "__Schema",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(ctx context.Context) (interface{}, error) {
		return obj.SubscriptionType(), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	rctx.Result = res

	if res == nil {
		return graphql.Null
	}

	return ec.___Type(ctx, field.Selections, res)
}

// nolint: vetshadow
func (ec *executionContext) ___Schema_directives(ctx context.Context, field graphql.CollectedField, obj *introspection.Schema) graphql.Marshaler {
	rctx := &graphql.ResolverContext{
		Object: "__Schema",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(ctx context.Context) (interface{}, error) {
		return obj.Directives(), nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]introspection.Directive)
	rctx.Result = res
	arr1 := graphql.Array{}
	for idx1 := range res {
		rctx := &graphql.ResolverContext{
			Index:  &idx1,
			Result: &res[idx1],
		}
		ctx := graphql.WithResolverContext(ctx, rctx)
		arr1 = append(arr1, func() graphql.Marshaler {

			return ec.___Directive(ctx, field.Selections, &res[idx1])
		}())
	}
	return arr1
}

var __TypeImplementors = []string{"__Type"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___Type(ctx context.Context, sel ast.SelectionSet, obj *introspection.Type) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, __TypeImplementors)

	out := graphql.NewOrderedMap(len(fields))
	invalid := false
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__Type")
		case "kind":
			out.Values[i] = ec.___Type_kind(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "name":
			out.Values[i] = ec.___Type_name(ctx, field, obj)
		case "description":
			out.Values[i] = ec.___Type_description(ctx, field, obj)
		case "fields":
			out.Values[i] = ec.___Type_fields(ctx, field, obj)
		case "interfaces":
			out.Values[i] = ec.___Type_interfaces(ctx, field, obj)
		case "possibleTypes":
			out.Values[i] = ec.___Type_possibleTypes(ctx, field, obj)
		case "enumValues":
			out.Values[i] = ec.___Type_enumValues(ctx, field, obj)
		case "inputFields":
			out.Values[i] = ec.___Type_inputFields(ctx, field, obj)
		case "ofType":
			out.Values[i] = ec.___Type_ofType(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	if invalid {
		return graphql.Null
	}
	return out
}

// nolint: vetshadow
func (ec *executionContext) ___Type_kind(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	rctx := &graphql.ResolverContext{
		Object: "__Type",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(ctx context.Context) (interface{}, error) {
		return obj.Kind(), nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	return graphql.MarshalString(res)
}

// nolint: vetshadow
func (ec *executionContext) ___Type_name(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	rctx := &graphql.ResolverContext{
		Object: "__Type",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(ctx context.Context) (interface{}, error) {
		return obj.Name(), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	rctx.Result = res

	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

// nolint: vetshadow
func (ec *executionContext) ___Type_description(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	rctx := &graphql.ResolverContext{
		Object: "__Type",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(ctx context.Context) (interface{}, error) {
		return obj.Description(), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	return graphql.MarshalString(res)
}

// nolint: vetshadow
func (ec *executionContext) ___Type_fields(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	rawArgs := field.ArgumentMap(ec.Variables)
	args := map[string]interface{}{}
	var arg0 bool
	if tmp, ok := rawArgs["includeDeprecated"]; ok {
		var err error
		arg0, err = graphql.UnmarshalBoolean(tmp)
		if err != nil {
			ec.Error(ctx, err)
			return graphql.Null
		}
	}
	args["includeDeprecated"] = arg0
	rctx := &graphql.ResolverContext{
		Object: "__Type",
		Args:   args,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(ctx context.Context) (interface{}, error) {
		return obj.Fields(args["includeDeprecated"].(bool)), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]introspection.Field)
	rctx.Result = res
	arr1 := graphql.Array{}
	for idx1 := range res {
		rctx := &graphql.ResolverContext{
			Index:  &idx1,
			Result: &res[idx1],
		}
		ctx := graphql.WithResolverContext(ctx, rctx)
		arr1 = append(arr1, func() graphql.Marshaler {

			return ec.___Field(ctx, field.Selections, &res[idx1])
		}())
	}
	return arr1
}

// nolint: vetshadow
func (ec *executionContext) ___Type_interfaces(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	rctx := &graphql.ResolverContext{
		Object: "__Type",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(ctx context.Context) (interface{}, error) {
		return obj.Interfaces(), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]introspection.Type)
	rctx.Result = res
	arr1 := graphql.Array{}
	for idx1 := range res {
		rctx := &graphql.ResolverContext{
			Index:  &idx1,
			Result: &res[idx1],
		}
		ctx := graphql.WithResolverContext(ctx, rctx)
		arr1 = append(arr1, func() graphql.Marshaler {

			return ec.___Type(ctx, field.Selections, &res[idx1])
		}())
	}
	return arr1
}

// nolint: vetshadow
func (ec *executionContext) ___Type_possibleTypes(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	rctx := &graphql.ResolverContext{
		Object: "__Type",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(ctx context.Context) (interface{}, error) {
		return obj.PossibleTypes(), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]introspection.Type)
	rctx.Result = res
	arr1 := graphql.Array{}
	for idx1 := range res {
		rctx := &graphql.ResolverContext{
			Index:  &idx1,
			Result: &res[idx1],
		}
		ctx := graphql.WithResolverContext(ctx, rctx)
		arr1 = append(arr1, func() graphql.Marshaler {

			return ec.___Type(ctx, field.Selections, &res[idx1])
		}())
	}
	return arr1
}

// nolint: vetshadow
func (ec *executionContext) ___Type_enumValues(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	rawArgs := field.ArgumentMap(ec.Variables)
	args := map[string]interface{}{}
	var arg0 bool
	if tmp, ok := rawArgs["includeDeprecated"]; ok {
		var err error
		arg0, err = graphql.UnmarshalBoolean(tmp)
		if err != nil {
			ec.Error(ctx, err)
			return graphql.Null
		}
	}
	args["includeDeprecated"] = arg0
	rctx := &graphql.ResolverContext{
		Object: "__Type",
		Args:   args,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(ctx context.Context) (interface{}, error) {
		return obj.EnumValues(args["includeDeprecated"].(bool)), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]introspection.EnumValue)
	rctx.Result = res
	arr1 := graphql.Array{}
	for idx1 := range res {
		rctx := &graphql.ResolverContext{
			Index:  &idx1,
			Result: &res[idx1],
		}
		ctx := graphql.WithResolverContext(ctx, rctx)
		arr1 = append(arr1, func() graphql.Marshaler {

			return ec.___EnumValue(ctx, field.Selections, &res[idx1])
		}())
	}
	return arr1
}

// nolint: vetshadow
func (ec *executionContext) ___Type_inputFields(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	rctx := &graphql.ResolverContext{
		Object: "__Type",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(ctx context.Context) (interface{}, error) {
		return obj.InputFields(), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]introspection.InputValue)
	rctx.Result = res
	arr1 := graphql.Array{}
	for idx1 := range res {
		rctx := &graphql.ResolverContext{
			Index:  &idx1,
			Result: &res[idx1],
		}
		ctx := graphql.WithResolverContext(ctx, rctx)
		arr1 = append(arr1, func() graphql.Marshaler {

			return ec.___InputValue(ctx, field.Selections, &res[idx1])
		}())
	}
	return arr1
}

// nolint: vetshadow
func (ec *executionContext) ___Type_ofType(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	rctx := &graphql.ResolverContext{
		Object: "__Type",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(ctx context.Context) (interface{}, error) {
		return obj.OfType(), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	rctx.Result = res

	if res == nil {
		return graphql.Null
	}

	return ec.___Type(ctx, field.Selections, res)
}

func (ec *executionContext) FieldMiddleware(ctx context.Context, obj interface{}, next graphql.Resolver) (ret interface{}) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = nil
		}
	}()
	res, err := ec.ResolverMiddleware(ctx, next)
	if err != nil {
		ec.Error(ctx, err)
		return nil
	}
	return res
}

func (ec *executionContext) introspectSchema() *introspection.Schema {
	return introspection.WrapSchema(parsedSchema)
}

func (ec *executionContext) introspectType(name string) *introspection.Type {
	return introspection.WrapTypeFromDef(parsedSchema, parsedSchema.Types[name])
}

var parsedSchema = gqlparser.MustLoadSchema(
	&ast.Source{Name: "schema.graphql", Input: `type HelloResponse {
  msg: String!
}

type Query {
  hello(name: String! = "World"): HelloResponse
}
`},
)
