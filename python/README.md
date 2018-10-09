# Hello Goose - Python

The php version of the application has no external dependencies and requires no configuration. Just run `cf push` to use.

If you are deploying in an environment where the buildpack will not be able to fetch libraries from the cheese shop, then you will have to vendor the dependencies. See [the docs](https://docs.cloudfoundry.org/buildpacks/python/index.html#vendoring).
