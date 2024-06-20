package app

import (
	"fmt"

	"github.com/CosmWasm/wasmd/x/wasm"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	bam "github.com/cosmos/cosmos-sdk/baseapp"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authAnte "github.com/cosmos/cosmos-sdk/x/auth/ante"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/capability"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	crisiskeeper "github.com/cosmos/cosmos-sdk/x/crisis/keeper"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"
	distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/cosmos/cosmos-sdk/x/evidence"
	evidencekeeper "github.com/cosmos/cosmos-sdk/x/evidence/keeper"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	feegrantkeeper "github.com/cosmos/cosmos-sdk/x/feegrant/keeper"
	feegrantmodule "github.com/cosmos/cosmos-sdk/x/feegrant/module"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/cosmos/cosmos-sdk/x/gov"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/cosmos/cosmos-sdk/x/mint"
	mintkeeper "github.com/cosmos/cosmos-sdk/x/mint/keeper"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/cosmos/cosmos-sdk/x/upgrade"
	upgradekeeper "github.com/cosmos/cosmos-sdk/x/upgrade/keeper"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	"github.com/cosmos/ibc-go/v4/modules/apps/transfer"
	ibctransferkeeper "github.com/cosmos/ibc-go/v4/modules/apps/transfer/keeper"
	ibctransfertypes "github.com/cosmos/ibc-go/v4/modules/apps/transfer/types"
	ibc "github.com/cosmos/ibc-go/v4/modules/core"
	ibchost "github.com/cosmos/ibc-go/v4/modules/core/24-host"
	ibcante "github.com/cosmos/ibc-go/v4/modules/core/ante"
	ibckeeper "github.com/cosmos/ibc-go/v4/modules/core/keeper"
	ibchooks "github.com/osmosis-labs/osmosis/x/ibc-hooks"
	ibchookstypes "github.com/osmosis-labs/osmosis/x/ibc-hooks/types"
	"github.com/spf13/cast"

	axelarParams "github.com/axelarnetwork/axelar-core/app/params"
	"github.com/axelarnetwork/axelar-core/x/ante"
	"github.com/axelarnetwork/axelar-core/x/auxiliary"
	auxiliarytypes "github.com/axelarnetwork/axelar-core/x/auxiliary/types"
	"github.com/axelarnetwork/axelar-core/x/axelarnet"
	axelarnetKeeper "github.com/axelarnetwork/axelar-core/x/axelarnet/keeper"
	axelarnetTypes "github.com/axelarnetwork/axelar-core/x/axelarnet/types"
	axelarbankkeeper "github.com/axelarnetwork/axelar-core/x/bank/keeper"
	"github.com/axelarnetwork/axelar-core/x/evm"
	evmKeeper "github.com/axelarnetwork/axelar-core/x/evm/keeper"
	evmTypes "github.com/axelarnetwork/axelar-core/x/evm/types"
	"github.com/axelarnetwork/axelar-core/x/multisig"
	multisigKeeper "github.com/axelarnetwork/axelar-core/x/multisig/keeper"
	multisigTypes "github.com/axelarnetwork/axelar-core/x/multisig/types"
	"github.com/axelarnetwork/axelar-core/x/nexus"
	nexusKeeper "github.com/axelarnetwork/axelar-core/x/nexus/keeper"
	nexusTypes "github.com/axelarnetwork/axelar-core/x/nexus/types"
	"github.com/axelarnetwork/axelar-core/x/permission"
	permissionKeeper "github.com/axelarnetwork/axelar-core/x/permission/keeper"
	permissionTypes "github.com/axelarnetwork/axelar-core/x/permission/types"
	"github.com/axelarnetwork/axelar-core/x/reward"
	rewardKeeper "github.com/axelarnetwork/axelar-core/x/reward/keeper"
	rewardTypes "github.com/axelarnetwork/axelar-core/x/reward/types"
	"github.com/axelarnetwork/axelar-core/x/snapshot"
	snapKeeper "github.com/axelarnetwork/axelar-core/x/snapshot/keeper"
	snapTypes "github.com/axelarnetwork/axelar-core/x/snapshot/types"
	"github.com/axelarnetwork/axelar-core/x/tss"
	tssKeeper "github.com/axelarnetwork/axelar-core/x/tss/keeper"
	tssTypes "github.com/axelarnetwork/axelar-core/x/tss/types"
	"github.com/axelarnetwork/axelar-core/x/vote"
	voteKeeper "github.com/axelarnetwork/axelar-core/x/vote/keeper"
	voteTypes "github.com/axelarnetwork/axelar-core/x/vote/types"

	// Override with generated statik docs
	_ "github.com/axelarnetwork/axelar-core/client/docs/statik"
)

