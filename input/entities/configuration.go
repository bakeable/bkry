package entities

func GetMetaData() []MetaData {
	return []MetaData{
		{
			TypeName: "UserProfile",
			Fields: userProfileFields,
		},
		{
			TypeName: "User",
			Fields: userFields,
		},
		{
			TypeName: "PortalEvent",
			Fields: portalEventFields,
		},
		{
			TypeName: "Media",
			Fields: mediaFields,
		},
		{
			TypeName: "AirlineSettings",
			Fields: airlineSettingsFields,
		},
		{
			TypeName: "DeliveryEntry",
			Fields: deliveryEntryFields,
		},
		{
			TypeName: "PrintingOrderSupplier",
			Fields: printingOrderSupplierFields,
		},
		{
			TypeName: "Product",
			Fields: productFields,
		},
		{
			TypeName: "ExaminationSettings",
			Fields: examinationSettingsFields,
		},
		{
			TypeName: "PriceConfiguration",
			Fields: priceConfigurationFields,
		},
		{
			TypeName: "PrintingOrderItem",
			Fields: printingOrderItemFields,
		},
		{
			TypeName: "PackingSlipPackage",
			Fields: packingSlipPackageFields,
		},
		{
			TypeName: "CloudFunction",
			Fields: cloudFunctionFields,
		},
		{
			TypeName: "ApiKey",
			Fields: apiKeyFields,
		},
		{
			TypeName: "Quotation",
			Fields: quotationFields,
		},
		{
			TypeName: "ExaminationProperty",
			Fields: examinationPropertyFields,
		},
		{
			TypeName: "AirlineOrderGrouping",
			Fields: airlineOrderGroupingFields,
		},
		{
			TypeName: "PriceLayer",
			Fields: priceLayerFields,
		},
		{
			TypeName: "AirlineOrderGroup",
			Fields: airlineOrderGroupFields,
		},
		{
			TypeName: "AirlineOrder",
			Fields: airlineOrderFields,
		},
		{
			TypeName: "AirlinePricing",
			Fields: airlinePricingFields,
		},
		{
			TypeName: "PrintingOrder",
			Fields: printingOrderFields,
		},
		{
			TypeName: "PackingSlip",
			Fields: packingSlipFields,
		},
		{
			TypeName: "Email",
			Fields: emailFields,
		},
		{
			TypeName: "BatchExport",
			Fields: exportFields,
		},
		{
			TypeName: "ExaminationTask",
			Fields: examinationTaskFields,
		},
		{
			TypeName: "Attribute",
			Fields: attributeFields,
		},
		{
			TypeName: "AttributeOptionSet",
			Fields: attributeOptionSetFields,
		},
		{
			TypeName: "AttributeOption",
			Fields: attributeOptionFields,
		},
	}
}
