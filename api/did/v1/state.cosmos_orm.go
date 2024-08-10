// Code generated by protoc-gen-go-cosmos-orm. DO NOT EDIT.

package didv1

import (
	context "context"
	ormlist "cosmossdk.io/orm/model/ormlist"
	ormtable "cosmossdk.io/orm/model/ormtable"
	ormerrors "cosmossdk.io/orm/types/ormerrors"
)

type AliasesTable interface {
	Insert(ctx context.Context, aliases *Aliases) error
	Update(ctx context.Context, aliases *Aliases) error
	Save(ctx context.Context, aliases *Aliases) error
	Delete(ctx context.Context, aliases *Aliases) error
	Has(ctx context.Context, id string) (found bool, err error)
	// Get returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	Get(ctx context.Context, id string) (*Aliases, error)
	HasBySubject(ctx context.Context, subject string) (found bool, err error)
	// GetBySubject returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	GetBySubject(ctx context.Context, subject string) (*Aliases, error)
	List(ctx context.Context, prefixKey AliasesIndexKey, opts ...ormlist.Option) (AliasesIterator, error)
	ListRange(ctx context.Context, from, to AliasesIndexKey, opts ...ormlist.Option) (AliasesIterator, error)
	DeleteBy(ctx context.Context, prefixKey AliasesIndexKey) error
	DeleteRange(ctx context.Context, from, to AliasesIndexKey) error

	doNotImplement()
}

type AliasesIterator struct {
	ormtable.Iterator
}

func (i AliasesIterator) Value() (*Aliases, error) {
	var aliases Aliases
	err := i.UnmarshalMessage(&aliases)
	return &aliases, err
}

type AliasesIndexKey interface {
	id() uint32
	values() []interface{}
	aliasesIndexKey()
}

// primary key starting index..
type AliasesPrimaryKey = AliasesIdIndexKey

type AliasesIdIndexKey struct {
	vs []interface{}
}

func (x AliasesIdIndexKey) id() uint32            { return 0 }
func (x AliasesIdIndexKey) values() []interface{} { return x.vs }
func (x AliasesIdIndexKey) aliasesIndexKey()      {}

func (this AliasesIdIndexKey) WithId(id string) AliasesIdIndexKey {
	this.vs = []interface{}{id}
	return this
}

type AliasesSubjectIndexKey struct {
	vs []interface{}
}

func (x AliasesSubjectIndexKey) id() uint32            { return 1 }
func (x AliasesSubjectIndexKey) values() []interface{} { return x.vs }
func (x AliasesSubjectIndexKey) aliasesIndexKey()      {}

func (this AliasesSubjectIndexKey) WithSubject(subject string) AliasesSubjectIndexKey {
	this.vs = []interface{}{subject}
	return this
}

type aliasesTable struct {
	table ormtable.Table
}

func (this aliasesTable) Insert(ctx context.Context, aliases *Aliases) error {
	return this.table.Insert(ctx, aliases)
}

func (this aliasesTable) Update(ctx context.Context, aliases *Aliases) error {
	return this.table.Update(ctx, aliases)
}

func (this aliasesTable) Save(ctx context.Context, aliases *Aliases) error {
	return this.table.Save(ctx, aliases)
}

func (this aliasesTable) Delete(ctx context.Context, aliases *Aliases) error {
	return this.table.Delete(ctx, aliases)
}

func (this aliasesTable) Has(ctx context.Context, id string) (found bool, err error) {
	return this.table.PrimaryKey().Has(ctx, id)
}

func (this aliasesTable) Get(ctx context.Context, id string) (*Aliases, error) {
	var aliases Aliases
	found, err := this.table.PrimaryKey().Get(ctx, &aliases, id)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &aliases, nil
}

func (this aliasesTable) HasBySubject(ctx context.Context, subject string) (found bool, err error) {
	return this.table.GetIndexByID(1).(ormtable.UniqueIndex).Has(ctx,
		subject,
	)
}

