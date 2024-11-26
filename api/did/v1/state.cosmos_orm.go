// Code generated by protoc-gen-go-cosmos-orm. DO NOT EDIT.

package didv1

import (
	context "context"
	ormlist "cosmossdk.io/orm/model/ormlist"
	ormtable "cosmossdk.io/orm/model/ormtable"
	ormerrors "cosmossdk.io/orm/types/ormerrors"
)

type AccountTable interface {
	Insert(ctx context.Context, account *Account) error
	Update(ctx context.Context, account *Account) error
	Save(ctx context.Context, account *Account) error
	Delete(ctx context.Context, account *Account) error
	Has(ctx context.Context, did string) (found bool, err error)
	// Get returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	Get(ctx context.Context, did string) (*Account, error)
	HasByControllerSubject(ctx context.Context, controller string, subject string) (found bool, err error)
	// GetByControllerSubject returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	GetByControllerSubject(ctx context.Context, controller string, subject string) (*Account, error)
	List(ctx context.Context, prefixKey AccountIndexKey, opts ...ormlist.Option) (AccountIterator, error)
	ListRange(ctx context.Context, from, to AccountIndexKey, opts ...ormlist.Option) (AccountIterator, error)
	DeleteBy(ctx context.Context, prefixKey AccountIndexKey) error
	DeleteRange(ctx context.Context, from, to AccountIndexKey) error

	doNotImplement()
}

type AccountIterator struct {
	ormtable.Iterator
}

func (i AccountIterator) Value() (*Account, error) {
	var account Account
	err := i.UnmarshalMessage(&account)
	return &account, err
}

type AccountIndexKey interface {
	id() uint32
	values() []interface{}
	accountIndexKey()
}

// primary key starting index..
type AccountPrimaryKey = AccountDidIndexKey

type AccountDidIndexKey struct {
	vs []interface{}
}

func (x AccountDidIndexKey) id() uint32            { return 0 }
func (x AccountDidIndexKey) values() []interface{} { return x.vs }
func (x AccountDidIndexKey) accountIndexKey()      {}

func (this AccountDidIndexKey) WithDid(did string) AccountDidIndexKey {
	this.vs = []interface{}{did}
	return this
}

type AccountControllerSubjectIndexKey struct {
	vs []interface{}
}

func (x AccountControllerSubjectIndexKey) id() uint32            { return 1 }
func (x AccountControllerSubjectIndexKey) values() []interface{} { return x.vs }
func (x AccountControllerSubjectIndexKey) accountIndexKey()      {}

func (this AccountControllerSubjectIndexKey) WithController(controller string) AccountControllerSubjectIndexKey {
	this.vs = []interface{}{controller}
	return this
}

func (this AccountControllerSubjectIndexKey) WithControllerSubject(controller string, subject string) AccountControllerSubjectIndexKey {
	this.vs = []interface{}{controller, subject}
	return this
}

type accountTable struct {
	table ormtable.Table
}

func (this accountTable) Insert(ctx context.Context, account *Account) error {
	return this.table.Insert(ctx, account)
}

func (this accountTable) Update(ctx context.Context, account *Account) error {
	return this.table.Update(ctx, account)
}

func (this accountTable) Save(ctx context.Context, account *Account) error {
	return this.table.Save(ctx, account)
}

func (this accountTable) Delete(ctx context.Context, account *Account) error {
	return this.table.Delete(ctx, account)
}

func (this accountTable) Has(ctx context.Context, did string) (found bool, err error) {
	return this.table.PrimaryKey().Has(ctx, did)
}

func (this accountTable) Get(ctx context.Context, did string) (*Account, error) {
	var account Account
	found, err := this.table.PrimaryKey().Get(ctx, &account, did)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &account, nil
}

func (this accountTable) HasByControllerSubject(ctx context.Context, controller string, subject string) (found bool, err error) {
	return this.table.GetIndexByID(1).(ormtable.UniqueIndex).Has(ctx,
		controller,
		subject,
	)
}

func (this accountTable) GetByControllerSubject(ctx context.Context, controller string, subject string) (*Account, error) {
	var account Account
	found, err := this.table.GetIndexByID(1).(ormtable.UniqueIndex).Get(ctx, &account,
		controller,
		subject,
	)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &account, nil
}

func (this accountTable) List(ctx context.Context, prefixKey AccountIndexKey, opts ...ormlist.Option) (AccountIterator, error) {
	it, err := this.table.GetIndexByID(prefixKey.id()).List(ctx, prefixKey.values(), opts...)
	return AccountIterator{it}, err
}

