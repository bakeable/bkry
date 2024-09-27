
import { AirlineOrder, AirlineOrderList, createAirlineOrderStore } from './airline_order'
import { AirlineOrderGroup, AirlineOrderGroupList, createAirlineOrderGroupStore } from './airline_order_group'
import { AirlineOrderGrouping, AirlineOrderGroupingList, createAirlineOrderGroupingStore } from './airline_order_grouping'
import { AirlinePricing, AirlinePricingList, createAirlinePricingStore } from './airline_pricing'
import { AirlineSettings, AirlineSettingsList, createAirlineSettingsStore } from './airline_settings'
import { ApiKey, ApiKeyList, createApiKeyStore } from './api_key'
import { Attribute, AttributeList, createAttributeStore } from './attribute'
import { AttributeOption, AttributeOptionList, createAttributeOptionStore } from './attribute_option'
import { AttributeOptionSet, AttributeOptionSetList, createAttributeOptionSetStore } from './attribute_option_set'
import { BatchExport, BatchExportList, createBatchExportStore } from './batch_export'
import { CloudFunction, CloudFunctionList, createCloudFunctionStore } from './cloud_function'
import { DeliveryEntry, DeliveryEntryList, createDeliveryEntryStore } from './delivery_entry'
import { Email, EmailList, createEmailStore } from './email'
import { ExaminationProperty, ExaminationPropertyList, createExaminationPropertyStore } from './examination_property'
import { ExaminationSettings, ExaminationSettingsList, createExaminationSettingsStore } from './examination_settings'
import { ExaminationTask, ExaminationTaskList, createExaminationTaskStore } from './examination_task'
import { Media, MediaList, createMediaStore } from './media'
import { PackingSlip, PackingSlipList, createPackingSlipStore } from './packing_slip'
import { PackingSlipPackage, PackingSlipPackageList, createPackingSlipPackageStore } from './packing_slip_package'
import { PortalEvent, PortalEventList, createPortalEventStore } from './portal_event'
import { PriceConfiguration, PriceConfigurationList, createPriceConfigurationStore } from './price_configuration'
import { PriceLayer, PriceLayerList, createPriceLayerStore } from './price_layer'
import { PrintingOrder, PrintingOrderList, createPrintingOrderStore } from './printing_order'
import { PrintingOrderItem, PrintingOrderItemList, createPrintingOrderItemStore } from './printing_order_item'
import { PrintingOrderSupplier, PrintingOrderSupplierList, createPrintingOrderSupplierStore } from './printing_order_supplier'
import { Product, ProductList, createProductStore } from './product'
import { Quotation, QuotationList, createQuotationStore } from './quotation'
import { User, UserList, createUserStore } from './user'
import { UserProfile, UserProfileList, createUserProfileStore } from './user_profile'
import { type JSONData } from "../base_classes/dto.d";
import {
  type EntityName,
  type EntityClass,
  type EntityTypeName,
} from '../types/types.gen.d';