func (this aliasesTable) GetBySubject(ctx context.Context, subject string) (*Aliases, error) {
	var aliases Aliases
	found, err := this.table.GetIndexByID(1).(ormtable.UniqueIndex).Get(ctx, &aliases,
		subject,
	)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &aliases, nil
}

func (this aliasesTable) List(ctx context.Context, prefixKey AliasesIndexKey, opts ...ormlist.Option) (AliasesIterator, error) {
	it, err := this.table.GetIndexByID(prefixKey.id()).List(ctx, prefixKey.values(), opts...)
	return AliasesIterator{it}, err
}

func (this aliasesTable) ListRange(ctx context.Context, from, to AliasesIndexKey, opts ...ormlist.Option) (AliasesIterator, error) {
	it, err := this.table.GetIndexByID(from.id()).ListRange(ctx, from.values(), to.values(), opts...)
	return AliasesIterator{it}, err
}

func (this aliasesTable) DeleteBy(ctx context.Context, prefixKey AliasesIndexKey) error {
	return this.table.GetIndexByID(prefixKey.id()).DeleteBy(ctx, prefixKey.values()...)
}

func (this aliasesTable) DeleteRange(ctx context.Context, from, to AliasesIndexKey) error {
	return this.table.GetIndexByID(from.id()).DeleteRange(ctx, from.values(), to.values())
}

func (this aliasesTable) doNotImplement() {}

var _ AliasesTable = aliasesTable{}

func NewAliasesTable(db ormtable.Schema) (AliasesTable, error) {
	table := db.GetTable(&Aliases{})
	if table == nil {
		return nil, ormerrors.TableNotFound.Wrap(string((&Aliases{}).ProtoReflect().Descriptor().FullName()))
	}
	return aliasesTable{table}, nil
}

type AssertionTable interface {
	Insert(ctx context.Context, assertion *Assertion) error
	Update(ctx context.Context, assertion *Assertion) error
	Save(ctx context.Context, assertion *Assertion) error
	Delete(ctx context.Context, assertion *Assertion) error
	Has(ctx context.Context, id string) (found bool, err error)
	// Get returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	Get(ctx context.Context, id string) (*Assertion, error)
	List(ctx context.Context, prefixKey AssertionIndexKey, opts ...ormlist.Option) (AssertionIterator, error)
	ListRange(ctx context.Context, from, to AssertionIndexKey, opts ...ormlist.Option) (AssertionIterator, error)
	DeleteBy(ctx context.Context, prefixKey AssertionIndexKey) error
	DeleteRange(ctx context.Context, from, to AssertionIndexKey) error

	doNotImplement()
}

type AssertionIterator struct {
	ormtable.Iterator
}

func (i AssertionIterator) Value() (*Assertion, error) {
	var assertion Assertion
	err := i.UnmarshalMessage(&assertion)
	return &assertion, err
}

type AssertionIndexKey interface {
	id() uint32
	values() []interface{}
	assertionIndexKey()
}

// primary key starting index..
type AssertionPrimaryKey = AssertionIdIndexKey

type AssertionIdIndexKey struct {
	vs []interface{}
}

func (x AssertionIdIndexKey) id() uint32            { return 0 }
func (x AssertionIdIndexKey) values() []interface{} { return x.vs }
func (x AssertionIdIndexKey) assertionIndexKey()    {}

func (this AssertionIdIndexKey) WithId(id string) AssertionIdIndexKey {
	this.vs = []interface{}{id}
	return this
}

type assertionTable struct {
	table ormtable.Table
}

func (this assertionTable) Insert(ctx context.Context, assertion *Assertion) error {
	return this.table.Insert(ctx, assertion)
}

func (this assertionTable) Update(ctx context.Context, assertion *Assertion) error {
	return this.table.Update(ctx, assertion)
}

func (this assertionTable) Save(ctx context.Context, assertion *Assertion) error {
	return this.table.Save(ctx, assertion)
}

func (this assertionTable) Delete(ctx context.Context, assertion *Assertion) error {
	return this.table.Delete(ctx, assertion)
}

