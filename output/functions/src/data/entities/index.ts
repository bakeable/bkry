
import { UserProfile, UserProfileList } from './user_profile'
import { User, UserList } from './user'
import { PortalEvent, PortalEventList } from './portal_event'
import { Media, MediaList } from './media'
import { AirlineSettings, AirlineSettingsList } from './airline_settings'
import { DeliveryEntry, DeliveryEntryList } from './delivery_entry'
import { PrintingOrderSupplier, PrintingOrderSupplierList } from './printing_order_supplier'
import { Product, ProductList } from './product'
import { ExaminationSettings, ExaminationSettingsList } from './examination_settings'
import { PriceConfiguration, PriceConfigurationList } from './price_configuration'
import { PrintingOrderItem, PrintingOrderItemList } from './printing_order_item'
import { PackingSlipPackage, PackingSlipPackageList } from './packing_slip_package'
import { CloudFunction, CloudFunctionList } from './cloud_function'
import { ApiKey, ApiKeyList } from './api_key'
import { Quotation, QuotationList } from './quotation'
import { ExaminationProperty, ExaminationPropertyList } from './examination_property'
import { AirlineOrderGrouping, AirlineOrderGroupingList } from './airline_order_grouping'
import { PriceLayer, PriceLayerList } from './price_layer'
import { AirlineOrderGroup, AirlineOrderGroupList } from './airline_order_group'
import { AirlineOrder, AirlineOrderList } from './airline_order'
import { AirlinePricing, AirlinePricingList } from './airline_pricing'
import { PrintingOrder, PrintingOrderList } from './printing_order'
import { PackingSlip, PackingSlipList } from './packing_slip'
import { Email, EmailList } from './email'
import { BatchExport, BatchExportList } from './batch_export'
import { ExaminationTask, ExaminationTaskList } from './examination_task'
import { Attribute, AttributeList } from './attribute'
import { AttributeOptionSet, AttributeOptionSetList } from './attribute_option_set'
import { AttributeOption, AttributeOptionList } from './attribute_option'
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

const entityTypeName: Record<string, EntityTypeName> = Object.freeze({
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
  User,
  UserList,
  PortalEvent,
  PortalEventList,
  Media,
  MediaList,
  AirlineSettings,
  AirlineSettingsList,
  DeliveryEntry,
  DeliveryEntryList,
  PrintingOrderSupplier,
  PrintingOrderSupplierList,
  Product,
  ProductList,
  ExaminationSettings,
  ExaminationSettingsList,
  PriceConfiguration,
  PriceConfigurationList,
  PrintingOrderItem,
  PrintingOrderItemList,
  PackingSlipPackage,
  PackingSlipPackageList,
  CloudFunction,
  CloudFunctionList,
  ApiKey,
  ApiKeyList,
  Quotation,
  QuotationList,
  ExaminationProperty,
  ExaminationPropertyList,
  AirlineOrderGrouping,
  AirlineOrderGroupingList,
  PriceLayer,
  PriceLayerList,
  AirlineOrderGroup,
  AirlineOrderGroupList,
  AirlineOrder,
  AirlineOrderList,
  AirlinePricing,
  AirlinePricingList,
  PrintingOrder,
  PrintingOrderList,
  PackingSlip,
  PackingSlipList,
  Email,
  EmailList,
  BatchExport,
  BatchExportList,
  ExaminationTask,
  ExaminationTaskList,
  Attribute,
  AttributeList,
  AttributeOptionSet,
  AttributeOptionSetList,
  AttributeOption,
  AttributeOptionList,
}
