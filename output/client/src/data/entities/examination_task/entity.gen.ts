import { type JSONData } from '../../base_classes/dto.d'
import { ExaminationTaskDto } from './dto.gen'
import { type IExaminationTask } from './entity.d'
import ExaminationTaskParsers from './parsers'
import ExaminationTaskValidators from './custom_validators'

export class ExaminationTask extends ExaminationTaskDto implements IExaminationTask {
  constructor(pathVars?: string | string[] | {
    id?: string,
  }, data?: JSONData) {
    super(data, ExaminationTaskParsers, ExaminationTaskValidators)

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
  ): Promise<ExaminationTask> {
    return (await super.get(pathVars)) as ExaminationTask
  }

  /// ///////////////////////////////////////
  /// /////////// STATIC METHODS ////////////
  /// ///////////////////////////////////////

  static override FromJSON(data?: JSONData): ExaminationTask {
    return new ExaminationTask(undefined, data)
  }

  static override FromJSONArray(data: JSONData[]): ExaminationTask[] {
    return data.map((item) => ExaminationTask.FromJSON(item))
  }

  static FromDto(dto: ExaminationTaskDto): ExaminationTask {
    return new ExaminationTask(undefined, dto.toJSON())
  }
}