func (this assertionTable) Has(ctx context.Context, id string) (found bool, err error) {
	return this.table.PrimaryKey().Has(ctx, id)
}

func (this assertionTable) Get(ctx context.Context, id string) (*Assertion, error) {
	var assertion Assertion
	found, err := this.table.PrimaryKey().Get(ctx, &assertion, id)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &assertion, nil
}

func (this assertionTable) List(ctx context.Context, prefixKey AssertionIndexKey, opts ...ormlist.Option) (AssertionIterator, error) {
	it, err := this.table.GetIndexByID(prefixKey.id()).List(ctx, prefixKey.values(), opts...)
	return AssertionIterator{it}, err
}

func (this assertionTable) ListRange(ctx context.Context, from, to AssertionIndexKey, opts ...ormlist.Option) (AssertionIterator, error) {
	it, err := this.table.GetIndexByID(from.id()).ListRange(ctx, from.values(), to.values(), opts...)
	return AssertionIterator{it}, err
}

func (this assertionTable) DeleteBy(ctx context.Context, prefixKey AssertionIndexKey) error {
	return this.table.GetIndexByID(prefixKey.id()).DeleteBy(ctx, prefixKey.values()...)
}

func (this assertionTable) DeleteRange(ctx context.Context, from, to AssertionIndexKey) error {
	return this.table.GetIndexByID(from.id()).DeleteRange(ctx, from.values(), to.values())
}

func (this assertionTable) doNotImplement() {}

var _ AssertionTable = assertionTable{}

func NewAssertionTable(db ormtable.Schema) (AssertionTable, error) {
	table := db.GetTable(&Assertion{})
	if table == nil {
		return nil, ormerrors.TableNotFound.Wrap(string((&Assertion{}).ProtoReflect().Descriptor().FullName()))
	}
	return assertionTable{table}, nil
}

type AttestationTable interface {
	Insert(ctx context.Context, attestation *Attestation) error
	Update(ctx context.Context, attestation *Attestation) error
	Save(ctx context.Context, attestation *Attestation) error
	Delete(ctx context.Context, attestation *Attestation) error
	Has(ctx context.Context, id string) (found bool, err error)
	// Get returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	Get(ctx context.Context, id string) (*Attestation, error)
	List(ctx context.Context, prefixKey AttestationIndexKey, opts ...ormlist.Option) (AttestationIterator, error)
	ListRange(ctx context.Context, from, to AttestationIndexKey, opts ...ormlist.Option) (AttestationIterator, error)
	DeleteBy(ctx context.Context, prefixKey AttestationIndexKey) error
	DeleteRange(ctx context.Context, from, to AttestationIndexKey) error

	doNotImplement()
}

type AttestationIterator struct {
	ormtable.Iterator
}

func (i AttestationIterator) Value() (*Attestation, error) {
	var attestation Attestation
	err := i.UnmarshalMessage(&attestation)
	return &attestation, err
}

type AttestationIndexKey interface {
	id() uint32
	values() []interface{}
	attestationIndexKey()
}

// primary key starting index..
type AttestationPrimaryKey = AttestationIdIndexKey

type AttestationIdIndexKey struct {
	vs []interface{}
}

func (x AttestationIdIndexKey) id() uint32            { return 0 }
func (x AttestationIdIndexKey) values() []interface{} { return x.vs }
func (x AttestationIdIndexKey) attestationIndexKey()  {}

func (this AttestationIdIndexKey) WithId(id string) AttestationIdIndexKey {
	this.vs = []interface{}{id}
	return this
}

type attestationTable struct {
	table ormtable.Table
}

func (this attestationTable) Insert(ctx context.Context, attestation *Attestation) error {
	return this.table.Insert(ctx, attestation)
}

func (this attestationTable) Update(ctx context.Context, attestation *Attestation) error {
	return this.table.Update(ctx, attestation)
}

func (this attestationTable) Save(ctx context.Context, attestation *Attestation) error {
	return this.table.Save(ctx, attestation)
}

