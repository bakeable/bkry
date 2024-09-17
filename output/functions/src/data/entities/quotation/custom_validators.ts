import { type Validator } from '../../base_classes/dto_validator.d'
import generatedValidators from './validators.gen'

const validators: Validator[] = [
  ...generatedValidators,
]

export default validators
