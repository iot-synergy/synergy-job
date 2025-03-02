// Code generated by ent, DO NOT EDIT.

package task

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/iot-synergy/synergy-job/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldUpdatedAt, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v uint8) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldStatus, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldName, v))
}

// TaskGroup applies equality check predicate on the "task_group" field. It's identical to TaskGroupEQ.
func TaskGroup(v string) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldTaskGroup, v))
}

// CronExpression applies equality check predicate on the "cron_expression" field. It's identical to CronExpressionEQ.
func CronExpression(v string) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldCronExpression, v))
}

// Pattern applies equality check predicate on the "pattern" field. It's identical to PatternEQ.
func Pattern(v string) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldPattern, v))
}

// Payload applies equality check predicate on the "payload" field. It's identical to PayloadEQ.
func Payload(v string) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldPayload, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldUpdatedAt, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v uint8) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v uint8) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...uint8) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...uint8) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v uint8) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v uint8) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v uint8) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v uint8) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldStatus, v))
}

// StatusIsNil applies the IsNil predicate on the "status" field.
func StatusIsNil() predicate.Task {
	return predicate.Task(sql.FieldIsNull(FieldStatus))
}

// StatusNotNil applies the NotNil predicate on the "status" field.
func StatusNotNil() predicate.Task {
	return predicate.Task(sql.FieldNotNull(FieldStatus))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Task {
	return predicate.Task(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Task {
	return predicate.Task(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Task {
	return predicate.Task(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Task {
	return predicate.Task(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Task {
	return predicate.Task(sql.FieldContainsFold(FieldName, v))
}

// TaskGroupEQ applies the EQ predicate on the "task_group" field.
func TaskGroupEQ(v string) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldTaskGroup, v))
}

// TaskGroupNEQ applies the NEQ predicate on the "task_group" field.
func TaskGroupNEQ(v string) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldTaskGroup, v))
}

// TaskGroupIn applies the In predicate on the "task_group" field.
func TaskGroupIn(vs ...string) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldTaskGroup, vs...))
}

// TaskGroupNotIn applies the NotIn predicate on the "task_group" field.
func TaskGroupNotIn(vs ...string) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldTaskGroup, vs...))
}

// TaskGroupGT applies the GT predicate on the "task_group" field.
func TaskGroupGT(v string) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldTaskGroup, v))
}

// TaskGroupGTE applies the GTE predicate on the "task_group" field.
func TaskGroupGTE(v string) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldTaskGroup, v))
}

// TaskGroupLT applies the LT predicate on the "task_group" field.
func TaskGroupLT(v string) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldTaskGroup, v))
}

// TaskGroupLTE applies the LTE predicate on the "task_group" field.
func TaskGroupLTE(v string) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldTaskGroup, v))
}

// TaskGroupContains applies the Contains predicate on the "task_group" field.
func TaskGroupContains(v string) predicate.Task {
	return predicate.Task(sql.FieldContains(FieldTaskGroup, v))
}

// TaskGroupHasPrefix applies the HasPrefix predicate on the "task_group" field.
func TaskGroupHasPrefix(v string) predicate.Task {
	return predicate.Task(sql.FieldHasPrefix(FieldTaskGroup, v))
}

// TaskGroupHasSuffix applies the HasSuffix predicate on the "task_group" field.
func TaskGroupHasSuffix(v string) predicate.Task {
	return predicate.Task(sql.FieldHasSuffix(FieldTaskGroup, v))
}

// TaskGroupEqualFold applies the EqualFold predicate on the "task_group" field.
func TaskGroupEqualFold(v string) predicate.Task {
	return predicate.Task(sql.FieldEqualFold(FieldTaskGroup, v))
}

// TaskGroupContainsFold applies the ContainsFold predicate on the "task_group" field.
func TaskGroupContainsFold(v string) predicate.Task {
	return predicate.Task(sql.FieldContainsFold(FieldTaskGroup, v))
}

// CronExpressionEQ applies the EQ predicate on the "cron_expression" field.
func CronExpressionEQ(v string) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldCronExpression, v))
}

// CronExpressionNEQ applies the NEQ predicate on the "cron_expression" field.
func CronExpressionNEQ(v string) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldCronExpression, v))
}

// CronExpressionIn applies the In predicate on the "cron_expression" field.
func CronExpressionIn(vs ...string) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldCronExpression, vs...))
}

// CronExpressionNotIn applies the NotIn predicate on the "cron_expression" field.
func CronExpressionNotIn(vs ...string) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldCronExpression, vs...))
}

// CronExpressionGT applies the GT predicate on the "cron_expression" field.
func CronExpressionGT(v string) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldCronExpression, v))
}

// CronExpressionGTE applies the GTE predicate on the "cron_expression" field.
func CronExpressionGTE(v string) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldCronExpression, v))
}

// CronExpressionLT applies the LT predicate on the "cron_expression" field.
func CronExpressionLT(v string) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldCronExpression, v))
}

// CronExpressionLTE applies the LTE predicate on the "cron_expression" field.
func CronExpressionLTE(v string) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldCronExpression, v))
}

// CronExpressionContains applies the Contains predicate on the "cron_expression" field.
func CronExpressionContains(v string) predicate.Task {
	return predicate.Task(sql.FieldContains(FieldCronExpression, v))
}

// CronExpressionHasPrefix applies the HasPrefix predicate on the "cron_expression" field.
func CronExpressionHasPrefix(v string) predicate.Task {
	return predicate.Task(sql.FieldHasPrefix(FieldCronExpression, v))
}

