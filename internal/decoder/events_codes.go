package decoder

func GetAllEventsCodes() map[string]string {
	eventToHash := map[string]string{
		"MintedToTreasury(address,uint256)":                                      "0xbfa21aa5d5f9a1f0120a95e7c0749f389863cbdbfff531aa7339077a5bc919de",
		"Mint(address,address,uint256,uint256,uint256)":                          "0x458f5fa412d0f69b08dd84872b0215675cc67bc1d5b6fd93300a1c3878b86196",
		"BorrowAllowanceDelegated(address,address,address,uint256)":              "0xda919360433220e13b51e8c211e490d148e61a3bd53de8c097194e458b97f3e1",
		"Initialized(address,address,address,uint8,string,string,bytes)":         "0x40251fbfb6656cfa65a00d7879029fec1fad21d28fdcff2f4f68f52795b74f2c",
		"ReserveUsedAsCollateralDisabled(address,address)":                       "0x44c58d81365b66dd4b1a7f36c25aa97b8c71c361ee4937adc1a00000227db5dd",
		"Withdraw(address,address,address,uint256)":                              "0x3115d1449a7b732c986cba18244e897a450f61e1bb8d589cd2e69e6c8924f9f7",
		"FlashLoan(address,address,address,uint256,uint8,uint256,uint16)":        "0xefefaba5e921573100900a3ad9cf29f222d995fb3b6045797eaea7521bd8d6f0",
		"Transfer(address,address,uint256)":                                      "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef",
		"UserEModeSet(address,uint8)":                                            "0xd728da875fc88944cbf17638bcbe4af0eedaef63becd1d1c57cc097eb4608d84",
		"Supply(address,address,address,uint256,uint16)":                         "0x2b627736bca15cd5381dcf80b0bf11fd197d01a037c52b927a881a10fb73ba61",
		"DeficitCovered(address,address,uint256)":                                "0x84b203e49f1a4b553088061534231969a68ad1c81be192205e96d23a206cb26a",
		"BalanceTransfer(address,address,uint256,uint256)":                       "0x4beccb90f994c31aced7a23b5611020728a23d8ec5cddd1a3e9d97b96fda8666",
		"MintUnbacked(address,address,address,uint256,uint16)":                   "0xf25af37b3d3ec226063dc9bdc103ece7eb110a50f340fe854bb7bc1b0676d7d0",
		"Repay(address,address,address,uint256,bool)":                            "0xa534c8dbe71f871f9f3530e97a74601fea17b426cae02e1c5aee42c96c784051",
		"IsolationModeTotalDebtUpdated(address,uint256)":                         "0xaef84d3b40895fd58c561f3998000f0583abb992a52fbdc99ace8e8de4d676a5",
		"ReserveDataUpdated(address,uint256,uint256,uint256,uint256,uint256)":    "0x804c9b842b2748a22bb64b345453a3de7ca54a6ca45ce00d415894979e22897a",
		"Burn(address,address,uint256,uint256,uint256)":                          "0x4cf25bc1d991c17529c25213d3cc0cda295eeaad5f13f361969b12ea48015f90",
		"BackUnbacked(address,address,uint256,uint256)":                          "0x281596e92b2d974beb7d4f124df30a0b39067b096893e95011ce4bdad798b759",
		"LiquidationCall(address,address,address,uint256,uint256,address,bool)":  "0xe413a321e8681d831f4dbccbca790d2952b56f977908e45be37335533e005286",
		"Borrow(address,address,address,uint256,uint8,uint256,uint16)":           "0xb3d084820fb1a9decffb176436bd02558d15fac9b0ddfed8c465bc7359d7dce0",
		"OwnershipTransferred(address,address)":                                  "0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0",
		"Approval(address,address,uint256)":                                      "0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925",
		"ReserveUsedAsCollateralEnabled(address,address)":                        "0x00058a56ea94653cdf4f152d227ace22d4c00ad99e2a43f58cb7d9e3feb295f2",
		"DeficitCreated(address,address,uint256)":                                "0x2bccfb3fad376d59d7accf970515eb77b2f27b082c90ed0fb15583dd5a942699",
		"Initialized(address,address,address,address,uint8,string,string,bytes)": "0xb19e051f8af41150ccccb3fc2c2d8d15f4a4cf434f32a559ba75fe73d6eea20b",
	}
	return eventToHash
}
