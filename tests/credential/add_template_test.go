package credential

import (
	"testing"

	"github.com/ProtoconNet/mitum-contract-tests/tests/util"
	"github.com/ProtoconNet/mitum-currency/v3/operation/test"
	currencytypes "github.com/ProtoconNet/mitum-currency/v3/types"

	"github.com/ProtoconNet/mitum-credential/operation/credential"
	"github.com/ProtoconNet/mitum-credential/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

//Use below variables for default node configuration values
//t.NetworkID    	: network id
//t.GenesisPriv  	: genesis account private key
//t.GenesisAddr  	: genesis account address
//t.GenesisCurrency : genesis currency

type testAddTemplate struct {
	suite.Suite
	credential.TestAddTemplateProcessor
	account     []test.Account
	sender      []test.Account
	contract    []test.Account
	creator     []test.Account
	currency    []currencytypes.CurrencyID
	accountKey  string // Private Key
	ownerKey    string // Private Key
	senderKey   string // Private Key
	contractKey string // Private Key
	creatorKey  string // Private Key
	owner       []test.Account
}

func (t *testAddTemplate) SetupTest() {
	opr := credential.NewTestAddTemplateProcessor(util.Encoders)
	t.TestAddTemplateProcessor = opr
	t.Setup()
	t.account = make([]test.Account, 1)
	t.owner = make([]test.Account, 1)
	t.sender = make([]test.Account, 1)
	t.contract = make([]test.Account, 1)
	t.creator = make([]test.Account, 1)
	t.currency = make([]currencytypes.CurrencyID, 1)
	t.accountKey = t.NewPrivateKey("account")
	t.ownerKey = t.NewPrivateKey("owner")
	t.senderKey = t.NewPrivateKey("sender")
	t.contractKey = t.NewPrivateKey("contract")
	t.creatorKey = t.NewPrivateKey("creator")
}

// func (t *testAddTemplate) Test01ErrorSenderNotFound() {
// 	err := t.Create().
// 		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, false).
// 		SetContractAccount(t.sender[0].Address(), t.contractKey, 1000, t.GenesisCurrency, t.contract, true).
// 		SetAccount(t.creatorKey, 1000, t.GenesisCurrency, t.creator, true).
// 		SetTemplate(
// 			"templateID",
// 			"templateName",
// 			types.Date("2024-01-01"),
// 			types.Date("2024-01-01"),
// 			types.Bool(true),
// 			types.Bool(true),
// 			"displayName",
// 			"subjectKey",
// 			"description",
// 		).
// 		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(), t.contract[0].Address(), t.creator[0].Address(), t.GenesisCurrency).
// 		RunPreProcess()

// 	if assert.NotNil(t.Suite.T(), err) {
// 		t.Suite.T().Log(err.Error())
// 	}
// }

// func (t *testAddTemplate) Test02ErrorSenderIscontract() {
// 	err := t.Create().
// 		SetAccount(t.ownerKey, 1000, t.GenesisCurrency, t.owner, true).
// 		SetContractAccount(t.owner[0].Address(), t.senderKey, 1000, t.GenesisCurrency, t.sender, true).
// 		SetContractAccount(t.owner[0].Address(), t.contractKey, 1000, t.GenesisCurrency, t.contract, true).
// 		SetAccount(t.creatorKey, 1000, t.GenesisCurrency, t.creator, true).
// 		SetTemplate(
// 			"templateID",
// 			"templateName",
// 			types.Date("2024-01-01"),
// 			types.Date("2024-01-01"),
// 			types.Bool(true),
// 			types.Bool(true),
// 			"displayName",
// 			"subjectKey",
// 			"description",
// 		).
// 		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(), t.contract[0].Address(), t.creator[0].Address(), t.GenesisCurrency).
// 		RunPreProcess()

// 	if assert.NotNil(t.Suite.T(), err) {
// 		t.Suite.T().Log(err.Error())
// 	}
// }

// func (t *testAddTemplate) Test03CreatorNotFound() {
// 	err := t.Create().
// 		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, true).
// 		SetContractAccount(t.sender[0].Address(), t.contractKey, 1000, t.GenesisCurrency, t.contract, true).
// 		SetAccount(t.creatorKey, 1000, t.GenesisCurrency, t.creator, false).
// 		SetTemplate(
// 			"templateID",
// 			"templateName",
// 			types.Date("2024-01-01"),
// 			types.Date("2024-01-01"),
// 			types.Bool(true),
// 			types.Bool(true),
// 			"displayName",
// 			"subjectKey",
// 			"description",
// 		).
// 		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(), t.contract[0].Address(), t.creator[0].Address(), t.GenesisCurrency).
// 		RunPreProcess()

// 	if assert.NotNil(t.Suite.T(), err) {
// 		t.Suite.T().Log(err.Error())
// 	}
// }

// func (t *testAddTemplate) Test04CreatorIscontract() {
// 	err := t.Create().
// 		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, true).
// 		SetContractAccount(t.sender[0].Address(), t.contractKey, 1000, t.GenesisCurrency, t.contract, true).
// 		SetContractAccount(t.sender[0].Address(), t.creatorKey, 1000, t.GenesisCurrency, t.creator, true).
// 		SetTemplate(
// 			"templateID",
// 			"templateName",
// 			types.Date("2024-01-01"),
// 			types.Date("2024-01-01"),
// 			types.Bool(true),
// 			types.Bool(true),
// 			"displayName",
// 			"subjectKey",
// 			"description",
// 		).
// 		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(), t.contract[0].Address(), t.creator[0].Address(), t.GenesisCurrency).
// 		RunPreProcess()

// 	if assert.NotNil(t.Suite.T(), err) {
// 		t.Suite.T().Log(err.Error())
// 	}
// }

// // !! The test below does not work properly because the process first checks the credential service and then checks for invalid signatures.
// // func (t *testAddTemplate) Test05InvalidSigning() {
// // 	err := t.Create().
// // 		SetAccount(t.accountKey, 1000, t.GenesisCurrency, t.account, true).
// // 		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, true).
// // 		SetContractAccount(t.sender[0].Address(), t.contractKey, 1000, t.GenesisCurrency, t.contract, true).
// // 		SetAccount(t.creatorKey, 1000, t.GenesisCurrency, t.creator, true).
// // 		SetTemplate(
// // 			"templateID",
// // 			"templateName",
// // 			types.Date("2024-01-01"),
// // 			types.Date("2024-01-01"),
// // 			types.Bool(true),
// // 			types.Bool(true),
// // 			"displayName",
// // 			"subjectKey",
// // 			"description",
// // 		).
// // 		MakeOperation(t.sender[0].Address(), t.account[0].Priv(), t.contract[0].Address(), t.creator[0].Address(), t.GenesisCurrency).
// // 		RunPreProcess()

// // 	if assert.NotNil(t.Suite.T(), err) {
// // 		t.Suite.T().Log(err.Error())
// // 	}
// // }

// func (t *testAddTemplate) Test06CurrencyNotFound() {
// 	err := t.Create().
// 		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, true).
// 		SetContractAccount(t.sender[0].Address(), t.contractKey, 1000, t.GenesisCurrency, t.contract, true).
// 		SetAccount(t.creatorKey, 1000, t.GenesisCurrency, t.creator, true).
// 		SetTemplate(
// 			"templateID",
// 			"templateName",
// 			types.Date("2024-01-01"),
// 			types.Date("2024-01-01"),
// 			types.Bool(true),
// 			types.Bool(true),
// 			"displayName",
// 			"subjectKey",
// 			"description",
// 		).
// 		SetCurrency("FOO", 1000, t.sender[0].Address(), t.currency, false).
// 		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(), t.contract[0].Address(), t.creator[0].Address(), t.currency[0]).
// 		RunPreProcess()

// 	if assert.NotNil(t.Suite.T(), err) {
// 		t.Suite.T().Log(err.Error())
// 	}
// }

// func (t *testAddTemplate) Test07ContractFound() {
// 	err := t.Create().
// 		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, true).
// 		SetContractAccount(t.sender[0].Address(), t.contractKey, 1000, t.GenesisCurrency, t.contract, false).
// 		SetAccount(t.creatorKey, 1000, t.GenesisCurrency, t.creator, true).
// 		SetTemplate(
// 			"templateID",
// 			"templateName",
// 			types.Date("2024-01-01"),
// 			types.Date("2024-01-01"),
// 			types.Bool(true),
// 			types.Bool(true),
// 			"displayName",
// 			"subjectKey",
// 			"description",
// 		).
// 		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(), t.contract[0].Address(), t.creator[0].Address(), t.GenesisCurrency).
// 		RunPreProcess()

// 	if assert.NotNil(t.Suite.T(), err) {
// 		t.Suite.T().Log(err.Error())
// 	}
// }

// func (t *testAddTemplate) Test08SenderNotAuthorized() {
// 	err := t.Create().
// 		SetAccount(t.accountKey, 1000, t.GenesisCurrency, t.account, true).
// 		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, true).
// 		SetContractAccount(t.sender[0].Address(), t.contractKey, 1000, t.GenesisCurrency, t.contract, true).
// 		SetAccount(t.creatorKey, 1000, t.GenesisCurrency, t.creator, true).
// 		SetTemplate(
// 			"templateID",
// 			"templateName",
// 			types.Date("2024-01-01"),
// 			types.Date("2024-01-01"),
// 			types.Bool(true),
// 			types.Bool(true),
// 			"displayName",
// 			"subjectKey",
// 			"description",
// 		).
// 		MakeOperation(t.account[0].Address(), t.account[0].Priv(), t.contract[0].Address(), t.creator[0].Address(), t.GenesisCurrency).
// 		RunPreProcess()

// 	if assert.NotNil(t.Suite.T(), err) {
// 		t.Suite.T().Log(err.Error())
// 	}
// }

// func (t *testAddTemplate) Test09ErrorSenderNotFound() {
// 	err := t.Create().
// 		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, true).
// 		SetContractAccount(t.sender[0].Address(), t.contractKey, 1000, t.GenesisCurrency, t.contract, true).
// 		SetAccount(t.creatorKey, 1000, t.GenesisCurrency, t.creator, true).
// 		SetTemplate(
// 			"!@#$%^&*(?/)",
// 			"templateName",
// 			types.Date("2024-01-01"),
// 			types.Date("2024-01-01"),
// 			types.Bool(true),
// 			types.Bool(true),
// 			"displayName",
// 			"subjectKey",
// 			"description",
// 		).
// 		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(), t.contract[0].Address(), t.creator[0].Address(), t.GenesisCurrency).
// 		IsValid()

// 	if assert.NotNil(t.Suite.T(), err) {
// 		t.Suite.T().Log(err.Error())
// 	}
// }

// func (t *testAddTemplate) Test10ErrorSenderNotFound() {
// 	err := t.Create().
// 		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, false).
// 		SetContractAccount(t.sender[0].Address(), t.contractKey, 1000, t.GenesisCurrency, t.contract, true).
// 		SetAccount(t.creatorKey, 1000, t.GenesisCurrency, t.creator, true).
// 		SetTemplate(
// 			"templateID",
// 			"templateName",
// 			types.Date("2024-01-01"),
// 			types.Date("2024-01-01"),
// 			types.Bool(true),
// 			types.Bool(true),
// 			"displayName",
// 			"subjectKey",
// 			"description",
// 		).
// 		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(), t.contract[0].Address(), t.creator[0].Address(), t.GenesisCurrency).
// 		RunPreProcess()

// 	if assert.NotNil(t.Suite.T(), err) {
// 		t.Suite.T().Log(err.Error())
// 	}
// }

// func (t *testAddTemplate) Test11SelfTargetedCreator() {
// 	err := t.Create().
// 		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, true).
// 		SetContractAccount(t.sender[0].Address(), t.contractKey, 1000, t.GenesisCurrency, t.contract, true).
// 		SetTemplate(
// 			"templateID",
// 			"templateName",
// 			types.Date("2024-01-01"),
// 			types.Date("2024-01-01"),
// 			types.Bool(true),
// 			types.Bool(true),
// 			"displayName",
// 			"subjectKey",
// 			"description",
// 		).
// 		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(), t.contract[0].Address(), t.contract[0].Address(), t.GenesisCurrency).
// 		IsValid()

// 	if assert.NotNil(t.Suite.T(), err) {
// 		t.Suite.T().Log(err.Error())
// 	}
// }

func (t *testAddTemplate) Test12SelfTargetedSender() {
	err := t.Create().
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, true).
		SetContractAccount(t.sender[0].Address(), t.contractKey, 1000, t.GenesisCurrency, t.contract, true).
		SetAccount(t.creatorKey, 1000, t.GenesisCurrency, t.creator, true).
		SetTemplate(
			"templateID",
			"templateName",
			types.Date("2024-01-01"),
			types.Date("2024-01-01"),
			types.Bool(true),
			types.Bool(true),
			"displayName",
			"subjectKey",
			"description",
		).
		MakeOperation(t.contract[0].Address(), t.contract[0].Priv(), t.contract[0].Address(), t.creator[0].Address(), t.GenesisCurrency).
		IsValid()

	if assert.NotNil(t.Suite.T(), err) {
		t.Suite.T().Log(err.Error())
	}
}

func TestAddTemplate(t *testing.T) {
	suite.Run(t, new(testAddTemplate))
}
