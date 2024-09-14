
import { UserProfile, UserProfileList, createUserProfileStore } from './user_profile'
import { User, UserList, createUserStore } from './user'
import { PortalEvent, PortalEventList, createPortalEventStore } from './portal_event'
import { Media, MediaList, createMediaStore } from './media'
import { AirlineSettings, AirlineSettingsList, createAirlineSettingsStore } from './airline_settings'
import { DeliveryEntry, DeliveryEntryList, createDeliveryEntryStore } from './delivery_entry'
import { PrintingOrderSupplier, PrintingOrderSupplierList, createPrintingOrderSupplierStore } from './printing_order_supplier'
import { Product, ProductList, createProductStore } from './product'
import { ExaminationSettings, ExaminationSettingsList, createExaminationSettingsStore } from './examination_settings'
import { PriceConfiguration, PriceConfigurationList, createPriceConfigurationStore } from './price_configuration'
import { PrintingOrderItem, PrintingOrderItemList, createPrintingOrderItemStore } from './printing_order_item'
import { PackingSlipPackage, PackingSlipPackageList, createPackingSlipPackageStore } from './packing_slip_package'
import { CloudFunction, CloudFunctionList, createCloudFunctionStore } from './cloud_function'
import { ApiKey, ApiKeyList, createApiKeyStore } from './api_key'
import { Quotation, QuotationList, createQuotationStore } from './quotation'
import { ExaminationProperty, ExaminationPropertyList, createExaminationPropertyStore } from './examination_property'
import { AirlineOrderGrouping, AirlineOrderGroupingList, createAirlineOrderGroupingStore } from './airline_order_grouping'
import { PriceLayer, PriceLayerList, createPriceLayerStore } from './price_layer'
import { AirlineOrderGroup, AirlineOrderGroupList, createAirlineOrderGroupStore } from './airline_order_group'
import { AirlineOrder, AirlineOrderList, createAirlineOrderStore } from './airline_order'
import { AirlinePricing, AirlinePricingList, createAirlinePricingStore } from './airline_pricing'
import { PrintingOrder, PrintingOrderList, createPrintingOrderStore } from './printing_order'
import { PackingSlip, PackingSlipList, createPackingSlipStore } from './packing_slip'
import { Email, EmailList, createEmailStore } from './email'
import { BatchExport, BatchExportList, createBatchExportStore } from './batch_export'
import { ExaminationTask, ExaminationTaskList, createExaminationTaskStore } from './examination_task'
import { Attribute, AttributeList, createAttributeStore } from './attribute'
import { AttributeOptionSet, AttributeOptionSetList, createAttributeOptionSetStore } from './attribute_option_set'
import { AttributeOption, AttributeOptionList, createAttributeOptionStore } from './attribute_option'
import { type JSONData } from "../base_classes/dto.d";
import {
  type EntityName,
  type EntityClass,
  type EntityTypeName,
} from '../types/types.gen.d';