const entityBuilder: Record<EntityName, ((pathVars?: any, data?: JSONData) => EntityClass)> = {
  airline_order: (
    pathVars?: string 
             | string[] 
             | {id: string }, 
    data?: JSONData
  ) => new AirlineOrder(pathVars, data),
  airline_order_group: (
    pathVars?: string 
             | string[] 
             | {id: string }, 
    data?: JSONData
  ) => new AirlineOrderGroup(pathVars, data),
  airline_order_grouping: (
    pathVars?: string 
             | string[] 
             | {id: string }, 
    data?: JSONData
  ) => new AirlineOrderGrouping(pathVars, data),
  airline_pricing: (
    pathVars?: string 
             | string[] 
             | {id: string }, 
    data?: JSONData
  ) => new AirlinePricing(pathVars, data),
  airline_settings: (
    pathVars?: string 
             | string[] 
             | {id: string }, 
    data?: JSONData
  ) => new AirlineSettings(pathVars, data),
  api_key: (
    pathVars?: string 
             | string[] 
             | {id: string }, 
    data?: JSONData
  ) => new ApiKey(pathVars, data),
  attribute: (
    pathVars?: string 
             | string[] 
             | {id: string }, 
    data?: JSONData
  ) => new Attribute(pathVars, data),
  attribute_option: (
    pathVars?: string 
             | string[] 
             | {id: string }, 
    data?: JSONData
  ) => new AttributeOption(pathVars, data),
  attribute_option_set: (
    pathVars?: string 
             | string[] 
             | {id: string }, 
    data?: JSONData
  ) => new AttributeOptionSet(pathVars, data),
  batch_export: (
    pathVars?: string 
             | string[] 
             | {id: string }, 
    data?: JSONData
  ) => new BatchExport(pathVars, data),
  cloud_function: (
    pathVars?: string 
             | string[] 
             | {id: string }, 
    data?: JSONData
  ) => new CloudFunction(pathVars, data),
  delivery_entry: (
    pathVars?: string 
             | string[] 
             | {id: string }, 
    data?: JSONData
  ) => new DeliveryEntry(pathVars, data),
  email: (
    pathVars?: string 
             | string[] 
             | {id: string }, 
    data?: JSONData
  ) => new Email(pathVars, data),
  examination_property: (
    pathVars?: string 
             | string[] 
             | {id: string }, 
    data?: JSONData
  ) => new ExaminationProperty(pathVars, data),
  examination_settings: (
    pathVars?: string 
             | string[] 
             | {id: string }, 
    data?: JSONData
  ) => new ExaminationSettings(pathVars, data),
  examination_task: (
    pathVars?: string 
             | string[] 
             | {id: string }, 
    data?: JSONData
  ) => new ExaminationTask(pathVars, data),
  media: (
    pathVars?: string 
             | string[] 
             | {id: string }, 
    data?: JSONData
  ) => new Media(pathVars, data),
  packing_slip: (
    pathVars?: string 
             | string[] 
             | {id: string }, 
    data?: JSONData
  ) => new PackingSlip(pathVars, data),
  packing_slip_package: (
    pathVars?: string 
             | string[] 
             | {id: string }, 
    data?: JSONData
  ) => new PackingSlipPackage(pathVars, data),
  portal_event: (
    pathVars?: string 
             | string[] 
             | {id: string }, 
    data?: JSONData
  ) => new PortalEvent(pathVars, data),
  price_configuration: (
    pathVars?: string 
             | string[] 
             | {id: string }, 
    data?: JSONData
  ) => new PriceConfiguration(pathVars, data),
  price_layer: (
    pathVars?: string 
             | string[] 
             | {id: string }, 
    data?: JSONData
  ) => new PriceLayer(pathVars, data),
  printing_order: (
    pathVars?: string 
             | string[] 
             | {id: string }, 
    data?: JSONData
  ) => new PrintingOrder(pathVars, data),
  printing_order_item: (
    pathVars?: string 
             | string[] 
             | {id: string }, 
    data?: JSONData
  ) => new PrintingOrderItem(pathVars, data),
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
  quotation: (
    pathVars?: string 
             | string[] 
             | {id: string }, 
    data?: JSONData
  ) => new Quotation(pathVars, data),
  user: (
    pathVars?: string 
             | string[] 
             | {id: string }, 
    data?: JSONData
  ) => new User(pathVars, data),
  user_profile: (
    pathVars?: string 
             | string[] 
             | {id: string }, 
    data?: JSONData
  ) => new UserProfile(pathVars, data),
}

const entityCollectionNameToEntityName: Record<string, EntityName> = Object.freeze({
  AirlineOrder: 'airline_order',
  AirlineOrderGroup: 'airline_order_group',
  AirlineOrderGrouping: 'airline_order_grouping',
  AirlinePricing: 'airline_pricing',
  AirlineSettings: 'airline_settings',
  ApiKey: 'api_key',
  Attribute: 'attribute',
  AttributeOption: 'attribute_option',
  AttributeOptionSet: 'attribute_option_set',
  BatchExport: 'batch_export',
  CloudFunction: 'cloud_function',
  DeliveryEntry: 'delivery_entry',
  Email: 'email',
  ExaminationProperty: 'examination_property',
  ExaminationSettings: 'examination_settings',
  ExaminationTask: 'examination_task',
  Media: 'media',
  PackingSlip: 'packing_slip',
  PackingSlipPackage: 'packing_slip_package',
  PortalEvent: 'portal_event',
  PriceConfiguration: 'price_configuration',
  PriceLayer: 'price_layer',
  PrintingOrder: 'printing_order',
  PrintingOrderItem: 'printing_order_item',
  PrintingOrderSupplier: 'printing_order_supplier',
  Product: 'product',
  Quotation: 'quotation',
  User: 'user',
  UserProfile: 'user_profile',
})