func initAppModules(
	keepers *KeeperCache,
	bApp *bam.BaseApp,
	encodingConfig axelarParams.EncodingConfig,
	appOpts servertypes.AppOptions,
	axelarnetModule axelarnet.AppModule,
) []module.AppModule {
	// NOTE: we may consider parsing `appOpts` inside module constructors. For the moment
	// we prefer to be more strict in what arguments the modules expect.
	var skipGenesisInvariants = cast.ToBool(appOpts.Get(crisis.FlagSkipGenesisInvariants))

	appCodec := encodingConfig.Codec

	appModules := []module.AppModule{
		genutil.NewAppModule(
			GetKeeper[authkeeper.AccountKeeper](keepers),
			GetKeeper[stakingkeeper.Keeper](keepers), bApp.DeliverTx, encodingConfig.TxConfig),
		auth.NewAppModule(appCodec,
			*GetKeeper[authkeeper.AccountKeeper](keepers), nil),
		vesting.NewAppModule(
			*GetKeeper[authkeeper.AccountKeeper](keepers),
			GetKeeper[bankkeeper.BaseKeeper](keepers)),

		// bank module accepts a reference to the base keeper, but panics when RegisterService is called on a reference, so we need to dereference it
		bank.NewAppModule(appCodec,
			*GetKeeper[bankkeeper.BaseKeeper](keepers),
			GetKeeper[authkeeper.AccountKeeper](keepers)),
		crisis.NewAppModule(GetKeeper[crisiskeeper.Keeper](keepers), skipGenesisInvariants),
		gov.NewAppModule(appCodec,
			*GetKeeper[govkeeper.Keeper](keepers),
			GetKeeper[authkeeper.AccountKeeper](keepers),
			GetKeeper[bankkeeper.BaseKeeper](keepers)),
		mint.NewAppModule(appCodec,
			*GetKeeper[mintkeeper.Keeper](keepers),
			GetKeeper[authkeeper.AccountKeeper](keepers)),
		slashing.NewAppModule(appCodec,
			*GetKeeper[slashingkeeper.Keeper](keepers),
			GetKeeper[authkeeper.AccountKeeper](keepers),
			GetKeeper[bankkeeper.BaseKeeper](keepers),
			GetKeeper[stakingkeeper.Keeper](keepers)),
		distr.NewAppModule(appCodec,
			*GetKeeper[distrkeeper.Keeper](keepers),
			GetKeeper[authkeeper.AccountKeeper](keepers),
			GetKeeper[bankkeeper.BaseKeeper](keepers),
			GetKeeper[stakingkeeper.Keeper](keepers)),
		staking.NewAppModule(appCodec,
			*GetKeeper[stakingkeeper.Keeper](keepers),
			GetKeeper[authkeeper.AccountKeeper](keepers),
			GetKeeper[bankkeeper.BaseKeeper](keepers)),
		upgrade.NewAppModule(
			*GetKeeper[upgradekeeper.Keeper](keepers)),
		evidence.NewAppModule(
			*GetKeeper[evidencekeeper.Keeper](keepers)),
		params.NewAppModule(
			*GetKeeper[paramskeeper.Keeper](keepers)),
		capability.NewAppModule(appCodec,
			*GetKeeper[capabilitykeeper.Keeper](keepers)),
	}

	// wasm module needs to be added in a specific order, so we cannot just append it at the end
	if IsWasmEnabled() {
		appModules = append(
			appModules,
			wasm.NewAppModule(
				appCodec,
				GetKeeper[wasm.Keeper](keepers),
				GetKeeper[stakingkeeper.Keeper](keepers),
				GetKeeper[authkeeper.AccountKeeper](keepers),
				GetKeeper[bankkeeper.BaseKeeper](keepers),
			),
		)
	}

	if IsIBCWasmHooksEnabled() {
		appModules = append(appModules, ibchooks.NewAppModule(GetKeeper[authkeeper.AccountKeeper](keepers)))
	}

	appModules = append(appModules,
		evidence.NewAppModule(*GetKeeper[evidencekeeper.Keeper](keepers)),
		ibc.NewAppModule(GetKeeper[ibckeeper.Keeper](keepers)),
		transfer.NewAppModule(*GetKeeper[ibctransferkeeper.Keeper](keepers)),
		feegrantmodule.NewAppModule(
			appCodec,
			GetKeeper[authkeeper.AccountKeeper](keepers),
			GetKeeper[bankkeeper.BaseKeeper](keepers),
			*GetKeeper[feegrantkeeper.Keeper](keepers),
			encodingConfig.InterfaceRegistry,
		),
		snapshot.NewAppModule(*GetKeeper[snapKeeper.Keeper](keepers)),
		multisig.NewAppModule(
			*GetKeeper[multisigKeeper.Keeper](keepers),
			GetKeeper[stakingkeeper.Keeper](keepers),
			GetKeeper[slashingkeeper.Keeper](keepers),
			GetKeeper[snapKeeper.Keeper](keepers),
			GetKeeper[rewardKeeper.Keeper](keepers),
			GetKeeper[nexusKeeper.Keeper](keepers),
		),
		tss.NewAppModule(
			*GetKeeper[tssKeeper.Keeper](keepers),
			GetKeeper[snapKeeper.Keeper](keepers),
			GetKeeper[nexusKeeper.Keeper](keepers),
			GetKeeper[stakingkeeper.Keeper](keepers),
			GetKeeper[multisigKeeper.Keeper](keepers),
		),
		vote.NewAppModule(*GetKeeper[voteKeeper.Keeper](keepers)),
		nexus.NewAppModule(
			*GetKeeper[nexusKeeper.Keeper](keepers),
			GetKeeper[snapKeeper.Keeper](keepers),
			GetKeeper[slashingkeeper.Keeper](keepers),
			GetKeeper[stakingkeeper.Keeper](keepers),
			GetKeeper[axelarnetKeeper.Keeper](keepers),
			GetKeeper[rewardKeeper.Keeper](keepers),
		),
		evm.NewAppModule(
			GetKeeper[evmKeeper.BaseKeeper](keepers),
			GetKeeper[voteKeeper.Keeper](keepers),
			GetKeeper[nexusKeeper.Keeper](keepers),
			GetKeeper[snapKeeper.Keeper](keepers),
			GetKeeper[stakingkeeper.Keeper](keepers),
			GetKeeper[slashingkeeper.Keeper](keepers),
			GetKeeper[multisigKeeper.Keeper](keepers),
		),
		axelarnetModule,
		reward.NewAppModule(
			*GetKeeper[rewardKeeper.Keeper](keepers),
			GetKeeper[nexusKeeper.Keeper](keepers),
			GetKeeper[mintkeeper.Keeper](keepers),
			GetKeeper[stakingkeeper.Keeper](keepers),
			GetKeeper[slashingkeeper.Keeper](keepers),
			GetKeeper[multisigKeeper.Keeper](keepers),
			GetKeeper[snapKeeper.Keeper](keepers),
			GetKeeper[bankkeeper.BaseKeeper](keepers),
			bApp.MsgServiceRouter(),
			bApp.Router(),
		),
		permission.NewAppModule(*GetKeeper[permissionKeeper.Keeper](keepers)),
		auxiliary.NewAppModule(encodingConfig.Codec, bApp.MsgServiceRouter()),
	)
	return appModules
}

