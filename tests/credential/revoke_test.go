package credential

import (
	"testing"

	"github.com/ProtoconNet/mitum-contract-tests/tests/util"
	"github.com/ProtoconNet/mitum-currency/v3/operation/test"
	currencytypes "github.com/ProtoconNet/mitum-currency/v3/types"

	"github.com/ProtoconNet/mitum-credential/operation/credential"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

//Use below variables for default node configuration values
//t.NetworkID    	: network id
//t.GenesisPriv  	: genesis account private key
//t.GenesisAddr  	: genesis account address
//t.GenesisCurrency : genesis currency

type testRevoke struct {
	suite.Suite
	credential.TestRevokeProcessor
	account     []test.Account
	sender      []test.Account
	contract    []test.Account
	holder      []test.Account
	currency    []currencytypes.CurrencyID
	ownerKey    string // Private Key
	senderKey   string // Private Key
	contractKey string // Private Key
	holderKey   string // Private Key
	accountKey  string
	owner       []test.Account
}

func (t *testRevoke) SetupTest() {
	opr := credential.NewTestRevokeProcessor(util.Encoders)
	t.TestRevokeProcessor = opr
	t.Setup()
	t.account = make([]test.Account, 1)
	t.owner = make([]test.Account, 1)
	t.sender = make([]test.Account, 1)
	t.contract = make([]test.Account, 1)
	t.holder = make([]test.Account, 1)
	t.currency = make([]currencytypes.CurrencyID, 1)
	t.ownerKey = t.NewPrivateKey("owner")
	t.senderKey = t.NewPrivateKey("sender")
	t.contractKey = t.NewPrivateKey("contract")
	t.holderKey = t.NewPrivateKey("holder")
	t.accountKey = t.NewPrivateKey("account")
}

// func (t *testRevoke) Test01ErrorSenderNotFound() {
// 	err := t.Create().
// 		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, false).
// 		SetContractAccount(t.sender[0].Address(), t.contractKey, 1000, t.GenesisCurrency, t.contract, true).
// 		SetAccount(t.holderKey, 1000, t.GenesisCurrency, t.holder, true).
// 		SetTemplate(
// 			"templateID",
// 			"id",
// 		).
// 		MakeItem(t.contract[0], t.holder[0], t.GenesisCurrency, t.Items()).
// 		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(), t.Items()).
// 		RunPreProcess()

// 	if assert.NotNil(t.Suite.T(), err) {
// 		t.Suite.T().Log(err.Error())
// 	}
// }

// func (t *testRevoke) Test02SenderIsCA() {
// 	err := t.Create().
// 		SetAccount(t.ownerKey, 1000, t.GenesisCurrency, t.owner, true).
// 		SetContractAccount(t.owner[0].Address(), t.senderKey, 1000, t.GenesisCurrency, t.sender, true).
// 		SetContractAccount(t.sender[0].Address(), t.contractKey, 1000, t.GenesisCurrency, t.contract, true).
// 		SetAccount(t.holderKey, 1000, t.GenesisCurrency, t.holder, true).
// 		SetTemplate(
// 			"templateID",
// 			"id",
// 		).
// 		MakeItem(t.contract[0], t.holder[0], t.GenesisCurrency, t.Items()).
// 		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(), t.Items()).
// 		RunPreProcess()

// 	if assert.NotNil(t.Suite.T(), err) {
// 		t.Suite.T().Log(err.Error())
// 	}
// }

// func (t *testRevoke) Test03InvalidSigning() {
// 	err := t.Create().
// 		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, true).
// 		SetContractAccount(t.sender[0].Address(), t.contractKey, 1000, t.GenesisCurrency, t.contract, true).
// 		SetAccount(t.holderKey, 1000, t.GenesisCurrency, t.holder, true).
// 		SetTemplate(
// 			"templateID",
// 			"id",
// 		).
// 		MakeItem(t.contract[0], t.holder[0], t.GenesisCurrency, t.Items()).
// 		MakeOperation(t.sender[0].Address(), t.holder[0].Priv(), t.Items()).
// 		RunPreProcess()

// 	if assert.NotNil(t.Suite.T(), err) {
// 		t.Suite.T().Log(err.Error())
// 	}
// }

// func (t *testRevoke) Test04CurrencyNotFound() {
// 	err := t.Create().
// 		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, true).
// 		SetContractAccount(t.sender[0].Address(), t.contractKey, 1000, t.GenesisCurrency, t.contract, true).
// 		SetAccount(t.holderKey, 1000, t.GenesisCurrency, t.holder, true).
// 		SetCurrency("FOO", 1000, t.sender[0].Address(), t.currency, false).
// 		SetTemplate(
// 			"templateID",
// 			"id",
// 		).
// 		MakeItem(t.contract[0], t.holder[0], t.currency[0], t.Items()).
// 		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(), t.Items()).
// 		RunPreProcess()

// 	if assert.NotNil(t.Suite.T(), err) {
// 		t.Suite.T().Log(err.Error())
// 	}
// }

// func (t *testRevoke) Test05ContractNotFound() {
// 	err := t.Create().
// 		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, true).
// 		SetContractAccount(t.sender[0].Address(), t.contractKey, 1000, t.GenesisCurrency, t.contract, false).
// 		SetAccount(t.holderKey, 1000, t.GenesisCurrency, t.holder, true).
// 		SetTemplate(
// 			"templateID",
// 			"id",
// 		).
// 		MakeItem(t.contract[0], t.holder[0], t.GenesisCurrency, t.Items()).
// 		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(), t.Items()).
// 		RunPreProcess()

// 	if assert.NotNil(t.Suite.T(), err) {
// 		t.Suite.T().Log(err.Error())
// 	}
// }

// func (t *testRevoke) Test06SenderNotAuthorized() {
// 	err := t.Create().
// 		SetAccount(t.accountKey, 1000, t.GenesisCurrency, t.account, true).
// 		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, true).
// 		SetContractAccount(t.sender[0].Address(), t.contractKey, 1000, t.GenesisCurrency, t.contract, true).
// 		SetAccount(t.holderKey, 1000, t.GenesisCurrency, t.holder, true).
// 		SetTemplate(
// 			"templateID",
// 			"id",
// 		).
// 		MakeItem(t.contract[0], t.holder[0], t.GenesisCurrency, t.Items()).
// 		MakeOperation(t.account[0].Address(), t.account[0].Priv(), t.Items()).
// 		RunPreProcess()

// 	if assert.NotNil(t.Suite.T(), err) {
// 		t.Suite.T().Log(err.Error())
// 	}
// }

// func (t *testRevoke) Test07CredentialStateNotFound() {
// 	err := t.Create().
// 		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, true).
// 		SetContractAccount(t.sender[0].Address(), t.contractKey, 1000, t.GenesisCurrency, t.contract, true).
// 		SetAccount(t.holderKey, 1000, t.GenesisCurrency, t.holder, true).
// 		SetTemplate(
// 			"templateID",
// 			"id",
// 		).
// 		MakeItem(t.contract[0], t.holder[0], t.GenesisCurrency, t.Items()).
// 		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(), t.Items()).
// 		RunPreProcess()

// 	if assert.NotNil(t.Suite.T(), err) {
// 		t.Suite.T().Log(err.Error())
// 	}
// }

// func (t *testRevoke) Test08SpecialString1() {
// 	err := t.Create().
// 		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, true).
// 		SetContractAccount(t.sender[0].Address(), t.contractKey, 1000, t.GenesisCurrency, t.contract, true).
// 		SetAccount(t.holderKey, 1000, t.GenesisCurrency, t.holder, true).
// 		SetTemplate(
// 			"!@#$%^&*(?/)",
// 			"id",
// 		).
// 		MakeItem(t.contract[0], t.holder[0], t.GenesisCurrency, t.Items()).
// 		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(), t.Items()).
// 		RunPreProcess()

// 	if assert.NotNil(t.Suite.T(), err) {
// 		t.Suite.T().Log(err.Error())
// 	}
// }

// func (t *testRevoke) Test09SpecialString2() {
// 	err := t.Create().
// 		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, true).
// 		SetContractAccount(t.sender[0].Address(), t.contractKey, 1000, t.GenesisCurrency, t.contract, true).
// 		SetAccount(t.holderKey, 1000, t.GenesisCurrency, t.holder, true).
// 		SetTemplate(
// 			"templateID",
// 			"!@#$%^&*(?/)",
// 		).
// 		MakeItem(t.contract[0], t.holder[0], t.GenesisCurrency, t.Items()).
// 		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(), t.Items()).
// 		RunPreProcess()

// 	if assert.NotNil(t.Suite.T(), err) {
// 		t.Suite.T().Log(err.Error())
// 	}
// }

func (t *testRevoke) Test10ErrorSenderNotFound() {
	err := t.Create().
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, true).
		SetContractAccount(t.sender[0].Address(), t.contractKey, 1000, t.GenesisCurrency, t.contract, true).
		SetAccount(t.holderKey, 1000, t.GenesisCurrency, t.holder, false).
		SetTemplate(
			"templateID",
			"id",
		).
		MakeItem(t.contract[0], t.holder[0], t.GenesisCurrency, t.Items()).
		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(), t.Items()).
		RunPreProcess()

	if assert.NotNil(t.Suite.T(), err) {
		t.Suite.T().Log(err.Error())
	}
}

