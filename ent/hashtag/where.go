// Code generated by ent, DO NOT EDIT.

package hashtag

import (
	"enttest/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Hashtag {
	return predicate.Hashtag(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Hashtag {
	return predicate.Hashtag(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Hashtag {
	return predicate.Hashtag(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Hashtag {
	return predicate.Hashtag(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Hashtag {
	return predicate.Hashtag(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Hashtag {
	return predicate.Hashtag(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Hashtag {
	return predicate.Hashtag(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Hashtag {
	return predicate.Hashtag(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Hashtag {
	return predicate.Hashtag(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// HasPost applies the HasEdge predicate on the "post" edge.
func HasPost() predicate.Hashtag {
	return predicate.Hashtag(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PostTable, PostFieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, PostTable, PostPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPostWith applies the HasEdge predicate on the "post" edge with a given conditions (other predicates).
func HasPostWith(preds ...predicate.Post) predicate.Hashtag {
	return predicate.Hashtag(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PostInverseTable, PostFieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, PostTable, PostPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasHashtagToPost applies the HasEdge predicate on the "hashtag_to_post" edge.
func HasHashtagToPost() predicate.Hashtag {
	return predicate.Hashtag(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(HashtagToPostTable, HashtagToPostFieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, HashtagToPostTable, HashtagToPostColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasHashtagToPostWith applies the HasEdge predicate on the "hashtag_to_post" edge with a given conditions (other predicates).
func HasHashtagToPostWith(preds ...predicate.HashtagToPost) predicate.Hashtag {
	return predicate.Hashtag(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(HashtagToPostInverseTable, HashtagToPostFieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, HashtagToPostTable, HashtagToPostColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Hashtag) predicate.Hashtag {
	return predicate.Hashtag(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Hashtag) predicate.Hashtag {
	return predicate.Hashtag(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Hashtag) predicate.Hashtag {
	return predicate.Hashtag(func(s *sql.Selector) {
		p(s.Not())
	})
}
