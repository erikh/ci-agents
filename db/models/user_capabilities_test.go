// Code generated by SQLBoiler 4.5.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testUserCapabilities(t *testing.T) {
	t.Parallel()

	query := UserCapabilities()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testUserCapabilitiesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserCapability{}
	if err = randomize.Struct(seed, o, userCapabilityDBTypes, true, userCapabilityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserCapability struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := UserCapabilities().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUserCapabilitiesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserCapability{}
	if err = randomize.Struct(seed, o, userCapabilityDBTypes, true, userCapabilityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserCapability struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := UserCapabilities().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := UserCapabilities().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUserCapabilitiesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserCapability{}
	if err = randomize.Struct(seed, o, userCapabilityDBTypes, true, userCapabilityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserCapability struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := UserCapabilitySlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := UserCapabilities().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUserCapabilitiesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserCapability{}
	if err = randomize.Struct(seed, o, userCapabilityDBTypes, true, userCapabilityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserCapability struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := UserCapabilityExists(ctx, tx, o.UserID, o.Name)
	if err != nil {
		t.Errorf("Unable to check if UserCapability exists: %s", err)
	}
	if !e {
		t.Errorf("Expected UserCapabilityExists to return true, but got false.")
	}
}

func testUserCapabilitiesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserCapability{}
	if err = randomize.Struct(seed, o, userCapabilityDBTypes, true, userCapabilityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserCapability struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	userCapabilityFound, err := FindUserCapability(ctx, tx, o.UserID, o.Name)
	if err != nil {
		t.Error(err)
	}

	if userCapabilityFound == nil {
		t.Error("want a record, got nil")
	}
}

func testUserCapabilitiesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserCapability{}
	if err = randomize.Struct(seed, o, userCapabilityDBTypes, true, userCapabilityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserCapability struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = UserCapabilities().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testUserCapabilitiesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserCapability{}
	if err = randomize.Struct(seed, o, userCapabilityDBTypes, true, userCapabilityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserCapability struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := UserCapabilities().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testUserCapabilitiesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	userCapabilityOne := &UserCapability{}
	userCapabilityTwo := &UserCapability{}
	if err = randomize.Struct(seed, userCapabilityOne, userCapabilityDBTypes, false, userCapabilityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserCapability struct: %s", err)
	}
	if err = randomize.Struct(seed, userCapabilityTwo, userCapabilityDBTypes, false, userCapabilityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserCapability struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = userCapabilityOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = userCapabilityTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := UserCapabilities().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testUserCapabilitiesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	userCapabilityOne := &UserCapability{}
	userCapabilityTwo := &UserCapability{}
	if err = randomize.Struct(seed, userCapabilityOne, userCapabilityDBTypes, false, userCapabilityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserCapability struct: %s", err)
	}
	if err = randomize.Struct(seed, userCapabilityTwo, userCapabilityDBTypes, false, userCapabilityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserCapability struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = userCapabilityOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = userCapabilityTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UserCapabilities().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func userCapabilityBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *UserCapability) error {
	*o = UserCapability{}
	return nil
}

func userCapabilityAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *UserCapability) error {
	*o = UserCapability{}
	return nil
}

func userCapabilityAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *UserCapability) error {
	*o = UserCapability{}
	return nil
}

func userCapabilityBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *UserCapability) error {
	*o = UserCapability{}
	return nil
}

func userCapabilityAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *UserCapability) error {
	*o = UserCapability{}
	return nil
}

func userCapabilityBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *UserCapability) error {
	*o = UserCapability{}
	return nil
}

func userCapabilityAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *UserCapability) error {
	*o = UserCapability{}
	return nil
}

func userCapabilityBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *UserCapability) error {
	*o = UserCapability{}
	return nil
}

func userCapabilityAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *UserCapability) error {
	*o = UserCapability{}
	return nil
}

func testUserCapabilitiesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &UserCapability{}
	o := &UserCapability{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, userCapabilityDBTypes, false); err != nil {
		t.Errorf("Unable to randomize UserCapability object: %s", err)
	}

	AddUserCapabilityHook(boil.BeforeInsertHook, userCapabilityBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	userCapabilityBeforeInsertHooks = []UserCapabilityHook{}

	AddUserCapabilityHook(boil.AfterInsertHook, userCapabilityAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	userCapabilityAfterInsertHooks = []UserCapabilityHook{}

	AddUserCapabilityHook(boil.AfterSelectHook, userCapabilityAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	userCapabilityAfterSelectHooks = []UserCapabilityHook{}

	AddUserCapabilityHook(boil.BeforeUpdateHook, userCapabilityBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	userCapabilityBeforeUpdateHooks = []UserCapabilityHook{}

	AddUserCapabilityHook(boil.AfterUpdateHook, userCapabilityAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	userCapabilityAfterUpdateHooks = []UserCapabilityHook{}

	AddUserCapabilityHook(boil.BeforeDeleteHook, userCapabilityBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	userCapabilityBeforeDeleteHooks = []UserCapabilityHook{}

	AddUserCapabilityHook(boil.AfterDeleteHook, userCapabilityAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	userCapabilityAfterDeleteHooks = []UserCapabilityHook{}

	AddUserCapabilityHook(boil.BeforeUpsertHook, userCapabilityBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	userCapabilityBeforeUpsertHooks = []UserCapabilityHook{}

	AddUserCapabilityHook(boil.AfterUpsertHook, userCapabilityAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	userCapabilityAfterUpsertHooks = []UserCapabilityHook{}
}

func testUserCapabilitiesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserCapability{}
	if err = randomize.Struct(seed, o, userCapabilityDBTypes, true, userCapabilityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserCapability struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UserCapabilities().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testUserCapabilitiesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserCapability{}
	if err = randomize.Struct(seed, o, userCapabilityDBTypes, true); err != nil {
		t.Errorf("Unable to randomize UserCapability struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(userCapabilityColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := UserCapabilities().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testUserCapabilityToOneUserUsingUser(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local UserCapability
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, userCapabilityDBTypes, false, userCapabilityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserCapability struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.UserID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.User().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := UserCapabilitySlice{&local}
	if err = local.L.LoadUser(ctx, tx, false, (*[]*UserCapability)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.User = nil
	if err = local.L.LoadUser(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testUserCapabilityToOneSetOpUserUsingUser(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a UserCapability
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, userCapabilityDBTypes, false, strmangle.SetComplement(userCapabilityPrimaryKeyColumns, userCapabilityColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*User{&b, &c} {
		err = a.SetUser(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.User != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.UserCapabilities[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.UserID != x.ID {
			t.Error("foreign key was wrong value", a.UserID)
		}

		if exists, err := UserCapabilityExists(ctx, tx, a.UserID, a.Name); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}

func testUserCapabilitiesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserCapability{}
	if err = randomize.Struct(seed, o, userCapabilityDBTypes, true, userCapabilityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserCapability struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testUserCapabilitiesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserCapability{}
	if err = randomize.Struct(seed, o, userCapabilityDBTypes, true, userCapabilityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserCapability struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := UserCapabilitySlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testUserCapabilitiesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserCapability{}
	if err = randomize.Struct(seed, o, userCapabilityDBTypes, true, userCapabilityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserCapability struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := UserCapabilities().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	userCapabilityDBTypes = map[string]string{`UserID`: `bigint`, `Name`: `character varying`}
	_                     = bytes.MinRead
)

func testUserCapabilitiesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(userCapabilityPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(userCapabilityAllColumns) == len(userCapabilityPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &UserCapability{}
	if err = randomize.Struct(seed, o, userCapabilityDBTypes, true, userCapabilityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserCapability struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UserCapabilities().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, userCapabilityDBTypes, true, userCapabilityPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize UserCapability struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testUserCapabilitiesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(userCapabilityAllColumns) == len(userCapabilityPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &UserCapability{}
	if err = randomize.Struct(seed, o, userCapabilityDBTypes, true, userCapabilityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserCapability struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UserCapabilities().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, userCapabilityDBTypes, true, userCapabilityPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize UserCapability struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(userCapabilityAllColumns, userCapabilityPrimaryKeyColumns) {
		fields = userCapabilityAllColumns
	} else {
		fields = strmangle.SetComplement(
			userCapabilityAllColumns,
			userCapabilityPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := UserCapabilitySlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testUserCapabilitiesUpsert(t *testing.T) {
	t.Parallel()

	if len(userCapabilityAllColumns) == len(userCapabilityPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := UserCapability{}
	if err = randomize.Struct(seed, &o, userCapabilityDBTypes, true); err != nil {
		t.Errorf("Unable to randomize UserCapability struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert UserCapability: %s", err)
	}

	count, err := UserCapabilities().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, userCapabilityDBTypes, false, userCapabilityPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize UserCapability struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert UserCapability: %s", err)
	}

	count, err = UserCapabilities().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