func (this accountTable) ListRange(ctx context.Context, from, to AccountIndexKey, opts ...ormlist.Option) (AccountIterator, error) {
	it, err := this.table.GetIndexByID(from.id()).ListRange(ctx, from.values(), to.values(), opts...)
	return AccountIterator{it}, err
}

func (this accountTable) DeleteBy(ctx context.Context, prefixKey AccountIndexKey) error {
	return this.table.GetIndexByID(prefixKey.id()).DeleteBy(ctx, prefixKey.values()...)
}

func (this accountTable) DeleteRange(ctx context.Context, from, to AccountIndexKey) error {
	return this.table.GetIndexByID(from.id()).DeleteRange(ctx, from.values(), to.values())
}

func (this accountTable) doNotImplement() {}

var _ AccountTable = accountTable{}

func NewAccountTable(db ormtable.Schema) (AccountTable, error) {
	table := db.GetTable(&Account{})
	if table == nil {
		return nil, ormerrors.TableNotFound.Wrap(string((&Account{}).ProtoReflect().Descriptor().FullName()))
	}
	return accountTable{table}, nil
}

type PublicKeyTable interface {
	Insert(ctx context.Context, publicKey *PublicKey) error
	InsertReturningNumber(ctx context.Context, publicKey *PublicKey) (uint64, error)
	LastInsertedSequence(ctx context.Context) (uint64, error)
	Update(ctx context.Context, publicKey *PublicKey) error
	Save(ctx context.Context, publicKey *PublicKey) error
	Delete(ctx context.Context, publicKey *PublicKey) error
	Has(ctx context.Context, number uint64) (found bool, err error)
	// Get returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	Get(ctx context.Context, number uint64) (*PublicKey, error)
	HasBySonrAddress(ctx context.Context, sonr_address string) (found bool, err error)
	// GetBySonrAddress returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	GetBySonrAddress(ctx context.Context, sonr_address string) (*PublicKey, error)
	HasByEthAddress(ctx context.Context, eth_address string) (found bool, err error)
	// GetByEthAddress returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	GetByEthAddress(ctx context.Context, eth_address string) (*PublicKey, error)
	HasByBtcAddress(ctx context.Context, btc_address string) (found bool, err error)
	// GetByBtcAddress returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	GetByBtcAddress(ctx context.Context, btc_address string) (*PublicKey, error)
	HasByDid(ctx context.Context, did string) (found bool, err error)
	// GetByDid returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	GetByDid(ctx context.Context, did string) (*PublicKey, error)
	List(ctx context.Context, prefixKey PublicKeyIndexKey, opts ...ormlist.Option) (PublicKeyIterator, error)
	ListRange(ctx context.Context, from, to PublicKeyIndexKey, opts ...ormlist.Option) (PublicKeyIterator, error)
	DeleteBy(ctx context.Context, prefixKey PublicKeyIndexKey) error
	DeleteRange(ctx context.Context, from, to PublicKeyIndexKey) error

	doNotImplement()
}

type PublicKeyIterator struct {
	ormtable.Iterator
}

func (i PublicKeyIterator) Value() (*PublicKey, error) {
	var publicKey PublicKey
	err := i.UnmarshalMessage(&publicKey)
	return &publicKey, err
}

type PublicKeyIndexKey interface {
	id() uint32
	values() []interface{}
	publicKeyIndexKey()
}

// primary key starting index..
type PublicKeyPrimaryKey = PublicKeyNumberIndexKey

type PublicKeyNumberIndexKey struct {
	vs []interface{}
}

func (x PublicKeyNumberIndexKey) id() uint32            { return 0 }
func (x PublicKeyNumberIndexKey) values() []interface{} { return x.vs }
func (x PublicKeyNumberIndexKey) publicKeyIndexKey()    {}

func (this PublicKeyNumberIndexKey) WithNumber(number uint64) PublicKeyNumberIndexKey {
	this.vs = []interface{}{number}
	return this
}

type PublicKeySonrAddressIndexKey struct {
	vs []interface{}
}

func (x PublicKeySonrAddressIndexKey) id() uint32            { return 1 }
func (x PublicKeySonrAddressIndexKey) values() []interface{} { return x.vs }
func (x PublicKeySonrAddressIndexKey) publicKeyIndexKey()    {}

