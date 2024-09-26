import { type JSONData } from '../../base_classes/dto.d'
import { BatchExportDto } from './dto.gen'
import { type IBatchExport } from './entity.d'
import BatchExportParsers from './parsers'
import BatchExportValidators from './custom_validators'

export class BatchExport extends BatchExportDto implements IBatchExport {
  constructor(pathVars?: string | string[] | {
    id?: string,
  }, data?: JSONData) {
    super(data, BatchExportParsers, BatchExportValidators)

    if (data !== undefined) {
      // Set data
      this.set(data)
    }

    this.buildPathObject(pathVars);
  }
  /// ///////////////////////////////////////
  /// //////// PRIVATE VARIABLES ////////////
  /// ///////////////////////////////////////

  /// ///////////////////////////////////////
  /// ///////// PUBLIC VARIABLES ////////////
  /// ///////////////////////////////////////

  /// ///////////////////////////////////////
  /// ////////////// METHODS ////////////////
  /// ///////////////////////////////////////

  public async get(
    pathVars?: string | string[] | {
    id?: string,
  } | undefined
  ): Promise<BatchExport> {
    return (await super.get(pathVars)) as BatchExport
  }

  /// ///////////////////////////////////////
  /// /////////// STATIC METHODS ////////////
  /// ///////////////////////////////////////

  static override FromJSON(data?: JSONData): BatchExport {
    return new BatchExport(undefined, data)
  }

  static override FromJSONArray(data: JSONData[]): BatchExport[] {
    return data.map((item) => BatchExport.FromJSON(item))
  }

  static FromDto(dto: BatchExportDto): BatchExport {
    return new BatchExport(undefined, dto.toJSON())
  }
}

