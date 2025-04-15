# recaptcha-module-level-poc

## Introduction
- This is a POC module which will show how to add a recaptcha checking code snippet at any method in a proto file.

- In this POC, we used Modular Monolithic architecture. It Consists of, 
  - A single HTTP server.
  - A middleware, which use to set google-recaptcha-response from request header to ongoing context.
  - recaptcha_module, this module will contain code to validate recaptcha.
  - Two Modules - dummy_1 and dummy2 (containing two endpoints each, on which we have option to enable recaptcha)

## Prerequisite
- Setup ENV's for recaptcha config (default value already added for POC)
  - recaptcha.recaptchaskip [default value: false]
  - recaptcha.recaptchasecretkey [default value: 6LeIxAcTAAAAAGG-vFI1TnRWxMZNFuojJ4WifJWe] // test key
 
- HTTP Request should contain request header with key - **g-recaptcha-response** (containing response from frontend application)

## Steps to Add Recaptcha option within each module

- Include this Proto file in your ENV: https://github.com/jatin-dev-appointy/recaptcha-go-plugin/blob/main/recaptcha/recaptcha.proto
- Import this proto in you Module's Proto and add below code snippet.

For example
  ```
  rpc GetDummy2Method2(GetDummy2Method2Request) returns (Dummy2Method2Object) {
    option (google.api.http) = {
      get: "/v1/dummy-2/method-2/{id=**}"
    };
    option (recaptcha.require_recaptcha) = true;
  }
  ```

- Add this custom plugin in your compilation method (https://github.com/jatin-dev-appointy/recaptcha-go-plugin)
  - Generate its Binary with name: **protoc-gen-recaptcha** and add in ENV
  - or download (for M chip): https://github.com/jatin-dev-appointy/recaptcha-go-plugin/blob/main/protoc-gen-recaptcha
 
- Recompile Proto using **Trishul** or **GenX** compile option


## Steps to Run POC in local
- Clone this Module.
- Make required changes in each Module's Proto file and Recompile.
- Run **main.go** as a Go Package.
