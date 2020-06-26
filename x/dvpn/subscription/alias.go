// nolint
// autogenerated code using github.com/rigelrozanski/multitool
// aliases generated for the following subdirectories:
// ALIASGEN: github.com/sentinel-official/hub/x/dvpn/subscription/types
// ALIASGEN: github.com/sentinel-official/hub/x/dvpn/subscription/keeper
// ALIASGEN: github.com/sentinel-official/hub/x/dvpn/subscription/querier
package subscription

import (
	"github.com/sentinel-official/hub/x/dvpn/subscription/keeper"
	"github.com/sentinel-official/hub/x/dvpn/subscription/querier"
	"github.com/sentinel-official/hub/x/dvpn/subscription/types"
)

const (
	Codespace                         = types.Codespace
	EventTypeSetPlan                  = types.EventTypeSetPlan
	EventTypeSetPlansCount            = types.EventTypeSetPlansCount
	EventTypeSetPlanStatus            = types.EventTypeSetPlanStatus
	EventTypeSetNodeAddressForPlan    = types.EventTypeSetNodeAddressForPlan
	EventTypeDeleteNodeAddressForPlan = types.EventTypeDeleteNodeAddressForPlan
	EventTypeSetSubscription          = types.EventTypeSetSubscription
	EventTypeSetSubscriptionsCount    = types.EventTypeSetSubscriptionsCount
	AttributeKeyAddress               = types.AttributeKeyAddress
	AttributeKeyID                    = types.AttributeKeyID
	AttributeKeyCount                 = types.AttributeKeyCount
	AttributeKeyStatus                = types.AttributeKeyStatus
	AttributeKeyPlan                  = types.AttributeKeyPlan
	ModuleName                        = types.ModuleName
	QuerierRoute                      = types.QuerierRoute
	QueryPlan                         = types.QueryPlan
	QueryPlans                        = types.QueryPlans
	QueryPlansOfProvider              = types.QueryPlansOfProvider
	QueryNodesOfPlan                  = types.QueryNodesOfPlan
)

var (
	// functions aliases
	RegisterCodec                 = types.RegisterCodec
	ErrorMarshal                  = types.ErrorMarshal
	ErrorUnmarshal                = types.ErrorUnmarshal
	ErrorUnknownMsgType           = types.ErrorUnknownMsgType
	ErrorUnknownQueryType         = types.ErrorUnknownQueryType
	ErrorInvalidField             = types.ErrorInvalidField
	ErrorProviderDoesNotExist     = types.ErrorProviderDoesNotExist
	ErrorPlanDoesNotExist         = types.ErrorPlanDoesNotExist
	ErrorNodeDoesNotExist         = types.ErrorNodeDoesNotExist
	ErrorUnauthorized             = types.ErrorUnauthorized
	ErrorDuplicateNode            = types.ErrorDuplicateNode
	ErrorNodeWasNotAdded          = types.ErrorNodeWasNotAdded
	ErrorInvalidPlanStatus        = types.ErrorInvalidPlanStatus
	ErrorPriceDoesNotExist        = types.ErrorPriceDoesNotExist
	ErrorInvalidNodeStatus        = types.ErrorInvalidNodeStatus
	NewGenesisState               = types.NewGenesisState
	DefaultGenesisState           = types.DefaultGenesisState
	PlanKey                       = types.PlanKey
	PlanIDForProviderKey          = types.PlanIDForProviderKey
	NodeAddressForPlanKey         = types.NodeAddressForPlanKey
	SubscriptionKey               = types.SubscriptionKey
	SubscriptionIDForAddressKey   = types.SubscriptionIDForAddressKey
	SubscriptionIDForPlanKey      = types.SubscriptionIDForPlanKey
	NewMsgAddPlan                 = types.NewMsgAddPlan
	NewMsgSetPlanStatus           = types.NewMsgSetPlanStatus
	NewMsgAddNodeForPlan          = types.NewMsgAddNodeForPlan
	NewMsgRemoveNodeForPlan       = types.NewMsgRemoveNodeForPlan
	NewQueryPlanParams            = types.NewQueryPlanParams
	NewQueryPlansOfProviderParams = types.NewQueryPlansOfProviderParams
	NewQueryNodesOfPlanParams     = types.NewQueryNodesOfPlanParams
	NewKeeper                     = keeper.NewKeeper
	Querier                       = querier.Querier

	// variable aliases
	ModuleCdc             = types.ModuleCdc
	RouterKey             = types.RouterKey
	StoreKey              = types.StoreKey
	PlansCountKey         = types.PlansCountKey
	PlanKeyPrefix         = types.PlanKeyPrefix
	SubscriptionsCountKey = types.SubscriptionsCountKey
	SubscriptionKeyPrefix = types.SubscriptionKeyPrefix
)

type (
	GenesisPlan                = types.GenesisPlan
	GenesisPlans               = types.GenesisPlans
	GenesisState               = types.GenesisState
	MsgAddPlan                 = types.MsgAddPlan
	MsgSetPlanStatus           = types.MsgSetPlanStatus
	MsgAddNodeForPlan          = types.MsgAddNodeForPlan
	MsgRemoveNodeForPlan       = types.MsgRemoveNodeForPlan
	MsgStartSubscription       = types.MsgStartSubscription
	MsgEndSubscription         = types.MsgEndSubscription
	Plan                       = types.Plan
	Plans                      = types.Plans
	QueryPlanParams            = types.QueryPlanParams
	QueryPlansOfProviderParams = types.QueryPlansOfProviderParams
	QueryNodesOfPlanParams     = types.QueryNodesOfPlanParams
	Subscription               = types.Subscription
	Subscriptions              = types.Subscriptions
	Keeper                     = keeper.Keeper
)
