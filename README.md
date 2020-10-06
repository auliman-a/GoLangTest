Site will be hosted on http://localhost:8080/

Graphql schema file is included in the repository (/graph/schema.graphqls)
Example Scenarios mentioned is availaible in Unit Test

Assumption: 
- Promotion for Macbook Pro will not be applied if there is not enough Raspberry Pi
- Promotion will be applied in Checkout Process
- Checkout Process will directly substract necessary quantity in Inventory

Script to run Unit Test using Testify library: "go test .\...\ -coverprofile=output_test"