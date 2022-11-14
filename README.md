# Check for using the correct package manager

A simple check of the usage of the correct package manager.

## Usage

No installation is needed. Just run with `npx`.

### As script in `package.json`

Use one of the examples.

```jsonc
"scripts": {
	"use-correct-pm": "npx use-correct-pm",       // PNPM
	"use-correct-pm": "npx use-correct-pm pnpm",  // PNPM
	"use-correct-pm": "npx use-correct-pm npm",   // NPM
	"use-correct-pm": "npx use-correct-pm yarn",  // YARN
},
```

### As shell script

#### Default (same as `PNPM`)

Use `npx` if you use `NPM` as package manager.\
Otherwise use `exec` if you have installed the package in your project, otherwise use `dlx`.

#### Check for `PNPM`

```sh
npx use-correct-pm
# or
npx use-correct-pm pnpm
```

#### Check for `NPM`

```sh
npx use-correct-pm npm
```

#### Check for `YARN`

```sh
npx use-correct-pm yarn
```

### As git hook

For example used in combination with [husky](https://typicode.github.io/husky/).

```sh
#!/usr/bin/env sh
. "$(dirname -- "$0")/_/husky.sh"

[ -n "$CI" ] && exit 0

npx use-correct-pm
```
