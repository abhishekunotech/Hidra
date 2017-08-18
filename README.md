# Hydra - Working Title for Middleware in Golang
### Overview

---

This is a middleware application that allows you to share APIs from individual components of the Felicity Platform. It acts as a layer of abstraction between the underlying components and the User Interface. 

As this type of an application is by design meant to be extended as more applications integrate with it. The design of the applicaiton is modular by design where each layer has separation of concerns. This allows for easy extensibility of each of the layers.


### Documentation and Reference

---

API Documentation : TBD

Docker Image : TBD

### Application Architecture

---

#### Overview
The broad underlying design theme of this application is a modular approach where each layer communicates with the layers below
and above it in a structured format and clearly stated separation of concenrs. This allows for easily adding modules on any of the layers and extending support for different types of technologies.

#### Server Layer

#### Request Response Abstract
This layer serves 2 purposes.

1) As a standardization layer between the Different servers and the layer below so the servers and the implementation
don't need to be tightly coupled and the code in the layers below can be reused.

2) This is a layer through which all traffic is expected to flow through. Both the request and response will flow through this layer.
This allows for(All are planned features) :
- All requests wide security validation & dropping requests that are not authorised.
- Hooks and logic to integrate logging systems, calls to monitoring and alarm systems.
- Transformation or enriching of all responses before they are returned back to the caller.

This layer accepts a standard Request object and returns a fixed Response object to the Server Layer


#### Application Custom Logic
Based on the application and version of the calling application the request is routed to the right method. 

This is the major part that would need to be extended to support future applications.


#### System Abstract Layer
The purpose of this layer is to have a standard request and response format for all types of databases. This layer simplifies the creation of application drivers. This layer exists to keep things modular within the application and also simplifies the creation of driver layer to support more databases.


