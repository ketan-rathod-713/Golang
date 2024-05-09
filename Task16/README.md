# Eccommerce application using graphql

- Product & Category Listing

## Queries Examples

```
query example {
  products: GetProducts{
    Id,
    Name, 
    Description,
    Price,
    Quantity,
    Category {
      Id,
      Name
    }
  },
  categories : getCategories{
    Id,
    Name,
    Products {
      Id,
      Name
    }
  }
}
```

## Mutation Examples

```
// create category
mutation createCategory {
  CreateCategory(name: "electronic-items"){
    Id,
    Name
  }
}

// create product
mutation createProduct{
  CreateProduct(name: "Notebook", price:500, quantity:10, description:"this is a product",Category: "663b89651f437c9bcf6d3862"){
    Id,
    Name
  }
}
```