// func mustReadWasmConfig(appOpts servertypes.AppOptions) wasmtypes.WasmConfig {
// 	wasmConfig, err := wasm.ReadWasmConfig(appOpts)
// 	if err != nil {
// 		panic(fmt.Sprintf("error while reading wasm config: %s", err))
// 	}
// 	return wasmConfig
// }

func initAnteHandlers(encodingConfig axelarParams.EncodingConfig, keys map[string]*sdk.KVStoreKey, keepers *KeeperCache, appOpts servertypes.AppOptions) sdk.AnteHandler {
	// The baseAnteHandler handles signature verification and transaction pre-processing
	baseAnteHandler, err := authAnte.NewAnteHandler(
		authAnte.HandlerOptions{
			AccountKeeper:   GetKeeper[authkeeper.AccountKeeper](keepers),
			BankKeeper:      GetKeeper[bankkeeper.BaseKeeper](keepers),
			SignModeHandler: encodingConfig.TxConfig.SignModeHandler(),
			FeegrantKeeper:  GetKeeper[feegrantkeeper.Keeper](keepers),
			SigGasConsumer:  authAnte.DefaultSigVerificationGasConsumer,
		},
	)
	if err != nil {
		panic(err)
	}

	customAnteDecorators := initCustomAnteDecorators(encodingConfig, keys, keepers, appOpts)
	anteDecorators := append([]sdk.AnteDecorator{ante.NewAnteHandlerDecorator(baseAnteHandler)}, customAnteDecorators...)

	return sdk.ChainAnteDecorators(anteDecorators...)
}