const entityBuilder: Record<EntityName, ((pathVars?: any, data?: JSONData) => EntityClass)> = {
  user_profile: (
    pathVars?: string 
             | string[] 
             | {id: string }, 
    data?: JSONData
  ) => new UserProfile(pathVars, data),
  user: (
    pathVars?: string 
             | string[] 
             | {id: string }, 
    data?: JSONData
  ) => new User(pathVars, data),
  portal_event: (
    pathVars?: string 
             | string[] 
             | {id: string }, 
    data?: JSONData
  ) => new PortalEvent(pathVars, data),
  media: (
    pathVars?: string 
             | string[] 
             | {id: string }, 
    data?: JSONData
  ) => new Media(pathVars, data),
  airline_settings: (
    pathVars?: string 
             | string[] 
             | {id: string }, 
    data?: JSONData
  ) => new AirlineSettings(pathVars, data),
  delivery_entry: (
    pathVars?: string 
             | string[] 
             | {id: string }, 
    data?: JSONData
  ) => new DeliveryEntry(pathVars, data),
  printing_order_supplier: (
    pathVars?: string 
             | string[] 
             | {id: string }, 
    data?: JSONData
  ) => new PrintingOrderSupplier(pathVars, data),
  product: (
    pathVars?: string 
             | string[] 
             | {id: string }, 
    data?: JSONData
  ) => new Product(pathVars, data),
  examination_settings: (
    pathVars?: string 
             | string[] 
             | {id: string }, 
    data?: JSONData
  ) => new ExaminationSettings(pathVars, data),
  price_configuration: (
    pathVars?: string 
             | string[] 
             | {id: string }, 
    data?: JSONData
  ) => new PriceConfiguration(pathVars, data),
  printing_order_item: (
    pathVars?: string 
             | string[] 
             | {id: string }, 
    data?: JSONData
  ) => new PrintingOrderItem(pathVars, data),
  packing_slip_package: (
    pathVars?: string 
             | string[] 
             | {id: string }, 
    data?: JSONData
  ) => new PackingSlipPackage(pathVars, data),
  cloud_function: (
    pathVars?: string 
             | string[] 
             | {id: string }, 
    data?: JSONData
  ) => new CloudFunction(pathVars, data),
  api_key: (
    pathVars?: string 
             | string[] 
             | {id: string }, 
    data?: JSONData
  ) => new ApiKey(pathVars, data),
  quotation: (
    pathVars?: string 
             | string[] 
             | {id: string }, 
    data?: JSONData
  ) => new Quotation(pathVars, data),
  examination_property: (
    pathVars?: string 
             | string[] 
             | {id: string }, 
    data?: JSONData
  ) => new ExaminationProperty(pathVars, data),
  airline_order_grouping: (
    pathVars?: string 
             | string[] 
             | {id: string }, 
    data?: JSONData
  ) => new AirlineOrderGrouping(pathVars, data),
  price_layer: (
    pathVars?: string 
             | string[] 
             | {id: string }, 
    data?: JSONData
  ) => new PriceLayer(pathVars, data),
  airline_order_group: (
    pathVars?: string 
             | string[] 
             | {id: string }, 
    data?: JSONData
  ) => new AirlineOrderGroup(pathVars, data),
  airline_order: (
    pathVars?: string 
             | string[] 
             | {id: string }, 
    data?: JSONData
  ) => new AirlineOrder(pathVars, data),
  airline_pricing: (
    pathVars?: string 
             | string[] 
             | {id: string }, 
    data?: JSONData
  ) => new AirlinePricing(pathVars, data),
  printing_order: (
    pathVars?: string 
             | string[] 
             | {id: string }, 
    data?: JSONData
  ) => new PrintingOrder(pathVars, data),
  packing_slip: (
    pathVars?: string 
             | string[] 
             | {id: string }, 
    data?: JSONData
  ) => new PackingSlip(pathVars, data),
  email: (
    pathVars?: string 
             | string[] 
             | {id: string }, 
    data?: JSONData
  ) => new Email(pathVars, data),
  batch_export: (
    pathVars?: string 
             | string[] 
             | {id: string }, 
    data?: JSONData
  ) => new BatchExport(pathVars, data),
  examination_task: (
    pathVars?: string 
             | string[] 
             | {id: string }, 
    data?: JSONData
  ) => new ExaminationTask(pathVars, data),
  attribute: (
    pathVars?: string 
             | string[] 
             | {id: string }, 
    data?: JSONData
  ) => new Attribute(pathVars, data),
  attribute_option_set: (
    pathVars?: string 
             | string[] 
             | {id: string }, 
    data?: JSONData
  ) => new AttributeOptionSet(pathVars, data),
  attribute_option: (
    pathVars?: string 
             | string[] 
             | {id: string }, 
    data?: JSONData
  ) => new AttributeOption(pathVars, data),
}

const entityCollectionNameToEntityName: Record<string, EntityName> = Object.freeze({
  UserProfile: 'user_profile',
  User: 'user',
  PortalEvent: 'portal_event',
  Media: 'media',
  AirlineSettings: 'airline_settings',
  DeliveryEntry: 'delivery_entry',
  PrintingOrderSupplier: 'printing_order_supplier',
  Product: 'product',
  ExaminationSettings: 'examination_settings',
  PriceConfiguration: 'price_configuration',
  PrintingOrderItem: 'printing_order_item',
  PackingSlipPackage: 'packing_slip_package',
  CloudFunction: 'cloud_function',
  ApiKey: 'api_key',
  Quotation: 'quotation',
  ExaminationProperty: 'examination_property',
  AirlineOrderGrouping: 'airline_order_grouping',
  PriceLayer: 'price_layer',
  AirlineOrderGroup: 'airline_order_group',
  AirlineOrder: 'airline_order',
  AirlinePricing: 'airline_pricing',
  PrintingOrder: 'printing_order',
  PackingSlip: 'packing_slip',
  Email: 'email',
  BatchExport: 'batch_export',
  ExaminationTask: 'examination_task',
  Attribute: 'attribute',
  AttributeOptionSet: 'attribute_option_set',
  AttributeOption: 'attribute_option',
})

