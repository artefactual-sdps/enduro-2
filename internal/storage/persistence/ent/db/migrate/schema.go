// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// LocationColumns holds the columns for the "location" table.
	LocationColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Size: 2048},
		{Name: "description", Type: field.TypeString, Size: 2048},
		{Name: "source", Type: field.TypeEnum, Enums: []string{"unspecified", "minio"}},
		{Name: "purpose", Type: field.TypeEnum, Enums: []string{"unspecified", "aip_store"}},
		{Name: "uuid", Type: field.TypeUUID, Unique: true},
		{Name: "config", Type: field.TypeJSON},
		{Name: "created_at", Type: field.TypeTime},
	}
	// LocationTable holds the schema information for the "location" table.
	LocationTable = &schema.Table{
		Name:       "location",
		Columns:    LocationColumns,
		PrimaryKey: []*schema.Column{LocationColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "location_name",
				Unique:  false,
				Columns: []*schema.Column{LocationColumns[1]},
				Annotation: &entsql.IndexAnnotation{
					Prefix: 50,
				},
			},
			{
				Name:    "location_uuid",
				Unique:  false,
				Columns: []*schema.Column{LocationColumns[5]},
			},
		},
	}
	// PackageColumns holds the columns for the "package" table.
	PackageColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Size: 2048},
		{Name: "aip_id", Type: field.TypeUUID, Unique: true},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"unspecified", "in_review", "rejected", "stored", "moving"}},
		{Name: "object_key", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "location_id", Type: field.TypeInt, Nullable: true},
	}
	// PackageTable holds the schema information for the "package" table.
	PackageTable = &schema.Table{
		Name:       "package",
		Columns:    PackageColumns,
		PrimaryKey: []*schema.Column{PackageColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "package_location_location",
				Columns:    []*schema.Column{PackageColumns[6]},
				RefColumns: []*schema.Column{LocationColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "pkg_aip_id",
				Unique:  false,
				Columns: []*schema.Column{PackageColumns[2]},
			},
			{
				Name:    "pkg_object_key",
				Unique:  false,
				Columns: []*schema.Column{PackageColumns[4]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		LocationTable,
		PackageTable,
	}
)

func init() {
	LocationTable.Annotation = &entsql.Annotation{
		Table: "location",
	}
	PackageTable.ForeignKeys[0].RefTable = LocationTable
	PackageTable.Annotation = &entsql.Annotation{
		Table: "package",
	}
}