func (this attestationTable) Delete(ctx context.Context, attestation *Attestation) error {
	return this.table.Delete(ctx, attestation)
}

func (this attestationTable) Has(ctx context.Context, id string) (found bool, err error) {
	return this.table.PrimaryKey().Has(ctx, id)
}

func (this attestationTable) Get(ctx context.Context, id string) (*Attestation, error) {
	var attestation Attestation
	found, err := this.table.PrimaryKey().Get(ctx, &attestation, id)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &attestation, nil
}

func (this attestationTable) List(ctx context.Context, prefixKey AttestationIndexKey, opts ...ormlist.Option) (AttestationIterator, error) {
	it, err := this.table.GetIndexByID(prefixKey.id()).List(ctx, prefixKey.values(), opts...)
	return AttestationIterator{it}, err
}

func (this attestationTable) ListRange(ctx context.Context, from, to AttestationIndexKey, opts ...ormlist.Option) (AttestationIterator, error) {
	it, err := this.table.GetIndexByID(from.id()).ListRange(ctx, from.values(), to.values(), opts...)
	return AttestationIterator{it}, err
}

func (this attestationTable) DeleteBy(ctx context.Context, prefixKey AttestationIndexKey) error {
	return this.table.GetIndexByID(prefixKey.id()).DeleteBy(ctx, prefixKey.values()...)
}

func (this attestationTable) DeleteRange(ctx context.Context, from, to AttestationIndexKey) error {
	return this.table.GetIndexByID(from.id()).DeleteRange(ctx, from.values(), to.values())
}

func (this attestationTable) doNotImplement() {}

var _ AttestationTable = attestationTable{}

func NewAttestationTable(db ormtable.Schema) (AttestationTable, error) {
	table := db.GetTable(&Attestation{})
	if table == nil {
		return nil, ormerrors.TableNotFound.Wrap(string((&Attestation{}).ProtoReflect().Descriptor().FullName()))
	}
	return attestationTable{table}, nil
}

type ControllerTable interface {
	Insert(ctx context.Context, controller *Controller) error
	Update(ctx context.Context, controller *Controller) error
	Save(ctx context.Context, controller *Controller) error
	Delete(ctx context.Context, controller *Controller) error
	Has(ctx context.Context, id string) (found bool, err error)
	// Get returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	Get(ctx context.Context, id string) (*Controller, error)
	List(ctx context.Context, prefixKey ControllerIndexKey, opts ...ormlist.Option) (ControllerIterator, error)
	ListRange(ctx context.Context, from, to ControllerIndexKey, opts ...ormlist.Option) (ControllerIterator, error)
	DeleteBy(ctx context.Context, prefixKey ControllerIndexKey) error
	DeleteRange(ctx context.Context, from, to ControllerIndexKey) error

	doNotImplement()
}

type ControllerIterator struct {
	ormtable.Iterator
}

func (i ControllerIterator) Value() (*Controller, error) {
	var controller Controller
	err := i.UnmarshalMessage(&controller)
	return &controller, err
}

type ControllerIndexKey interface {
	id() uint32
	values() []interface{}
	controllerIndexKey()
}

// primary key starting index..
type ControllerPrimaryKey = ControllerIdIndexKey

type ControllerIdIndexKey struct {
	vs []interface{}
}

func (x ControllerIdIndexKey) id() uint32            { return 0 }
func (x ControllerIdIndexKey) values() []interface{} { return x.vs }
func (x ControllerIdIndexKey) controllerIndexKey()   {}

func (this ControllerIdIndexKey) WithId(id string) ControllerIdIndexKey {
	this.vs = []interface{}{id}
	return this
}

type controllerTable struct {
	table ormtable.Table
}

func (this controllerTable) Insert(ctx context.Context, controller *Controller) error {
	return this.table.Insert(ctx, controller)
}

func (this controllerTable) Update(ctx context.Context, controller *Controller) error {
	return this.table.Update(ctx, controller)
}