func initCustomAnteDecorators(
	encodingConfig axelarParams.EncodingConfig,
	keys map[string]*sdk.KVStoreKey,
	keepers *KeeperCache,
	appOpts servertypes.AppOptions,
) []sdk.AnteDecorator {
	var anteDecorators []sdk.AnteDecorator

	// unwrap batch messages, must be done before any other custom decorators
	anteDecorators = append(anteDecorators, ante.NewBatchDecorator(encodingConfig.Codec))

	// enforce wasm limits earlier in the ante handler chain
	if IsWasmEnabled() {
		wasmConfig, err := wasm.ReadWasmConfig(appOpts)
		if err != nil {
			panic(fmt.Sprintf("error while reading wasm config: %s", err))
		}
		// wasmConfig := mustReadWasmConfig(appOpts)
		wasmAnteDecorators := []sdk.AnteDecorator{
			wasmkeeper.NewLimitSimulationGasDecorator(wasmConfig.SimulationGasLimit),
			wasmkeeper.NewCountTXDecorator(keys[wasm.StoreKey]),
		}

		anteDecorators = append(anteDecorators, wasmAnteDecorators...)
	}

	anteDecorators = append(anteDecorators,
		ibcante.NewAnteDecorator(GetKeeper[ibckeeper.Keeper](keepers)),
		ante.NewCheckRefundFeeDecorator(
			encodingConfig.InterfaceRegistry,
			GetKeeper[authkeeper.AccountKeeper](keepers),
			GetKeeper[stakingkeeper.Keeper](keepers),
			GetKeeper[snapKeeper.Keeper](keepers),
			GetKeeper[rewardKeeper.Keeper](keepers),
		),
		ante.NewAnteHandlerDecorator(
			initMessageAnteDecorators(encodingConfig, keepers).ToAnteHandler()),
	)
	return anteDecorators
}

func initMessageAnteDecorators(encodingConfig axelarParams.EncodingConfig, keepers *KeeperCache) ante.MessageAnteHandler {
	return ante.ChainMessageAnteDecorators(
		ante.NewLogMsgDecorator(encodingConfig.Codec),
		ante.NewCheckCommissionRate(GetKeeper[stakingkeeper.Keeper](keepers)),
		ante.NewUndelegateDecorator(
			GetKeeper[multisigKeeper.Keeper](keepers),
			GetKeeper[nexusKeeper.Keeper](keepers),
			GetKeeper[snapKeeper.Keeper](keepers),
		),

		ante.NewCheckProxy(GetKeeper[snapKeeper.Keeper](keepers)),
		ante.NewRestrictedTx(GetKeeper[permissionKeeper.Keeper](keepers)),
	)
}

