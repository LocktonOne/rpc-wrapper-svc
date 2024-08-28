package helpers

//
//func GetAddressAllowedContractsRegistry(r *http.Request) (common.Address, error) {
//	address := handlers.RegistryConfig(r).Address
//
//	registry, err := master_contracts_registry.NewMasterContractsRegistry(address, handlers.EthRPCConfig(r).EthClient())
//	if err != nil {
//		return common.Address{}, err
//	}
//	return registry.GetMasterAccessManagement(&bind.CallOpts{})
//}
