# GOREST

This is a simple REST HTTP VC Framework written in GO

Design paradigm is similar to Django/Rails

* Routes are defined in the go.routes file
* Controllers and their methods are defiend in controllers directory
* Views are defiend in the views directory

To compile, run `` go build goserver ``

### TODO

* Redesign View rendering
* Add support for different types of content responses (JSON, Javascript...)
* Passing Parameters (Both URL and Headers)
* Session Support
* Add basic ORM Support
* Add Unit Tests