func (this controllerTable) Save(ctx context.Context, controller *Controller) error {
	return this.table.Save(ctx, controller)
}

func (this controllerTable) Delete(ctx context.Context, controller *Controller) error {
	return this.table.Delete(ctx, controller)
}

func (this controllerTable) Has(ctx context.Context, id string) (found bool, err error) {
	return this.table.PrimaryKey().Has(ctx, id)
}

func (this controllerTable) Get(ctx context.Context, id string) (*Controller, error) {
	var controller Controller
	found, err := this.table.PrimaryKey().Get(ctx, &controller, id)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &controller, nil
}

func (this controllerTable) List(ctx context.Context, prefixKey ControllerIndexKey, opts ...ormlist.Option) (ControllerIterator, error) {
	it, err := this.table.GetIndexByID(prefixKey.id()).List(ctx, prefixKey.values(), opts...)
	return ControllerIterator{it}, err
}

func (this controllerTable) ListRange(ctx context.Context, from, to ControllerIndexKey, opts ...ormlist.Option) (ControllerIterator, error) {
	it, err := this.table.GetIndexByID(from.id()).ListRange(ctx, from.values(), to.values(), opts...)
	return ControllerIterator{it}, err
}

func (this controllerTable) DeleteBy(ctx context.Context, prefixKey ControllerIndexKey) error {
	return this.table.GetIndexByID(prefixKey.id()).DeleteBy(ctx, prefixKey.values()...)
}

func (this controllerTable) DeleteRange(ctx context.Context, from, to ControllerIndexKey) error {
	return this.table.GetIndexByID(from.id()).DeleteRange(ctx, from.values(), to.values())
}

func (this controllerTable) doNotImplement() {}

var _ ControllerTable = controllerTable{}

func NewControllerTable(db ormtable.Schema) (ControllerTable, error) {
	table := db.GetTable(&Controller{})
	if table == nil {
		return nil, ormerrors.TableNotFound.Wrap(string((&Controller{}).ProtoReflect().Descriptor().FullName()))
	}
	return controllerTable{table}, nil
}

type DelegationTable interface {
	Insert(ctx context.Context, delegation *Delegation) error
	Update(ctx context.Context, delegation *Delegation) error
	Save(ctx context.Context, delegation *Delegation) error
	Delete(ctx context.Context, delegation *Delegation) error
	Has(ctx context.Context, id string) (found bool, err error)
	// Get returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	Get(ctx context.Context, id string) (*Delegation, error)
	List(ctx context.Context, prefixKey DelegationIndexKey, opts ...ormlist.Option) (DelegationIterator, error)
	ListRange(ctx context.Context, from, to DelegationIndexKey, opts ...ormlist.Option) (DelegationIterator, error)
	DeleteBy(ctx context.Context, prefixKey DelegationIndexKey) error
	DeleteRange(ctx context.Context, from, to DelegationIndexKey) error

	doNotImplement()
}

type DelegationIterator struct {
	ormtable.Iterator
}

func (i DelegationIterator) Value() (*Delegation, error) {
	var delegation Delegation
	err := i.UnmarshalMessage(&delegation)
	return &delegation, err
}

type DelegationIndexKey interface {
	id() uint32
	values() []interface{}
	delegationIndexKey()
}

// primary key starting index..
type DelegationPrimaryKey = DelegationIdIndexKey

type DelegationIdIndexKey struct {
	vs []interface{}
}

func (x DelegationIdIndexKey) id() uint32            { return 0 }
func (x DelegationIdIndexKey) values() []interface{} { return x.vs }
func (x DelegationIdIndexKey) delegationIndexKey()   {}

func (this DelegationIdIndexKey) WithId(id string) DelegationIdIndexKey {
	this.vs = []interface{}{id}
	return this
}

type delegationTable struct {
	table ormtable.Table
}

func (this delegationTable) Insert(ctx context.Context, delegation *Delegation) error {
	return this.table.Insert(ctx, delegation)
}

func (this delegationTable) Update(ctx context.Context, delegation *Delegation) error {
	return this.table.Update(ctx, delegation)
}

