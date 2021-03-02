// Code generated by entc, DO NOT EDIT.

package userrole

import (
	"repro/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.UserRole {
	return predicate.UserRole(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.UserRole {
	return predicate.UserRole(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.UserRole {
	return predicate.UserRole(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.UserRole {
	return predicate.UserRole(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.UserRole {
	return predicate.UserRole(func(s *sql.Selector) {
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
func IDGT(id int) predicate.UserRole {
	return predicate.UserRole(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.UserRole {
	return predicate.UserRole(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.UserRole {
	return predicate.UserRole(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.UserRole {
	return predicate.UserRole(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// TargetID applies equality check predicate on the "target_id" field. It's identical to TargetIDEQ.
func TargetID(v int) predicate.UserRole {
	return predicate.UserRole(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTargetID), v))
	})
}

// TargetType applies equality check predicate on the "target_type" field. It's identical to TargetTypeEQ.
func TargetType(v string) predicate.UserRole {
	return predicate.UserRole(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTargetType), v))
	})
}

// ActorType applies equality check predicate on the "actor_type" field. It's identical to ActorTypeEQ.
func ActorType(v string) predicate.UserRole {
	return predicate.UserRole(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldActorType), v))
	})
}

// ActorID applies equality check predicate on the "actor_id" field. It's identical to ActorIDEQ.
func ActorID(v int) predicate.UserRole {
	return predicate.UserRole(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldActorID), v))
	})
}

// TargetIDEQ applies the EQ predicate on the "target_id" field.
func TargetIDEQ(v int) predicate.UserRole {
	return predicate.UserRole(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTargetID), v))
	})
}

// TargetIDNEQ applies the NEQ predicate on the "target_id" field.
func TargetIDNEQ(v int) predicate.UserRole {
	return predicate.UserRole(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTargetID), v))
	})
}

// TargetIDIn applies the In predicate on the "target_id" field.
func TargetIDIn(vs ...int) predicate.UserRole {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserRole(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTargetID), v...))
	})
}

// TargetIDNotIn applies the NotIn predicate on the "target_id" field.
func TargetIDNotIn(vs ...int) predicate.UserRole {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserRole(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTargetID), v...))
	})
}

// TargetIDGT applies the GT predicate on the "target_id" field.
func TargetIDGT(v int) predicate.UserRole {
	return predicate.UserRole(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTargetID), v))
	})
}

// TargetIDGTE applies the GTE predicate on the "target_id" field.
func TargetIDGTE(v int) predicate.UserRole {
	return predicate.UserRole(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTargetID), v))
	})
}

// TargetIDLT applies the LT predicate on the "target_id" field.
func TargetIDLT(v int) predicate.UserRole {
	return predicate.UserRole(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTargetID), v))
	})
}

// TargetIDLTE applies the LTE predicate on the "target_id" field.
func TargetIDLTE(v int) predicate.UserRole {
	return predicate.UserRole(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTargetID), v))
	})
}

// TargetTypeEQ applies the EQ predicate on the "target_type" field.
func TargetTypeEQ(v string) predicate.UserRole {
	return predicate.UserRole(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTargetType), v))
	})
}

// TargetTypeNEQ applies the NEQ predicate on the "target_type" field.
func TargetTypeNEQ(v string) predicate.UserRole {
	return predicate.UserRole(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTargetType), v))
	})
}

// TargetTypeIn applies the In predicate on the "target_type" field.
func TargetTypeIn(vs ...string) predicate.UserRole {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserRole(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTargetType), v...))
	})
}

// TargetTypeNotIn applies the NotIn predicate on the "target_type" field.
func TargetTypeNotIn(vs ...string) predicate.UserRole {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserRole(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTargetType), v...))
	})
}

// TargetTypeGT applies the GT predicate on the "target_type" field.
func TargetTypeGT(v string) predicate.UserRole {
	return predicate.UserRole(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTargetType), v))
	})
}

// TargetTypeGTE applies the GTE predicate on the "target_type" field.
func TargetTypeGTE(v string) predicate.UserRole {
	return predicate.UserRole(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTargetType), v))
	})
}

// TargetTypeLT applies the LT predicate on the "target_type" field.
func TargetTypeLT(v string) predicate.UserRole {
	return predicate.UserRole(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTargetType), v))
	})
}

// TargetTypeLTE applies the LTE predicate on the "target_type" field.
func TargetTypeLTE(v string) predicate.UserRole {
	return predicate.UserRole(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTargetType), v))
	})
}

// TargetTypeContains applies the Contains predicate on the "target_type" field.
func TargetTypeContains(v string) predicate.UserRole {
	return predicate.UserRole(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTargetType), v))
	})
}

// TargetTypeHasPrefix applies the HasPrefix predicate on the "target_type" field.
func TargetTypeHasPrefix(v string) predicate.UserRole {
	return predicate.UserRole(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTargetType), v))
	})
}

// TargetTypeHasSuffix applies the HasSuffix predicate on the "target_type" field.
func TargetTypeHasSuffix(v string) predicate.UserRole {
	return predicate.UserRole(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTargetType), v))
	})
}

// TargetTypeEqualFold applies the EqualFold predicate on the "target_type" field.
func TargetTypeEqualFold(v string) predicate.UserRole {
	return predicate.UserRole(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTargetType), v))
	})
}

