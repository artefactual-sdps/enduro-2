// Code generated by ent, DO NOT EDIT.

package pkg

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/artefactual-sdps/enduro/internal/persistence/ent/db/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Pkg {
	return predicate.Pkg(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Pkg {
	return predicate.Pkg(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Pkg {
	return predicate.Pkg(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Pkg {
	return predicate.Pkg(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Pkg {
	return predicate.Pkg(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Pkg {
	return predicate.Pkg(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Pkg {
	return predicate.Pkg(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Pkg {
	return predicate.Pkg(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Pkg {
	return predicate.Pkg(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Pkg {
	return predicate.Pkg(sql.FieldEQ(FieldName, v))
}

// WorkflowID applies equality check predicate on the "workflow_id" field. It's identical to WorkflowIDEQ.
func WorkflowID(v string) predicate.Pkg {
	return predicate.Pkg(sql.FieldEQ(FieldWorkflowID, v))
}

// RunID applies equality check predicate on the "run_id" field. It's identical to RunIDEQ.
func RunID(v uuid.UUID) predicate.Pkg {
	return predicate.Pkg(sql.FieldEQ(FieldRunID, v))
}

// AipID applies equality check predicate on the "aip_id" field. It's identical to AipIDEQ.
func AipID(v uuid.UUID) predicate.Pkg {
	return predicate.Pkg(sql.FieldEQ(FieldAipID, v))
}

// LocationID applies equality check predicate on the "location_id" field. It's identical to LocationIDEQ.
func LocationID(v uuid.UUID) predicate.Pkg {
	return predicate.Pkg(sql.FieldEQ(FieldLocationID, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v int8) predicate.Pkg {
	return predicate.Pkg(sql.FieldEQ(FieldStatus, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Pkg {
	return predicate.Pkg(sql.FieldEQ(FieldCreatedAt, v))
}

// StartedAt applies equality check predicate on the "started_at" field. It's identical to StartedAtEQ.
func StartedAt(v time.Time) predicate.Pkg {
	return predicate.Pkg(sql.FieldEQ(FieldStartedAt, v))
}

// CompletedAt applies equality check predicate on the "completed_at" field. It's identical to CompletedAtEQ.
func CompletedAt(v time.Time) predicate.Pkg {
	return predicate.Pkg(sql.FieldEQ(FieldCompletedAt, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Pkg {
	return predicate.Pkg(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Pkg {
	return predicate.Pkg(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Pkg {
	return predicate.Pkg(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Pkg {
	return predicate.Pkg(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Pkg {
	return predicate.Pkg(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Pkg {
	return predicate.Pkg(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Pkg {
	return predicate.Pkg(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Pkg {
	return predicate.Pkg(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Pkg {
	return predicate.Pkg(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Pkg {
	return predicate.Pkg(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Pkg {
	return predicate.Pkg(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Pkg {
	return predicate.Pkg(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Pkg {
	return predicate.Pkg(sql.FieldContainsFold(FieldName, v))
}

// WorkflowIDEQ applies the EQ predicate on the "workflow_id" field.
func WorkflowIDEQ(v string) predicate.Pkg {
	return predicate.Pkg(sql.FieldEQ(FieldWorkflowID, v))
}

// WorkflowIDNEQ applies the NEQ predicate on the "workflow_id" field.
func WorkflowIDNEQ(v string) predicate.Pkg {
	return predicate.Pkg(sql.FieldNEQ(FieldWorkflowID, v))
}

// WorkflowIDIn applies the In predicate on the "workflow_id" field.
func WorkflowIDIn(vs ...string) predicate.Pkg {
	return predicate.Pkg(sql.FieldIn(FieldWorkflowID, vs...))
}

// WorkflowIDNotIn applies the NotIn predicate on the "workflow_id" field.
func WorkflowIDNotIn(vs ...string) predicate.Pkg {
	return predicate.Pkg(sql.FieldNotIn(FieldWorkflowID, vs...))
}

// WorkflowIDGT applies the GT predicate on the "workflow_id" field.
func WorkflowIDGT(v string) predicate.Pkg {
	return predicate.Pkg(sql.FieldGT(FieldWorkflowID, v))
}

// WorkflowIDGTE applies the GTE predicate on the "workflow_id" field.
func WorkflowIDGTE(v string) predicate.Pkg {
	return predicate.Pkg(sql.FieldGTE(FieldWorkflowID, v))
}

// WorkflowIDLT applies the LT predicate on the "workflow_id" field.
func WorkflowIDLT(v string) predicate.Pkg {
	return predicate.Pkg(sql.FieldLT(FieldWorkflowID, v))
}

// WorkflowIDLTE applies the LTE predicate on the "workflow_id" field.
func WorkflowIDLTE(v string) predicate.Pkg {
	return predicate.Pkg(sql.FieldLTE(FieldWorkflowID, v))
}

// WorkflowIDContains applies the Contains predicate on the "workflow_id" field.
func WorkflowIDContains(v string) predicate.Pkg {
	return predicate.Pkg(sql.FieldContains(FieldWorkflowID, v))
}

// WorkflowIDHasPrefix applies the HasPrefix predicate on the "workflow_id" field.
func WorkflowIDHasPrefix(v string) predicate.Pkg {
	return predicate.Pkg(sql.FieldHasPrefix(FieldWorkflowID, v))
}

// WorkflowIDHasSuffix applies the HasSuffix predicate on the "workflow_id" field.
func WorkflowIDHasSuffix(v string) predicate.Pkg {
	return predicate.Pkg(sql.FieldHasSuffix(FieldWorkflowID, v))
}

// WorkflowIDEqualFold applies the EqualFold predicate on the "workflow_id" field.
func WorkflowIDEqualFold(v string) predicate.Pkg {
	return predicate.Pkg(sql.FieldEqualFold(FieldWorkflowID, v))
}

// WorkflowIDContainsFold applies the ContainsFold predicate on the "workflow_id" field.
func WorkflowIDContainsFold(v string) predicate.Pkg {
	return predicate.Pkg(sql.FieldContainsFold(FieldWorkflowID, v))
}

// RunIDEQ applies the EQ predicate on the "run_id" field.
func RunIDEQ(v uuid.UUID) predicate.Pkg {
	return predicate.Pkg(sql.FieldEQ(FieldRunID, v))
}

// RunIDNEQ applies the NEQ predicate on the "run_id" field.
func RunIDNEQ(v uuid.UUID) predicate.Pkg {
	return predicate.Pkg(sql.FieldNEQ(FieldRunID, v))
}

// RunIDIn applies the In predicate on the "run_id" field.
func RunIDIn(vs ...uuid.UUID) predicate.Pkg {
	return predicate.Pkg(sql.FieldIn(FieldRunID, vs...))
}

// RunIDNotIn applies the NotIn predicate on the "run_id" field.
func RunIDNotIn(vs ...uuid.UUID) predicate.Pkg {
	return predicate.Pkg(sql.FieldNotIn(FieldRunID, vs...))
}

// RunIDGT applies the GT predicate on the "run_id" field.
func RunIDGT(v uuid.UUID) predicate.Pkg {
	return predicate.Pkg(sql.FieldGT(FieldRunID, v))
}

// RunIDGTE applies the GTE predicate on the "run_id" field.
func RunIDGTE(v uuid.UUID) predicate.Pkg {
	return predicate.Pkg(sql.FieldGTE(FieldRunID, v))
}

// RunIDLT applies the LT predicate on the "run_id" field.
func RunIDLT(v uuid.UUID) predicate.Pkg {
	return predicate.Pkg(sql.FieldLT(FieldRunID, v))
}

// RunIDLTE applies the LTE predicate on the "run_id" field.
func RunIDLTE(v uuid.UUID) predicate.Pkg {
	return predicate.Pkg(sql.FieldLTE(FieldRunID, v))
}

// AipIDEQ applies the EQ predicate on the "aip_id" field.
func AipIDEQ(v uuid.UUID) predicate.Pkg {
	return predicate.Pkg(sql.FieldEQ(FieldAipID, v))
}

// AipIDNEQ applies the NEQ predicate on the "aip_id" field.
func AipIDNEQ(v uuid.UUID) predicate.Pkg {
	return predicate.Pkg(sql.FieldNEQ(FieldAipID, v))
}

// AipIDIn applies the In predicate on the "aip_id" field.
func AipIDIn(vs ...uuid.UUID) predicate.Pkg {
	return predicate.Pkg(sql.FieldIn(FieldAipID, vs...))
}

// AipIDNotIn applies the NotIn predicate on the "aip_id" field.
func AipIDNotIn(vs ...uuid.UUID) predicate.Pkg {
	return predicate.Pkg(sql.FieldNotIn(FieldAipID, vs...))
}

// AipIDGT applies the GT predicate on the "aip_id" field.
func AipIDGT(v uuid.UUID) predicate.Pkg {
	return predicate.Pkg(sql.FieldGT(FieldAipID, v))
}

// AipIDGTE applies the GTE predicate on the "aip_id" field.
func AipIDGTE(v uuid.UUID) predicate.Pkg {
	return predicate.Pkg(sql.FieldGTE(FieldAipID, v))
}

// AipIDLT applies the LT predicate on the "aip_id" field.
func AipIDLT(v uuid.UUID) predicate.Pkg {
	return predicate.Pkg(sql.FieldLT(FieldAipID, v))
}

// AipIDLTE applies the LTE predicate on the "aip_id" field.
func AipIDLTE(v uuid.UUID) predicate.Pkg {
	return predicate.Pkg(sql.FieldLTE(FieldAipID, v))
}

// LocationIDEQ applies the EQ predicate on the "location_id" field.
func LocationIDEQ(v uuid.UUID) predicate.Pkg {
	return predicate.Pkg(sql.FieldEQ(FieldLocationID, v))
}

// LocationIDNEQ applies the NEQ predicate on the "location_id" field.
func LocationIDNEQ(v uuid.UUID) predicate.Pkg {
	return predicate.Pkg(sql.FieldNEQ(FieldLocationID, v))
}

// LocationIDIn applies the In predicate on the "location_id" field.
func LocationIDIn(vs ...uuid.UUID) predicate.Pkg {
	return predicate.Pkg(sql.FieldIn(FieldLocationID, vs...))
}

// LocationIDNotIn applies the NotIn predicate on the "location_id" field.
func LocationIDNotIn(vs ...uuid.UUID) predicate.Pkg {
	return predicate.Pkg(sql.FieldNotIn(FieldLocationID, vs...))
}

// LocationIDGT applies the GT predicate on the "location_id" field.
func LocationIDGT(v uuid.UUID) predicate.Pkg {
	return predicate.Pkg(sql.FieldGT(FieldLocationID, v))
}

// LocationIDGTE applies the GTE predicate on the "location_id" field.
func LocationIDGTE(v uuid.UUID) predicate.Pkg {
	return predicate.Pkg(sql.FieldGTE(FieldLocationID, v))
}

// LocationIDLT applies the LT predicate on the "location_id" field.
func LocationIDLT(v uuid.UUID) predicate.Pkg {
	return predicate.Pkg(sql.FieldLT(FieldLocationID, v))
}

// LocationIDLTE applies the LTE predicate on the "location_id" field.
func LocationIDLTE(v uuid.UUID) predicate.Pkg {
	return predicate.Pkg(sql.FieldLTE(FieldLocationID, v))
}

// LocationIDIsNil applies the IsNil predicate on the "location_id" field.
func LocationIDIsNil() predicate.Pkg {
	return predicate.Pkg(sql.FieldIsNull(FieldLocationID))
}

// LocationIDNotNil applies the NotNil predicate on the "location_id" field.
func LocationIDNotNil() predicate.Pkg {
	return predicate.Pkg(sql.FieldNotNull(FieldLocationID))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v int8) predicate.Pkg {
	return predicate.Pkg(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v int8) predicate.Pkg {
	return predicate.Pkg(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...int8) predicate.Pkg {
	return predicate.Pkg(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...int8) predicate.Pkg {
	return predicate.Pkg(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v int8) predicate.Pkg {
	return predicate.Pkg(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v int8) predicate.Pkg {
	return predicate.Pkg(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v int8) predicate.Pkg {
	return predicate.Pkg(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v int8) predicate.Pkg {
	return predicate.Pkg(sql.FieldLTE(FieldStatus, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Pkg {
	return predicate.Pkg(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Pkg {
	return predicate.Pkg(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Pkg {
	return predicate.Pkg(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Pkg {
	return predicate.Pkg(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Pkg {
	return predicate.Pkg(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Pkg {
	return predicate.Pkg(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Pkg {
	return predicate.Pkg(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Pkg {
	return predicate.Pkg(sql.FieldLTE(FieldCreatedAt, v))
}

// StartedAtEQ applies the EQ predicate on the "started_at" field.
func StartedAtEQ(v time.Time) predicate.Pkg {
	return predicate.Pkg(sql.FieldEQ(FieldStartedAt, v))
}

// StartedAtNEQ applies the NEQ predicate on the "started_at" field.
func StartedAtNEQ(v time.Time) predicate.Pkg {
	return predicate.Pkg(sql.FieldNEQ(FieldStartedAt, v))
}

// StartedAtIn applies the In predicate on the "started_at" field.
func StartedAtIn(vs ...time.Time) predicate.Pkg {
	return predicate.Pkg(sql.FieldIn(FieldStartedAt, vs...))
}

// StartedAtNotIn applies the NotIn predicate on the "started_at" field.
func StartedAtNotIn(vs ...time.Time) predicate.Pkg {
	return predicate.Pkg(sql.FieldNotIn(FieldStartedAt, vs...))
}

// StartedAtGT applies the GT predicate on the "started_at" field.
func StartedAtGT(v time.Time) predicate.Pkg {
	return predicate.Pkg(sql.FieldGT(FieldStartedAt, v))
}

// StartedAtGTE applies the GTE predicate on the "started_at" field.
func StartedAtGTE(v time.Time) predicate.Pkg {
	return predicate.Pkg(sql.FieldGTE(FieldStartedAt, v))
}

// StartedAtLT applies the LT predicate on the "started_at" field.
func StartedAtLT(v time.Time) predicate.Pkg {
	return predicate.Pkg(sql.FieldLT(FieldStartedAt, v))
}

// StartedAtLTE applies the LTE predicate on the "started_at" field.
func StartedAtLTE(v time.Time) predicate.Pkg {
	return predicate.Pkg(sql.FieldLTE(FieldStartedAt, v))
}

// StartedAtIsNil applies the IsNil predicate on the "started_at" field.
func StartedAtIsNil() predicate.Pkg {
	return predicate.Pkg(sql.FieldIsNull(FieldStartedAt))
}

// StartedAtNotNil applies the NotNil predicate on the "started_at" field.
func StartedAtNotNil() predicate.Pkg {
	return predicate.Pkg(sql.FieldNotNull(FieldStartedAt))
}

// CompletedAtEQ applies the EQ predicate on the "completed_at" field.
func CompletedAtEQ(v time.Time) predicate.Pkg {
	return predicate.Pkg(sql.FieldEQ(FieldCompletedAt, v))
}

// CompletedAtNEQ applies the NEQ predicate on the "completed_at" field.
func CompletedAtNEQ(v time.Time) predicate.Pkg {
	return predicate.Pkg(sql.FieldNEQ(FieldCompletedAt, v))
}

// CompletedAtIn applies the In predicate on the "completed_at" field.
func CompletedAtIn(vs ...time.Time) predicate.Pkg {
	return predicate.Pkg(sql.FieldIn(FieldCompletedAt, vs...))
}

// CompletedAtNotIn applies the NotIn predicate on the "completed_at" field.
func CompletedAtNotIn(vs ...time.Time) predicate.Pkg {
	return predicate.Pkg(sql.FieldNotIn(FieldCompletedAt, vs...))
}

// CompletedAtGT applies the GT predicate on the "completed_at" field.
func CompletedAtGT(v time.Time) predicate.Pkg {
	return predicate.Pkg(sql.FieldGT(FieldCompletedAt, v))
}

// CompletedAtGTE applies the GTE predicate on the "completed_at" field.
func CompletedAtGTE(v time.Time) predicate.Pkg {
	return predicate.Pkg(sql.FieldGTE(FieldCompletedAt, v))
}

// CompletedAtLT applies the LT predicate on the "completed_at" field.
func CompletedAtLT(v time.Time) predicate.Pkg {
	return predicate.Pkg(sql.FieldLT(FieldCompletedAt, v))
}

// CompletedAtLTE applies the LTE predicate on the "completed_at" field.
func CompletedAtLTE(v time.Time) predicate.Pkg {
	return predicate.Pkg(sql.FieldLTE(FieldCompletedAt, v))
}

// CompletedAtIsNil applies the IsNil predicate on the "completed_at" field.
func CompletedAtIsNil() predicate.Pkg {
	return predicate.Pkg(sql.FieldIsNull(FieldCompletedAt))
}

// CompletedAtNotNil applies the NotNil predicate on the "completed_at" field.
func CompletedAtNotNil() predicate.Pkg {
	return predicate.Pkg(sql.FieldNotNull(FieldCompletedAt))
}

// HasPreservationActions applies the HasEdge predicate on the "preservation_actions" edge.
func HasPreservationActions() predicate.Pkg {
	return predicate.Pkg(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PreservationActionsTable, PreservationActionsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPreservationActionsWith applies the HasEdge predicate on the "preservation_actions" edge with a given conditions (other predicates).
func HasPreservationActionsWith(preds ...predicate.PreservationAction) predicate.Pkg {
	return predicate.Pkg(func(s *sql.Selector) {
		step := newPreservationActionsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Pkg) predicate.Pkg {
	return predicate.Pkg(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Pkg) predicate.Pkg {
	return predicate.Pkg(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Pkg) predicate.Pkg {
	return predicate.Pkg(sql.NotPredicates(p))
}