const entityNameToTypeName: Record<EntityName, EntityTypeName> = Object.freeze({
  user_profile: 'UserProfile',
  user: 'User',
  portal_event: 'PortalEvent',
  media: 'Media',
  airline_settings: 'AirlineSettings',
  delivery_entry: 'DeliveryEntry',
  printing_order_supplier: 'PrintingOrderSupplier',
  product: 'Product',
  examination_settings: 'ExaminationSettings',
  price_configuration: 'PriceConfiguration',
  printing_order_item: 'PrintingOrderItem',
  packing_slip_package: 'PackingSlipPackage',
  cloud_function: 'CloudFunction',
  api_key: 'ApiKey',
  quotation: 'Quotation',
  examination_property: 'ExaminationProperty',
  airline_order_grouping: 'AirlineOrderGrouping',
  price_layer: 'PriceLayer',
  airline_order_group: 'AirlineOrderGroup',
  airline_order: 'AirlineOrder',
  airline_pricing: 'AirlinePricing',
  printing_order: 'PrintingOrder',
  packing_slip: 'PackingSlip',
  email: 'Email',
  batch_export: 'BatchExport',
  examination_task: 'ExaminationTask',
  attribute: 'Attribute',
  attribute_option_set: 'AttributeOptionSet',
  attribute_option: 'AttributeOption',
})

const entityTypeName: Record<EntityTypeName, EntityTypeName> = Object.freeze({
  UserProfile: 'UserProfile',
  User: 'User',
  PortalEvent: 'PortalEvent',
  Media: 'Media',
  AirlineSettings: 'AirlineSettings',
  DeliveryEntry: 'DeliveryEntry',
  PrintingOrderSupplier: 'PrintingOrderSupplier',
  Product: 'Product',
  ExaminationSettings: 'ExaminationSettings',
  PriceConfiguration: 'PriceConfiguration',
  PrintingOrderItem: 'PrintingOrderItem',
  PackingSlipPackage: 'PackingSlipPackage',
  CloudFunction: 'CloudFunction',
  ApiKey: 'ApiKey',
  Quotation: 'Quotation',
  ExaminationProperty: 'ExaminationProperty',
  AirlineOrderGrouping: 'AirlineOrderGrouping',
  PriceLayer: 'PriceLayer',
  AirlineOrderGroup: 'AirlineOrderGroup',
  AirlineOrder: 'AirlineOrder',
  AirlinePricing: 'AirlinePricing',
  PrintingOrder: 'PrintingOrder',
  PackingSlip: 'PackingSlip',
  Email: 'Email',
  BatchExport: 'BatchExport',
  ExaminationTask: 'ExaminationTask',
  Attribute: 'Attribute',
  AttributeOptionSet: 'AttributeOptionSet',
  AttributeOption: 'AttributeOption',
})

export {
  entityBuilder,
  entityCollectionNameToEntityName,
  entityNameToTypeName,
  entityTypeName,
  UserProfile,
  UserProfileList,
  createUserProfileStore,
  User,
  UserList,
  createUserStore,
  PortalEvent,
  PortalEventList,
  createPortalEventStore,
  Media,
  MediaList,
  createMediaStore,
  AirlineSettings,
  AirlineSettingsList,
  createAirlineSettingsStore,
  DeliveryEntry,
  DeliveryEntryList,
  createDeliveryEntryStore,
  PrintingOrderSupplier,
  PrintingOrderSupplierList,
  createPrintingOrderSupplierStore,
  Product,
  ProductList,
  createProductStore,
  ExaminationSettings,
  ExaminationSettingsList,
  createExaminationSettingsStore,
  PriceConfiguration,
  PriceConfigurationList,
  createPriceConfigurationStore,
  PrintingOrderItem,
  PrintingOrderItemList,
  createPrintingOrderItemStore,
  PackingSlipPackage,
  PackingSlipPackageList,
  createPackingSlipPackageStore,
  CloudFunction,
  CloudFunctionList,
  createCloudFunctionStore,
  ApiKey,
  ApiKeyList,
  createApiKeyStore,
  Quotation,
  QuotationList,
  createQuotationStore,
  ExaminationProperty,
  ExaminationPropertyList,
  createExaminationPropertyStore,
  AirlineOrderGrouping,
  AirlineOrderGroupingList,
  createAirlineOrderGroupingStore,
  PriceLayer,
  PriceLayerList,
  createPriceLayerStore,
  AirlineOrderGroup,
  AirlineOrderGroupList,
  createAirlineOrderGroupStore,
  AirlineOrder,
  AirlineOrderList,
  createAirlineOrderStore,
  AirlinePricing,
  AirlinePricingList,
  createAirlinePricingStore,
  PrintingOrder,
  PrintingOrderList,
  createPrintingOrderStore,
  PackingSlip,
  PackingSlipList,
  createPackingSlipStore,
  Email,
  EmailList,
  createEmailStore,
  BatchExport,
  BatchExportList,
  createBatchExportStore,
  ExaminationTask,
  ExaminationTaskList,
  createExaminationTaskStore,
  Attribute,
  AttributeList,
  createAttributeStore,
  AttributeOptionSet,
  AttributeOptionSetList,
  createAttributeOptionSetStore,
  AttributeOption,
  AttributeOptionList,
  createAttributeOptionStore,
}
