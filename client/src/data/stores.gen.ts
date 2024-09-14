
import { createUserProfileStore } from './entities/user_profile/store.gen';
import { createUserStore } from './entities/user/store.gen';
import { createPortalEventStore } from './entities/portal_event/store.gen';
import { createMediaStore } from './entities/media/store.gen';
import { createAirlineSettingsStore } from './entities/airline_settings/store.gen';
import { createDeliveryEntryStore } from './entities/delivery_entry/store.gen';
import { createPrintingOrderSupplierStore } from './entities/printing_order_supplier/store.gen';
import { createProductStore } from './entities/product/store.gen';
import { createExaminationSettingsStore } from './entities/examination_settings/store.gen';
import { createPriceConfigurationStore } from './entities/price_configuration/store.gen';
import { createPrintingOrderItemStore } from './entities/printing_order_item/store.gen';
import { createPackingSlipPackageStore } from './entities/packing_slip_package/store.gen';
import { createCloudFunctionStore } from './entities/cloud_function/store.gen';
import { createApiKeyStore } from './entities/api_key/store.gen';
import { createQuotationStore } from './entities/quotation/store.gen';
import { createExaminationPropertyStore } from './entities/examination_property/store.gen';
import { createAirlineOrderGroupingStore } from './entities/airline_order_grouping/store.gen';
import { createPriceLayerStore } from './entities/price_layer/store.gen';
import { createAirlineOrderGroupStore } from './entities/airline_order_group/store.gen';
import { createAirlineOrderStore } from './entities/airline_order/store.gen';
import { createAirlinePricingStore } from './entities/airline_pricing/store.gen';
import { createPrintingOrderStore } from './entities/printing_order/store.gen';
import { createPackingSlipStore } from './entities/packing_slip/store.gen';
import { createEmailStore } from './entities/email/store.gen';
import { createBatchExportStore } from './entities/batch_export/store.gen';
import { createExaminationTaskStore } from './entities/examination_task/store.gen';
import { createAttributeStore } from './entities/attribute/store.gen';
import { createAttributeOptionSetStore } from './entities/attribute_option_set/store.gen';
import { createAttributeOptionStore } from './entities/attribute_option/store.gen';
import { type EntityTypeName } from './types/types.gen';

const stores: Record<EntityTypeName, any> = {
    UserProfile: createUserProfileStore,
    User: createUserStore,
    PortalEvent: createPortalEventStore,
    Media: createMediaStore,
    AirlineSettings: createAirlineSettingsStore,
    DeliveryEntry: createDeliveryEntryStore,
    PrintingOrderSupplier: createPrintingOrderSupplierStore,
    Product: createProductStore,
    ExaminationSettings: createExaminationSettingsStore,
    PriceConfiguration: createPriceConfigurationStore,
    PrintingOrderItem: createPrintingOrderItemStore,
    PackingSlipPackage: createPackingSlipPackageStore,
    CloudFunction: createCloudFunctionStore,
    ApiKey: createApiKeyStore,
    Quotation: createQuotationStore,
    ExaminationProperty: createExaminationPropertyStore,
    AirlineOrderGrouping: createAirlineOrderGroupingStore,
    PriceLayer: createPriceLayerStore,
    AirlineOrderGroup: createAirlineOrderGroupStore,
    AirlineOrder: createAirlineOrderStore,
    AirlinePricing: createAirlinePricingStore,
    PrintingOrder: createPrintingOrderStore,
    PackingSlip: createPackingSlipStore,
    Email: createEmailStore,
    BatchExport: createBatchExportStore,
    ExaminationTask: createExaminationTaskStore,
    Attribute: createAttributeStore,
    AttributeOptionSet: createAttributeOptionSetStore,
    AttributeOption: createAttributeOptionStore,
}
    

export default stores