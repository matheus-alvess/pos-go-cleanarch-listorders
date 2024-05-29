package graphql

const Schema = `
schema {
    query: Query
    mutation: Mutation
}

type Query {
    listOrders: [Order!]!
}

type Mutation {
    createOrder(price: Float!, tax: Float!): Order!
}

type Order {
    id: Int!
    price: Float!
    tax: Float!
}
`