func (this PublicKeySonrAddressIndexKey) WithSonrAddress(sonr_address string) PublicKeySonrAddressIndexKey {
	this.vs = []interface{}{sonr_address}
	return this
}

type PublicKeyEthAddressIndexKey struct {
	vs []interface{}
}

func (x PublicKeyEthAddressIndexKey) id() uint32            { return 2 }
func (x PublicKeyEthAddressIndexKey) values() []interface{} { return x.vs }
func (x PublicKeyEthAddressIndexKey) publicKeyIndexKey()    {}

func (this PublicKeyEthAddressIndexKey) WithEthAddress(eth_address string) PublicKeyEthAddressIndexKey {
	this.vs = []interface{}{eth_address}
	return this
}

type PublicKeyBtcAddressIndexKey struct {
	vs []interface{}
}

func (x PublicKeyBtcAddressIndexKey) id() uint32            { return 3 }
func (x PublicKeyBtcAddressIndexKey) values() []interface{} { return x.vs }
func (x PublicKeyBtcAddressIndexKey) publicKeyIndexKey()    {}

func (this PublicKeyBtcAddressIndexKey) WithBtcAddress(btc_address string) PublicKeyBtcAddressIndexKey {
	this.vs = []interface{}{btc_address}
	return this
}

type PublicKeyDidIndexKey struct {
	vs []interface{}
}

func (x PublicKeyDidIndexKey) id() uint32            { return 4 }
func (x PublicKeyDidIndexKey) values() []interface{} { return x.vs }
func (x PublicKeyDidIndexKey) publicKeyIndexKey()    {}

func (this PublicKeyDidIndexKey) WithDid(did string) PublicKeyDidIndexKey {
	this.vs = []interface{}{did}
	return this
}

type publicKeyTable struct {
	table ormtable.AutoIncrementTable
}

func (this publicKeyTable) Insert(ctx context.Context, publicKey *PublicKey) error {
	return this.table.Insert(ctx, publicKey)
}

func (this publicKeyTable) Update(ctx context.Context, publicKey *PublicKey) error {
	return this.table.Update(ctx, publicKey)
}

func (this publicKeyTable) Save(ctx context.Context, publicKey *PublicKey) error {
	return this.table.Save(ctx, publicKey)
}

func (this publicKeyTable) Delete(ctx context.Context, publicKey *PublicKey) error {
	return this.table.Delete(ctx, publicKey)
}

func (this publicKeyTable) InsertReturningNumber(ctx context.Context, publicKey *PublicKey) (uint64, error) {
	return this.table.InsertReturningPKey(ctx, publicKey)
}

func (this publicKeyTable) LastInsertedSequence(ctx context.Context) (uint64, error) {
	return this.table.LastInsertedSequence(ctx)
}

func (this publicKeyTable) Has(ctx context.Context, number uint64) (found bool, err error) {
	return this.table.PrimaryKey().Has(ctx, number)
}

func (this publicKeyTable) Get(ctx context.Context, number uint64) (*PublicKey, error) {
	var publicKey PublicKey
	found, err := this.table.PrimaryKey().Get(ctx, &publicKey, number)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &publicKey, nil
}

func (this publicKeyTable) HasBySonrAddress(ctx context.Context, sonr_address string) (found bool, err error) {
	return this.table.GetIndexByID(1).(ormtable.UniqueIndex).Has(ctx,
		sonr_address,
	)
}

func (this publicKeyTable) GetBySonrAddress(ctx context.Context, sonr_address string) (*PublicKey, error) {
	var publicKey PublicKey
	found, err := this.table.GetIndexByID(1).(ormtable.UniqueIndex).Get(ctx, &publicKey,
		sonr_address,
	)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &publicKey, nil
}

func (this publicKeyTable) HasByEthAddress(ctx context.Context, eth_address string) (found bool, err error) {
	return this.table.GetIndexByID(2).(ormtable.UniqueIndex).Has(ctx,
		eth_address,
	)
}

func (this publicKeyTable) GetByEthAddress(ctx context.Context, eth_address string) (*PublicKey, error) {
	var publicKey PublicKey
	found, err := this.table.GetIndexByID(2).(ormtable.UniqueIndex).Get(ctx, &publicKey,
		eth_address,
	)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &publicKey, nil
}

func (this publicKeyTable) HasByBtcAddress(ctx context.Context, btc_address string) (found bool, err error) {
	return this.table.GetIndexByID(3).(ormtable.UniqueIndex).Has(ctx,
		btc_address,
	)
}

