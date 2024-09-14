import { type JSONData } from '../../base_classes/dto.d'
import type IDto from '../../base_classes/dto.d'
import { type Media } from '../media'

/// ///////////////////////////////////////////////////////////////////////////////////////////////////
/// WARNING! THIS FILE IS AUTOMATICALLY GENERATED, ANY CHANGES WILL BE LOST ON NEXT GENERATION
/// ///////////////////////////////////////////////////////////////////////////////////////////////////


// Types
export type PrintingOrderCorrespondenceTemplateKey = "CUSTOMER_APPROVE_SAMPLE" | "CUSTOMER_PRINTING_ORDER_RECEIVED" | "CUSTOMER_TRACK_AND_TRACE" | "SUPPLIER_PRINTING_ORDER_RECEIVED" | "SUPPLIER_SAMPLE_ACCEPTED" | "SUPPLIER_SAMPLE_REJECTED"
export type PrintingOrderStatus = "new" | "waiting_for_sample" | "waiting_for_approval" | "sample_rejected" | "in_production" | "shipped"

// Interfaces
export interface ICorrespondence {
    // ID of the correspondence
      id: string
    // Key of the correspondence
      key: PrintingOrderCorrespondenceTemplateKey
    // Flag indicating if the correspondence was sent
      sent: boolean
    // Timestamp when the correspondence was sent
      sentAt: number
    // The index of the status log change that triggered the correspondence
      statusLogIndex: number
    // Timestamp of the correspondence
      timestamp: number
}
export interface ICustomerAddress {
    // City of the address
      city: string
    // Country of the address
      country: string
    // Address line
      line: string
    // Postal code of the address
      postalCode: string
}
export interface IFiles {
    // The artwork file
      artwork: Media
    // The media file
      media: Media
    // The order confirmation file
      orderConfirmation: Media
}
export interface IProductOptions {
    // ID of the product option
      id: string
    // Key of the product option
      key: string
    // Value of the product option
      value: string
}
export interface IStatusLog {
    // Previous status of the printing order
      prevStatus: PrintingOrderStatus
    // Status of the printing order
      status: PrintingOrderStatus
    // Timestamp of the status change
      timestamp: number
}

// Data Transfer Object Interface
export interface IPrintingOrderDto extends IDto {
  /// ///////////////////////////////////////
  /// //////// PRIVATE VARIABLES ////////////
  /// ///////////////////////////////////////
  /**
   * The path to the corresponding API endpoint with variables marked as {?}
   */
  _path: string

  /// ///////////////////////////////////////
  /// ///////// PUBLIC VARIABLES ////////////
  /// ///////////////////////////////////////

    // Flag indicating if the printing order has been approved
    approved: boolean

    // The timestamp of the approved
    approvedTimestamp: number

    // Flag indicating whether the artwork is approved
    artworkApproved: boolean

    // Flag indicating whether the artwork is a URL or a file
    artworkIsUrl: boolean

    // Flag indicating whether to automatically send correspondence
    autoCorrespondence: boolean

    // The name of the company associated with the printing order
    companyName: string

    // The correspondence tracking
    correspondence: ICorrespondence[]

    // The address of the customer
    customerAddress: ICustomerAddress

    // Email of the customer
    customerEmail: string

    // Name of the customer
    customerName: string

    // Description of the printing order
    description: string

    // The 10 digit e-mail code
    emailCode: string

    // The files associated with the printing order
    files: IFiles

    // Whether the printing order is from Airtable
    isAirtableOrder: boolean

    // The item index of the printing order
    itemIndex: number

    // Type of the printing order
    _kind: string

    // Notes of the printing order
    notes: string

    // Number of units ordered per quantity
    numberOfUnits: number

    // The order number of the printing order
    orderNumber: string

    // The status of the order in Magento
    orderStatus: string

    // The options of the product
    productOptions: IProductOptions[]

    // Quantity of the printing order
    quantity: number

    // Quantity ordered of the printing order
    quantityOrdered: number

    // Whether the sample was rejected
    rejected: boolean

    // The reason for rejection
    rejectionReason: string

    // SKU of the printing order
    sku: string

    // Status of the printing order
    status: PrintingOrderStatus

    // The ordering index of the status
    statusIndex: number

    // The log of previous statuses
    statusLog: IStatusLog[]

    // Name of the supplier contact associated with the printing order
    supplierContactName: string

    // ID of the supplier associated with the printing order
    supplierId: string

    // Name of the supplier associated with the printing order
    supplierName: string

    // The supplier reference
    supplierReference: string

    // Tracking URL of the printing order
    trackingUrl: string

    // Whether the tracking URL is approved
    trackingUrlApproved: boolean

  /// ///////////////////////////////////////
  /// ///////// PUBLIC VARIABLES ////////////
  /// ///////////////////////////////////////

  /**
   * Set the PrintingOrder's properties
   */
  set: (data: JSONData) => void
}
