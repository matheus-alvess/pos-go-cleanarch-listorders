package graphql

const Schema = `
schema {
    query: Query
}

type Query {
    listOrders: [Order!]!
}

type Order {
    id: Int!
    price: Float!
    tax: Float!
}
`
