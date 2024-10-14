package repository

import "github.com/techforge-lat/sqlcraft"

const tableName = "projects"

var selectColumns = []string{"id", "display_name", "description", "created_at", "updated_at"}

var (
	createQuery = sqlcraft.InsertInto(tableName).WithColumns("id", "display_name", "description", "created_at")
	updateQuery = sqlcraft.Update(tableName).WithColumns("display_name", "description", "updated_at")
	deleteQuery = sqlcraft.DeleteFrom(tableName)
	getQuery    = sqlcraft.Select(selectColumns...).From(tableName)
	listQuery   = sqlcraft.Select(selectColumns...).From(tableName)
)
