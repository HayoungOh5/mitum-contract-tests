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

type testCreateService struct {
	suite.Suite
	credential.TestCreateServiceProcessor
	account     []test.Account
	sender      []test.Account
	contract    []test.Account
	currency    []currencytypes.CurrencyID
	accountKey  string // Private Key
	ownerKey    string // Private Key
	senderKey   string // Private Key
	contractKey string // Private Key
	creatorKey  string // Private Key
	owner       []test.Account
}

func (t *testCreateService) SetupTest() {
	opr := credential.NewTestCreateServiceProcessor(util.Encoders)
	t.TestCreateServiceProcessor = opr
	t.Setup()
	t.account = make([]test.Account, 1)
	t.owner = make([]test.Account, 1)
	t.sender = make([]test.Account, 1)
	t.contract = make([]test.Account, 1)
	t.currency = make([]currencytypes.CurrencyID, 1)
	t.accountKey = t.NewPrivateKey("account")
	t.ownerKey = t.NewPrivateKey("owner")
	t.senderKey = t.NewPrivateKey("sender")
	t.contractKey = t.NewPrivateKey("contract")
}

func (t *testCreateService) Test01ErrorSenderNotFound() {
	err := t.Create().
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, false).
		SetContractAccount(t.sender[0].Address(), t.contractKey, 1000, t.GenesisCurrency, t.contract, true).
		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(), t.contract[0].Address(), t.GenesisCurrency).
		RunPreProcess()

	if assert.NotNil(t.Suite.T(), err) {
		t.Suite.T().Log(err.Error())
	}
}

func (t *testCreateService) Test02SenderIsCA() {
	err := t.Create().
		SetAccount(t.ownerKey, 1000, t.GenesisCurrency, t.owner, true).
		SetContractAccount(t.owner[0].Address(), t.senderKey, 1000, t.GenesisCurrency, t.sender, true).
		SetContractAccount(t.sender[0].Address(), t.contractKey, 1000, t.GenesisCurrency, t.contract, true).
		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(), t.contract[0].Address(), t.GenesisCurrency).
		RunPreProcess()

	if assert.NotNil(t.Suite.T(), err) {
		t.Suite.T().Log(err.Error())
	}
}

func (t *testCreateService) Test03ContractNotFound() {
	err := t.Create().
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, true).
		SetContractAccount(t.sender[0].Address(), t.contractKey, 1000, t.GenesisCurrency, t.contract, false).
		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(), t.contract[0].Address(), t.GenesisCurrency).
		RunPreProcess()

	if assert.NotNil(t.Suite.T(), err) {
		t.Suite.T().Log(err.Error())
	}
}

func (t *testCreateService) Test04SenderNotAuthorized() {
	err := t.Create().
		SetAccount(t.accountKey, 1000, t.GenesisCurrency, t.account, true).
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, true).
		SetContractAccount(t.sender[0].Address(), t.contractKey, 1000, t.GenesisCurrency, t.contract, true).
		MakeOperation(t.account[0].Address(), t.account[0].Priv(), t.contract[0].Address(), t.GenesisCurrency).
		RunPreProcess()

	if assert.NotNil(t.Suite.T(), err) {
		t.Suite.T().Log(err.Error())
	}
}

func (t *testCreateService) Test05CurrencyNotFound() {
	err := t.Create().
		SetAccount(t.accountKey, 1000, t.GenesisCurrency, t.account, true).
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, true).
		SetContractAccount(t.sender[0].Address(), t.contractKey, 1000, t.GenesisCurrency, t.contract, true).
		SetCurrency("FOO", 1000, t.sender[0].Address(), t.currency, false).
		MakeOperation(t.account[0].Address(), t.account[0].Priv(), t.contract[0].Address(), t.currency[0]).
		RunPreProcess()

	if assert.NotNil(t.Suite.T(), err) {
		t.Suite.T().Log(err.Error())
	}
}

func (t *testCreateService) Test06InvalidSigning() {
	err := t.Create().
		SetAccount(t.accountKey, 1000, t.GenesisCurrency, t.account, true).
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, true).
		SetContractAccount(t.sender[0].Address(), t.contractKey, 1000, t.GenesisCurrency, t.contract, true).
		MakeOperation(t.sender[0].Address(), t.account[0].Priv(), t.contract[0].Address(), t.GenesisCurrency).
		RunPreProcess()

	if assert.NotNil(t.Suite.T(), err) {
		t.Suite.T().Log(err.Error())
	}
}

func (t *testCreateService) Test07SelfTargetedSender() {
	err := t.Create().
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, true).
		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(), t.sender[0].Address(), t.GenesisCurrency).
		RunPreProcess()

	if assert.NotNil(t.Suite.T(), err) {
		t.Suite.T().Log(err.Error())
	}
}

func TestCreateService(t *testing.T) {
	suite.Run(t, new(testCreateService))
}