func (this delegationTable) Save(ctx context.Context, delegation *Delegation) error {
	return this.table.Save(ctx, delegation)
}

func (this delegationTable) Delete(ctx context.Context, delegation *Delegation) error {
	return this.table.Delete(ctx, delegation)
}

func (this delegationTable) Has(ctx context.Context, id string) (found bool, err error) {
	return this.table.PrimaryKey().Has(ctx, id)
}

func (this delegationTable) Get(ctx context.Context, id string) (*Delegation, error) {
	var delegation Delegation
	found, err := this.table.PrimaryKey().Get(ctx, &delegation, id)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &delegation, nil
}

func (this delegationTable) List(ctx context.Context, prefixKey DelegationIndexKey, opts ...ormlist.Option) (DelegationIterator, error) {
	it, err := this.table.GetIndexByID(prefixKey.id()).List(ctx, prefixKey.values(), opts...)
	return DelegationIterator{it}, err
}

func (this delegationTable) ListRange(ctx context.Context, from, to DelegationIndexKey, opts ...ormlist.Option) (DelegationIterator, error) {
	it, err := this.table.GetIndexByID(from.id()).ListRange(ctx, from.values(), to.values(), opts...)
	return DelegationIterator{it}, err
}

func (this delegationTable) DeleteBy(ctx context.Context, prefixKey DelegationIndexKey) error {
	return this.table.GetIndexByID(prefixKey.id()).DeleteBy(ctx, prefixKey.values()...)
}

func (this delegationTable) DeleteRange(ctx context.Context, from, to DelegationIndexKey) error {
	return this.table.GetIndexByID(from.id()).DeleteRange(ctx, from.values(), to.values())
}

func (this delegationTable) doNotImplement() {}

var _ DelegationTable = delegationTable{}

func NewDelegationTable(db ormtable.Schema) (DelegationTable, error) {
	table := db.GetTable(&Delegation{})
	if table == nil {
		return nil, ormerrors.TableNotFound.Wrap(string((&Delegation{}).ProtoReflect().Descriptor().FullName()))
	}
	return delegationTable{table}, nil
}

type ServiceTable interface {
	Insert(ctx context.Context, service *Service) error
	Update(ctx context.Context, service *Service) error
	Save(ctx context.Context, service *Service) error
	Delete(ctx context.Context, service *Service) error
	Has(ctx context.Context, id string) (found bool, err error)
	// Get returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	Get(ctx context.Context, id string) (*Service, error)
	List(ctx context.Context, prefixKey ServiceIndexKey, opts ...ormlist.Option) (ServiceIterator, error)
	ListRange(ctx context.Context, from, to ServiceIndexKey, opts ...ormlist.Option) (ServiceIterator, error)
	DeleteBy(ctx context.Context, prefixKey ServiceIndexKey) error
	DeleteRange(ctx context.Context, from, to ServiceIndexKey) error

	doNotImplement()
}

type ServiceIterator struct {
	ormtable.Iterator
}

func (i ServiceIterator) Value() (*Service, error) {
	var service Service
	err := i.UnmarshalMessage(&service)
	return &service, err
}

type ServiceIndexKey interface {
	id() uint32
	values() []interface{}
	serviceIndexKey()
}

// primary key starting index..
type ServicePrimaryKey = ServiceIdIndexKey

type ServiceIdIndexKey struct {
	vs []interface{}
}

func (x ServiceIdIndexKey) id() uint32            { return 0 }
func (x ServiceIdIndexKey) values() []interface{} { return x.vs }
func (x ServiceIdIndexKey) serviceIndexKey()      {}

func (this ServiceIdIndexKey) WithId(id string) ServiceIdIndexKey {
	this.vs = []interface{}{id}
	return this
}

type serviceTable struct {
	table ormtable.Table
}

func (this serviceTable) Insert(ctx context.Context, service *Service) error {
	return this.table.Insert(ctx, service)
}

