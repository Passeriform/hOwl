# Physical Constructs

## Keyword

Any space separated phrase

## Hierarchy

Activator: `<Enter-Tab>`

The hierarchy root becomes heading/root-node.

## Sibling

Activator: `<Enter>`

## Points

Activator: `<Enter>`

Special kind of sibling with no leaf properties.

## Comparison

Compare siblings for same stuff. Append rest to individual points.

## Definition

Activator: `-`

## Implication

Activator: `=>`

## Preview Definition

Activator: `()`

## Flow

Activator: `->`

## Variables

Activator: `create`

Assume every variable exists.
Create scoped variables.

## Loop

Activator: `such that`

Child hierarchy runs for loop.

## Iterator

Activator: `iterate`

## Lists

Activator: `[]`

## Checkboxes

Activator: `[X]` `[ ]`

## Question

Activator: `?`

# Arithmetic

# Modes

-   Docs
-   Art
-   Math
-   Algo

# For my sanity: context switching from rust.

-   cmd has only binary representations.
-   pkg and internal are made imagining the following scenario:

MyBigProject
├── MySmallProjectX
│   ├── pkg
│   │   └── SharedPackageX
│   └── internal
│       └── InternalPackageX
└── MySmallProjectY
    ├── pkg
    │   └── SharedPackageX
    └── internal
        └── InternalPackageY

-   Internal is sharable with the current package only.
-   Pkg is sharable across projects.
-   Heavy confusion between package and modules.
-   Packages are mostly main file and supplement files like a small CPP project.
-   Modules are comprehensive and contextual libraries like a CPP project with a makefile or maybe a Rust project with cargo support.
-   Modules are referred and much better explored if searched for VGO
-   Don't commit vendor unless its integral to the project and expected to be discontinued. Basically its a licensing/archiving system (hence the name).
-   Build the files using the command: go build ./...
-   Run tests using command: go test ./...
-   Currently the files aren't recognized as part of a module by go-tool, so must explicitly specify the files being used.
