package utils

import (
	"fmt"
	"strings"
)

type QueryBuilder struct {
	query string
}

func (b *QueryBuilder) Build() string {
	return b.query
}

func (b *QueryBuilder) Select(args ...string) *QueryBuilder {
	b.query = "SELECT " + strings.Join(args, ", ")
	return b
}

func (b *QueryBuilder) From(table string) *QueryBuilder {
	b.query += " FROM " + table
	return b
}

func (b *QueryBuilder) Where(filters map[string]string) *QueryBuilder {
	var pairs []string
	for k, v := range filters {
		pairs = append(pairs, fmt.Sprintf("%s=%s", k, v))
	}
	b.query += " WHERE " + strings.Join(pairs, " AND ")
	return b
}

func (b *QueryBuilder) OrderBy(orders map[string]bool) *QueryBuilder {
	var pairs []string
	for k, v := range orders {
		if v {
			pairs = append(pairs, fmt.Sprintf("%s ASC", k))
		} else {
			pairs = append(pairs, fmt.Sprintf("%s DESC", k))
		}
	}
	b.query += " ORDER BY " + strings.Join(pairs, ", ")
	return b
}

func (b *QueryBuilder) InsertInto(table string) *QueryBuilder {
	b.query = "INSERT INTO " + table
	return b
}

func (b *QueryBuilder) Fields(fields ...string) *QueryBuilder {
	b.query += " (" + strings.Join(fields, ", ") + ")"
	return b
}

func (b *QueryBuilder) Values(values int) *QueryBuilder {
	var dValues []string
	for i := range values {
		dValues = append(dValues, fmt.Sprintf("$%d", i+1))
	}
	b.query += " VALUES (" + strings.Join(dValues, ", ") + ")"
	return b
}