func (this serviceTable) Update(ctx context.Context, service *Service) error {
	return this.table.Update(ctx, service)
}

func (this serviceTable) Save(ctx context.Context, service *Service) error {
	return this.table.Save(ctx, service)
}

func (this serviceTable) Delete(ctx context.Context, service *Service) error {
	return this.table.Delete(ctx, service)
}

func (this serviceTable) Has(ctx context.Context, id string) (found bool, err error) {
	return this.table.PrimaryKey().Has(ctx, id)
}

func (this serviceTable) Get(ctx context.Context, id string) (*Service, error) {
	var service Service
	found, err := this.table.PrimaryKey().Get(ctx, &service, id)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &service, nil
}

func (this serviceTable) List(ctx context.Context, prefixKey ServiceIndexKey, opts ...ormlist.Option) (ServiceIterator, error) {
	it, err := this.table.GetIndexByID(prefixKey.id()).List(ctx, prefixKey.values(), opts...)
	return ServiceIterator{it}, err
}

func (this serviceTable) ListRange(ctx context.Context, from, to ServiceIndexKey, opts ...ormlist.Option) (ServiceIterator, error) {
	it, err := this.table.GetIndexByID(from.id()).ListRange(ctx, from.values(), to.values(), opts...)
	return ServiceIterator{it}, err
}

func (this serviceTable) DeleteBy(ctx context.Context, prefixKey ServiceIndexKey) error {
	return this.table.GetIndexByID(prefixKey.id()).DeleteBy(ctx, prefixKey.values()...)
}

func (this serviceTable) DeleteRange(ctx context.Context, from, to ServiceIndexKey) error {
	return this.table.GetIndexByID(from.id()).DeleteRange(ctx, from.values(), to.values())
}

func (this serviceTable) doNotImplement() {}

var _ ServiceTable = serviceTable{}

func NewServiceTable(db ormtable.Schema) (ServiceTable, error) {
	table := db.GetTable(&Service{})
	if table == nil {
		return nil, ormerrors.TableNotFound.Wrap(string((&Service{}).ProtoReflect().Descriptor().FullName()))
	}
	return serviceTable{table}, nil
}

type StateStore interface {
	AliasesTable() AliasesTable
	AssertionTable() AssertionTable
	AttestationTable() AttestationTable
	ControllerTable() ControllerTable
	DelegationTable() DelegationTable
	ServiceTable() ServiceTable

	doNotImplement()
}

type stateStore struct {
	aliases     AliasesTable
	assertion   AssertionTable
	attestation AttestationTable
	controller  ControllerTable
	delegation  DelegationTable
	service     ServiceTable
}

func (x stateStore) AliasesTable() AliasesTable {
	return x.aliases
}

func (x stateStore) AssertionTable() AssertionTable {
	return x.assertion
}

func (x stateStore) AttestationTable() AttestationTable {
	return x.attestation
}

func (x stateStore) ControllerTable() ControllerTable {
	return x.controller
}

func (x stateStore) DelegationTable() DelegationTable {
	return x.delegation
}

func (x stateStore) ServiceTable() ServiceTable {
	return x.service
}

func (stateStore) doNotImplement() {}

var _ StateStore = stateStore{}

func NewStateStore(db ormtable.Schema) (StateStore, error) {
	aliasesTable, err := NewAliasesTable(db)
	if err != nil {
		return nil, err
	}

	assertionTable, err := NewAssertionTable(db)
	if err != nil {
		return nil, err
	}

	attestationTable, err := NewAttestationTable(db)
	if err != nil {
		return nil, err
	}

	controllerTable, err := NewControllerTable(db)
	if err != nil {
		return nil, err
	}

	delegationTable, err := NewDelegationTable(db)
	if err != nil {
		return nil, err
	}

	serviceTable, err := NewServiceTable(db)
	if err != nil {
		return nil, err
	}

	return stateStore{
		aliasesTable,
		assertionTable,
		attestationTable,
		controllerTable,
		delegationTable,
		serviceTable,
	}, nil
}
