import { type JSONData } from '../../base_classes/dto.d'
import { MediaDto } from './dto.gen'
import { type IMedia } from './entity.d'
import MediaParsers from './parsers'
import MediaValidators from './custom_validators'

export class Media extends MediaDto implements IMedia {
  constructor(pathVars?: string | string[] | {
    id?: string,
  }, data?: JSONData) {
    super(data, MediaParsers, MediaValidators)

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
  ): Promise<Media> {
    return (await super.get(pathVars)) as Media
  }

  /// ///////////////////////////////////////
  /// /////////// STATIC METHODS ////////////
  /// ///////////////////////////////////////

  static override FromJSON(data?: JSONData): Media {
    return new Media(undefined, data)
  }

  static override FromJSONArray(data: JSONData[]): Media[] {
    return data.map((item) => Media.FromJSON(item))
  }

  static FromDto(dto: MediaDto): Media {
    return new Media(undefined, dto.toJSON())
  }
}