func initMessageRouter(keepers *KeeperCache) nexusTypes.MessageRouter {
	messageRouter := nexusTypes.NewMessageRouter().
		AddRoute(evmTypes.ModuleName, evmKeeper.NewMessageRoute()).
		AddRoute(axelarnetTypes.ModuleName, axelarnetKeeper.NewMessageRoute(
			*GetKeeper[axelarnetKeeper.Keeper](keepers),
			GetKeeper[axelarnetKeeper.IBCKeeper](keepers),
			GetKeeper[feegrantkeeper.Keeper](keepers),
			axelarbankkeeper.NewBankKeeper(GetKeeper[bankkeeper.BaseKeeper](keepers)),
			GetKeeper[nexusKeeper.Keeper](keepers),
			GetKeeper[authkeeper.AccountKeeper](keepers),
		))

	if IsWasmEnabled() {
		messageRouter.AddRoute(wasm.ModuleName, nexusKeeper.NewMessageRoute(
			GetKeeper[nexusKeeper.Keeper](keepers),
			GetKeeper[authkeeper.AccountKeeper](keepers),
			GetKeeper[wasmkeeper.PermissionedKeeper](keepers),
		))
	}
	return messageRouter
}

func orderMigrations() []string {
	migrationOrder := []string{
		// auth module needs to go first
		authtypes.ModuleName,
		// sdk modules
		upgradetypes.ModuleName,
		capabilitytypes.ModuleName,
		crisistypes.ModuleName,
		govtypes.ModuleName,
		stakingtypes.ModuleName,
		ibctransfertypes.ModuleName,
		ibchost.ModuleName,
		banktypes.ModuleName,
		distrtypes.ModuleName,
		slashingtypes.ModuleName,
		minttypes.ModuleName,
		genutiltypes.ModuleName,
		evidencetypes.ModuleName,
		feegrant.ModuleName,
		paramstypes.ModuleName,
		vestingtypes.ModuleName,
	}

	// wasm module needs to be added in a specific order, so we cannot just append it at the end
	if IsWasmEnabled() {
		migrationOrder = append(migrationOrder, wasm.ModuleName)
	}

	if IsIBCWasmHooksEnabled() {
		migrationOrder = append(migrationOrder, ibchookstypes.ModuleName)
	}

	// axelar modules
	migrationOrder = append(migrationOrder,
		multisigTypes.ModuleName,
		tssTypes.ModuleName,
		rewardTypes.ModuleName,
		voteTypes.ModuleName,
		evmTypes.ModuleName,
		nexusTypes.ModuleName,
		permissionTypes.ModuleName,
		snapTypes.ModuleName,
		axelarnetTypes.ModuleName,
		auxiliarytypes.ModuleName,
	)
	return migrationOrder
}

func orderBeginBlockers() []string {
	// During begin block slashing happens after distr.BeginBlocker so that
	// there is nothing left over in the validator fee pool, so as to keep the
	// CanWithdrawInvariant invariant.
	// NOTE: staking module is required if HistoricalEntries param > 0
	beginBlockerOrder := []string{
		// upgrades should be run first
		upgradetypes.ModuleName,
		capabilitytypes.ModuleName,
		crisistypes.ModuleName,
		govtypes.ModuleName,
		stakingtypes.ModuleName,
		ibctransfertypes.ModuleName,
		ibchost.ModuleName,
		authtypes.ModuleName,
		banktypes.ModuleName,
		distrtypes.ModuleName,
		slashingtypes.ModuleName,
		minttypes.ModuleName,
		genutiltypes.ModuleName,
		evidencetypes.ModuleName,
		feegrant.ModuleName,
		paramstypes.ModuleName,
		vestingtypes.ModuleName,
	}

	// wasm module needs to be added in a specific order, so we cannot just append it at the end
	if IsWasmEnabled() {
		beginBlockerOrder = append(beginBlockerOrder, wasm.ModuleName)
	}

	if IsIBCWasmHooksEnabled() {
		beginBlockerOrder = append(beginBlockerOrder, ibchookstypes.ModuleName)
	}

	// axelar custom modules
	beginBlockerOrder = append(beginBlockerOrder,
		rewardTypes.ModuleName,
		nexusTypes.ModuleName,
		permissionTypes.ModuleName,
		multisigTypes.ModuleName,
		tssTypes.ModuleName,
		evmTypes.ModuleName,
		snapTypes.ModuleName,
		axelarnetTypes.ModuleName,
		voteTypes.ModuleName,
		auxiliarytypes.ModuleName,
	)
	return beginBlockerOrder
}

