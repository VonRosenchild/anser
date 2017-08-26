/*
Anser Migrations

The anser package defines supports three differnt migration definition
forms: the SimpleMigration for single document updates, using
MongoDB's update query form, the ManualMigration which takes as input
a single document and provides an MongoDB session operation for manual
migration, and the AggregateProducer

Migrations themselves are executed as amboy.Jobs either serially or in
an amboy Queue. Although round-trippable serialization is not a strict
requirement of running migrations as amboy Jobs, these migrations
support round-trippable BSON serialization and thus distributed
queues.

Simple

Use simple migrations to rename a field in a document or change the
structure of a document using MongoDB queries. Prefer these operations
to running actual queries for the rate-limiting properties of the
Anser executor.

Manual

Use manual migrations when you need to perform a migration operation
that requires application logic, results in the creation of new
documents, or requires destructive modification of the source
document.

Stream

Use stream migrations for processing using application logic, an
iterator of documents. This is similar to the manual migration but
allows reduce-like operations, or even destructive operations.
*/

package anser

import (
	"github.com/mongodb/amboy"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Migration is a type alias for amboy.Job, used to identify
// migration-operations as distinct from other kinds of amboy.Jobs
type Migration amboy.Job

// ManualMigrationOperation defines the function object that performs
// the transformation in the manual migration migrations. Register
// these functions using RegisterManualMigrationOperation.
//
// Implementors of ManualMigrationOperations are responsible for
// implementating idempotent operations.
type ManualMigrationOperation func(*mgo.Session, bson.Raw) error