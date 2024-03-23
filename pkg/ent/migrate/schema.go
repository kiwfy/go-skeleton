// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CartsColumns holds the columns for the "carts" table.
	CartsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "modified_at", Type: field.TypeTime},
		{Name: "user_cart", Type: field.TypeString, Nullable: true},
	}
	// CartsTable holds the schema information for the "carts" table.
	CartsTable = &schema.Table{
		Name:       "carts",
		Columns:    CartsColumns,
		PrimaryKey: []*schema.Column{CartsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "carts_users_cart",
				Columns:    []*schema.Column{CartsColumns[2]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CartsTable,
		UsersTable,
	}
)

func init() {
	CartsTable.ForeignKeys[0].RefTable = UsersTable
}