func orderEndBlockers() []string {
	endBlockerOrder := []string{
		crisistypes.ModuleName,
		govtypes.ModuleName,
		stakingtypes.ModuleName,
		ibctransfertypes.ModuleName,
		ibchost.ModuleName,
		feegrant.ModuleName,
		capabilitytypes.ModuleName,
		authtypes.ModuleName,
		banktypes.ModuleName,
		distrtypes.ModuleName,
		slashingtypes.ModuleName,
		minttypes.ModuleName,
		genutiltypes.ModuleName,
		evidencetypes.ModuleName,
		paramstypes.ModuleName,
		upgradetypes.ModuleName,
		vestingtypes.ModuleName,
	}

	// wasm module needs to be added in a specific order, so we cannot just append it at the end
	if IsWasmEnabled() {
		endBlockerOrder = append(endBlockerOrder, wasm.ModuleName)
	}

	if IsIBCWasmHooksEnabled() {
		endBlockerOrder = append(endBlockerOrder, ibchookstypes.ModuleName)
	}

	// axelar custom modules
	endBlockerOrder = append(endBlockerOrder,
		multisigTypes.ModuleName,
		tssTypes.ModuleName,
		evmTypes.ModuleName,
		nexusTypes.ModuleName,
		rewardTypes.ModuleName,
		snapTypes.ModuleName,
		axelarnetTypes.ModuleName,
		permissionTypes.ModuleName,
		voteTypes.ModuleName,
		auxiliarytypes.ModuleName,
	)
	return endBlockerOrder
}

func orderModulesForGenesis() []string {
	// Sets the order of Genesis - Order matters, genutil is to always come last
	// NOTE: The genutils module must occur after staking so that pools are
	// properly initialized with tokens from genesis accounts.
	genesisOrder := []string{
		capabilitytypes.ModuleName,
		authtypes.ModuleName,
		banktypes.ModuleName,
		distrtypes.ModuleName,
		stakingtypes.ModuleName,
		slashingtypes.ModuleName,
		govtypes.ModuleName,
		minttypes.ModuleName,
		crisistypes.ModuleName,
		genutiltypes.ModuleName,
		evidencetypes.ModuleName,
		ibchost.ModuleName,
		evidencetypes.ModuleName,
		ibctransfertypes.ModuleName,
		feegrant.ModuleName,
		paramstypes.ModuleName,
		upgradetypes.ModuleName,
		vestingtypes.ModuleName,
	}

	// wasm module needs to be added in a specific order, so we cannot just append it at the end
	if IsWasmEnabled() {
		genesisOrder = append(genesisOrder, wasm.ModuleName)
	}

	if IsIBCWasmHooksEnabled() {
		genesisOrder = append(genesisOrder, ibchookstypes.ModuleName)
	}

	genesisOrder = append(genesisOrder,
		snapTypes.ModuleName,
		multisigTypes.ModuleName,
		tssTypes.ModuleName,
		nexusTypes.ModuleName,
		evmTypes.ModuleName, // Run evm end blocker after nexus so GMP calls routed to EVM get processed within the same block
		voteTypes.ModuleName,
		axelarnetTypes.ModuleName,
		rewardTypes.ModuleName,
		permissionTypes.ModuleName,
		auxiliarytypes.ModuleName,
	)
	return genesisOrder
}
