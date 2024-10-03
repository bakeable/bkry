
import { Author, AuthorTransporter, AuthorHandler, createAuthorStore } from './author'
import { Category, CategoryTransporter, CategoryHandler, createCategoryStore } from './category'
import { GameModus, GameModusTransporter, GameModusHandler, createGameModusStore } from './game_modus'
import { Media, MediaTransporter, MediaHandler, createMediaStore } from './media'
import { ProductPackage, ProductPackageTransporter, ProductPackageHandler, createProductPackageStore } from './product_package'
import { Purchase, PurchaseTransporter, PurchaseHandler, createPurchaseStore } from './purchase'
import { QuestionBundle, QuestionBundleTransporter, QuestionBundleHandler, createQuestionBundleStore } from './question_bundle'
import { QuestionContext, QuestionContextTransporter, QuestionContextHandler, createQuestionContextStore } from './question_context'
import { Tag, TagTransporter, TagHandler, createTagStore } from './tag'
import { User, UserTransporter, UserHandler, createUserStore } from './user'
import { CategoryLocalisation, CategoryLocalisationTransporter, CategoryLocalisationHandler, createCategoryLocalisationStore } from './category_localisation'
import { GameModusLocalisation, GameModusLocalisationTransporter, GameModusLocalisationHandler, createGameModusLocalisationStore } from './game_modus_localisation'
import { ProductPackageLocalisation, ProductPackageLocalisationTransporter, ProductPackageLocalisationHandler, createProductPackageLocalisationStore } from './product_package_localisation'
import { QuestionBundleLocalisation, QuestionBundleLocalisationTransporter, QuestionBundleLocalisationHandler, createQuestionBundleLocalisationStore } from './question_bundle_localisation'
import { QuestionContextLocalisation, QuestionContextLocalisationTransporter, QuestionContextLocalisationHandler, createQuestionContextLocalisationStore } from './question_context_localisation'
import { TagLocalisation, TagLocalisationTransporter, TagLocalisationHandler, createTagLocalisationStore } from './tag_localisation'
import { type JSONData } from "../base_classes/dto.d";
import {
  type EntityName,
  type EntityClass,
  type EntityTypeName,
} from '../types/types.gen.d';

const entityCollectionNameToEntityName: Record<string, EntityName> = Object.freeze({
  Author: 'author',
  Category: 'category',
  GameModus: 'game_modus',
  Media: 'media',
  ProductPackage: 'product_package',
  Purchase: 'purchase',
  QuestionBundle: 'question_bundle',
  QuestionContext: 'question_context',
  Tag: 'tag',
  User: 'user',
  CategoryLocalisation: 'category_localisation',
  GameModusLocalisation: 'game_modus_localisation',
  ProductPackageLocalisation: 'product_package_localisation',
  QuestionBundleLocalisation: 'question_bundle_localisation',
  QuestionContextLocalisation: 'question_context_localisation',
  TagLocalisation: 'tag_localisation',
})

const entityNameToTypeName: Record<EntityName, EntityTypeName> = Object.freeze({
  author: 'Author',
  category: 'Category',
  game_modus: 'GameModus',
  media: 'Media',
  product_package: 'ProductPackage',
  purchase: 'Purchase',
  question_bundle: 'QuestionBundle',
  question_context: 'QuestionContext',
  tag: 'Tag',
  user: 'User',
  category_localisation: 'CategoryLocalisation',
  game_modus_localisation: 'GameModusLocalisation',
  product_package_localisation: 'ProductPackageLocalisation',
  question_bundle_localisation: 'QuestionBundleLocalisation',
  question_context_localisation: 'QuestionContextLocalisation',
  tag_localisation: 'TagLocalisation',
})

const entityTypeName: Record<EntityTypeName, EntityTypeName> = Object.freeze({
  Author: 'Author',
  Category: 'Category',
  GameModus: 'GameModus',
  Media: 'Media',
  ProductPackage: 'ProductPackage',
  Purchase: 'Purchase',
  QuestionBundle: 'QuestionBundle',
  QuestionContext: 'QuestionContext',
  Tag: 'Tag',
  User: 'User',
  CategoryLocalisation: 'CategoryLocalisation',
  GameModusLocalisation: 'GameModusLocalisation',
  ProductPackageLocalisation: 'ProductPackageLocalisation',
  QuestionBundleLocalisation: 'QuestionBundleLocalisation',
  QuestionContextLocalisation: 'QuestionContextLocalisation',
  TagLocalisation: 'TagLocalisation',
})

export {
  entityCollectionNameToEntityName,
  entityNameToTypeName,
  entityTypeName,
  Author,
  AuthorTransporter,
  AuthorHandler,
  createAuthorStore,
  Category,
  CategoryTransporter,
  CategoryHandler,
  createCategoryStore,
  GameModus,
  GameModusTransporter,
  GameModusHandler,
  createGameModusStore,
  Media,
  MediaTransporter,
  MediaHandler,
  createMediaStore,
  ProductPackage,
  ProductPackageTransporter,
  ProductPackageHandler,
  createProductPackageStore,
  Purchase,
  PurchaseTransporter,
  PurchaseHandler,
  createPurchaseStore,
  QuestionBundle,
  QuestionBundleTransporter,
  QuestionBundleHandler,
  createQuestionBundleStore,
  QuestionContext,
  QuestionContextTransporter,
  QuestionContextHandler,
  createQuestionContextStore,
  Tag,
  TagTransporter,
  TagHandler,
  createTagStore,
  User,
  UserTransporter,
  UserHandler,
  createUserStore,
  CategoryLocalisation,
  CategoryLocalisationTransporter,
  CategoryLocalisationHandler,
  createCategoryLocalisationStore,
  GameModusLocalisation,
  GameModusLocalisationTransporter,
  GameModusLocalisationHandler,
  createGameModusLocalisationStore,
  ProductPackageLocalisation,
  ProductPackageLocalisationTransporter,
  ProductPackageLocalisationHandler,
  createProductPackageLocalisationStore,
  QuestionBundleLocalisation,
  QuestionBundleLocalisationTransporter,
  QuestionBundleLocalisationHandler,
  createQuestionBundleLocalisationStore,
  QuestionContextLocalisation,
  QuestionContextLocalisationTransporter,
  QuestionContextLocalisationHandler,
  createQuestionContextLocalisationStore,
  TagLocalisation,
  TagLocalisationTransporter,
  TagLocalisationHandler,
  createTagLocalisationStore,
}