// TargetTypeContainsFold applies the ContainsFold predicate on the "target_type" field.
func TargetTypeContainsFold(v string) predicate.UserRole {
	return predicate.UserRole(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTargetType), v))
	})
}

// ActorTypeEQ applies the EQ predicate on the "actor_type" field.
func ActorTypeEQ(v string) predicate.UserRole {
	return predicate.UserRole(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldActorType), v))
	})
}

// ActorTypeNEQ applies the NEQ predicate on the "actor_type" field.
func ActorTypeNEQ(v string) predicate.UserRole {
	return predicate.UserRole(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldActorType), v))
	})
}

// ActorTypeIn applies the In predicate on the "actor_type" field.
func ActorTypeIn(vs ...string) predicate.UserRole {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserRole(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldActorType), v...))
	})
}

// ActorTypeNotIn applies the NotIn predicate on the "actor_type" field.
func ActorTypeNotIn(vs ...string) predicate.UserRole {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserRole(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldActorType), v...))
	})
}

// ActorTypeGT applies the GT predicate on the "actor_type" field.
func ActorTypeGT(v string) predicate.UserRole {
	return predicate.UserRole(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldActorType), v))
	})
}

// ActorTypeGTE applies the GTE predicate on the "actor_type" field.
func ActorTypeGTE(v string) predicate.UserRole {
	return predicate.UserRole(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldActorType), v))
	})
}

// ActorTypeLT applies the LT predicate on the "actor_type" field.
func ActorTypeLT(v string) predicate.UserRole {
	return predicate.UserRole(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldActorType), v))
	})
}

// ActorTypeLTE applies the LTE predicate on the "actor_type" field.
func ActorTypeLTE(v string) predicate.UserRole {
	return predicate.UserRole(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldActorType), v))
	})
}

// ActorTypeContains applies the Contains predicate on the "actor_type" field.
func ActorTypeContains(v string) predicate.UserRole {
	return predicate.UserRole(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldActorType), v))
	})
}

// ActorTypeHasPrefix applies the HasPrefix predicate on the "actor_type" field.
func ActorTypeHasPrefix(v string) predicate.UserRole {
	return predicate.UserRole(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldActorType), v))
	})
}

// ActorTypeHasSuffix applies the HasSuffix predicate on the "actor_type" field.
func ActorTypeHasSuffix(v string) predicate.UserRole {
	return predicate.UserRole(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldActorType), v))
	})
}

// ActorTypeEqualFold applies the EqualFold predicate on the "actor_type" field.
func ActorTypeEqualFold(v string) predicate.UserRole {
	return predicate.UserRole(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldActorType), v))
	})
}

// ActorTypeContainsFold applies the ContainsFold predicate on the "actor_type" field.
func ActorTypeContainsFold(v string) predicate.UserRole {
	return predicate.UserRole(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldActorType), v))
	})
}

// ActorIDEQ applies the EQ predicate on the "actor_id" field.
func ActorIDEQ(v int) predicate.UserRole {
	return predicate.UserRole(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldActorID), v))
	})
}

// ActorIDNEQ applies the NEQ predicate on the "actor_id" field.
func ActorIDNEQ(v int) predicate.UserRole {
	return predicate.UserRole(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldActorID), v))
	})
}

// ActorIDIn applies the In predicate on the "actor_id" field.
func ActorIDIn(vs ...int) predicate.UserRole {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserRole(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldActorID), v...))
	})
}

// ActorIDNotIn applies the NotIn predicate on the "actor_id" field.
func ActorIDNotIn(vs ...int) predicate.UserRole {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserRole(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldActorID), v...))
	})
}

// ActorIDGT applies the GT predicate on the "actor_id" field.
func ActorIDGT(v int) predicate.UserRole {
	return predicate.UserRole(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldActorID), v))
	})
}

// ActorIDGTE applies the GTE predicate on the "actor_id" field.
func ActorIDGTE(v int) predicate.UserRole {
	return predicate.UserRole(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldActorID), v))
	})
}

// ActorIDLT applies the LT predicate on the "actor_id" field.
func ActorIDLT(v int) predicate.UserRole {
	return predicate.UserRole(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldActorID), v))
	})
}

// ActorIDLTE applies the LTE predicate on the "actor_id" field.
func ActorIDLTE(v int) predicate.UserRole {
	return predicate.UserRole(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldActorID), v))
	})
}

// HasRole applies the HasEdge predicate on the "role" edge.
func HasRole() predicate.UserRole {
	return predicate.UserRole(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RoleTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, RoleTable, RoleColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRoleWith applies the HasEdge predicate on the "role" edge with a given conditions (other predicates).
func HasRoleWith(preds ...predicate.Role) predicate.UserRole {
	return predicate.UserRole(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RoleInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, RoleTable, RoleColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.UserRole {
	return predicate.UserRole(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.UserRole {
	return predicate.UserRole(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.UserRole) predicate.UserRole {
	return predicate.UserRole(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.UserRole) predicate.UserRole {
	return predicate.UserRole(func(s *sql.Selector) {
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
func Not(p predicate.UserRole) predicate.UserRole {
	return predicate.UserRole(func(s *sql.Selector) {
		p(s.Not())
	})
}