const entityNameToTypeName: Record<EntityName, EntityTypeName> = Object.freeze({
  airline_order: 'AirlineOrder',
  airline_order_group: 'AirlineOrderGroup',
  airline_order_grouping: 'AirlineOrderGrouping',
  airline_pricing: 'AirlinePricing',
  airline_settings: 'AirlineSettings',
  api_key: 'ApiKey',
  attribute: 'Attribute',
  attribute_option: 'AttributeOption',
  attribute_option_set: 'AttributeOptionSet',
  batch_export: 'BatchExport',
  cloud_function: 'CloudFunction',
  delivery_entry: 'DeliveryEntry',
  email: 'Email',
  examination_property: 'ExaminationProperty',
  examination_settings: 'ExaminationSettings',
  examination_task: 'ExaminationTask',
  media: 'Media',
  packing_slip: 'PackingSlip',
  packing_slip_package: 'PackingSlipPackage',
  portal_event: 'PortalEvent',
  price_configuration: 'PriceConfiguration',
  price_layer: 'PriceLayer',
  printing_order: 'PrintingOrder',
  printing_order_item: 'PrintingOrderItem',
  printing_order_supplier: 'PrintingOrderSupplier',
  product: 'Product',
  quotation: 'Quotation',
  user: 'User',
  user_profile: 'UserProfile',
})

const entityTypeName: Record<EntityTypeName, EntityTypeName> = Object.freeze({
  AirlineOrder: 'AirlineOrder',
  AirlineOrderGroup: 'AirlineOrderGroup',
  AirlineOrderGrouping: 'AirlineOrderGrouping',
  AirlinePricing: 'AirlinePricing',
  AirlineSettings: 'AirlineSettings',
  ApiKey: 'ApiKey',
  Attribute: 'Attribute',
  AttributeOption: 'AttributeOption',
  AttributeOptionSet: 'AttributeOptionSet',
  BatchExport: 'BatchExport',
  CloudFunction: 'CloudFunction',
  DeliveryEntry: 'DeliveryEntry',
  Email: 'Email',
  ExaminationProperty: 'ExaminationProperty',
  ExaminationSettings: 'ExaminationSettings',
  ExaminationTask: 'ExaminationTask',
  Media: 'Media',
  PackingSlip: 'PackingSlip',
  PackingSlipPackage: 'PackingSlipPackage',
  PortalEvent: 'PortalEvent',
  PriceConfiguration: 'PriceConfiguration',
  PriceLayer: 'PriceLayer',
  PrintingOrder: 'PrintingOrder',
  PrintingOrderItem: 'PrintingOrderItem',
  PrintingOrderSupplier: 'PrintingOrderSupplier',
  Product: 'Product',
  Quotation: 'Quotation',
  User: 'User',
  UserProfile: 'UserProfile',
})

export {
  entityBuilder,
  entityCollectionNameToEntityName,
  entityNameToTypeName,
  entityTypeName,
  AirlineOrder,
  AirlineOrderList,
  createAirlineOrderStore,
  AirlineOrderGroup,
  AirlineOrderGroupList,
  createAirlineOrderGroupStore,
  AirlineOrderGrouping,
  AirlineOrderGroupingList,
  createAirlineOrderGroupingStore,
  AirlinePricing,
  AirlinePricingList,
  createAirlinePricingStore,
  AirlineSettings,
  AirlineSettingsList,
  createAirlineSettingsStore,
  ApiKey,
  ApiKeyList,
  createApiKeyStore,
  Attribute,
  AttributeList,
  createAttributeStore,
  AttributeOption,
  AttributeOptionList,
  createAttributeOptionStore,
  AttributeOptionSet,
  AttributeOptionSetList,
  createAttributeOptionSetStore,
  BatchExport,
  BatchExportList,
  createBatchExportStore,
  CloudFunction,
  CloudFunctionList,
  createCloudFunctionStore,
  DeliveryEntry,
  DeliveryEntryList,
  createDeliveryEntryStore,
  Email,
  EmailList,
  createEmailStore,
  ExaminationProperty,
  ExaminationPropertyList,
  createExaminationPropertyStore,
  ExaminationSettings,
  ExaminationSettingsList,
  createExaminationSettingsStore,
  ExaminationTask,
  ExaminationTaskList,
  createExaminationTaskStore,
  Media,
  MediaList,
  createMediaStore,
  PackingSlip,
  PackingSlipList,
  createPackingSlipStore,
  PackingSlipPackage,
  PackingSlipPackageList,
  createPackingSlipPackageStore,
  PortalEvent,
  PortalEventList,
  createPortalEventStore,
  PriceConfiguration,
  PriceConfigurationList,
  createPriceConfigurationStore,
  PriceLayer,
  PriceLayerList,
  createPriceLayerStore,
  PrintingOrder,
  PrintingOrderList,
  createPrintingOrderStore,
  PrintingOrderItem,
  PrintingOrderItemList,
  createPrintingOrderItemStore,
  PrintingOrderSupplier,
  PrintingOrderSupplierList,
  createPrintingOrderSupplierStore,
  Product,
  ProductList,
  createProductStore,
  Quotation,
  QuotationList,
  createQuotationStore,
  User,
  UserList,
  createUserStore,
  UserProfile,
  UserProfileList,
  createUserProfileStore,
}
