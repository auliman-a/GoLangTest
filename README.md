Site will be hosted on http://localhost:8080/

Graphql schema file is included in the repository (/graph/schema.graphqls)

Example Scenarios mentioned is availaible in Unit Test

Assumption: 
- Promotion for Macbook Pro will not be applied if there is not enough Raspberry Pi in Inventory
- Promotion will be applied in Checkout Process
- Checkout Process will directly substract necessary quantity in Inventory

Script to run Unit Test using Testify library: 
```bash
go test .\...\ -coverprofile=output_test
```
Graphql example:
- Add to Cart

```bash
mutation {
  addToCart(input:{sku: "43N23P", quantity: 1})
}
```

- Checkout
```bash
mutation {
  checkout{
    totalPrice
    cartItems{
      sku
      name
      quantity
      amount
    }
  }
}
```


Additional graphql to check inventory data and current cart content
- Get Inventory Data
```bash
query {
  getInventoryData{
    sku
    name
    price
    quantity
  }
}
```

- Get Cart data
```bash
query {
  getCartItem {
    totalPrice
    cartItems {
      sku
      quantity
      name
      amount
    }
  }
}

```