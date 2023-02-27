// Code generated by ent, DO NOT EDIT.

package maintenanceentry

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/thechosenlan/homebox/backend/internal/data/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldEQ(FieldUpdatedAt, v))
}

// ItemID applies equality check predicate on the "item_id" field. It's identical to ItemIDEQ.
func ItemID(v uuid.UUID) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldEQ(FieldItemID, v))
}

// Date applies equality check predicate on the "date" field. It's identical to DateEQ.
func Date(v time.Time) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldEQ(FieldDate, v))
}

// ScheduledDate applies equality check predicate on the "scheduled_date" field. It's identical to ScheduledDateEQ.
func ScheduledDate(v time.Time) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldEQ(FieldScheduledDate, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldEQ(FieldName, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldEQ(FieldDescription, v))
}

// Cost applies equality check predicate on the "cost" field. It's identical to CostEQ.
func Cost(v float64) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldEQ(FieldCost, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldLTE(FieldUpdatedAt, v))
}

// ItemIDEQ applies the EQ predicate on the "item_id" field.
func ItemIDEQ(v uuid.UUID) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldEQ(FieldItemID, v))
}

// ItemIDNEQ applies the NEQ predicate on the "item_id" field.
func ItemIDNEQ(v uuid.UUID) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldNEQ(FieldItemID, v))
}

// ItemIDIn applies the In predicate on the "item_id" field.
func ItemIDIn(vs ...uuid.UUID) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldIn(FieldItemID, vs...))
}

// ItemIDNotIn applies the NotIn predicate on the "item_id" field.
func ItemIDNotIn(vs ...uuid.UUID) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldNotIn(FieldItemID, vs...))
}

// DateEQ applies the EQ predicate on the "date" field.
func DateEQ(v time.Time) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldEQ(FieldDate, v))
}

// DateNEQ applies the NEQ predicate on the "date" field.
func DateNEQ(v time.Time) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldNEQ(FieldDate, v))
}

// DateIn applies the In predicate on the "date" field.
func DateIn(vs ...time.Time) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldIn(FieldDate, vs...))
}

// DateNotIn applies the NotIn predicate on the "date" field.
func DateNotIn(vs ...time.Time) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldNotIn(FieldDate, vs...))
}

// DateGT applies the GT predicate on the "date" field.
func DateGT(v time.Time) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldGT(FieldDate, v))
}

// DateGTE applies the GTE predicate on the "date" field.
func DateGTE(v time.Time) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldGTE(FieldDate, v))
}

// DateLT applies the LT predicate on the "date" field.
func DateLT(v time.Time) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldLT(FieldDate, v))
}

// DateLTE applies the LTE predicate on the "date" field.
func DateLTE(v time.Time) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldLTE(FieldDate, v))
}

// DateIsNil applies the IsNil predicate on the "date" field.
func DateIsNil() predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldIsNull(FieldDate))
}

// DateNotNil applies the NotNil predicate on the "date" field.
func DateNotNil() predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldNotNull(FieldDate))
}

// ScheduledDateEQ applies the EQ predicate on the "scheduled_date" field.
func ScheduledDateEQ(v time.Time) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldEQ(FieldScheduledDate, v))
}

// ScheduledDateNEQ applies the NEQ predicate on the "scheduled_date" field.
func ScheduledDateNEQ(v time.Time) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldNEQ(FieldScheduledDate, v))
}

// ScheduledDateIn applies the In predicate on the "scheduled_date" field.
func ScheduledDateIn(vs ...time.Time) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldIn(FieldScheduledDate, vs...))
}

// ScheduledDateNotIn applies the NotIn predicate on the "scheduled_date" field.
func ScheduledDateNotIn(vs ...time.Time) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldNotIn(FieldScheduledDate, vs...))
}

// ScheduledDateGT applies the GT predicate on the "scheduled_date" field.
func ScheduledDateGT(v time.Time) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldGT(FieldScheduledDate, v))
}

// ScheduledDateGTE applies the GTE predicate on the "scheduled_date" field.
func ScheduledDateGTE(v time.Time) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldGTE(FieldScheduledDate, v))
}

// ScheduledDateLT applies the LT predicate on the "scheduled_date" field.
func ScheduledDateLT(v time.Time) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldLT(FieldScheduledDate, v))
}

// ScheduledDateLTE applies the LTE predicate on the "scheduled_date" field.
func ScheduledDateLTE(v time.Time) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldLTE(FieldScheduledDate, v))
}

// ScheduledDateIsNil applies the IsNil predicate on the "scheduled_date" field.
func ScheduledDateIsNil() predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldIsNull(FieldScheduledDate))
}

// ScheduledDateNotNil applies the NotNil predicate on the "scheduled_date" field.
func ScheduledDateNotNil() predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldNotNull(FieldScheduledDate))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldContainsFold(FieldName, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldContainsFold(FieldDescription, v))
}

// CostEQ applies the EQ predicate on the "cost" field.
func CostEQ(v float64) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldEQ(FieldCost, v))
}

// CostNEQ applies the NEQ predicate on the "cost" field.
func CostNEQ(v float64) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldNEQ(FieldCost, v))
}

// CostIn applies the In predicate on the "cost" field.
func CostIn(vs ...float64) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldIn(FieldCost, vs...))
}

// CostNotIn applies the NotIn predicate on the "cost" field.
func CostNotIn(vs ...float64) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldNotIn(FieldCost, vs...))
}

// CostGT applies the GT predicate on the "cost" field.
func CostGT(v float64) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldGT(FieldCost, v))
}

// CostGTE applies the GTE predicate on the "cost" field.
func CostGTE(v float64) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldGTE(FieldCost, v))
}

// CostLT applies the LT predicate on the "cost" field.
func CostLT(v float64) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldLT(FieldCost, v))
}

// CostLTE applies the LTE predicate on the "cost" field.
func CostLTE(v float64) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(sql.FieldLTE(FieldCost, v))
}

// HasItem applies the HasEdge predicate on the "item" edge.
func HasItem() predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ItemTable, ItemColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasItemWith applies the HasEdge predicate on the "item" edge with a given conditions (other predicates).
func HasItemWith(preds ...predicate.Item) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ItemInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ItemTable, ItemColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.MaintenanceEntry) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.MaintenanceEntry) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(func(s *sql.Selector) {
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
func Not(p predicate.MaintenanceEntry) predicate.MaintenanceEntry {
	return predicate.MaintenanceEntry(func(s *sql.Selector) {
		p(s.Not())
	})
}
