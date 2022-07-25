package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmic-horizon/coho/x/coho/types"
	"github.com/cosmos/cosmos-sdk/client"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	return cmd
}

//   rpc TransferModuleOwnership(MsgTransferModuleOwnership) returns (MsgTransferModuleOwnershipResponse);
//   rpc WhitelistNftContracts(MsgWhitelistNftContracts) returns (MsgWhitelistNftContractsResponse);
//   rpc RemoveWhitelistedNftContracts(MsgRemoveWhitelistedNftContracts) returns (MsgRemoveWhitelistedNftContractsResponse);
//   rpc DepositNft(MsgDepositNft) returns (MsgDepositNftResponse);
//   rpc WithdrawUpdatedNft(MsgWithdrawUpdatedNft) returns (MsgWithdrawUpdatedNftResponse);
