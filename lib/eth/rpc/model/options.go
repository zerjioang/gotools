package model

type RpcOptions struct {
	DefaultAccount                string `json:"default_account"`
	DefaultBlock                  string `json:"default_block"`
	DefaultGas                    int    `json:"default_gas"`
	DefaultGasPrice               int    `json:"default_gas_price"`
	TransactionBlockTimeout       int    `json:"transaction_block_timeout"`
	TransactionConfirmationBlocks int    `json:"transaction_confirmation_blocks"`
	TransactionPollingTimeout     int    `json:"transaction_polling_timeout"`
}
