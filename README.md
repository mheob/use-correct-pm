# Check for using the correct package manager

A simple check of the usage of the correct package manager.

## Install

The local installation is optional. It is only needed if you want to have the package inside your project. This could improve performance, for example.

### With NPM

```sh
npm install -D use-correct-pm
```

### With YARN

```sh
yarn add -D use-correct-pm
```

### With PNPM

```sh
pnpm add -D use-correct-pm
```

## Usage

### As script in `package.json`

Use one of the examples.

```json
"scripts": {
	"use-correct-pm": "npx use-correct-pm npm",        // NPM
	"use-correct-pm": "yarn dlx use-correct-pm yarn",  // YARN
	"use-correct-pm": "yarn exec use-correct-pm yarn", // YARN - locally installed
	"use-correct-pm": "pnpm dlx use-correct-pm",       // PNPM
	"use-correct-pm": "pnpm dlx use-correct-pm pnpm",  // PNPM - alternative
	"use-correct-pm": "pnpm exec use-correct-pm",      // PNPM - locally installed
	"use-correct-pm": "pnpm exec use-correct-pm pnpm", // PNPM - alternative locally installed
},
```

### As shell script

#### Default (same as `PNPM`)

Use `npx` if you use `NPM` as package manager.\
Otherwise use `exec` if you have installed the package in your project, otherwise use `dlx`.

```sh
npx use-correct-pm
```

#### `NPM`

```sh
npx use-correct-pm npm
```

#### `YARN`

```sh
yarn dlx use-correct-pm yarn
```

#### `PNPM`

```sh
pnpm dlx use-correct-pm pnpm
```

### As git hook

For example used in combination with [husky](https://typicode.github.io/husky/).

```sh
#!/usr/bin/env sh
. "$(dirname -- "$0")/_/husky.sh"

[ -n "$CI" ] && exit 0

pnpm dlx use-correct-pm
```
