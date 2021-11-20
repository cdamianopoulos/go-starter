# Go utility packages `utl`
This repository contains utility packages that can be reused in other services and projects. It aims to provide robust, well tested & production proven code.

## Re-usability
By extracting often repeated & useful code amongst our repositories into a central location we can take advantage of:
* A bug found and fixed in one repository can be quickly applied to other repositories by updating the Go module dependencies.
* A reference for existing code. e.g.: Database connections, email validation, etc.

## Keep It Simple Stupid
* Don't add additional logic to existing functions or methods if it isn't required by all descendants.
  * Instead, **make a copy** and add the additional logic that's required to it. It's better to have a variety of available utilities rather than impacting CPU and/or memory consumption in production.
* Think carefully about function/method signatures. Is it consistent with existing APIs or the Go builtin library?
  * Changing a function/method signature could have a significant impact on technical dept and available developer time.


## Before pushing changes
* Format Go code using [`gofumpt`](https://github.com/mvdan/gofumpt)
* Check for spelling mistakes grammatical errors in function names, returned values, etc.
  * Goland has a code inspection tool available in the main menu. `Code` > `Inspect Code...` > `Whole project`
* Lint the code using [`golangci-lint`](https://github.com/golangci/golangci-lint) to discover any additional issues.