// CronExpressionHasSuffix applies the HasSuffix predicate on the "cron_expression" field.
func CronExpressionHasSuffix(v string) predicate.Task {
	return predicate.Task(sql.FieldHasSuffix(FieldCronExpression, v))
}

// CronExpressionEqualFold applies the EqualFold predicate on the "cron_expression" field.
func CronExpressionEqualFold(v string) predicate.Task {
	return predicate.Task(sql.FieldEqualFold(FieldCronExpression, v))
}

// CronExpressionContainsFold applies the ContainsFold predicate on the "cron_expression" field.
func CronExpressionContainsFold(v string) predicate.Task {
	return predicate.Task(sql.FieldContainsFold(FieldCronExpression, v))
}

// PatternEQ applies the EQ predicate on the "pattern" field.
func PatternEQ(v string) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldPattern, v))
}

// PatternNEQ applies the NEQ predicate on the "pattern" field.
func PatternNEQ(v string) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldPattern, v))
}

// PatternIn applies the In predicate on the "pattern" field.
func PatternIn(vs ...string) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldPattern, vs...))
}

// PatternNotIn applies the NotIn predicate on the "pattern" field.
func PatternNotIn(vs ...string) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldPattern, vs...))
}

// PatternGT applies the GT predicate on the "pattern" field.
func PatternGT(v string) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldPattern, v))
}

// PatternGTE applies the GTE predicate on the "pattern" field.
func PatternGTE(v string) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldPattern, v))
}

// PatternLT applies the LT predicate on the "pattern" field.
func PatternLT(v string) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldPattern, v))
}

// PatternLTE applies the LTE predicate on the "pattern" field.
func PatternLTE(v string) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldPattern, v))
}

// PatternContains applies the Contains predicate on the "pattern" field.
func PatternContains(v string) predicate.Task {
	return predicate.Task(sql.FieldContains(FieldPattern, v))
}

// PatternHasPrefix applies the HasPrefix predicate on the "pattern" field.
func PatternHasPrefix(v string) predicate.Task {
	return predicate.Task(sql.FieldHasPrefix(FieldPattern, v))
}

// PatternHasSuffix applies the HasSuffix predicate on the "pattern" field.
func PatternHasSuffix(v string) predicate.Task {
	return predicate.Task(sql.FieldHasSuffix(FieldPattern, v))
}

// PatternEqualFold applies the EqualFold predicate on the "pattern" field.
func PatternEqualFold(v string) predicate.Task {
	return predicate.Task(sql.FieldEqualFold(FieldPattern, v))
}

// PatternContainsFold applies the ContainsFold predicate on the "pattern" field.
func PatternContainsFold(v string) predicate.Task {
	return predicate.Task(sql.FieldContainsFold(FieldPattern, v))
}

// PayloadEQ applies the EQ predicate on the "payload" field.
func PayloadEQ(v string) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldPayload, v))
}

// PayloadNEQ applies the NEQ predicate on the "payload" field.
func PayloadNEQ(v string) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldPayload, v))
}

// PayloadIn applies the In predicate on the "payload" field.
func PayloadIn(vs ...string) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldPayload, vs...))
}

// PayloadNotIn applies the NotIn predicate on the "payload" field.
func PayloadNotIn(vs ...string) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldPayload, vs...))
}

// PayloadGT applies the GT predicate on the "payload" field.
func PayloadGT(v string) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldPayload, v))
}

// PayloadGTE applies the GTE predicate on the "payload" field.
func PayloadGTE(v string) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldPayload, v))
}

// PayloadLT applies the LT predicate on the "payload" field.
func PayloadLT(v string) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldPayload, v))
}

// PayloadLTE applies the LTE predicate on the "payload" field.
func PayloadLTE(v string) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldPayload, v))
}

// PayloadContains applies the Contains predicate on the "payload" field.
func PayloadContains(v string) predicate.Task {
	return predicate.Task(sql.FieldContains(FieldPayload, v))
}

// PayloadHasPrefix applies the HasPrefix predicate on the "payload" field.
func PayloadHasPrefix(v string) predicate.Task {
	return predicate.Task(sql.FieldHasPrefix(FieldPayload, v))
}

// PayloadHasSuffix applies the HasSuffix predicate on the "payload" field.
func PayloadHasSuffix(v string) predicate.Task {
	return predicate.Task(sql.FieldHasSuffix(FieldPayload, v))
}

// PayloadEqualFold applies the EqualFold predicate on the "payload" field.
func PayloadEqualFold(v string) predicate.Task {
	return predicate.Task(sql.FieldEqualFold(FieldPayload, v))
}

// PayloadContainsFold applies the ContainsFold predicate on the "payload" field.
func PayloadContainsFold(v string) predicate.Task {
	return predicate.Task(sql.FieldContainsFold(FieldPayload, v))
}

// HasTaskLogs applies the HasEdge predicate on the "task_logs" edge.
func HasTaskLogs() predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TaskLogsTable, TaskLogsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTaskLogsWith applies the HasEdge predicate on the "task_logs" edge with a given conditions (other predicates).
func HasTaskLogsWith(preds ...predicate.TaskLog) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		step := newTaskLogsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Task) predicate.Task {
	return predicate.Task(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Task) predicate.Task {
	return predicate.Task(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Task) predicate.Task {
	return predicate.Task(sql.NotPredicates(p))
}
