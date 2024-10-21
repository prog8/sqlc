// Code generated by sqlc-pg-gen. DO NOT EDIT.

package contrib

import (
	"github.com/prog8/sqlc/internal/sql/ast"
	"github.com/prog8/sqlc/internal/sql/catalog"
)

var funcsHstore = []*catalog.Function{
	{
		Name: "akeys",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "hstore"},
			},
		},
		ReturnType: &ast.TypeName{Name: "text[]"},
	},
	{
		Name: "avals",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "hstore"},
			},
		},
		ReturnType: &ast.TypeName{Name: "text[]"},
	},
	{
		Name: "defined",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "hstore"},
			},
			{
				Type: &ast.TypeName{Name: "text"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "delete",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "hstore"},
			},
			{
				Type: &ast.TypeName{Name: "hstore"},
			},
		},
		ReturnType: &ast.TypeName{Name: "hstore"},
	},
	{
		Name: "delete",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "hstore"},
			},
			{
				Type: &ast.TypeName{Name: "text"},
			},
		},
		ReturnType: &ast.TypeName{Name: "hstore"},
	},
	{
		Name: "delete",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "hstore"},
			},
			{
				Type: &ast.TypeName{Name: "text[]"},
			},
		},
		ReturnType: &ast.TypeName{Name: "hstore"},
	},
	{
		Name: "each",
		Args: []*catalog.Argument{
			{
				Name: "hs",
				Type: &ast.TypeName{Name: "hstore"},
			},
		},
		ReturnType: &ast.TypeName{Name: "record"},
	},
	{
		Name: "exist",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "hstore"},
			},
			{
				Type: &ast.TypeName{Name: "text"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "exists_all",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "hstore"},
			},
			{
				Type: &ast.TypeName{Name: "text[]"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "exists_any",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "hstore"},
			},
			{
				Type: &ast.TypeName{Name: "text[]"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "fetchval",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "hstore"},
			},
			{
				Type: &ast.TypeName{Name: "text"},
			},
		},
		ReturnType: &ast.TypeName{Name: "text"},
	},
	{
		Name: "ghstore_in",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "cstring"},
			},
		},
		ReturnType: &ast.TypeName{Name: "ghstore"},
	},
	{
		Name: "ghstore_out",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ghstore"},
			},
		},
		ReturnType: &ast.TypeName{Name: "cstring"},
	},
	{
		Name: "hs_concat",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "hstore"},
			},
			{
				Type: &ast.TypeName{Name: "hstore"},
			},
		},
		ReturnType: &ast.TypeName{Name: "hstore"},
	},
	{
		Name: "hs_contained",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "hstore"},
			},
			{
				Type: &ast.TypeName{Name: "hstore"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "hs_contains",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "hstore"},
			},
			{
				Type: &ast.TypeName{Name: "hstore"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "hstore",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "record"},
			},
		},
		ReturnType: &ast.TypeName{Name: "hstore"},
	},
	{
		Name: "hstore",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "text"},
			},
			{
				Type: &ast.TypeName{Name: "text"},
			},
		},
		ReturnType: &ast.TypeName{Name: "hstore"},
	},
	{
		Name: "hstore",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "text[]"},
			},
		},
		ReturnType: &ast.TypeName{Name: "hstore"},
	},
	{
		Name: "hstore",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "text[]"},
			},
			{
				Type: &ast.TypeName{Name: "text[]"},
			},
		},
		ReturnType: &ast.TypeName{Name: "hstore"},
	},
	{
		Name: "hstore_cmp",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "hstore"},
			},
			{
				Type: &ast.TypeName{Name: "hstore"},
			},
		},
		ReturnType: &ast.TypeName{Name: "integer"},
	},
	{
		Name: "hstore_eq",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "hstore"},
			},
			{
				Type: &ast.TypeName{Name: "hstore"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "hstore_ge",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "hstore"},
			},
			{
				Type: &ast.TypeName{Name: "hstore"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "hstore_gt",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "hstore"},
			},
			{
				Type: &ast.TypeName{Name: "hstore"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "hstore_hash",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "hstore"},
			},
		},
		ReturnType: &ast.TypeName{Name: "integer"},
	},
	{
		Name: "hstore_hash_extended",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "hstore"},
			},
			{
				Type: &ast.TypeName{Name: "bigint"},
			},
		},
		ReturnType: &ast.TypeName{Name: "bigint"},
	},
	{
		Name: "hstore_in",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "cstring"},
			},
		},
		ReturnType: &ast.TypeName{Name: "hstore"},
	},
	{
		Name: "hstore_le",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "hstore"},
			},
			{
				Type: &ast.TypeName{Name: "hstore"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "hstore_lt",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "hstore"},
			},
			{
				Type: &ast.TypeName{Name: "hstore"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "hstore_ne",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "hstore"},
			},
			{
				Type: &ast.TypeName{Name: "hstore"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "hstore_out",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "hstore"},
			},
		},
		ReturnType: &ast.TypeName{Name: "cstring"},
	},
	{
		Name: "hstore_send",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "hstore"},
			},
		},
		ReturnType: &ast.TypeName{Name: "bytea"},
	},
	{
		Name: "hstore_to_array",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "hstore"},
			},
		},
		ReturnType: &ast.TypeName{Name: "text[]"},
	},
	{
		Name: "hstore_to_json",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "hstore"},
			},
		},
		ReturnType: &ast.TypeName{Name: "json"},
	},
	{
		Name: "hstore_to_json_loose",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "hstore"},
			},
		},
		ReturnType: &ast.TypeName{Name: "json"},
	},
	{
		Name: "hstore_to_jsonb",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "hstore"},
			},
		},
		ReturnType: &ast.TypeName{Name: "jsonb"},
	},
	{
		Name: "hstore_to_jsonb_loose",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "hstore"},
			},
		},
		ReturnType: &ast.TypeName{Name: "jsonb"},
	},
	{
		Name: "hstore_to_matrix",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "hstore"},
			},
		},
		ReturnType: &ast.TypeName{Name: "text[]"},
	},
	{
		Name: "hstore_version_diag",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "hstore"},
			},
		},
		ReturnType: &ast.TypeName{Name: "integer"},
	},
	{
		Name: "isdefined",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "hstore"},
			},
			{
				Type: &ast.TypeName{Name: "text"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isexists",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "hstore"},
			},
			{
				Type: &ast.TypeName{Name: "text"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "populate_record",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "anyelement"},
			},
			{
				Type: &ast.TypeName{Name: "hstore"},
			},
		},
		ReturnType: &ast.TypeName{Name: "anyelement"},
	},
	{
		Name: "skeys",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "hstore"},
			},
		},
		ReturnType: &ast.TypeName{Name: "text"},
	},
	{
		Name: "slice",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "hstore"},
			},
			{
				Type: &ast.TypeName{Name: "text[]"},
			},
		},
		ReturnType: &ast.TypeName{Name: "hstore"},
	},
	{
		Name: "slice_array",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "hstore"},
			},
			{
				Type: &ast.TypeName{Name: "text[]"},
			},
		},
		ReturnType: &ast.TypeName{Name: "text[]"},
	},
	{
		Name: "svals",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "hstore"},
			},
		},
		ReturnType: &ast.TypeName{Name: "text"},
	},
	{
		Name: "tconvert",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "text"},
			},
			{
				Type: &ast.TypeName{Name: "text"},
			},
		},
		ReturnType: &ast.TypeName{Name: "hstore"},
	},
}

func Hstore() *catalog.Schema {
	s := &catalog.Schema{Name: "pg_catalog"}
	s.Funcs = funcsHstore
	return s
}
