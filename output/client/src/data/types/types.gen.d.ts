import {
  type Author,
  type Category,
  type GameModus,
  type Media,
  type ProductPackage,
  type Purchase,
  type QuestionBundle,
  type QuestionContext,
  type Tag,
  type User,
  type CategoryLocalisation,
  type GameModusLocalisation,
  type ProductPackageLocalisation,
  type QuestionBundleLocalisation,
  type QuestionContextLocalisation,
  type TagLocalisation,
} from '../entities';

type EntityTypeName = 
 | 'Author'
 | 'Category'
 | 'GameModus'
 | 'Media'
 | 'ProductPackage'
 | 'Purchase'
 | 'QuestionBundle'
 | 'QuestionContext'
 | 'Tag'
 | 'User'
 | 'CategoryLocalisation'
 | 'GameModusLocalisation'
 | 'ProductPackageLocalisation'
 | 'QuestionBundleLocalisation'
 | 'QuestionContextLocalisation'
 | 'TagLocalisation'


type EntityName =  | 'author'
 | 'category'
 | 'game_modus'
 | 'media'
 | 'product_package'
 | 'purchase'
 | 'question_bundle'
 | 'question_context'
 | 'tag'
 | 'user'
 | 'category_localisation'
 | 'game_modus_localisation'
 | 'product_package_localisation'
 | 'question_bundle_localisation'
 | 'question_context_localisation'
 | 'tag_localisation'


type EntityClass =
    | Author
    | Category
    | GameModus
    | Media
    | ProductPackage
    | Purchase
    | QuestionBundle
    | QuestionContext
    | Tag
    | User
    | CategoryLocalisation
    | GameModusLocalisation
    | ProductPackageLocalisation
    | QuestionBundleLocalisation
    | QuestionContextLocalisation
    | TagLocalisation;

interface IEntityClassMap {
  Author: Author;
  Category: Category;
  GameModus: GameModus;
  Media: Media;
  ProductPackage: ProductPackage;
  Purchase: Purchase;
  QuestionBundle: QuestionBundle;
  QuestionContext: QuestionContext;
  Tag: Tag;
  User: User;
  CategoryLocalisation: CategoryLocalisation;
  GameModusLocalisation: GameModusLocalisation;
  ProductPackageLocalisation: ProductPackageLocalisation;
  QuestionBundleLocalisation: QuestionBundleLocalisation;
  QuestionContextLocalisation: QuestionContextLocalisation;
  TagLocalisation: TagLocalisation;
}