# app-engine-diving
To set up dev env [follow this guide](https://cloud.google.com/appengine/docs/standard/go/download)  
Testing is done using the [local dev server](https://cloud.google.com/appengine/docs/standard/go/tools/using-local-server)

To test:
1. Navigate to the src directory
2. run `$ dev_appserver.py app.yaml`

The goal of this repo is to build the foundation of hemres-appengine:
* Load in XML export file from MacDive and Subsurface
* File loading can be done logged in or anonymous
* When logged in a user can see all of there dives
  * This will only show summary data such as
  * Date and time of dive
  * Duration of dive
  * Location of dive
  * Max and Min depth
  * Max and Min temp
  * The list view will only have the, "name", time, and duration of the dive