// func (t *testRevoke) Test11SelfTargetedHolder() {
// 	err := t.Create().
// 		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, true).
// 		SetContractAccount(t.sender[0].Address(), t.contractKey, 1000, t.GenesisCurrency, t.contract, true).
// 		SetTemplate(
// 			"templateID",
// 			"id",
// 		).
// 		MakeItem(t.contract[0], t.contract[0], t.GenesisCurrency, t.Items()).
// 		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(), t.Items()).
// 		IsValid()

// 	if assert.NotNil(t.Suite.T(), err) {
// 		t.Suite.T().Log(err.Error())
// 	}
// }

// func (t *testRevoke) Test12SelfTargetedSender() {
// 	err := t.Create().
// 		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, true).
// 		SetContractAccount(t.sender[0].Address(), t.contractKey, 1000, t.GenesisCurrency, t.contract, true).
// 		SetAccount(t.holderKey, 1000, t.GenesisCurrency, t.holder, true).
// 		SetTemplate(
// 			"templateID",
// 			"id",
// 		).
// 		MakeItem(t.sender[0], t.holder[0], t.GenesisCurrency, t.Items()).
// 		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(), t.Items()).
// 		IsValid()

// 	if assert.NotNil(t.Suite.T(), err) {
// 		t.Suite.T().Log(err.Error())
// 	}
// }

func TestRevoke(t *testing.T) {
	suite.Run(t, new(testRevoke))
}