func (this publicKeyTable) GetByBtcAddress(ctx context.Context, btc_address string) (*PublicKey, error) {
	var publicKey PublicKey
	found, err := this.table.GetIndexByID(3).(ormtable.UniqueIndex).Get(ctx, &publicKey,
		btc_address,
	)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &publicKey, nil
}

func (this publicKeyTable) HasByDid(ctx context.Context, did string) (found bool, err error) {
	return this.table.GetIndexByID(4).(ormtable.UniqueIndex).Has(ctx,
		did,
	)
}

func (this publicKeyTable) GetByDid(ctx context.Context, did string) (*PublicKey, error) {
	var publicKey PublicKey
	found, err := this.table.GetIndexByID(4).(ormtable.UniqueIndex).Get(ctx, &publicKey,
		did,
	)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &publicKey, nil
}

func (this publicKeyTable) List(ctx context.Context, prefixKey PublicKeyIndexKey, opts ...ormlist.Option) (PublicKeyIterator, error) {
	it, err := this.table.GetIndexByID(prefixKey.id()).List(ctx, prefixKey.values(), opts...)
	return PublicKeyIterator{it}, err
}

func (this publicKeyTable) ListRange(ctx context.Context, from, to PublicKeyIndexKey, opts ...ormlist.Option) (PublicKeyIterator, error) {
	it, err := this.table.GetIndexByID(from.id()).ListRange(ctx, from.values(), to.values(), opts...)
	return PublicKeyIterator{it}, err
}

func (this publicKeyTable) DeleteBy(ctx context.Context, prefixKey PublicKeyIndexKey) error {
	return this.table.GetIndexByID(prefixKey.id()).DeleteBy(ctx, prefixKey.values()...)
}

func (this publicKeyTable) DeleteRange(ctx context.Context, from, to PublicKeyIndexKey) error {
	return this.table.GetIndexByID(from.id()).DeleteRange(ctx, from.values(), to.values())
}

func (this publicKeyTable) doNotImplement() {}

var _ PublicKeyTable = publicKeyTable{}

func NewPublicKeyTable(db ormtable.Schema) (PublicKeyTable, error) {
	table := db.GetTable(&PublicKey{})
	if table == nil {
		return nil, ormerrors.TableNotFound.Wrap(string((&PublicKey{}).ProtoReflect().Descriptor().FullName()))
	}
	return publicKeyTable{table.(ormtable.AutoIncrementTable)}, nil
}

type VerificationTable interface {
	Insert(ctx context.Context, verification *Verification) error
	Update(ctx context.Context, verification *Verification) error
	Save(ctx context.Context, verification *Verification) error
	Delete(ctx context.Context, verification *Verification) error
	Has(ctx context.Context, did string) (found bool, err error)
	// Get returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	Get(ctx context.Context, did string) (*Verification, error)
	HasByIssuerSubject(ctx context.Context, issuer string, subject string) (found bool, err error)
	// GetByIssuerSubject returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	GetByIssuerSubject(ctx context.Context, issuer string, subject string) (*Verification, error)
	HasByControllerDidMethodIssuer(ctx context.Context, controller string, did_method string, issuer string) (found bool, err error)
	// GetByControllerDidMethodIssuer returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	GetByControllerDidMethodIssuer(ctx context.Context, controller string, did_method string, issuer string) (*Verification, error)
	HasByVerificationTypeSubjectIssuer(ctx context.Context, verification_type string, subject string, issuer string) (found bool, err error)
	// GetByVerificationTypeSubjectIssuer returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	GetByVerificationTypeSubjectIssuer(ctx context.Context, verification_type string, subject string, issuer string) (*Verification, error)
	List(ctx context.Context, prefixKey VerificationIndexKey, opts ...ormlist.Option) (VerificationIterator, error)
	ListRange(ctx context.Context, from, to VerificationIndexKey, opts ...ormlist.Option) (VerificationIterator, error)
	DeleteBy(ctx context.Context, prefixKey VerificationIndexKey) error
	DeleteRange(ctx context.Context, from, to VerificationIndexKey) error

	doNotImplement()
}

type VerificationIterator struct {
	ormtable.Iterator
}

func (i VerificationIterator) Value() (*Verification, error) {
	var verification Verification
	err := i.UnmarshalMessage(&verification)
	return &verification, err
}

type VerificationIndexKey interface {
	id() uint32
	values() []interface{}
	verificationIndexKey()
}

