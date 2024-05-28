package graphql

const Schema = `
schema {
    query: Query
}

type Query {
    listOrders: [Order!]!
}

type Order {
    id: ID!
    price: Float!
    tax: Float!
}
`
