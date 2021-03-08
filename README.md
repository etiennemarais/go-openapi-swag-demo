# go-openapi-swag-demo

![Api Docs](https://github.com/overhq/etiennemarais/go-openapi-swag-demo/Swagger/badge.svg)

## Summary

Trying to write documentation by hand is a tedious ask. This optimizes writing documentation for the writer as close to the code as possible by using annotations
to auto generate api docs.

Ths issue with most outputs are that it would just live in a file somewhere and that sucks. By leveraging github actions and a little magic, we can send the generated file to a [repo that auto builds and deploys](https://github.com/etiennemarais/go-openapi-swag-ui-demo) the documentation making it more up to date.