// primary key starting index..
type VerificationPrimaryKey = VerificationDidIndexKey

type VerificationDidIndexKey struct {
	vs []interface{}
}

func (x VerificationDidIndexKey) id() uint32            { return 0 }
func (x VerificationDidIndexKey) values() []interface{} { return x.vs }
func (x VerificationDidIndexKey) verificationIndexKey() {}

func (this VerificationDidIndexKey) WithDid(did string) VerificationDidIndexKey {
	this.vs = []interface{}{did}
	return this
}

type VerificationIssuerSubjectIndexKey struct {
	vs []interface{}
}

func (x VerificationIssuerSubjectIndexKey) id() uint32            { return 1 }
func (x VerificationIssuerSubjectIndexKey) values() []interface{} { return x.vs }
func (x VerificationIssuerSubjectIndexKey) verificationIndexKey() {}

func (this VerificationIssuerSubjectIndexKey) WithIssuer(issuer string) VerificationIssuerSubjectIndexKey {
	this.vs = []interface{}{issuer}
	return this
}

func (this VerificationIssuerSubjectIndexKey) WithIssuerSubject(issuer string, subject string) VerificationIssuerSubjectIndexKey {
	this.vs = []interface{}{issuer, subject}
	return this
}

type VerificationControllerDidMethodIssuerIndexKey struct {
	vs []interface{}
}

func (x VerificationControllerDidMethodIssuerIndexKey) id() uint32            { return 2 }
func (x VerificationControllerDidMethodIssuerIndexKey) values() []interface{} { return x.vs }
func (x VerificationControllerDidMethodIssuerIndexKey) verificationIndexKey() {}

func (this VerificationControllerDidMethodIssuerIndexKey) WithController(controller string) VerificationControllerDidMethodIssuerIndexKey {
	this.vs = []interface{}{controller}
	return this
}

func (this VerificationControllerDidMethodIssuerIndexKey) WithControllerDidMethod(controller string, did_method string) VerificationControllerDidMethodIssuerIndexKey {
	this.vs = []interface{}{controller, did_method}
	return this
}

func (this VerificationControllerDidMethodIssuerIndexKey) WithControllerDidMethodIssuer(controller string, did_method string, issuer string) VerificationControllerDidMethodIssuerIndexKey {
	this.vs = []interface{}{controller, did_method, issuer}
	return this
}

type VerificationVerificationTypeSubjectIssuerIndexKey struct {
	vs []interface{}
}

func (x VerificationVerificationTypeSubjectIssuerIndexKey) id() uint32            { return 3 }
func (x VerificationVerificationTypeSubjectIssuerIndexKey) values() []interface{} { return x.vs }
func (x VerificationVerificationTypeSubjectIssuerIndexKey) verificationIndexKey() {}

func (this VerificationVerificationTypeSubjectIssuerIndexKey) WithVerificationType(verification_type string) VerificationVerificationTypeSubjectIssuerIndexKey {
	this.vs = []interface{}{verification_type}
	return this
}

func (this VerificationVerificationTypeSubjectIssuerIndexKey) WithVerificationTypeSubject(verification_type string, subject string) VerificationVerificationTypeSubjectIssuerIndexKey {
	this.vs = []interface{}{verification_type, subject}
	return this
}

func (this VerificationVerificationTypeSubjectIssuerIndexKey) WithVerificationTypeSubjectIssuer(verification_type string, subject string, issuer string) VerificationVerificationTypeSubjectIssuerIndexKey {
	this.vs = []interface{}{verification_type, subject, issuer}
	return this
}

type verificationTable struct {
	table ormtable.Table
}

func (this verificationTable) Insert(ctx context.Context, verification *Verification) error {
	return this.table.Insert(ctx, verification)
}

func (this verificationTable) Update(ctx context.Context, verification *Verification) error {
	return this.table.Update(ctx, verification)
}

func (this verificationTable) Save(ctx context.Context, verification *Verification) error {
	return this.table.Save(ctx, verification)
}

func (this verificationTable) Delete(ctx context.Context, verification *Verification) error {
	return this.table.Delete(ctx, verification)
}

func (this verificationTable) Has(ctx context.Context, did string) (found bool, err error) {
	return this.table.PrimaryKey().Has(ctx, did)
}

func (this verificationTable) Get(ctx context.Context, did string) (*Verification, error) {
	var verification Verification
	found, err := this.table.PrimaryKey().Get(ctx, &verification, did)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &verification, nil
}

