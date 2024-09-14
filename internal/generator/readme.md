# Project Code Generator

## Overview
This code generator automates the creation of various components for data entities defined within our project, streamlining the development process and ensuring consistency across database operations, API endpoints, and more.

## How It Works
The generator utilizes entity definitions from `data/entities/metadata.go`, where each entity is meticulously defined to guide the generation of repositories, configurations, operations, foreign keys, controllers, and routes.

### Key Components
- **Entities**: Central data models defined in `metadata.go`.
- **Repositories**: Facilitates all database operations.
- **Config Functions**: Define entity-specific configurations for database interactions.
- **Operations**: Include hooks for custom logic insertion around CRUD operations.
- **Foreign Keys**: Define relationships between entities based on `metadata.go`.
- **Controllers & Routes**: Automatically generated for handling HTTP requests.

### Customization
To customize, developers can rename generated controllers from `.gen.go` to `.go` to prevent overwrites during subsequent generations.

### The `MetaData` Struct
Outlined in `metadata.go`, it includes crucial information for code generation, such as entity names, database paths, and relationship mappings.

## Getting Started
Update `metadata.go` with your entity definitions, run the generator script, and customize the generated code by editing the `.go` files as needed.