func (this verificationTable) HasByIssuerSubject(ctx context.Context, issuer string, subject string) (found bool, err error) {
	return this.table.GetIndexByID(1).(ormtable.UniqueIndex).Has(ctx,
		issuer,
		subject,
	)
}

func (this verificationTable) GetByIssuerSubject(ctx context.Context, issuer string, subject string) (*Verification, error) {
	var verification Verification
	found, err := this.table.GetIndexByID(1).(ormtable.UniqueIndex).Get(ctx, &verification,
		issuer,
		subject,
	)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &verification, nil
}

func (this verificationTable) HasByControllerDidMethodIssuer(ctx context.Context, controller string, did_method string, issuer string) (found bool, err error) {
	return this.table.GetIndexByID(2).(ormtable.UniqueIndex).Has(ctx,
		controller,
		did_method,
		issuer,
	)
}

func (this verificationTable) GetByControllerDidMethodIssuer(ctx context.Context, controller string, did_method string, issuer string) (*Verification, error) {
	var verification Verification
	found, err := this.table.GetIndexByID(2).(ormtable.UniqueIndex).Get(ctx, &verification,
		controller,
		did_method,
		issuer,
	)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &verification, nil
}

func (this verificationTable) HasByVerificationTypeSubjectIssuer(ctx context.Context, verification_type string, subject string, issuer string) (found bool, err error) {
	return this.table.GetIndexByID(3).(ormtable.UniqueIndex).Has(ctx,
		verification_type,
		subject,
		issuer,
	)
}

func (this verificationTable) GetByVerificationTypeSubjectIssuer(ctx context.Context, verification_type string, subject string, issuer string) (*Verification, error) {
	var verification Verification
	found, err := this.table.GetIndexByID(3).(ormtable.UniqueIndex).Get(ctx, &verification,
		verification_type,
		subject,
		issuer,
	)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &verification, nil
}

func (this verificationTable) List(ctx context.Context, prefixKey VerificationIndexKey, opts ...ormlist.Option) (VerificationIterator, error) {
	it, err := this.table.GetIndexByID(prefixKey.id()).List(ctx, prefixKey.values(), opts...)
	return VerificationIterator{it}, err
}

func (this verificationTable) ListRange(ctx context.Context, from, to VerificationIndexKey, opts ...ormlist.Option) (VerificationIterator, error) {
	it, err := this.table.GetIndexByID(from.id()).ListRange(ctx, from.values(), to.values(), opts...)
	return VerificationIterator{it}, err
}

func (this verificationTable) DeleteBy(ctx context.Context, prefixKey VerificationIndexKey) error {
	return this.table.GetIndexByID(prefixKey.id()).DeleteBy(ctx, prefixKey.values()...)
}

func (this verificationTable) DeleteRange(ctx context.Context, from, to VerificationIndexKey) error {
	return this.table.GetIndexByID(from.id()).DeleteRange(ctx, from.values(), to.values())
}

func (this verificationTable) doNotImplement() {}

var _ VerificationTable = verificationTable{}

func NewVerificationTable(db ormtable.Schema) (VerificationTable, error) {
	table := db.GetTable(&Verification{})
	if table == nil {
		return nil, ormerrors.TableNotFound.Wrap(string((&Verification{}).ProtoReflect().Descriptor().FullName()))
	}
	return verificationTable{table}, nil
}

type StateStore interface {
	AccountTable() AccountTable
	PublicKeyTable() PublicKeyTable
	VerificationTable() VerificationTable

	doNotImplement()
}

type stateStore struct {
	account      AccountTable
	publicKey    PublicKeyTable
	verification VerificationTable
}

func (x stateStore) AccountTable() AccountTable {
	return x.account
}

func (x stateStore) PublicKeyTable() PublicKeyTable {
	return x.publicKey
}

func (x stateStore) VerificationTable() VerificationTable {
	return x.verification
}

func (stateStore) doNotImplement() {}

var _ StateStore = stateStore{}

func NewStateStore(db ormtable.Schema) (StateStore, error) {
	accountTable, err := NewAccountTable(db)
	if err != nil {
		return nil, err
	}

	publicKeyTable, err := NewPublicKeyTable(db)
	if err != nil {
		return nil, err
	}

	verificationTable, err := NewVerificationTable(db)
	if err != nil {
		return nil, err
	}

	return stateStore{
		accountTable,
		publicKeyTable,
		verificationTable,
	}, nil
}
